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
	ClientCertPath    string
	ClientCertKeyPath string
	ClientCAPath      string
	MiningDelay       int
	CachePassword     string
	cache             cache.CacheStore
}

func (m Miner) Mine() error {

	// Open connection to redis
	redisCache, err := cache.NewRedisCache(m.RedisUrl, m.CachePassword)
	if err != nil {
		return fmt.Errorf("Error creating cache: %v\n", err)
	}
	m.cache = redisCache

	tickerChan := time.NewTicker(1 * time.Second).C

	workDoneChan := make(chan bool)
	go func() {
		for {
			if m.MiningDelay > 0 {
				delay := time.Duration(m.MiningDelay) * time.Millisecond
				time.Sleep(delay)
			}

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
				m.cache.HashSet("wallet", []string{minedData, hashedData})
			}

			workDoneChan <- true
		}
	}()

	var workCount = 0
	for {
		select {
		case <-tickerChan:
			fmt.Printf("%d units of work done, updating hash counter\n", workCount)
			_, err := m.cache.Increment("hashes", uint64(workCount))
			if err != nil {
				fmt.Printf("Error storing in cache: %v\n", err)
			}

			workCount = 0

		case <-workDoneChan:
			workCount++
		}
	}
}

func (m Miner) getRandomBytes() (string, error) {
	client := rc.GrpcRngClient{ServerAddress: m.RngUrl, CertFile: m.ClientCertPath, KeyFile: m.ClientCertKeyPath, CAFile: m.ClientCAPath}
	data, err := client.GenerateRandomString(32)
	if err != nil {
		return "", err
	}

	return data, nil
}

func (m Miner) hashData(data string) (string, error) {
	client := hc.GrpcHasherClient{ServerAddress: m.HasherUrl, CertFile: m.ClientCertPath, KeyFile: m.ClientCertKeyPath, CAFile: m.ClientCAPath}
	hashed, err := client.Hash(data)
	if err != nil {
		return "", err
	}

	return hashed, nil
}
