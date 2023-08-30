/*package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var HistoryExpense []uint
var MonthExpense uint //当月花销
func main() {
	rootCmd := &cobra.Command{
		Use:   "min",
		Short: "一个简单的记账工具",
		Long:  "现在才开始！",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("一个简单的记账工具...")
		},
	}

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "添加一份花销",
		Long:  "添加花销命令",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			//count := args[0]
			count, _ := strconv.Atoi(args[0])
			MonthExpense += uint(count)
			fmt.Println("成功添加一笔花销 -> ", count)
			fmt.Println("当月总花销 -> ", MonthExpense)
		},
	}

	rootCmd.AddCommand(addCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)

	}

}
*/

package main

import (
	"min/api"
	"min/model"
	"min/service"
)

func main() {
	api.StartCron()
	model.Init()

	//serve := service.NewRouter()

	service.Run(":3000")
}
