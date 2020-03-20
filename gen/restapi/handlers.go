package restapi

import (
	"context"
	"fmt"
	"time"

	"gen/models"
	"gen/restapi/operations"

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-pg/pg"
	"github.com/google/uuid"
)

func DoWork(task *models.Task) {
	task.Status = models.StatusRunning
	_, err := dbconn.Model(task).WherePK().Update()
	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()
	go func(c context.Context, id uuid.UUID) {
		runinQueue <- id
	}(ctx, task.UUID)

	time.Sleep(time.Minute * 5)
	task.Status = models.StatusFinished
	_, err = dbconn.Model(task).WherePK().Update()
	if err != nil {
		fmt.Println(err)
	}
	go func(id uuid.UUID) {
		finishedQueue <- id
	}(task.UUID)
}
func CreateTaskHandler(params operations.CreateTaskParams) middleware.Responder {
	id := uuid.New()
	task := &models.Task{
		UUID:   id,
		Status: models.StatusCreated,
	}
	_, err := dbconn.Model(task).Insert()
	if err != nil {
		fmt.Println(err)
		return operations.NewCreateTaskInternalServerError()
	}
	go DoWork(task)
	return operations.NewCreateTaskAccepted().WithPayload(strfmt.UUID(task.UUID.String()))
}

func Observ(msg uuid.UUID, observable chan uuid.UUID) chan uuid.UUID {
	uuidchan := make(chan uuid.UUID, 1)
	go func() {
		for {
			for val := range observable {
				if val == msg {
					uuidchan <- val
				}
			}
		}
	}()
	return uuidchan
}
func GetFinishedTaskHandler(params operations.GetFinishedTaskParams) middleware.Responder {
	id, err := uuid.Parse(params.TaskID.String())
	if err != nil {
		return operations.NewCreateTaskAccepted()
	}
	task := &models.Task{
		UUID: id,
	}
	err = dbconn.Select(task)
	if err == pg.ErrNoRows {
		return operations.NewGetFinishedTaskNotFound()
	}
	if err != nil {
		return operations.NewGetFinishedTaskInternalServerError()
	}
	if task.Status != models.StatusFinished {
		ch := Observ(task.UUID, finishedQueue)
		for {
			select {
			case <-time.After(5 * time.Minute):
				return operations.NewGetFinishedTaskRequestTimeout()
			case <-ch:
				t, err := strfmt.ParseDateTime(task.CreatedAt.Format("2006-01-02T15:04"))
				if err != nil {
					fmt.Println(err)
				}
				return operations.NewGetFinishedTaskOK().WithPayload(&models.TaskStatus{
					Status:    task.Status,
					Timestamp: t,
				})
			}
		}
	}

	t, err := strfmt.ParseDateTime(task.CreatedAt.Format("2006-01-02T15:04"))
	if err != nil {
		fmt.Println(err)
	}

	return operations.NewGetFinishedTaskOK().WithPayload(&models.TaskStatus{
		Status:    task.Status,
		Timestamp: t,
	})
}

func GetTaskSyncHandler(params operations.GetTaskSyncParams) middleware.Responder {
	id, err := uuid.Parse(params.TaskID.String())
	if err != nil {
		fmt.Println(err)
	}
	task := &models.Task{
		UUID: id,
	}
	ch := Observ(task.UUID, runinQueue)
	select {
	case <-ch:
		t, err := strfmt.ParseDateTime(task.CreatedAt.Format("2006-01-02T15:04"))
		if err != nil {
			fmt.Println(err)
		}
		return operations.NewGetTaskSyncOK().WithPayload(&models.TaskStatus{
			Status:    task.Status,
			Timestamp: t,
		})
	default:
		close(ch)
	}
	err = dbconn.Select(task)
	if err == pg.ErrNoRows {
		return operations.NewGetTaskSyncNotFound()
	}
	if err != nil {
		return operations.NewGetTaskSyncInternalServerError()
	}
	t, err := strfmt.ParseDateTime(task.CreatedAt.Format("2006-01-02T15:04"))
	if err != nil {
		fmt.Println(err)
	}
	return operations.NewGetTaskSyncOK().WithPayload(&models.TaskStatus{
		Status:    task.Status,
		Timestamp: t,
	})
}
