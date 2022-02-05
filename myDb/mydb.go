package mydb

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Kivanc10/task-app-for-westerOps/models"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var cfg models.Config

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(cfg *models.Config, filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func readEnv(cfg *models.Config, filepath string) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}

/*
func init() { //---------
	readFile(&cfg,"config.yml")
	readEnv(&cfg,"config.yml")
}
*/
func dsn1(filepath string) string {
	readFile(&cfg, filepath)
	readEnv(&cfg, filepath)
	return fmt.Sprintf("%s:%s@tcp(%s)/", cfg.Database.Username, cfg.Database.Password, cfg.Server.Host) //  username, password, host
}

func dsn2(filepath string) string {
	readFile(&cfg, filepath)
	readEnv(&cfg, filepath)
	// fmt.Println(cfg.Database, " ", cfg.Database.DbName)
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", cfg.Database.Username, cfg.Database.Password, cfg.Server.Host, cfg.Database.DbName) // username, password, host, database_name1
}

/*
func isTestRun() bool {
	return flag.Lookup("test.v").Value.(flag.Getter).Get().(bool)
}
*/

func Connect() *gorm.DB {
	var filepath string
	if flag.Lookup("test.v") == nil { // normal execution
		filepath = "config.yml"
	} else { // test execution
		filepath = "../config.yml"
	}

	db, err := gorm.Open(mysql.Open(dsn1(filepath)), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Exec("CREATE DATABASE IF NOT EXISTS " + cfg.Database.DbName)
	db, err = gorm.Open(mysql.Open(dsn2(filepath)), &gorm.Config{})

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

func UpdateStatueOfCompleteById(id int64, completed string, db *gorm.DB) (models.Todo, error) {
	var todo models.Todo
	db.Where(map[string]interface{}{"id": id}).Find(&todo)
	if todo.Context != "" && todo.Owner != "" {
		//db.Where("id = ?", id).Update("completed", completed)
		todo.Completed = completed
		db.Save(&todo)
		return todo, nil
	} else {
		return models.Todo{}, errors.New("İlgili todo bulunamadı")
	}
}
