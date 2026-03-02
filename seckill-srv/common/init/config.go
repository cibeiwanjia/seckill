package init

import (
	"fmt"
	"seckill/seckill-srv/common/config"

	"github.com/spf13/viper"
)

func ConfigInit() {
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("./common")

	err := viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config.AppCfg)
	if err != nil {
		return
	}
	fmt.Println("配置文件加载成功", config.AppCfg)
}
