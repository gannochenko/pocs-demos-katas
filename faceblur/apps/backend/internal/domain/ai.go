package domain

type Coordinate struct {
	X int32
	Y int32
}

type FaceDetection struct {
	TopLeft *Coordinate
	BottomRight *Coordinate
}
