# Seraph
[![Status](https://github.com/alainrk/seraph/workflows/Go/badge.svg)](https://github.com/alainrk/seraph/actions)

Offline, multi-vaults, terminal-based passwords/secrets keeper.

![Seraph Matrix Banner](https://github.com/alainrk/seraph/raw/master/example/matrix.png)

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

### Test
```
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
