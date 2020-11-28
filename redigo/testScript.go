package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func testScript()  {
	var conn,err = redis.Dial("tcp","127.0.0.1:6379")
	if err != nil{
		panic(err)
	}

	var m = `-- KEYS CITY_LIMIT CITY_ID DAILY_LIMIT_KEY

-- 获取当前该城市的每日保养限制
local limitNum = redis.call("HMGET",KEYS[1],KEYS[2])
-- 如果未设置则返回
if limitNum == nil then
	return "city_limit_not_found"
end

local  limitNum = tonumber(limitNum)

--获取该城市的当日保养数
local dayCount = redis.call("HMGET", KEYS[1], KEYS[2]);
if dayCount == nil then
	dayCount = 0
end

dayCount = tonumber(dayCount)
if dayCount + 1 <= limitNum then
	redis.call("HINCRBY", KEYS[3],KEYS[2],1) return OK 
end

return LIMITED`
	script := redis.NewScript(3,m)
	s,err := redis.String(	script.Do(conn, "cityLimit", "xian", "2020-01-01"))
	if err != nil{
		panic(err)
	}

	fmt.Print(s)
}
