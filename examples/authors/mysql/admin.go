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

	info.AddField(ID, id, db.int64).FieldSortable()

	info.AddField(Name, name, db.string)

	info.AddField(Bio, bio, db.sql.NullString)

	info.SetTable(authors).SetTitle(Authors).SetDescription(Authors)

	formList := authorsTable.GetForm()
	formList.AddField(ID, id, db.int64, form.Text)
	formList.AddField(Name, name, db.string, form.Text)
	formList.AddField(Bio, bio, db.sql.NullString, form.Text)

	formList.SetTable(authors).SetTitle(Authors).SetDescription(Authors)

	return
}
