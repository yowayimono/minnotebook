package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"min/model"
	"os"
	"strconv"
)

func AddExpense(c *gin.Context) {
	fmt.Println("添加花销逻辑...")
	numberStr := c.Param("number")

	//fmt.Println(numberStr)
	number, err := strconv.Atoi(numberStr)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid number"})
		return
	}

	//更新数据
	model.CurrentExpense.CurrentDayExpense += number
	model.CurrentExpense.CurrentWeekExpense += number
	model.CurrentExpense.CurrentMonthExpense += number

	model.CurrentExpense.Prenumber = number

	// 将结构体转换为 JSON 字符串
	jsonBytes, err := json.Marshal(model.CurrentExpense)
	if err != nil {
		fmt.Println("JSON marshaling failed:", err)
		return
	}

	// 将 JSON 字节切片转换为字符串
	jsonString := string(jsonBytes)
	message := jsonString
	//jsonString += "\n"
	//message := fmt.Sprintf("ADD %v %v %v %v %v", model.CurrentExpense.CurrentMonthExpense, model.CurrentExpense.CurrentWeekExpense, model.CurrentExpense.CurrentDayExpense, model.CurrentExpense.Prenumber, number)
	fmt.Println(message)

	LogCmd(message)
	model.CurrentExpense.Retimes++
	// 在这里使用 number 变量执行相关操作，例如添加费用

	messaged := model.Message{
		SuccessInfo: "添加成功： " + strconv.Itoa(int(number)),
		SpareMoney:  "少花点吧！这个月已经花了： " + strconv.Itoa(model.CurrentExpense.CurrentMonthExpense),
	}

	c.JSON(200, messaged)
	//c.JSON(200, gin.H{"当月已花": model.CurrentExpense.CurrentMonthExpense,
	//	"message": fmt.Sprintf("已添加: %d", number)})

}

func Delete(c *gin.Context) {

	model.CurrentExpense.CurrentDayExpense -= model.CurrentExpense.Prenumber
	model.CurrentExpense.CurrentWeekExpense -= model.CurrentExpense.Prenumber
	model.CurrentExpense.CurrentMonthExpense -= model.CurrentExpense.Prenumber
	// 将结构体转换为 JSON 字符串
	jsonBytes, err := json.Marshal(model.CurrentExpense)
	if err != nil {
		fmt.Println("JSON marshaling failed:", err)
		return
	}

	messaged := model.Message{
		SuccessInfo: "删除成功： " + strconv.Itoa(int(model.CurrentExpense.Prenumber)),
		SpareMoney:  "少花点吧！这个月已经花了： " + strconv.Itoa(model.CurrentExpense.CurrentMonthExpense),
	}

	c.JSON(200, messaged)
	model.CurrentExpense.Retimes++
	model.CurrentExpense.Prenumber = 0
	// 将 JSON 字节切片转换为字符串
	jsonString := string(jsonBytes)
	message := jsonString
	//jsonString += "\n"
	fmt.Println(jsonString)
	//message := fmt.Sprintf("DELETE %v %v %v %v %v", model.CurrentExpense.CurrentMonthExpense, model.CurrentExpense.CurrentWeekExpense, model.CurrentExpense.CurrentDayExpense, model.CurrentExpense.Prenumber, 0)
	LogCmd(message)

	fmt.Println("删除上一条添加记录...")
}

func LogCmd(message string) {
	// 打开文件以供写入，如果文件不存在则创建，如果文件存在则追加写入
	logFile, err := os.OpenFile("app.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Failed to open log file: %v", err)
		return
	}
	defer logFile.Close()

	// 日志内容

	// 写入日志文件
	data := fmt.Sprintf("%s\n", message)
	_, err = logFile.WriteString(data)
	if err != nil {
		fmt.Printf("Failed to write to log file: %v", err)
		return
	}

	fmt.Println("Log file updated.")
}

func GetMonth(c *gin.Context) {
	message := model.Message{
		SuccessInfo: "这个月已经花了这么多： " + strconv.Itoa(int(model.CurrentExpense.CurrentMonthExpense)),
		SpareMoney:  "少花点吧！",
	}

	c.JSON(200, message)

	//c.JSON(200, gin.H{"你这个月已经花了这么多钱": model.CurrentExpense.CurrentMonthExpense})
}

