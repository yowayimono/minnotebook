package model

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Persist() {
	LastLog, err := GetLastLine()
	if err != nil {
		return
	}
	fmt.Println(err)
	err = json.Unmarshal([]byte(LastLog), &CurrentExpense)
	if err != nil {
		//log.Fatal("Error parsing log file:", err)
	}

}

func GetLastLine() (string, error) {
	// 打开文件
	file, err := os.Open("./app.log")
	if err != nil {
		log.Fatal("Error opening file:", err)
		return "", err
	}
	defer file.Close()

	// 判断文件大小是否为0

	// 创建一个 Scanner 对象来逐行读取文件
	scanner := bufio.NewScanner(file)

	var lastLine string

	// 逐行读取文件
	for scanner.Scan() {
		lastLine = scanner.Text()
	}

	// 检查是否有扫描错误
	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning file:", err)
	}
	return lastLine, nil
}
