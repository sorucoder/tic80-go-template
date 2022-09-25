# tic80-go

This is an unofficial WASM binding for Go to make TIC-80 Cartridges.

## Usage

The included `tic80` package follows the native [TIC-80 API](https://github.com/nesbox/TIC-80/wiki/API) as closely as possible, including optional arguments.
For functions that have optional arguments, you can either use the defaults by passing `nil`, like so:

```go
tic80.Print("HELLO WORLD FROM GO!", 65, 84, nil)
```

Or, you can pass an instance of the corresponding `tic80.`**`<APIName>`**`Options`, chaining its methods to configure it, like so:

```go
tic80.Spr(1+t%60/30*2, x, y, tic80.NewSpriteOptions().AddTransparentColor(14).SetScale(3).SetSize(2, 2))
```

## Building

* Run `make build` to just build the WASM code.
* Run `make cartridge` to build the WASM code and cartridge.
* Run `make`, `make all`, or `make run` to build the WASM code, cartridge, and run the cartridge.