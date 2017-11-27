package common

func StartUp() {
	initConfig()
	initKeys()
	initLogger()
	createDbSession()
	addIndexes()
}
