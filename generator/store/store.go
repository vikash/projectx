package store

import (
	"fmt"
	"github.com/vikash/projectx/generator/config"
	"os"
	"strings"
	"text/template"
)

func Create(e config.Entity, folderName string) error {
	// Store
	file, err := os.Create(folderName + "/" + e.Name + ".go")
	if err != nil {
		return fmt.Errorf("can not create %s. Error: %s", file.Name(), err.Error())
	}
	defer file.Close()

	tmpl := template.Must(template.New("").Funcs(config.FuncMap).ParseFS(templateFile, "store.tmpl"))
	insertQuery, valueString := insertQueryForEntity(e)
	updateQuery, updateValues := updateQueryForEntity(e)
	err = tmpl.ExecuteTemplate(file, "store.tmpl", map[string]interface{}{
		"entity":       e,
		"insertQuery":  insertQuery,
		"valueString":  valueString,
		"updateQuery":  updateQuery,
		"updateValues": updateValues,
	})
	if err != nil {
		return fmt.Errorf("can not parse store.tmpl. Error: %s", err.Error())
	}

	return nil
}

func insertQueryForEntity(e config.Entity) (query string, value string) {
	fieldNames := ""
	questionMarks := ""
	valueString := ""
	for _, f := range e.Fields {
		fieldNames += "`" + config.SnakeCase(f["name"]) + "`, "
		questionMarks += "?, "
		valueString += config.CamelCase(e.Name) + "." + config.PascalCase(f["name"]) + ", "
	}
	fieldNames = strings.TrimRight(fieldNames, ", ")
	questionMarks = strings.TrimRight(questionMarks, ", ")
	valueString = strings.TrimRight(valueString, ", ")

	query = fmt.Sprintf("insert into `%s` (%s) values (%s)", config.SnakeCase(e.Name), fieldNames, questionMarks)
	return query, valueString
}

func updateQueryForEntity(e config.Entity) (query string, value string) {
	queryPart := ""
	for _, f := range e.Fields {
		queryPart += "`" + config.SnakeCase(f["name"]) + "` = ?, "
		value += config.CamelCase(e.Name) + "." + config.PascalCase(f["name"]) + ", "
	}
	queryPart = strings.TrimRight(queryPart, ", ")
	value = strings.TrimRight(value, ", ")
	query = fmt.Sprintf("update %s SET %s where id=?", config.SnakeCase(e.Name), queryPart)

	return query, value
}
