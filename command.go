package cobra

import (
	"context"
	"errors"

	"github.com/spf13/cobra"
)

type CLI struct {
	rootCmd *cobra.Command
}

func New(rootCmd *cobra.Command) *CLI {
	return &CLI{rootCmd: rootCmd}
}

func (c *CLI) String() string {
	return "Cobra CLI App"
}

func (c *CLI) Invoke(ctx context.Context) error {
	if c.rootCmd == nil {
		return errors.New("root command is nil")
	}
	return c.rootCmd.ExecuteContext(ctx)
}
