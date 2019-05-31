# cnat
Nats / Nats streaming command line utility for publishing and subscribing to messages.

### Installation

#### Go users

```
go get -u github.com/anjmao/cnat
```

#### Other users

Download release binary from https://github.com/anjmao/cnat/releases and add to PATH.

### Usage

#### Initialize cnat with configuration file which is written to HOME/.cnat/config.json

Initialize with cluser, client and url.
```
cnat init -cluster test-cluster -client cnat-client -url nats://localhost:4222
```

Initialize with default values.

```
cnat init
```

#### Subscribe to subjects

Subscribe to subjects.
```
cnat sub topic1 topic2
```

Subscribe to all known subjects. Internally sub command calls stan monitoring api to get all known subjects.
```
cnat sub
```

#### Publish to subject

```
cnat pub topic1 '{"name": "cnat"}'
```