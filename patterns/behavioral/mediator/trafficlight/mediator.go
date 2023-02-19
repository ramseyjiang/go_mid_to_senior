package trafficlight

type Mediator interface {
	canPass() bool
	noticeTrainComing()
	setCarPass(flag bool)
}

type LightsManager struct {
	carPass bool
}

func (s *LightsManager) canPass() bool {
	if s.carPass {
		s.carPass = false
		return true
	}

	return false
}

func (s *LightsManager) noticeTrainComing() {
	if s.carPass {
		s.carPass = false
	}
}

func (s *LightsManager) setCarPass(flag bool) {
	s.carPass = flag
}