func GetWeek(c *gin.Context) {
	message := model.Message{
		SuccessInfo: "这周已经花了这么多: " + strconv.Itoa(int(model.CurrentExpense.CurrentWeekExpense)),
		SpareMoney:  "少花点吧！",
	}

	c.JSON(200, message)

	//c.JSON(200, gin.H{"你这周花了这么多钱": model.CurrentExpense.CurrentWeekExpense})
}

func GetDay(c *gin.Context) {
	message := model.Message{
		SuccessInfo: "今天花了这么多： " + strconv.Itoa(int(model.CurrentExpense.CurrentDayExpense)),
		SpareMoney:  "少花点吧！",
	}

	c.JSON(200, message)

	//c.JSON(200, gin.H{"你今天花了这么多钱": model.CurrentExpense.CurrentDayExpense})
}

func GetTimes(c *gin.Context) {
	message := model.Message{
		SuccessInfo: "这么多次啦！: " + strconv.Itoa(int(model.CurrentExpense.Retimes)),
		SpareMoney:  "少花点吧！",
	}

	c.JSON(200, message)
	//c.JSON(200, gin.H{"这么多次": model.CurrentExpense.Retimes})
}

func MulGetMonth(c *gin.Context) {
	count := c.Param("count")

	//fmt.Println(numberStr)
	cnt, err := strconv.Atoi(count)

	//test
	//model.CurrentExpense.HistoryMonth = append(model.CurrentExpense.HistoryMonth, "1323", "1213", "123124", "12312", "1231")
	size := len(model.CurrentExpense.HistoryMonth)

	if size < cnt {
		message := model.Message{
			SuccessInfo: "没这么久！",
			SpareMoney:  "少花点吧！",
		}

		c.JSON(200, message)
	}

	if err != nil {
		fmt.Println("error:", err)
	}

	MonthExpense := model.CurrentExpense.HistoryMonth[size-cnt:]

	c.JSON(200, MonthExpense)
}

func MulGetWeek(c *gin.Context) {
	count := c.Param("count")

	//fmt.Println(numberStr)
	cnt, err := strconv.Atoi(count)

	//test
	//model.CurrentExpense.HistoryWeek = append(model.CurrentExpense.HistoryWeek, "1323", "1213", "123124", "12312", "1231")
	size := len(model.CurrentExpense.HistoryWeek)

	if size < cnt {
		message := model.Message{
			SuccessInfo: "没这么久！",
			SpareMoney:  "少花点吧！",
		}

		c.JSON(200, message)
	}

	if err != nil {
		fmt.Println("error:", err)
	}

	WeekExpense := model.CurrentExpense.HistoryWeek[size-cnt:]

	c.JSON(200, WeekExpense)
}

func MulGetDay(c *gin.Context) {
	count := c.Param("count")

	//fmt.Println(numberStr)
	cnt, err := strconv.Atoi(count)

	//test
	//model.CurrentExpense.HistoryDay = append(model.CurrentExpense.HistoryDay, "1323", "1213", "123124", "12312", "1231")
	size := len(model.CurrentExpense.HistoryDay)

	if size < cnt {
		message := model.Message{
			SuccessInfo: "没这么久！",
			SpareMoney:  "少花点吧！",
		}

		c.JSON(200, message)
	}

	if err != nil {
		fmt.Println("error:", err)
	}

	DayExpense := model.CurrentExpense.HistoryDay[size-cnt:]

	c.JSON(200, DayExpense)
}

/*
func SetExpense() gin.HandlerFunc {

	return func(c *gin.Context) {
		model.Once.Do(func() {
			expense := model.Expense{
				CurrentMonthExpense: 0,                  //当月花销
				HistoryExpense:      make([]int, 10000), //历史花销

				CurrentWeekExpense: 0,                  //当周花销
				HistoryWeek:        make([]int, 10000), //历史每周花销

				CurrentDayExpense: 0,                  //当日花销
				HistoryDay:        make([]int, 10000), //历史每周花销

				Prenumber: 0, //上次花销
			}

			c.Set("Expense", expense)
			fmt.Println("已设置信息上下文！")
			c.Next()
		})

	}
}


func GetExpense(c *gin.Context) model.Expense {
	val, exi := c.Get("Expense")

	if !exi {
		c.JSON(500, gin.H{"error": "获取上下文错误!"})
	}

	Data := val.(model.Expense)

	return Data
}
*/
