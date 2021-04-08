# Seraph
[![experimental](http://badges.github.io/stability-badges/dist/experimental.svg)](http://github.com/badges/stability-badges)

Rewriting in go of an old secrets vault.
Experimental. DON'T USE IT YET.

# Usage
```
$ ./seraph -h

Usage of ./seraph:
  -d    decrypt mode
  -e    encrypt mode
  -p value
        passphrase [INSECURE method, use interactive mode instead]
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
