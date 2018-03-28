package fec

type CommitteeF struct {
	Id string `json:"id"`
}

type DirectFECEntityF struct {
	Committee *CommitteeF `json:"committee"`
	Candidate *Candidate `json:"candidate"`
}

type ScheduleF {
	Schedule
	Ref *ReferenceScheduleTransaction `json:"ref"`

	Designated *YesNo `json:"designated"`
	DesignatingCommittee *Committee `json:"designating_committee"`
	SubordinateCommittee *Committee `json:"subordinate_committee"`
	SubordinateAddress *Address `json:"subordinate_address"`
	Payee *CounterParty `json:"payee"`

	Date *Date `json:"date"`
	Amount float64 `json:"amount"`
	Aggregate float64 `json:"aggregate"`
	Purpose string `json:"purpose"`
	Category string `json:"category"`
	FECPayee *DirectFECEntityF `json:"fec_payee"`
	Memo *Memo `json:"memo"`
}

func (s *ScheduleF) Link(report *Report) {
	if s.Ref != nil {
		s.Ref.Link(report)
	}
}

func (s *ScheduleF) Unlink(report *Report) {
	if s.Ref != nil {
		s.Ref.Unlink(report)
	}
}

