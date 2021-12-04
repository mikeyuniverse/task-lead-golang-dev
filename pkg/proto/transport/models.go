package transport

type Sorting struct {
	SortingType SortingType
	OrderType   OrderType
}

type Pagging struct {
	Limit int32
	Start int32
}

var sortTypes = map[string]int32{
	"NAME":  0,
	"PRICE": 1,
}

var sortOrders = map[string]int32{
	"ASC":  0,
	"DESC": 1,
}

type SortingType struct {
	Name  bool // По алфавиту
	Price bool // По цене
}

func (s *SortingType) GetType() string {
	if s.Name {
		return "NAME"
	} else if s.Price {
		return "PRICE"
	}
	return ""
}

func (s *SortingType) ToPB() int32 {
	sort := s.GetType()
	return sortTypes[sort]
}

type OrderType struct {
	Ascending  bool // По возрастанию
	Descending bool // По убыванию
}

func (o *OrderType) GetType() (string, bool) {
	if o.Ascending {
		return "ASC", true
	} else if o.Descending {
		return "DESC", true
	}
	return "", false
}

func (s *OrderType) ToPB() int32 {
	ord, exists := s.GetType()
	if !exists {
		// Что делать с ошибкой?
		return 0
	}
	return sortOrders[ord]
}
