import xml.etree.ElementTree as xmlparser
import json
import requests
import dataclasses
import inquirer
import clipboard

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
            latest=latest,
            release=release,
            version=versions,
            last_updated=last_updated
        )
    )


def fetch_metadata(group_id: str, artifact_id: str, source: str):
    sources = load_sources()
    url = build_metadata_url(sources[source], group_id, artifact_id)
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

def load_contents():
    with open("./contents.json") as file:
        return json.load(file)

def dependency_notation(group_id: str, artifact_id: str, version: str) -> str:
    return f"{group_id}:{artifact_id}:{version}"

def dependency_noun(language: str, dependency_notation) -> str:
    if language == "kotlin":
        return f"implementation(\"{dependency_notation}\")"
    elif language == "groovy":
        return f"implementation \"{dependency_notation}\""

if __name__ == "__main__":
    questions = [
        inquirer.List(
            'project_language',
            message="Language",
            choices=["Kotlin", "Java"],
            default="Kotlin"
        ),
        inquirer.List(
            'gradle_language',
            message="Gradle Language",
            choices=["Groovy", "Kotlin"],
            default="Groovy"
        ),
    ]

    answers = inquirer.prompt(questions)

    project_language = answers["project_language"].lower()
    gradle_language = answers["gradle_language"].lower()

    questions = []
    contents = load_contents()

    for content in contents:
        choices = contents[content]

        question = inquirer.Checkbox(
            name=content,
            message=f"{content} (Press spacebar to select)",
            choices=choices
        )
        questions.append(question)
    
    answers = inquirer.prompt(questions)

    output = []
    for content in contents:
        for item in answers[content]:
            source = contents[content][item]["source"]
            for dependency in contents[content][item][project_language]:
                group_id = dependency["groupId"]
                artifact_id = dependency["artifactId"]

                metadata = None
                if "source" in dependency:
                    metadata = fetch_metadata(group_id, artifact_id, dependency["source"])
                else:
                    metadata = fetch_metadata(group_id, artifact_id, source)

                if metadata:
                    dependency_str = dependency_noun(
                        gradle_language,
                        dependency_notation(group_id, artifact_id, metadata.versioning.latest)
                    )

                    output.append(dependency_str)


    clipboard.copy("\n".join(output))
    print("\n".join(output))
    print("\nCopied to clipboard")