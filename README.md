<div style="text-align:center"><img src="https://www.blast.hk/attachments/259340/" /></div>

# LuBu
**LuBu** is a simle lua bundler which allows you to compress .lua and .dll modules into one.

## Examples
See [`example-project`](https://github.com/chaposcripts/lubu/tree/main/example-project) folder

## Usage
1. install lubu.exe from the latest release
2. create `lubu.json` file in your project folder
3. set up your config
4. use `lubu.exe lubu.json`

## LuBu Config
LuBu config must have fields "modules", "modules" and "output". Also you can add some constants in "const"
1. `main` - `string` - path to main file
2. `modules` - `map[string]string` - **.lua / .dll** modules list, where key is module name (used in `require()`), value is a path to module
4. `output` - `string` - path to bundled file
5. `const` - `map[string]interface{}` - constants list, where key is variable name and value is a constant value. **Only string, number and bool are supported**
6. `watcher_delay` - `float64` - delay for "watcher" in milliseconds. Watcher will check files modification time with this interval. If one of files was changed lubu will re-bundle your script.
  
```json
{
    "modules": {
        "sum": "src/sum.lua",
        "mul": "src/mul.lua",
        "lfs": "src/lfs.dll"
    },
    "const": {
        "VERSION": "1.1a"
    },
    "main": "src/init.lua",
    "output": "dist/release.lua",
    "watcher_delay": 250
}
```
Also you can find auto json generator in `example` folder. Run it using `go run generate-lubu-config.go`
## Project folder example
```
my-project/
├── src/
│   ├── lfs.dll
│   ├── init.lua
│   ├── add.lua
│   └── mul.lua
├── lubu.exe
└── lubu.json
```

## Building
1. `git clone https://github.com/chaposcripts/lubu/`
2. `cd lubu`
3. `go build` or use `go build && lubu.exe lubu.json` to run it after building

## TODO
* [x] add .dll modules support (Done!)
* [ ] add resource bundler (ttf, png, etc.)
