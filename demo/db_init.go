package demo


import (
	"github.com/mabetle/mdb"
)

var(

	DEMO_TABLE_CREATE_SQL=`
		create table demo_table(
			id int not null primary key,
			DemoName varchar(50) not null,
			DemoAge int not null
		)
	`

	DEMO_INSERT_SQL=`
		insert into demo_table (id,DemoName,DemoAge)
		values (1,'demo', 30)
	`

	DEMO_INSERT_SQL2=`
		insert into demo_table (id,DemoName,DemoAge)
		values (2,'demo2', 10)
	`


)


func InitDB(db *mdb.Sql){
	db.Exec("drop table demo_table")
	db.Exec(DEMO_TABLE_CREATE_SQL)
	db.Exec(DEMO_INSERT_SQL)
	db.Exec(DEMO_INSERT_SQL2)
}






