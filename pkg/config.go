package pkg

import (
	"encoding/json"
	"io/ioutil"
)

/*
* @Author:hanyajun
* @Date:2019/5/25 22:52
* @Name:pkg
* @Function: 配置
 */

type Config struct {
	AccessKey     string `json:"access_key"`
	SecretKey     string `json:"secret_key"`
	Zone          int    `json:"zone"`
	Bucket        string `json:"Bucket"`
	UseHTTPS      bool   `json:"use_https"`
	UseCdnDomains bool   `json:"use_cdn_domains"`
	Domain        string `json:"domain"`
}

func LoadConfig() *Config {
	var config Config
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(b)
	}

	err = json.Unmarshal(b, &config)
	if err != nil {
		panic(err)
	}
	return &config
}
