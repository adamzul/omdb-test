package config

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/joho/godotenv"
)

var ArangoDbConfig, ArangoPaperChainDbConfig, RedisConfig, MysqlConfig, MysqlGatewayConfig map[string]interface{}
var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
	}

	return &dbConfig
}

func DBURL(dbConfig *DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&locLocal",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DbName,
	)
}
func GetCon() *gorm.DB {
	return DB
}

func init() {
	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
	}
}
