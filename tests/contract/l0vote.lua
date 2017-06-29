-- 用合约来完成一个投票系统
local L0 = require("L0")

-- 合约创建时会被调用一次, 完成数据初始化
function L0Init(func, args)
    print("L0Init")
    --初始化参与人
    for k, v in ipairs(voters) 
    do
        L0.PutState("voter:" .. v,  voter:new(v, ""))
    end
    
    --初始化候选项
    for k, v in ipairs(proposals) 
    do
        L0.PutState("proposal:" .. v, proposal:new(v))
    end
    return true, "ok"
end

-- 每次合约执行都调用
function L0Invoke(func, args)
    print("L0Invoke",func,args[0],args[1])
    if func == "vote" then
        local vote = args[0]
        local proposal = args[1]
        local voter = L0.GetState("voter:" .. vote)
        local proposaler = L0.GetState("proposal:" .. proposal)
        if voter == nil or voter.voted or proposaler == nil then
            return false, "voter or proposaler is nil "
        end
        voter.vote = vote
        voter.voted = true
        voter.voteTime = os.time()
        L0.PutState("voter:" .. v,  voter)
        proposaler.voteCount = proposaler.voteCount + 1
        L0.PutState("proposal:" .. v, proposaler)
    end
    return true, "ok"
end

-- 每次合约查询都调用
function L0Query(func, args)
    print("L0Query")
    if func == "vote" then
        local vote = args[0]
        local voter = L0.GetState("voter:" .. vote)
        return true, voter
    elseif func == "proposal" then
        local proposal = args[1]
        local proposaler = L0.GetState("proposal:" .. proposal)
        return true, proposaler
    elseif func == "max" then
        local proposal = nil
        for k, v in ipairs(proposaler) do

          local proposaler = L0.GetState("proposal:" .. v)
          if proposal == nil or proposaler.voteCount > proposal.voteCount then 
            proposal = proposaler
          end
        end
        return true, proposal
    end
    return false,"not fund func"..func
end

-- 可投票人
local voters = { "张三", "李四", "王五", "赵六", "孙七", "周八", "吴九", "郑十"}
-- 可候选人
local proposals = {"秦皇岛", "大连", "三亚"}

--投票人结构
voter = {
    name = "", -- 投票人名字
    vote = "", -- 投票候选人
    voteTime = 0, -- 投票时间
}

function voter:new(name, vote)
    o = {}
    setmetatable(o, self)
    self.__index = self
    self.name = name
    self.vote = vote
    return o
end

--候选人结构
proposal = {
    name = "", -- 候选人名字
    voteCount = 0, -- 候选人中票数
}

function proposal:new(name)
    o = {}
    setmetatable(o, self)
    self.__index = self
    self.name = name
    return o
end