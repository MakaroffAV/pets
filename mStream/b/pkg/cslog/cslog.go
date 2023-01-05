package cslog

import (
	"b/cnf"
	"fmt"
	"time"
)

func Fail(c string) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Show `FAIL` log
	*/

	fmt.Printf(
		"INFO -- %s -- %s \n",
		time.Now().Format("2006-01-02 15:04:05"),
		cnf.Logs[c],
	)

}

func Info(c string) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Show `INFO` log
	*/

	fmt.Printf(
		"INFO -- %s -- %s \n",
		time.Now().Format("2006-01-02 15:04:05"),
		cnf.Logs[c],
	)

}
