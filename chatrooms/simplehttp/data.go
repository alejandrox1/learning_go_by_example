package main

import (
	"database/sql"

	"github.com/alejandrox1/setup_sqldb"
	_ "github.com/lib/pq"
)


type Text interface {
	Retrieve(id int) (err error)
	Create() (err error)
	Update() (err error)
	Delete() (err error)
}

type Post struct {
	Db      *sql.DB
	Id      int     `json:"id"`
	Content string  `json:"content"`
	Author  string  `json:"author"`
}

func (p *Post) Retrieve(id int) (err error) {
	err = p.Db.QueryRow("select id, content, author from posts where id = $1",
		id).Scan(&p.Id, &p.Content, &p.Author)
	return
}

func (p *Post) Create() (err error) {
	err = p.Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id",
		p.Content, p.Author).Scan(&p.Id)
	return
}

func (p *Post) Update() (err error) {
	_, err = p.Db.Exec("update posts set content = $2, author= $3 where id = $1",
		p.Id, p.Content, p.Author)
	return
}

func (p *Post) Delete() (err error) {
	_, err = p.Db.Exec("delete from posts where id = $1", p.Id)
	return
}


var Db *sql.DB
func init() {
	sqlSetup := setup_sqldb.SQLSetup{
		DriverName: "postgres",
		SQLScript:  "setup.sql",
	}
	Db, _ = sqlSetup.Init()
	//if err != nil { panic(err)	}
}
