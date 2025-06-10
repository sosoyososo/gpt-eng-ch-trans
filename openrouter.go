package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func openrouter(content string, config OpenrouterConfig) {

	prompt := buildTranslationPrompt(content)

	requestBody := map[string]interface{}{
		"model": config.Model,
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", config.APIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", "Bearer "+config.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(responseBody, &result); err != nil {
		panic(err)
	}

	if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
		if choice, ok := choices[0].(map[string]interface{}); ok {
			if message, ok := choice["message"].(map[string]interface{}); ok {
				if content, ok := message["content"].(string); ok {
					fmt.Println(content)					
					return
				}
			}
		}
	}

	fmt.Println("无法解析API响应:")
	fmt.Println(string(responseBody))
}
