package mydb

import (
	"log"
	"testing"

	"github.com/Kivanc10/task-app-for-westerOps/models"
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	//var expected *gorm.DB
	defer func() {
		if err := recover(); err != nil {
			log.Println("an error occured while connecting to db")
			t.Errorf(" ------- unable to connect to db ---------")
		}
	}()
	Connect()
	// if db := Connect(); db != nil {
	// 	log.Println("an error occured while connecting to db")
	// }

}

func TestInsert(t *testing.T) {
	//var db *gorm.DB
	db := Connect()
	temp := models.Todo{
		Context:   "Deneme",
		Owner:     "John Doe",
		Completed: "false",
	}
	defer func() {
		if err := recover(); err != nil {
			log.Println("an error occured while connecting to db")
			t.Errorf(" ------- unable to insert element into db ---------")
		}
	}()

	id, err := insert(db, temp)
	assert.Nil(t, err)
	assert.NotEqual(t, id, -1)

}
