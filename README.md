<p align="center"><img width=50% src="https://github.com/veskoy/gomas/blob/master/media/Logo.png"></p>

<p align="center">
<a href="https://travis-ci.org/veskoy/gomas"><img src="https://travis-ci.org/veskoy/gomas.svg?branch=master" alt="Build Status"></a>
<a href="https://goreportcard.com/report/github.com/veskoy/gomas"><img src="https://goreportcard.com/badge/github.com/veskoy/gomas" alt="Go Report Card"></a>
<a href="https://github.com/veskoy/gomas/issues"><img src="https://img.shields.io/github/issues/veskoy/gomas.svg" alt="GitHub issues"></a>
<a href="https://raw.githubusercontent.com/veskoy/gomas/master/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="GitHub license"></a>
</p>

## Table Of Contents
* [Table Of Contents](#table-of-contents)
* [Basic Overview](#basic-overview)
* [Installation](#installation)
* [Usage](#usage)
* [Changelog](#changelog)
* [Contributing](#contributing)
* [License](#license)
* [Contributors](#contributors)

## Basic Overview
***Gomas*** is a Master Server written in [Go](https://golang.org/) for some of [Valve](http://www.valvesoftware.com/)'s multiplayer games. It is an implementation of the [Master Server Query Protocol](https://developer.valvesoftware.com/wiki/Master_Server_Query_Protocol). The main objective of Gomas is to allow the community to run 3rd party quality master servers.

## Installation
Gomas hasn't been released yet. This means there are no stable executables that you can use out of the box. Gomas is in its early stage and goes through rapid development. For that reason the only way you can run and try it out is by building it from source (Linux / MacOS):

    git clone https://github.com/veskoy/gomas.git
    cd gomas
    go get -v -t ./...
    go build -o ./gomasd ./cmd/gomasd

    This will make an executable "./gomasd" which you can use to start a Master Server.

## Usage

**Default - the master server will listen on 127.0.0.1:27010:**

    ./gomasd

**Start a master server that will listen on specified ip address and port:**

    ./gomasd -ip=xxx.xxx.xxx.xxx -port=xxxxx

**Truncate/Seed/Reset database on startup:**

    ./gomasd -ip=xxx.xxx.xxx.xxx -port=xxxxx -db=truncate
    ./gomasd -ip=xxx.xxx.xxx.xxx -port=xxxxx -db=seed
    ./gomasd -ip=xxx.xxx.xxx.xxx -port=xxxxx -db=reset

## Changelog
To see what has changed in recent versions of Gomas, see the [CHANGELOG](https://github.com/veskoy/gomas/blob/master/CHANGELOG.md).

## Problems
Please report and follow the resolution of any encountered problems in the [issue tracker](https://github.com/veskoy/gomas/issues).

## Contributing
Thank you for considering contributing to Gomas. Any contributions are always welcomed as long as they follow our [contributing guidelines](https://github.com/veskoy/gomas/blob/master/CONTRIBUTING.md).

## License
Gomas is open-sourced software licensed under the [MIT license](http://opensource.org/licenses/MIT).

## Contributors
* **Author:** [Veselin Stoyanov](https://github.com/veskoy)
