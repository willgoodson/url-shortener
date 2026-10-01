# URL Shortener

A small web app that turns a long URL into a short link. Open the short link and it redirects to the original URL.

## Overview

The goal of this project was to build a minimal URL shortener in Go. You paste a URL into the form and get back a short link with a random 8-character slug. Slug-to-URL mappings live in Redis, and visiting `/{slug}` redirects to the stored URL. The frontend is a single HTML page that uses [htmx](https://htmx.org) to submit the form and show the result without a page reload.

## Tech Stack

| Layer    | Technology              |
|----------|-------------------------|
| Language | Go                      |
| Router   | [chi](https://github.com/go-chi/chi) |
| Database | Redis                   |
| Frontend | HTML + htmx             |

## Project Structure

```
main.go             HTTP server: form page, link creation, redirects
static/index.html   Form UI (htmx)
```

## Getting Started

**Prerequisites:** Go 1.20+, and Redis running on `localhost:6379`

```bash
# Start Redis (if you don't have it installed locally)
docker run -d --name redis -p 6379:6379 redis

# Run the app
go run .
```

Then open http://localhost:3000.

## Routes

| Method | Route     | Description                                                 |
|--------|-----------|-------------------------------------------------------------|
| GET    | `/`       | URL entry form                                              |
| POST   | `/new`    | Form field `url` → returns an HTML link to the short URL    |
| GET    | `/{slug}` | Redirects to the stored URL (or back to `/` if it isn't found) |

## Future Work

- [ ] Validate submitted URLs (allow only `http`/`https`)
- [ ] Create one Redis client at startup and reuse it, instead of connecting on every request
- [ ] Read the Redis address, port and public base URL from environment variables instead of hardcoding `localhost`
- [ ] Use `302`/`307` redirects instead of `308` so browsers don't permanently cache a link that might change
- [ ] Optional custom slugs and link expiration (Redis TTL)
- [ ] Click analytics per link
- [ ] Rate limiting on `/new`
- [ ] Add a Dockerfile and `docker-compose.yml` (app + Redis)
- [ ] Add handler tests
