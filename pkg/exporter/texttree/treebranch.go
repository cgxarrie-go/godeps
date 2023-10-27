package texttree

type treeBranch int

const (
	bar treeBranch = iota
	item
	lastItem
)

var treeBranchNames = map[treeBranch]string{
	bar:      "│ ",
	item:     "├ ",
	lastItem: "└ ",
}

// Name return the name of the treeBranch
func (t treeBranch) Name() string {
	return treeBranchNames[t]
}
