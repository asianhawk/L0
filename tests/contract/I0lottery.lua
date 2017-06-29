-- 用合约来完成 福彩
local L0 = require("L0")

-- 合约创建时会被调用一次, 完成数据初始化
function L0Init()
    L0.PutState("balance", 0)
    L0.PutState("state", "ongoing")
    L0.PutState("winner", "")
    return true
end

-- 每次合约执行都调用
function L0Invoke(func, args)
    if func == "buy" then
    elseif func == "finish" then   
    end

    return true
end

-- 每次合约查询都调用
function L0Query(func, args)
    if func == "all"  then 

    end
    return true
end