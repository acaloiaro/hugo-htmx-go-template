---
title: "Hello, World!"
date: 2023-06-24T15:58:00-06:00
draft: false
---

## Let's dynamically fetch content from our API.

What you're currently reading is typical Hugo content authored in Markdown. Below the break we see content loaded dynamically via HTMX from our API server.

Open `content/posts/hello-world.md` to follow along with the code.

---

{{< htmx.inline >}}
<div
  hx-get="{{ .Site.Params.apiBaseUrl }}/hello_world"
  hx-trigger="load"
  hx-vals='js:{"name": new URLSearchParams(window.location.search).get("name")}'
  hx-on="htmx:configRequest: console.log('detail:', event.detail); event.detail.headers='';" />
  <p>Content loading from API...</p>
</div>
{{< /htmx.inline >}}

---

## Let's add a form

Writing markdown content is great, but sometimes you want to add hypermedia controls: buttons, forms, etc.

See `func helloWorldForm` from `server.go` to see what's going on behind the scenes with this form.

Here we are using a Hugo concept called an "inline shortcode". This allows us to change the base URL for our API server, depending on whether we're in development or production.


{{< htmx.inline >}}
<!-- Make sure the api base URL is set -->
<form hx-post="{{ .Site.Params.apiBaseUrl }}/hello_world_form">
  <label>Name</label>
  <input type="text" name="name">
  <br/>
  <button>Submit</button>
</form>
{{< /htmx.inline >}}


---

## Let's add a simple button that loads content

{{< htmx.inline >}}
<div
  hx-get="/content.html"
  hx-trigger="click" />
  <button>Load content</button>
</div>
{{< /htmx.inline >}}


