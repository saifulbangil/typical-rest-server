package typpostgres

import (
	"database/sql"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

func (m *Module) seedCmd(c *typcore.BuildContext) *cli.Command {
	return &cli.Command{
		Name:   "seed",
		Usage:  "Data seeding",
		Action: c.ActionFunc(m.seed),
	}
}

func (m *Module) seed(cfg Config) (err error) {
	var conn *sql.DB
	if conn, err = sql.Open("postgres", m.dataSource(cfg)); err != nil {
		return
	}
	defer conn.Close()
	files, _ := ioutil.ReadDir(m.SeedSource)
	for _, f := range files {
		sqlFile := m.SeedSource + "/" + f.Name()
		log.Infof("Execute seed '%s'", sqlFile)
		var b []byte
		if b, err = ioutil.ReadFile(sqlFile); err != nil {
			log.Error(err.Error())
			continue
		}
		if _, err = conn.Exec(string(b)); err != nil {
			log.Error(err.Error())
			continue
		}
	}
	return
}
