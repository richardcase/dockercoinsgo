package main

import (
	"github.com/richardcase/dockercoinsgo/miner"
)

func main() {
	miner := miner.Miner{
		IntervalInSeconds: 1,
		RedisUrl:          "localhost:6379",
		HasherUrl:         "localhost:50051",
		RngUrl:            "localhost:50052"}

	err := miner.Mine()
	if err != nil {
		panic(err)
	}
}
