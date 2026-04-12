# Gator (Blog Aggregator)

A CLI-based RSS feed aggregator written in Go. Gator allows users to register, log in, and subscribe to their favorite RSS feeds. It uses a PostgreSQL database for persistent storage of users, feeds, and followings.

## Features

- **User Management**: Register and switch between different users.
- **Feed Management**: Add RSS feeds with names and URLs.
- **Subscription Tracking**: Follow or unfollow feeds for specific users.
- **Feed Listing**: View all registered feeds and their creators.
- **RSS Integration**: Fetches feed content (currently basic functionality).

## Prerequisites

Before installing Gator, ensure you have the following:

- **Go**: 1.26 or higher.
- **PostgreSQL**: A running instance with a database for Gator.
- **sqlc** (Optional): Only needed for regenerating database code from SQL queries.

## Installation

To install Gator, clone the repository and build the binary:

```bash
git clone https://github.com/b0nehenge/blog-aggregator.git
cd blog-aggregator
go build -o gator
```

You can then move the `gator` binary to a directory in your `PATH` for easier access.

## Configuration

Gator requires a configuration file named `.gatorconfig.json` in your home directory. This file stores your database connection URL and the currently active user.

Example `.gatorconfig.json`:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": "your_username"
}
```

## Usage

Run Gator using the following command structure:

```bash
./gator <command> [args...]
```

### Commands

- `register <name>`: Create a new user and set them as the current user.
- `login <name>`: Switch the current user to an existing one.
- `users`: List all registered users.
- `addfeed <name> <url>`: Add a new feed (requires being logged in). Automatically follows the new feed.
- `feeds`: List all registered feeds.
- `follow <feed_url>`: Follow an existing feed by its URL (requires being logged in).
- `following`: List all feeds followed by the current user (requires being logged in).
- `unfollow <feed_url>`: Unfollow a feed by its URL (requires being logged in).
- `agg <time_between_reqs>`: Continuously fetch and store posts from all followed feeds. Example: `agg 1m`.
- `browse [limit]`: View the latest posts from feeds followed by the current user. Default limit is 2.
- `reset`: Clear all users and feeds from the database (use with caution!).

## License

This project is open-source and available under the MIT License.
