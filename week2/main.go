package main

import (
	"database/sql"
	"log"

	"github.com/pkg/errors"
)

func main() {
	db, err := sql.Open("driver-name", "database=test1")
	if err != nil {
		log.Fatal(err)
	}
	// 个人认为sql.ErrNoRows是否向上抛取决于业务，如果业务认为没有查到结果就是个错误就携带业务信息往上抛，如果认为没有查到结果属于正常case就降级处理
	// 1.业务认为没有查到结果就是个错误就携带业务信息往上抛：
	username,err1:=query1(id,db)
	if 

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
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("data not found")
		return "", nil
	} else if err != nil {
		return "", errors.Wrapf(err, "failed to select username %s from users where id=%d", username, id)
	}
	return username, nil
}
