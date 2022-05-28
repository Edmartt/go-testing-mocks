# GO-TESTING-MOCKS

The main purpose of this project is understanding the importance of knowing about what to do
when our services main core depends on third party services like http services, databases, etc.

For reaching our goal of testing the service layer we make use of mocks. With a mock we can replace those third parties services calls
with our own logic according to the data we expect from those services.


### Running Project

```
go mod download
```

or

```
go mod tidy
```


```
go run main.go
```
When the project is running a hardcoded user will be created in a SQLite database. If we run the project again with the same
hardcoded user we'll have an error of nil because the CreateUser function has a check in case the user already exists.


### Running Unit Test with a Coverage file as Output

```
go test -v --coverprofile=coverage.out ./... ./...
```

### Watching the coverage in browser

```
go tool cover -html=coverage.out
```
