package gomysql

import (
    "database/sql"
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)

dsn := "localhost:root@tcp(localhost:3306)/liga2"
db, err := sql.Open("mysql", dsn)
if err != nil {
    // Handle error
}
defer db.Close()

func getUsers(c *gin.Context) {
    rows, err := db.Query("SELECT * FROM klasemen")
    if err != nil {
        // Handle error
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        err := rows.Scan(&id_klasemen, &liga, &kode_tim, &tim, &total_poin,
						&kalah, &menang, &seri, &posisi, &jumlah_pemain)
        if err != nil {
            // Handle error
        }
        users = append(klasemen, user)
    }

    c.JSON(http.StatusOK, klasemen)
}

router := gin.Default()
router.GET("/klasemen", getUsers)
router.Run(":8080")