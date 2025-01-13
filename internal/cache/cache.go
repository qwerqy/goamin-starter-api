package cache

import "github.com/redis/go-redis/v9"

type Storage struct {
	// Users interface {
	// 	Get(context.Context, int64) (*store.User, error)
	// 	Set(context.Context, *store.User) error
	// 	Delete(context.Context, int64)
	// }
}

func NewCacheStorage(rdb *redis.Client) Storage {
	return Storage{
		// Users: &UserStore{rdb: rdb},
	}
}
