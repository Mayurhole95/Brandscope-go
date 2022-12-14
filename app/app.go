package app

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/golang-boilerplate/config"
	"go.uber.org/zap"
)

var (
	db     *sqlx.DB
	logger *zap.SugaredLogger
)

func Init() {
	InitLogger()

	err := initDB()
	if err != nil {
		panic(err)
	}
}

func InitLogger() {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	logger = zapLogger.Sugar()
}

func initDB() (err error) {
	dbConfig := config.Database()

	fmt.Printf("%+v", dbConfig)
	fmt.Printf("Driver : ", dbConfig.Driver())
	fmt.Printf("Conn URL", dbConfig.ConnectionURL())
	db, err = sqlx.Open(dbConfig.Driver(), dbConfig.ConnectionURL())
	if err != nil {
		fmt.Println("Error : ", err.Error())
		return
	}

	if err = db.Ping(); err != nil {
		fmt.Println("Error : ", err.Error())

		return
	}

	db.SetMaxIdleConns(dbConfig.MaxPoolSize())
	db.SetMaxOpenConns(dbConfig.MaxOpenConns())
	db.SetConnMaxLifetime(time.Duration(dbConfig.MaxLifeTimeMins()) * time.Minute)

	return
}

func GetDB() *sqlx.DB {
	return db
}

func GetLogger() *zap.SugaredLogger {
	return logger
}

func Close() {
	logger.Sync()
	db.Close()
}
