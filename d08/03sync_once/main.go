package main

import (
	"image"
	"sync"
)

// sync.once

var icons map[string]image.Image
var loadIconsOnce sync.Once

func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}
}

func loadIcon(s string) image.Image {
	return Icon("")
}

func Icon(name string) image.Image {
	// Icon 被多个goroutine调用时不是并发安全的
	//if icons == nil {
	//	loadIcons()
	//}

	//Icon 是并发安全的
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}
