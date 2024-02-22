package initialize

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/go-redis/redis"
	"go-backend/global"
)

func InitRedis() {
	addr := fmt.Sprintf("%s:%d", global.GlobConfig.Redis.Host, global.GlobConfig.Redis.Port)
	// 生成redis客户端
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: global.GlobConfig.Redis.Password, // no password set
		DB:       global.GlobConfig.Redis.DB,       // use default DB
	})
	// 链接redis
	_, err := global.Redis.Ping().Result()
	if err != nil {
		color.Red("[InitRedis] connect redis error:\n" + err.Error())
	} else {
		color.Blue("Redis inits successfully\n")
	}
}
