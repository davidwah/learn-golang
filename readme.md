# Assignment 1

## Task
Membuat service CLI dengan menampilkan data teman-teman di kelas.  
**data teman**
```
data := []Teman{
		{
			Nama:      "Agus",
			Alamat:    "Jalan Raya Utara 01",
			Pekerjaan: "Software Engineer",
			Alasan:    "Golang untuk membuat aplikasi backend",
		},
		{
			Nama:      "Badur",
			Alamat:    "Jalan Raya Selatan 02",
			Pekerjaan: "Software Engineer",
			Alasan:    "Golang untuk membuat aplikasi backend",
		},
		{
			Nama:      "Hilman",
			Alamat:    "Jalan Raya Barat 03",
			Pekerjaan: "Software Engineer",
			Alasan:    "Golang untuk membuat aplikasi backend",
		},
		{
			Nama:      "Zulfikar",
			Alamat:    "Jalan Raya Timur 04",
			Pekerjaan: "Software Engineer",
			Alasan:    "Golang untuk membuat aplikasi backend",
		},
	}
```


## Run
Menjalankan service dengan perintah `go run main.go <no-data-teman>`
```
...
$ go run main.go 1
Data teman dengan absen no: 1
--------
Nama:  Agus
Alamat:  Jalan Raya Utara 01
Pekerjaan:  Software Engineer
Alasan memilih kelas Golang:  Golang untuk membuat aplikasi backend

...
$ go run main.go 2
Data teman dengan absen no: 2
--------
Nama:  Badur
Alamat:  Jalan Raya Selatan 02
Pekerjaan:  Software Engineer
Alasan memilih kelas Golang:  Golang untuk membuat aplikasi backend

...
$ go run main.go 10
Data absen peserta tidak ditemukan


```