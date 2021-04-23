package internal

import (
	"app/internal/env"
	"app/internal/runner"
	"app/pkg/ansi"
	"app/pkg/fs"
	"app/pkg/monkey"
	"app/pkg/monkey/object"
	"app/pkg/ssh"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Run(args []string) error {
	flag := flag.NewFlagSet(args[0], flag.ExitOnError)
	file := flag.String("f", "deploy.mky", "Specify deployer file")
	dryrun := flag.Bool("dry", false, "dry run : only print action")

	flag.Parse(args[1:])

	nonFlagArgs := flag.Args()
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
		PrintHelp(flag, e.Tasks)
		return nil
	}
	_, ok := e.Tasks[command]
	if !ok {
		PrintHelp(flag, e.Tasks)
		return fmt.Errorf("Invalid command : %s", command)
	}

	eval.SetEnv("getTask", &object.Builtin{Fn: e.GetTask})
	var r runner.Runner
	if *dryrun {
		r = runner.NewDryRun(os.Stdout)
	} else {
		run := runner.New()
		host_path, ok := e.Store["host_path"]
		if !ok {
			host_path = "~/"
		}
		host_addr, ok := e.Store["host_addr"]
		if !ok {
			return fmt.Errorf("you must set the key 'host_addr'")
		}
		host_port, ok := e.Store["host_port"]
		if !ok {
			host_port = "22"
		}

		host_user, ok := e.Store["host_user"]
		if !ok {
			return fmt.Errorf("you must set the key 'host_user'")
		}
		host_private_key_name, ok := e.Store["host_private_key"]
		if !ok {
			host_private_key_name = "~/.ssh/id_rsa"
		}
		host_private_key, err := ioutil.ReadFile(host_private_key_name)
		if err != nil {
			return err
		}
		run.Ssh, err = ssh.New(ssh.Login{
			Addr:       host_addr,
			Port:       host_port,
			User:       host_user,
			PrivateKey: string(host_private_key),
		})
		if err != nil {
			return err
		}
		run.Local = fs.NewLocal("./")
		run.DistantPath = host_path
		r = run
	}
	eval.SetEnv("run", &object.Builtin{Fn: env.Run(r)})
	eval.SetEnv("runLocally", &object.Builtin{Fn: env.RunLocally(r)})
	eval.SetEnv("upload", &object.Builtin{Fn: env.Upload(r)})
	eval.SetEnv("download", &object.Builtin{Fn: env.Download(r)})

	prog := fmt.Sprintf(`getTask("%s")()`, command)

	return object.ToError(eval.Eval(strings.NewReader(prog)))
}

func PrintHelp(fs *flag.FlagSet, tasks map[string]env.Task) {
	fmt.Println(fs.Name())
	fmt.Println()
	fmt.Println(ansi.Yellow("Usage:"))
	fmt.Println()
	fmt.Println("   deployer [option] Command")
	fmt.Println()
	fmt.Println(ansi.Yellow("Options:"))
	fmt.Println()
	fs.PrintDefaults()
	fmt.Println()
	fmt.Println(ansi.Yellow("Available commands"))
	fmt.Println()
	fmt.Println(ansi.Green("   init   ") + "Initialize deployer in your project")
	fmt.Println(ansi.Green("   list   ") + "Lists commands")

	for _, t := range tasks {
		fmt.Printf(ansi.Green("   %s")+"   %s\n", t.Title, t.Description)
	}
	fmt.Printf("\n\n")
}
