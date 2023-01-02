package main

/**
    @date: 2023/1/2
**/

type GameObject struct {
	width  int
	height int
	x      float64
	y      float64
}

func (gameObj *GameObject) Width() int {
	return gameObj.width
}

func (gameObj *GameObject) Height() int {
	return gameObj.height
}

func (gameObj *GameObject) X() float64 {
	return gameObj.x
}

func (gameObj *GameObject) Y() float64 {
	return gameObj.y
}
// 然后定义一个接口Entity：

type Entity interface {
	Width() int
	Height() int
	X() float64
	Y() float64
}