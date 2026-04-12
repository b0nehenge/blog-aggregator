package main

import (
	"context"
	"fmt"
	"time"

	"github.com/b0nehenge/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}
	fmt.Printf("Feed: %+v\n", feed)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	ctx := context.Background()

	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("couldn't fetch feeds: %w", err)
	}
	for _, feed := range feeds {
		printFeed(feed)
	}
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %v <name> <url>", cmd.Name)
	}

	u := s.cfg.CurrentUserName
	ctx := context.Background()

	user, err := s.db.GetUser(ctx, u)
	if err != nil {
		return fmt.Errorf("couldn't get user: %w", err)
	}

	_, err = s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		Url:       cmd.Args[1],
		Name:      cmd.Args[0],
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	return nil
}

func printFeed(user database.GetFeedsRow) {
	fmt.Printf(" * Name:    %v\n", user.Name)
	fmt.Printf(" * Url:    %v\n", user.Url)
	fmt.Printf(" * User:    %v\n", user.User)
}
