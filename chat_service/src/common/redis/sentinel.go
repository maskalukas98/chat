package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

type Sentinel struct {
	masterName        string
	sentinelAddress   string
	redisMasterClient *redis.Client
}

func NewSentinel(masterName string, sentinelAddress string) *Sentinel {
	sentinel := &Sentinel{
		masterName:      masterName,
		sentinelAddress: sentinelAddress,
	}

	sentinel.connectToMaster(sentinel.getMasterAddress())

	return sentinel
}

func (r *Sentinel) GetRedisClient() *redis.Client {
	return r.redisMasterClient
}

func (r *Sentinel) getMasterAddress() string {
	client := redis.NewSentinelClient(&redis.Options{
		Addr:     r.sentinelAddress,
		Password: "",
	})

	defer client.Close()

	ctx := context.Background()
	masterAddr, err := client.GetMasterAddrByName(ctx, r.masterName).Result()
	if err != nil {
		log.Fatal("Could not fetch master redis address.", err)
	}

	return masterAddr[0] + ":" + masterAddr[1]
}

func (r *Sentinel) connectToMaster(redisAddress string) {
	r.redisMasterClient = redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "",
	})
}
