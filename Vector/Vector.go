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
	Vector2
	Z float64
}

type Vector3I struct {
	Vector2I
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

func (v1 Vector2) Equals(v2 Vector2) bool {
	return v1.X == v2.X && v1.Y == v2.Y
}

func (v1 Vector3) Add(v2 Vector3) Vector3 {
	return Vector3{
		Vector2: v1.Vector2.Add(v2.Vector2),
		Z:       v1.Z + v2.Z,
	}
}

func (v1 Vector3) Sub(v2 Vector3) Vector3 {
	return Vector3{
		Vector2: v1.Vector2.Sub(v2.Vector2),
		Z:       v1.Z - v2.Z,
	}
}

func (v1 Vector3) Equals(v2 Vector3) bool {
	return v1.Vector2.Equals(v2.Vector2) && v1.Z == v2.Z
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

func (v1 Vector2I) Equals(v2 Vector2I) bool {
	return v1.X == v2.X && v1.Y == v2.Y
}

func (v1 Vector3I) Add(v2 Vector3I) Vector3I {
	return Vector3I{
		Vector2I: v1.Vector2I.Add(v2.Vector2I),
		Z:        v1.Z + v2.Z,
	}
}

func (v1 Vector3I) Sub(v2 Vector3I) Vector3I {
	return Vector3I{
		Vector2I: v1.Vector2I.Sub(v2.Vector2I),
		Z:        v1.Z - v2.Z,
	}
}

func (v1 Vector3I) Equals(v2 Vector3I) bool {
	return v1.Vector2I.Equals(v2.Vector2I) && v1.Z == v2.Z
}
