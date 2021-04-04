package provider

import (
	"github.com/go-redis/redis/v8"

	"github.com/tariqc80/oui-challenge/cmd/gqlgen/graph/model"
	"github.com/tariqc80/oui-challenge/internal/config"
)

//Redis struct for Redis provider
type Redis struct {
	c *redis.Client
}

// NewRedis creates a redis client and returns a new redis provider
func NewRedis(c *config.Config) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     c.RedisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &Redis{
		c: client,
	}
}

// StoreSet stores a Set object in cache
func (r *Redis) StoreSet(s *model.Set) error {
	// serialize object and store in redis in a hash keyed by the id
	// HSET sets {s.id} {serialized object data}
	return nil
}

// RetrieveSet gets a Set object from cache by the given id
func (r *Redis) RetrieveSet(id int64) (*model.Set, error) {
	// fetch set from hash and unserialize
	// HGET sets {s.id}
	return nil, nil
}

func (r *Redis) RetrieveSetCollection() ([]*model.Set, error) {
	// fetch enitre hash iterate over each and unserialize
	// HGETALL sets
	return nil, nil
}

func (r *Redis) StoreSetCollection(sets []*model.Set) error {
	// serialize object data and store all in the hash
	// HMSET sets {s1.id} {s1 serialized data}  {s2.id} {s2 serialized data}  {s3.id} {s3 serialized data} ...
	return nil
}
