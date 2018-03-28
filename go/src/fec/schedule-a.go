package fec

import (
	"strings"
)

type Contributor struct {
	CounterParty
	Employer string `json:"employer,omitempty" fec:"skip"`
	Occupation string `json:"occupation,omitempty" fec:"skip"`
}

type Contribution struct {
	Date *Date `json:"date,omitempty"`
	Amount float64 `json:"amount"`
	Aggregate float64 `json:"aggregate"`
	Purpose string `json:"purpose,omitempty"`
}

type ScheduleA struct {
	Schedule
	Ref *ReferenceScheduleTransaction `json:"ref"`
	Contributor *Contributor `json:"contributor,omitempty"`
	Election *SimpleElection `json:"election,omitempty"`
	Contribution *Contribution `json:"contribution,omitempty"`
	Donor *FECEntity `json:"donor,omitempty"`
	Memo *Memo `json:"memo,omitempty"`
	AccountRef string `json:"account_ref,omitempty"`
}

func (s *ScheduleA) Link(report *Report) {
	if s.Ref != nil {
		s.Ref.Link(report)
	}
}

func (s *ScheduleA) Unlink(report *Report) {
	if s.Ref != nil {
		s.Ref.Unlink(report)
	}
}

func (s *ScheduleA) MarshalFECField(fields *[]string) error {
	Marshal(s.Type, fields)
	Marshal(s.FilerId, fields)
	Marshal(s.TransactionId, fields)
	Marshal(s.RefTransactionId, fields)
	Marshal(s.RefScheduleName, fields)
	Marshal(s.Contributor, fields)
	Marshal(s.Election, fields)
	Marshal(s.Contribution, fields)
	if s.Contributor != nil {
		Marshal(s.Contributor.Employer, fields)
		Marshal(s.Contributor.Occupation, fields)
	} else {
		Marshal("", fields)
		Marshal("", fields)
	}
	Marshal(s.Donor, fields)
	Marshal(s.Memo, fields)
	Marshal(s.AccountRef, fields)
	return nil
}

func (s *ScheduleA) UnmarshalFECField(fields *[]string) error {
	s.Contributor = &Contributor{}
	s.Election = &SimpleElection{}
	s.Contribution = &Contribution{}
	s.Donor = &FECEntity{}
	s.Memo = &Memo{}
	Unmarshal(&s.Type, fields)
	Unmarshal(&s.FilerId, fields)
	Unmarshal(&s.TransactionId, fields)
	Unmarshal(&s.RefTransactionId, fields)
	Unmarshal(&s.RefScheduleName, fields)
	Unmarshal(s.Contributor, fields)
	Unmarshal(s.Election, fields)
	Unmarshal(s.Contribution, fields)
	if s.Contributor != nil {
		Unmarshal(&s.Contributor.Employer, fields)
		Unmarshal(&s.Contributor.Occupation, fields)
	} else {
		Unmarshal(new(string), fields)
		Unmarshal(new(string), fields)
	}
	Unmarshal(s.Donor, fields)
	Unmarshal(s.Memo, fields)
	Unmarshal(&s.AccountRef, fields)
	return nil
}

/*
func (s *ScheduleA) Fields() []string {
	data := DumpFECFields(s.Type, s.FilerId, s.TransactionId, s.RefTransactionId, s.RefScheduleName, s.Contributor, s.Election, s.Contribution)
	if s.Contributor == nil {
		data = append(data, "", "")
	} else {
		data = append(data, DumpFECFields(s.Contributor.Employer, s.Contributor.Occupation)...)
	}
	data = append(data, DumpFECFields(s.Donor, s.Memo, s.AccountRef)...)
	return data
}

func (s *ScheduleA) Consume(fields []string) ([]string, error) {
	orig_fields := fields
	var err error
	if len(fields) < 5 {
		return orig_fields, IncompleteFieldError
	}
	stype := fields[0]
	filer := fields[1]
	trans := fields[2]
	refTrans := fields[3]
	refSched := fields[4]
	fields = fields[5:]
	cont := &Contributor{}
	fields, err = cont.Consume(fields)
	if err != nil {
		return orig_fields, err
	}
	elec := &SimpleElection{}
	fields, err = elec.Consume(fields)
	if err != nil {
		return orig_fields, err
	}
	contrib := &Contribution{}
	fields, err = contrib.Consume(fields)
	if err != nil {
		return orig_fields, err
	}
	if len(fields) < 2 {
		return orig_fields, IncompleteFieldError
	}
	cont.Employer = fields[0]
	cont.Occupation = fields[1]
	fields = fields[2:]
	donor := &FECEntity{}
	fields, err = donor.Consume(fields)
	if err != nil {
		return orig_fields, err
	}
	memo := &Memo{}
	fields, err = memo.Consume(fields)
	if err != nil {
		return orig_fields, err
	}
	if len(fields) < 1 {
		return orig_fields, IncompleteFieldError
	}
	s.Type = stype
	s.FilerId = filer
	s.TransactionId = trans
	s.RefTransactionId = refTrans
	s.RefScheduleName = refSched
	s.Contributor = cont
	s.Election = elec
	s.Contribution = contrib
	s.Donor = donor
	s.Memo = memo
	s.AccountRef = fields[0]
	return fields[1:], nil
}
*/
func (s *ScheduleA) Record() []byte {
	fields := []string{}
	Marshal(s, &fields)
	return []byte(strings.Join(fields, FIELD_SEP) + RECORD_SEP)
}

func (s *ScheduleA) Name() string {
	return "SA"
}

