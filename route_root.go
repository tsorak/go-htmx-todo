package main

import (
	"net/http"
)

func routeRoot(resp Responder, r *http.Request) {
	if !accepts("text/html", r) {
		resp.SetStatus(http.StatusNotAcceptable)
		return
	}

	htmx := `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8" />
			<title>HTMX Todo</title>
			<script src="https://cdn.tailwindcss.com"></script>
			<script src="https://unpkg.com/htmx.org@1.9.6"></script>
		</head>
		<body class="dark:bg-neutral-900 dark:text-neutral-200">
			<main class="flex flex-col gap-2">
				<h1 class="text-4xl">Go HTMX</h1>
				` + todoForm() + `
				<div id="todo-list" class="flex flex-wrap gap-2">
				` + allTodosTags() + `
				</div>
			</main>
		</body>
	</html>
	`
	// <button hx-get="/motd" hx-swap="outerHTML">Refresh</button>
	resp.SendString(htmx)
}
