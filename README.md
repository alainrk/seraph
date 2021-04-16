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

![ScreenExample02](https://github.com/alainrk/seraph/raw/master/example/02.png)
![ScreenExample03](https://github.com/alainrk/seraph/raw/master/example/03.png)
![ScreenExample04](https://github.com/alainrk/seraph/raw/master/example/04.png)
![ScreenExample05](https://github.com/alainrk/seraph/raw/master/example/05.png)
![ScreenExample06](https://github.com/alainrk/seraph/raw/master/example/06.png)

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
