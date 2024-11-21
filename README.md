# TODO-лист 😃

## Описание проекта
Веб-сервер, который реализует функциональность простейшего планировщика задач.

## Список выполенных заданий со звёздочкой
Все задания выполнены
- Возможность определять порт из переменной окружения TODO_PORT.
- Возможность определять путь к файлу базы данных через переменную окружения TODO_DBFILE.
- Правила повторения задач W и M.
- Возможность выбрать задачи через строку поиска.
- Аутентификация
- Создание докер образа

## Инструкция
### Запуску кода локально 
```
cd <проект>  
```
```golang
go run ./cmd/todo/main.go
```   
### Докер
```
docker build -t todo-list . 
```  
docker run -d -p 7540:7540 todo-list  
адрес http://localhost:7540

## Инструкция по запуску тестов
### Параметры для теста файл tests/settings.go
Port = 7540  
DBFile = "../db/scheduler.db"  
FullNextDate = true  
Search = true  
Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.P4Lqll22jQQJ1eMJikvNg5HKG-cKB0hUZA9BZFIG7Jk"

### Тесты
cd <проект>  
```golang
go test ./tests
```
```golang
go test -run ^TestApp$ ./tests
```
```golang
go test -run ^TestDB$ ./tests
```
```golang
go test -run ^TestNextDate$ ./tests
```
```golang
go test -run ^TestAddTask$ ./tests
```
```golang
go test -run ^TestTasks$ ./tests
```
```golang
go test -run ^TestTask$ ./tests
```
```golang
go test -run ^TestEditTask$ ./tests
```
```golang
go test -run ^TestDone$ ./tests
```
```golang
go test -run ^TestDelTask$ ./tests
```

