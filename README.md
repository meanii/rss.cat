# RSS.cat Telegram Bot

RSS.cat is a Telegram bot that allows users to subscribe to RSS feeds and receive updates directly in Telegram. It supports adding feeds, listing your feeds, bot statistics, and more.

## Features
- Add RSS feeds and get updates automatically
- List your added feeds
- View bot statistics (/stats)
- Help and usage instructions

## Getting Started

### Prerequisites
- Docker & Docker Compose
- Telegram Bot Token (from @BotFather)

### Build and Run with Docker Compose
1. Clone this repository:
   ```sh
   git clone https://github.com/meanii/rss.cat.git
   cd rss.cat
   ```
2. Set your Telegram bot token in an `.env` file:
   ```sh
   echo "TOKEN=your-telegram-bot-token" > .env
   ```
3. Build and start the bot:
   ```sh
   docker compose up --build
   ```

### Manual Build (without Docker)
1. Install Go 1.24+
2. Build:
   ```sh
   go build -o rsscat main.go
   ./rsscat
   ```

## Usage
- `/start` — Welcome message
- `/add <rss-url>` — Add a new RSS feed
- `/myfeeds` — List your added RSS feeds
- `/stats` — Show bot statistics
- `/help` — Show help message

## Environment Variables
- `TOKEN` — Your Telegram bot token

## Data Persistence
- The SQLite database (`rss.cat.db`) is persisted via Docker volume in Compose setup.

## License
MIT

## Contributing
Pull requests and issues are welcome!
