package main

import (
	"osmantheus/petals"
)

func main() {
	//非running的pod
	petals.ListNotRunngPod()
	//pod不满足调协的controller
	petals.ListEoughController()

}
