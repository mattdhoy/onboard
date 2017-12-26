package onboard

import (
	"github.com/codyoss/nap"
)

func main() {
	app := nap.New()
	app.GET("/")
	app.Serve()
}
