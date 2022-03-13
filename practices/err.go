package practices

import (
	"fmt"

	"golang_learn/customizepkgs/goerr"
)

func TriggerErrorPkg() {
	fmt.Println(ResourceNotFound("1234", "message", "User", nil).Error())
}

// ResourceNotFound error abstraction
func ResourceNotFound(id, message string, kind string, cause error) goerr.Error {
	data := map[string]interface{}{"kind": kind, "id": id}
	return goerr.NewGoError("ResourceNotFound", message, data, cause).
		SetComponent(goerr.ErrService).SetResponseType(goerr.NotFound)
}
