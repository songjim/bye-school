// DataSource定义各种数据源

package utiles

import (
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql" // only use form gorm
	"github.com/jinzhu/gorm"
	"time"
)

// DB 暴走漫画mysql
var (
	DB         *gorm.DB
	Pool       *redis.Pool
	SubPool    *redis.Pool
	Redis1Poll *redis.Pool
	GoEnv      string
)

func init() {
	dbPath := getEnv("MYSQL_HOST", "127.0.0.1:3306")
	dbUser := getEnv("MYSQL_USER", "root")
	dbPass := getEnv("MYSQL_PASSWORD", "")
	DB = initMysql(dbPath, dbUser, dbPass)

	GoEnv = getEnv("GO_ENV", "development")

	// redis2 umem
	redisServer := getEnv("REDIS_HOST", "redis://127.0.0.1:6379/0")
	Pool = initRedis(redisServer)

	// subtitle redis - umem
	subRedisServer := getEnv("SUB_REDIS_HOST", "redis://127.0.0.1:6379/0")
	SubPool = initRedis(subRedisServer)

	redis1Server := getEnv("REDIS1_HOST", "redis://127.0.0.1:6379/0")
	Redis1Poll = initRedis(redis1Server)
}

func initRedis(url string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(url)
			if err != nil {
				return nil, err
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func initMysql(host, user, pass string) *gorm.DB {
	tdb, err := gorm.Open("mysql", user+":"+pass+"@tcp("+host+")/bye_school?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	tdb.DB().SetMaxIdleConns(10)
	tdb.DB().SetMaxOpenConns(100)

	return tdb
}
