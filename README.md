<h1 align="center">Brick</h1>

<!-- <p align="center">
    <img alt="License" src="https://img.shields.io/badge/License-Apache%202.0-blue.svg"/>
</p> -->

<p align="center">
Compose latest android library dependencies
</p>

## Preview

<p align="center">
    <img src="images/preview.png" width="70%"/>
</p>

## Getting Started

```bash
pip install -r requirements.txt
```

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
python main.py
```

### Subcommands

- ### get
  - Arguments  
    `{section}:{group}` `{section}:{group}` ...
       
    e.g jetpack:Activity  
    
  - Options  
    - `-l`, `--lang` : Project Language (kotlin or java), defaults kotlin
    - `-g`, `--gradle` : Gradle Language (groovy or kotlin), default groovy

  - Examples
    ```bash
    $ python main.py get jetpack:Activity jetpack:Appcompat
    implementation "androidx.appcompat:appcompat:1.3.0-beta01"
    implementation "androidx.activity:activity-ktx:1.3.0-alpha04"

    $ python main.py get jetpack:Appcompat --gradle=kotlin 
    implementation("androidx.appcompat:appcompat:1.3.0-beta01")
    ```