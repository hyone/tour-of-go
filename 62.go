package main

import (
  "code.google.com/p/go-tour/pic"
  "image"
  "image/color"
)

type Image struct {
  width int;
  height int;
}

func (i Image) Bounds() image.Rectangle {
  return image.Rect(0, 0, i.width, i.height)
}

func (i Image) ColorModel() color.Model {
  return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
  v := uint8(x*x + y*y)
  return color.RGBA {v, v, 255, 255}
}


func main () {
  m := Image { width: 500, height: 500 }
  pic.ShowImage(m)
}
