package api

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"min/model"
	"time"
)

func StartCron() {
	c := cron.New()

	MonthlyTask(c)
	WeeklyTask(c)
	DaylyTask(c)
	Mtest(c)
	c.Start()
}

func Mtest(c *cron.Cron) {
	_, err := c.AddFunc("* * * * *", func() {
		// 每十秒执行的任务
		fmt.Println("每一分钟执行的定时任务执行于:", time.Now())
	})
	if err != nil {
		log.Fatal("添加每一分钟执行的定时任务失败:", err)
	}
}

func MonthlyTask(c *cron.Cron) {

	_, err := c.AddFunc("0 0 1 * *", MonthlyFresh)
	if err != nil {
		fmt.Println("error:", err)
	}

}

func MonthlyFresh() {
	model.MonthReset()
}

func WeeklyTask(c *cron.Cron) {
	_, err := c.AddFunc("0 0 * * 0", WeeklyFresh)
	if err != nil {
		fmt.Println("error:", err)
	}
}
func WeeklyFresh() {
	model.WeekReset()
}

func DaylyTask(c *cron.Cron) {
	_, err := c.AddFunc("0 0 * * *", DaylyFresh)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func DaylyFresh() {
	model.DayReset()
}
