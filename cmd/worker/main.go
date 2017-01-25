package main

import (
	"flag"

	"github.com/richardcase/dockercoinsgo/miner"
)

func main() {
	hasherAddr := flag.String("hash-addr", "localhost:50051", "The published address of the hasher service")
	rngAddr := flag.String("rng-addr", "localhost:50052", "The published address of the rng service")
	cacheAddr := flag.String("cache-addr", "localhost:6379", "The published address of the cache")
	clientCert := flag.String("cert", "", "[Required]. Path to the certificate for the file")
	clientKey := flag.String("key", "", "[Required]. Path to the certificate key")
	clientCA := flag.String("ca", "", "[Required]. Path to the CA file")
	miningDelay := flag.Int("delay", 100, "Delay in milliseconds between mining iterations. 100ms by default.")
	cachePassword := flag.String("cache-pwd", "", "Password to use when connecting to the cache")

	flag.Parse()

	miner := miner.Miner{
		IntervalInSeconds: 1,
		RedisUrl:          *cacheAddr,
		HasherUrl:         *hasherAddr,
		RngUrl:            *rngAddr,
		ClientCertPath:    *clientCert,
		ClientCertKeyPath: *clientKey,
		ClientCAPath:      *clientCA,
		MiningDelay:       *miningDelay,
		CachePassword:     *cachePassword,
	}

	err := miner.Mine()
	if err != nil {
		panic(err)
	}
}
