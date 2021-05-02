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

<details>
  <summary>Jetpack</summary>

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
</details>

<details>
  <summary>DI</summary>

  - dagger2
  - dagger-hilt
  - koin
</details>

<details>
  <summary>Networking</summary>

  - retrofit
  - okhttp
  - okhttp-bom
  - okhttp-mockwebserver
  - fast-android-networking
  - volley
  - cronet
</details>

<details>
  <summary>Firebase</summary>

  - Firebase Android BoM (Bill of Materials)
  - AdMob
  - Analytics
  - App Indexing
  - Authentication
  - Cloud Firestore
  - Cloud Functions for Firebase Client SDK
  - Cloud Messaging
  - Cloud Storage
  - Crashlytics
  - Dynamic Links
  - Firebase ML Model Downloader API
  - In-App Messaging
  - In-App Messaging Display
  - Performance Monitoring
  - Realtime Database
  - Remote Config
</details>

<details>
  <summary>Coroutines</summary>

  - kotlinx-coroutines-bom
  - kotlinx-coroutines-core
  - kotlinx-coroutines-core-common
  - kotlinx-coroutines-test
  - kotlinx-coroutines-debug
  - kotlinx-coroutines-play-services
</details>

<details>
  <summary>Rx</summary>

  - rxjava3
  - rxjava3-snapshot
  - rxkotlin3
  - rxandroid
</details>

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
  invalid: require 'source' field in (root)
  invalid: require 'name' field in (root)
  invalid: require 'content' field in dagger-hilt (java) (line: 6, col: 9)
  invalid: require 'java' or 'kotlin' field in dagger2 (line: 16, col: 5)
  invalid: require 'document' field in koin (line: 19, col: 5)
  invalid: require 'name' field in (line: 27, col: 5)
  Error: 6 issues found
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