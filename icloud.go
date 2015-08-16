package main
import (
	"github.com/mig2/icloud/engine"
	"github.com/lwsanty/gophotocloud/photos"
)
func main() {
	eng, err := engine.NewEngine("login", "pass")
	if err != nil {
		panic(err)
	}

	err2 := photos.NewP(eng)
	if err2 != nil {
		panic(err2)
	}

}


