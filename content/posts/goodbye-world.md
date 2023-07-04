---
title: "Goodbye, World!"
date: 2023-06-23T15:58:00-06:00
draft: false
---

## Goodbye world is a `templ` template example

Every time `Click Me!` is clicked, a request is sent to the API server, which renders `templ GoodbyeWorld()` from `partial/templates.templ`.

{{< html.inline >}}
<button
  hx-get="{{ .Site.Params.apiBaseUrl }}/goodbyeworld.html"
  hx-trigger="click"
  hx-target="#goodbye"
  hx-swap="beforeend">
  Click Me!
</button>
<div id="goodbye"></div>
{{< /html.inline >}}

