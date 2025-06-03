package pokecache

import (
	"sync"
	"time"
)



type Cache struct{
	cacheEntry map[string]cacheEntry
	mu sync.Mutex
}

type cacheEntry struct{
	createdAt time.Time
	val []byte
}


func NewCache(key string, val []byte){
	newCache = 
}

func cache.Add(){

}

func cache.Get(){

}

func cache.reapLoop(){

}