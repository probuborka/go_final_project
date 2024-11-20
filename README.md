# Файлы для итогового задания

В директории `tests` находятся тесты для проверки API, которое должно быть реализовано в веб-сервере.

Директория `web` содержит файлы фронтенда.

//
docker build -t todo-list .
docker run -d -p 7540:7540 todo-list

//
go clean -testcache
//
go test -run ^TestApp$ ./tests
go test -run ^TestDB$ ./tests
go test -run ^TestNextDate$ ./tests
go test -run ^TestAddTask$ ./tests
go test -run ^TestTasks$ ./tests
go test -run ^TestTask$ ./tests
go test -run ^TestEditTask$ ./tests

go test -run ^TestDone$ ./tests
go test -run ^TestDelTask$ ./tests

go test ./tests