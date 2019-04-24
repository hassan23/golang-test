package main
import (
        "fmt"
        "github.com/Pallinder/go-randomdata"
	"github.com/DataDog/datadog-go"
)
func main(){
        fmt.Println("Running the TestApp")
        fmt.Println(randomdata.SillyName())
}
