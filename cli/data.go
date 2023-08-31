package cli

import (
	"fmt"
	"reflect"
	"time"
)

type Message struct {
	SuccessInfo string `json:"sucessinfo"`

	SpareMoney string `json:"sparemoney"`
}

func Log(message Message) {
	green := "\033[32m"
	cyan := "\033[36m"
	pink := "\033[95m"
	reset := "\033[0m"

	fmt.Println(cyan, "[----------------------------------------------------------]", reset)

	//m := reflect.TypeOf(message)
	v := reflect.ValueOf(message)
	for i := 0; i < v.NumField(); i++ {
		//field := m.Field(i)
		//fmt.Println("字段名：", field.Name)
		value := v.Field(i)
		line := fmt.Sprintf("      %s", value)
		fmt.Println(green, line, reset)
	}
	//fmt.Println(pink, "\n  __  __   ___   _  _ \n |  \\/  | |_ _| | \\| |\n | |\\/| |  | |  | .` |\n |_|  |_| |___| |_|\\_|    @yowayimono\n                      \n", reset)
	fmt.Println(pink, "\n  __  __   ___   _  _ \n |  \\/  | |_ _| | \\| |\n | |\\/| |  | |  | .` |\n |_|  |_| |___| |_|\\_|    @yowayimono  ", time.Now().Format("2006-01-02 15:04:05"), reset)
	fmt.Println(cyan, "[----------------------------------------------------------]", reset)
}
