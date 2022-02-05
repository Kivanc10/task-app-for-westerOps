# task-app-for-westerOps
Todo app for company named WesterOps

base_url = `http://157.245.79.10:8080/`


## add new todo
  - `POST` --> `http://157.245.79.10:8080/addTodo`
  
  - `curl -X POST http://157.245.79.10:8080/addTodo -H "Content-Type: application/json" -d '{"owner": "user_westerops", "context": "from web","completed" : "false"}'`
      

## get all todos
  - `GET` ---> `http://157.245.79.10:8080/todos`
  - `curl -X GET http://157.245.79.10:8080/todos`
  

## get all completed todos
  - `GET` ---> `http://157.245.79.10:8080/todos/completed`
  - `curl -X GET http://157.245.79.10:8080/todos/completed`
 

## get all uncompleted todos
  - `GET` ---> `http://157.245.79.10:8080/todos/uncompleted`
  - `curl -X GET http://157.245.79.10:8080/todos/uncompleted`

## delete todo by id
  - `DELETE` --> `http://157.245.79.10:8080/delete/{id}`
  - `curl -X DELETE http://157.245.79.10:8080/delete/3`
 
## update todo's state of complete
  - `POST` ---> `http://157.245.79.10:8080/update/{id}`
  - `curl -X POST http://157.245.79.10:8080/update/3 -H "Content-Type: application/json" -d '{"completed" : "true"}'`
  
