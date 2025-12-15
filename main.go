package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func bubbleSortIterative(arr []float64) {
	n := len(arr)
	var i, j int

	for i = 0; i < n-1; i++ {
		for j = 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func bubbleSortRecursive(arr []float64, n int) {
	var i int

	if n <= 1 {
		return
	}
	for i = 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
	bubbleSortRecursive(arr, n-1)
}

func readTemperatureCSV(filename string) ([]float64, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var temps []float64
	var i int

	for i = 1; i < len(records); i++ {
		tempStr := records[i][7] // temperature_celsius

		temp, err := strconv.ParseFloat(tempStr, 64)
		if err == nil {
			temps = append(temps, temp)
		}
	}

	return temps, nil
}


func main() {
	data, err := readTemperatureCSV("weather.csv")
	
	if err != nil {
		fmt.Println("Gagal baca CSV:", err)
		return
	}

	iterData := make([]float64, len(data))
	recData := make([]float64, len(data))
	copy(iterData, data)
	copy(recData, data)

	startIter := time.Now()
	bubbleSortIterative(iterData)
	iterTime := time.Since(startIter)

	startRec := time.Now()
	bubbleSortRecursive(recData, len(recData))
	recTime := time.Since(startRec)

	fmt.Println("BUBBLE SORT SUHU CELCIUS")
	fmt.Println("Jumlah data :", len(data))
	fmt.Printf("Suhu terendah : %.2f °C\n", iterData[0])
	fmt.Printf("Suhu tertinggi: %.2f °C\n", iterData[len(iterData)-1])

	fmt.Println("\nRunning Time:")
	fmt.Println("Iteratif  :", iterTime)
	fmt.Println("Rekursif :", recTime)
}
