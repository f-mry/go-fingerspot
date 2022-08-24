package fingerspot

type SwitchDemux struct {
	All            func(message interface{})
	RealtimeAttlog func(message *RealtimeAttLogMessage)
	GetUserInfo    func(message *GetUserInfoMessage)
	Other          func(message interface{})
}

func NewSwitchDemux() SwitchDemux {
	return SwitchDemux{
		All:            func(message interface{}) {},
		RealtimeAttlog: func(message *RealtimeAttLogMessage) {},
		GetUserInfo:    func(message *GetUserInfoMessage) {},
		Other:          func(message interface{}) {},
	}
}

func (s SwitchDemux) Handle(message interface{}) {
	s.All(message)
	switch msg := message.(type) {
	case *RealtimeAttLogMessage:
		s.RealtimeAttlog(msg)
	case *GetUserInfoMessage:
		s.GetUserInfo(msg)
	default:
		s.Other(msg)
	}

}
