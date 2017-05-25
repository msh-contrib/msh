package main

import (
	"flag"
	"fmt"
	"github.com/hzlmn/msh/graph"
	"github.com/mvdan/sh/syntax"
	//"bufio"
	//"bufio"
	"bytes"
	"io/ioutil"
	"log"
)

var source = flag.String("i", "", "define entry point file")

func readFile(file string) (data []byte) {
	data, readError := ioutil.ReadFile(file)
	if readError != nil {
		log.Fatal(readError)
	}

	return
}

func main() {
	flag.Parse()

	if len(*source) < 1 {
		log.Fatal("-i input flag should be defined")
	}

	fmt.Println("works", *source)

	fileData := readFile(*source)

	//output := flag.String("o", "out.sh", "output file")
	//fileData, parseError := syntax.Parse()
	params := map[string]interface{}{
		"param": "test",
	}

	fileReader := bytes.NewReader(fileData)
	parsedFile, parseError := syntax.Parse(fileReader, "shell", syntax.PosixConformant)
	if parseError != nil {
		log.Fatal(parseError)
	}

	fmt.Println(parsedFile.Stmts[0])

	//syntax.Fprint(os.Stdout, parsedFile)

	childNode := graph.NewNode("child", params)
	node := graph.NewNode("key", params)

	syntax.Walk(parsedFile, func(node syntax.Node) bool {
		switch x := node.(type) {
		case *syntax.CallExpr:
			fmt.Println("call expr", x.Args[1].Parts[0])
			break
		case *syntax.ParamExp:
			fmt.Println("parsed param expression", x.Exp.Word)
			break
		case *syntax.FuncDecl:
			fmt.Println("found func declaration", x.Name)
			break
		}

		return true
	})

	// for _, val := range parsedFile.Stmts {
	// 	fmt.Println("statemtnt", val.Cmd)
	// }

	node.AddEdge(childNode)

	graph := graph.New()
	graph.AttachNode(node)

	fmt.Println(graph.GetNodes())
}
