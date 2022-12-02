/**
 *@filename       setting.go
 *@Description
 *@author          liyajun
 *@create          2022-12-02 19:31
 */

package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(ConfigUnit)

type ConfigUnit struct {
	*AppConfig    `mapstructure:"app"`
	*AlipayConfig `mapstructure:"alipay"`
}

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Port    string `mapstructure:"port"`
	Version string `mapstructure:"version"`
}

type AlipayConfig struct {
	Appid            string `mapstructure:"Appid"`
	NotifyURL        string `mapstructure:"notifyURL"`
	ReturnURL        string `mapstructure:"returnURL"`
	ProductCode      string `mapstructure:"productCode"`
	IsProduction     bool   `mapstructure:"isProduction"`
	AlipayPublicKey  string `mapstructure:"alipayPublicKey"`
	AlipayPrivateKey string `mapstructure:"alipayPrivateKey"`
}

func Init() (err error) {
	viper.SetConfigFile("./config/config.yaml")
	err = viper.ReadInConfig()
	if err != nil { // 处理读取配置文件的错误
		panic(fmt.Errorf("viper.ReadInConfig() failed: %s \n", err))
	}
	if err := viper.Unmarshal(Conf); err != nil { //解析配置失败
		panic(fmt.Errorf("viper.Unmarshal() failed: %s \n", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		//配置文件修改之后自动修改Conf
		fmt.Printf("config file changed %s \n:", e.Name)
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("viper.Unmarshal() failed: %s \n", err))
		}
	})
	return
}
