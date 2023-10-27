package texttree

import (
	"fmt"
	"strings"

	"github.com/cgxarrie-go/godeps/pkg/app"
)

type exporter struct {
	pkg app.Package
}

// Export implements app.Exporter.
func (e exporter) Export() ([]byte, error) {
	result := []string{}

	result = append(result, e.pkg.Name())
	result = append(result, e.getTree(e.pkg, 1)...)

	return []byte(strings.Join(result, "\n")), nil
}

func (e exporter) getTree(p app.Package, level int) []string {
	result := []string{}

	numbDeps := len(p.Dependencies())
	currentDep := 0
	for _, dep := range p.Dependencies() {
		currentDep++
		prefix := e.getItemPrefix(level, currentDep == numbDeps)
		item := fmt.Sprintf("%s%s", prefix, dep.Name())
		result = append(result, item)
		result = append(result,
			e.getTree(dep, level+1)...)
	}

	return result
}

func NewExporter(p app.Package) app.Exporter {
	return exporter{
		pkg: p,
	}
}

func (e exporter) getItemPrefix(level int, isLast bool) string {
	prefix := strings.Repeat(bar.Name(), level-1)
	if isLast {
		prefix += lastItem.Name()
	} else {
		prefix += item.Name()
	}
	return prefix
}
