package init

type App struct {
	Addr  string
	Debug bool
}

func initAppConf() *App {
	// load conf
	conf := &App{}
	return conf
}
