package conf

import (
	"encoding/json"
	"os"
)

type Conf struct {
	DstUrl string 
	FreshInterval int
}

func ReadConfig(confpath string) (*Conf, error) {
	file, _ := os.Open(confpath)
	decoder := json.NewDecoder(file)
	config := Conf{}
	err := decoder.Decode(&config)

	return &config, err
}
