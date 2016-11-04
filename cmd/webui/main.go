package main

import (
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
	staticPath := os.Getenv("DCKR_WEB_STATIC")
	if staticPath == "" {
		fmt.Println("You must set the DCKR_WEB_STATIC environment variable")
		os.Exit(1)
	}

	cacheAddress := os.Getenv("DCKR_CACHE_ADDR")
	if cacheAddress == "" {
		fmt.Println("You must set the DCKR_CACHE_ADDRC environment variable")
		os.Exit(1)
	}

	coinsCache, _ := cache.NewRedisCache(cacheAddress)

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

	test := path.Join(staticPath, "/index.html")
	fmt.Println(test)

	router.StaticFile("/", path.Join(staticPath, "/index.html"))
	router.StaticFile("/index.html", path.Join(staticPath, "/index.html"))
	router.StaticFile("/d3.min.js", path.Join(staticPath, "/d3.min.js"))
	router.StaticFile("/jquery-1.11.3.min.js", path.Join(staticPath, "/jquery-1.11.3.min.js"))
	router.StaticFile("/rickshaw.min.css", path.Join(staticPath, "/rickshaw.min.css"))
	router.StaticFile("/rickshaw.min.js", path.Join(staticPath, "/rickshaw.min.js"))

	router.Run(":8000")
}
