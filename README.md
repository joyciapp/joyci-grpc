# joyci-grpc
This repo has all sources regarding the GRPC server

# GRPC

## Protocol Buffers
To compile protobuffers, in the project's root folder run the command bellow:
```
$ protoc -I grpc grpc/proto/core.proto --go_out=plugins=grpc:grpc
```

# Releases

To release a new version:
```
$ git tag -a vx.x.x
$ git push origin vx.x.x
```