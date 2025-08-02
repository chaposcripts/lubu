--[[
	Bundled Using LuBu - Simple Lua Bundler
	LuBu: https://github.com/chaposcripts/lubu
]]

local __G = _G;

-- Constants
LUBU_BUNDLED = true;
LUBU_BUNDLED_AT = tonumber("1754149639");
n = "2";
ARTUR_DILBAROV = "PIDORAS, GOVNOJUY, INCEL I DOLBAEB";
TRUTH = "vktr shplv i l u";
b = "true";


-- Module "utils" (from D:\dev\lubu\example-project\src\utils.lua)
package['preload']['utils'] = (function()
Utils = {};

Utils['printTable'] = function(t)
    for k, v in pairs(t) do
        print(k, '=', v);
    end
end
end);

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

-- Init (from D:\dev\lubu\example-project\src\init.lua) 
LUBU_ENTRY_POINT = (function()
require('utils');
local MyMath = require('my-math');

local numbers = {
	tonumber("1"),
	tonumber("10"),
	tonumber("81"),
	tonumber("28"),
	tonumber("19")
};

print('Numbers:');
Utils['printTable'](numbers);
print('Sum:', MyMath['sum'](table.unpack(numbers)));

-- ObfuscationPrepare example:
local Developer = {
	name = 'Dmitry',
	age = tonumber("21"),
	---@OBFIGNORE
	height = 191,
	---@ENDOBFIGNORE
	weight = tonumber("80"),
	citizenship = 'Russian Federation',
	skills = {}
};

Developer['say'] = function(self, ...)
	print(('%s says: %s'):format(self.name, table['concat']({ ... }, ' ')));
end

Developer['learn'] = function(self, skill)
	table['insert'](self.skills, skill);
end

Developer['hear'] = function(self, text)
	local text = text['lower'](text);
	if (text == 'how old are you?') then
		self['say'](self, ('i\'m %d years old!'):format(self.age));
	elseif (text == 'which skills do you have?') then
		if (#self.skills == 0) then
			self['say'](self, 'I don\'t have any skills :(');
		else
			self['say'](self, 'I have', table.concat(self.skills, ', '), 'skills!');
		end
	end
end

Developer['say'](Developer, 'Hello!');
Developer['hear'](Developer, 'which skills do you have?')
end);
LUBU_ENTRY_POINT();