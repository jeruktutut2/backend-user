package configuration

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Database struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"datanase"`
	Port     int    `yaml:"port"`
}

type Timezone struct {
	Timezone string `yaml:"timezone"`
}

type Configuration struct {
	Database Database `yaml:"database"`
	Timezone Timezone `yaml:"timezone"`
}

func NewConfiguration() (configuration Configuration) {
	file, err := ioutil.ReadFile("configuration/configuration.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// configuration := Configuration{}
	err = yaml.Unmarshal(file, &configuration)
	return configuration
}
