package main

import "strings"

func isDoneElem(todo Todo) string {
	hxtags := `hx-patch="/todo?id=` + todo.ID + `" hx-target="#TODO-` + todo.ID + `" hx-swap="outerHTML"`

	if todo.Done {
		return `
			<span ` + hxtags + ` class="bg-green-600 p-1 rounded-md">Done</span>
		`
	}
	return `
		<span ` + hxtags + ` class="bg-red-600 p-1 rounded-md break-none">Not Done</span>
	`
}

func singleTodoTag(todo Todo) string {
	return `
		<div id="TODO-` + todo.ID + `" class="bg-yellow-600 p-2">
			<h3 class="text-lg font-bold">` + todo.Title + `</h3>
			<p class="mb-1">` + todo.Desc + `</p>
			<div class="flex gap-2 justify-between">
				` + isDoneElem(todo) + `
				<button hx-delete="/todo?id=` + todo.ID + `" hx-target="#TODO-` + todo.ID + `" hx-swap="outerHTML" class="bg-red-600 p-1 rounded-md font-bold">
					DELETE
				</button>
			</div>
		</div>
	`
	// <input type="checkbox" ` + checked(todo.Done) + ` hx-patch="/todo?id=` + todo.ID + `" hx-target="#TODO-` + todo.ID + `" hx-swap="outerHTML">
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
