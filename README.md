# gitRats CLI

> Official command-line client for [gitRats](https://github.com/rafael-vincensi/gitrats), a social accountability and gamified check-in platform for developers. Built in Go with Cobra, `gitrats-cli` lets you register, authenticate via email/password or GitHub OAuth, join groups, track your daily streak and weekly GitHub commits, and submit photo-verified coding check-ins straight from your terminal.

## Tech Stack

<div align="center">

[![Go](https://img.shields.io/badge/Go-1.27-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Cobra](https://img.shields.io/badge/Cobra-CLI-00ADD8?style=for-the-badge&logo=gnubash&logoColor=white)](https://github.com/spf13/cobra)
[![GitHub OAuth](https://img.shields.io/badge/GitHub_OAuth-181717?style=for-the-badge&logo=github&logoColor=white)](https://docs.github.com/en/apps/oauth-apps)
[![gitRats API](https://img.shields.io/badge/gitRats-Backend_API-brightgreen?style=for-the-badge&logo=spring&logoColor=white)](https://github.com/rafael-vincensi/gitrats)

</div>

## Ecosystem

| Component | Repository | Stack | Description |
|-----------|------------|-------|-------------|
| **CLI Client** | This repository | Go, Cobra | Terminal client to register, authenticate, join groups, view stats, and submit photo check-ins |
| **Backend API** | [rafael-vincensi/gitrats](https://github.com/rafael-vincensi/gitrats) | Java 21, Spring Boot, PostgreSQL | Core REST API, authentication, streak tracking, GitHub commit aggregation, and Cloudinary uploads |

## Features

- Create an account and authenticate via email/password or browser-based GitHub OAuth
- Local session and API URL persistence across terminal runs
- View personal profile details, current/highest check-in streak, total check-ins, and weekly GitHub commits
- List all joined groups with invite codes, total weekly check-ins, and total weekly commits formatted in clean terminal tables
- Join groups quickly using an invite code
- Submit photo-verified coding check-ins (`multipart/form-data`) directly from local image files

## Architecture

The CLI is structured around Cobra commands (`cmd/`) backed by internal packages for HTTP communication (`internal/api`) and local credential persistence (`internal/config`):

```text
gitrats-cli/
├── cmd/
│   ├── root.go       # Root Cobra command definition
│   ├── register.go   # User registration (POST /users)
│   ├── login.go      # Email/password & GitHub OAuth authentication
│   ├── me.go         # Profile, streaks, and weekly GitHub commits (GET /users/me)
│   ├── groups.go     # Group listing & invite-code join (GET /groups, POST /groups/{code}/join)
│   └── checkin.go    # Group resolution & multipart photo check-in upload
├── internal/
│   ├── api/          # HTTP client, JWT injection, and multipart/form-data builder
│   └── config/       # Local config & JWT persistence
└── main.go           # CLI entry point
```

## Commands Reference

### `gitrats register`

Create a new gitRats account.

| Flag | Shorthand | Required | Description |
|------|-----------|----------|-------------|
| `--email` | `-e` | Yes | Account email address |
| `--password` | `-p` | Yes | Account password |
| `--username` | `-u` | No | Public username |
| `--name` | `-n` | No | Display name |

```bash
gitrats register -u rafinha -n "rafinha goat" -e rafinhagoat@example.com -p secret123
```

---

### `gitrats login`

Authenticate with the gitRats API and save the JWT locally. Supports standard credentials or GitHub OAuth via your default browser.

| Flag | Shorthand | Required | Description |
|------|-----------|----------|-------------|
| `--email` | `-e` | With password | Account email address |
| `--password` | `-p` | With email | Account password |
| `--github` | — | No | Authenticate using GitHub OAuth in the browser |
| `--client-id` | — | With `--github` | GitHub OAuth Client ID (defaults to `GITRATS_GITHUB_CLIENT_ID`) |

```bash
# Email and password
gitrats login -e dev@example.com -p secret123

# GitHub OAuth
gitrats login --github --client-id <your-github-client-id>
```

---

### `gitrats me`

Display your authenticated profile, current and highest daily streak, total check-ins, and weekly GitHub commit count.

```bash
gitrats me
```

---

### `gitrats groups`

List all groups you belong to alongside their invite codes, weekly check-ins, and weekly GitHub commits.

```bash
gitrats groups
```

#### `gitrats groups join [invite-code]`

Join an existing group (up to 20 members) using its invite code.

```bash
gitrats groups join ABC123
```

---

### `gitrats checkin`

Publish a new photo-verified coding check-in to a group using its invite code.

| Flag | Shorthand | Required | Description |
|------|-----------|----------|-------------|
| `--group` | `-g` | Yes | Target group invite code |
| `--title` | `-t` | Yes | Check-in title |
| `--image` | `-i` | Yes | Path to the check-in image file |
| `--description` | `-d` | No | Activity description |

```bash
gitrats checkin \
  -g ABC123 \
  -t "Implemented OAuth flow" \
  -d "Added --github flag to gitrats login" \
  -i ./screenshot.png
```

## Getting Started

### Prerequisites

- Go 1.27 or higher
- A running instance of the [gitRats Backend API](https://github.com/rafael-vincensi/gitrats)

### Installation

#### Using `go install`

```bash
go install github.com/rafael-vincensi/gitrats-cli@latest
```

#### Build from Source

```bash
git clone https://github.com/rafael-vincensi/gitrats-cli.git
cd gitrats-cli
go build -o gitrats .
```

### Environment Variables

| Variable | Description |
|----------|-------------|
| `GITRATS_GITHUB_CLIENT_ID` | Default GitHub OAuth Client ID used when running `gitrats login --github` without `--client-id` |

---

<div align="center">

Built by [rafael-vincensi](https://github.com/rafael-vincensi)

</div>
