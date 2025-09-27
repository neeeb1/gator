package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/neeeb1/gator/internal/cli"
	"github.com/neeeb1/gator/internal/config"
	"github.com/neeeb1/gator/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config: %v\n", err)
		return
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Printf("error opening database: %v", err)
		return
	}
	dbQueries := database.New(db)

	state := cli.State{
		Config: &cfg,
		Db:     dbQueries,
	}

	commands := cli.Commands{
		Cmds: make(map[string]func(*cli.State, cli.Command) error),
	}
	commands.Register("login", cli.HandlerLogin)
	commands.Register("register", cli.HandlerRegister)
	commands.Register("reset", cli.HandlerReset)
	commands.Register("users", cli.HandlerUsers)
	commands.Register("agg", cli.HandlerAgg)
	commands.Register("addfeed", cli.MiddlewareLoggedIn(cli.HandlerAddFeed))
	commands.Register("feeds", cli.HandlerFeeds)
	commands.Register("follow", cli.MiddlewareLoggedIn(cli.HandlerFollow))
	commands.Register("following", cli.MiddlewareLoggedIn(cli.HandlerFollowing))
	commands.Register("unfollow", cli.MiddlewareLoggedIn(cli.HandlerUnfollow))
	commands.Register("browse", cli.MiddlewareLoggedIn(cli.HandlerBrowse))

	args := os.Args
	if len(args) < 2 {
		fmt.Printf("error: not enough arguments\n")
		os.Exit(1)
	}

	cmd := cli.Command{
		Name:      args[1],
		Arguments: args[2:],
	}

	err = commands.Run(&state, cmd)
	if err != nil {
		fmt.Printf("error running specified command: %v\n", err)
		os.Exit(1)
	}
}
