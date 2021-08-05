package initialize

import (
	"github.com/BurntSushi/toml"
	"net/http"
)

type TomlConfig struct {
	Title string
	DB    database    `toml:"database"`
	HS    http.Server `toml:"httpserver"`
}

type database struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	ConnMax  int `toml:"connection_max"`
	Enabled  bool
}

func InitConfig(filepath string) (config TomlConfig) {
	if _, err := toml.DecodeFile(filepath, &config); err != nil {
		panic(err)
	}
	return
}
