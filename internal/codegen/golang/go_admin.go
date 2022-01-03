package golang

import "sync"

type AdminTableGenerator struct {
	TableName string
	Method    string
}

// technical debt: handling sql custom type is developing in go-admin

var adminTypesMap = map[string]string{
	"int64":          "Int",
	"int32":          "Int",
	"string":         "Text",
	"sql.NullString": "Text",
	"time.Time":      "Timestamp",
}
var adminMapLock sync.Mutex
