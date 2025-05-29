package internal

type Command struct {
	// name of the command (e.g. curl)
	Name string `json:"command"`
	// arguments/flags that should be passed to the command
	// such as "-X POST" or "--request GET"
	Args []string `json:"args"`
}
