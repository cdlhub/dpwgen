# dpwgen

Diceware passphrase generator command line tool written in Go. It generates "random" passphrase from a wordlist file based on EFF article [Deep Dive: EFF's New Wordlists for Random Passphrases](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases).

## Prerequisites

* Go language: Follow installation steps at Golang documentation [Getting started](https://golang.org/doc/install) page.

## Installation

Use Go tool to install `dpwgen` executable on any platform (Linux, macOS, or Windows):

```sh
go get -v github.com/cdlhub/dpwgen
```

Or, build executable from sources:

1. Use `git` to download sources:

    ```sh
    git clone https://github.com/cdlhub/dpwgen.git
    ```

1. Build the package:

    ```sh
    cd dpwgen
    go build
    ```

## Usage

You need a diceware wordlist file as described in [Deep Dive: EFF's New Wordlists for Random Passphrases](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases). Two EFF files are provided in the repository:

* [eff_large_wordlist.txt](eff_large_wordlist.txt).
* [eff_short_wordlist_2_0.txt](eff_short_wordlist_2_0.txt).

The number of dice the tool needs to throw is automatically detected from the file format.

You can then generate a passphrase with the number of words you want (e.g. 6):

```sh
dpwgen -n 6 eff_large_wordlist.txt
```

```txt
whiff surname footgear overfill bust expel
```

Use `-help` to display usage:

```sh
dpwgen -help
```

```txt
Usage:
	dpwgen -version
	dpwgen [ -n N ] <pass-file-name> 

  -n uint
    	number of words in generated password (default 4)
  -version
    	Print version number
```

## Tests

Run unit tests from the repository directory:

```sh
go test -v ./...
```

## Version and Change log

Program version is set in [about.go](about.go) with `VERSION` constant.

See [CHANGELOG.md](CHANGELOG.md).

## License

This project is licensed under the GNU GENERAL PUBLIC LICENSE Version 3.

```txt
    dpwgen  Copyright (C) 2019 Camille Daum-Lobko
    This program comes with ABSOLUTELY NO WARRANTY.
    This is free software, and you are welcome to redistribute it
    under certain conditions.
```

See the [LICENSE](LICENSE) for details.

## Authors

See [AUTHORS](AUTHORS).