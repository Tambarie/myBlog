package models

import (
	"fmt"
	"github.com/Tambarie/myBlog/pkg/postgresql"
	_ "github.com/lib/pq"
)

type Blog struct {
	ID string
	Author string
	Title string
	Body string
	Time string
}

func (b *Blog) CreateBlog()  {
	_, err := postgresql.Db.Exec(`INSERT INTO blog(ID, Author,Title,Body,Time) VALUES ($1,$2,$3,$4,$5)
`,b.ID,b.Author,b.Title,b.Body,b.Time)
	if err != nil{
		panic(err)
	}
}



func (b *Blog) GetContent(id string) *Blog {
	row := postgresql.Db.QueryRow("select ID, Author,Title,Body,Time from blog where Id = $1",id)
	err := row.Scan(&b.ID,&b.Author,&b.Title,&b.Body,&b.Time)
	if err != nil{
		panic(err)
	}
	return b
}

func (b *Blog) GetAllContents()(blogPosts  []Blog, err error)   {
	rows, err := postgresql.Db.Query(`SELECT ID,Author,Title,Body,Time FROM blog`)
	if err != nil{
		return
	}
	for rows.Next(){
		bp := Blog{}
		err = rows.Scan(&bp.ID,&bp.Author,&bp.Title,&bp.Body,&bp.Time)
		if err != nil{
			return
		}
		blogPosts = append(blogPosts,bp)
	}
	rows.Close()
	return

}

func (b *Blog) Delete(val string) (err error) {
	_, err = postgresql.Db.Query(`delete from blog where ID = $1`,val)
	return err
}

func (b *Blog) UpdateForm()(err error)  {
	Table := "blog"
	stmt, err := postgresql.Db.Prepare(fmt.Sprintf("UPDATE %s SET Author = $1, Title = $2, Body = $3 WHERE ID = $4",Table))
	if err != nil{
		panic(err)
	}
	_, err = stmt.Exec(b.Author, b.Title, b.Body, b.ID)
	return
}