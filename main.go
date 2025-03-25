package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gitsight/go-vcsurl"
	"github.com/go-git/go-git/v5"
	"github.com/urfave/cli/v3"
)

func homeDir() string {
	home := os.Getenv("KLONE_HOME")
	if home == "" {
		home = filepath.Join(os.Getenv("HOME"), "klone")
	}
	return home
}

func main() {
	cmd := &cli.Command{
		Name:  "klone",
		Usage: "clone git repository using url as destination directory",
		Action: func(_ context.Context, cmd *cli.Command) error {
			url := cmd.Args().Get(0)
			if url == "" {
				return cli.Exit("klone help for details", 1)
			}
			info, err := vcsurl.Parse(url)
			if err != nil {
				return cli.Exit(err, 1)
			}
			path := filepath.Join(homeDir(), string(info.Host), info.Username, info.Name)
			fmt.Printf("cloning to %s\n", path)
			err = os.MkdirAll(path, 0755)
			if err != nil {
				return cli.Exit(err, 1)
			}
			_, err = git.PlainClone(path, false, &git.CloneOptions{
				URL:      url,
				Progress: os.Stdout,
			})
			if err != nil {
				return cli.Exit(err, 1)
			}
			fmt.Printf("to change directory: cd %s\n", path)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
