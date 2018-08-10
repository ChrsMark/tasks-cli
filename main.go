package main

import (
	"github.com/gophersize/task/cmd"
	"github.com/mitchellh/go-homedir"
	"path/filepath"
	"github.com/gophersize/task/db"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.Init(dbPath)
	if err != nil {
		panic(err)
	}
	cmd.RootCmd.Execute()
}
