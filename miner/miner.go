package miner

import (
	"fmt"
	"strings"
	"time"

	"github.com/richardcase/dockercoinsgo/cache"
	hc "github.com/richardcase/dockercoinsgo/hasher/client/grpc"
	rc "github.com/richardcase/dockercoinsgo/rng/client/grpc"
)

type Miner struct {
	IntervalInSeconds int
	RedisUrl          string
	HasherUrl         string
	RngUrl            string
	redisPool         cache.CacheStore
}

func (m Miner) Mine() error {

	// Open connection to redis
	m.redisPool = cache.NewRedisCache(m.RedisUrl, "", cache.FOREVER)

	tickerChan := time.NewTicker(1 * time.Second).C

	workDoneChan := make(chan bool)
	go func() {
		for {
			time.Sleep(time.Millisecond * 100)
			minedData, err := m.getRandomBytes()
			if err != nil {
				fmt.Printf("WARN: Error getting random bytes: %v\n", err)
				continue
			}
			hashedData, err := m.hashData(minedData)
			if err != nil {
				fmt.Printf("WARN: Error hashing data: %v\n", err)
				continue
			}

			if strings.HasPrefix(hashedData, "123") {
				fmt.Println("Coin found")
				m.redisPool.HashSet("wallet", []string{minedData, hashedData}, cache.FOREVER)
			}

			workDoneChan <- true
		}
	}()

	var workCount = 0
	for {
		select {
		case <-tickerChan:
			fmt.Printf("%d units of work done, updating hash counter\n", workCount)
			//TODO increment hash counter
			m.redisPool.Increment("hashes", uint64(workCount))

			workCount = 0

		case <-workDoneChan:
			workCount++
		}
	}
}

func (m Miner) getRandomBytes() (string, error) {
	client := rc.GrpcRngClient{ServerAddress: m.RngUrl}
	data, err := client.GenerateRandomString(32)
	if err != nil {
		return "", err
	}

	return data, nil
}

func (m Miner) hashData(data string) (string, error) {
	client := hc.GrpcHasherClient{ServerAddress: m.HasherUrl}
	hashed, err := client.Hash(data)
	if err != nil {
		return "", err
	}

	return hashed, nil
}
