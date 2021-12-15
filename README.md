## How to use this repo

This repo is part of the Marketplace team's [Golang onboarding guide](https://auth0team.atlassian.net/wiki/spaces/eco/pages/2643725202/WIP+Marketplace+Golang+onboarding+guide).
Follow along using our [step-by-step walkthrough](https://auth0team.atlassian.net/wiki/spaces/eco/pages/2646311412/WIP+Build+your+own+Golang+grocery+list)! 

## Getting started

### Prereqs 

1. You will need to install both Go and the `protoc` compiler (version 3):
- [Go installation](https://go.dev/doc/install)
- [protoc installation](https://grpc.io/docs/protoc-installation/)

2. Install the protobuf and Twirp plugins:
```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
$ go get google.golang.org/grpc
$ go get github.com/twitchtv/twirp/protoc-gen-twirp
```

3. Update your `PATH` so the `protoc` compiler can find the plugins:
```
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

Relevant docs:
- [gRPC quickstart](https://grpc.io/docs/languages/go/quickstart/)
- [Twirp quickstart](https://twitchtv.github.io/twirp/docs/install.html)


### Database setup with Docker

You will need a Postgres instance running on the default port with a db called `root` and a table called `groceries`:
```
$ docker run --name grocery-list -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres
$ docker exec -it grocery-list psql
$ create table groceries ( id SERIAL PRIMARY KEY, item TEXT, quantity INT );
```