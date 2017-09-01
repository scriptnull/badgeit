Thanks for your interest in contributing to badgeit.

`badgeit` is planned and maintained with [github issue board](https://github.com/scriptnull/badgeit/issues). Please feel free to [check it](https://github.com/scriptnull/badgeit/issues). It can be used to 
1. Report and know bugs
1. Suggesting and discovering new enhancements to work on
1. What's next for badgeit.

## Workflow

#### Setup 
```bash
git clone git@github.com:scriptnull/badgeit.git
cd badgeit
make init
go get -v ./...
```

#### Build
```bash
make build
```

#### Run tests
```bash
make test
make test-contracts # test all the contracts 
make test-formatters # test all the formatters
```

## Guidelines
1. Please comment on an issue, if you are willing to work on it.
2. If you don't know where to start, check the issues with `beginner-friendly` label.
3. Feel free to reach out, if you are struck with something. There is always somebody to help and something to learn.

## Reporting bugs
Please search the issue board if there it is already reported. Provide detailed incident history like screenshots, input data etc.

## Adding Contracts
> These issues have `new contract` label

Contracts are conditions on which a one or more badges can be generated.

For example, `github` checks if the given repository is a git repository and has remotes pointing to github.com. If the contract is satisfied, then it returns badges like [![github closed issues](https://img.shields.io/github/issues-closed/atom/atom.svg)](https://github.com/atom/atom) [![github open pr](https://img.shields.io/github/issues-pr/atom/atom.svg)](https://github.com/atom/atom) [![github closed pr](https://img.shields.io/github/issues-pr-closed/atom/atom.svg)](https://github.com/atom/atom) [![github contributors](https://img.shields.io/github/contributors/atom/atom.svg)](https://github.com/atom/atom)

#### Steps
1. Add `contract_NAME.go` and `contract_NAME_test.go` in [contracts package](https://github.com/scriptnull/badgeit/tree/master/contracts).
1. Refer [similar contracts](https://github.com/scriptnull/badgeit/blob/master/contracts/contract_npm.go) to model code.
1. Make sure to write required tests. You can use `samples` folder to hold any test sample data for your contract.
1. Use `Makefile` to define task for running tests, generating the badges for the contract and initializing the sample data (if needed).
1. Add sample badges in README that the contract can output.

## Adding Formatters
> These issues have `new formatter` label

Formatters define the way the badges are presented.

Steps
1. Add `formatter_NAME.go` and `formatter_NAME_test.go` in [fotmatters package](https://github.com/scriptnull/badgeit/tree/master/formatters).
1. Refer [similar formatters](https://github.com/scriptnull/badgeit/blob/master/formatters/formatter_all.go) to model code.
1. Add short description in [Formatters section of README](https://github.com/scriptnull/badgeit#formatters)

## Code of Conduct
Please adhere to the [Code of Conduct](https://github.com/scriptnull/badgeit/blob/master/CODE_OF_CONDUCT.md).