package fec

import (
	"strings"
)

type Expenditure struct {
	Date *Date `json:"date,omitempty"`
	Amount float64 `json:"amount"`
	Refund *float64 `json:"refund"`
	Purpose string `json:"purpose,omitempty"`
	Category string `json:"category,omitempty"`
}

type ScheduleB struct {
	Schedule
	Ref *ReferenceScheduleTransaction `json:"ref"`
	Payee *CounterParty `json:"payee,omitempty"`
	Election *SimpleElection `json:"election,omitempty"`
	Expenditure *Expenditure `json:"contribution,omitempty"`
	Beneficiary *FECEntity `json:"donor,omitempty"`
	Memo *Memo `json:"memo,omitempty"`
	AccountRef string `json:"account_ref,omitempty"`
}

func (s *ScheduleB) Link(report *Report) {
	if s.Ref != nil {
		s.Ref.Link(report)
	}
}

func (s *ScheduleA) Unlink(report *Report) {
	if s.Ref != nil {
		s.Ref.Unlink(report)
	}
}

func (s *ScheduleB) Record() []byte {
	fields := []string{}
	Marshal(s, &fields)
	return []byte(strings.Join(fields, FIELD_SEP) + RECORD_SEP)
}

func (s *ScheduleB) Name() string {
	return "SB"
}

