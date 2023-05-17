package migration

import (
	"fmt"
	"github.com/vikash/projectx/generator/config"
	"os"
	"text/template"
)

func Create(d config.Domain, folderName string) error {
	// Parse migration templates
	tmpl := template.Must(template.New("").Funcs(config.FuncMap).ParseFS(templateFS, "*.tmpl"))

	migrationList := map[int]string{}

	for i, entity := range d.Entities {
		// Create the file
		migrationName := fmt.Sprintf("M%d", 20230517000000+i)
		fileName := fmt.Sprintf("%s/%s_initial_%s.go", folderName, migrationName, entity.Name)
		file, err := os.Create(fileName)
		if err != nil {
			return fmt.Errorf("can not create %s. Error: %s", fileName, err.Error())
		}
		defer file.Close()

		err = tmpl.ExecuteTemplate(file, "migration_1.tmpl", map[string]interface{}{
			"entity":        entity,
			"migrationName": migrationName,
			"createQuery":   createQueryForEntity(entity),
			"deleteQuery":   deleteQueryForEntity(entity),
		})
		if err != nil {
			return fmt.Errorf("can not parse migration_1.tmpl. Error: %s", err.Error())
		}

		migrationList[20230517000000+i] = migrationName
	}

	// Create common migration file
	file, err := os.Create(folderName + "/000_all.go")
	if err != nil {
		return fmt.Errorf("can not create %s. Error: %s", file.Name(), err.Error())
	}
	defer file.Close()
	err = tmpl.ExecuteTemplate(file, "migration_0.tmpl", migrationList)
	if err != nil {
		return fmt.Errorf("can not parse migration_0.tmpl. Error: %s", err.Error())
	}

	return nil
}

func createQueryForEntity(e config.Entity) string {
	return ""
}

func deleteQueryForEntity(e config.Entity) string {
	return ""
}
