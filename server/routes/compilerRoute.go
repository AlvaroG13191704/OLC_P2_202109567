package routes

import (
	"fmt"
	"server/compiler/compiler"
	"server/compiler/parser"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gofiber/fiber/v2"
)

func AnalyzeAndParseCode() fiber.Handler {
	return func(c *fiber.Ctx) error {

		// // Errors
		var TotalErrors []compiler.Error

		// get the code from the request body
		code := string(c.Body())
		fmt.Println(code)

		fmt.Println("Parsing code.............")

		input := antlr.NewInputStream(code)                                    // convert string to stream
		lexer := parser.NewGrammarLexer(input)                                 // create lexer
		tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel) // create tokens
		p := parser.NewGrammarParser(tokens)                                   // create parser

		// new list of errors
		sintacticErrors := &compiler.NewCustomErrorListener{}
		// add the errors to the parser
		p.RemoveErrorListeners()
		p.AddErrorListener(sintacticErrors)

		p.BuildParseTrees = true // tell the parser to build parse trees
		tree := p.Start_()       // parse the input

		// create the visitor
		visitor := compiler.NewVisitor()
		visitor.Visit(tree) // visit the tree

		// append the errors
		TotalErrors = append(TotalErrors, sintacticErrors.Errors...)

		// append the errors
		TotalErrors = append(TotalErrors, visitor.Errors...)

		fmt.Println("errors: ", TotalErrors)

		// if errors is null return []
		if TotalErrors == nil {
			TotalErrors = []compiler.Error{}
		}

		// get the complete code
		resultString := visitor.Generator.GetFinalCode()

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"result": resultString,
			"errors": TotalErrors,
			"symbol": visitor.TableSymbol,
		})
	}
}
