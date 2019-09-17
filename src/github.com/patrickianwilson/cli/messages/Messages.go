package messages

const GENERIC_USAGE_MESSAGE string = `Incorrect Usage!  Expected command format 'cli <command> [options]' where:

command is one of {workspace, module, template}

here are some examples:

1. Create a project workspace (a bundle of modules making up a single releasable unit)
	project workspace create

2. Create a module in the workspace
	project module create --name SampleModule --tempate java --flavor lib

3. List all available templates
	project template list


`
