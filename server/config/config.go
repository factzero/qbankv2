package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 数据数据库配置
type DBconf struct {
	Driver   string `yaml:"driver"`
	Hostname string `yaml:"hostname"`
	Hostport string `yaml:"hostport"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Prefix   string `yaml:"prefix"`
}

// 应用配置
type App struct {
	Port              string `yaml:"port"`
	Version           string `yaml:"version"`
	Env               string `yaml:"env"`
	Apisecret         string `yaml:"apisecret"`
	Allowurl          string `yaml:"allowurl"`
	TokenOutTime      string `yaml:"tokenouttime"`
	CPUnum            string `yaml:"cpunum"`
	Domain            string `yaml:"domain"`
	Vueobjroot        string `yaml:"vueobjroot"`
	Rootview          string `yaml:"rootview"`
	RunlogType        string `yaml:"runlogtype"`
	NoVerifyTokenRoot string `yaml:"noVerifyTokenRoot"`
	NoVerifyAPIRoot   string `yaml:"noVerifyAPIRoot"`
	NoVerifyToken     string `yaml:"noVerifyToken"`
	NoVerifyAPI       string `yaml:"noVerifyAPI"`
}

// JWT验证
type Jwt struct {
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"`
	JwtTtl int64  `mapstructure:"jwt_ttl" json:"jwt_ttl" yaml:"jwt_ttl"` // token 有效期（秒）
}

// 日志文件
type Log struct {
	Level      string `mapstructure:"level" json:"level" yaml:"level"`
	RootDir    string `mapstructure:"root_dir" json:"root_dir" yaml:"root_dir"`
	Filename   string `mapstructure:"filename" json:"filename" yaml:"filename"`
	Format     string `mapstructure:"format" json:"format" yaml:"format"`
	ShowLine   bool   `mapstructure:"show_line" json:"show_line" yaml:"show_line"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"` // MB
	MaxAge     int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`    // day
	Compress   bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}

type Config struct {
	DBconf DBconf `yaml:"dbconf"`
	App    App    `yaml:"app"`
	Jwt    Jwt    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Log    Log    `mapstructure:"log" json:"log" yaml:"log"`
}

// 读取Yaml配置文件，并转换成Config对象  struct结构
func (config *Config) InitializeConfig() *Config {
	// 获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	vip := viper.New()
	vip.AddConfigPath(path + "/config") // 设置读取的文件路径
	vip.SetConfigName("settings")       // 设置读取的文件名
	vip.SetConfigType("yaml")           // 设置文件的类型
	// 尝试进行配置读取
	if err := vip.ReadInConfig(); err != nil {
		panic(err)
	}
	// 监听配置文件
	vip.WatchConfig()
	vip.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := vip.Unmarshal(&config); err != nil {
			fmt.Println(err)
		}
	})

	err = vip.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return config
}
