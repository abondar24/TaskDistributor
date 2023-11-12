package service

import (
	"github.com/abondar24/TaskDistributor/taskData/data"
	"github.com/abondar24/TaskDistributor/taskStore/dao"
	"github.com/abondar24/TaskDistributor/taskStore/model"
)

type TaskDbService struct {
	taskDAO        dao.TaskDao
	taskHistoryDAO dao.TaskHistoryDao
}

type TaskService interface {
	SaveUpdateTask(task *data.Task) error

	GetTaskById(id *string) (*data.Task, error)
	GetTasksByStatus(status *data.TaskStatus, offset *int, limit *int) ([]*data.Task, error)
	GetTaskHistory(id *string) (*data.TaskHistory, error)
}

func NewTaskService(taskDao dao.TaskDao, historyDao dao.TaskHistoryDao) *TaskDbService {
	return &TaskDbService{
		taskDAO:        taskDao,
		taskHistoryDAO: historyDao,
	}
}

func (ts *TaskDbService) SaveUpdateTask(task *data.Task) error {

	if task.Status == data.TASK_CREATED {
		taskDTO := &model.TaskDTO{
			Id:        task.ID,
			Name:      task.Name,
			CreatedAt: task.CreateTime,
		}
		err := ts.taskDAO.SaveTask(taskDTO)
		if err != nil {
			return err
		}
	}

	taskHistoryEntry := &model.TaskHistoryDTO{
		TaskId:    task.ID,
		Status:    task.Status,
		UpdatedAt: task.UpdateTime,
	}

	err := ts.taskHistoryDAO.SaveTaskEntry(taskHistoryEntry)
	if err != nil {
		return err
	}

	return nil
}

func (ts *TaskDbService) GetTaskById(id *string) (*data.Task, error) {
	task, err := ts.taskDAO.GetTaskById(id)
	if err != nil {
		return nil, err
	}

	taskEntry, err := ts.taskHistoryDAO.GetTaskById(id)
	if err != nil {
		return nil, err
	}

	result := &data.Task{
		ID:         task.Id,
		Name:       task.Name,
		Status:     taskEntry.Status,
		CreateTime: task.CreatedAt,
		UpdateTime: taskEntry.UpdatedAt,
	}

	return result, nil
}

func (ts *TaskDbService) GetTasksByStatus(status *data.TaskStatus, offset *int, limit *int) ([]*data.Task, error) {
	result := make([]*data.Task, 0)
	statusEntries, err := ts.taskHistoryDAO.GetTasksByStatus(status, offset, limit)
	if err != nil {
		return nil, err
	}

	ids := make([]*string, 0)

	for _, se := range *statusEntries {
		ids = append(ids, &se.TaskId)
	}

	tasks, err := ts.taskDAO.GetTasksByIds(ids)
	if err != nil {
		return nil, err
	}

	for _, t := range *tasks {
		task := &data.Task{
			ID:         t.Id,
			Name:       t.Name,
			CreateTime: t.CreatedAt,
		}
		result = append(result, task)
	}

	for i, se := range *statusEntries {
		result[i].Status = se.Status
		result[i].UpdateTime = se.UpdatedAt
	}

	return result, nil
}

func (ts *TaskDbService) GetTaskHistory(id *string) (*data.TaskHistory, error) {
	task, err := ts.taskDAO.GetTaskById(id)
	if err != nil {
		return nil, err
	}

	taskHistory, err := ts.taskHistoryDAO.GetTaskHistory(id)
	if err != nil {
		return nil, err
	}

	statusHistory := make([]data.TaskStatusEntry, 0)

	for _, th := range *taskHistory {
		statusEntry := data.TaskStatusEntry{
			Status:    th.Status,
			UpdatedAt: th.UpdatedAt,
		}
		statusHistory = append(statusHistory, statusEntry)
	}

	result := &data.TaskHistory{
		ID:            task.Id,
		Name:          task.Name,
		CreateTime:    task.CreatedAt,
		StatusHistory: statusHistory,
	}

	return result, nil
}
