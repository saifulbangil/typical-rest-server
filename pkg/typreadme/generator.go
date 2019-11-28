package typreadme

import (
	"fmt"
	"io"

	"github.com/iancoleman/strcase"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typctx"

	"github.com/typical-go/typical-go/pkg/typmodule"
	"github.com/typical-go/typical-rest-server/pkg/typreadme/markdown"
)

// Generator responsible to generate readme
type Generator struct{}

// Generate readme
func (Generator) Generate(ctx *typctx.Context, w io.Writer) (err error) {
	md := &markdown.Markdown{Writer: w}
	md.Comment("Autogenerated by Typical-Go. DO NOT EDIT.")
	md.H1(ctx.Name)
	if ctx.Description != "" {
		md.Writeln(ctx.Description)
	}
	md.H3("Usage")
	md.H3("Configuration")
	configuration(md, ctx)
	md.Hr()
	md.H2("Development Guide")
	md.H3("Prerequisite")
	prerequisite(md)
	md.H3("Build & Run")
	buildAndRun(md)
	md.H3("Test")
	test(md)
	md.H3("Release the destribution")
	releaseDistribution(md)
	md.H3("Command")
	return
}

func prerequisite(md *markdown.Markdown) {
	md.Writeln("Install [Go](https://golang.org/doc/install) (It is recommend to install via [Homebrew](https://brew.sh/) `brew install go`)")
}

func buildAndRun(md *markdown.Markdown) {
	md.Writeln("Use `./typicalw run` to build and run the project.")
}

func test(md *markdown.Markdown) {
	md.Writeln("Use `./typicalw test` to test the project.")
}

func releaseDistribution(md *markdown.Markdown) {
	md.Writeln("Use `./typicalw release` to make the release. You can find the binary at `release` folder.")
	md.Writeln("Learn more [Release Distribution](https://typical-go.github.io/learn-more/build-tool/release-distribution.html)")
}

func module(md *markdown.Markdown, module interface{}) {
	if name := typmodule.Name(module); name != "" {
		md.H3(strcase.ToCamel(name))
	}
	if description := typmodule.Description(module); description != "" {
		md.Writeln(description)
	}

	cmd := typbuildtool.Command(nil, module)
	if cmd != nil && len(cmd.Subcommands) > 0 {
		md.WriteString("Commands:\n")
		var cmdHelps []string
		for _, subcmd := range cmd.Subcommands {
			cmdHelps = append(cmdHelps,
				fmt.Sprintf("`./typicalw %s %s`: %s", cmd.Name, subcmd.Name, subcmd.Usage))
		}
		md.UnorderedList(cmdHelps...)
	}
}

func configuration(md *markdown.Markdown, ctx *typctx.Context) {
	md.WriteString("| Name | Type | Default | Required |\n")
	md.WriteString("|---|---|---|:---:|\n")
	// TODO: sort by name
	for _, module := range ctx.AllModule() {
		if configurer, ok := module.(typcfg.Configurer); ok {
			for _, field := range configurer.Configure().Fields() {
				var required string
				if field.Required {
					required = "Yes"
				}
				md.WriteString(fmt.Sprintf("|%s|%s|%s|%s|\n",
					field.Name, field.Type, field.Default, required))
			}
		}
	}
	md.WriteString("\n")
}
