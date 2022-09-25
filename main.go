package main

import (
	"tic80-go/tic80"
	_ "unsafe"
)

var (
	t int = 0
	x int = 96
	y int = 24
)

//go:export BOOT
func BOOT() {
	tic80.Start()
}

//go:export TIC
func TIC() {
	if tic80.Btn(0) {
		y--
	}
	if tic80.Btn(1) {
		y++
	}
	if tic80.Btn(2) {
		x--
	}
	if tic80.Btn(3) {
		x++
	}

	tic80.Cls(13)
	tic80.Spr(1+t%60/30*2, x, y, tic80.NewSpriteOptions().AddTransparentColor(14).SetScale(3).SetSize(2, 2))
	tic80.Print("HELLO WORLD FROM GO!", 65, 84, nil)
	t++
}
