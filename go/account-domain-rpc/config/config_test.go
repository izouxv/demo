package config

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"testing"
)

type graphite struct {
	Host string
	Port int
	Pool string
}

func Test_Config(t *testing.T) {
	c, _ := NewConfigFromFile("./config.yaml")

	//getStruct
	var result []graphite
	v, err := c.DIY("metrics")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
	err = mapstructure.Decode(v, &result)
	fmt.Println(result)

	var te []graphite
	c.GetStruct("metrics", &te)
	fmt.Println(te)
	//getSection
	/*
		v, err := c.GetSection("metrics")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(v["host"])
	*/
}
