package martinipkg

import "github.com/go-martini/martini"

// martini is a go framework page, which includes router, logger, handler, recovery and static in it.

func Trigger() {
	m := martini.Classic()

	// Access http://localhost:3000, it will output "Hello world."
	m.Get("/", func() string {
		return "Hello world!"
	})

	// Access http://localhost:3000/test, it will output "Hello test"
	m.Get("/:name", func(params martini.Params) string {
		return "Hello " + params["name"]
	})

	m.Run()
}
