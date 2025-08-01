

local __G = _G;


LUBU_BUNDLED = true;
LUBU_BUNDLED_AT = tonumber("1754059361");



LUBU_ENTRY_POINT = (function()

local x = tonumber("5")  





print("Hello")   print("World")

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



local anotherNumber = 123; 


t['a']();
t['funcs']['a']();
t['funcs']['method'](t['funcs']);
print([[
    test3
    1
    2
]]);
end);
LUBU_ENTRY_POINT();