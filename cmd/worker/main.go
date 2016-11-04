package main

import (
	"flag"

	"github.com/richardcase/dockercoinsgo/miner"
)

func main() {
	hasherAddr := flag.String("hash-addr", "localhost:50051", "The published hostname of the service")
	clientCert := flag.String("cert", "", "[Required]. Path to the certificate for the file")
	clientKey := flag.String("key", "", "[Required]. Path to the certificate key")
	clientCA := flag.String("ca", "", "[Required]. Path to the CA file")
	miningDelay := flag.Int("delay", 100, "Delay in milliseconds between mining iterations. 100ms by default.")
	flag.Parse()

	miner := miner.Miner{
		IntervalInSeconds: 1,
		RedisUrl:          "localhost:6379",
		HasherUrl:         *hasherAddr,
		RngUrl:            "localhost:50052",
		ClientCertPath:    *clientCert,
		ClientCertKeyPath: *clientKey,
		ClientCAPath:      *clientCA,
		MiningDelay:       *miningDelay,
	}

	err := miner.Mine()
	if err != nil {
		panic(err)
	}
}
