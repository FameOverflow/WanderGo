package main

import (
	rou "SparkForge/Api"
	con "SparkForge/Config"
)

func main() {
	con.InitLogging()
	con.ConnectToDb()
	rou.Start()
}
