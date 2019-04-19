package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Bacon ipsum dolor amet porchetta short ribs short loin, spare ribs t-bone kielbasa bresaola\n")
	fmt.Fprint(w, "tail ribeye pastrami flank doner. Turducken shankle kevin, landjaeger rump bresaola \n")
	// Don't forget to flush!
	w.Flush()
}
