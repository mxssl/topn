# topn

topn finds top N numbers in a file

## Task

Write a program, topN, that given an arbitrarily large file and a number, N, containing individual numbers on each line (e.g. 200Gb file), will output the largest N numbers, highest first.

## Install

- [Install go](https://golang.org/dl/)
- `make build`

## Usage

### Generate file

`topn generate --lines N --name filename` - generate `file` with `N` lines of random numbers.

### Find the top N of largest numbers

`topn run --file filename --top N`

### Example

- Generate file `numbers.txt` with `1000` lines of random numbers

```sh
./topn generate --name numbers.txt --lines 1000
File numbers.txt with 1000 lines of random numbers was successfully generated!
```

- Find top 5 largest numbers in file `numbers.txt`

```sh
./topn run --file numbers.txt --top 5
[9219151525891662339 9215619702456294450 9207997407084704522 9207450753580197605 9173624551887931713]
```

## Test

```sh
go test -cover -v ./...
?   	github.com/mxssl/topn	[no test files]
=== RUN   TestCase1
expected:  [100 90 80 70 60]
output:  [100 90 80 70 60]
--- PASS: TestCase1 (0.00s)
PASS
coverage: 42.3% of statements
ok  	github.com/mxssl/topn/cmd	0.260s	coverage: 42.3% of statements
```
