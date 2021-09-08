package arguments

import "github.com/urfave/cli/v2"

const (
	Platform = "platform"
	Target   = "target"
	Filename = "filename"
	Locale   = "locale"
	Document = "document"
	Sheet    = "sheet"
)

type Arguments struct {
	Platform string
	Target   string
	Filename string
	Locale   string
	Document string
	Sheet    string
}

func Get(c *cli.Context) *Arguments {
	return &Arguments{
		Platform: c.String(Platform),
		Target:   c.String(Target),
		Filename: c.String(Filename),
		Locale:   c.String(Locale),
		Document: c.String(Document),
		Sheet:    c.String(Sheet),
	}
}
