package main

import "github.com/shivamhw/golang/interfaces/pkg/store"



func ex1(st store.Store) {
	st.Read("fileEX1")
	st.(*store.MockStore).ChangeId("newID")
}
