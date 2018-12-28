package util

import (
	"fmt"
	"log"
	//	"strings"
	"time"

	"github.com/go-redis/redis"
)

//type RedisClient struct {
//	client redis.Client
//}

var Cli = new(redis.Client)

const (
	Addr     = "127.0.0.1"
	Port     = "6379"
	Password = ""
	DB       = 0

	REDIS_CODE_TIMEOUT = 60 * 60 * 2
	REDIS_CODE_PREFIX  = "OAUTH_CODE_"
)

//TODO panic when client is nil
func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     Addr + ":" + Port,
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong, err)

	Cli = client
}

//var Redis = NewRedisClient(Addr, Password, DB)

//func NewRedisClient(Addr string, Password string, DB int) *RedisClient {
//	c := redis.NewClient(&redis.Options{
//		Addr:     "127.0.0.1" + ":" + "6379",
//		Password: "",
//		DB:       0,
//	})

//	pong, err := c.Ping().Result()
//	fmt.Println(pong, err)
//	return &RedisClient{c}
//	//	return &RedisClient{client}
//}

func RedisSave(client_id string, data []byte, time_out time.Duration) {
	key := REDIS_CODE_PREFIX + client_id
	//	data, _ := json.Marshal(value)
	fmt.Println(string(data))
	//	Second Millisecond
	err := Cli.Set(key, data, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func RedisRemove(client_id string) {
	key := REDIS_CODE_PREFIX + client_id
	err := Cli.Del(key).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func RedisGet(client_id string) (string, string, error) {
	key := REDIS_CODE_PREFIX + client_id
	value, err := Cli.Get(key).Result()
	if err != nil {
		fmt.Println(err)
	}
	if value == "" {
		return "", "", nil
	}

	//	var data interface{}
	//	err = json.Unmarshal([]byte(value), data)
	//	if err != nil {
	//		fmt.Println(err)
	//	}

	//	tmps := strings.Split(value, ",")
	return client_id, value, nil
}
