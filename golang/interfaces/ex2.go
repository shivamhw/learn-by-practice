package main

import (
	"fmt"

	"github.com/shivamhw/golang/interfaces/pkg/store"
)

func ex2(st store.Store) {
	switch v := st.(type){
	case *store.MockStore:
		fmt.Print("Type is *store.MockStore\n")
		v.GetId()
	}
}