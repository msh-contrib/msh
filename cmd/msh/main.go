package main

import (
	//"bytes"
	"flag"
	"fmt"
	"github.com/hzlmn/msh/graph"
	//"github.com/mvdan/sh/interp"
	"github.com/mvdan/sh/syntax"
	//"bufio"
	"github.com/davecgh/go-spew/spew"
	"github.com/hzlmn/msh/resolver"
	"log"
	//"os"
)

var source = flag.String("i", "", "define entry point file")

func composeFiles(list ...*syntax.File) *syntax.File {
	newFile := &syntax.File{}

	for _, f := range list {
		for _, stmt := range f.Stmts {
			newFile.Stmts = append(newFile.Stmts, stmt)
		}
	}

	return newFile
}

func main() {
	flag.Parse()

	if len(*source) < 1 {
		log.Fatal("-i input flag should be defined")
	}

	depsGraph := graph.MakeGraph(*source)

	resolvedGraph := resolver.Resolve(depsGraph.GetNode(graph.NormalizePath(*source)))
	fmt.Println("resolved", spew.Sdump(resolvedGraph))

	for _, node := range resolvedGraph {
		fmt.Println(node.GetParams()["data"])
	}

	// fmt.Println("works", *source)

	// fileData := utils.ReadFile(*source)
	// sourceReader := bytes.NewReader(fileData)

	// mainFile := utils.ReadFile("./test2.sh")
	// mainReader := bytes.NewReader(mainFile)

	// file, fileError := syntax.Parse(sourceReader, "shell", syntax.PosixConformant)
	// mainParsed, mainFileError := syntax.Parse(mainReader, "shell", syntax.PosixConformant)

	// // fmt.Println("mainFile", string(mainFile))
	// // fmt.Println("module", string(fileData))

	// if mainFileError != nil {
	// 	log.Fatal("Main read error")
	// }

	// if fileError != nil {
	// 	log.Fatal("Error while parsing source file")
	// }

	// syntax.Walk(mainParsed, func(node syntax.Node) bool {
	// 	switch x := node.(type) {
	// 	case *syntax.CallExpr:
	// 		fmt.Println("call exp", spew.Sdump(x))
	// 		for _, value := range x.Args {
	// 			fmt.Println("arg", spew.Sdump(value))
	// 		}
	// 		break
	// 	}

	// 	return true
	// })

	// newFile := composeFiles(mainParsed, file)

	// syntax.Fprint(os.Stdout, newFile)

	// scriptRunner := &interp.Runner{
	// 	File:   file,
	// 	Stdout: os.Stdout,
	// 	Stderr: os.Stderr,
	// }

	// runnerError := scriptRunner.Run()

	// if runnerError != nil {
	// 	fmt.Println("Script running error", runnerError)
	// }
	//syntax.Fprint(os.Stdout, parsedFile)

	//graph.CollectNodes(*source)
	// syntax.Walk(parsedFile, func(node syntax.Node) bool {
	// 	switch x := node.(type) {
	// 	case *syntax.CallExpr:
	// 		fmt.Println("call expr", x.Args[1].Parts[0])
	// 		break
	// 	case *syntax.ParamExp:
	// 		fmt.Println("parsed param expression", x.Exp.Word)
	// 		break
	// 	case *syntax.FuncDecl:
	// 		fmt.Println("found func declaration", spew.Sdump(x))
	// 		break
	// 	}

	// 	return true
	// })

	// for _, val := range parsedFile.Stmts {
	// 	fmt.Println("statemtnt", val.Cmd)
	// }

	// node.AddEdge(childNode)

	// graph := graph.New()
	// graph.AttachNode(node)

	// fmt.Println(graph.GetNodes())
}
