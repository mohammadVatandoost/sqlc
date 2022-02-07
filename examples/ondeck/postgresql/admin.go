package ondeck

import (
	"context"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetCityTable(ctx *context.Context) (cityTable table.Table) {
	cityTable = table.NewDefaultTable(table.DefaultConfig())

	info := cityTable.GetInfo()

	info.AddField("Slug", "slug", db.Text)

	info.AddField("Name", "name", db.Text)

	info.SetTable("city").SetTitle("Citys").SetDescription("Citys")

	formList := cityTable.GetForm()
	formList.AddField("Slug", "slug", db.Text, form.Text)
	formList.AddField("Name", "name", db.Text, form.Text)

	formList.SetTable("city").SetTitle("Citys").SetDescription("Citys")

	return
}
func GetVenueTable(ctx *context.Context) (venueTable table.Table) {
	venueTable = table.NewDefaultTable(table.DefaultConfig())

	info := venueTable.GetInfo()

	info.AddField("ID", "id", db.Int).FieldSortable()

	info.AddField("Status", "status", db.Text)

	info.AddField("Statuses", "statuses", db.Text)

	info.AddField("Slug", "slug", db.Text)

	info.AddField("Name", "name", db.Text)

	info.AddField("City", "city", db.Text)

	info.AddField("SpotifyPlaylist", "spotifyplaylist", db.Text)

	info.AddField("SongkickID", "songkickid", db.Text)

	info.AddField("Tags", "tags", db.Text)

	info.AddField("CreatedAt", "createdat", db.Timestamp)

	info.SetTable("venue").SetTitle("Venues").SetDescription("Venues")

	formList := venueTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Text)
	formList.AddField("Status", "status", db.Text, form.Text)
	formList.AddField("Statuses", "statuses", db.Text, form.Text)
	formList.AddField("Slug", "slug", db.Text, form.Text)
	formList.AddField("Name", "name", db.Text, form.Text)
	formList.AddField("City", "city", db.Text, form.Text)
	formList.AddField("SpotifyPlaylist", "spotifyplaylist", db.Text, form.Text)
	formList.AddField("SongkickID", "songkickid", db.Text, form.Text)
	formList.AddField("Tags", "tags", db.Text, form.Text)
	formList.AddField("CreatedAt", "createdat", db.Timestamp, form.Text)

	formList.SetTable("venue").SetTitle("Venues").SetDescription("Venues")

	return
}
