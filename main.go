package main

import (
	"encoding/json"
	"fmt"

	"github.com/ehsanghaffar/github-stats-golang/api"
)

func main() {

	user := api.GetUserData("ehsanghaffar")

	fmt.Println(PrettyPrint(user))

}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
