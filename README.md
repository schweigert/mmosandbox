[![Coverage Status](https://coveralls.io/repos/github/schweigert/mmosandbox/badge.svg)](https://coveralls.io/github/schweigert/mmosandbox)
[![Build Status](https://travis-ci.org/schweigert/mmosandbox.svg?branch=master)](https://travis-ci.org/schweigert/mmosandbox)

# mmosandbox
MMORPG Playground

## Deps

Setup:

```
    go get github.com/kisielk/godepgraph
    sudo pacman -S graphviz
```

```
    make deps
```

## Domain

![deps](domain/deps.png)

## Services

### apps/wauth

![deps](apps/wauth/deps.png)

### apps/wgame

![deps](apps/wgame/deps.png)

### apps/wweb

![deps](apps/wweb/deps.png)

## Clients

### clients/wclient

![deps](clients/wclient/deps.png)

# Deploy

## beacon

Beacon send `docker stats` log to graphite server ¯\\_(ツ)_/¯:

```
NODE=node METRIC_HOST=localhost METRIC_PORT=2003 go run beacon/beacon.go
```