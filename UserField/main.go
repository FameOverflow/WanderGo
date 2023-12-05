package main

import (
	dbf "SparkForge/Database"
	l "SparkForge/Log"
	rou "SparkForge/Routes"
)

func main() {
	l.InitLogging()
	dbf.ConnectToDb()
	rou.Start()
}
