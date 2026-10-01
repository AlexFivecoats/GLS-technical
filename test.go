package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// send http request with the login info to get the employee data
	url := "https://test.gls.com/salary-base.json"

	username := "user1"
	password := "password1"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Request creation error:", err)
		return
	}

	req.SetBasicAuth(username, password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed:", resp.Status)
		return
	}

	// Parse the response. Put the employee data into an array for further use.
	/*
		{
			"data": [
				{
				"name": "Employee One",
				"office": "City One",
				"title": "Title One",
				"salary": "Salary One"
				},
				{
				"name": "Employee Two",
				"office": "City Two",
				"title": "Title Two",
				"salary": "Salary Two"
				},
				.
				.
				.
			]
		}
	*/
	var result map[string][]map[string]string

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("JSON error:", err)
		return
	}

	employees := result["data"]

	fmt.Println("Number of employees:", len(employees))

	// Calculate the min, median, max of employee salaries.
	var salaries []float64 = make([]float64, len(employees))

	replacer := strings.NewReplacer("$", "", ",", "")
	for i, employee := range employees {
		temp := replacer.Replace(employee["salary"])
		salfloat, _ := strconv.ParseFloat(temp, 64)
		salaries[i] = salfloat
	}

	sort.Float64s(salaries)

	min := salaries[0]
	max := salaries[len(salaries)-1]
	middle := len(salaries) / 2
	var median float64

	if middle%2 == 1 {
		median = salaries[middle]
	} else {
		median = (salaries[middle] + salaries[middle-1]) / 2
	}

	fmt.Printf("Employee Salaries\n  Min: $%.2f\n  Max: $%.2f\n  Median: $%.2f", min, max, median)
}
