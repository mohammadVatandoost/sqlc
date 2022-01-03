package booktest

import (
	"context"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAuthorTable(ctx *context.Context) (authorsTable table.Table) {
	authorsTable = table.NewDefaultTable(table.DefaultConfig())

	info := authorsTable.GetInfo()

	info.AddField("AuthorID", "authorid", db.Text).FieldSortable()

	info.AddField("Name", "name", db.Text)

	info.SetTable("authors").SetTitle("Authors").SetDescription("Authors")

	formList := authorsTable.GetForm()
	formList.AddField("AuthorID", "authorid", db.Text, form.Text)
	formList.AddField("Name", "name", db.Text, form.Text)

	formList.SetTable("authors").SetTitle("Authors").SetDescription("Authors")

	return
}
func GetBookTable(ctx *context.Context) (booksTable table.Table) {
	booksTable = table.NewDefaultTable(table.DefaultConfig())

	info := booksTable.GetInfo()

	info.AddField("BookID", "bookid", db.Text).FieldSortable()

	info.AddField("AuthorID", "authorid", db.Text).FieldSortable()

	info.AddField("Isbn", "isbn", db.Text)

	info.AddField("BookType", "booktype", db.Text)

	info.AddField("Title", "title", db.Text)

	info.AddField("Yr", "yr", db.Text).FieldSortable()

	info.AddField("Available", "available", db.Text)

	info.AddField("Tags", "tags", db.Text)

	info.SetTable("books").SetTitle("Books").SetDescription("Books")

	formList := booksTable.GetForm()
	formList.AddField("BookID", "bookid", db.Text, form.Text)
	formList.AddField("AuthorID", "authorid", db.Text, form.Text)
	formList.AddField("Isbn", "isbn", db.Text, form.Text)
	formList.AddField("BookType", "booktype", db.Text, form.Text)
	formList.AddField("Title", "title", db.Text, form.Text)
	formList.AddField("Yr", "yr", db.Text, form.Text)
	formList.AddField("Available", "available", db.Text, form.Text)
	formList.AddField("Tags", "tags", db.Text, form.Text)

	formList.SetTable("books").SetTitle("Books").SetDescription("Books")

	return
}
