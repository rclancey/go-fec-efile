package fec


type ScheduleE {
	Schedule
	Ref *ReferenceScheduleTransaction `json:"ref"`
	Payee *CounterParty `json:"payee"`
	Election *SimpleElection `json:"election"`
	DisseminationDate *Date `json:"dissemination_date"`
	Amount float64 `json:"amount"`
	DisbursementDate *Date `json:"disbursement_date"`
	YearToDateAmount float64 `json:"year_to_date_amount"`
	Purpose string `json:"purpose"`
	Category string `json:"category"`
	PayeeFECId string `json:"payee_fec_id"`
	SupportOppose string `json:"support_oppose"`
	Candidate *SwapCandidate string `json:"candidate"`
	Signature *Signature `json:"signature"`
	Memo *Memo `json:"memo"`
}

func (s *ScheduleE) Link(report *Report) {
	if s.Ref != nil {
		s.Ref.Link(report)
	}
}

func (s *ScheduleE) Unlink(report *Report) {
	if s.Ref != nil {
		s.Ref.Unlink(report)
	}
}

