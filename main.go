package main

import (
	"fmt"
	"os"

	"github.com/Nekroze/tuireframe/pkg/display"
	"github.com/Nekroze/tuireframe/pkg/input"
	"github.com/docopt/docopt-go"
)

var version = `1.0.0`
var usage = fmt.Sprintf(`tuiframe v%s - sketch terminal user interfaces.

Usage:
  tuiframe <screen> <meta>
  tuiframe -h | --help
  tuiframe --version

Options:
  -h --help     Show this screen.
  --version     Show version.
`, version)

type Config struct {
	Screen string
	Meta   string
}

var parser = &docopt.Parser{
	HelpHandler:  docopt.PrintHelpOnly,
	OptionsFirst: true,
}

func main() {
	opts, err := docopt.ParseArgs(usage, os.Args[1:], version)
	if err != nil {
		panic(err)
	}
	var config Config
	if err = opts.Bind(&config); err != nil {
		panic(err)
	}
	if err = loadAndDisplay(config); err != nil {
		panic(err)
	}
}

func loadAndDisplay(config Config) error {
	screen, err := input.LoadScreen(config.Screen)
	if err != nil {
		return err
	}
	meta, err := input.LoadMeta(config.Meta)
	if err != nil {
		return err
	}
	meta.ApplyToScreen(screen)
	return display.Screen(screen)
}
