print('Current script directory:', require('lfs').currentdir());
local sum, mul = require('sum'), require('mul');

local a, b, c = 1, 5, 10;
print(sum(a, b));
print(sum(b, c));
print(mul(a, b));
print(mul(c, b));