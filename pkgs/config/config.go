package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type StockInfo struct {
	Code       string  `yaml:"code"`
	Alias      string  `yaml:"alias"`
	HoldPrice  float64 `yaml:"holdPrice"`
	HoldNumber int     `yaml:"holdNumber"`
}

type Common struct {
	UpdateInterval time.Duration `yaml:"updateInterval"`
}

type MainConfig struct {
	Global Common `yaml:"Global"`
	Stocks []StockInfo
}

func InitConf(configPath string) MainConfig {
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	c := MainConfig{}
	err = viper.Unmarshal(&c)
	if err != nil {
		panic(fmt.Errorf("Error unmarshaling stocksUntil: %s \n", err))
	}

	// 剔除 HoldPrice 或 HoldNumber 小于 0 的股票
	var validStocks []StockInfo
	for _, stock := range c.Stocks {
		if stock.HoldPrice >= 0 && stock.HoldNumber >= 0 {
			validStocks = append(validStocks, stock)
		}
	}
	c.Stocks = validStocks

	return c
}
