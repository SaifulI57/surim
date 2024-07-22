package app

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/urfave/cli/v3"
)

const Name = "surim"

var version = "v0.1"

func init() {
	cli.VersionPrinter = func(c *cli.Command) {
		blue := color.New(color.FgBlue)
		orange := color.New(color.FgHiMagenta)
		fmt.Fprintf(color.Output, "\n%s: version %s, a simple suricata merger rules.\n\n", orange.Sprintf(Name), blue.Sprintf(version))
	}
}

func New() *cli.Command {
	app := &cli.Command{
		Name:    Name,
		Usage:   "A simple suricata merger rules",
		Version: version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Usage:   "Input suricata rules",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Output merged rules",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if !cmd.IsSet("input") && !cmd.IsSet("output") {
				cli.VersionPrinter(cmd)
				return nil
			}
			var wg sync.WaitGroup

			filesRules := make(chan string, 100)
			mergedContent := make(chan string, 100)
			var contentBuilder strings.Builder

			go func() {
				filepath.Walk(cmd.String("input"), func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}

					if !info.IsDir() && strings.HasSuffix(info.Name(), ".rules") {
						filesRules <- path
					}
					return nil
				})
				close(filesRules)
			}()
			go func() {
				for file := range filesRules {
					wg.Add(1)
					go func(filename string) {
						defer wg.Done()
						data, err := os.ReadFile(filename)
						if err != nil {
							fmt.Printf("Error reading file %s: %v\n", filename, err)
							return
						}
						mergedContent <- string(data)
					}(file)
				}
				wg.Wait()
				close(mergedContent)
			}()
			var lenRules = 0
			for content := range mergedContent {
				lenRules += 1
				contentBuilder.WriteString(content)
			}
			if lenRules < 1 {
				fmt.Printf("\nNo Suricata rules found in the input directory.\n")
				return nil
			}
			dir, file := filepath.Split(cmd.String("output"))
			if dir == "" {
				file = cmd.String("output")
			} else {
				if _, err := os.Stat(dir); os.IsNotExist(err) {
					fmt.Printf("\nDirectory does not exist. Creating: %s\n", dir)
					if err := os.MkdirAll(dir, os.ModePerm); err != nil {
						fmt.Printf("\nFailed to create directory: %v\n", err)
						return nil
					}
				} else if err != nil {
					fmt.Printf("\nError checking directory status: %v\n", err)
					return nil
				}
			}
			err := os.WriteFile(filepath.Join(dir, file), []byte(contentBuilder.String()), 0644)
			if err != nil {
				fmt.Fprintf(color.Output, "%s", color.New(color.FgHiRed).Sprint("\nFailed to write output file.\n"))
				return err
			} else {
				currentDir, err := os.Getwd()
				if err != nil {
					fmt.Fprintf(color.Output, "\nFile saved to %s\n", color.New(color.FgGreen).Sprint(cmd.String("output")))
					return nil
				}
				fmt.Fprintf(color.Output, "\nFile saved to %s\n", color.New(color.FgHiCyan).Sprint(filepath.Join(currentDir, cmd.String("output"))))
			}

			return nil
		},
		EnableShellCompletion: true,
	}
	return app
}
