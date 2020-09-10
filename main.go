package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shakilbd009/go-sql-client/sqlClient"
)

const (
	getUserQuery = "select * from users where id=%d;"
)

var (
	dbClient sqlClient.SqlClient
)

type User struct {
	Email string
	Id    int64
}

func init() {
	var err error
	dbClient, err = sqlClient.Open("", "")
	if err != nil {
		panic(err)
	}
}

func main() {

}

func GetUser(id int) (*User, error) {
	rows, err := dbClient.Query(fmt.Sprintf(getUserQuery, id))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var usr User
	for rows.HasNext() {
		if err := rows.Scan(&usr.Id, &usr.Email); err != nil {
			return nil, err
		}
	}
	return &usr, nil
}
