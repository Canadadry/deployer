package internal

import (
	"app/internal/env"
	"app/internal/runner"
	"app/pkg/monkey"
	"app/pkg/monkey/object"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Run(args []string) error {
	fs := flag.NewFlagSet(args[0], flag.ExitOnError)
	file := fs.String("f", "deploy.mky", "Specify deployer file")
	_ = fs.Bool("dry", false, "dry run : only print action")

	fs.Parse(args[1:])

	nonFlagArgs := fs.Args()
	command := "list"
	if len(nonFlagArgs) > 0 {
		command = nonFlagArgs[0]
	}

	if command == "init" {
		return ioutil.WriteFile(*file, []byte(`
task("test","mydesc",fn(){
    print("test")
})
`), 0644)
	}

	f, err := os.Open(*file)
	if err != nil {
		return err
	}
	e := env.New()
	eval := monkey.New(map[string]object.BuiltinFunction{
		"set":  e.Set,
		"get":  e.Get,
		"task": e.AddTask,
	})
	err = object.ToError(eval.Eval(f))
	if err != nil {
		return err
	}

	if command == "list" {
		PrintHelp(fs, e.Tasks)
		return nil
	}
	_, ok := e.Tasks[command]
	if !ok {
		PrintHelp(fs, e.Tasks)
		return fmt.Errorf("Invalid command : %s", command)
	}

	eval.SetEnv("getTask", &object.Builtin{Fn: e.GetTask})
	r := runner.NewDryRun(os.Stdout)
	eval.SetEnv("run", &object.Builtin{Fn: env.Run(r)})
	eval.SetEnv("runLocally", &object.Builtin{Fn: env.RunLocally(r)})
	eval.SetEnv("upload", &object.Builtin{Fn: env.Upload(r)})
	eval.SetEnv("download", &object.Builtin{Fn: env.Download(r)})

	prog := fmt.Sprintf(`getTask("%s")()`, command)

	return object.ToError(eval.Eval(strings.NewReader(prog)))
}

func PrintHelp(fs *flag.FlagSet, tasks map[string]env.Task) {
	fmt.Println(fs.Name())
	fmt.Print(`
Usage:
Command [option]

Options:

`)
	fs.PrintDefaults()
	fmt.Print(`
Available commands:

  init   Initialize deployer in your project
  list   Lists commands`)

	for _, t := range tasks {
		fmt.Printf("\n  %s   %s", t.Title, t.Description)
	}
	fmt.Printf("\n\n")
}
