package model

import (
	"fmt"
	"strconv"
	"time"
)

/*
var CurrentMonthExpense int //当月花销
var HistoryExpense []int    //历史花销

var CurrentWeekExpense int //当周花销
var HistoryWeek []int      //历史每周花销

var CurrentDayExpense int //当日花销
var HistoryDay []int      //历史每周花销

var Prenumber int //上次花销
*/

type Expense struct {
	CurrentMonthExpense int      //当月花销
	HistoryMonth        []string //历史花销

	CurrentWeekExpense int      //当周花销
	HistoryWeek        []string //历史每周花销

	CurrentDayExpense int      //当日花销
	HistoryDay        []string //历史每周花销

	Prenumber int //上次花销

	Retimes uint
}

func NewExpense() *Expense {
	return &Expense{
		CurrentMonthExpense: 0,                     //当月花销
		HistoryMonth:        make([]string, 10000), //历史花销

		CurrentWeekExpense: 0,                     //当周花销
		HistoryWeek:        make([]string, 10000), //历史每周花销

		CurrentDayExpense: 0,                     //当日花销
		HistoryDay:        make([]string, 10000), //历史每周花销

		Prenumber: 0, //上次花销
		Retimes:   0,
	}
}

var CurrentExpense Expense

func Init() {
	CurrentExpense := NewExpense()
	Persist()
	fmt.Println("初始化信息结构...")
	fmt.Println(CurrentExpense)
}

func MonthReset() {
	// 添加时间戳
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	//拼接字符串
	currentmonth := fmt.Sprintf("%s %s", strconv.Itoa(CurrentExpense.CurrentMonthExpense), timestamp)

	//添加上个月数据
	CurrentExpense.HistoryMonth = append(CurrentExpense.HistoryMonth, currentmonth)

	CurrentExpense.CurrentMonthExpense = 0
}

func WeekReset() {
	// 添加时间戳
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	//拼接字符串
	currentweek := fmt.Sprintf("%s元 %s", strconv.Itoa(CurrentExpense.CurrentWeekExpense), timestamp)

	//添加上个月数据
	CurrentExpense.HistoryMonth = append(CurrentExpense.HistoryMonth, currentweek)

	CurrentExpense.CurrentMonthExpense = 0

}

func DayReset() {
	// 添加时间戳
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	//拼接字符串
	currentday := fmt.Sprintf("%s %s", strconv.Itoa(CurrentExpense.CurrentDayExpense), timestamp)

	//添加上个月数据
	CurrentExpense.HistoryMonth = append(CurrentExpense.HistoryMonth, currentday)

	CurrentExpense.CurrentMonthExpense = 0
}
