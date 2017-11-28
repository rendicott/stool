package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Verifier string `yaml:"verifier"`
	TestPath string `yaml:"testpath"`
}

type ConfigReader interface {
	Read(path string) ([]byte, error)
}

type FileConfigReader struct {
}

func (f FileConfigReader) Read(path string) ([]byte, error) {
	return readFile(path)
}

func readFile(path string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(path)
	return bytes, err

}

func FileExists(path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}
	return nil
}

func LoadConfig(path string, cfReader ConfigReader) (Config, error) {
	// read the file
	if err := FileExists(path); err != nil {
		fmt.Println("error in LoadConfig : could not find file" + path)
		return Config{}, err
	}
	payload, err := cfReader.Read(path)
	if err != nil {
		fmt.Println("error reading file at " + path)
		return Config{}, err
	}
	config, err := ParseConfigFile(payload)
	if err != nil {
		fmt.Println("error parsing config file ")
		return Config{}, err
	}
	// parse the contents
	return config, err
}

func ParseConfigFile(payload []byte) (Config, error) {
	config := Config{}
	err := yaml.Unmarshal(payload, &config)

	if config.Verifier == "" {
		return config, errors.New("error: field 'verifier' not found in config.yml")
	}
	if config.TestPath == "" {
		return config, errors.New("error: field 'testpath' not found in config.yml")
	}

	return config, err

}
