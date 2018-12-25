package data

import (
	"github.com/go-redis/redis"
)


type RedisClient struct{
	client redis.Client
}
//var Cli = new(redis.Client)

//TODO panic when client is nil
//func init() {
//	client := redis.NewClient(&redis.Options{
//		Addr:     "127.0.0.1" + ":" + "6379",
//		Password: "",
//		DB:       0,
//	})
//	Cli = client
//}
var Redis = NewRedisClient(Addr, Password, DB);

const (
	Addr := "127.0.0.1:6379"
	Password := ""
	DB := 0
)

func NewRedisClient(Addr string, Password string, DB int) *RedisClient{
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1" + ":" + "6379",
		Password: "",
		DB:       0,
	})
	return &RedisClient{client:client}
}

func (this *RedisClient)Save(clientId string, code string, timeOut string) {
	key := util.REDIS_CODE_PREFIX + code
	err := Cli.Set(key, clientId, time.Millisecond*timeOut).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func (this *RedisClient)Remove(code string, codePrefix string) {
	key := codePrefix + code
	err := Cli.Del(key).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func (this *RedisClient)Get(code string codePrefix string) (string, string, string) {
	key := codePrefix + code
	value, err := Cli.Get(key).Result()
	if err != nil {
		fmt.Println(err)
	}
	if value == "" {
		return "", "", ""
	}
	tmps := strings.Split(value, ",")
	return code, tmps[0], tmps[1]
}
