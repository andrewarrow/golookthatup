package client

//import "github.com/dontpanic92/wxGo/wx"
import "fmt"
import "time"

func (f *TheFrame) checkQueryEvery() {
	for {
		time.Sleep(time.Second * 1)
		val := ui_query.GetValue()
		fmt.Println(val)
	}
}
