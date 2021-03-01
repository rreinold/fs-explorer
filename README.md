# fs-explorer

## Purpose

List and view contents of files and directories over REST API.

## Setup
### Bare Metal
```
go mod download
go run fs-explorer.go
```
### Docker
```
docker build -t fs-explorer .
docker run -p 8080:8080 fs-explorer
```

## Usage

```
Usage of fs-explorer:
  -d string
    	Directory to host (Default: '.' ) (default ".")
```

## API

TODO Add OpenAPI Doc

## Credit

1. API Response structure derived from [directory-tree npm package](https://www.npmjs.com/package/directory-tree)
