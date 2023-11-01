package service

type TaskService interface {
	CreateTask(name *string) (string, error)
	UpdateTask(id *string, completed *bool) error

	DeleteTask(id *string) error
}
