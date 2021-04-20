# Seraph
[![YourActionName Actions Status](https://github.com/alainrk/seraph/workflows/Go/badge.svg)](https://github.com/alainrk/seraph/actions)
[![experimental](http://badges.github.io/stability-badges/dist/experimental.svg)](http://github.com/badges/stability-badges)

Multi-vaults secrets keeper

![Seraph Matrix Banner](https://github.com/alainrk/seraph/raw/master/example/matrix.png)

Experimental. DON'T USE IT YET.

## Usage
```
$ ./seraph
# Follow interactive instructions
```
### Secrets handling
![Seraph GIF Example](https://github.com/alainrk/seraph/raw/master/example/seraph.gif)

### Multiple vaults
![Multiple Vaults GIF Example](https://github.com/alainrk/seraph/raw/master/example/vaults.gif)

## Dev

### Mod init
```
go mod init github.com/alainrk/seraph
```

### Test everything recursively
```
chmod +x test.sh
./test.sh
```

### Build everything recursively
```
go build ./...
```

### Cross compilation for multiple arch
```
chmod +x crosscompile.sh
./crosscompile.sh
```
