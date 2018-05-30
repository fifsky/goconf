package goconf

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

type Config struct {
	Path string
	Ext  string
}

func (c *Config) Get(key string) (*gjson.Result, error) {
	keys := strings.Split(key, ".")

	if len(keys) < 2 {
		return nil, errors.New("config XPath is at least two paragraphs")
	}

	configFile := c.Path + keys[0] + c.Ext

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return nil, errors.New("config path not found:" + configFile)
	}

	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, errors.Wrap(err, "config file read err")
	}

	result := gjson.ParseBytes(b)
	ret := result.Get(strings.Join(keys[1:], "."))
	return &ret, nil
}

func (c *Config) MustGet(key string) *gjson.Result {
	ret, err := c.Get(key)
	if err != nil {
		return &gjson.Result{}
	}

	return ret
}

func NewConfig(path string) (*Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, errors.New("config path not found:" + path)
	}
	configPath, _ := filepath.Abs(path)
	config := &Config{
		Path: configPath + "/",
		Ext:  ".json",
	}
	return config, nil
}