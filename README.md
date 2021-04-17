# Seraph
[![YourActionName Actions Status](https://github.com/alainrk/seraph/workflows/Go/badge.svg)](https://github.com/alainrk/seraph/actions)
[![experimental](http://badges.github.io/stability-badges/dist/experimental.svg)](http://github.com/badges/stability-badges)

Ongoing password and secrets vaults manager.

Experimental. DON'T USE IT YET.

# Usage [NEW]
```
$ ./seraph
# Follow interactive instructions
```

## First POC example

![Seraph GIF Example](https://github.com/alainrk/seraph/raw/master/example/seraph.gif)

# Usage [OLD]
```
$ ./seraph -h

Usage of ./seraph:
  -d    decrypt mode
  -e    encrypt mode
  -p value
        password [INSECURE method, use interactive mode instead]
```

# Dev
## Mod init
```
go mod init github.com/alainrk/seraph
```

## Build everything recursively
```
go build ./...
```

## Test everything recursively
```
go test ./...
```
