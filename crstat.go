package main

import "fmt"
import _ "database/sql"
import _ "github.com/mattn/go-sqlite3"

func Crawstat() {
	var theExecTimeDbS ExecTimeDbS
	theExecTimeDbS.Boot()
	theExecTimeDbS.LoadConfigure()
	row, _ := theExecTimeDbS.Dbc.Query("SELECT uid,count(*) FROM lastcrawfeed GROUP BY uid")
	for row.Next() {
		var uid string
		var count int
		row.Scan(&uid, &count)
		//fmt.Print(uid)
		fmt.Printf("uid %v:count %v\n", uid, count)

	}
	row, _ = theExecTimeDbS.Dbc.Query("SELECT count(*) FROM lastcrawfeed")
	for row.Next() {
		var count int
		row.Scan(&count)
		//fmt.Print(uid)
		fmt.Printf("countall %v\n", count)

	}
}
