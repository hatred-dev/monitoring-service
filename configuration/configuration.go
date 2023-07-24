package configuration

import (
	"fmt"
)

var (
	MDBConfig  *MongoConfig
	TGConfig   *TelegramConfig
	UptimeConf *UptimeConfiguration
	AppConf    *AppConfiguration
)

type MongoConfig struct {
	MongoHost string
	MongoPort string
	MongoUser string
	MongoPass string
}

func (db *MongoConfig) DatabaseUrl() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s",
		db.MongoUser,
		db.MongoPass,
		db.MongoHost,
		db.MongoPort,
	)
}

func initDBConfig() {
	prefix := "mongo_"
	MDBConfig = &MongoConfig{
		MongoHost: getEnvOrDefault(prefix, "host", "localhost"),
		MongoPort: getEnvOrDefault(prefix, "port", "27017"),
		MongoUser: getEnvOrDefault(prefix, "user", ""),
		MongoPass: getEnvOrDefault(prefix, "password", ""),
	}
}

type TelegramConfig struct {
	BotToken string
	ChatId   string
}

func initTelegramConfig() {
	TGConfig = &TelegramConfig{
		BotToken: getEnvOrDefault("", "bot_token", ""),
		ChatId:   getEnvOrDefault("", "chat_id", ""),
	}
}

type UptimeConfiguration struct {
	UptimeUrl string
}

func initUptimeConfig() {
	UptimeConf = &UptimeConfiguration{UptimeUrl: getEnvOrDefault("", "uptime_url", "")}
}

type AppConfiguration struct {
	RootPrefix string
	ApiKey     string
}

func initAppConfiguration() {
	AppConf = &AppConfiguration{
		RootPrefix: getEnvOrDefault("", "root_prefix", ""),
		ApiKey:     getEnvOrDefault("", "api_key", ""),
	}
}

func InitConfigurations() {
	initDBConfig()
	initTelegramConfig()
	initUptimeConfig()
	initAppConfiguration()
}
