package main

type Serializer interface {
	Serialize(any) ([]byte, error)
}

func main() {
}
