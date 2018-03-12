package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ademuanthony/simple/lexer"
	"github.com/ademuanthony/simple/parser"
	"github.com/ademuanthony/simple/evaluator"
	"github.com/ademuanthony/simple/object"
	"os"
	"strings"
	"errors"
)

const PROMPT = "simple> "

const MONKEY_FACEc = `
			    __,__
		    .--.  .-"	"-.  .--.
		 / ..  \/ .-.  .-.  \/ ..  \
		| |  '|  /    Y    \  |'  | |
		| \   \  \  0 | 0  /  /   / |
		 \  '- ,\.-"""""""-./, -'  /
		   ''-'  /_  ^ ^  _\  '-''
		        | \._   _./ |
			\  \ '~' /  /
			 '._'-=-'_.'
			   '-----'
	`
const SIMPLE = `
			    
            --------        ----------  ----                ----          -------           ------                  ------------------
        ---------------       ------    -------          -------       -------------        ------                  ------------------
    -----           -----     ------    --------        --------     -----------------      ------                  ------------------
    -----            ----     ------    ----------    ----------    --------    -------     ------                  ------
    -----                     ------    ----------- ------------    ------      -------     ------                  ------
     -----                    ------    ------------------------    --------   -------      ------                  ------
      -----                   ------    ------            ------    ----------------        ------                  ------
        -----                 ------    ------            ------    -------------           ------                  ------------------
          -----               ------    ------            ------    ------                  ------                  ------------------
            -----             ------    ------            ------    ------                  ------                  ------------------
              -----           ------    ------            ------    ------                  ------                  ------
                -----         ------    ------            ------    ------                  ------                  ------
                   -----      ------    ------            ------    ------                  ------                  ------
                    -----     ------    ------            ------    ------                  ------                  ------
                    -----     ------    ------            ------    ------                  ------                  ------
    ------         ------     ------    ------            ------    ------                  --------------------    ------------------
     -----------------        ------    ------            ------    ------                  --------------------    ------------------
         ---------          ----------  ------            ------    ------                  --------------------    ------------------
	`


func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "simply") {
			cmd, err := getSimplyCommand(line)
			if err != nil{
				printParserErrors(out, []string{cmd})
				return
			}
			if cmd == "run" {
				filename, err := getFilenameArg(line)
				if err != nil{
					printParserErrors(out, []string{cmd})
					return
				}
				line, err = readSourceFromFile(filename)
				if err != nil{
					printParserErrors(out, []string{cmd})
					return
				}
			}
		}

		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		/*evaluated :=*/
			evaluator.Eval(program, env)
		/*if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}*/
	}
}


func getSimplyCommand(input string) (string, error) {
	arr := strings.Split(input, " ")
	if len(arr) < 2{
		return "", errors.New("Invalid command supplied")
	}
	cmd := arr[1]

	return cmd, nil
}

func getFilenameArg(input string) (string, error) {
	arr := strings.Split(input, " ")
	if len(arr) < 3{
		return "", errors.New("No file supplied")
	}
	fileName := arr[2]

	return fileName, nil
}

func readSourceFromFile(fileName string) (content string, err error)  {
	f, err := os.Open(fileName)
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		content += scanner.Text()
	}
	return
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, SIMPLE)
	io.WriteString(out, "Woops! We ran into some troubles here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
