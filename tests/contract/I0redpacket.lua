-- 用合约来完成红包
local L0 = require("L0")

-- 合约创建时会被调用一次, 完成数据初始化
function L0Init()
    local number  = 5
    local money  = 100
    L0.PutState("number", number)
    L0.PutState("money", money)
    return true
end

-- 每次合约执行都调用
function L0Invoke(func, args)
    if func == "snatch" then
        cnt :=  L0.GetState("number") --剩余个数
        if cnt == 0 {
            return false --抢光了
        }

        receiver := args[0]
        if L0.GetState(receiver) > 0 {
            return false --已抢过
        }

        money := L0.GetState("money") --剩余金额
        balance := money
        L0.PutState(receiver, balance)
        L0.Transfer(receiver, balance)
        L0.PutState("money", money-balance)

        L0.PutState("receivers",  L0.GetState("receivers") .. "," .. receiver)
    end

    return true
end

-- 每次合约查询都调用
function L0Query(func, args)
    if func == "all"  then 
       receivers = L0.GetState("receivers")
    --     while (true) do
    --     local pos = string.find(receivers, ",");
    --     if (not pos) then
    --         sub_str_tab[#sub_str_tab + 1] = str;
    --         break;
    --     end
    --     local receiver = string.sub(receivers, 1, pos - 1);
    --     str = string.sub(str, pos + 1, #str);
    -- end
       
    end
    return true
end