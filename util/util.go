package util

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

//int64 to string
func TranToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

//string to int64
func TranToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

//struct to string
func TranToStringFromStruct(x interface{}) string {
	b, e := json.Marshal(x)
	if e != nil {
		fmt.Println("err:", e)
	}
	return string(b)
}

func GetResponse() Response {
	return Response{
		Status: http.StatusOK,
		Code:   0,
	}
}

func ReadConfig(configPath, configName, configType string, key ...string) []string {
	config := viper.New()
	config.AddConfigPath(configPath) //设置读取的文件路径
	config.SetConfigName(configName) //设置读取的文件名
	config.SetConfigType(configType) //设置文件的类型

	//尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	//文件读取出来的内容:
	res := make([]string, 0, len(key))
	if len(key) > 0 {
		for _, v := range key {
			res = append(res, config.Get(v).(string))
		}
	}
	return res

}
