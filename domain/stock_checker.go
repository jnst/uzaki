package domain

type StockChecker interface {
	CreateURL() string
	CheckStock() (bool, error)
	String() string
}
