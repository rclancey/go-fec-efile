package fec

import (
	"fmt"
	"strings"
)

type Election struct {
	Code string `json:"code,omitempty"`
	Date *Date `json:"date,omitempty"`
	State string `json:"state,omitempty"`
}

type Period struct {
	Start *Date `json:"start,omitempty"`
	End *Date `json:"end,omitempty"`
}

type F3XIndividuals struct {
	Itemized []*ScheduleA `json:"itemized,omitempty" fec:"skip"`
	XXXItemizedTotal float64 `json:"-"`
	Unitemized float64 `json:"unitemized"` // Line 11.a.ii
	XXXTotalValue float64 `json:"-"`
}

// Line 11.a.i
func (f *F3XIndividuals) ItemizedTotal() float64 {
	t := float64(0.0)
	for _, item := range f.Itemized {
		if item.Contribution != nil {
			t += item.Contribution.Amount
		}
	}
	return t
}

// Line 11.a.iii
func (f *F3XIndividuals) Total() float64 {
	return f.ItemizedTotal() + f.Unitemized
}

type F3XContributions struct {
	Individuals *F3XIndividuals `json:"individuals,omitempty"`
	PartyCommittees float64 `json:"party_committees"` // Line 11.b
	PACs float64 `json:"pacs"` // Line 11.c
	XXXTotal float64
}

// Line 11.d
func (f *F3XContributions) Total() float64 {
	t := float64(0.0)
	if f.Individuals != nil {
		t += f.Individuals.Total()
	}
	t += f.PartyCommittees + f.PACs
	return t
}

type F3XTransfers struct {
	NonFederal float64 `json:"non_federal"` // Line 18.a
	LevinFunds float64 `json:"levin_funds"` // Line 18.b
	XXXTotal float64
}

// Line 18.c
func (f *F3XTransfers) Total() float64 {
	return f.NonFederal + f.LevinFunds
}

type F3XReceipts struct {
	Contributions *F3XContributions `json:"contributions,omitempty"`
	AffiliateTransfers float64 `json:"affiliate_transfers"` // Line 12
	Loans float64 `json:"loans"` // Line 13
	LoanPayments float64 `json:"loan_payments"` // Line 14
	Offsets float64 `json:"offsets"` // Line 15
	Refunds float64 `json:"refunds"` // Line 16
	OtherFederal float64 `json:"other_federal"` // Line 17
	Transfers *F3XTransfers `json:"transfers"`
	XXXTotal float64
	XXXFederalTotal float64
}

// Line 19
func (f *F3XReceipts) Total() float64 {
	t := float64(0.0)
	if f.Contributions != nil {
		t += f.Contributions.Total()
	}
	t += f.AffiliateTransfers
	t += f.Loans
	t += f.LoanPayments
	t += f.Offsets
	t += f.Refunds
	t += f.OtherFederal
	if f.Transfers != nil {
		t += f.Transfers.Total()
	}
	return t
}

// Line 20
func (f *F3XReceipts) FederalTotal() float64 {
	t := f.Total()
	if f.Transfers != nil {
		t -= f.Transfers.Total()
	}
	return t
}

type F3XOperatingExpenditures struct {
	AllocatedFederal float64 `json:"allocated_federal"` // Line 21.a.i
	AllocatedNonFederal float64 `json:"allocated_nonfederal"` // Line 21.a.ii
	OtherFederal float64 `json:"other_federal"` // Line 21.b
	XXXTotal float64 // Line 21.c
}

// Line 21.c
func (f *F3XOperatingExpenditures) Total() float64 {
	return f.AllocatedFederal + f.AllocatedNonFederal + f.OtherFederal
}

type F3XRefunds struct {
	Individuals float64 `json:"individuals"` // Line 28.a
	PartyCommittees float64 `json:"party_committees"` // Line 28.b
	PACs float64 `json:"pacs"` // Line 28.c
	XXXTotal float64 // Line 28.d
}

// Line 28.d
func (f F3XRefunds) Total() float64 {
	return f.Individuals + f.PartyCommittees + f.PACs
}

type F3XFederalElection struct {
	AllocatedFederal float64 `json:"allocated_federal"` // Line 30.a.i
	AllocatedLevin float64 `json:"allocated_levin"` // Line 30.a.ii
	PureFederal float64 `json:"pure_federal"` // Line 30.b
	XXXTotal float64 // Line 30.c
}

