package configs

import (
	"fmt"
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug *bool   `yaml:"is_debug" env-required:"true"`
	Listen  Listen  `yaml:"listen"`
	Storage Storage `yaml:"storage"`
	PWD     string
}

type Storage struct {
	Psql Psql `yaml:"psql"`
}

type Listen struct {
	Type   string `yaml:"type" env-required:"true"`
	BindIP string `yaml:"bind_ip" env-required:"true"`
	Port   string `yaml:"port" env-required:"true"`
}

type Psql struct {
	Host          string `yaml:"host" env-required:"true"`
	Port          string `yaml:"port" env-required:"true"`
	Database      string `yaml:"database" env-required:"true"`
	Username      string `yaml:"username" env-required:"true"`
	Password      string `yaml:"password" env-required:"true"`
	PgPoolMaxConn int    `yaml:"pg_pool_max_conn" env-required:"true"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		pathConfig := pwd + "/config_cart.yml"

		fmt.Println("read application configuration: pwd: ", pwd)

		instance = &Config{PWD: pwd}
		if err = cleanenv.ReadConfig(pathConfig, instance); err != nil {
			_, err = cleanenv.GetDescription(instance, nil)
			if err != nil {
				fmt.Printf("read application configuration: err: %v \n", err)
			}
		}
	})

	return instance
}
