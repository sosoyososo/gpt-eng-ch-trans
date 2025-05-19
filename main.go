package main

import (
	"embed"
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Deepseek   DeepseekConfig   `json:"deepseek"`
	Openrouter OpenrouterConfig `json:"openrouter"`
}

//go:embed config.json
var configFiles embed.FS

func loadConfig() (*Config, error) {	
	configData, err := configFiles.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(configData, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("用法: go run main.go [deepseek|openrouter] [要翻译的内容]")
		return
	}

	config, err := loadConfig()
	if err != nil {
		fmt.Printf("加载配置文件失败: %v\n", err)
		return
	}

	model := os.Args[1]
	content := os.Args[2]

	switch model {
	case "deepseek":
		deepseek(content, config.Deepseek)
	case "openrouter":
		openrouter(content, config.Openrouter)
	default:
		fmt.Printf("未知模型: %s\n", model)
	}
}
