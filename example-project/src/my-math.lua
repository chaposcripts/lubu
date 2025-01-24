local MyMath = {};

function MyMath.sum(...)
    local result = 0;
    for _, num in ipairs({ ... }) do
        result = result + num;
    end
    return result;
end

return MyMath;