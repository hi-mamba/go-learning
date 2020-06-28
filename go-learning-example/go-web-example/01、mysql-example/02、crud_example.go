package init

import "fmt"

func InsertUser() bool {
	tx, err = DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}
	//准备sql语句
	stmt, err = tx.Prepare("")
	if err != nil{
		fmt.Println("Prepare fail")
		return false
	}


}

 