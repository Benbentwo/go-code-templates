package {{ .Folder | toLower  }}

import (
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/spf13/cobra"
)

// options for the command
type {{.CommandUse | title }}Options struct {
	*	opts.CommonOptions
}


func NewCmd{{ .CommandUse | title }}(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &{{.CommandUse | title }}Options {
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "{{ .CommandUse | toLower }}",
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			helper.CheckErr(err)
		},
	}
	options.Add{{.CommandUse | title }}Flags(cmd)
	// the line below (Section to...) is for the generate-function command to add a template_command to.
	// the line above this and below can be deleted.
	// DO NOT DELETE THE FOLLOWING LINE:
	// Section to add commands to:

	return cmd
}

// Run implements this command
func (o *{{.CommandUse | title }}Options) Run() error {
	return o.Cmd.Help()
}


func (o *{{.CommandUse | title }}Options) Add{{.CommandUse | title }}Flags(cmd *cobra.Command) {
	o.Cmd = cmd
}