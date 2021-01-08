package main

import (
	"osmantheus/petals"
)

func main() {
	//todo args选择检测项

	//非running的pod
	petals.ListNotRunngPod()
	//pod不满足调协的controller
	petals.ListEoughController()
	//测试节点之间的网络互通
	petals.TestNodeConnect()

	//todo 测试节点的io
	//todo 测试节点的cpu
	//todo 测试节点的缓存

	//todo 测试所有节点的pv读写是否正常
}
