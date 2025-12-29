# Analisis Bubble Sort pada Data Suhu

Projek ini mengimplementasikan algoritma Bubble Sort iteratif dan Bubble Sort rekursif untuk mengurutkan data suhu.  
Tujuannya adalah untuk membandingkan efisiensi kedua algoritma dalam hal waktu eksekusi pada berbagai ukuran dataset.  

Program ini membandingkan algoritma Bubble Sort versi iteratif dan rekursif menggunakan data suhu (temperature_celsius) yang diambil dari file CSV (`weather.csv`).  

---

## Studi Kasus
Data suhu dipilih sebagai studi kasus karena merupakan data numerik nyata dengan jumlah besar, sehingga cocok untuk mengamati pengaruh ukuran input terhadap running time algoritma sorting.

---

## Fitur Program
- Membaca data suhu dari CSV (`weather.csv`).  
- Mengurutkan data suhu menggunakan:
  - Bubble Sort Iteratif
  - Bubble Sort Rekursif
- Menampilkan:
  - Suhu terendah
  - Suhu tertinggi
- Menghitung dan membandingkan running time kedua algoritma
- Menampilkan grafik running time  

---

## Algoritma yang Digunakan
- Bubble Sort Iteratif
- Bubble Sort Rekursif

Kedua algoritma memiliki kompleksitas waktu:
- Best Case: O(n)
- Worst Case: O(n²)
- Average Case: O(n²)

---

## Cara Menjalankan
1. Pastikan **Go** sudah terinstall.  
2. Letakkan file `weather.csv` di folder proyek.  
3. Jalankan di terminal:
```bash
go run main.go
