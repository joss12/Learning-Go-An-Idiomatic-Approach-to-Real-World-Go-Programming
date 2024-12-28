package main

import (
	"context"
	"database/sql"
	"fmt"
	// "io"
	// "log"
	// "os"
)

var pl = fmt.Println

// func deferExpl() int {
// 	a := 10
// 	defer func(val int) {
// 		pl("First:", val)
// 	}(a)
// 	a = 20
// 	defer func(val int) {
// 		pl("second:", val)
// 	}(a)
// 	a = 30
// 	pl("exiting:", a)
// 	return a
// }

func DoSomeInsert(ctx context.Context, db *sql.DB, value1, value2 string) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err == nil {
			err = tx.Commit()
		}
		if err != nil {
			tx.Rollback()
		}
	}()
	_, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) values $1", value1)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// if len(os.Args) < 2 {
	// 	log.Fatal("on file specified")
	// }
	// f, err := os.Open(os.Args[1])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// data := make([]byte, 2048)
	// for{
	// 	count, err := f.Read(data)
	// 	os.Stdout.Write(data[:count])
	// 	if err != nil{
	// 		if err != io.EOF{
	// 			log.Fatal(err)
	// 		}
	// 		break
	// 	}
	// }
	// deferExpl()
	
	pl(DoSomeInsert())
}
