package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	todopb "api.todo/protobuf/todo/v1"
	"api.todo/todo-cli/internal/app"
	"api.todo/todo-cli/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:          "todo-cli",
	SilenceUsage: true,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	ctx := context.Background()

	var cliApp app.App
	ctx = app.ContextWithApp(ctx, &cliApp)
	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", ".todo-cli.yaml", "config file (default is .todo-cli.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find current directory.
		currentDir, err := os.Getwd()
		cobra.CheckErr(err)

		viper.AddConfigPath(currentDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".todo-cli")
	}
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	cobra.CheckErr(err)
	err = viper.Unmarshal(&config.Conf)
	cobra.CheckErr(err)
}

var Completed_value = map[string]bool{
	"on":  true,
	"off": false,
}

func readCompleted(s string, r io.Reader) (string, error) {
	var readCount = 0
	for {
		if readCount > 0 {
			fmt.Println("please type again")
		}
		readCount++
		fmt.Print(s)
		input, err := readString(r)
		if err != nil {
			return "", err
		}
		if _, ok := Completed_value[strings.ToLower(input)]; ok {
			return input, nil
		}
	}
}
func readPriority(s string, r io.Reader) (string, error) {
	var readCount = 0
	for {
		if readCount > 0 {
			fmt.Println("please type again")
		}
		readCount++
		fmt.Print(s)
		input, err := readString(r)
		if err != nil {
			return "", err
		}
		i, err := strconv.Atoi(input)
		if err != nil {
			continue
		}
		if v, ok := todopb.Priority_name[int32(i)]; ok {
			return v, nil
		}
	}
}

func readTodoTaskId(s string, r io.Reader) (string, error) {
	var readCount = 0
	for {
		if readCount > 0 {
			fmt.Println("please type again")
		}
		readCount++
		fmt.Print(s)
		input, err := readString(r)
		if err != nil {
			return "", err
		}
		if input != "" {
			return input, nil
		}
	}
}

func readContents(s string, r io.Reader) (string, error) {
	fmt.Print(s)
	input, err := readString(r)
	if err != nil {
		return "", err
	}
	return input, nil
}

func readString(r io.Reader) (string, error) {
	br := bufio.NewReader(r)
	s, err := br.ReadString(byte('\n'))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(s), nil
}

func printJsonD(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
