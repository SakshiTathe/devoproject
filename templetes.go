package main

import "html/template"

var pageTmpl = template.Must(template.New("page").Parse(`<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8" />
<meta name="viewport" content="width=device-width, initial-scale=1" />
<title>Simple List</title>
<style>
  :root { font-family: system-ui, -apple-system, Segoe UI, Roboto, Arial, sans-serif; }
  body { margin: 0; background: #f6f7fb; }
  .wrap { max-width: 720px; margin: 40px auto; padding: 24px; background: #fff; border-radius: 16px; box-shadow: 0 6px 24px rgba(0,0,0,.08); }
  h1 { margin: 0 0 16px; font-size: 24px; }
  form { display: flex; gap: 8px; margin-bottom: 16px; }
  input[type="text"] { flex: 1; padding: 12px 14px; border: 1px solid #dcddea; border-radius: 12px; outline: none; }
  input[type="text"]:focus { border-color: #8aa4ff; box-shadow: 0 0 0 3px rgba(138,164,255,.25); }
  button { padding: 12px 16px; border: 0; border-radius: 12px; cursor: pointer; background: #3b82f6; color: #fff; font-weight: 600; }
  ul { list-style: none; padding: 0; margin: 0; display: grid; gap: 8px; }
  li { display: flex; align-items: center; justify-content: space-between; padding: 10px 12px; border: 1px solid #e6e7f2; border-radius: 10px; background: #fafbff; }
  .empty { color: #666; padding: 8px; }
  .row { display: flex; gap: 8px; align-items: center; }
  .del { background: #ef4444; }
</style>
</head>
<body>
  <div class="wrap">
    <h1>Simple List</h1>

    <form method="POST" action="/add" autocomplete="off">
      <input type="text" name="text" placeholder="Add an item…" />
      <button type="submit">Add</button>
    </form>

    {{if .Items}}
      <ul>
        {{range .Items}}
        <li>
          <span>#{{.ID}} &middot; {{.Text}}</span>
          <form method="POST" action="/delete" class="row">
            <input type="hidden" name="id" value="{{.ID}}" />
            <button class="del">Delete</button>
          </form>
        </li>
        {{end}}
      </ul>
    {{else}}
      <div class="empty">No items yet. Add your first one above!</div>
    {{end}}
  </div>
</body>
</html>`))
