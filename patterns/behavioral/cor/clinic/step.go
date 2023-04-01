package clinic

type Handler interface {
	execute(*Patient) []string
	setNext(Handler)
}

type Patient struct {
	name              string
	bookDone          bool
	doctorCheckUpDone bool
	paymentDone       bool
	pharmacyDone      bool
	record            []string
}
