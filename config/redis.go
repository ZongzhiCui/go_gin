package config

type redisConf struct {
	Addr string
	//Addr:     "localhost:6379",
	Password string // no password set
	DB       int    // use default DB
}

var RedisConf = func() *redisConf {
	return &redisConf{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}
}
