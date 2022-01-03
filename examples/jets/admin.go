package jets

import (
	"context"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetJetTable(ctx *context.Context) (jetsTable table.Table) {
	jetsTable = table.NewDefaultTable(table.DefaultConfig())

	info := jetsTable.GetInfo()

	info.AddField("ID", "id", db.Int).FieldSortable()

	info.AddField("PilotID", "pilotid", db.Int).FieldSortable()

	info.AddField("Age", "age", db.Int).FieldSortable()

	info.AddField("Name", "name", db.Text)

	info.AddField("Color", "color", db.Text)

	info.SetTable("jets").SetTitle("Jets").SetDescription("Jets")

	formList := jetsTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Text)
	formList.AddField("PilotID", "pilotid", db.Int, form.Text)
	formList.AddField("Age", "age", db.Int, form.Text)
	formList.AddField("Name", "name", db.Text, form.Text)
	formList.AddField("Color", "color", db.Text, form.Text)

	formList.SetTable("jets").SetTitle("Jets").SetDescription("Jets")

	return
}
func GetLanguageTable(ctx *context.Context) (languagesTable table.Table) {
	languagesTable = table.NewDefaultTable(table.DefaultConfig())

	info := languagesTable.GetInfo()

	info.AddField("ID", "id", db.Int).FieldSortable()

	info.AddField("Language", "language", db.Text)

	info.SetTable("languages").SetTitle("Languages").SetDescription("Languages")

	formList := languagesTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Text)
	formList.AddField("Language", "language", db.Text, form.Text)

	formList.SetTable("languages").SetTitle("Languages").SetDescription("Languages")

	return
}
func GetPilotTable(ctx *context.Context) (pilotsTable table.Table) {
	pilotsTable = table.NewDefaultTable(table.DefaultConfig())

	info := pilotsTable.GetInfo()

	info.AddField("ID", "id", db.Int).FieldSortable()

	info.AddField("Name", "name", db.Text)

	info.SetTable("pilots").SetTitle("Pilots").SetDescription("Pilots")

	formList := pilotsTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Text)
	formList.AddField("Name", "name", db.Text, form.Text)

	formList.SetTable("pilots").SetTitle("Pilots").SetDescription("Pilots")

	return
}
func GetPilotLanguageTable(ctx *context.Context) (pilot_languagesTable table.Table) {
	pilot_languagesTable = table.NewDefaultTable(table.DefaultConfig())

	info := pilot_languagesTable.GetInfo()

	info.AddField("PilotID", "pilotid", db.Int).FieldSortable()

	info.AddField("LanguageID", "languageid", db.Int).FieldSortable()

	info.SetTable("pilot_languages").SetTitle("PilotLanguages").SetDescription("PilotLanguages")

	formList := pilot_languagesTable.GetForm()
	formList.AddField("PilotID", "pilotid", db.Int, form.Text)
	formList.AddField("LanguageID", "languageid", db.Int, form.Text)

	formList.SetTable("pilot_languages").SetTitle("PilotLanguages").SetDescription("PilotLanguages")

	return
}
