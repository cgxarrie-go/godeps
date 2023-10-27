package app

import (
	"fmt"
	"go/build"
	"strings"
)

type Package struct {
	name         string
	dependencies []Package
	cycleFound   bool
}

func NewPackage(name string) Package {
	return Package{
		name:         name,
		dependencies: []Package{},
		cycleFound:   false,
	}
}

func (p Package) Name() string {
	return p.name
}

func (p Package) Dependencies() []Package {
	return p.dependencies
}

func (p Package) CycleFound() bool {
	return p.cycleFound
}

func (p *Package) LoadDependencies() error {
	return p.loadDependencies(p.name, true)
}

func (p *Package) loadDependencies(ref string, isFirst bool) error {
	pkg, err := build.Import(p.name, "", build.ImportComment)
	if err != nil {
		return fmt.Errorf("importing package %s: %w", p.name, err)
	}

	if pkg.Goroot {
		return nil
	}

	if strings.Contains(pkg.Dir, "/vendor/") {
		p.name = fmt.Sprintf("%s (vendor)", p.name)
		return nil
	}

	if p.name == ref && !isFirst {
		p.cycleFound = true
		return nil
	}

	for _, imp := range pkg.Imports {
		dep := NewPackage(imp)
		err = dep.loadDependencies(ref, false)
		if err != nil {
			return err
		}
		if dep.cycleFound {
			p.cycleFound = true
		}
		p.dependencies = append(p.dependencies, dep)
	}
	return nil
}

func (p *Package) FindLoop() {
	for i := 0; i < len(p.dependencies); i++ {
		if !p.dependencies[i].cycleFound {
			p.dependencies = append(p.dependencies[:i], p.dependencies[i+1:]...)
			i--
		} else {
			p.dependencies[i].FindLoop()
		}
	}
}
