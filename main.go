package main

var db TodoDB

func main() {
	db = CreateTodoDB()

	db.addTodo(Todo{
		ID:    generateULID(),
		Title: "Learn Go",
		Desc:  "Learn Go and build a web app",
		Done:  false,
	})

	startWebserver()
}
