package main

import (
	"fmt"
	"github.com/YAITS/api/server"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"go.uber.org/zap"
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

	apiServer := server.NewServer(ginPort, logger)

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
