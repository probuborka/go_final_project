package service

type service struct {
	Task task
}

func New(taskDB dbTask) *service {
	return &service{
		Task: newTask(taskDB),
	}
}
