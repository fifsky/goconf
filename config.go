package goconf

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

//config file path and ext,Ext default .json
type Config struct {
	Path string
	Ext  string
}

//if key does not exist, return error
//key: dev.user.name  dev is filename in config path
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

//Ignore the error and return a zero value when it does not exist
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
