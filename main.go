package main

import (
	"gopoke/internal/pokecache"
	"time"
)

func main() {
	cache := pokecache.NewCache(30 * time.Second)
	startRepl(cache)
}
