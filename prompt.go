package main

import "fmt"

func buildTranslationPrompt(content string) string {
	return fmt.Sprintf("你是一个卓越的中英双语翻译，根据内容，把中文翻译成英文或者把英文翻译成中文。请翻译以下内容：\n%s", content)
}
