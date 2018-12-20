# Version Checker
## Overview
This will return a webpage. I built this with the intention to return a specific webpage for an application which returns the version. Also I wanted to write something in Golang :)

## Build
Build the app using a multi-stage Dockerfile. The first is used to build the Go app and the second is used to run it. 

To build it, run the following:
```bash
docker build -t versionchecker:latest .
```

To run it:
```bash
docker run -e URL="https://thing.com/page/here" --name server versionchecker:latest
```

