var Cfg *ini.File

func init() {
	var err error
	Cfg, err = macaron.SetConfig("conf/app.ini")
	if err != nil {
		panic(err)
	}
}

// Usando o valor
Cfg.Section("").Key("db_type").Value()
