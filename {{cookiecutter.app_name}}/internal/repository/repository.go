package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository struct {
	db  *gorm.DB
	rdb *redis.Client
	oss *minio.Client
}

func NewRepository(db *gorm.DB, rdb *redis.Client, oss *minio.Client) *Repository {
	return &Repository{
		db:  db,
		rdb: rdb,
		oss: oss,
	}
}

func NewMinIO(conf *viper.Viper) *minio.Client {

	addr := conf.GetString("db.minio.addr")
	opt := &minio.Options{
		Creds: credentials.NewStaticV4(
			conf.GetString("db.minio.user"),
			conf.GetString("db.minio.password"),
			""),
		Secure: conf.GetBool("db.minio.ssl"),
	}

	minioClient, err := minio.New(addr, opt)
	if err != nil {
		panic(err)
	}
	return minioClient
}

func NewDB(conf *viper.Viper) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.GetString("db.mysql.user"),
		conf.GetString("db.mysql.password"),
		conf.GetString("db.mysql.addr"),
		conf.GetString("db.mysql.name"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
func NewRedis(conf *viper.Viper) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.GetString("db.redis.addr"),
		Password: conf.GetString("db.redis.password"),
		DB:       conf.GetInt("db.redis.db"),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("redis error: %s", err.Error()))
	}

	return rdb
}

func (r *Repository) RGet(ctx context.Context, key string, dest interface{}) error {
	buf, err := r.rdb.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, dest)
}

func (r *Repository) RSet(ctx context.Context,
	key string,
	value interface{},
	expiration time.Duration) error {

	buf, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.rdb.Set(ctx, key, buf, expiration).Err()
}
