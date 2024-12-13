--[[
																						
LLLLLLLLLLL                               BBBBBBBBBBBBBBBBB                     
L:::::::::L                               B::::::::::::::::B                    
L:::::::::L                               B::::::BBBBBB:::::B                   
LL:::::::LL                               BB:::::B     B:::::B                  
L:::::L               uuuuuu    uuuuuu    B::::B     B:::::Buuuuuu    uuuuuu  
L:::::L               u::::u    u::::u    B::::B     B:::::Bu::::u    u::::u  
L:::::L               u::::u    u::::u    B::::BBBBBB:::::B u::::u    u::::u  
L:::::L               u::::u    u::::u    B:::::::::::::BB  u::::u    u::::u  
L:::::L               u::::u    u::::u    B::::BBBBBB:::::B u::::u    u::::u  
L:::::L               u::::u    u::::u    B::::B     B:::::Bu::::u    u::::u  
L:::::L               u::::u    u::::u    B::::B     B:::::Bu::::u    u::::u  
L:::::L         LLLLLLu:::::uuuu:::::u    B::::B     B:::::Bu:::::uuuu:::::u  
LL:::::::LLLLLLLLL:::::Lu:::::::::::::::uuBB:::::BBBBBB::::::Bu:::::::::::::::uu
L::::::::::::::::::::::L u:::::::::::::::uB:::::::::::::::::B  u:::::::::::::::u
L::::::::::::::::::::::L  uu::::::::uu:::uB::::::::::::::::B    uu::::::::uu:::u
LLLLLLLLLLLLLLLLLLLLLLLL    uuuuuuuu  uuuuBBBBBBBBBBBBBBBBB       uuuuuuuu  uuuu

	Bundled usin LuBu - https://github.com/chaposcripts/lubu
]]

LUBU_BUNDLED = true;

--LuBu Constants
DEV = false;

--Module add (X:\dev\lubu\example-project/src/add.lua)
package.preload['add'] = (function(...)
return function(a, b)
    return a + b;
end
end)

--Module mul (X:\dev\lubu\example-project/src/mul.lua)
package.preload['mul'] = (function(...)
return function(a, b)
    return a * b;
end
end)

--Entry Point main (X:\dev\lubu\example-project/src/init.lua)
local entry = (function(...)
local add = require('add');
local mul = require('mul');

print(add(10, 5));
print(mul(2, 2));
end)
entry();