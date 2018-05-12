package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const (
	//ConfigFile ...
	ConfigFile = "./app/config.json"
)

//Config type defines the application configuration variables.
type Config struct {
	Keys  Keys  `json:"keys"`
	DB    DB    `json:"db"`
	Cache Cache `json:"cache"`
	Store Store `json:"store"`
}

//Keys defines the application's rsa key location/path.
type Keys struct {
	SignKey   string `json:"privkey"`
	VerifyKey string `json:"pubkey"`
}

//DB defines the application's database connection information(host, port, username,...).
type DB struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

//Cache defines the cache database connection information(host, port, username, ...).
type Cache struct {
	Host     string `json:"host"`
	Password string `json:"password"`
}

//Store defines the object store connection information.
type Store struct {
	Host      string `json:"host"`
	AccessKey string `json:"access-key"`
}

//NewAppConfig takes a file and returns a new configuration type
func NewAppConfig(file []byte) *Config {
	var config *Config
	err := json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

//LoadConfig read the configuration from the configuration file
func LoadConfig(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	//r := bytes.NewReader(b)

	return b, nil
}
