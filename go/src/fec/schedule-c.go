package fec

import (
	"strings"
)

type ScheduleC struct {
	Schedule
	ReceiptLineNumber string `json:"receipt_line_number"`
	Lender *CounterParty `json:"lender"`
	Election *SimpleElection `json:"election"`
	LoanAmount float64 `json:"loan_amount"`
	LoanPaymentToDate float64 `json:"loan_payment_to_date"`
	LoanBalance float64 `json:"loan_balance"`
	LoanIncurredDate *Date `json:"loan_incurred_date"`
	LoanDueDate string `json:"loan_due_date"`
	LoanInterestRate string `json:"loan_interest_rate"`
	Secured *YesNo `json:"secured"`
	PersonalFunds *YesNo `json:"personal_funds"`
	FECLender *DirectFECEntity `json:"fec_lender"`
	Memo *Memo `json:"memo,omitempty"`
}

type Collateral struct {
	Pledged *YesNo `json:"pledged"`
	Description string `json:"description"`
	Value float64 `json:"value"`
	PerfectedInterest *YesNo `json:"perfected_interest"`
}

type FutureIncomeAccount struct {
	DateEstablished *Date `json:"date_established"`
	Location string `json:"location"`
	Address *Address `json:"address"`
	AuthDate *Date `json:"auth_date"`
}

type FutureIncome struct {
	Pledged *YesNo `json:"pledged"`
	Description string `json:"description"`
	Value float64 `json:"value"`
	Account *FutureIncomeAccount `json:"account"`
}

type ScheduleC1 struct {
	Schedule
	Ref *ReferenceBareTransaction `json:"ref"`
	Type string `json:"type"`
	FilerId string `json:"filer_id"`
	TransactionId string `json:"transaction_id"`
	BackReferenceTransactionId string `json:"back_reference_transaction_id"`
	Lender *CounterParty `json:"lender"`
	Election *SimpleElection `json:"election"`
	LoanAmount float64 `json:"loan_amount"`
	LoanInterestRate string `json:"loan_interest_rate"`
	LoanIncurredDate *Date `json:"loan_incurred_date"`
	LoanDueDate string `json:"loan_due_date"`
	Restructured *YesNo `json:"restructured"`
	OriginalLoanDate *Date `json:"original_loan_date"`
	CreditAmount float64 `json:"credit_amount"`
	TotalBalance float64 `json:"total_balance"`
	OthersLiable *YesNo `json:"others_liable"`
	Collateral *Collateral `json:"collateral"`
	FutureIncome *FutureIncome `json:"future_income"`
	Basis string `json:"basis"`
	TreasurerSignature *Signature `json:"treasurer_signature"`
	AuthorizedRepresentative *TitledSignature `json:"authorized_representative"`
}

func (s *ScheduleC1) Link(report *Report) {
	if s.Ref != nil {
		s.Ref.Link(report)
	}
}

func (s *ScheduleC1) Unlink(report *Report) {
	if s.Ref != nil {
		s.Ref.Unlink(report)
	}
}

type ScheduleC2 struct {
	Schedule
	Ref *ReferenceBareTransaction `json:"ref"`
	Name *Name `json:"name"`
	Address *Address `json:"address"`
	Employer string `json:"employer"`
	Occupation string `json:"occupation"`
	Amount float64 `json:"amount"`
}

func (s *ScheduleC2) Link(report *Report) {
	if s.Ref != nil {
		s.Ref.Link(report)
	}
}

func (s *ScheduleC2) Unlink(report *Report) {
	if s.Ref != nil {
		s.Ref.Unlink(report)
	}
}

