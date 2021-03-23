<h1 align="center">Brick</h1>

<!-- <p align="center">
    <img alt="License" src="https://img.shields.io/badge/License-Apache%202.0-blue.svg"/>
</p> -->

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

## How to use

```bash
# Windows
./brick.exe

# Linux
./brick
```

## Subcommands

### `get`

> Get latest library

- Arguments  
    `{section}:{group}` `{section}:{group}` ...

    e.g jetpack:activity

- Options
  - `-l`, `--lang` : Project Language (`kotlin` or `java`), defaults `kotlin`
  - `-g`, `--gradle` : Gradle Language (`groovy` or `kotlin`), default `groovy`
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
