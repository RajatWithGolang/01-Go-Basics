package main
import (
	"fmt"
	s "strings"
)
var p1 = fmt.Println
func main(){

	p1("Len: ", len("hello Rajat"))
	p1("Char:", "hello"[1])
	 p1("ToLower:   ", s.ToLower("TEST"))
    p1("ToUpper:   ", s.ToUpper("test"))
	p1("Contains:  ", s.Contains("test", "es"))
    p1("Count:     ", s.Count("test", "t"))
    p1("HasPrefix: ", s.HasPrefix("test", "te"))
    p1("HasSuffix: ", s.HasSuffix("test", "st"))
    p1("Index:     ", s.Index("test", "e"))
    p1("Repeat:    ", s.Repeat("a", 5))
    p1("Replace:   ", s.Replace("foo", "o", "0", -1))
    p1("Replace:   ", s.Replace("foo", "o", "0", 1))
    p1("Split:     ", s.Split("a-b-c-d-e", "-"))
   
	
}