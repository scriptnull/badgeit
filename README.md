<h1 align="center">
<img width="400" src="https://raw.githubusercontent.com/scriptnull/badgeit/master/art.png" />
</h1>

[WIP] Hassle-free badges for your READMEs.

## Install


## Usage
```bash
$ badgeit path-to-source-repo
badge-1
badge-2
badge-n
```

> If path is not specified, current working directory will be default path.

## Formatters
Use `-f` options to send in the available formatting options.

- all: all the badges
- min: 7 most favourable badges
- cat: category wise badges

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