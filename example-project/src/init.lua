local lfs = require('lfs');
local MyMath = require('my-math');
local Utils = require('utils');

print('Script path:', lfs.currentdir());

local numbers = { 1, 10, 71, 22, 39 };
print('Numbers:');
Utils.printTable(numbers);
print('Sum of this numbers is', MyMath.sum(table.unpack(numbers)));