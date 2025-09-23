package config

import (
	"encoding/json"
	"os"

	"github.com/woshilapp/netprotector/server/global"
)

func ReadConfig() error {
	// read file
	data, err := os.ReadFile("./data/config.json")
	if err != nil {
		return err
	}

	// parse json
	err = json.Unmarshal(data, &global.Cfg)
	if err != nil {
		return err
	}

	return nil
}

func ReadRules() error {
	// read file
	data, err := os.ReadFile("./data/rules.json")
	if err != nil {
		return err
	}

	// parse json
	err = json.Unmarshal(data, &global.Rule)
	if err != nil {
		return err
	}

	return nil
}

func ReadUsers() error {
	// read file
	data, err := os.ReadFile("./data/users.json")
	if err != nil {
		return err
	}

	// parse json
	err = json.Unmarshal(data, &global.Users)
	if err != nil {
		return err
	}

	return nil
}

func WriteConfig() error {
	data, err := json.MarshalIndent(global.Cfg, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("./data/config.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func WriteRules() error {
	data, err := json.MarshalIndent(global.Rule, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("./data/rules.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func WriteUsers() error {
	data, err := json.MarshalIndent(global.Users, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("/data/users.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}
