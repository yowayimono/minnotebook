package cli

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func AddExpense(addr string, number string) {
	url := "http://" + addr + "/app/v1/add/" + number
	resp, err := http.Post(url, "", nil)
	if err != nil {
		log.Fatal("Error adding expense:", err)
	}
	defer resp.Body.Close()

	var message Message

	err = json.NewDecoder(resp.Body).Decode(&message)

	if err != nil {
		log.Fatal("Error decoding JSON response:", err)
	}

	if resp.StatusCode == http.StatusOK {
		//fmt.Println("添加成功.")
		Log(message)
	} else {
		fmt.Println("Failed to add expense. Status code:", resp.StatusCode)
	}

}

func DeleteExpense(addr string) {
	url := "http://" + addr + "/app/v1/delete"

	//url := "http://localhost:3031/app/v1/delete"
	resp, err := http.Post(url, "", nil)
	if err != nil {
		log.Fatal("Error deleting expense:", err)
	}
	defer resp.Body.Close()
	var message Message

	err = json.NewDecoder(resp.Body).Decode(&message)

	if err != nil {
		log.Fatal("Error decoding JSON response:", err)
	}

	if resp.StatusCode == http.StatusOK {
		//fmt.Println("添加成功.")
		Log(message)
	} else {
		fmt.Println("Failed to delete expense. Status code:", resp.StatusCode)
	}
}

func GetMonthExpenses(addr string) {
	url := "http://" + addr + "/app/v1/get/month"
	resp, err := http.Post(url, "", nil)
	if err != nil {
		log.Fatal("Error getting month expenses:", err)
	}
	defer resp.Body.Close()

	var message Message

	err = json.NewDecoder(resp.Body).Decode(&message)

	if err != nil {
		log.Fatal("Error decoding JSON response:", err)
	}

	if resp.StatusCode == http.StatusOK {
		//fmt.Println("添加成功.")
		Log(message)
	} else {
		fmt.Println("Failed to delete expense. Status code:", resp.StatusCode)
	}
}

func GetWeekExpenses(addr string) {
	url := "http://" + addr + "/app/v1/get/week"
	resp, err := http.Post(url, "", nil)
	if err != nil {
		log.Fatal("Error getting week expenses:", err)
	}
	defer resp.Body.Close()

	var message Message

	err = json.NewDecoder(resp.Body).Decode(&message)

	if err != nil {
		log.Fatal("Error decoding JSON response:", err)
	}

	if resp.StatusCode == http.StatusOK {
		//fmt.Println("添加成功.")
		Log(message)
	} else {
		fmt.Println("Failed to delete expense. Status code:", resp.StatusCode)
	}
}

func GetDayExpenses(addr string) {
	url := "http://" + addr + "/app/v1/get/day"
	resp, err := http.Post(url, "", nil)
	if err != nil {
		log.Fatal("Error getting day expenses:", err)
	}
	defer resp.Body.Close()

	var message Message

	err = json.NewDecoder(resp.Body).Decode(&message)

	if err != nil {
		log.Fatal("Error decoding JSON response:", err)
	}

	if resp.StatusCode == http.StatusOK {
		//fmt.Println("添加成功.")
		Log(message)
	} else {
		fmt.Println("Failed to delete expense. Status code:", resp.StatusCode)
	}
}

func GetMonthExpensesWithCount(addr string, count string) {
	url := "http://" + addr + "/app/v1/get/month/" + count
	resp, err := http.Post(url, "", nil)
	if err != nil {
		log.Fatal("Error getting month expenses:", err)
	}
	defer resp.Body.Close()
	var message Message

	err = json.NewDecoder(resp.Body).Decode(&message)

	if err != nil {
		log.Fatal("Error decoding JSON response:", err)
	}

	if resp.StatusCode == http.StatusOK {
		//fmt.Println("添加成功.")
		Log(message)
	} else {
		fmt.Println("Failed to delete expense. Status code:", resp.StatusCode)
	}
}

func GetWeekExpensesWithCount(addr string, count string) {
	url := "http://" + addr + "/app/v1/get/week/" + count
	resp, err := http.Post(url, "", nil)
	if err != nil {
		log.Fatal("Error getting week expenses:", err)
	}
	defer resp.Body.Close()

	var message Message

	err = json.NewDecoder(resp.Body).Decode(&message)

	if err != nil {
		log.Fatal("Error decoding JSON response:", err)
	}

	if resp.StatusCode == http.StatusOK {
		//fmt.Println("添加成功.")
		Log(message)
	} else {
		fmt.Println("Failed to delete expense. Status code:", resp.StatusCode)
	}
}

func GetDayExpensesWithCount(addr string, count string) {
	url := "http://" + addr + "/app/v1/get/day/" + count
	resp, err := http.Post(url, "", nil)
	if err != nil {
		log.Fatal("Error getting day expenses:", err)
	}
	defer resp.Body.Close()
	var message Message

	err = json.NewDecoder(resp.Body).Decode(&message)

	if err != nil {
		log.Fatal("Error decoding JSON response:", err)
	}

	if resp.StatusCode == http.StatusOK {
		//fmt.Println("添加成功.")
		Log(message)
	} else {
		fmt.Println("Failed to delete expense. Status code:", resp.StatusCode)
	}
}

func GetExpenseTimes(addr string) {
	url := "http://" + addr + "/app/v1/get/times"
	resp, err := http.Post(url, "", nil)
	if err != nil {
		log.Fatal("Error getting expense times:", err)
	}
	defer resp.Body.Close()

	var message Message

	err = json.NewDecoder(resp.Body).Decode(&message)

	if err != nil {
		log.Fatal("Error decoding JSON response:", err)
	}

	if resp.StatusCode == http.StatusOK {
		//fmt.Println("添加成功.")
		Log(message)
	} else {
		fmt.Println("Failed to delete expense. Status code:", resp.StatusCode)
	}
}
