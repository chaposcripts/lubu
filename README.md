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
7. `resource` - `map[string]string` - list of resources. All resources will be converted to bytes, and after running the script, source files will be created from these bytes. The key is the path that will be used to create the file after running the script. The value is the path to the compressed file.
8. `prepare_for_obfuscation` - `bool` - preparing for script obfuscation

#### Bundle config example
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
    "out": "dist/release.lua",
    "watcher_delay": 250,
    "prepare_for_obfuscation": true
}
```

<details>
  <summary>Obfuscation Prepare example</summary>  
  
### Preparing for obfuscation
1. Change all `number` values to `tonumber("NUMBER")`
2. Change all table functions defenitions and calls

Also you can add "ingnoring" blocks:
```lua
---@OBFIGNORE
local anotherNumber = 123; -- this number will NOT replaced to "tonumber" cuz of "ignoring zone"
---@ENDOBFIGNORE
```
#### Before:
```lua
local localNum = 1;
globalNum = 99;

local t = {
    [1] = 1,
    ['2'] = 'two',
    funcs = {}
};

function t.a() end
function t.funcs.a() end
function t.funcs:method()
    print(tostring(self));
end

print('Number inside string will NOT replaced to "tonumber": 999');
for i = 1, 100 do
    print(i .. '%');
end


---@OBFIGNORE
local anotherNumber = 123; -- this number will NOT replaced to "tonumber" cuz of "ignoring zone"
---@ENDOBFIGNORE

t.a();
t.funcs.a();
t.funcs:method();
print([[
    test3
    1
    2
]]);
```
#### After:
```lua
local localNum = tonumber("1");
globalNum = tonumber("99");

local t = {
    [tonumber("1")] = tonumber("1"),
    ['2'] = 'two',
    funcs = {}
};

t['a'] = function() end
t.funcs['a'] = function() end
t['funcs']['method'] = function(self)
    print(tostring(self));
end

print('Number inside string will NOT replaced to "tonumber": 999');
for i = tonumber("1"), tonumber("100") do
    print(i .. '%');
end


---@OBFIGNORE
local anotherNumber = 123; -- this number will NOT replaced to "tonumber" cuz of "ignoring zone"
---@ENDOBFIGNORE

t['a']();
t['funcs']['a']();
t['funcs']['method'](t['funcs']);
print([[
    test3
    1
    2
]]);
```
</details>

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
* [x] add resource bundler (ttf, png, etc.)
* [ ] add "-g" parameter, wich will automatically generate bundle config
