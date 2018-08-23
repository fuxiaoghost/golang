package main
import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var step = flag.String("s", " ", "separator")
func main() {
	flag.Parse();
	fmt.Print(strings.Join(flag.Args(), *step));
	if !*n {
		fmt.Println();
	}
}