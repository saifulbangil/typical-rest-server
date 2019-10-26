package buildtool

import (
	"io"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/kelseyhightower/envconfig"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typicmd/buildtool/markdown"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typictx"
)

const (
	configTemplate = `| Key | Type | Default | Required | Description |	
|---|---|---|---|---|{{range .}}
|{{usage_key .}}|{{usage_type .}}|{{usage_default .}}|{{usage_required .}}|{{usage_description .}}|{{end}}`
)

// Readme detail
type Readme struct {
	*typictx.Context
	Title       string
	Description string
}

// Markdown to return the markdown
func (r Readme) Markdown(w io.Writer) *markdown.Markdown {
	md := &markdown.Markdown{Writer: w}
	md.Comment("Autogenerated by Typical-Go. DO NOT EDIT.")
	md.Heading1(r.Title)
	md.Writeln(r.Description)
	r.prerequisite(md)
	r.runInstruction(md)
	r.configuration(md)
	r.releaseDistribution(md)
	return md
}

func (r Readme) prerequisite(md *markdown.Markdown) {
	md.Heading2("Prerequisite")
	md.OrderedList(
		"Install [Go](https://golang.org/doc/install) or `brew install go`",
	)
}

func (r Readme) runInstruction(md *markdown.Markdown) {
	md.Heading2("Run")
	md.Writeln("Use `./typicalw run` to compile and run local development. [Learn more](https://typical-go.github.io/learn-more/wrapper.html)")
}

func (r Readme) releaseDistribution(md *markdown.Markdown) (err error) {
	md.Heading2("Release Distribution")
	md.Writeln("Use `./typicalw release` to make the release. You can find the binary at `release` folder. [Learn more](https://typical-go.github.io/learn-more/release.html)")
	return
}

func (r Readme) configuration(md *markdown.Markdown) {
	md.Heading2("Configuration")
	for _, module := range r.Modules {
		if module.Name != "" {
			md.Heading3(strcase.ToCamel(module.Name))
		}
		var builder strings.Builder
		cfg := module.Config
		envconfig.Usagef(cfg.Prefix(), cfg.Spec(), &builder, configTemplate)
		md.Writeln(builder.String())
	}
}
