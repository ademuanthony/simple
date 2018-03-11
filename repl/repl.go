package repl

import (
	"bufio"
	"fmt"
	"io"

	"bitbucket.org/laverita/enyo/lexer"
	"bitbucket.org/laverita/enyo/parser"
	"bitbucket.org/laverita/enyo/evaluator"
	"bitbucket.org/laverita/enyo/object"
)

const PROMPT = ">> "

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


func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, SIMPLE)
	io.WriteString(out, "Woops! We ran into some troubles here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
