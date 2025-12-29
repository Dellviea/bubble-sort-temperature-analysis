package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
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

func sortHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	nStr := r.URL.Query().Get("n")
	n, err := strconv.Atoi(nStr)
	if err != nil || n <= 0 {
		http.Error(w, "Parameter n tidak valid", http.StatusBadRequest)
		return
	}

	data, err := readTemperatureCSV("weather.csv")
	if err != nil {
		http.Error(w, "Gagal membaca file CSV", http.StatusInternalServerError)
		return
	}

	if n > len(data) {
		n = len(data)
	}

	dataN := make([]float64, n)
	copy(dataN, data[:n])

	iterData := make([]float64, n)
	recData := make([]float64, n)

	repeat := 1
	if n <= 100 {
		repeat = 1000000
	} else if n <= 1000 {
		repeat = 1000
	} else if n <= 10000 {
		repeat = 100
	}

	startIter := time.Now()
	for i := 0; i < repeat; i++ {
		copy(iterData, dataN)
		bubbleSortIterative(iterData)
	}
	iterTime := time.Since(startIter)

	startRec := time.Now()
	for i := 0; i < repeat; i++ {
		copy(recData, dataN)
		bubbleSortRecursive(recData, len(recData))
	}
	recTime := time.Since(startRec)

	response := map[string]interface{}{
	"total_data": len(data),
	"n":          n,
	"repeat":     repeat,
	"min":        iterData[0],
	"max":        iterData[len(iterData)-1],
	"iterative_ms": float64(iterTime.Nanoseconds()) / 1e6 / float64(repeat),
	"recursive_ms": float64(recTime.Nanoseconds()) / 1e6 / float64(repeat),
	}

	json.NewEncoder(w).Encode(response)
}


func main() {
	http.HandleFunc("/sort", sortHandler)

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
