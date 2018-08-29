package config

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml"
	"io/ioutil"
	"log"
	"path/filepath"
	"github.com/cihub/seelog"
)

type Config struct {
	conf *map[string]interface{}
}

//解析文件，取出所有参数
func NewConfigFromFile(path string) *Config {
	seelog.Info("NewConfigFromFile...")
	filename, err := filepath.Abs(path)
	if err != nil {
		log.Fatal("can not find the file:", filename)
		panic(err)
	}
	data, err := ioutil.ReadFile(filename)
	//将解析出的参数转为map的形式
	m := make(map[string]interface{})
	if err != nil {
		log.Fatal("config init failed!,", err)
		panic(err)
	}
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatal("config unmarshal failed!", err)
		panic(err)
	}
	return &Config{&m}
}

// It accepts 1, 1.0, t, T, TRUE, true, True, YES, yes, Yes,Y, y, ON, on, On,
// 0, 0.0, f, F, FALSE, false, False, NO, no, No, N,n, OFF, off, Off.
// Any other value returns an error.
func ParseBool(val interface{}) (value bool, err error) {
	if val != nil {
		switch v := val.(type) {
		case bool:
			return v, nil
		case string:
			switch v {
			case "1", "t", "T", "true", "TRUE", "True", "YES", "yes", "Yes", "Y", "y", "ON", "on", "On":
				return true, nil
			case "0", "f", "F", "false", "FALSE", "False", "NO", "no", "No", "N", "n", "OFF", "off", "Off":
				return false, nil
			}
		case int8, int32, int64:
			strV := fmt.Sprintf("%s", v)
			if strV == "1" {
				return true, nil
			} else if strV == "0" {
				return false, nil
			}
		case float64:
			if v == 1 {
				return true, nil
			} else if v == 0 {
				return false, nil
			}
		}
		return false, fmt.Errorf("parsing %q: invalid syntax", val)
	}
	return false, fmt.Errorf("parsing <nil>: invalid syntax")
}

// Bool returns the boolean value for a given key.
func (c *Config) Bool(key string) (bool, error) {
	v, err := c.getData(key)
	if err != nil {
		return false, err
	}
	return ParseBool(v)
}

// Int returns the integer value for a given key.
func (c *Config) Int(key string) (int, error) {
	if v, err := c.getData(key); err != nil {
		return 0, err
	} else if vv, ok := v.(int); ok {
		return vv, nil
	} else if vv, ok := v.(int64); ok {
		return int(vv), nil
	}
	return 0, errors.New("not int value")
}

// DefaultInt returns the integer value for a given key.
// if err != nil return defaltval
func (c *Config) DefaultInt(key string, defaultval int) int {
	v, err := c.Int(key)
	if err != nil {
		return defaultval
	}
	return v
}

// Int64 returns the int64 value for a given key.
func (c *Config) Int64(key string) (int64, error) {
	if v, err := c.getData(key); err != nil {
		return 0, err
	} else if vv, ok := v.(int64); ok {
		return vv, nil
	}
	return 0, errors.New("not bool value")
}

// String returns the string value for a given key.
func (c *Config) String(key string) string {
	if value, err := c.getData(key); err != nil {
		return ""
	} else if vv, ok := value.(string); ok {
		return vv
	}
	return ""
}

// GetSection returns map for the given key
func (c *Config) GetSection(key string) (map[interface{}]interface{}, error) {
	if v, err := c.getData(key); err != nil {
		return nil, errors.New("not exist setction")
	} else if vv, ok := v.(map[interface{}]interface{}); ok {
		return vv, nil
	}
	/* check the type of v
	v, _ := c.getData(key)
	switch v.(type) {
	default:
		fmt.Printf("%T",v)
	}
	*/
	return nil, errors.New("section not map[interface{}]interface{}")
}

// getStruct returns struct for the give key and given struct type
func (c *Config) GetStruct(key string, out interface{}) error {
	v, err := c.getData(key)
	if err != nil {
		return err
	}
	err = mapstructure.Decode(v, out)
	if err != nil {
		return err
	}
	return nil
}

// DIY returns the raw value by a given key.
func (c *Config) DIY(key string) (v interface{}, err error) {
	return c.getData(key)
}

func (c *Config) getData(key string) (interface{}, error) {

	if len(key) == 0 {
		return nil, errors.New("key is empty")
	}

	if v, ok := (*c.conf)[key]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("not exist key %q", key)
}
