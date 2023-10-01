package main

import "net/http"

func routeTodoDone(resp Responder, r *http.Request) {
	id := r.URL.Query().Get("id")
	todo := db.getTodoByID(id)

	todo.Done = !todo.Done
	db.updateTodoByID(id, todo)

	resp.SendString(singleTodoTag(todo))
}
