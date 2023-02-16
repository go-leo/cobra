package cobra

import (
	"context"
	"errors"

	"github.com/spf13/cobra"
)

type Cobra struct {
	rootCmd *cobra.Command
}

func New(rootCmd *cobra.Command) *Cobra {
	return &Cobra{rootCmd: rootCmd}
}

func (c *Cobra) String() string {
	return "Cobra App"
}

func (c *Cobra) Invoke(ctx context.Context) error {
	if c.rootCmd == nil {
		return errors.New("root command is nil")
	}
	return c.rootCmd.ExecuteContext(ctx)
}
