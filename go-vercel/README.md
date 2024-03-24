# how to run

```bash
go run .
```

to test, run

```bash
go run ./test
```

this folder contains the backend for IICBES.

It has two entry point. for local development, the entry point is `server.go`. It uses ECHO framework. for production on Vercel, it uses `/api` folder as the entry point.

# defining route

define a route inside the `server.go` file for local development.
define a route inside the `/api` folder for production on Vercel by creating a new file with the name of the route.

DO BOTH when creating a new endpoint.

To reduce code duplication, both handler is created inside `handler` folder.

ECHO doesn't have hot reload. you can install `air` or `fresh` (google it) to have hot reload.

(drop vercel deployment for now. not enough human resource)
