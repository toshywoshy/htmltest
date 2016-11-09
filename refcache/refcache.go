package refcache

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type RefCache struct {
	refStore     map[string]CachedRef
	rwMutex      *sync.RWMutex
	cacheExpires time.Duration
}

func NewRefCache(storePath string, cacheExpiresStr string) *RefCache {
	rS := &RefCache{}
	_ = storePath
	rS.rwMutex = &sync.RWMutex{}
	rS.cacheExpires, _ = time.ParseDuration(cacheExpiresStr)

	if !rS.ReadStore(storePath) {
		rS.refStore = make(map[string]CachedRef)
	}

	return rS
}

func (rS *RefCache) ReadStore(storePath string) bool {
	// Read in RefCache
	f, err := os.Open(storePath)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println(err.Error())
			os.Exit(2)
		}
		return false
	}
	defer f.Close()

	var refStore map[string]CachedRef
	err = json.NewDecoder(f).Decode(&refStore)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	rS.refStore = refStore
	return true
}

func (rS *RefCache) WriteStore(storePath string) {
	// Write out RefCache
	os.Mkdir(path.Dir(storePath), 0777)
	f, err := os.Create(storePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(&rS.refStore)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}

type CachedRef struct {
	StatusCode int
	LastSeen   time.Time
	// Body byte[] // For when we do hash checking on external documents
}

func (rS *RefCache) Get(urlStr string) (*CachedRef, bool) {
	rS.rwMutex.RLock()
	val, ok := rS.refStore[urlStr]
	rS.rwMutex.RUnlock()
	if ok {
		// In cache, check if cache has expired
		if time.Now().Before(val.LastSeen.Add(rS.cacheExpires)) {
			// All ok!
			return &val, true
		} else {
			// Nope, cache has expired
			return nil, false
		}
	} else {
		return nil, false
	}
}

func (rS *RefCache) Save(urlStr string, statusCode int) {
	if !(statusCode == http.StatusPartialContent || statusCode == http.StatusOK) {
		// Don't cache failed results
		fmt.Println(statusCode)
		return
	}
	cR := CachedRef{
		StatusCode: statusCode,
		LastSeen:   time.Now(),
	}
	rS.rwMutex.Lock()
	rS.refStore[urlStr] = cR
	rS.rwMutex.Unlock()
}
