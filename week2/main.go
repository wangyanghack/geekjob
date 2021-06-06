package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

func main() {
	db, err := sql.Open("driver-name", "database=test1")
	if err != nil {
		log.Fatal(err)
	}
	var id int
	// 分两种情况，个人认为sql.ErrNoRows是否向上抛取决于业务，如果业务认为没有查到结果就是个错误就携带业务信息往上抛，如果认为没有查到结果属于正常情况就降级处理
	// case1.如果认为没有查到结果属于正常case就降级处理：
	username1, err1 := query1(id, db)
	if err1 != nil {
		fmt.Printf("original error: %T %v\n", errors.Cause(err1), errors.Cause(err1))
		fmt.Printf("stack trace:\n%+v\n", err1)
		os.Exit(1)
	}

	// case2.如果业务认为没有查到结果就是个错误就携带业务信息往上抛：
	username2, err2 := query2(id, db)
	if err2 != nil {
		fmt.Printf("original error: %T %v\n", errors.Cause(err2), errors.Cause(err2))
		fmt.Printf("stack trace:\n%+v\n", err2)
		os.Exit(1)
	}

}

func query1(id int, db *sql.DB) (string, error) {
	var username string
	err := db.QueryRow("select username from users where id = ?", id).Scan(&username)
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("data not found")
		return "", nil
	} else if err != nil {
		return "", errors.Wrapf(err, "failed to select username %s from users where id=%d", username, id)
	}
	return username, nil
}

func query2(id int, db *sql.DB) (string, error) {
	var username string
	err := db.QueryRow("select username from users where id = ?", id).Scan(&username)
	if err != nil {
		return "", errors.Wrapf(err, "failed to select username %s from users where id=%d", username, id)
	}
	return username, nil
}
