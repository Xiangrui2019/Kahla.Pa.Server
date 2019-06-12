package services

import (
	"Kahla.PublicAddress.Server/models"
	"encoding/json"
	"io/ioutil"
)

func SaveConfig(config *models.Config) ([]byte, error) {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func SaveConfigToFile(filename string, config *models.Config) error {
	data, err := SaveConfig(config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0600)
	if err != nil {
		return err
	}
	return nil
}

func LoadConfig(data []byte) (*models.Config, error) {
	config := new(models.Config)
	err := json.Unmarshal(data, config)
	if err != nil {
		return config, err
	}
	return config, nil
}

func LoadConfigFromFile(filename string) (*models.Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	config, err := LoadConfig(data)
	if err != nil {
		return config, err
	}
	return config, nil
}
