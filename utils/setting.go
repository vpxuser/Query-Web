package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpHost string
	HttpPort string
	JwtKey 	 string
	UseTls 	 bool
	CrtPath  string
	KeyPath	 string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("Config File Reading Error :", err)
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpHost = file.Section("server").Key("HttpHost").MustString("nu1l.online")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8081")
	JwtKey = file.Section("server").Key("JwtKey").MustString("Ekxy3Dxk5i")
	UseTls = file.Section("server").Key("UseTls").MustBool(true)
	CrtPath = file.Section("server").Key("CrtPath").MustString("/etc/v2ray-agent/tls/nu1l.online.crt")
	KeyPath = file.Section("server").Key("KeyPath").MustString("/etc/v2ray-agent/tls/nu1l.online.key")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("query")
	DbPassWord = file.Section("database").Key("DbPassWord").String()
	DbName = file.Section("database").Key("DbName").MustString("query")
}