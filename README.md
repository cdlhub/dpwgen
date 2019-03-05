# dpwgen

Diceware passphrase generator command line tool written in Go. It generates "random" passphrase from a wordlist file based on EFF article [Deep Dive: EFF's New Wordlists for Random Passphrases](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases).

## Prerequisites

* Go language: Follow installation steps at Golang documentation [Getting started](https://golang.org/doc/install) page.

## Installation

### Get Executable

Use Go tool to install `dpwgen` executable:

```sh
go get -v github.com/cdlhub/dpwgen
```

### Build from Sources

Use `git` to download sources:

```sh
git clone https://github.com/cdlhub/dpwgen.git
```

And then build the package:

```sh
cd dpwgen
go build
```

## Usage

You need a diceware wordlist file as described in [Deep Dive: EFF's New Wordlists for Random Passphrases](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases), alsop provided in the repository.

You can then generate a passphrase with the number of words you want (i.e. 6):

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

Run:

```sh
go test -v ./...
```

## License

Licensed under GNU GENERAL PUBLIC LICENSE Version 3.

```txt
    dpwgen  Copyright (C) 2019 Camille Daum-Lobko
    This program comes with ABSOLUTELY NO WARRANTY.
    This is free software, and you are welcome to redistribute it
    under certain conditions.
```

See full [LICENSE terms](LICENSE).

## Authors

See [AUTHORS](AUTHORS).