// Line 30.c
func (f *F3XFederalElection) Total() float64 {
	return f.AllocatedFederal + f.AllocatedLevin + f.PureFederal
}

type F3XDisbursements struct {
	OperatingExpenditures *F3XOperatingExpenditures `json:"operating_expenditures"`
	AffiliateTransfers float64 `json:"affiliate_transfers"` // Line 22
	Contributions float64 `json:"contributions"` // Line 23
	IndependentExpenditures float64 `json:"independent_expenditures"` // Line 24
	CoordinatedExpenditures float64 `json:"coordinated_expenditures"` // Line 25
	LoanPayments float64 `json:"loan_payments"` // Line 26
	Loans float64 `json:"loans"` // Line 27
	Refunds *F3XRefunds `json:"refunds"`
	Other float64 `json:"other"` // Line 29
	FederalElection *F3XFederalElection `json:"federal_election"`
	XXXTotal float64 // Line 31
	XXXFederalTotal float64 // Line 32
}

// Line 31
func (f *F3XDisbursements) Total() float64 {
	t := float64(0.0)
	if f.OperatingExpenditures != nil {
		t += f.OperatingExpenditures.Total()
	}
	t += f.AffiliateTransfers
	t += f.Contributions
	t += f.IndependentExpenditures
	t += f.CoordinatedExpenditures
	t += f.LoanPayments
	t += f.Loans
	if f.Refunds != nil {
		t += f.Refunds.Total()
	}
	t += f.Other
	if f.FederalElection != nil {
		t += f.FederalElection.Total()
	}
	return t
}

// Line 32
func (f *F3XDisbursements) FederalTotal() float64 {
	t := f.Total()
	if f.OperatingExpenditures != nil {
		t -= f.OperatingExpenditures.AllocatedNonFederal
	}
	if f.FederalElection != nil {
		t -= f.FederalElection.AllocatedLevin
	}
	return t
}

type F3XPeriod struct {
	BeginningCashOnHand float64 `json:"cash_on_hand"` // Line 6.b
	XXXTotalReciepts float64 // Line 6.c
	XXXCashSubtotal float64 // Line 6.d
	XXXTotalDisbursements float64 // Line 7
	XXXEndingCashOnHand float64 // Line 8
	DebtsTo float64 `json:"debts_to,omitempty"` // Line 9
	DebtsBy float64 `json:"debts_by,omitempty"` // Line 10 
	Receipts *F3XReceipts `json:"receipts,omitempty"` // Lines 11-20
	Disbursements *F3XDisbursements `json:"disbursements,omitempty"` // Lines 21-32
	XXXTotalContributions float64 // Line 33
	XXXTotalContributionRefunds float64 // Line 34
	XXXNetContributions float64 // Line 35
	XXXTotalFederalOperatingExpenditures float64 // Line 36
	XXXOffsetsToOperatingExpenditures float64 // Line 37
	XXXNetOperatingExpenditures float64 // Line 38
}

// Line 6.c
func (f *F3XPeriod) TotalReceipts() float64 {
	if f.Receipts == nil {
		return 0.0
	}
	return f.Receipts.Total()
}

// Line 6.d
func (f *F3XPeriod) CashSubtotal() float64 {
	return f.BeginningCashOnHand + f.TotalReceipts()
}

// Line 7
func (f *F3XPeriod) TotalDisbursements() float64 {
	if f.Disbursements == nil {
		return 0.0
	}
	return f.Disbursements.Total()
}

// Line 8
func (f *F3XPeriod) EndingCashOnHand() float64 {
	return f.CashSubtotal() - f.TotalDisbursements()
}

// Line 33
func (f *F3XPeriod) TotalContributions() float64 {
	if f.Receipts == nil || f.Receipts.Contributions == nil {
		return 0.0
	}
	return f.Receipts.Contributions.Total()
}

// Line 34
func (f *F3XPeriod) TotalContributionRefunds() float64 {
	if f.Disbursements == nil || f.Disbursements.Refunds == nil {
		return 0.0
	}
	return f.Disbursements.Refunds.Total()
}

// Line 35
func (f *F3XPeriod) NetContributions() float64 {
	return f.TotalContributions() - f.TotalContributionRefunds()
}

// Line 36
func (f *F3XPeriod) TotalFederalOperatingExpenditures() float64 {
	if f.Disbursements == nil || f.Disbursements.OperatingExpenditures == nil {
		return 0.0
	}
	return f.Disbursements.OperatingExpenditures.AllocatedFederal + f.Disbursements.OperatingExpenditures.OtherFederal
}

