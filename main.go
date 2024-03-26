package main

import (
	"github.com/Taliker/Gokedex/repl"
	"github.com/Taliker/Gokedex/repl/commands"
	"time"
)

func main() {
	var config = &commands.Config{}
	config.Cache = config.Cache.NewCache(5 * time.Minute)
	repl.StartREPL(config)
}
