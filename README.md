<h1 align="center">Brick</h1>

<p align="center">
    <img alt="Build Status" src="https://github.com/namhyun-gu/brick/actions/workflows/build.yml/badge.svg"/>
    <img alt="GitHub release" src="https://img.shields.io/github/v/release/namhyun-gu/brick">
    <!-- <img alt="License" src="https://img.shields.io/badge/License-Apache%202.0-blue.svg"/> -->
</p>

<p align="center">
Compose latest android dependencies
</p>

## Preview

<p align="center">
    <img src="images/preview.png" width="70%"/>
</p>

## Getting Started

Download [latest release](https://github.com/namhyun-gu/brick/releases)

## Supported libraries

- Jetpack
    - activity
    - appcompat
    - camera
    - compose
    - fragment
    - hilt
    - lifecycle
    - material
    - navigation
    - paging
    - room
    - work

- DI
    - dagger2
    - dagger-hilt
    - koin

- Networking
    - retrofit
    - okhttp
    - okhttp-bom
    - okhttp-mockwebserver
    - fast-android-networking
    - volley
    - cronet

- Firebase
    - [All products](https://firebase.google.com/support/release-notes/android)

- Kotlin Coroutines
    - kotlinx-coroutines-bom
    - kotlinx-coroutines-core
    - kotlinx-coroutines-core-common
    - kotlinx-coroutines-test
    - kotlinx-coroutines-debug
    - kotlinx-coroutines-play-services

- Rx
    - rxjava3
    - rxjava3-snapshot
    - rxkotlin3
    - rxandroid

## How to use

```bash
# Windows
./brick.exe [command]

# Linux
./brick [command]
```

> ⚠ Require run `update` command before first run.

## Commands

### `get`

> Get latest library

- Arguments  
  `{section}:{group}` `{section}:{group}` ...

  e.g jetpack:activity

- Options
    - `-l`, `--lang` : Project Language (`kotlin` or `java`), defaults `kotlin`
    - `-g`, `--gradle` : Gradle Language (`groovy` or `kotlin`), defaults `groovy`
    - Examples
      ```bash
      $ ./brick get jetpack:activity jetpack:appcompat
      implementation "androidx.appcompat:appcompat:1.3.0-beta01"
      implementation "androidx.activity:activity-ktx:1.3.0-alpha04"
  
      $ ./brick get jetpack:appcompat --gradle=kotlin 
      implementation("androidx.appcompat:appcompat:1.3.0-beta01")
        ```

### `doc`

> Open document in browser

- Argument  
  `{section}:{group}`

  e.g jetpack:activity

- Examples
  ```bash
  $ ./brick doc jetpack:activity
  Opening https://developer.android.com/jetpack/androidx/releases/activity in your browser.
  ```

### `list`

> Print supported libraries

- Options
    - `-s`, `--section`: Section Name, defaults `""`

- Examples
  ```bash
  $ ./brick list
  └── jetpack
    ├── activity
    ├── appcompat
    ├── camera
    ├── compose
    ├── fragment
    ├── hilt-navigation-compose
    ├── hilt-navigation-fragment
    ├── hilt-workmanager
    ├── lifecycle
    ...

  $ ./brick list --section jetpack
  jetpack
  ├── activity
  ├── appcompat
  ├── camera
  ├── compose
  ├── fragment
  ├── hilt-navigation-compose
  ├── hilt-navigation-fragment
  ├── hilt-workmanager
  ├── lifecycle
  ...
  ```

### `update`

> Update bucket caches

- Examples
  ```bash
  $ ./brick update
  Updating namhyun-gu:brick...
  ```

### `valid`

> Validate `yaml` data file

- Examples
  ```bash
  $ ./brick valid test.yaml
  invalid: require 'source' field
  invalid dependency (index: 0) in group (index: 0): require 'content' field
  invalid group (index: 1): require 'java' or 'kotlin' field
  invalid group (index: 2): require 'document' field
  invalid group (index: 4): require 'name' field
  Error: 5 issues found
  ```

### `bucket`

> Management bucket

- `add`

  > Add new bucket

  - Argument  
    `{owner}:{repo}@{branch}`  

    e.g namhyun-gu:brick@main
  - Options
    - `-p`, `--path`: Bucket default path, defaults `""`
  - Example
    ```bash
    $ ./brick bucket add namhyun-gu:brick@main --path data/
    Added namhyun-gu:brick to bucket.
    ```

- `remove`

  > Remove bucket

  - Argument  
    `{owner}:{repo}`  
  - Example
    ```bash
    $ ./brick bucket remove namhyun-gu:brick
    Removed namhyun-gu:brick to bucket.
    ```

- `list`

  > Print bucket list

  - Example
    ```bash
    $ ./brick bucket list
    namhyun-gu:brick@main data/
    Found 1 buckets.
    ```