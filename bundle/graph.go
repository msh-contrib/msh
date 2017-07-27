package bundle

import (
	"github.com/hzlmn/graph"
	"github.com/hzlmn/msh/utils"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

// NormalizePath normalize file path
func NormalizePath(fpath string) string {
	currDir, _ := os.Getwd()

	if filepath.IsAbs(fpath) {
		return fpath
	}

	return filepath.Join(currDir, fpath)
}

// Ctx defines walk function context of execution
type Ctx struct {
	Graph    *graph.Graph
	Reg      *regexp.Regexp
	Dir      string
	CurrFile string
	ParentID string
}

// walk recursively walks files and creates deps graph
// based on include statements
func walk(ctx *Ctx) {
	normPath := NormalizePath(ctx.CurrFile)

	if ctx.Graph.GetNode(normPath) != nil {
		// Ignore if node already presented
		return
	}

	fdata := string(utils.ReadFile(normPath))

	ctx.Graph.AddNode(normPath, map[string]interface{}{
		"data": ctx.Reg.ReplaceAllString(fdata, "# replaced import statement for $1"),
	})

	if ctx.ParentID != "" {
		ctx.Graph.AddEdge(ctx.ParentID, normPath)
	}

	//result := reqRegExp.FindAllStringSubmatch(entryData, -1)

	for _, match := range ctx.Reg.FindAllStringSubmatch(fdata, -1) {
		ctx.ParentID = normPath
		ctx.CurrFile = match[1]
		walk(ctx)
	}

}

// MakeGraph run dependency graph creating process starting from entry node
func MakeGraph(entry string) *graph.Graph {
	graph := graph.New()

	currDir, _ := os.Getwd()
	reqRegExp, regError := regexp.Compile("include\\s+'(.*?)'")

	if regError != nil {
		log.Fatal("Error while compiling RegExp", regError)
	}

	walk(&Ctx{
		Graph:    graph,
		Reg:      reqRegExp,
		Dir:      currDir,
		CurrFile: entry,
	})

	return graph
}
