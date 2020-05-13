package config

import (
  "github.com/spf13/viper"
  "log"
)

var c Config

func init() {
  viper.SetConfigType("yaml")
  viper.SetConfigName("config")
  viper.AddConfigPath("/mnt")
  viper.AddConfigPath(".")
}

func GetConfig() Config {
  return c
}


func Get(key string) interface{} {
  return viper.Get(key)
}

func GetBool(key string) bool {
  return viper.GetBool(key)
}

func GetInt(key string) int {
  return viper.GetInt(key)
}

func GetString(key string) string {
  return viper.GetString(key)
}

func GetStringMapString(key string) map[string]string {
  return viper.GetStringMapString(key)
}

func GetStringMap(key string) map[string]interface{} {
  return viper.GetStringMap(key)
}

func GetStringSlice(key string) []string {
  return viper.GetStringSlice(key)
}

func IsSet(key string) bool {
  return viper.IsSet(key)
}

func Set(key string, value interface{}) {
  viper.Set(key, value)
}

type ServeConfig struct {
  GrpcHostPort    string `mapstructure:"grpcHostPort"`
  RestfulHostPort string `mapstructure:"restfulHostPort"`
}

type CommonConfig struct {
  PoolFreeInterval uint32 `mapstructure:"pool_free_interval"`
  Version          string `mapstructure:"version"`
  StaticPath       string `mapstructure:"staticPath"`
}

type SSLConfig struct {
  SSLMode string `mapstructure:"sslmode"`
  SSLCert string `mapstructure:"sslcert"`
  SSLKey  string `mapstructure:"sslkey"`
}

type DBConfig struct {
  ReadHost  string            `mapstructure:"readhost"`
  ReadPort  int               `mapstructure:"readport"`
  WriteHost string            `mapstructure:"writehost"`
  WritePort int               `mapstructure:"writeport"`
  Username  string            `mapstructure:"username"`
  Password  string            `mapstructure:"password"`
  Database  string            `mapstructure:"database"`
}
type Config struct {
  Serve       ServeConfig    `mapstructure:"serve"`
  DbServer    DBConfig       `mapstructure:"postgresqlServer"`
  Common      CommonConfig   `mapstructure:"common"`
}

func SetConfigPath(path string) {
  viper.SetConfigFile(path)
}

func ReadConfig() {
  err := viper.ReadInConfig()

  if err != nil {
    log.Fatal(err.Error())
  }

  err = viper.Unmarshal(&c)

  if err != nil {
    log.Fatal(err)
  }
}
