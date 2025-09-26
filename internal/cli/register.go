package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/neeeb1/gator/internal/database"
)

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("no arguments found")
	} else if len(cmd.Arguments) > 1 {
		return fmt.Errorf("expected 1 argument, but recieved %d", len(cmd.Arguments))
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Arguments[0],
	}

	newUser, err := s.Db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to create user %s", params.Name)
	}

	s.Config.SetUser(newUser.Name)
	fmt.Println("new user was created")
	fmt.Printf("user name: %s\ncreated at: %s\nupdated at: %s\nuuid: %s\n", newUser.Name, newUser.CreatedAt, newUser.UpdatedAt, newUser.ID)

	return nil
}
