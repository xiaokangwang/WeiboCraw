package main

import "fmt"
import "database/sql"
import _ "github.com/mattn/go-sqlite3"
import "io/ioutil"
import "os"

func Dbafc(argm []string) {

	var err error
	db, err := sql.Open("sqlite3", "file:weibopulldb.sqlite")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	switch argm[0] {

	case "ExecRawSql", "ERS", "ers":

		if len(argm) < 2 {
			fmt.Println("assert fail: Arg len")
			return
		}

		fd, err := os.Open(argm[1])
		defer fd.Close()

		if err != nil {
			fmt.Println(err)
		}

		ftx, err := ioutil.ReadAll(fd)
		if err != nil {
			fmt.Println(err)
		}

		execsql := string(ftx)

		transaction, _ := db.Begin()

		res, err := transaction.Exec(execsql)

		if err != nil {
			fmt.Println(err)
			ir, _ := res.RowsAffected()

			err = transaction.Rollback()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Rollbacked SUCCESSFULLY.")
			}
			fmt.Printf("%v Rows was Affected, HOWEVER tx was rollbacked as an error was throwen.", ir)
		}

		ir, _ := res.RowsAffected()

		err = transaction.Commit()

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Commited SUCCESSFULLY.")
		}

		fmt.Printf("%v Rows was Affected.", ir)

	case "SetConf", "SF", "sf":

		if len(argm) < 3 {
			fmt.Println("assert fail: Arg len")
			return
		}

		_, err := db.Exec("DELETE FROM programma_configure WHERE confname=?;", argm[1])

		if err != nil {
			fmt.Println(err)
		}

		_, err = db.Exec("INSERT INTO programma_configure(confname,cfncval) VALUES (?,?)", argm[1], argm[2])

		if err != nil {
			fmt.Println(err)
		}

	case "CatConf", "CF", "cf":

		if len(argm) < 2 {
			fmt.Println("assert fail: Arg len")
			return
		}

		row := db.QueryRow("SELECT cfncval FROM programma_configure WHERE confname=?", argm[1])
		var tocat string
		err := row.Scan(&tocat)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(tocat)
	}
}
