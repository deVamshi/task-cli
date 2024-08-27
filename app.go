package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

type Task struct {
	Title     string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type application struct {
	TaskList []Task
}

func (app *application) loadTasks() error {
	fileData, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = json.Unmarshal(fileData, &app.TaskList)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (app *application) saveTasks() error {

	fileData, err := json.Marshal(app.TaskList)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(fileData) == 0 {
		return errors.New("no tasks")
	}

	err = os.WriteFile("data.json", fileData, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (app *application) Add(ctx *cli.Context) error {

	if !ctx.Args().Present() {
		return errors.New("error: task cannot be empty")
	}

	tasks := ctx.Args().Slice()

	for _, t := range tasks {
		newTask := Task{
			Title:     t,
			Status:    "TODO",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		app.TaskList = append(app.TaskList, newTask)
	}

	fmt.Printf("Added %d task(s)\n", ctx.Args().Len())

	return nil
}

func (app *application) List(ctx *cli.Context) error {

	fmt.Println("ID\tTitle\tStatus\t")

	for i, t := range app.TaskList {
		fmt.Printf("%d\t%s\t%s\n", i+1, t.Title, t.Status)
	}

	return nil
}

func (app *application) Delete(ctx *cli.Context) error {
	return nil
}

type TaskModel interface {
	Add(t Task) (bool, error)
	Update(id int) (bool, error)
	Delete(id int) (bool, error)
	MarkInProgress(id int) (bool, error)
	MarkDone(id int) (bool, error)
}
