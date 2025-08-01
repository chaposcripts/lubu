--[[
	Bundled Using LuBu - Simple Lua Bundler
	LuBu: https://github.com/chaposcripts/lubu
]]

local __G = _G;

-- Constants
LUBU_BUNDLED = true;
LUBU_BUNDLED_AT = 1754054962;


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


-- -- table
-- t = {
--     fieldWithNumber = tonumber("3.14"),
--     [tonumber("2")] = 'test',
--     ['34'] = 'test2',
--     ["22"] = 'a',
--     fns = {},
-- };


-- t.fns['test'] = function()

-- end

-- t['fns']['method'] = function(self)

-- end

-- t['a'] = function()

-- end

-- t['b'] = function()

-- end

-- t['c'] = function(n)
--     print('Your number is (it may be 99)', n)
-- end

-- t['z'] = function(self)

-- end

-- -- table field index
-- print(tonumber("183482"))
-- print(t['a'](tonumber("1")));
-- print(t.fieldWithNumber);
-- print(t['b'](t, 'hello'))
-- t['c'](
--     tonumber("1")
-- )

-- for i = tonumber("1"), tonumber("10") do
--     print(i)
-- end
end);
LUBU_ENTRY_POINT();