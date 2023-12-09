package main

import (
	con "SparkForge/configs"
	rou "SparkForge/router"
)

func main() {
	con.InitLogging()
	con.ConnectToDb()
	rou.Start()
}
