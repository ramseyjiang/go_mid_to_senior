package clinic

type Step interface {
	execute(*Patient) []string
	setNext(Step)
}

type Patient struct {
	name              string
	bookDone          bool
	doctorCheckUpDone bool
	paymentDone       bool
	pharmacyDone      bool
	record            []string
}
