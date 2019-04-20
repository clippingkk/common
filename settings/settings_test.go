package settings

func (suitor *Suitor) TestFetchFile()  {
	err := suitor.loader.fetchFile()
	suitor.NoError(err)
	suitor.Equal("test.dev.json", suitor.loader.file)
}

func (suitor *Suitor)TestLoadSettings()  {
	var cfg *AppConfig

	suitor.loader.container = &AppConfig{}
	err := suitor.loader.loadSettings()
	suitor.NoError(err)

	cfg, ok := suitor.loader.container.(*AppConfig)
	suitor.True(ok)
	suitor.Equal("test-app", cfg.AppName)
	suitor.Equal(true, cfg.Debug)
	suitor.Equal("test", cfg.AppEnv)
	suitor.Equal("root@tcp(localhost:3306)/test", cfg.Database.DSN)
}

func (suitor *Suitor)TestInitialize() {
	cfg := &AppConfig{}
	err := Initialize("test.dev.json", cfg)
	suitor.NoError(err)
	suitor.Equal("test-app", cfg.AppName)
	suitor.Equal(true, cfg.Debug)
	suitor.Equal("test", cfg.AppEnv)
	suitor.Equal("root@tcp(localhost:3306)/test", cfg.Database.DSN)
}