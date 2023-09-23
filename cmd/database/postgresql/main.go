package main
 

// This is importing for a side effect. The underscore is in the alias position. This allows you to circumvent 
// the error (or automatic removal if using something like goimports) of an unused import.

// When github.com/lib/pq is imported it runs an init() function that registers the postgres database driver.
// https://github.com/lib/pq/blob/d34b9ff171c21ad295489235aec8b6626023cd04/conn.go#L48-L50


import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)
 
const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "postgres5432"
    dbname   = "mytestdb"
)
 
func main() {
        // connection string
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
         
        // open database
    db, err := sql.Open("postgres", psqlconn)
    CheckError(err)
     
        // close database
    defer db.Close()
 
        // check db
    err = db.Ping()
    CheckError(err)
 
    fmt.Println("Connected!")
}
 
func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}
