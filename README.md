# Gator 🐊  
A simple CLI-based news & blog aggregator built with Go and PostgreSQL.  
Gator lets you subscribe to RSS feeds, follow/unfollow them, and browse the latest posts — all from your terminal.

---

## Requirements
- [Go](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)

---

## Installation
Clone the repository and build the binary:

```bash
git clone https://github.com/yourusername/gator.git
cd gator
go install
```

## Configuration
You will have to manually create a PostgreSQL database called **gator**. Run the migrations in `sql/schema` (in order).

Gator requires a config file in your home directory:
`~/.gatorconfig.json`

It must look like this:
```
{
    "db_url": "connection_string"
}
```

The <connection_string> looks like this:
```
"postgres://username:@localhost:5432/gator?sslmode=disable"
```

Make sure to use your username and test the connection with:
```bash
psql <connection_string>
```

## Usage
Run commands with:
`gator <command> [arguments]`

### Commands
- register <name> – Register a new user
- login <name> – Log in as an existing user
- users – List all users (highlights the logged-in user)
- addfeed <name> <url> – Add an RSS feed
- feeds – List all available feeds
- follow <url> – Follow a feed
- following – Show feeds you follow
- unfollow <url> – Unfollow a feed
- browse [limit] – Browse posts from followed feeds (default limit: 2)
- agg <time> – Continuously fetch feeds at a given interval (e.g., 1m, 10s). ⚠️ Typically run this in a separate terminal.