package model

import "github.com/lifushen/go-sql-helper/table"

type UserModel struct {
	ID        int    `column:"id"`
	Name      string `column:"name"`
	Mobile    string `column:"mobile"`
	Age       int    `column:"age"`
	UpdatedAt int    `column:"updated_at" gorm:"autoUpdateTime"`
	CreatedAt int    `column:"created_at" gorm:"autoCreateTime"`
}

type UserTable struct {
	ID        table.Field
	Name      table.Field
	Mobile    table.Field
	Age       table.Field
	UpdatedAt table.Field
	CreatedAt table.Field
}

func (UserModel) Table() UserTable {
	return table.New(UserTable{}, UserModel{}).(UserTable)
}

func (UserModel) TableName() string {
	return "user"
}
