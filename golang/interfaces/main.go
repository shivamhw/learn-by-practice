package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"reflect"

	"github.com/shivamhw/golang/interfaces/pkg/store"
)



func main() {
    var s store.Store
	s = store.NewMockStore(fmt.Sprintf("%d", rand.Int32()%1000))
	ex1(s)
	s.(*store.MockStore).GetId()
	ex2(s)

	// checking type of interface using reflect
	t := reflect.TypeOf(s)
	fmt.Print(t)
	fmt.Print(t.Kind())

	// marshaling interfacejj
	d, err := json.Marshal(s)
	if err != nil {
		fmt.Print("error:", err)
	}
	fmt.Println(string(d))
}