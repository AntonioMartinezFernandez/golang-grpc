# Client/Server gRPC PoC with Golang

## Start Services

With Docker:

```
docker compose up -d
```

With Air:

1. Copy `example.env` to `.env` and customize values

2. Launch server:

```
air -c .\grpc-server.air.toml
```

3. Launch client:

```
air -c .\grpc-client.air.toml
```

## Steps to implement a gRPC service

1. Define the protocol buffer (.proto file) with the structures and methods (this definition is like an interface for remote procedures)
2. Compile the protocol buffer with `protoc`

### Protoc example command

Execute the command from the location folder of the .proto file:

```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld.proto
```

## Requirements

1. Docker desktop
2. Golang runtime
3. Protocol Buffer compiler -protoc- (install with `brew` -Mac- or `chocolatey` -Win-)
4. protoc-gen

Install protoc-gen:

```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## Docs

- [gRPC Quick start with Golang](https://grpc.io/docs/languages/go/quickstart/)
- [Protocol Buffers Language Guide -proto 3-](https://protobuf.dev/programming-guides/proto3/)
