--[[Bundled using LuBu - Simple Lua Bundler]]

--Module add (src/add.lua)
package.preload['add'] = (function(...)
return function(a, b)
    return a + b;
end
end)

--Module mul (src/mul.lua)
package.preload['mul'] = (function(...)
return function(a, b)
    return a * b;
end
end)

--Entry Point main (src/init.lua)
local entry = (function(...)
local add = require('add');
local mul = require('mul');

print(add(10, 5));
print(mul(2, 2));
end)
entry();