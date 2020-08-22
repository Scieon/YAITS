package main

import (
	"database/sql"
	"fmt"
	"github.com/YAITS/api/server"
	"github.com/YAITS/api/persistence"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"os"
)

const (
	defaultConfigPath = "./conf/conf.toml"
)

var logger *zap.SugaredLogger

func main() {
	err := readConfig(defaultConfigPath)
	logger := GetLogger()

	if err != nil {
		logger.Errorf("read config error: %s", err)
		return
	}

	ginPort := fmt.Sprintf(":%d", viper.GetInt64("server.port"))

	storage, err := initDB()
	if err != nil {
		logger.Errorf("error initializing database: %s", err.Error())
		os.Exit(1)
	}

	apiServer := server.NewServer(ginPort, logger, storage)

	// start server
	if err := apiServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal(err.Error())
		panic(err)
	}
}

func GetLogger() *zap.SugaredLogger {
	if logger == nil {
		l, _ := zap.NewDevelopment()
		logger = l.Sugar()
	}

	return logger
}

func readConfig(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	viper.SetConfigType("toml")

	viper.SetDefault("server.host", "127.0.0.1")
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "3306")
	return viper.ReadConfig(f)
}

func initDB() (*persistence.MysqlStorage, error) {
	host := viper.GetString("db.host")
	password := viper.GetString("db.password")
	user := viper.GetString("db.user")
	dbName := viper.GetString("db.database")

	conConfig := mysql.Config{
		User:                 user,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 host,
		DBName:               dbName,
		MaxAllowedPacket:     0,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := sql.Open("mysql", conConfig.FormatDSN())

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	dbStorage := persistence.NewMysqlStorage(db)

	return dbStorage, nil
}
