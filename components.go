package main

import "strings"

func singleTodoTag(todo Todo) string {
	return `
		<div id="TODO-` + todo.ID + `">
			<h3>` + todo.Title + `</h3>
			<p>` + todo.Desc + `</p>
			<button hx-delete="/todo?id=` + todo.ID + `" hx-target="#TODO-` + todo.ID + `" hx-swap="outerHTML">
				Delete
			</button>
		</div>
	`
}

func allTodosTags() string {
	todos := db.getTodosSorted()
	var tags []string
	for _, todo := range todos {
		tags = append(tags, singleTodoTag(todo))
	}

	return strings.Join(tags, "\n")
}

func todoForm() string {
	htmx, _ := getFileContentsIntoString("htmx/form.html")
	return htmx
}
