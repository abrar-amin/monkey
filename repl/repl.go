package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ": "

func Start(in io.Reader, out io.Writer){
	scanner := bufio.NewScanner(in)

	//Infinite loop shorthand 
	for 
	{
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		// No more input (EOF)
		if !scanned {
			return 
		}

		line := scanner.Text()
		// Create a new lexer for each line ...
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken(){
			fmt.Printf("%+v\n", tok)
		}
	}
}