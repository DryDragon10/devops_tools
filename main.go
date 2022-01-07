package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

type Price struct {
	Name  string `yaml:"Name"`
	Price int64  `yaml:"Price"`
}

func main() {
	filename := "./prices.txt"
	bts, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bts), "\n")
	prices := []Price{}
	for _, line := range lines {
		parts := strings.Split(line, "	")
		if len(parts) != 2 {
			continue
		}
		price, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			fmt.Printf("fail to parse %+v\n", parts)
			continue
		}
		prices = append(prices, Price{
			Name:  parts[0],
			Price: price,
		})
	}

	out, err := yaml.Marshal(&prices)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(out))
}
