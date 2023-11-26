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

## Requirements

1. Docker desktop
2. Golang runtime
3. Protocol Buffer compiler -protoc- (install with `brew` -Mac- or `chocolatey` -Win-)

## Docs

- [gRPC Quick start with Golang](https://grpc.io/docs/languages/go/quickstart/)
