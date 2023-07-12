# hugo-htmx-go-template

Make static Hugo sites dynamic with htmx and go

This is a project template combining [Hugo](https://gohugo.io), [htmx](https://htmx.org), and an optional API server written in Go, using `html/template` or [templ](https://github.com/a-h/templ/) for HTML rendering.

## Why?

Hugo is a fantastic static site building tool, and there are few things about Hugo that can or should be improved.

The existence of this project template does not suggest that all static sites should be dynamic. If your site _can_ be static, it _should_ be static.

Yet there are instances in which one might want to add dynamic functionality to static Hugo sites. That is the purpose of this project template. Not to make all static sites dynamic, but to provide a simple solution to add islands of dynamic behavior to static sites.

This allows us to build fast, easily deployable HTML content, with the added ability of meeting a new class dynamic behavior needs.

Example use cases include

- Contact forms
- Comment systems
- Up/Down vote systems
- You know ... website stuff

You shouldn't have to reach for a SaaS product to offer dynamic content on your static sites.

## About

**First of all, what is meant by "project template"?**

This git repository is to be used as a template from which to create Hugo sites. It does not impose any constraints on how you use Hugo. It is, itself, a simple hugo project created with `hugo site new ...`. You can run `hugo serve` from this project's root directory and it will behave like every other Hugo site you've developed.

**So what _does_ this template provide?**

This template provides example code and simple developer tooling for running Hugo along with an optional API server, written in Go, using `html/template`.

**What developer tooling is provided?**

1. `bin/develop`: This utility both starts the hugo server (`hugo server`) and the API server (`go run server.go`) for development. If the user has `air` installed (https://github.com/cosmtrek/air), API server code will hot-reload when changes to server.go are made.
2. `bin/build`: This utility builds a binary with your entire Hugo site embedded within. This allows Hugo sites to be deployed as a self-contained binary.

## Getting Started

**Clone this project template and initialize it**:

```bash
git clone git@github.com:acaloiaro/hugo-htmx-go-template.git ./my_new_site
cd my_new_site
# The theme 'ananke' is used here because it's the one used by Hugo's quickstart guide: https://gohugo.io/getting-started/quick-start/
# You can use any theme with this project template, so long as you add htmx to its included scripts: <script type="text/javascript"src="/htmx.js"></script>
# fetch-deps.sh does this automatically for the ananke theme
git clone https://github.com/theNewDynamic/gohugo-theme-ananke ./themes/ananke
bin/fetch-deps.sh
go build -o bin/develop internal/cmd/develop/main.go
go build -o bin/build internal/cmd/build/main.go
mkdir public && touch public/.empty
```

If you'll be using the API server, it's useful to install `air` to hot redeploy Go code changes when running `bin/develop`:

```bash
go install github.com/cosmtrek/air@latest
```

To make changes to `templ` templates (.templ files), first install `templ`. Using `templ` is optional. The [goodbye world](https://github.com/acaloiaro/hugo-htmx-go-template/blob/main/server.go#L83) example shows `templ` in action.

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

## Run

To start the development server(s), run `bin/develop`

Hugo will run on its develop port at [http://localhost:1313](http://localhost:1313) and the API server runs on [http://localhost:1314](http://localhost:1314).

Example content is available at: [http://localhost:1313/posts/hello-world/](http://localhost:1313/posts/hello-world/)

`bin/build` can be used to build fat binaries of your Hugo site, which will be available at  [http://localhost:1314](http://localhost:1314) after running `build/server`.

![screenshot](https://user-images.githubusercontent.com/3331648/248586236-1ad03704-4f13-418c-aa9a-3122742c6b8c.png)

## Changing `templ` templates

To use `templ` templates, you will need to `templ` installed. The `air` configuration watches for changes to `templ` files and automatically builds them.

## Deploy

Hugo sites made from this template project can be deployed in two ways.

### Traditional deployment

A traditional deployment is any deployment where Hugo content is deployed to an HTTP server that serves static content. Deploying Hugo content with this project template remains the same as a usual Hugo site.

However, to use the API server in production, you'll need to tell your static site content where to send HTTP requests.

Configure this in `config/production/hugo.toml`

If you've made your API available at `https://api.example.com`, add the following to `hugo.toml`:

```
[Params]
  apiBaseUrl = 'https://api.example.com'
```

The API server binary can be had by using `bin/build` to generate a fat binary, or simply compiling `server.go` directly: `go -o build/server server.go`

### Fat binary deployment

Fat binary deployment is made possible by `bin/build`. The `build/server` executable contains both the API server code and Hugo static site content.

No additional configuration is necessary to deploy a fat binary. Simply copy `build/server` to the system that will run it.

**Note:** To build fat binaries for different systems/architectures than the system performing the build, use the `GOOS` and `GOARCH` to change system/architecture, e.g.

Build a fat binary for OpenBSD and AMD64 architecture
```bash
GOOS=openbsd GOARCH=amd64 bin/build
```

