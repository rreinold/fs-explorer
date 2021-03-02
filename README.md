# fs-explorer

## Overview

List and view contents of files and directories over REST API.

### Support
Supported OSs:
- Linux
- OS X

Tested On:
- Mac OS X 10.15.5
- Alpine Linux 3.13.2

### Approach

Make use of Golang's native libraries, `io` and `os` to access the filesystem. This was selected in place of a system call to run `ls` command in order to preclude any arbitrary code injections at the command line. Responses follow an [existing npm approach](https://www.npmjs.com/package/directory-tree) for representing filesystems in JSON.

## Usage

Runs on default gin port: 8080

```
Usage of fs-explorer:
  -d string
    	Directory to host (Default: "." )
```

### 1. Run script

```
$ ./run.sh
```
 ### 2.Run in Docker, from Docker Hub (Recommended)

```
docker run -p 8080:8080 rreinold/fs-explorer:0.1.0
```

### 2. Run on Bare Metal
```
go run fs-explorer.go -d foo
```

### 3. Run in Docker, from image
```
GOOS=linux go build fs-explorer.go
docker build -t fs-explorer:dev .
docker run -p 8080:8080 fs-explorer:dev
```

## API

[View OpenAPI Spec](https://github.com/rreinold/fs-explorer/blob/master/openapi.yml)

### Example

```
$ curl -s localhost:8080/bar

{
  "name": "bar",
  "owner": 503,
  "size": 128,
  "permissions": "-rwxr-xr-x",
  "isDir": true,
  "links": [
    { "name": "bar1", "isDir": false, "path": "/bar/bar1", "href": "/bar/bar1", "type": "GET" },
    { "name": "baz", "isDir": true, "path": "/bar/baz", "href": "/bar/baz", "type": "GET" }
  ],
  "path": "/bar",
  "contents": ""
}
```
## Testing

Docker images are bundled with and host a test directory: `foo`

### Unit Tests

Unit tests are available, and are run with:

```
go test ./...
```

### System Tests

This is an outstanding item, which should rely on a testing harness that support REST calls.

## Roadmap

- For v1.0.0, prepend basepath of 'v1' for backwards compatibility
- Add concurrency on os.Stat calls for fetching multiple file details
- Add System Tests


## Credit

1. API Response structure derived from [directory-tree npm package](https://www.npmjs.com/package/directory-tree)
