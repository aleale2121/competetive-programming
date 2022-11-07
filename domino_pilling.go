package main
     
import (
	"bufio"
	"fmt"
	"os"
)
 
func maxDomino(m, n int) int {
	return n*(m/2) + (m%2)*(n/2)
}


 
var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }
 
func main() {
	defer writer.Flush()
 
	var a, b int
	scanf("%d %d\n", &a, &b)
	printf("%d\n", maxDomino(a,b))
}