package belajar_golang_redis

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0,
})

func TestConnection(t *testing.T) {
	assert.NotNil(t, client)

	err := client.Close()
	assert.Nil(t, err)
}

var ctx = context.Background()

func TestPing(t *testing.T) {
	result, err := client.Ping(ctx).Result()
	assert.Nil(t, err)
	assert.Equal(t, "PONG", result)
}

func TestString(t *testing.T) {
	client.SetEx(ctx, "name", "Ucup Surucup", 3*time.Second)

	result, err := client.Get(ctx, "name").Result()
	assert.Nil(t, err)
	assert.Equal(t, "Ucup Surucup", result)

	time.Sleep(5 * time.Second)

	result, err = client.Get(ctx, "name").Result()
	assert.NotNil(t, err)
	assert.Equal(t, "", result)
}

func TestList(t *testing.T) {
	client.RPush(ctx, "names", "Eko")
	client.RPush(ctx, "names", "Kurniawan")
	client.RPush(ctx, "names", "Khannedy")

	assert.Equal(t, "Eko", client.LPop(ctx, "names").Val())
	assert.Equal(t, "Kurniawan", client.LPop(ctx, "names").Val())
	assert.Equal(t, "Khannedy", client.LPop(ctx, "names").Val())

	client.Del(ctx, "names")
}

func TestSet(t *testing.T) {
	client.SAdd(ctx, "students", "Eko")
	client.SAdd(ctx, "students", "Eko")
	client.SAdd(ctx, "students", "Kurniawan")
	client.SAdd(ctx, "students", "Kurniawan")
	client.SAdd(ctx, "students", "Khannedy")
	client.SAdd(ctx, "students", "Khannedy")

	assert.Equal(t, int64(3), client.SCard(ctx, "students").Val())
	assert.Equal(t, []string{"Eko", "Kurniawan", "Khannedy"}, client.SMembers(ctx, "students").Val())
}

func TestSortedSet(t *testing.T) {
	client.ZAdd(ctx, "scores", redis.Z{Score: 100, Member: "Eko"})
	client.ZAdd(ctx, "scores", redis.Z{Score: 85, Member: "Budi"})
	client.ZAdd(ctx, "scores", redis.Z{Score: 95, Member: "Joko"})

	assert.Equal(t, []string{"Budi", "Joko", "Eko"}, client.ZRange(ctx, "scores", 0, -1).Val())

	assert.Equal(t, "Eko", client.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "Joko", client.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "Budi", client.ZPopMax(ctx, "scores").Val()[0].Member)
}

func TestHash(t *testing.T) {
	client.HSet(ctx, "user:1", "id", "1")
	client.HSet(ctx, "user:1", "name", "Ucup")
	client.HSet(ctx, "user:1", "email", "ucup@example.com")

	user := client.HGetAll(ctx, "user:1").Val()

	assert.Equal(t, "1", user["id"])
	assert.Equal(t, "Ucup", user["name"])
	assert.Equal(t, "ucup@example.com", user["email"])

	client.Del(ctx, "user:1")
}

func TestGeoPoint(t *testing.T) {
	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name:      "Toko A",
		Longitude: 107.5563043368091,
		Latitude:  -6.930552777104311,
	})
	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name:      "Toko B",
		Longitude: 107.53965318309565,
		Latitude:  -6.944611150396564,
	})

	distance := client.GeoDist(ctx, "sellers", "Toko A", "Toko B", "km").Val()
	assert.Equal(t, 2.4135, distance)

	sellers := client.GeoSearch(ctx, "sellers", &redis.GeoSearchQuery{
		Longitude:  107.55029618856199,
		Latitude:   -6.935835366593564,
		Radius:     5,
		RadiusUnit: "km",
	}).Val()

	assert.Equal(t, []string{"Toko B", "Toko A"}, sellers)
}

func TestHyperLogLog(t *testing.T) {
	client.PFAdd(ctx, "visitors", "eko", "kurniawan", "khannedy")
	client.PFAdd(ctx, "visitors", "eko", "budi", "joko")
	client.PFAdd(ctx, "visitors", "rully", "budi", "joko")

	total := client.PFCount(ctx, "visitors").Val()
	assert.Equal(t, int64(6), total)
}

func TestPipeline(t *testing.T) {
	_, err := client.Pipelined(ctx, func(pipeliner redis.Pipeliner) error {
		pipeliner.SetEx(ctx, "name", "Ucup", 5*time.Second)
		pipeliner.SetEx(ctx, "address", "Indonesia", 5*time.Second)

		return nil
	})
	assert.Nil(t, err)

	assert.Equal(t, "Ucup", client.Get(ctx, "name").Val())
	assert.Equal(t, "Indonesia", client.Get(ctx, "address").Val())
}

func TestTransaction(t *testing.T) {
	_, err := client.TxPipelined(ctx, func(pipeliner redis.Pipeliner) error {
		pipeliner.SetEx(ctx, "name", "Otong", 5*time.Second)
		pipeliner.SetEx(ctx, "address", "Cirebon", 5*time.Second)

		return nil
	})
	assert.Nil(t, err)

	assert.Equal(t, "Otong", client.Get(ctx, "name").Val())
	assert.Equal(t, "Cirebon", client.Get(ctx, "address").Val())
}

func TestPublishStream(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := client.XAdd(ctx, &redis.XAddArgs{
			Stream: "members",
			Values: map[string]any{
				"name":    "Ucup",
				"address": "Indonesia",
			},
		}).Err()
		assert.Nil(t, err)
	}
}

func TestCreateConsumerGroup(t *testing.T) {
	client.XGroupCreate(ctx, "members", "group-1", "0")
	client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-1")
	client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-2")
}

func TestConsumeStream(t *testing.T) {
	streams := client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    "group-1",
		Consumer: "consumer-1",
		Streams:  []string{"members", ">"},
		Count:    2,
		Block:    5 * time.Second,
	}).Val()

	for _, stream := range streams {
		for _, message := range stream.Messages {
			fmt.Println(message.ID)
			fmt.Println(message.Values)
		}
	}
}

func TestSubscribePubSub(t *testing.T) {
	subscriber := client.Subscribe(ctx, "channel-1")
	defer subscriber.Close()

	for i := 0; i < 10; i++ {
		message, err := subscriber.ReceiveMessage(ctx)
		assert.Nil(t, err)
		fmt.Println(message.Payload)
	}
}

func TestTestPublishPubSub(t *testing.T) {
	for i := 10; i < 10; i++ {
		err := client.Publish(ctx, "channel-1", "Hello "+strconv.Itoa(i))
		assert.Nil(t, err)
	}
}
