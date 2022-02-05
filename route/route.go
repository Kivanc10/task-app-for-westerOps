package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Kivanc10/task-app-for-westerOps/models"
	mydb "github.com/Kivanc10/task-app-for-westerOps/myDb"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = mydb.Connect()
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse error :%v\n", err)
		return
	}

	owner := r.FormValue("owner")
	context := r.FormValue("context")
	completed := r.FormValue("completed")
	//completed, _ := strconv.ParseBool(strCompl)
	if owner == "" || context == "" || completed == "" { // via postman
		rBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			fmt.Fprintf(w, "An error occured during the get credentials")
		}
		if todo, err := models.ProcessToJson(rBody); err != nil { // process to json
			w.WriteHeader(http.StatusBadGateway)
			fmt.Fprintf(w, "An error occured during the get todo struct")
		} else {
			if todo.Owner != "" && todo.Context != "" && todo.Completed != "" {
				todo, err = mydb.CreateTodo(todo.Context, todo.Owner, todo.Completed, db)
				if err != nil {
					w.WriteHeader(http.StatusNotAcceptable)
					fmt.Fprintf(w, err.Error())
				}
				json.NewEncoder(w).Encode(todo)
			}
		}
	}

}

func getAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := mydb.GetAllTodos(db)
	json.NewEncoder(w).Encode(todos)
}

func getAllUncompletedTodos(w http.ResponseWriter, r *http.Request) {
	uncTodos := mydb.GetAllUncompletedTodos(db)
	json.NewEncoder(w).Encode(uncTodos)
}

func getAllCompletedTodos(w http.ResponseWriter, r *http.Request) {
	compTodos := mydb.GetAllCompletedTodos(db)
	json.NewEncoder(w).Encode(compTodos)
}

func deleteTodoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	i, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		panic(err)
	}
	if _, err := mydb.DeleteTodoById(i, db); err != nil { // todo
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode("Todo başarılı bir şekilde kaldırıldı")

	}
}

func updateStateOfComplete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	i, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		panic(err)
	}

	completed := r.FormValue("completed")
	if completed == "" { // via postman
		rBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			fmt.Fprintf(w, "An error occured during the get credentials")
		}
		if todo, err := models.ProcessToJson(rBody); err != nil {
			w.WriteHeader(http.StatusBadGateway)
			fmt.Fprintf(w, "An error occured during the get todo struct")
		} else {
			if todo.Completed != "" {
				todo, err := mydb.UpdateStatueOfCompleteById(i, todo.Completed, db)
				if err != nil {
					w.WriteHeader(http.StatusNotAcceptable)
					fmt.Fprintf(w, err.Error())
				} else {
					json.NewEncoder(w).Encode(todo)
				}

			}
		}
	}
}

func HandleRequest() {
	//db = mydb.Connect()
	fmt.Println("connected to db")
	r := mux.NewRouter()
	r.HandleFunc("/addTodo", addTodo).Methods("POST")
	r.HandleFunc("/todos", getAllTodos)
	r.HandleFunc("/todos/uncompleted", getAllUncompletedTodos)
	r.HandleFunc("/todos/completed", getAllCompletedTodos)
	r.HandleFunc("/delete/{id}", deleteTodoById).Methods("DELETE")
	r.HandleFunc("/update/{id}", updateStateOfComplete).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
