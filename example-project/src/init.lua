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


-- -- table
-- t = {
--     fieldWithNumber = tonumber("3.14"),
--     [2] = 'test',
--     ['34'] = 'test2',
--     ["22"] = 'a',
--     fns = {},
-- };


-- function t.fns.test()

-- end

-- function t.fns:method()

-- end

-- function t.a()

-- end

-- t.b = function()

-- end

-- function t.c(n)
--     print('Your number is (it may be 99)', n)
-- end

-- function t:z()

-- end

-- -- table field index
-- print(183482)
-- print(t.a(1));
-- print(t.fieldWithNumber);
-- print(t:b('hello'))
-- t.c(
--     1
-- )

-- for i = 1, 10 do
--     print(i)
-- end