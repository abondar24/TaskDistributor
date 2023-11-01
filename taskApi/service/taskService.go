package service

type TaskService interface {
	CreateTask(name *string) (string, error)
	UpdateTask(id *string) error

	DeleteTask(id *string) error
}
