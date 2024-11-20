package service

type service struct {
	Task          task
	Authorization authorization
}

func New(taskDB dbTask) *service {
	return &service{
		Task:          newTask(taskDB),
		Authorization: newAuthorization(),
	}
}
