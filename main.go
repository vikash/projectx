package main

import (
	"github.com/vikash/gofr/pkg/gofr"
	"github.com/vikash/projectx/generator"
	"github.com/vikash/projectx/generator/config"
	"gopkg.in/yaml.v3"

	"os"
)

func main() {
	app := gofr.NewCMD()
	app.SubCommand("", func(c *gofr.Context) (interface{}, error) {
		c.Logger.Debug("Reading Entity file entities.yaml")
		bytes, err := os.ReadFile("entities.yaml")
		if err != nil {
			return nil, err

		}

		c.Logger.Debug("Unmarshalling Configuration")
		conf := config.Config{}
		err = yaml.Unmarshal(bytes, &conf)
		if err != nil {
			return nil, err
		}
		c.Logger.Infof("%d domains found", len(conf.Domains))

		// Set default globals
		if conf.Global.GenFolder == "" {
			conf.Global.GenFolder = "gen"
		}

		for i, d := range conf.Domains {
			c.Logger.Infof("%d. Parsing %s domain", i+1, d.Name)
			err := generator.CreateDomainCode(conf.Global, &d)
			if err != nil {
				c.Logger.Error(err)
			}
		}

		return nil, nil
	})
	app.Run()
}
