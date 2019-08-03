package object

type Object interface {
	CompareWith(obj interface{}) int
}
