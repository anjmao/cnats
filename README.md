# cnats
NATS / NATS streaming command line utility for publishing and subscribing to messages.

### Installation

#### Go users

```
go get -u github.com/anjmao/cnats
```

#### Other users

Download release binary from https://github.com/anjmao/cnat/releases and add to PATH.

### Usage

#### Initialize cnat with configuration file which is written to HOME/.cnat/config.json

Initialize with cluser, client and url.
```
cnats init -cluster test-cluster -client cnat-client -url nats://localhost:4222
```

Initialize with default values.

```
cnats init
```

#### Subscribe to subjects

Subscribe to nats streaming subjects.
```
cnats ssub topic1 topic2
```

Subscribe to all known subjects. Internally sub command calls stan monitoring api to get all known subjects.
```
cnats ssub
```

Subscribe to nats subjects.
```
cnats sub topic1 topic2
```


#### Publish to subject

Publish to nats streaming.
```
cnats spub topic1 '{"name": "cnat"}'
```

Publish to nats.
```
cnats pub topic1 '{"name": "cnat"}'
```

### Release

Before running `make release` make sure you have these tools installed.

```
GO111MODULE=off go get github.com/c4milo/github-release
GO111MODULE=off go get github.com/mitchellh/gox 
```