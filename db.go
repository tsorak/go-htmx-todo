package main

import (
	"sort"
)

type Todo struct {
	ID    string
	Title string
	Desc  string
	Done  bool
}

type TodoDB struct {
	_todos map[string]Todo
}

func CreateTodoDB() TodoDB {
	db := TodoDB{}
	db._todos = make(map[string]Todo)
	return db
}

func (db TodoDB) getTodos() []Todo {
	var todos []Todo

	for _, v := range db._todos {
		todos = append(todos, v)
	}

	return todos
}

func (db TodoDB) getTodosSorted() []Todo {
	var todos []Todo

	var sortedIds []string
	for k := range db._todos {
		sortedIds = append(sortedIds, k)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(sortedIds)))

	for _, k := range sortedIds {
		todos = append(todos, db._todos[k])
	}

	return todos
}

func (db TodoDB) addTodo(todo Todo) {
	_, exists := db._todos[todo.ID]
	if exists {
		return
	}

	db._todos[todo.ID] = todo
}

func (db TodoDB) updateTodoByID(id string, todo Todo) {
	db._todos[id] = todo
}

func (db TodoDB) exists(id string) bool {
	_, exists := db._todos[id]
	return exists
}

func (db TodoDB) deleteTodoByID(id string) bool {
	if !db.exists(id) {
		return false
	}

	delete(db._todos, id)

	return true
}
