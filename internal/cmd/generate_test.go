package cmd

import (
	"context"
	"testing"

	"github.com/spf13/cobra"
)

func Test_Generate(t *testing.T)  {
	dir := "/home/divar/Tutorial/sqlc/sqlc/examples/authors"
	rootCmd := &cobra.Command{}
	stderr := rootCmd.ErrOrStderr()
	_, err := Generate(context.Background(), Env{}, dir, "", stderr)
	if err != nil {
		t.Fatalf("can not generate: %v \n", err.Error())
	}
}
