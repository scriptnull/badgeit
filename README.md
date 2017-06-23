<h1 align="center">
<img width="400" src="https://raw.githubusercontent.com/scriptnull/badgeit/master/art.png" />
</h1>

[WIP] Hassle-free badges for your READMEs.

## Install
Download the binary from [latest release](https://github.com/scriptnull/badgeit/releases).

## Usage
```bash
$ badgeit path-to-source-repo
badge-1
badge-2
badge-n
```

> If path is not specified, current working directory will be default path.

## Samples
Here are some samples generated by badgeit, for various kinds of projects.
### npm
[![npm weekly downloads](https://img.shields.io/npm/dw/express.svg)](https://npmjs.com/package/express) [![npm monthly downloads](https://img.shields.io/npm/dm/express.svg)](https://npmjs.com/package/express) [![npm yearly downloads](https://img.shields.io/npm/dy/express.svg)](https://npmjs.com/package/express) [![npm total downloads](https://img.shields.io/npm/dt/express.svg)](https://npmjs.com/package/express) [![npm version](https://img.shields.io/npm/v/express.svg)](https://npmjs.com/package/express) [![npm next version](https://img.shields.io/npm/v/express/next.svg)](https://npmjs.com/package/express) [![npm canary version](https://img.shields.io/npm/v/express/canary.svg)](https://npmjs.com/package/express) [![license badge](https://img.shields.io/npm/l/express.svg)](https://npmjs.com/package/express) [![snyk - known vulnerabilities](https://snyk.io/test/npm/express/badge.svg)](https://snyk.io/test/npm/express)

### github
[![github all releases](https://img.shields.io/github/downloads/atom/atom/total.svg)](https://github.com/atom/atom) [![github latest release](https://img.shields.io/github/downloads/atom/atom/latest/total.svg)](https://github.com/atom/atom) [![github tag](https://img.shields.io/github/tag/atom/atom.svg)](https://github.com/atom/atom) [![github release](https://img.shields.io/github/release/atom/atom.svg)](https://github.com/atom/atom) [![github pre release](https://img.shields.io/github/release/atom/atom/all.svg)](https://github.com/atom/atom) [![github fork](https://img.shields.io/github/forks/atom/atom.svg?style=social&label=Fork)](https://github.com/atom/atom) [![github stars](https://img.shields.io/github/stars/atom/atom.svg?style=social&label=Star)](https://github.com/atom/atom) [![github watchers](https://img.shields.io/github/watchers/atom/atom.svg?style=social&label=Watch)](https://github.com/atom/atom) [![github open issues](https://img.shields.io/github/issues/atom/atom.svg)](https://github.com/atom/atom) [![github closed issues](https://img.shields.io/github/issues-closed/atom/atom.svg)](https://github.com/atom/atom) [![github open pr](https://img.shields.io/github/issues-pr/atom/atom.svg)](https://github.com/atom/atom) [![github closed pr](https://img.shields.io/github/issues-pr-closed/atom/atom.svg)](https://github.com/atom/atom) [![github contributors](https://img.shields.io/github/contributors/atom/atom.svg)](https://github.com/atom/atom) [![github license](https://img.shields.io/github/license/atom/atom.svg)](https://github.com/atom/atom)

### gitter
[![gitter chat room](https://badges.gitter.im/scriptnull/badgeit.svg)](https://gitter.im/scriptnull/badgeit)

## Formatters
Use `-f` options to send in the available formatting options.

- all: all the badges
- min: 7 most favourable badges [todo](https://github.com/scriptnull/badgeit/issues/9)
- cat: category wise badges [todo](https://github.com/scriptnull/badgeit/issues/10)

> If not specified, `all` is the default formatter

## Delimit
Use `-d` to mention the character sequence between the badges.

```bash
$ badgeit -d "\n"
badge-1
badge-2
badge-n
```

> If not specified, blank space will be the default

## Style
Use `-s` to mention the style of the badges generated with [shields.io](http://shields.io/).

```bash
$ badgeit -s flat
```

Available styles are
- `plastic`: ![plastic](https://img.shields.io/badge/style-plastic-green.svg?style=plastic)
- `flat`: ![flat](https://img.shields.io/badge/style-flat-green.svg?style=flat)
- `flat-square`:  ![flat-square](https://img.shields.io/badge/style-flat--squared-green.svg?style=flat-square)
- `social`: ![social](https://img.shields.io/badge/style-social-green.svg?style=social)

## Contribute
Please refer [Contributing section](https://github.com/scriptnull/badgeit/blob/master/CONTRIBUTING.md) for guidelines on contributing.