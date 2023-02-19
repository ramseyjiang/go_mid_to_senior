package trafficlight

type Vehicle interface {
	coming() string
	passed() string
	passing() string
}

type Train struct {
	mediator Mediator
}

func (t *Train) coming() string {
	if !t.mediator.canPass() {
		return "Train: Coming, cars cannot through, please wait"
	}
	t.mediator.noticeTrainComing()
	t.mediator.setCarPass(false)
	return "Train: is coming, cars wait."
}

func (t *Train) passed() string {
	t.mediator.setCarPass(true)
	return "Train: Left"
}

func (t *Train) passing() string {
	t.mediator.setCarPass(false)
	return "Train: No car, train is passing"
}

type Cars struct {
	mediator Mediator
}

func (c *Cars) coming() string {
	if !c.mediator.canPass() {
		return "Cars: Please wait, train is coming"
	}
	return "Cars: No train, cars can pass"
}

func (c *Cars) passed() string {
	c.mediator.noticeTrainComing()
	return "All car passed"
}

func (c *Cars) passing() string {
	c.coming()
	return "Cars: Allowed pass"
}
