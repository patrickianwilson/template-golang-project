package cmd

import (
	"errors"
	"flag"
	"fmt"
	"github.com/patrickianwilson/cli/common"
)

type SampleCmd struct {
	flagSet *flag.FlagSet
	context *common.Context
	name    string
}

func (*SampleCmd) GetCmdName() string {
	return "sample-cmd"
}

func (*SampleCmd) GetCmdGroup() string {
	return "sample-group"
}

func (cmd *SampleCmd) Init(args []string, context *common.Context) error {
	cmd.context = context
	cmd.flagSet = flag.NewFlagSet("module-checkout", flag.ContinueOnError)

	cmd.flagSet.StringVar(&cmd.name, "name", "", "The name of the thing")
	cmd.flagSet.Parse(args)
	return nil
}

func (cmd *SampleCmd) Validate() error {
	if cmd.name == "" {
		return errors.New("name is required")
	}
	return nil
}

func (cmd *SampleCmd) PrintHelp() {
	fmt.Println("This command does something")
	cmd.flagSet.PrintDefaults()
}

func (*SampleCmd) DoAction() error {
	panic("implement me")
}

func NewSampleCmd() *SampleCmd {
	return &SampleCmd{}
}
