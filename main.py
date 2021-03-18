import xml.etree.ElementTree as xmlparser
import json
import requests
import dataclasses
import inquirer
import clipboard
import os
import yaml
import click
import sys


@dataclasses.dataclass
class MetadataVersions:
    latest: str
    release: str
    version: list[str]
    last_updated: str


@dataclasses.dataclass
class Metadata:
    group_id: str
    artifact_id: str
    versioning: MetadataVersions


@dataclasses.dataclass
class Dependency:
    content: str
    type: str = "implementation"


@dataclasses.dataclass
class Content:
    name: str
    document: str
    java: list[Dependency]
    kotlin: list[Dependency]


@dataclasses.dataclass
class Section:
    name: str
    source: str
    contents: dict[str, Content]


def parse_metadata(metadata: str) -> Metadata:
    element = xmlparser.fromstring(metadata)

    group_id = element.find("groupId").text
    artifact_id = element.find("artifactId").text

    versioning = element.find("versioning")
    latest = versioning.find("latest").text
    release = versioning.find("release").text

    versions = []
    for v in versioning.iter("version"):
        versions.append(v.text)

    last_updated = versioning.find("lastUpdated").text

    return Metadata(
        group_id=group_id,
        artifact_id=artifact_id,
        versioning=MetadataVersions(
            latest=latest, release=release, version=versions, last_updated=last_updated
        ),
    )


def fetch_metadata(group_id: str, artifact_id: str, source: str):
    sources = load_sources()
    source_url = sources[source] if source in sources else source
    url = build_metadata_url(source_url, group_id, artifact_id)
    response = requests.get(url)

    if response.status_code == 200:
        return parse_metadata(response.text)
    else:
        return None


def build_metadata_url(source: str, group_id: str, artifact_id: str) -> str:
    source += "/" + group_id.replace(".", "/")
    source += "/" + artifact_id
    source += "/maven-metadata.xml"
    return source


def load_sources():
    with open("./sources.json") as file:
        return json.load(file)


def load_sections() -> dict[str, Section]:
    contents = read_sections()
    return parse_sections(contents)


def read_sections():
    contents = {}
    for content in os.listdir("./contents"):
        content_path = os.path.join("./contents", content)
        content_name, _ = os.path.splitext(os.path.basename(content))

        with open(content_path) as content_file:
            content_yaml = yaml.load(content_file, Loader=yaml.BaseLoader)
            contents[content_name] = content_yaml

    return contents


def parse_sections(contents: dict) -> dict[str, Section]:
    sections = {}
    for content in contents:
        section = contents[content]
        section_name = section["name"]
        section_source = section["source"]
        section_contents = {}

        for item in section["content"]:
            name = item["name"]
            document = item["document"]
            java = parse_dependencies(item["java"])
            kotlin = parse_dependencies(item["kotlin"]) if "kotlin" in item else []

            section_contents[name] = Content(
                name=name, document=document, java=java, kotlin=kotlin
            )

        sections[content] = Section(
            name=section_name, source=section_source, contents=section_contents
        )

    return sections


def parse_dependencies(items: list) -> list[Dependency]:
    dependencies = []
    for dep in items:
        dependency_object = None
        if isinstance(dep, dict):
            content = dep["content"]
            if "type" in dep:
                type = dep["type"]
                dependency_object = Dependency(content, type)
            else:
                dependency_object = Dependency(content)
        else:
            dependency_object = Dependency(dep)
        dependencies.append(dependency_object)
    return dependencies


def generate_dependencies(
    sections: dict[str, Section],
    project_language: str,
    gradle_language: str,
    select_items: list[tuple[str, str]],
) -> list[str]:
    output = []

    for select_item in select_items:
        section_key, group = select_item

        if section_key not in sections or group not in sections[section_key].contents:
            continue

        section = sections[section_key]
        group_content = sections[section_key].contents[group]

        if project_language == "kotlin":
            dependencies = (
                group_content.kotlin if group_content.kotlin else group_content.java
            )
        else:
            dependencies = group_content.java

        for dependency in dependencies:
            group_id, artifact_id = dependency.content.split(":")

            metadata = fetch_metadata(group_id, artifact_id, section.source)
            if metadata:
                dependency_str = dependency_noun(
                    gradle_language,
                    dependency.type,
                    dependency_notation(
                        group_id, artifact_id, metadata.versioning.latest
                    ),
                )
                output.append(dependency_str)

    return output


def dependency_notation(group_id: str, artifact_id: str, version: str) -> str:
    return f"{group_id}:{artifact_id}:{version}"


def dependency_noun(language: str, prefix: str, dependency_notation: str) -> str:
    if language == "kotlin":
        return f'{prefix}("{dependency_notation}")'
    elif language == "groovy":
        return f'{prefix} "{dependency_notation}"'


def copy_clipboard(content: str):
    clipboard.copy(content)


@click.group()
def cli():
    pass


@click.command()
@click.argument("dependencies", nargs=-1)
@click.option("-l", "--lang", "lang", default="kotlin", type=str)
@click.option("-g", "--gradle", "gradle", default="groovy", type=str)
def get(dependencies: list[str], lang, gradle):
    sections = load_sections()

    select_items = []
    for dependency in dependencies:
        section, group = dependency.split(":")
        select_items.append((section, group))

    output = generate_dependencies(sections, lang.lower(), gradle.lower(), select_items)

    print("\n".join(output))


@click.command("ui")
def open_ui():
    click.echo("Unsupported operation")


def interactive():
    sections = load_sections()

    questions = [
        inquirer.List(
            "project_language",
            message="Language",
            choices=["Kotlin", "Java"],
            default="Kotlin",
        ),
        inquirer.List(
            "gradle_language",
            message="Gradle Language",
            choices=["Groovy", "Kotlin"],
            default="Groovy",
        ),
    ]

    answers = inquirer.prompt(questions)

    project_language = answers["project_language"].lower()
    gradle_language = answers["gradle_language"].lower()

    questions = []

    for key, section in sections.items():
        question = inquirer.Checkbox(
            name=key,
            message=f"{section.name} (Press spacebar to select)",
            choices=list(map(lambda content: content.name, section.contents.values())),
        )

        questions.append(question)

    answers = inquirer.prompt(questions)

    select_items = []
    for section, items in answers.items():
        for item in items:
            select_items.append((section, item))

    output = generate_dependencies(
        sections, project_language, gradle_language, select_items
    )
    output = "\n".join(output)
    print(output)
    copy_clipboard(output)
    print("\nCopied to clipboard!")


if __name__ == "__main__":
    if len(sys.argv) >= 2:
        cli.add_command(get)
        cli.add_command(open_ui)
        cli()
    else:
        interactive()