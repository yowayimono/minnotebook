package main

import (
	"fmt"
	"min-cli/cli"

	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command.")
		return
	}

	addr := "localhost:3000"

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a number.")
			return
		}
		number := os.Args[2]
		cli.AddExpense(addr, number)
	case "delete":
		cli.DeleteExpense(addr)
	case "get/month":
		cli.GetMonthExpenses(addr)
	case "get/week":
		cli.GetWeekExpenses(addr)
	case "get/day":
		cli.GetDayExpenses(addr)
	case "get/month/":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a count.")
			return
		}
		count := os.Args[2]
		cli.GetMonthExpensesWithCount(addr, count)
	case "get/week/":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a count.")
			return
		}
		count := os.Args[2]
		cli.GetWeekExpensesWithCount(addr, count)
	case "get/day/":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a count.")
			return
		}
		count := os.Args[2]
		cli.GetDayExpensesWithCount(addr, count)
	case "get/times":
		cli.GetExpenseTimes(addr)
	default:
		fmt.Println("Unknown command.")
	}
}
