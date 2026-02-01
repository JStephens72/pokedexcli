package pokeapi

import (
	"net/http"
	"time"

	"github.com/JStephens72/pokedexcli/internal/pokecache"
)

// Client
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

// NewClient
func NewClient(timeout time.Duration) Client {
	cache := pokecache.NewCache(5 * time.Second)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: cache,
	}
}
