package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	ServerAddress      string `json:"serverAddress"`
	TimeoutSeconds     int    `json:"timeoutSeconds"`
	Certs              string `json:"certsTLS"`
	Mailbox_paths_list string `json:"mailbox_paths_list"`
	New_mail_path      string `json:"new_mail_path"`
	Junk_path          string `json:"junk_path"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cfg := &Config{}
	if err := json.NewDecoder(file).Decode(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
