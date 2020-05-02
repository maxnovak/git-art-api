## Description
Generates a repository that will do some art in github's work bar - you can choose the year.

## How to use:
1. Install [Golang](https://golang.org/dl/)
2. Run command: `go run src/generate-art.go`
3. Enter prompted data and a new local repository will be created
4. Create a repository with the same name in github
5. Push the contents of the repository up to the new github repository
6. Observe your new github art :)

## Running Tests:
To run all tests in the src folder run:
```go
go test ./src/...
```
To run tests and get a coverage report run:
```go
go test ./src/... -coverage
```

## Matrix files
If you wish to use the matrix design behavior you will need to supply a file for processing.  The file should looks something like this:
```json
{
    "matrix":[
        [0,1,1,1,0,10,5],
        [1,0,1,0,1,0,1]
    ]
}
```