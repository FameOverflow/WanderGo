package main

import (
	dbf "SparkForge/Database"
	l "SparkForge/Log"
	svr "SparkForge/Server"
)

func main() {
	l.InitLogging()
	dbf.ConnectToDb()
	svr.Start()
}
