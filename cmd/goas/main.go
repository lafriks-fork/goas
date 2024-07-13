package main

import (
	"fmt"
	"os"

	"github.com/lafriks-fork/goas"

	"github.com/spf13/cobra"
)

var version = "v1.0.0"

var rootCmd = &cobra.Command{
	Use:     "goas <module-path> <main-file-path>",
	Short:   "goas is OpenAPI Specification generator for Go",
	Long:    "goas is a tool to generate OpenAPI Specification file from Go source code.",
	Version: version,

	RunE: action,
}

var (
	handlerPath string
	output      string
	debug       bool
)

func action(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("<module-path> and <main-file-path> are required")
	}

	var (
		modulePath   = args[0]
		mainFilePath = args[1]
	)

	p, err := goas.NewParser(modulePath, mainFilePath, handlerPath, debug)
	if err != nil {
		return err
	}

	if output == "-" {
		return p.CreateOAS(os.Stdout)
	}

	return p.CreateOASFile(output)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&handlerPath, "handler-path", "", "goas only search handleFunc comments under the path")
	rootCmd.Flags().StringVar(&output, "output", "-", "output file")
	rootCmd.Flags().BoolVar(&debug, "debug", false, "show debug message")
}
