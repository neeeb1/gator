# Gator 🐊
[![Go](https://img.shields.io/badge/Go-1.25+-blue.svg)](https://golang.org/)  
[![Postgres](https://img.shields.io/badge/Postgres-15+-blueviolet.svg)](https://www.postgresql.org/)  

A **command-line RSS feed aggregator** built with Go and PostgreSQL.  
Gator was built as a Boot.Dev project to demonstrate CLI design, SQL persistence, and working with external RSS feeds.  

---

## ✨ Features

- User registration and login  
- Add and manage RSS feeds  
- Follow feeds per user  
- Browse the latest posts from followed feeds  
- Background feed aggregation at a set interval  
- PostgreSQL persistence  

---

## 📦 Tech Stack

- **Go** (≥ 1.25)  
- **PostgreSQL** (≥ 15)  
- **JSON Config** for user and database settings  

---

## ⚙️ Setup & Installation

```bash
# Install via Go
go install github.com/neeeb1/gator@latest
```

Create a `.gatorconfig.json` file in your home directory (`~/.gatorconfig.json`):

```json
{
  "db_url": "postgres://user:password@localhost:5432/gator_db?sslmode=disable",
  "current_user_string": "example-user"
}
```

> ⚠️ This file is not auto-generated. An example is included in this repo.  

---

## 🚀 Usage

Gator is run entirely from the command line.  

```bash
gator <command> [options]
```

### 🔹 User Commands
```bash
gator register       # Add a new user
gator login          # Set current user
gator users          # List all users
```

### 🔹 Feed Commands
```bash
gator addfeed        # Add a new RSS feed (name + URL)
gator feeds          # List all available feeds
gator follow         # Follow a feed for the current user
gator following      # List feeds current user is following
```

### 🔹 Aggregation & Browsing
```bash
gator agg Xs        # Fetch feeds every Xs
gator browse Y      # Browse latest Y posts for current user
```

### 🔹 Danger Zone
```bash
gator reset          # ⚠️ Clears ALL data from the database
```

---

## 🛠️ Database

- PostgreSQL stores users, feeds, and subscriptions  
- Simple schema designed for learning purposes  

---

## 🧪 Testing

Run any available tests with:

```bash
go test ./...
```
