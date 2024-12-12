# adventofcode-2024
https://adventofcode.com/2024

# Usage

Each day is a self-enclosed go program. Input is read off stdin, pipe in either the sample or input to run. Command line arg to main.go is whether to invoke part 1 or 2 of each day's problem.

```
cat dayX/[sample|input].txt | go run dayX/main.go [1|2]
```

To run tests, `go test` will need need the local path to contain the main.go file.

```
cd dayX
go test
```