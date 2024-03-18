package main

import (
	"fmt"
	"os"
	"strconv"
)

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Input tidak sesuai.")
		return
	}

	absen, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Input tidak valid.")
		return
	}

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

	if absen > 0 && absen <= len(data) {
		fmt.Println("Data teman dengan absen no:", absen)
		fmt.Println("--------")
		fmt.Println("Nama: ", data[absen-1].Nama)
		fmt.Println("Alamat: ", data[absen-1].Alamat)
		fmt.Println("Pekerjaan: ", data[absen-1].Pekerjaan)
		fmt.Println("Alasan memilih kelas Golang: ", data[absen-1].Alasan)
	} else {
		fmt.Println("Data absen peserta tidak ditemukan")
	}

}
