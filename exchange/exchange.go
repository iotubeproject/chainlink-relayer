package exchange

type Exchange interface {
	GetPrice() (uint64, error)
}
