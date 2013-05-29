package cocoa

type CCSize struct{
	width float32
	height float32
}

func (s* CCSize) Init() {
	s.width = 0.0
	s.height = 0.0
}

func (s* CCSize) InitWithXY(x float32, y float32) {
	s.width = x
	s.height = y
}

func (s* CCSize) InitWithPoint(other CCSize) {
	s.width = other.width
	s.height = other.height
}

