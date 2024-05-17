Start: go run main.go

Testing: 
        Postman:
        GET http://localhost:8000/
        "Welcome to the TODO API!"
        
        POST http://localhost:8000/todos
        Body -> raw -> JSON
        Input:
        "{
        "title": "Shopping list",
        "content": "Milk, bread, eggs"
        }
        Output: {"id":1,"title":"Shopping list","content":"Milk, bread, eggs"}

        GET especifico de ID http://localhost:8000/todos/1     (1 = id do "todo")
        Output: {"id":1,"title":"Shopping list","content":"Milk, bread, eggs"}

        GET all http://localhost:8000/todos/all
        Output: {"total_count":1,"todos":[{"id":1,"title":"Shopping list","content":"Milk, bread, eggs"}]}

        PUT (Editar todo) http://localhost:8000/todos/1
        Body -> raw -> JSON
        Input:
        {
        "title": "Shopping list",
        "content": "Milk, bread, eggs, apples"
        }
        Output:
        "{"id":1,"title":"Shopping list","content":"Milk, bread, eggs, apples"}"

        DELETE http://localhost:8000/todos/1    (1 = id do "todo")
        
        
        

        
        
        

        
