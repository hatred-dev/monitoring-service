package configuration

import (
    "fmt"
)

type Config struct {
    mongoConfig       MongoConfig
    telegramConfig    TelegramConfig
    applicationConfig AppConfiguration
    gamecatalogConfig GameCtalogConfiguration
}

var config *Config

func GetMongoConfig() MongoConfig {
    return config.mongoConfig
}

func GetTelegramConfig() TelegramConfig {
    return config.telegramConfig
}

func GetApplicationConfig() AppConfiguration {
    return config.applicationConfig
}

type MongoConfig struct {
    MongoHost string
    MongoPort string
    MongoUser string
    MongoPass string
}

func (db MongoConfig) DatabaseUrl() string {
    return fmt.Sprintf("mongodb://%s:%s@%s:%s",
        db.MongoUser,
        db.MongoPass,
        db.MongoHost,
        db.MongoPort,
    )
}

func initDBConfig() MongoConfig {
    prefix := "mongo_"
    return MongoConfig{
        MongoHost: getEnvOrDefault(prefix, "host", str("localhost")),
        MongoPort: getEnvOrDefault(prefix, "port", str("27017")),
        MongoUser: getEnvOrDefault(prefix, "user", nil),
        MongoPass: getEnvOrDefault(prefix, "password", nil),
    }
}

type TelegramConfig struct {
    BotToken string
    ChatId   string
}

func initTelegramConfig() TelegramConfig {
    return TelegramConfig{
        BotToken: getEnvOrDefault("", "bot_token", nil),
        ChatId:   getEnvOrDefault("", "chat_id", nil),
    }
}

type AppConfiguration struct {
    RootPrefix string
    ApiKey     string
}

func initAppConfiguration() AppConfiguration {
    return AppConfiguration{
        RootPrefix: getEnvOrDefault("", "root_prefix", str("")),
        ApiKey:     getEnvOrDefault("", "api_key", str("")),
    }
}

type GameCtalogConfiguration struct {
    Url string
}

func initGameCatalogConfig() GameCtalogConfiguration {
    return GameCtalogConfiguration{Url: getEnvOrDefault("", "game_catalog_url", nil)}
}

func InitConfigurations() *Config {
    config = &Config{
        mongoConfig:       initDBConfig(),
        telegramConfig:    initTelegramConfig(),
        applicationConfig: initAppConfiguration(),
        gamecatalogConfig: initGameCatalogConfig(),
    }

    return config
}
