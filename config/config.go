package config

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// loadConfig - функция заполняющая структуру конфига (structConfg)
// из файла (Path) или возвращающая ошибку если это сделать не удалось
func LoadConfig(Path string, structConfg interface{}) error {
	configBytes, err := os.ReadFile(Path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(configBytes, structConfg)
	if err != nil {
		return err
	}
	return nil
}

// getConfigPathFromArgs - Функция парсит аргументы заданные в командной строке
// и получает путь к файлу конфигурации иначе завершает программу с кодом 1
func GetConfigPathFromArgs() string {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <path_to_config_file>\n", os.Args[0])
		flag.PrintDefaults()
	}
	var configPath string
	flag.Parse()
	if configPath = flag.Arg(0); configPath == "" {
		flag.Usage()
		os.Exit(1)
	}
	return configPath
}
