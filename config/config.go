package config

// simply loading a config.json located in the root that should contain connection info

import (
	"encoding/json"
	"log"
	"os"

	. "issues-commenter/utilities"
)

var (
	Token        string
	Report_path	 string
	config		*config_struct
)

type config_struct struct {
	Token        string	`json:"token"`
	Report_path	 string `json:"report_path"`
}

func Read_config() error {
	log.Println("Reading config file...")
	file, err := os.ReadFile("./config.json")
	CE(err)
	
	err = json.Unmarshal(file, &config)
	CE(err)

	Token = config.Token
	Report_path = config.Report_path

	return nil
}
