package fec

type H1Fixed struct {
	Presidential bool `json:"presidential"`
	Senate bool `json:"senate"`
	Neither bool `json:"neither"`
}

func (f *H1Fixed) MarshalFECField(fields *[]string) error {
	if f.Presidential && f.Senate {
		*fields = append(*fields, "", "X", "", "")
	} else if f.Presidential {
		*fields = append(*fields, "X", "", "", "")
	} else if f.Senate {
		*fields = append(*fields, "", "", "X", "")
	} else if f.Neither {
		*fields = append(*fields, "", "", "", "X")
	} else {
		*fields = append(*fields, "", "", "", "")
	}
	return nil
}

func (f *H1Fixed) UnmarshalFECField(fields *[]string) error {
	if len(*fields) < 4 {
		return IncompleteFieldError
	}
	data = (*fields)[:4]
	if data[0] != "" {
		f.Presidential = true
		f.Senate = false
		f.Neither = false
	} else if data[1] != "" {
		f.Presidential = true
		f.Senate = true
		f.Neither = false
	} else if data[2] != "" {
		f.Presidential = false
		f.Senate = true
		f.Neither = false
	} else if data[3] != "" {
		f.Presidential = false
		f.Senate = false
		f.Neither = true
	} else {
		f.Presidential = false
		f.Senate = false
		f.Neither = false
	}
	*fields = (*fields)[4:]
	return nil
}

type H1Segregated struct {
	FederalPercent *float64 `json:"federal_percent"`
	NonFederalPercent *float64 `json:"non_federal_percent"`
	Administrative bool `json:"administrative"`
	VoterDrive bool `json:"voter_drive"`
	PublicCommunications bool `json:"public_communication"`
}

type ScheduleH1 {
	Schedule
	Fixed *H1Fixed `json:"fixed"`
	Segregated *H1Segregated `json:"segregated"`
}

func (s *ScheduleH1) Validate() error {
	if s.Segregated != nil && (s.Segregated.FederalPercent != nil || s.Segregated.NonFederalPercent != nil) {
		s.Fixed = nil
	} else if s.Fixed != nil && (s.Presidential || s.Senate || s.Neither) {
		s.Segregated = nil
	}
	return nil
}

type ScheduleH2 struct {
	Schedule
	Activity string `json:"activity"`
	DirectFundraising *YesNo `json:"direct_fundraising"`
	DirectCandidateSupport *YesNo `json:"direct_candidate_support"`
	RatioCode string `json:"ratio_code"`
	FederalPercentage float64 `json:"federal_percentage"`
	NonFederalPercentage float64 `json:"non_federal_percentage"`
}

type ScheduleH3 struct {
	Schedule
	Ref *ReferenceBareTransaction `json:"ref"`
	AccountName string `json:"account_name"`
	EventType string `json:"event_type"`
	Activity string `json:"activity"`
	ReceiptDate *Date `json:"receipt_date"`
	TotalTransfers float64 `json:"total_transfers"`
	Amount float64 `json:"amount"`
}

func (s *ScheduleH3) Link(report *Report) {
	if s.Ref != nil {
		s.Ref.Link(report)
	}
}

func (s *ScheduleH3) Unlink(report *Report) {
	if s.Ref != nil {
		s.Ref.Unlink(report)
	}
}

type ScheduleH4 struct {
	Schedule
	Ref *ReferenceScheduleTransaction `json:"ref"`
	Payee *CounterParty `json:"payee"`
	Account string `json:"account"`
	Date *Date `json:"date"`
	XXXTotal float64 `json:"total"`
	FederalAmount float64 `json:"federal_amount"`
	NonFederalAmount float64 `json:"non_federal_amount"`
	YearToDateTotal float64 `json:"year_to_date_total"`
	Purpose string `json:"purpose"`
	Category string `json:"category"`
	AdministrativeOnly bool `json:"administrative_only"`
	DirectFundraising bool `json:"direct_fundraising"`
	ExemptActivity bool `json:"exempt_activity"`
	VoterDrive bool `json:"voter_drive"`
	DirectCandidateSupport bool `json:"direct_candidate_support"`
	PublicCommunications bool `json:"public_communications"`
	Memo *Memo `json:"memo"`
}

func (s *ScheduleH4) Link(report *Report) {
	if s.Ref != nil {
		s.Ref.Link(report)
	}
}

func (s *ScheduleH4) Unlink(report *Report) {
	if s.Ref != nil {
		s.Ref.Unlink(report)
	}
}

func (s *ScheduleH4) Total() float64 {
	return s.FederalAmount + s.NonFederalAmount
}

type ScheduleH5 struct {
	Schedule
	AccountName string `json:"account_name"`
	Date *Date `json:"date"`
	XXXTotal float64 `json:"amount"`
	VoterRegistration float64 `json:"voter_registration"`
	VoterId float64 `json:"voter_id"`
	GOTV float64 `json:"gotv"`
	GenericCampaign float64 `json:"generic_campaign"`
}

func (s *ScheduleH5) Total() float64 {
	return s.VoterRegistration + s.VoterId + s.GOTV + s.GenericCampaign
}

type ScheduleH6 struct {
	Schedule
	Ref *ReferenceScheduleTransaction `json:"ref"`
	Payee *CounterParty `json:"payee"`
	Account string `json:"account"`
	Date *Date `json:"date"`
	XXXTotal float64 `json:"total"`
	FederalAmount float64 `json:"federal_amount"`
	LevinAmount float64 `json:"levin_amount"`
	YearToDateTotal float64 `json:"year_to_date_total"`
	Purpose string `json:"purpose"`
	Category string `json:"category"`
	VoterRegistration bool `json:"voter_registration"`
	GOTV bool `json:"gotv"`
	VoterId bool `json:"voter_id"`
	GenericCampaign bool `json:"generic_campaign"`
	Memo *Memo `json:"memo"`
}

func (s *ScheduleH6) Link(report *Report) {
	if s.Ref != nil {
		s.Ref.Link(report)
	}
}

func (s *ScheduleH6) Unlink(report *Report) {
	if s.Ref != nil {
		s.Ref.Unlink(report)
	}
}

func (s *ScheduleH6) Total() float64 {
	return s.FederalAmount + s.LevinAmount
}

