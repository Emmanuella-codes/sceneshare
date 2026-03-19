# SceneShare

SceneShare lets a user share a specific moment from a video with a short link.

The working path is:
- capture a YouTube timestamp from the browser extension
- create a short link through the Go API
- open the link in the Next.js app
- preview the shared moment
- redirect through the API to the platform deep link

Current platform support:
- `YouTube`: supported
- `Netflix`, `Prime`, other streaming platforms: not supported yet

## Repo Layout

```text
apps/
  api/        Go API for creating, resolving, and redirecting links
  web/        Next.js app for landing page and link preview/redirect UI
  extension/  WXT browser extension for capturing and sharing moments
packages/
  types/      Shared TypeScript types used across apps
```

## Architecture

The current flow is:

1. The extension reads the current YouTube video id, title, thumbnail, and playback time.
2. The extension sends that payload to the API with `POST /api/v1/links`.
3. The API stores the link metadata and returns a short link response.
4. The short link opens the web app preview route.
5. The web app fetches the link metadata and shows a preview card.
6. The preview page redirects to the API redirect handler.
7. The API redirects the user to YouTube with the saved timestamp.

## Requirements

- Node.js `>= 22`
- Yarn `4`
- Go `1.25.x`
- Postgres

## Workspace

This repo uses Yarn workspaces from the root.

Install dependencies from the root:

```bash
yarn install
```

## Local Environment

Current local ports:

- web: `http://localhost:3001`
- api: `http://localhost:4006`

## Running Locally

Start each app in its own terminal.

API:
```bash
cd apps/api
go run main.go
```

Web:
```bash
cd apps/web
yarn dev
```

Extension dev:
```bash
cd apps/extension
yarn dev
```

## Testing the Flow

1. Start the API, web app, and extension dev server.
2. Load the extension in Chrome.
3. Open a YouTube watch page.
4. Pause on a timestamp.
5. Open the extension popup.
6. Create a link.
7. Open the generated short link.
8. Confirm the web preview page renders.
9. Confirm the redirect lands on YouTube at the expected timestamp.

