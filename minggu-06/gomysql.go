package gomysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Membuat koneksi ke database
	db, err := sql.Open("mysql", "localhost:root@tcp(localhost:3306)/;liga2")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Membaca data dari tabel
	rows, err := db.Query("SELECT * FROM tim")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Menampilkan data
	for rows.Next() {
		var id int
		var nama string
		err = rows.Scan(&id, &nama)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(id, nama)
	}

	// Menangani error saat membaca data
	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
}
