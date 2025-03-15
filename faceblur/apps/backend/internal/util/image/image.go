package image

import (
	"fmt"
	"image"
)

type BoundingBox struct {
	Label          string
	Confidence     float32
	X1, Y1, X2, Y2 float32
}

func (b *BoundingBox) String() string {
	return fmt.Sprintf("Object %s (confidence %f): (%f, %f), (%f, %f)",
		b.Label, b.Confidence, b.X1, b.Y1, b.X2, b.Y2)
}

// This loses precision, but recall that the BoundingBox has already been
// scaled up to the original image's dimensions. So, it will only lose
// fractional pixels around the edges.
func (b *BoundingBox) toRect() image.Rectangle {
	return image.Rect(int(b.X1), int(b.Y1), int(b.X2), int(b.Y2)).Canon()
}

// Returns the area of b in pixels, after converting to an image.Rectangle.
func (b *BoundingBox) rectArea() int {
	size := b.toRect().Size()
	return size.X * size.Y
}

func (b *BoundingBox) intersection(other *BoundingBox) float32 {
	r1 := b.toRect()
	r2 := other.toRect()
	intersected := r1.Intersect(r2).Canon().Size()
	return float32(intersected.X * intersected.Y)
}

func (b *BoundingBox) union(other *BoundingBox) float32 {
	intersectArea := b.intersection(other)
	totalArea := float32(b.rectArea() + other.rectArea())
	return totalArea - intersectArea
}

// This won't be entirely precise due to conversion to the integral rectangles
// from the image.Image library, but we're only using it to estimate which
// boxes are overlapping too much, so some imprecision should be OK.
func (b *BoundingBox) iou(other *BoundingBox) float32 {
	return b.intersection(other) / b.union(other)
}
