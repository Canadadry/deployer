package internal

import (
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
	if len(nonFlagArgs) == 0 {
		PrintHelp(fs)
		return fmt.Errorf("You must specify a command")
	}

	command := nonFlagArgs[0]

	if command == "list" {
		PrintHelp(fs)
		return nil
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
	eval := monkey.New(map[string]object.BuiltinFunction{
		"set":  Set,
		"get":  Get,
		"task": AddTask,
	})
	err = monkey.ToError(eval.Eval(f))
	if err != nil {
		return err
	}

	_, ok := Tasks[command]
	if !ok {
		PrintHelp(fs)
		return fmt.Errorf("Invalid command : %s", command)
	}

	eval.SetEnv("getTask", &object.Builtin{Fn: GetTask})

	prog := fmt.Sprintf(`getTask("%s")()`, command)

	return monkey.ToError(eval.Eval(strings.NewReader(prog)))
}

func PrintHelp(fs *flag.FlagSet) {
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

	for _, t := range Tasks {
		fmt.Printf("\n  %s   %s", t.Title, t.Description)
	}
	fmt.Printf("\n\n")
}
