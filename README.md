# go-api-sample-todo
This is a simple RestAPI using gin and gorm.

## Usage
### Run Server
```
$ git clone https://github.com/Ixy-194/go-api-sample-todo.git
$ cd go-api-sample-todo
$ docker-compose up -d 
$ docker exec -d go-api-sample-todo go run cmd/go-api-sample-todo/main.go
```
### Unit tests
```
$ go-api-sample-todo % docker exec go-api-sample-todo go test ./... 
```

### End points
| Method  | Path | Description |
| ------------- | ------------- | ------------- |
| GET  | /todo  | Get all task list |
| GET  | /todo/{id}  | Get a task |
| POST  | /todo | Create a new task |
| PUT  | /todo/{id}  | Update a task |
| DELETE  | /todo/{id}  | Delete a task |

### API call samples
```
# Get all task list
$ curl -i -XGET localhost/todo

# Get a task
$ curl -i -XGET localhost/todo/1

# Create a new task
$ curl -i localhost/todo -H "Content-Type: application/json" -X POST -d '{"task": "test1"}' 

# Update a task
$ curl -i localhost/todo/1 -H "Content-Type: application/json" -X PUT -d '{"task": "test1","status": "done"}'

# Delete a task
$ curl -i localhost/todo/1 -X DELETE

```
## Other
### Login mysql
```
$ docker exec -it mysql-db mysql -utodo-app -ptodo-app -Dtodo-app
```
