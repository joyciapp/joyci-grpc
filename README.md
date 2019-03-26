# joyci-grpc
This repo has all sources regarding the GRPC server

# GRPC

## Protocol Buffers
To compile protobuffers, in the project's root folder run the command bellow:
```
$ protoc -I grpc grpc/proto/core.proto --go_out=plugins=grpc:grpc
```

# Docker
To run the server inside a docker container
```
$ docker run --rm -it -v $(pwd):/tmp/joyci-grpc -w /tmp/joyci-grpc golang:1.11 /bin/bash
$ go build
$ go install
$ joyci-grpc
```

It should output something like
```
2019/03/26 13:23:07 JoyCI GRPC server started at  :50051
```

# Releases

To release a new version:
```
$ git tag -a vx.x.x
$ git push origin vx.x.x
```