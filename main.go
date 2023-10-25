package main

type Object struct {
	Nil        *string
	Bool       bool
	Str        string
	Int        int
	Int16      int16
	Int32      int32
	Int64      int64
	Float32    float32
	Float64    float64
	ArrStr     []string
	ArrInt     []int
	ArrFloat32 []float32
	Map        map[string]Object
}

func main() {
	// msgpack.Marshal(&Object{})
}
