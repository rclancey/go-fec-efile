package fec

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

var UnknownFieldTypeError = errors.New("Unknown field type")
var IncompleteFieldError = errors.New("Incomplete field")
var RecordNotFound = errors.New("Record not found")

type YesNo bool

func (yn YesNo) MarshalFECField(fields *[]string) error {
	v := "N"
	if yn {
		v = "Y"
	}
	*fields = append(*fields, v)
	return nil
}

func (yn *YesNo) UnmarshalFECField(fields *[]string) error {
	if len(*fields) < 1 {
		return IncompleteFieldError
	}
	s := strings.ToLower(strings.TrimSpace((*fields)[0]))
	switch s {
	case "y":
		*yn = YesNo(true)
	case "x":
		*yn = YesNo(true)
	case "t":
		*yn = YesNo(true)
	default:
		*yn = YesNo(false)
	*fields = (*fields)[1:]
	return nil
}

type Date time.Time

func (d *Date) MarshalFECField(fields *[]string) error {
	if d == nil || time.Time(*d).IsZero() {
		*fields = append(*fields, "")
	} else {
		*fields = append(*fields, time.Time(*d).Format("20060102"))
	}
	return nil
}

func (d *Date) UnmarshalFECField(fields *[]string) error {
	if len(*fields) < 1 {
		return IncompleteFieldError
	}
	s := (*fields)[0]
	if strings.TrimSpace(s) == "" {
		var t time.Time
		*d = Date(t)
	} else {
		t, err := time.Parse("20060102", s)
		if err != nil {
			return nil
		}
		*d = Date(t)
	}
	*fields = (*fields)[1:]
	return nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	if time.Time(*d).IsZero() {
		return json.Marshal(nil)
	}
	return json.Marshal(time.Time(*d).Format("20060102"))
}

func (d *Date) UnmarshalJSON(b []byte) error {
	var s *string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if s == nil {
		var t time.Time
		*d = Date(t)
	} else {
		v, err := time.Parse("20060102", *s)
		if err != nil {
			return err
		}
		*d = Date(v)
	}
	return nil
}

type Name struct {
	Last string `json:"last,omitempty"`
	First string `json:"first,omitempty"`
	Middle string `json:"middle,omitempty"`
	Prefix string `json:"prefix,omitempty"`
	Suffix string `json:"suffix,omitempty"`
}

func (n *Name) String() string {
	parts := []string{}
	if n.Prefix != "" {
		parts = append(parts, n.Prefix)
	}
	if n.First != "" {
		parts = append(parts, n.First)
	}
	if n.Middle != "" {
		parts = append(parts, n.Middle)
	}
	if n.Last != "" {
		parts = append(parts, n.Last)
	}
	if n.Suffix != "" {
		parts = append(parts, n.Suffix);
	}
	return strings.Join(parts, " ");
}

type TitledName struct {
	Name
	Title string `json:"title,omitempty"`
}

type Address struct {
	Street []string `json:"street,omitempty"`
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
	ZipCode string `json:"zipcode,omitempty"`
}

func (a *Address) MarshalFECField (fields *[]string) error {
	if a == nil {
		*fields = append(*fields, "", "", "", "", "")
		return nil
	}
	st1 := ""
	st2 := ""
	if len(a.Street) > 0 {
		st1 = a.Street[0]
		if len(a.Street) > 1 {
			st2 = strings.Join(a.Street[1:], " ")
		}
	}
	*fields = append(*fields, st1, st2, a.City, a.State, a.ZipCode)
	return nil
}

func (a *Address) UnmarshalFECField(fields *[]string) error {
	if len(*fields) < 5 {
		return IncompleteFieldError
	}
	xfields := *fields
	a.Street = []string{}
	if xfields[0] != "" || xfields[1] != "" {
		a.Street = append(a.Street, xfields[0])
		if xfields[1] != "" {
			a.Street = append(a.Street, xfields[1])
		}
	}
	a.City = xfields[2]
	a.State = xfields[3]
	a.ZipCode = xfields[4]
	*fields = xfields[5:]
	return nil
}

func (a *Address) String() string {
	lines := []string{}
	if a.Street[0] != "" {
		lines = append(lines, a.Street[0])
		if a.Street[1] != "" {
			lines = append(lines, a.Street[1])
		}
	}
	if a.City != "" {
		line := a.City
		if a.State != "" {
			line += ", " + a.State
			if a.ZipCode != "" {
				line += " " + a.ZipCode
			}
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

type Signature struct {
	Name *Name `json:"name"`
	Date *Date `json:"date"`
}

type TitledSignature struct {
	Name *TitledName `json:"name"`
	Date *Date `json:"date"`
}

type RecordHeader struct {
	Form string `json:"form"`
	FilerId string `json:"filer_id"`
	TransactionId string `json:"transaction_id"`
}

type Schedule struct {
	Header *RecordHeader `json:"header"`
}

func (s *Schedule) Id() string {
	if s.Header == nil {
		return ""
	}
	return s.Header.TransactionId
}

func (s *Schedule) Schedule() string {
	if s.Header == nil {
		return ""
	}
	return s.Header.Form
}

type ReferenceScheduleTransaction struct {
	XXXTransactionId string `json:"transaction_id"`
	XXXSchedule string `json:"schedule"`
	Transaction FECRecord `json:"-" fec:"skip"`
}

func (r *ReferenceScheduleTransaction) TransactionId() string {
	if r.Transaction != nil {
		return r.Transaction.Id()
	}
	return r.XXXTransactionId
}

func (r *ReferenceScheduleTransaction) Schedule() string {
	if r.Transaction != nil {
		return r.Transaction.Schedule()
	}
}

func (r *ReferenceScheduleTransaction) Link(report *Report) {
	if r.Transaction == nil {
		r.Transaction = report.PopRecord(r.XXXTransactionId)
	}
}

func (r *ReferenceScheduleTransaction) Unlink(report *Report) {
	if r.Transaction != nil {
		sch, isa := r.Transaction.(Schedule)
		if isa && sch.Header != nil {
			r.XXXTransactionId = sch.Header.TransactionId
			r.XXXSchedule = sch.Header.Type
			report.Records = append(report.Records, r.Transaction)
			r.Transaction = nil
		}
	}
}

type ReferenceBareTransaction struct {
	XXXTransactionId string `json:"transaction_id"`
	Transaction FECRecord `json:"-" fec:"skip"`
}

func (r *ReferenceBareTransaction) TransactionId() string {
	if r.Transaction != nil {
		return r.Transaction.Id()
	}
	return r.XXXTransactionId
}

func (r *ReferenceBareTransaction) Link(report *Report) {
	if r.Transaction == nil {
		r.Transaction = report.PopRecord(r.XXXTransactionId)
	}
}

func (r *ReferenceBareTransaction) Unlink(report *Report) {
	if r.Transaction != nil {
		sch, isa := r.Transaction.(Schedule)
		if isa && sch.Header != nil {
			r.XXXTransactionId = sch.Header.TransactionId
			report.Records = append(report.Records, r.Transaction)
			r.Transaction = nil
		}
	}
}

