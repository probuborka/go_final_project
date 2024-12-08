# TODO-лист
## Описание проекта
Веб-сервер, который реализует функциональность простейшего планировщика задач.  
## Чистая архитектура
За основу взял идею чистой архитектуры Роберта Мартина (Дяди Боба).  
Понравилось как это видит Тигран Ханагян, видео на ютуб https://www.youtube.com/watch?v=hDwqFRUuykQ&t=48s  
## Список выполенных заданий со звёздочкой
Все задания выполнены
- Возможность определять порт из переменной окружения TODO_PORT.
- Возможность определять путь к файлу базы данных через переменную окружения TODO_DBFILE.
- Правила повторения задач W и M.
- Возможность выбрать задачи через строку поиска.
- Аутентификация
- Создание докер образа
## Инструкция по запуску кода локально 
Запуск кода  
```golang
go run ./cmd/todo/main.go
```   
По умолчанию порт "7540", его можно задать переменной окружения TODO_PORT   
```
export TODO_PORT=<порт>
```  
По умолчанию путь к файлу БД "./db/scheduler.db", его можно задать переменной окружения TODO_DBFILE  
```
export TODO_DBFILE=<путь к файлу БД>
```  
Для использования механизма аутентификации необходимо задать пароль в переменную окружения TODO_PASSWORD  
```
export TODO_PASSWORD=<пароль>
``` 
Открываем TODO-лист в браузере  
```
http://localhost:7540
```
## Инструкция по сборке и запуску проекта через докер 
```
docker build -t todo-list . 
```  
```
docker run -d -p 7540:7540 todo-list  
```
```
http://localhost:7540
```
## Инструкция по запуску тестов
### Параметры для теста (tests/settings.go)
Port = 7540  
DBFile = "../db/scheduler.db"  
FullNextDate = true  
Search = true  
Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.P4Lqll22jQQJ1eMJikvNg5HKG-cKB0hUZA9BZFIG7Jk"
### Запуск тестов
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

