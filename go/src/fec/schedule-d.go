package fec

type ScheduleD {
	Schedule
	Creditor *CounterParty `json:"creditor"`
	Purpose string `json:"purpose"`
	BeginningBalance float64 `json:"beginning_balance"`
	IncurredAmount float64 `json:"incurred_amount"`
	PaymentAmount float64 `json:"payment_amount"`
	EndingBalance float64 `json:"ending_balance"`
}
