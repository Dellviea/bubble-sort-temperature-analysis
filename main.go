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
		tempStr := records[i][7] 

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
		fmt.Println("Gagal membaca file CSV:", err)
		return
	}

	var n int
	fmt.Print("Masukkan jumlah data yang diuji (n): ")
	fmt.Scan(&n)

	if n > len(data) {
	n = len(data)
	}

	dataN := make([]float64, n)
	copy(dataN, data[:n])

	iterData := make([]float64, n)
	recData := make([]float64, n)
	copy(iterData, dataN)
	copy(recData, dataN)

	startIter := time.Now()
	bubbleSortIterative(iterData)
	iterTime := time.Since(startIter)

	startRec := time.Now()
	bubbleSortRecursive(recData, len(recData))
	recTime := time.Since(startRec)

	fmt.Println("Total data suhu :", len(data))
	fmt.Println("Jumlah data yang diuji (n):", n)
	fmt.Println("\n=== HASIL SORTING ===")
	fmt.Printf("Suhu terendah : %.2f °C\n", iterData[0])
	fmt.Printf("Suhu tertinggi: %.2f °C\n", iterData[len(iterData)-1])

	fmt.Println("\n=== RUNNING TIME (ms) ===")
	fmt.Printf("Iteratif  : %.2f ms\n", float64(iterTime.Microseconds())/1000)
	fmt.Printf("Rekursif : %.2f ms\n", float64(recTime.Microseconds())/1000)

}
