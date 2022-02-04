package mydb

import (
	"errors"
	"fmt"
	"log"

	"github.com/Kivanc10/task-app-for-westerOps/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	username       = "root"
	password       = "123456"
	host           = "mysql"
	database_name1 = "task_db"
)

func dsn1() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/", username, password, host)
}

func dsn2() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, host, database_name1)
}

func Connect() *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn1()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Exec("CREATE DATABASE IF NOT EXISTS " + database_name1)
	db, err = gorm.Open(mysql.Open(dsn2()), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("it is connect to the db ", db)
	fmt.Println("connecttedddd")
	return db
}

func insert(db *gorm.DB, todo models.Todo) (int64, error) {
	if isAlreadyExist(todo.Context, todo.Owner, db) {
		return -1, errors.New("BU TODO ZATEN MEVCUT")
	} else {
		result := db.Create(&todo)
		if result.Error != nil {
			panic(result.Error)
		}
		return todo.Id, result.Error
	}

}

// is already exist()
func isAlreadyExist(context, owner string, db *gorm.DB) bool {
	temp := models.Todo{}
	//fake := db.Where("owner = ? AND context = ?", owner, context).First(&temp)
	db.Where(map[string]interface{}{"owner": owner, "context": context}).Find(&temp)
	fmt.Println("fonded temp todo --> ", temp)
	//fmt.Println("fake db ---> ", fake)
	if temp.Context != "" && temp.Owner != "" && temp.Completed != "" {
		return true
	} else {
		return false
	}

}

func GetAllTodos(db *gorm.DB) []models.Todo {
	var todos []models.Todo
	result := db.Find(&todos)
	fmt.Println(result)
	return todos
}

func GetAllUncompletedTodos(db *gorm.DB) []models.Todo {
	var todos []models.Todo
	result := db.Where(map[string]interface{}{"completed": "false"}).Find(&todos)
	fmt.Println(result)
	return todos
}

func GetAllCompletedTodos(db *gorm.DB) []models.Todo {
	var todos []models.Todo
	result := db.Where(map[string]interface{}{"completed": "true"}).Find(&todos)
	fmt.Println(result)
	return todos
}

func CreateTodo(context, owner, completed string, db *gorm.DB) (*models.Todo, error) {
	db.AutoMigrate(&models.Todo{}) //--
	var todo models.Todo
	todo.Context = context
	todo.Owner = owner
	todo.Completed = completed
	id, err := insert(db, todo)
	if err != nil {
		log.Println("an error occured during the inserting ", err)
		return &models.Todo{}, err
	}
	todo.Id = id
	return &todo, nil
}

func DeleteTodoById(id int64, db *gorm.DB) (models.Todo, error) {
	var todo models.Todo
	db.Where(map[string]interface{}{"id": id}).Find(&todo)
	if todo.Context != "" && todo.Owner != "" {
		//db.Delete(todo)
		db.Where("id = ?", id).Delete(&models.Todo{})
		return todo, nil
	} else {
		return models.Todo{}, errors.New("İlgili todo bulunamadı")
	}

}
