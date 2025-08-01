-- numbers
local num = 1;
globalNum = 99;

-- table
t = {
    fieldWithNumber = tonumber("3.14"),
    [2] = 'test',
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


function t.fns.test()

end

function t.fns:method()

end

function t.a()

end

t.b = function()

end

function t.c(n)
    print('Your number is (it may be 99)', n)
end

function t:z()

end

-- table field index
print(183482)
print(t.a(1));
print(t.fieldWithNumber);
print(t:b('hello'))
t.c(
    1
)

for i = 1, 10 do
    print(i)
end