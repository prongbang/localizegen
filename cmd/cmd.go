package cmd

import (
	"fmt"
	"github.com/prongbang/localizegen/internal/banner"
	"github.com/prongbang/localizegen/internal/localizegen"
	"github.com/prongbang/localizegen/pkg/arguments"
	"os"
)

func Run(args *arguments.Arguments) error {
	banner.Print()
	fmt.Println("--> START")
	fmt.Printf("# Platform: %s\n", args.Platform)
	if args.Locale != "" {
		fmt.Printf("# Language: %s\n", args.Locale)
	} else {
		fmt.Println("# Language: all")
	}

	hasRequire := 0
	if args.Platform == "" {
		hasRequire++
		fmt.Println("> -platform is empty, please use -platform android, ios, flutter")
	}
	if args.Document == "" {
		hasRequire++
		fmt.Println("> -document is empty, please use -document document-id")
	}
	if args.Sheet == "" {
		hasRequire++
		fmt.Println("> -sheet is empty, please use -sheet sheet-id")
	}
	if hasRequire > 0 {
		fmt.Println("<-- DONE")
		os.Exit(0)
	}
	fmt.Println("# Generate: localization")

	uc := localizegen.New()
	state := uc.Process(localizegen.Localization{
		Platform:   args.Platform,
		DocumentID: args.Document,
		SheetID:    args.Sheet,
		Target:     args.Target,
		Filename:   args.Filename,
		Locale:     args.Locale,
	})
	for n := range state {
		if n.Status == "Error" {
			fmt.Println(">", n.Message)
		} else {
			fmt.Println("> Create", n.Filename, "->", n.Status)
		}
	}
	fmt.Println("<-- DONE")
	return nil
}
