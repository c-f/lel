# LEL

```bash
TLDR; Explode your notes.
A PoC for a better keepnote
```

![LEL - no need for more](./docs/LEL-short.png)

LEL is a visualization layer and helper for relevant documentation and note taking focused on IT operations.

LEL is built with <3 in Go and React by [iNvist](https://twitter.com/invist) .

[Website](https://editor.l3l.lol) |
[Documentation](./docu/guide.md) |
[Installation Guide](./docs/installation.md)

[![GoDoc](https://godoc.org/github.com/c-f/lel?status.svg)](https://godoc.org/github.com/c-f/lel)
[![Build Status](https://api.travis-ci.org/c-f/lel.svg?branch=master)](https://travis-ci.org/c-f/lel)
[![Go Report Card](https://goreportcard.com/badge/github.com/c-f/lel)](https://goreportcard.com/report/github.com/c-f/lel)

## Overview

![STATUS:BETA](https://img.shields.io/badge/Status-beta-blue)

The tool is heavily focused to help structuring relevant infos for Redteam, Pentest, Research, Bugbounty and Ops. Go create your own **explosion**

https://editor.l3l.lol

---

**Installation**:
[here](#install)

**Usage**:
[here](#usage)

**Dependencies**:
[here](docs/dependencies.md).

**Development**:
[Dev-Docs](docs/development.md)

**Bugs**:
[More information](docs/bugs.md)

**Automation**:
[Automation](docs/automation.md)

**API**:
[API](docs/api.md)

**Guide**:
[guide](docs/guide.md)

---

## Install

> LEL can be cross compiled and is therefore available for Windows, Linux

LEL can be downloaded through github releases

### Docker

In this repo you find a `docker-compose.yml`, which can be used to start LEL:

```bash
docker pull invist/lel:v0.0.1
docker-compose up
```

The official Repo for lel is [invist/lel:v0.0.1](https://hub.docker.com/r/invist/lel).

## Usage

> Because sharing is Caring

[More examples](docs/guide.md) can be found in the guide:

```bash
./lel -docu <docupath> -editor /usr/bin/code -addr 127.0.0.1:8888
```

## Metatags

Metatags are the explosion parts of LEL:

More details and documentation can be found [here](./docs/metatags.md)

| Metatag | Required | Short definition                            |
| ------- | -------- | ------------------------------------------- |
| @ref    | -        | Link between entities                       |
| @todo   | -        | Indicates the start of a todo list          |
| @tags   | -        | Give your notes some context and group them |
| @name   | yes      | Id parameter to reference the entity.       |
| @label  | -        | Graph label for node                        |
| @entity | yes      | Specifiy which entity a note should be.     |
| @icon   | -        | Lets you specify your own icon              |
|         | -        |                                             |
| ?       | -        | Do you think we missed one ? -> gh issue <3 |

````
```
@name example.com, 127.0.0.1, HOST.ad.example.com
@entity server
@label example.com
@tags entrypoint, linux, dmz, germany,

# references
@ref (c2)->[:webshell-1]->(this)
@ref (this)->[access]->(proxy)
@ref (this)->[access]->(intra.example.com)
```

# Example.com
> Some information about the host

## Postexploitation

```
@todo post
run +mimikatz +postexploit
collect +creds
deploy webshell

```
````

## FAQ

- **Q:** My External editor is not working ?

  **A** Please specify an external editor with the `-editor` parameter and click the toggleButton in the LeL settings.

- **Q:** Can I use LEL for private/corporate projects ?

  **A:** Absolutely

- **Q:** What does LEL stand for ?

  **A:** Of course LEL stands for the [_Lower explosion limit_](https://en.wikipedia.org/wiki/Flammability_limit).

  _The lowest concentration (percentage) of a gas or a vapor in air capable of producing a flash of fire in the presence of an ignition source (arc, flame, heat)_

- **Q:** But why should i use LEL when i have xyz

  **A** You can choose whatever editor you like. In fact LEL is supporting external editors :) Try it out LEL can coexist between sublime, VS-code, Seafile, Gitlab

- **Q:** Why a Logo / Website / Stuff ?

  **A** Just for fun ¯\_(ツ)\_/¯ - giving something back to the community: https://editor.l3l.lol/

- **Q:** But XYZ is missing !

  **A** No problem - write an issue and we add it to the feature board.
