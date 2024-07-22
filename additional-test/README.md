NAMA : Muhammad Habibullah

1. Swap A dan B
```text
A = A + B
B = A - B
A = A - B
```

2. Untuk menemukan bola basket yang beratnya berbeda di antara 8 bola, minimal perlu **3 kali** penimbangan:

	1. Pisahkan 8 bola menjadi 3 grup:
		- Grup 1: Bola 1, 2, 3
		- Grup 2: Bola 4, 5, 6
		- Grup 3: Bola 7, 8

	2. Timbang Grup 1 vs Grup 2:
		- Jika beratnya sama:  Bola yang beratnya berbeda ada di Grup 3.
		- (Grup 1) / 3 = Common Weight*
		- (Grup 2) / 3 = Common Weight

	4. Timbang bola 7 :
		- Jika beratnya berbeda dgn Common Weight: maka bola 7 adalah bola yang beratnya berbeda.
		- Jika beratnya sama dgn Common Weight: maka bola 8 adalah bola yang beratnya berbeda.

3. Ambil masing-masing koin emas dari kantong emas :

	- dari kantong pertama, ambil 1 koin emas
	- dari kantong kedua, ambil 2 koin emas
	- dari kantong ketiga, ambil 3 koin emas
	- dan seterusnya ... maka total adalah 55 koin emas

	55 koin x 10 gram = 550 gram

	- Jika beratnya 549, maka kantong emas palsu adalah kantong pertama.
	- Jika beratnya 548, maka kantong emas palsu adalah kantong kedua.
	- dan seterusnya ...

	maka hanya butuh **1 kali** penimbangan.

4. number to words
	
	```go
	// definikan unit nya
	var units = []string{"", "Satu", "Dua", "Tiga", "Empat", "Lima", "Enam", "Tujuh", "Delapan", "Sembilan"}

	var teens = []string{"Sepuluh", "Sebelas", "Dua Belas", "Tiga Belas", "Empat Belas", "Lima Belas", "Enam Belas", "Tujuh Belas", "Delapan Belas", "Sembilan Belas"}
	
	var tens = []string{"", "", "Dua Puluh", "Tiga Puluh", "Empat Puluh", "Lima Puluh", "Enam Puluh", "Tujuh Puluh", "Delapan Puluh", "Sembilan Puluh"}
	var scales = []string{"", "Ribu", "Juta", "Milyar"}
	```

	```go
	// Hitung modulo dari angka
	func convertNumberToWords(n int) string {
		if n == 0 {
			return "Nol"
		}

		result := ""
		scaleIndex := 0

		for n > 0 {
			if n%1000 != 0 {
				result = numberToWords(n%1000) + " " + scales[scaleIndex] + " " + result
			}
			n /= 1000
			scaleIndex++
		}

		return strings.TrimSpace(result)
	}
	```

	```go
	// Hitung satuan jika modulo kurang dari 1000
	func numberToWords(n int) string {
		if n == 0 {
			return "Nol"
		}
		if n < 10 {
			return units[n]
		}
		if n < 20 {
			return teens[n-10]
		}
		if n < 100 {
			return tens[n/10] + " " + units[n%10]
		}
		if n < 1000 {
			if n/100 == 1 {
				return "Seratus " + numberToWords(n%100)
			}
			return units[n/100] + " Ratus " + numberToWords(n%100)
		}
		return ""
	}
	```

	```go
	// run program
	func main() {
		number := 1000300123
		fmt.Println(convertNumberToWords(number))
	}
	```