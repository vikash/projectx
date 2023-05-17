package store

import (
	"fmt"
	"github.com/vikash/projectx/generator/config"
	"os"
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
	err = tmpl.ExecuteTemplate(file, "store.tmpl", e)
	if err != nil {
		return fmt.Errorf("can not parse store.tmpl. Error: %s", err.Error())
	}

	return nil
}
