# WC

A similar version of the Unix command line tool wc!

wc is a simple Go CLI built with Cobra that will be capable of doing the followings:

Read files and calculates:
- file size in bytes
- number of words in the file
- number of lines in the file
- number of characters in the file

## Installation

> Make sure you have **[Go](https://go.dev)** installed

- Clone the repository.
```bash
    git clone https://github.com/EmmanuelOmoiya/wc
```
- Install all the packages.
```bash
    go get . && go install
```

- Run the cli.
```bash
    wc
```

## Usage
```txt
wc count -c test.txt
```

## Commands

```txt
Print  characters, line, word, and  byte counts for each FILE, and a total line if more than one FILE is specified.  A word is a non-zero-length sequence of characters delimited by white space.

Usage:
  wc [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  count       count - print characters, line, word, and byte counts for each file
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.wc.yaml)
  -h, --help            help for wc

Use "wc [command] --help" for more information about a command.
```

## Subcommands

```txt
 Print  characters, line, word, and  byte counts for each FILE, and a total line if more than one FILE is specified.  A word is a non-zero-length sequence of characters delimited by white space.

Usage:
  wc wc [flags]

Flags:
  -c, --byte        Display the number of bytes in that file
  -m, --character   Display the number of characters in that file
  -h, --help        help for wc
  -l, --line        Display the number of lines in that file
  -w, --word        Display the number of words in that file

Global Flags:
      --config string   config file (default is $HOME/.wc.yaml)
```