package main

import "net/http"

func routeTodoAdd(resp Responder, r *http.Request) {
	title := r.FormValue("title")
	desc := r.FormValue("desc")
	done := r.FormValue("done")

	if done == "" {
		done = "false"
	}

	if title == "" || desc == "" || done == "" {
		resp.SetStatus(http.StatusBadRequest)
		return
	}

	todo := Todo{
		ID:    generateULID(),
		Title: title,
		Desc:  desc,
		Done:  done == "true",
	}

	db.addTodo(todo)

	htmx := singleTodoTag(todo)

	resp.SetHeader("Content-Type", "text/html").SetStatus(http.StatusOK).SendString(htmx)
}
