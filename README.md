# Stockyard Saloon

**Community forum — threads, replies, categories, moderation, no JavaScript framework bloat**

Part of the [Stockyard](https://stockyard.dev) family of self-hosted developer tools.

## Quick Start

```bash
docker run -p 9310:9310 -v saloon_data:/data ghcr.io/stockyard-dev/stockyard-saloon
```

Or with docker-compose:

```bash
docker-compose up -d
```

Open `http://localhost:9310` in your browser.

## Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `9310` | HTTP port |
| `DATA_DIR` | `./data` | SQLite database directory |
| `SALOON_LICENSE_KEY` | *(empty)* | Pro license key |

## Free vs Pro

| | Free | Pro |
|-|------|-----|
| Limits | 3 categories, 50 threads | Unlimited categories and threads |
| Price | Free | $4.99/mo |

Get a Pro license at [stockyard.dev/tools/](https://stockyard.dev/tools/).

## Category

Creator & Small Business

## License

Apache 2.0
