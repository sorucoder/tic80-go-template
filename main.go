package main

import (
	"github.com/sorucoder/tic80"
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
	if tic80.Btn(tic80.BUTTON_UP) {
		y--
	}
	if tic80.Btn(tic80.BUTTON_DOWN) {
		y++
	}
	if tic80.Btn(tic80.BUTTON_LEFT) {
		x--
	}
	if tic80.Btn(tic80.BUTTON_RIGHT) {
		x++
	}

	tic80.Cls(13)
	tic80.Spr(1+t%60/30*2, x, y, tic80.NewSpriteOptions().AddTransparentColor(14).SetScale(3).SetSize(2, 2))
	tic80.Print("HELLO WORLD FROM GO!", 65, 84, nil)
	t++
}
