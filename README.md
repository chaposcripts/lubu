<div style="text-align:center"><img src="https://www.blast.hk/attachments/259340/" /></div>

# LuBu
**LuBu** is a simle lua bundler which allows you to combine many lua scripts into one.

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
2. `modules` - `map[string]string` - modules list, where key is module name (using in `require()`), value is a path to module
3. `output` - `string` - path to bundled file
4. `const` - `map[string]interface{}` - constants list, where key is variable name and value is a constant value. **Only string, number and bool are supported**
  
```json
{
    "main": "src/init.lua",
    "modules": {
        "add": "src/add.lua",
        "mul": "src/mul.lua"
    },
    "output": "dist/release.lua",
    "const": {
        "DEV": false,
        "BUNDLED": true,
        "VERSION": "1.0.0"
    }
}
```
Also you can find auto json generator in `example` folder. Run it using `go run generate-lubu-config.go`
## Project folder example
```
my-project/
├── src/
│   ├── init.lua
│   ├── add.lua
│   └── mul.lua
└── lubu.json
```

## Building
1. `git clone https://github.com/chaposcripts/lubu/`
2. `cd lubu`
3. `go build` or use `go build && lubu.exe lubu.json` to run it after building
