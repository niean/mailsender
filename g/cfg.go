package g

import (
	"encoding/json"
	"github.com/toolkits/file"
	"log"
	"sync"
)

type HttpConfig struct {
	Enable bool   `json:"enable"`
	Listen string `json:"listen"`
}

type RpcConfig struct {
	Enable bool   `json:"enable"`
	Listen string `json:"listen"`
}

type MailConfig struct {
	Enable            bool   `json:"enable"`
	SendConcurrent    int    `json:"sendConcurrent"`
	MaxQueueSize      int    `json:"maxQueueSize"`
	FromUser          string `json:"fromUser"`
	MailServerHost    string `json:"mailServerHost"`
	MailServerPort    int    `json:"mailServerPort"`
	MailServerAccount string `json:"mailServerAccount"`
	MailServerPasswd  string `json:"mailServerPasswd"`
}

type GlobalConfig struct {
	Debug bool        `json:"debug"`
	Http  *HttpConfig `json:"http"`
	Rpc   *RpcConfig  `json:"rpc"`
	Mail  *MailConfig `json:"mail"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	configLock = new(sync.RWMutex)
)

func GetConfig() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func LoadConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent. maybe you need `mv cfg.example.json cfg.json`")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}

	configLock.Lock()
	defer configLock.Unlock()
	config = &c

	log.Println("g.ParseConfig ok, file ", cfg)
}
