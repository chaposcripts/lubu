require('utils');
local MyMath = require('my-math');

local numbers = {
	1,
	10,
	81,
	28,
	19
};

print('Numbers:');
Utils.printTable(numbers);
print('Sum:', MyMath.sum(table.unpack(numbers)));

-- ObfuscationPrepare example:
local Developer = {
	name = 'Dmitry',
	age = 21,
	---@OBFIGNORE
	height = 191,
	---@ENDOBFIGNORE
	weight = 80,
	citizenship = 'Russian Federation',
	skills = {}
};

function Developer:say(...)
	print(('%s says: %s'):format(self.name, table.concat({ ... }, ' ')));
end

function Developer:learn(skill)
	table.insert(self.skills, skill);
end

function Developer:hear(text)
	local text = text:lower();
	if (text == 'how old are you?') then
		self:say(('i\'m %d years old!'):format(self.age));
	elseif (text == 'which skills do you have?') then
		if (#self.skills == 0) then
			self:say('I don\'t have any skills :(');
		else
			self:say('I have', table.concat(self.skills, ', '), 'skills!');
		end
	end
end

Developer:say('Hello!');
Developer:hear('which skills do you have?')