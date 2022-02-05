package route

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Kivanc10/task-app-for-westerOps/models"
)

func Test_getAllTodos(t *testing.T) {
	//HandleRequest()
	req, err := http.NewRequest("GET", "/todos", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAllTodos)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func Test_addTodo(t *testing.T) {
	sample_todo := models.Todo{
		Context:   "sample context",
		Owner:     "john doe",
		Completed: "false",
	}
	byte_todo, _ := json.Marshal(sample_todo)
	req, err := http.NewRequest("GET", "/addTodo", bytes.NewReader(byte_todo))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addTodo)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	data, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("an error occured during the read the body")
	}
	var result models.Todo
	json.Unmarshal(data, &result)
	assert.Equal(t, sample_todo.Context, result.Context)
	assert.Equal(t, sample_todo.Owner, result.Owner)
	//json.Unmarshal(rr.Body,&result)
	//expected := `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"id":19,"context":"something different","owner":"leyla","completed":"true"}`
}

func Test_deleteTodoById(t *testing.T) {
	// id := map[string]interface{}{
	// 	"id" : 20,
	// }
	// byteId,_ := json.Marshal(id)
	id := "20"
	path := fmt.Sprintf("/delete/%s", id)
	req, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addTodo)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := "Todo başarılı bir şekilde kaldırıldı"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func Test_getAllUncompletedTodos(t *testing.T) {

}

func Test_getAllCompletedTodos(t *testing.T) {

}

func Test_updateStateOfComplete(t *testing.T) {

}
