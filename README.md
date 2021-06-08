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
brick.exe [command]

# Linux
brick [command]
```

> ⚠ Require run `update` command before first run.

## Commands

To see a list of commands, run:

```
brick help
```

The current commands are:

```
bucket      Management buckets
doc         Open library document in browser
get         Get latest library
help        Help about any command
list        Print supported libraries
update      Update bucket caches
valid       Validate configuration file
```

## Examples

- Get latest library

```bash
$ ./brick get jetpack:activity jetpack:appcompat
implementation "androidx.appcompat:appcompat:1.3.0-beta01"
implementation "androidx.activity:activity-ktx:1.3.0-alpha04"
  
$ ./brick get jetpack:appcompat --gradle=kotlin 
implementation("androidx.appcompat:appcompat:1.3.0-beta01")
```

- Open library document in browser

```bash
$ ./brick doc jetpack:activity
Opening https://developer.android.com/jetpack/androidx/releases/activity in your browser.
```

- Print supported libraries

```bash
$ ./brick list
└── jetpack
  ├── activity
  ├── appcompat
  ├── camera
  ├── compose
  ├── fragment
      ...

$ ./brick list --section jetpack
jetpack
├── activity
├── appcompat
├── camera
├── compose
├── fragment
    ...
```

- Update bucket caches

```bash
$ ./brick update
Updating namhyun-gu:brick...
```

- Validate configuration file

```bash
$ ./brick valid test.yaml
invalid: require 'source' field in (root)
invalid: require 'name' field in (root)
invalid: require 'content' field in dagger-hilt (java) (line: 6, col: 9)
invalid: require 'java' or 'kotlin' field in dagger2 (line: 16, col: 5)
invalid: require 'document' field in koin (line: 19, col: 5)
invalid: unable 'io.insert-koin:koin-android' (source: null) in koin (java) (line: 21, col: 9)
invalid: require 'name' field in (line: 27, col: 5)
Error: 7 issues found
```

- Management buckets

```bash
# Add bucket
$ ./brick bucket add namhyun-gu:brick@main
Added namhyun-gu:brick to bucket.

# Add bucket with Path
$ ./brick bucket add namhyun-gu:brick@main --path data/
Added namhyun-gu:brick to bucket.

# Remove bucket
$ ./brick bucket remove namhyun-gu:brick
Removed namhyun-gu:brick to bucket.

# Print buckets
$ ./brick bucket list
namhyun-gu:brick@main data/
```

## License

```
Copyright 2021 Namhyun, Gu

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```