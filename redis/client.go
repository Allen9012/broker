package redis

import (
	"github.com/Allen9012/broker"
	"github.com/go-redis/redis/v8"
)

type ClusterClient struct {
	*broker.BaseComponent
	*redis.ClusterClient
}

func NewClusterClient(conf *Config) *ClusterClient {
	c := &ClusterClient{
		BaseComponent: broker.NewBaseComponent(),
		ClusterClient: redis.NewClusterClient(&redis.ClusterOptions{
			Addrs: conf.Addrs,

			//Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
			// To route commands by latency or randomly, enable one of the following.
			//RouteByLatency: true,
			//RouteRandomly: true,
		}),
	}
	return c
}
