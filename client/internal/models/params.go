package models

type Pagging struct {
	Limit int
}

type SortingType struct {
	Alphabet bool
}

type OrderType struct {
	Ascending  bool // По возрастанию
	Descending bool // По убыванию
}

type Sorting struct {
	SortingType SortingType
	OrderType   OrderType
}
