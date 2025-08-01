-- Это однострочный комментарий
local x = 5  -- inline комментарий

--[[
  Это многострочный
  комментарий
]]

--[=[
  Еще один вариант
  многострочного комментария
]=]

print("Hello")  --[[ коммент ]] print("World")
-- numbers
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