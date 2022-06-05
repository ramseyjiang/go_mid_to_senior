package main

import (
	"log"

	"github.com/ramseyjiang/go_mid_to_senior/customizepkgs/goerr"
)

func main() {
	log.Println(ResourceNotFound("1234", "message", "User", nil).Error())
}

// ResourceNotFound error abstraction
func ResourceNotFound(id string, message string, kind string, cause error) goerr.Error {
	data := map[string]interface{}{"kind": kind, "id": id}
	return goerr.NewGoError("ResourceNotFound", message, data, cause).
		SetComponent(goerr.ErrService).SetResponseType(goerr.NotFound)
}
