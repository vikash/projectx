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
	err = tmpl.ExecuteTemplate(file, "store.tmpl", map[string]interface{}{
		"entity":      e,
		"insertQuery": insertQuery,
		"valueString": valueString,
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
