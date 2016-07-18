package common

func StartUp() {
	initConfig()
	initKeys()
	setLogLevel(Level(AppConfig.LogLevel))
	createDbSession()
	addIndexes()
}
