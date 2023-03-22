package cmd

import (
	"fmt"
	"os"

	todopb "api.todo/protobuf/todo/v1"
	"api.todo/todo-cli/internal/app"
	"api.todo/todo-cli/internal/config"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:  "create",
	RunE: createTask,
}

func createTask(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	app := app.AppFromContext(ctx)
	app.URL = fmt.Sprintf("%s:%s", config.Conf.URL, config.Conf.Port)

	contents, err := readContents("本文\n> ", os.Stdin)
	if err != nil {
		return err
	}
	priorityValue, err := readPriority("優先度（0:未指定, 1:高, 2:中, 3:低）\n> ", os.Stdin)
	if err != nil {
		return err
	}
	priorityEnum, ok := todopb.Priority_value[priorityValue]
	if !ok {
		priorityEnum = int32(todopb.Priority_PRIORITY_UNSPECIFIED)
	}
	err = app.Init(ctx)
	if err != nil {
		return err
	}
	res, err := app.Client.CreateTodoTask(ctx, &todopb.CreateTodoTaskRequest{
		TodoTask: &todopb.TodoTask{
			Contents: contents,
			Priority: todopb.Priority(priorityEnum),
		},
	})
	if err != nil {
		return err
	}
	fmt.Printf("\n\x1b[32m%s\x1b[0m\n", "created!")
	printJsonD(res.TodoTask)

	return nil
}

func init() {
	rootCmd.AddCommand(createCmd)
}
