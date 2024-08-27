package main

import (
	"time"

	"github.com/urfave/cli/v2"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TaskModel interface {
	Add(t Task) (bool, error)
	Update(id int) (bool, error)
	Delete(id int) (bool, error)
	MarkInProgress(id int) (bool, error)
	MarkDone(id int) (bool, error)
}

type application struct {
	TaskList []Task
}

func (app *application) loadTasks() error {
	return nil
}

func (app *application) saveTasks() error {
	return nil
}

func (app *application) Add(ctx *cli.Context) error {
	return nil
}

func (app *application) List(ctx *cli.Context) error {
	return nil
}

func (app *application) Delete(ctx *cli.Context) error {
	return nil
}
