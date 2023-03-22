package cmd

import (
	"fmt"

	todopb "api.todo/protobuf/todo/v1"
	"api.todo/todo-cli/internal/app"
	"api.todo/todo-cli/internal/config"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:  "list",
	RunE: listTask,
}

func listTask(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	app := app.AppFromContext(ctx)
	app.URL = fmt.Sprintf("%s:%s", config.Conf.URL, config.Conf.Port)

	err := app.Init(ctx)
	if err != nil {
		return err
	}
	res, err := app.Client.ListTodoTasks(ctx, &todopb.ListTodoTasksRequest{})
	if err != nil {
		return err
	}
	fmt.Printf("\n\x1b[32m%s\x1b[0m\n", "list!")
	printJsonD(res.Items)

	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
}
