package main

import "encoding/json"
import "os"

import "github.com/kdgregory/gocommons"


var DEFAULT_CONFIG_FILE = gocommons.PrependHomeDir(".s34go.ini")

type Config struct {
    PublicKey string
    SecretKey string
}


func ReadConfigFile(configPath string) (Config, error) {
    exists,err := gocommons.FileExists(configPath)
    if !exists {
        return Config{}, err
    }

    configFile,err := os.Open(configPath)
    if err != nil {
        return Config{}, err
    }
    defer configFile.Close()

    decoder := json.NewDecoder(configFile)
    config := Config{}
    err = decoder.Decode(&config)
    return config,err
}


// updates an existing configuration object, replacing items with set values from the
// passed configuration
func (config *Config) Merge(updates Config) {

    config.PublicKey = gocommons.DefaultIfBlank(updates.PublicKey, config.PublicKey)
    config.SecretKey = gocommons.DefaultIfBlank(updates.SecretKey, config.SecretKey)
}
