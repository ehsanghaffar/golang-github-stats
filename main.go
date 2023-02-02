package main

import (
	// "fmt"
	"log"
	"net/http"

	// "github.com/ehsanghaffar/github-stats-golang/api"
	"github.com/ehsanghaffar/github-stats-golang/charts"
	"github.com/ehsanghaffar/github-stats-golang/utils"
)

func main() {

	// user := api.GetUserData("ehsanghaffar")
	// fmt.Println(utils.PrettyPrint(user))

	example := []charts.Exampler{
		charts.BarType{},
		charts.PieType{},
	}

	for _, e := range example {
		e.Example()
	}

	// Run webserver for serve charts
	fs := http.FileServer(http.Dir("html"))
	log.Println("Server runing at http://localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", utils.LogRequest(fs)))

}
