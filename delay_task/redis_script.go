package delay_task

const delayQueuePushRedisScript = `

-- 添加readyTime到zset
local count = redis.call("zadd", KEYS[1], tonumber(ARGV[3]), ARGV[1])
-- 消息已经存在
if count == 0 then 
   return 0
end
-- 添加body到hash
redis.call("hsetnx", KEYS[2], ARGV[1], ARGV[2])
return 1
`

const delayQueueDelRedisScript = `

-- 删除zset和hash关于这条消息的内容
redis.call("zrem", KEYS[1], ARGV[1])
redis.call("hdel", KEYS[2], ARGV[1])
return 1
`
