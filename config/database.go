package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	// Database Host
	Host string `env:"DB_HOST"`
	// Database Port
	Port string `env:"DB_PORT"`
	// Database Username
	Username string `env:"DB_USERNAME"`
	// Database Password
	Password string `env:"DB_PASSWORD"`
	// Database Name
	Name string `env:"DB_NAME"`
	// Database Charset
	Charset string `env:"DB_CHARSET"`
	// Database Collation
	Collation string `env:"DB_COLLATION"`
	// Database Timezone
	Timezone string `env:"DB_TIMEZONE"`
	// Database Debug
	Debug bool `env:"DB_DEBUG"`
	// Database Migration
	Migration bool `env:"DB_MIGRATION"`
	// Database Seeding
	Seeding bool `env:"DB_SEEDING"`
	// Database Connection
	Connection string `env:"DB_CONNECTION"`
	// Database Max Open Connection
	MaxOpenConnection int `env:"DB_MAX_OPEN_CONNECTION"`
	// Database Max Idle Connection
	MaxIdleConnection int `env:"DB_MAX_IDLE_CONNECTION"`
	// Database Max Life Time
	MaxLifeTime int `env:"DB_MAX_LIFE_TIME"`
}

// (db *Database) GetDSN is a function to get database DSN
func (db *Database) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&collation=%s&parseTime=true&loc=%s",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
		db.Charset,
		db.Collation,
		db.Timezone,
	)
}

// (db *Database) Init is a function to initialize connection to database
func (db *Database) Init() (*gorm.DB, error) {
	logCfg := logger.Config{
		SlowThreshold:             time.Second,  // Slow SQL threshold
		LogLevel:                  logger.Error, // Log level
		Colorful:                  false,        // Disable color
		ParameterizedQueries:      true,         // Enable parameterized queries
		IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
	}
	if db.Debug {
		logCfg.LogLevel = logger.Info
		logCfg.Colorful = true
		logCfg.ParameterizedQueries = false
		logCfg.IgnoreRecordNotFoundError = false
	}

	mysqlCfg := mysql.Config{
		DSN: db.GetDSN(),
	}
	gormCfg := &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logCfg,
		),
	}

	gormDB, err := gorm.Open(mysql.New(mysqlCfg), gormCfg)
	if err != nil {
		return nil, err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(db.MaxOpenConnection)
	sqlDB.SetMaxIdleConns(db.MaxIdleConnection)
	sqlDB.SetConnMaxLifetime(time.Duration(db.MaxLifeTime) * time.Second)

	return gormDB, nil
}
