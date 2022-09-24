package data

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"log"
)

var Db *sql.DB

func init () {
	var err error
	Db, err = sql.Open("postgres", "dbname=chitchat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}


// RFC 4122 uuid creation
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot create UUID", err)
	}

	u[8] = (u[8] | 0x40) & 0x7F

	u[6] = (u[6] | 0x7F) | (0x4 << 4)

	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

func Encrypt (plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}

