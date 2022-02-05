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

func Test_isAlreadyExist(t *testing.T) {
	db := Connect()
	sample := models.Todo{
		Context: "sample",
		Owner:   "john doe",
	}
	expected := false
	actual := isAlreadyExist(sample.Context, sample.Owner, db)
	assert.Equal(t, expected, actual)
}

func TestGetAllTodos(t *testing.T) {
	db := Connect()
	//var todos []models.Todo
	todos := GetAllTodos(db)
	if assert.NotEqual(t, len(todos), 0) {
		if assert.Equal(t, todos[0].Context, "") {
			t.Errorf("an error occured during the fetch the todos")
		}
	} else {
		log.Printf("there is no todos")
	}
}

func TestGetAllUncompletedTodos(t *testing.T) {
	db := Connect()
	uncTodos := GetAllUncompletedTodos(db)
	if assert.NotEqual(t, len(uncTodos), 0) {
		if assert.NotEqual(t, uncTodos[0].Completed, false) {
			t.Errorf("an error occured during the fetching uncompleted todos")
		}
	} else {
		log.Printf("there is no uncompleted todos")
	}
}

func TestGetAllCompletedTodos(t *testing.T) {
	db := Connect()
	uncTodos := GetAllCompletedTodos(db)
	if assert.NotEqual(t, len(uncTodos), 0) {
		if assert.NotEqual(t, uncTodos[0].Completed, true) {
			t.Errorf("an error occured during the fetching completed todos")
		}
	} else {
		log.Printf("there is no completed todos")
	}
}

func TestCreateTodo(t *testing.T) {
	db := Connect()
	sample_entry := map[string]string{
		"context":   "sample context",
		"owner":     "john doe",
		"completed": "false",
	}
	todo, err := CreateTodo(sample_entry["context"], sample_entry["owner"], sample_entry["completed"], db)
	assert.Nil(t, err, nil)
	assert.Equal(t, sample_entry["context"], todo.Context)
}

func TestDeleteTodoById(t *testing.T) {
	db := Connect()
	id := 3
	todo, err := DeleteTodoById(int64(id), db)
	assert.Nil(t, err, nil)

	if assert.NotEqual(t, todo.Context, "") {
		log.Printf("todo bulunamadı")
	}
}

func TestUpdateStatueOfCompleteById(t *testing.T) {
	db := Connect()
	id := int64(3)
	completed := "false"
	todo, err := UpdateStatueOfCompleteById(id, completed, db)
	assert.Nil(t, err)
	if assert.Equal(t, todo.Context, "") {
		log.Println("ilgili todo yok")
	} else {
		assert.Equal(t, todo.Completed, "false")
	}

}
