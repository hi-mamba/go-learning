package obj

/**
定义 员工对象，以及一些数据

https://coolshell.cn/articles/21164.html#%E5%91%98%E5%B7%A5%E4%BF%A1%E6%81%AF
*/
type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   int
}

var List = []Employee{
	{"Hao", 44, 0, 8000},
	{"Bob", 34, 10, 5000},
	{"Alice", 23, 5, 9000},
	{"Jack", 26, 0, 4000},
	{"Tom", 48, 9, 7500},
	{"Marry", 29, 0, 6000},
	{"Mike", 32, 8, 4000},
}
