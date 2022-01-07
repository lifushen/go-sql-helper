package test

import (
	"github.com/lifushen/go-sql-helper/test/model"
	"gorm.io/gorm"
	"testing"
)

func TestQuery(t *testing.T) {

	var userModel = model.UserModel{}
	var userTable = userModel.Table()

	t.Run("TestQuery1", func(t *testing.T) {
		var db gorm.DB
		db.Debug().Where(userTable.ID.EQ(1)).Take(&userModel)
	})

	t.Run("TestQuery2", func(t *testing.T) {
		var db gorm.DB
		db.Debug()
		db.Where(userTable.Age.BETWEEN(10, 20))
		db.Order(userTable.Age.DESC())
		db.Find(&userModel)
	})
}
