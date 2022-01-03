package authors

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

	info.AddField("ID", "id", db.Int).FieldSortable()

	info.AddField("Name", "name", db.Text)

	info.AddField("Bio", "bio", db.Text)

	info.SetTable("authors").SetTitle("Authors").SetDescription("Authors")

	formList := authorsTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Text)
	formList.AddField("Name", "name", db.Text, form.Text)
	formList.AddField("Bio", "bio", db.Text, form.Text)

	formList.SetTable("authors").SetTitle("Authors").SetDescription("Authors")

	return
}
