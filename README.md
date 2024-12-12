# LuBu
LuBu is a simple Lua Bundler

## Usage
1. install lubu.exe from the latest release
2. create `lubu.json` file in your project folder
3. set up your config
```json
  {
    "main":"src/init.lua",
    "modules":{
        "first-module": "src/module1.lua",
        "second-module": "src/module2.lua"
    },
    "output":"dist/release.lua"
}
```
4. use `lubu.exe lubu.json`

## Project folder example
```
my-project/
├── src/
│   ├── init.lua
│   ├── module1.lua
│   └── module2.lua
└── lubu.json
```
