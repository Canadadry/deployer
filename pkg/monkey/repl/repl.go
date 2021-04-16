package repl

import (
	"app/pkg/monkey/evaluator"
	"app/pkg/monkey/lexer"
	"app/pkg/monkey/object"
	"app/pkg/monkey/parser"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

const PROMPT = ">> "

func Start(in io.Reader) error {
	content, err := ioutil.ReadAll(in)
	if err != nil {
		return fmt.Errorf("Can't read in put %w\n", err)
	}

	l := lexer.New(string(content))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		msg := strings.Join(p.Errors(), "\n")
		return fmt.Errorf("Parse errors : \n%s", msg)
	}

	env := object.NewEnvironment()
	result := evaluator.Eval(program, env)
	if result != nil {
		if result.Type() == object.ERROR_OBJ {
			return fmt.Errorf("Runtime error : %w", result.Inspect())
		}
	}
	return nil
}

func StartInterActive(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
