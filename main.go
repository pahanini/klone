package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gitsight/go-vcsurl"
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

			path := filepath.Join(homeDir(), string(info.Host), info.FullName)
			fmt.Printf("cloning to %s\n", path)
			err = os.MkdirAll(path, 0755)
			if err != nil {
				return cli.Exit(err, 1)
			}

			var execCmd *exec.Cmd
			execCmd = exec.Command("git", "clone", url, path)
			execCmd.Stdout = os.Stdout
			execCmd.Stderr = os.Stderr
			err = execCmd.Run()
			if err != nil {
				return cli.Exit(err, 1)
			}

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
