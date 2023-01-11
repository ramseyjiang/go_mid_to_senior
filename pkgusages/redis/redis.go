package redispkg

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

/**
For redis usual commands.

//Make redis works in the background
$brew services restart redis

$redis-server -v         //check redis version
Redis server v=6.2.6 sha=00000000:0 malloc=libc bits=64 build=c6f3693d1aced7d9

check redis is working or not
$redis-cli ping

If test the redis works or not in go, please run "redis server" first.
$redis-server

$ redis-cli
redis 127.0.0.1:6379> ping
PONG
redis 127.0.0.1:6379> set mykey somevalue
OK
redis 127.0.0.1:6379> get mykey
"somevalue"

$redis-cli shutdown
*/
func getRedisValues() (string, string, string) {
	// Every Redis command accepts a context that you can use to set timeouts or propagate some information
	ctx := context.Background()

	// NewClient returns a client to the Redis Server specified by Options. Password and DB can be removed, it is up to you.
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	// Check the connection
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to Redis!")

	err = client.Set(ctx, "customise-key", "customise-value", 0).Err()
	if err != nil {
		panic(err)
	}

	// the first way
	val, err := client.Get(ctx, "customise-key").Result()
	if err != nil {
		panic(err)
	}

	// the second way
	val1, err := client.Do(ctx, "get", "customise-key").Text()
	if err != nil {
		panic(err)
	}

	// try to get a key does not exist
	val2, err := client.Get(ctx, "key2").Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}

	return val, val1, val2
}
