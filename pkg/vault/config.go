package vault

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	ConfigEnvVar     = "DSV_SDK_CONFIG"
	ConfigFileEnvVar = "DSV_SDK_CONFIG_FILE"
)

func UnmarshalConfig(configJson []byte) (*Configuration, error) {
	config := new(Configuration)

	if err := json.Unmarshal(configJson, config); err != nil {
		return nil, err
	} else {
		return config, nil
	}
}

func ParseConfig(filePath string) (*Configuration, error) {
	if content, err := ioutil.ReadFile(filePath); err != nil {
		return nil, err
	} else {
		return UnmarshalConfig(content)
	}
}

func GetConfigFromEnv() (*Configuration, error) {
	if v := os.Getenv(ConfigEnvVar); v != "" {
		return UnmarshalConfig([]byte(v))
	} else if v := os.Getenv(ConfigFileEnvVar); v != "" {
		return ParseConfig(v)
	} else {
		return nil, nil
	}
}
