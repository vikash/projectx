package generator

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/vikash/projectx/generator/config"
	"github.com/vikash/projectx/generator/handlers"
	"github.com/vikash/projectx/generator/migration"
	"github.com/vikash/projectx/generator/store"
	fs "github.com/vikash/projectx/templates"
	"go/format"
	"os"
	"strings"
	"text/template"
)

func CreateDomainCode(g config.Global, d *config.Domain) error {
	if strings.TrimSpace(d.Name) == "" {
		return errors.New("Domain requires a name. Can not genenerate code.")
	}

	// If there are no entities, no generation
	if len(d.Entities) == 0 {
		return fmt.Errorf("some entities are required in a domain for code generatio. 0 found in %s", d.Name)
	}

	folderName := g.GenFolder + "/" + config.CamelCase(d.Name) + "-service"
	err := ensureGenDirectory(folderName)
	if err != nil {
		return fmt.Errorf("can not create directory %s for code. Error: %s ", folderName, err.Error())
	}

	// Parse the common files into folder
	err = parseTemplateToFolder(folderName, g, d)
	if err != nil {
		return fmt.Errorf("could not generate code in %s. Error: %s", folderName, err.Error())
	}

	// Create entity handlers
	for _, e := range d.Entities {
		if e.Name == "" {
			return errors.New("entity requires a name.")
		}

		e.Fields = append([]config.Field{map[string]string{
			"name": "id",
			"type": "string",
		}}, e.Fields...)

		err := handlers.Create(e, folderName+"/handler", g, *d)
		if err != nil {
			return fmt.Errorf("Could not generate handler for '%s'. Error: %s", e.Name, err.Error())
		}

		err = store.Create(e, folderName+"/store")
		if err != nil {
			return fmt.Errorf("Could not generate handler for '%s'. Error: %s", e.Name, err.Error())
		}
	}

	// Create migrations for all entities
	err = migration.Create(d, folderName+"/migrations")
	if err != nil {
		return fmt.Errorf("could not generate migration for '%s'. Error: %s", d.Name, err.Error())
	}

	return nil
}

func parseTemplateToFolder(folderName string, g config.Global, d *config.Domain) error {

	var templates = map[string]string{
		"main.tmpl":  "main.go",
		"gomod.tmpl": "go.mod",
		"gosum.tmpl": "go.sum",
		"env.tmpl":   "configs/.env",
	}

	tmpl := template.Must(template.New("").Funcs(config.FuncMap).ParseFS(fs.FS, "*.tmpl"))

	for t, f := range templates {
		file, err := os.Create(folderName + "/" + f)
		if err != nil {
			return fmt.Errorf("can not create %s. Error: %s", f, err.Error())
		}
		defer file.Close()

		var buf bytes.Buffer
		err = tmpl.ExecuteTemplate(&buf, t, struct {
			D config.Domain
			G config.Global
		}{*d, g})
		if err != nil {
			return fmt.Errorf("can not parse template '%s'. Error: %s", t, err.Error())
		}

		// Format go source code
		var formattedCode []byte
		if strings.HasSuffix(f, ".go") {
			formattedCode, err = format.Source(buf.Bytes())
			if err != nil {
				fmt.Errorf("Unable to format the go source. Error: %s", err.Error())
			}
		} else {
			formattedCode = buf.Bytes()
		}
		file.Write(formattedCode)
	}

	return nil
}

func ensureGenDirectory(path string) error {
	path = path
	// Main Directory
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	os.Mkdir(path+"/configs", os.ModePerm)
	os.Mkdir(path+"/handler", os.ModePerm)
	os.Mkdir(path+"/store", os.ModePerm)
	os.Mkdir(path+"/migrations", os.ModePerm)

	return nil
}
