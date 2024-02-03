# pdf

A Go CLI utility to manipulate pdf files.

## Prerequisites

The tools `ghostscript` and `imagemagick` must be installed in the system.

## Install

### Go install

Run `go install` to download the binary to the go's binary folder:

```bash
go install github.com/nu12/pdf@latest
```

Note: go's binary folder (tipically `~/go/bin`) should be added to your PATH.

### Releases

Download a tagged release binary for your OS (ubuntu, macos, windows) placing it in a folder in your PATH and make it executable (may require elevated permissions):

### From source

Clone this repo and compile the source code:

```bash
git clone github.com/nu12/pdf
cd pdf
go build -o pdf main.go
```

Move binary to a bin folder in your PATH (may require elevated permissions):
```bash
mv pdf /usr/local/bin/
```

## Usage

```
Append PDF files

Usage:
  pdf [command]

Available Commands:
  append      Append PDFs and images into a single PDF file
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     Show current version

Flags:
  -h, --help                         help for pdf
  -t, --temporary-directory string   Temporary directory (default "/tmp")

Use "pdf [command] --help" for more information about a command.
```

### Docker

Run the following docker command to run `pdf` without a local instalation (we must mount a volume so that the container has access to the local filesystem):

```
docker run --rm -v $(pwd):/app/files ghcr.io/nu12/pdf append files/input1 [files/input2 files/input3 ... files/inputN] files/output [flags]
```
