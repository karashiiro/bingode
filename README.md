# bingode
Binary data provider for [godestone](https://github.com/xivapi/godestone).

## Installation
`go get github.com/karashiiro/bingode`

## Usage
```go
bin := bingode.New()
s := godestone.NewScraper(bin, godestone.EN)
```

## Contributing

### Dependencies
  * A Unix-like shell environment (Git Bash works for this)
  * [`flatc`](https://google.github.io/flatbuffers)
  * [`go-bindata`](https://github.com/go-bindata/go-bindata)

### Updating

Clone the repository recursively and update the submodule to the latest commit:

```sh
git clone --recursive https://github.com/karashiiro/bingode
cd bingode/ && git submodule update --remote --merge
```

Run the `generate.sh` script:

```sh
./generate.sh
```

