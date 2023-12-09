package main

import (
	rou "SparkForge/Router"
	con "SparkForge/Config"
)

func main() {
	con.InitLogging()
	con.ConnectToDb()
	rou.Start()
}
