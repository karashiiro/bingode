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
Make sure to checkout the submodules if you are changing pack information.

### Dependencies
  * [`flatc`](https://google.github.io/flatbuffers)
  * [`go-bindata`](https://github.com/go-bindata/go-bindata)
