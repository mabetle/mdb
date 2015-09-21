package models

type DemoTable struct {
	Id       int64
	DemoName string
	DemoAge  int64
}

func (m DemoTable) TableName() string {
	return "demo_table"
}

type DemoXorm struct {
	Id   string `xorm:"varchar(50) not null pk 'ID'"`
	Name string `xorm:"varchar(50) not null default '' 'Name'"`
}

func (m DemoXorm) TableName() string {
	return "demo_xorm"
}
