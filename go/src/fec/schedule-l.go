package fec

type ScheduleLReceiptsFromPersons struct {
	Itemized float64
	Unitemized float64
	XXXTotal float64
}

func (x *ScheduleLReceiptsFromPersons) Total() float64 {
	return x.Itemized + x.Unitemized
}

type ScheduleLReceipts struct {
	ReceiptsFromPersions *ScheduleLReceiptsFromPersons
	Other float64
	XXXTotal float64
}

func (x *ScheduleLReceipts) Total() float64 {
	t := float64(0.0)
	if x.ReceiptsFromPersions != nil {
		t += x.ReceiptsFromPersions.Total()
	}
	t += x.Other
	return t
}

type CategorizedScheduleLDisbursements struct {
	VoterRegistration float64
	VoterId float64
	GOTV float64
	GenericCampaign float64
	XXXTotal float64
}

func (x *CategorizedScheduleLDisbursements) Total() float64 {
	return x.VoterRegistration + x.VoterId + x.GOTV + x.GenericCampaign
}

type ScheduleLDisbursements struct {
	Categorized *CategorizedScheduleLDisbursements
	Other float64
	XXXTotal float64
}

func (x *ScheduleLDisbursements) Total() float64 {
	t := float64(0.0)
	if x.Categorized != nil {
		t += x.Categorized.Total()
	}
	t += x.Other
	return t
}

type ScheduleLColumn struct {
	Receipts *ScheduleLReceipts `json:"receipts"`
	Disbursements *ScheduleLDisbursements `json:"disbursements"`
	BeginningCashOnHand float64 `json:"cash_on_hand"`
	XXXTotalReceipts float64
	XXXCashSubtotal float64
	XXXTotalDisbursements float64
	XXXEndingCashOnHand float64
}

func (x *ScheduleLColumn) TotalReceipts() float64 {
	if x.Receipts == nil {
		return 0.0
	}
	return x.Receipts.Total()
}

func (x *ScheduleLColumn) CashSubtotal() float64 {
	return x.BeginningCashOnHand + x.TotalReciepts()
}

func (x *ScheduleLColumn) TotalDisbursements() float64 {
	if x.Disbursements == nil {
		return 0.0
	}
	return x.Disbursements.Total()
}

func (x *ScheduleLColumn) EndingCashOnHand() float64 {
	return x.CashSubtotal() - x.TotalDisbursements()
}

type ScheduleL struct {
	Schedule
	RecordId string `json:"record_id"`
	AccountName string `json:"account_name"`
	CoveragePeriod *Period `json:"coverage_period"`
	ThisPeriod *ScheduleLColumn `json:"this_period"`
	YearToDate *ScheduleLColumn `json:"year_to_date"`
}

