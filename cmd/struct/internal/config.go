package internal

var (
	// FlagJSONUseSnake export json with snake(default camel case)
	FlagJSONUseSnake bool
)

// struct scene
const (
	SceneGorm int = iota
	SceneCommon
)

// mysql type to golang type
var (
	typeMap = map[string]string{
		"int":                "int",
		"integer":            "int",
		"tinyint":            "int8",
		"smallint":           "int16",
		"mediumint":          "int32",
		"bigint":             "int64",
		"int unsigned":       "uint",
		"integer unsigned":   "uint",
		"tinyint unsigned":   "uint8",
		"smallint unsigned":  "uint16",
		"mediumint unsigned": "uint32",
		"bigint unsigned":    "uint64",
		"bit":                "byte",
		"bool":               "bool",
		"enum":               "string",
		"set":                "string",
		"varchar":            "string",
		"char":               "string",
		"tinytext":           "string",
		"mediumtext":         "string",
		"text":               "string",
		"longtext":           "string",
		"blob":               "string",
		"tinyblob":           "string",
		"mediumblob":         "string",
		"longblob":           "string",
		"date":               "time.Time",
		"datetime":           "time.Time",
		"timestamp":          "time.Time",
		"time":               "time.Time",
		"float":              "float64",
		"double":             "float64",
		"decimal":            "float64",
		"binary":             "string",
		"varbinary":          "string",
	}
)
