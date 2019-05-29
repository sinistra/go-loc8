package main

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"log"
	"time"
)

type MessageDetailRecord struct {
	AccountId    	int
	SubAccountId    int
}

func main() {

	dbb, _ := sql.Open("postgres", "postgres://postgres:mysecretpassword@localhost:5432/loc8?sslmode=disable")
	dbb.SetMaxIdleConns(10)
	dbb.SetMaxOpenConns(10)
	dbb.SetConnMaxLifetime(0)

	a := time.Now()
	txn, err := dbb.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, _ := txn.Prepare(pq.CopyIn("messagedetailrecord", "accountid", "subaccountid")) // MessageDetailRecord is the table name
	m := &MessageDetailRecord{
		AccountId:    123456,
		SubAccountId: 123434,
	}
	mList := make([]*MessageDetailRecord, 0, 100)
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		mList = append(mList, m)
	}
	fmt.Println(m)
	for _, user := range mList {
		_, err := stmt.Exec(int64(user.AccountId), int64(user.SubAccountId))
		if err != nil {
			log.Fatal(err)
		}
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = txn.Commit()
	if err != nil {
		log.Fatal(err)
	}
	delta := time.Now().Sub(a)
	fmt.Println(delta.Nanoseconds())
	fmt.Println("Program finished successfully")
}
