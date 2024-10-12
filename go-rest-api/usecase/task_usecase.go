package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type ITaskUsecase interface {
	GetAll(userID uint) ([]model.TaskResponse, error)
	GetByID(userID uint, taskID uint) (model.TaskResponse, error)
	Create(task model.Task) (model.TaskResponse, error)
	Update(task model.Task, userID uint, taskID uint) (model.TaskResponse, error)
	Delete(userID uint, taskID uint) error
}

type taskUsecase struct {
	tr repository.ITaskRepository
	tv validator.ITaskValidator
}

func NewTaskUsecase(tr repository.ITaskRepository, tv validator.ITaskValidator) ITaskUsecase {
	return &taskUsecase{tr, tv}
}

func (tu *taskUsecase) GetAll(userID uint) ([]model.TaskResponse, error) {
	tasks := []model.Task{}
	if err := tu.tr.GetAll(&tasks, userID); err != nil {
		return nil, err
	}

	resTasks := []model.TaskResponse{}
	for _, t := range tasks {
		r := model.TaskResponse{
			ID:        t.ID,
			Title:     t.Title,
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
		}
		resTasks = append(resTasks, r)
	}
	return resTasks, nil
}

func (tu *taskUsecase) GetByID(userID uint, taskID uint) (model.TaskResponse, error) {
	task := model.Task{}
	if err := tu.tr.GetByID(&task, userID, taskID); err != nil {
		return model.TaskResponse{}, err
	}

	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) Create(task model.Task) (model.TaskResponse, error) {
	if err := tu.tv.Validate(task); err != nil {
		return model.TaskResponse{}, err
	}

	if err := tu.tr.Create(&task); err != nil {
		return model.TaskResponse{}, err
	}

	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) Update(task model.Task, userID uint, taskID uint) (model.TaskResponse, error) {
	if err := tu.tv.Validate(task); err != nil {
		return model.TaskResponse{}, err
	}

	if err := tu.tr.Update(&task, userID, taskID); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) Delete(userID uint, taskID uint) error {
	if err := tu.tr.Delete(userID, taskID); err != nil {
		return err
	}
	return nil
}
