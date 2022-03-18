package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func mySql() {
	// 连接
	db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// 创建表
	{
		query := `
		CREATE TABLE users (
			id INT AUTO_INCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id)
		);`
		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	// 插入/新增
	{
		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, "username", "password", time.Now())
		if err != nil {
			log.Fatal(err)
		}
		id, err := result.LastInsertId()
		fmt.Println(id)
	}

	// 查询单行
	{
		var (
			id       int
			username string
			password string
			createAt time.Time
		)
		query := `select id, username, password, create_at from users where id = ?`
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createAt); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, username, password, createAt)
	}

	// 查询多行
	{
		type user struct {
			id       int
			username string
			password string
			createAt time.Time
		}

		rows, err := db.Query(`select id, username, password, create_at from users`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []user
		for rows.Next() {
			var u user
			err := rows.Scan(&u.id, &u.username, &u.password, &u.createAt)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v", users)
	}

	// 删除
	{
		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}

type Todo struct {
	Title string
	Done  bool
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("layout.html"))

		_ = TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}

		tmpl.Execute(w, "data goes here")
		//fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	// 静态资源
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":80", nil)
}
