package redispkg

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func Trigger() {
	// NewClient returns a client to the Redis Server specified by Options.
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	ExampleClient(ctx, rdb)

	v, err := CustomCmd(ctx, rdb, "custom_cmd").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", v)

	v1, err := rdb.Do(ctx, "get", "custom_cmd").Text()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", v1)

	val, err := rdb.Get(ctx, "custom_cmd").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key is custom_cmd, value is ", val)
}

/**
For redis usual commands.
$brew cleanup redis

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

func ExampleClient(ctx context.Context, rdb *redis.Client) {
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

// In go redis, it has lots of similar commands as NewStringCmd can be used to create new commands.

func CustomCmd(ctx context.Context, rdb *redis.Client, key string) *redis.StringCmd {
	cmd := redis.NewStringCmd(ctx, "get", key)
	err := rdb.Set(ctx, key, "custom command value", 0).Err()
	if err != nil {
		panic(err)
	}
	_ = rdb.Process(ctx, cmd)
	return cmd
}
