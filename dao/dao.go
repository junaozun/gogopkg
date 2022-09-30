package dao

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/junaozun/gogopkg/config"
	"github.com/junaozun/gogopkg/etcdx"
	"github.com/junaozun/gogopkg/redisx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Dao 数据对象访问
type Dao struct {
	DB    *gorm.DB // db
	Redis redisx.IClient
	Etcd  etcdx.IClient
}

// NewDao 构造
func NewDao(daoConfig []interface{}) (*Dao, error) {
	dao := new(Dao)
	for _, cfg := range daoConfig {
		if mysqlCfg, ok := cfg.(*config.MysqlConfig); ok {
			db, err := NewDB(mysqlCfg)
			if err != nil {
				return nil, err
			}
			dao.DB = db
		} else if redisCfg, ok := cfg.(*config.RedisConfig); ok {
			redis, err := NewRedis(redisCfg)
			if err != nil {
				return nil, err
			}
			dao.Redis = redis
		} else if etcdCfg, ok := cfg.(*config.EtcdConfig); ok {
			etcd, err := NewEtcd(etcdCfg)
			if err != nil {
				return nil, err
			}
			dao.Etcd = etcd
		}
	}
	return dao, nil
}

func NewDB(mysqlConfig *config.MysqlConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.User,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Name)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			// 将标准输出作为Writer
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				// 设定慢查询时间阈值为10ms
				SlowThreshold: 10 * time.Microsecond,
				// 设置日志级别，只有Warn和Info级别会输出慢查询日志
				LogLevel: logger.Warn,
			},
		),
	})
}

func NewRedis(redisConfig *config.RedisConfig) (redisx.IClient, error) {
	return redisx.NewClient(redisx.Config{
		Server: redisConfig.Server,
		Index:  redisConfig.Index,
		Auth:   redisConfig.Auth,
	})
}

func NewEtcd(etcdConfig *config.EtcdConfig) (etcdx.IClient, error) {
	return etcdx.NewClientWithConfig(etcdx.Config{
		Servers:        etcdConfig.Servers,
		DialTimeout:    etcdConfig.DialTimeout,
		RequestTimeout: etcdConfig.RequestTimeout,
	})
}
