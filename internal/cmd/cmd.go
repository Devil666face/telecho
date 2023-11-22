package cmd

import (
	"fmt"

	"github.com/Devil666face/telecho/internal/config"
	"github.com/Devil666face/telecho/internal/reader"
	"github.com/spf13/cobra"
)

type Cli struct {
	rootcmd    *cobra.Command
	config     *config.Config
	configPath string
}

func New(vers string) *Cli {
	cli := Cli{}
	cli.rootcmd = &cobra.Command{
		Use:   "telecho",
		Short: "Analog of echo but in telegram chat",
		Long: `Analog of echo but in telegram chat
Get input data from linux pipe or args
and send in your telegram chat throw telegram bot`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.entrypoint(args)
		},
		// PreRunE: func(cmd *cobra.Command, args []string) error {
		// 	var err error
		// 	if cli.config, err = config.New(""); err != nil {
		// 		return err
		// 	}
		// 	return err
		// },
		Example: `  telecho "This test message"
  telecho Send alert from telecho
  cat file.txt | telecho
  echo "$VARIABLE" | telecho`,
		Version: vers,
	}
	cli.rootcmd.PersistentFlags().StringVarP(&cli.configPath, "config", "c", "", "path to config file")
	// .telecho.env
	// telecho.yml
	// telecho.yaml
	// ~/.telecho.env
	// ~/telecho.yml
	// ~/telecho.yaml
	// ~/.config/telecho/.telecho.env
	// ~/.config/telecho/telecho.yml
	// ~/.config/telecho/telecho.yaml
	return &cli
}

func (c *Cli) entrypoint(args []string) error {
	var err error
	if c.config, err = config.New(c.configPath); err != nil {
		return err
	}
	r := reader.New(args)
	input, err := r.Read()
	if err != nil {
		return err
	}
	fmt.Println(input, c.config)
	return nil
}

func (c *Cli) Execute() error {
	return c.rootcmd.Execute()
}