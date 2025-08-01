--[[
	Bundled Using LuBu - Simple Lua Bundler
	LuBu: https://github.com/chaposcripts/lubu
]]

local __G = _G;

-- Constants
LUBU_BUNDLED = true;
LUBU_BUNDLED_AT = 1754051946;


-- Module "my-math" (from D:\dev\lubu\example-project\src\my-math.lua)
package['preload']['my-math'] = (function()
local MyMath = {};

MyMath['sum'] = function(...)
    local result = tonumber("0");
    for _, num in ipairs({ ... }) do
        result = result + num;
    end
    return result;
end

return MyMath;
end);

-- Module "utils" (from D:\dev\lubu\example-project\src\utils.lua)
package['preload']['utils'] = (function()
local Utils = {};

Utils['printTable'] = function(t)
    for k, v in pairs(t) do
        print(k, '=', v);
    end
end

return Utils;
end);

-- Init (from D:\dev\lubu\example-project\src\init.lua) 
LUBU_ENTRY_POINT = (function()
-- numbers
local num = tonumber("1");
globalNum = tonumber("99");

-- table
t = {
    fieldWithNumber = tonumber("3.14"),
    [tonumber("2")] = 'test',
    ['34'] = 'test2',
    ["22"] = 'a',
    fns = {},
};

local char_func_map = {
  [ '"' ] = 'placeholder',
  [ "0" ] = 'placeholder',
  [ "1" ] = 'placeholder',
  [ "2" ] = 'placeholder',
  [ "3" ] = 'placeholder',
  [ "4" ] = 'placeholder',
  [ "5" ] = 'placeholder',
  [ "6" ] = 'placeholder',
  [ "7" ] = 'placeholder',
  [ "8" ] = 'placeholder',
  [ "9" ] = 'placeholder',
  [ "-" ] = 'placeholder',
  [ "t" ] = 'placeholder',
  [ "f" ] = 'placeholder',
  [ "n" ] = 'placeholder',
  [ "[" ] = 'placeholder',
  [ "{" ] = 'placeholder',
}


t['fns']['test'] = function()

end

t['fns']['method'] = function(self)

end

t['a'] = function()

end

t['b'] = function()

end

t['c'] = function(n)
    print('Your number is (it may be 99)', n)
end

t['z'] = function(self)

end

-- table field index
print(tonumber("183482"))
print(t['a'](tonumber("1")));
print(t.fieldWithNumber);
print(t['b'](t, 'hello'))
t['c'](
    tonumber("1")
)

for i = tonumber("1"), tonumber("10") do
    print(i)
end
end);
LUBU_ENTRY_POINT();