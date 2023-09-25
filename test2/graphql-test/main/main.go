package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

const repeat = 1000

func main() {

	requestTimes := []int64{}

	for i := 0; i < repeat; i++ {
		log.Println(i)
		getDataReq := getDataRequest()
		start := time.Now().UnixNano()

		getData(getDataReq)

		diff := time.Now().UnixNano() - start
		requestTimes = append(requestTimes, diff)
	}

	file, err := os.Create("cloud_data.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Values"}
	if err := writer.Write(header); err != nil {
		fmt.Println("Error writing header to CSV:", err)
		return
	}

	for _, value := range requestTimes {
		record := []string{fmt.Sprintf("%d", value)}
		if err := writer.Write(record); err != nil {
			fmt.Println("Error writing record to CSV:", err)
			return
		}
	}
}
