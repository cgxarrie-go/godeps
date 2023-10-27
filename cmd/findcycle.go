package cmd

import (
	"fmt"
	"os"

	"github.com/cgxarrie-go/godeps/pkg/app"
	"github.com/cgxarrie-go/godeps/pkg/exporter/texttree"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var findCycleCmd = &cobra.Command{
	Use:   "cycle",
	Short: "find the dependencies cycle in a package",
	Long:  `find the dependencies cycle in a package`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		p := app.NewPackage(args[0])
		err := p.LoadDependencies()
		if err != nil {
			fmt.Printf("error loading dependencies: %v\n", err)
			os.Exit(1)
		}

		p.FindLoop()

		if !p.CycleFound() {
			fmt.Printf("no cycle found")
			os.Exit(0)
		}

		e := texttree.NewExporter(p)
		b, err := e.Export()
		if err != nil {
			fmt.Printf("error exporting dependencies: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(b))

	},
}

func init() {
	rootCmd.AddCommand(findCycleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
