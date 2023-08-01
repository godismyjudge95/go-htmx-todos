package todos

import (
	db "danielweaver.dev/go-todo/database"
	"danielweaver.dev/go-todo/utils"
)

type Todo struct {
	ID        int    `json:"id" form:"id"`
	Title     string `json:"title" form:"title"`
	Completed bool   `json:"completed" form:"completed"`
}

func GetAllTodos() []Todo {
	SQL := db.Open()

	rows, err := SQL.Query("SELECT * FROM todos")
	utils.CheckError(err, "Error starting retrieval of todos from the database")
	defer rows.Close()

	todos := make([]Todo, 0)

	for rows.Next() {
		newTodo := Todo{}
		err = rows.Scan(&newTodo.ID, &newTodo.Title, &newTodo.Completed)
		utils.CheckError(err, "Error parsing todo row from the database")

		todos = append(todos, newTodo)
	}

	err = rows.Err()
	utils.CheckError(err, "Error finishing retrieval of todos from the database")

	SQL.Close()

	return todos
}

func StoreTodo(todo *Todo) {
	SQL := db.Open()

	stmt, err := SQL.Prepare("INSERT INTO todos (Title, Completed) VALUES (?, ?)")
	utils.CheckError(err, "Error preparing to insert todo into the database")

	_, err = stmt.Exec(todo.Title, todo.Completed)
	utils.CheckError(err, "Error inserting todo into the database")
	defer stmt.Close()

	SQL.Close()
}

func UpdateTodo(todo *Todo) {
	SQL := db.Open()

	stmt, err := SQL.Prepare("UPDATE todos SET Title = ?, Completed = ? WHERE ID = ?")
	utils.CheckError(err, "Error preparing to update todo in the database")

	_, err = stmt.Exec(todo.Title, todo.Completed, todo.ID)
	utils.CheckError(err, "Error updating todo in the database")
	defer stmt.Close()

	SQL.Close()
}

func DeleteTodo(todo *Todo) {
	SQL := db.Open()

	stmt, err := SQL.Prepare("DELETE FROM todos WHERE ID = ?")
	utils.CheckError(err, "Error preparing to delete todo from the database")

	_, err = stmt.Exec(todo.ID)
	utils.CheckError(err, "Error deleting todo from the database")
	defer stmt.Close()

	SQL.Close()
}

func CreateTable() {
	SQL := db.Open()

	stmt, err := SQL.Prepare(`
		CREATE TABLE IF NOT EXISTS todos (
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			Title VARCHAR(64) NULL,
			Completed INTEGER NULL
		);
	`)
	utils.CheckError(err, "Error preparing to create the todos database table")
	defer stmt.Close()

	_, err = stmt.Exec()
	utils.CheckError(err, "Error creating the todos database table")

	SQL.Close()
}
