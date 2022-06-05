package importpkg

import (
	"log"

	"github.com/ramseyjiang/go_mid_to_senior/customizepkgs/goerr"
	"github.com/ramseyjiang/go_mid_to_senior/customizepkgs/greet"
)

func TriggerImportPkg() {
	log.Println(ResourceNotFound("1234", "message", "User", nil).Error())
	Greet()
}

func Greet() {
	log.Println(greet.Morning)
	log.Println(greet.SayHi())
}

// ResourceNotFound error abstraction
func ResourceNotFound(id string, message string, kind string, cause error) goerr.Error {
	data := map[string]interface{}{"kind": kind, "id": id}
	return goerr.NewGoError("ResourceNotFound", message, data, cause).
		SetComponent(goerr.ErrService).SetResponseType(goerr.NotFound)
}