// Line 37
func (f *F3XPeriod) OffsetsToOperatingExpenditures() float64 {
	if f.Receipts == nil {
		return 0.0
	}
	return f.Receipts.Offsets
}

// Line 38
func (f *F3XPeriod) NetOperatingExpenditures() float64 {
	return f.TotalFederalOperatingExpenditures() - f.OffsetsToOperatingExpenditures()
}

type F3XYear struct {
	BeginningCashOnHand float64 `json:"cash_on_hand"` // Line 6.a
	Year int `json:"year"`
	XXXTotalReciepts float64 // Line 6.c
	XXXCashSubtotal float64 // Line 6.d
	XXXTotalDisbursements float64 // Line 7
	XXXEndingCashOnHand float64 // Line 8
	Receipts *F3XReceipts `json:"receipts,omitempty"` // Lines 11-20
	Disbursements *F3XDisbursements `json:"disbursements,omitempty"` // Lines 21-32
	XXXTotalContributions float64 // Line 33
	XXXTotalContributionRefunds float64 // Line 34
	XXXNetContributions float64 // Line 35
	XXXTotalFederalOperatingExpenditures float64 // Line 36
	XXXOffsetsToOperatingExpenditures float64 // Line 37
	XXXNetOperatingExpenditures float64 // Line 38
}

// Line 6.c
func (f *F3XYear) TotalReceipts() float64 {
	if f.Receipts == nil {
		return 0.0
	}
	return f.Receipts.Total()
}

// Line 6.d
func (f *F3XYear) CashSubtotal() float64 {
	return f.BeginningCashOnHand + f.TotalReceipts()
}

// Line 7
func (f *F3XYear) TotalDisbursements() float64 {
	if f.Disbursements == nil {
		return 0.0
	}
	return f.Disbursements.Total()
}

// Line 8
func (f *F3XYear) EndingCashOnHand() float64 {
	return f.CashSubtotal() - f.TotalDisbursements()
}

// Line 33
func (f *F3XYear) TotalContributions() float64 {
	if f.Receipts == nil || f.Receipts.Contributions == nil {
		return 0.0
	}
	return f.Receipts.Contributions.Total()
}

// Line 34
func (f *F3XYear) TotalContributionRefunds() float64 {
	if f.Disbursements == nil || f.Disbursements.Refunds == nil {
		return 0.0
	}
	return f.Disbursements.Refunds.Total()
}

// Line 35
func (f *F3XYear) NetContributions() float64 {
	return f.TotalContributions() - f.TotalContributionRefunds()
}

// Line 36
func (f *F3XYear) TotalFederalOperatingExpenditures() float64 {
	if f.Disbursements == nil || f.Disbursements.OperatingExpenditures == nil {
		return 0.0
	}
	return f.Disbursements.OperatingExpenditures.AllocatedFederal + f.Disbursements.OperatingExpenditures.OtherFederal
}

// Line 37
func (f *F3XYear) OffsetsToOperatingExpenditures() float64 {
	if f.Receipts == nil {
		return 0.0
	}
	return f.Receipts.Offsets
}

// Line 38
func (f *F3XYear) NetOperatingExpenditures() float64 {
	return f.TotalFederalOperatingExpenditures() - f.OffsetsToOperatingExpenditures()
}

type F3X struct {
	Type string `json:"type"`
	FilerId string `json:"filer_id"`
	CommitteeName string `json:"committee_name,omitempty"`
	ChangeOfAddress bool `json:"change_of_address"`
	Address *Address `json:"address,omitempty"`
	ReportCode string `json:"report_code,omitempty"`
	Election *Election `json:"election,omitempty"`
	Period *Period `json:"period,omitempty"`
	QualifiedCommittee bool `json:"qualified_committee"`
	TreasurerSignature *Signature `json:"treasurer_signature"`
	ThisPeriod *F3XPeriod `json:"this_period"`
	YearToDate *F3XYear `json:"year_to_date"`
}

func (f *F3X) Id() string {
	return ""
}

func (f *F3X) Schedule() string {
	return f.Type
}

func (f *F3X) Record() []byte {
	fields := []string{}
	err := Marshal(f, &fields)
	if err != nil {
		fmt.Println("error marshaling F3X to fields:", err)
	}
	return []byte(strings.Join(fields, FIELD_SEP) + RECORD_SEP)
}

func (f *F3X) Name() string {
	return "F3X"
}

