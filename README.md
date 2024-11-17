# Файлы для итогового задания

В директории `tests` находятся тесты для проверки API, которое должно быть реализовано в веб-сервере.

Директория `web` содержит файлы фронтенда.


go test -run ^TestApp$ ./tests
go test -run ^TestDB$ ./tests
go test -run ^TestNextDate$ ./tests
go test -run ^TestAddTask$ ./tests