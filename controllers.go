package main

func testRedis() {
	println(rdxGet("ac"))
	println(rdxSet("ac", "{Name: 'ac', Age: 25}"))
}