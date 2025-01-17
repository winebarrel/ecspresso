package ecspresso

import "context"

type RegisterOption struct {
	DryRun *bool `help:"dry run" default:"false"`
	Output *bool `help:"output the registered task definition as JSON" default:"false"`
}

func (opt RegisterOption) DryRunString() string {
	if *opt.DryRun {
		return dryRunStr
	}
	return ""
}

func (d *App) Register(ctx context.Context, opt RegisterOption) error {
	ctx, cancel := d.Start(ctx)
	defer cancel()

	d.Log("Starting register task definition %s", opt.DryRunString())
	td, err := d.LoadTaskDefinition(d.config.TaskDefinitionPath)
	if err != nil {
		return err
	}
	if *opt.DryRun {
		d.Log("task definition:")
		d.LogJSON(td)
		d.Log("DRY RUN OK")
		return nil
	}

	newTd, err := d.RegisterTaskDefinition(ctx, td)
	if err != nil {
		return err
	}

	if *opt.Output {
		d.LogJSON(newTd)
	}
	return nil
}
