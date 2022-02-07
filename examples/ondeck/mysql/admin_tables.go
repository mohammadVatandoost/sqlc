package ondeck

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

// Generators is a map of table models.
//
// The key of generators is the prefix of table info url.
// The corresponding value is the Form and TableName data.

var Generators = map[string]table.Generator{
	"city":  GetCityTable,
	"venue": GetVenueTable,
}
