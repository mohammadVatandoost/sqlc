package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
)

func Test_Generate(t *testing.T)  {
	dir := "/home/divar/Tutorial/sqlc/sqlc/examples/authors"
	rootCmd := &cobra.Command{}
	stderr := rootCmd.ErrOrStderr()
	output, err := Generate(context.Background(), Env{}, dir, "", stderr)
	if err != nil {
		t.Fatalf("can not generate: %v \n", err.Error())
	}

	for filename, source := range output {
		os.MkdirAll(filepath.Dir(filename), 0755)
		if err := os.WriteFile(filename, []byte(source), 0644); err != nil {
			fmt.Fprintf(stderr, "%s: %s\n", filename, err)
			os.Exit(1)
		}
	}
}
