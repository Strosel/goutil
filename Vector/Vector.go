package vector

type Vector2 struct {
	X float64
	Y float64
}

type Vector2I struct {
	X int
	Y int
}

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

type Vector3I struct {
	X int
	Y int
	Z int
}

func (v1 Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

func (v1 Vector2) Sub(v2 Vector2) Vector2 {
	return Vector2{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
}

func (v1 Vector3) Add(v2 Vector3) Vector3 {
	return Vector3{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
}

func (v1 Vector3) Sub(v2 Vector3) Vector3 {
	return Vector3{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
}

func (v1 Vector2I) Add(v2 Vector2I) Vector2I {
	return Vector2I{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

func (v1 Vector2I) Sub(v2 Vector2I) Vector2I {
	return Vector2I{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
}

func (v1 Vector3I) Add(v2 Vector3I) Vector3I {
	return Vector3I{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
}

func (v1 Vector3I) Sub(v2 Vector3I) Vector3I {
	return Vector3I{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
}
