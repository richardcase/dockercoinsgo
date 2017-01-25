package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/richardcase/dockercoinsgo/cache"
	"gopkg.in/gin-gonic/gin.v1"
)

type Summary struct {
	Coins  string `json:"coins"`
	Hashes int    `json:"hashes"`
	Now    int64  `json:"now"`
}

func main() {
	staticPath := flag.String("static-path", "", "[Required] Path to the static web contnet")
	cacheAddress := flag.String("cache-addr", "localhost:6379", "The published address of the cache")
	cachePassword := flag.String("cache-pwd", "", "Password to use when connecting to the cache")

	flag.Parse()

	if *staticPath == "" {
		fmt.Println("You must supply the path to the static web content")
		os.Exit(1)
	}

	if *cacheAddress == "" {
		fmt.Println("You must supply the cache address")
		os.Exit(1)
	}

	coinsCache, _ := cache.NewRedisCache(*cacheAddress, *cachePassword)

	router := gin.Default()

	router.GET("/json", func(c *gin.Context) {
		val, err := coinsCache.GetInt("hashes")
		if err != nil {
			log.Printf("Error getting hashes from the cache: %v\n", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		t := time.Now()
		summary := Summary{Coins: "", Hashes: val, Now: t.Unix()}
		c.JSON(http.StatusOK, summary)
	})

	test := path.Join(*staticPath, "/index.html")
	fmt.Println(test)

	router.StaticFile("/", path.Join(*staticPath, "/index.html"))
	router.StaticFile("/index.html", path.Join(*staticPath, "/index.html"))
	router.StaticFile("/d3.min.js", path.Join(*staticPath, "/d3.min.js"))
	router.StaticFile("/jquery-1.11.3.min.js", path.Join(*staticPath, "/jquery-1.11.3.min.js"))
	router.StaticFile("/rickshaw.min.css", path.Join(*staticPath, "/rickshaw.min.css"))
	router.StaticFile("/rickshaw.min.js", path.Join(*staticPath, "/rickshaw.min.js"))

	router.Run(":8000")
}
