package fec

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strings"
)

func NewRecord(formType string) (FECRecord, error) {
	switch {
	case formType == "HDR":
		return &HDR{}, nil
	case strings.HasPrefix(formType, "F3X"):
		return &F3X{}, nil
	case strings.HasPrefix(formType, "SA"):
		return &ScheduleA{}, nil
	case strings.HasPrefix(formType, "SB"):
		return &ScheduleB{}, nil
	case strings.HasPrefix(formType, "SC1"):
		return &ScheduleC1{}, nil
	case strings.HasPrefix(formType, "SC2"):
		return &ScheduleC2{}, nil
	case strings.HasPrefix(formType, "SC"):
		return &ScheduleC{}, nil
	case strings.HasPrefix(formType, "SD"):
		return &ScheduleD{}, nil
	case strings.HasPrefix(formType, "SE"):
		return &ScheduleE{}, nil
	case strings.HasPrefix(formType, "SF"):
		return &ScheduleF{}, nil
	case strings.HasPrefix(formType, "H1"):
		return &ScheduleH1{}, nil
	case strings.HasPrefix(formType, "H2"):
		return &ScheduleH2{}, nil
	case strings.HasPrefix(formType, "H3"):
		return &ScheduleH3{}, nil
	case strings.HasPrefix(formType, "H4"):
		return &ScheduleH4{}, nil
	case strings.HasPrefix(formType, "H5"):
		return &ScheduleH5{}, nil
	case strings.HasPrefix(formType, "H6"):
		return &ScheduleH6{}, nil
	case strings.HasPrefix(formType, "SL"):
		return &ScheduleL{}, nil
	}
	return nil, errors.New("Unknown record type " + formType)
}

type Report struct {
	Records []FECRecord `json:"records"`
}

func (r *Report) PopRecord(id string) FECRecord {
	for i, rec := range r.Records {
		sch, isa := rec.(Schedule)
		if isa && sch.Header != nil {
			if sch.Header.TransactionId == id {
				r.Records = append(r.Records[:i], r.Records[i+1:]...)
				return rec
			}
		}
	}
	return nil
}

func (r *Report) GetRecord(name string) FECRecord {
	for _, rec := range r.Records {
		if rec.Name() == name {
			return rec
		}
	}
	return nil
}

func (r *Report) GetRecords(name string) []FECRecord {
	recs := []FECRecord{}
	for _, rec := range r.Records {
		if rec.Name() == name {
			recs = append(recs, rec)
		}
	}
	return recs
}

func (r *Report) Data() ([]byte, error) {
	data := []byte{}
	for i := 0; i < len(r.Records); i++ {
		rec := r.Records[0]
		link, isa := rec.(Linkable)
		if isa {
			rec.Unlink(r)
		}
		data = append(data, rec.Record()...)
	}
	for 
	return data, nil
}

func ReadReport(data []byte) (*Report, error) {
	var err error
	r := &Report{
		Records: []FECRecord{},
	}
	reader := bufio.NewReader(bytes.NewReader(data))
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		recbytes := scanner.Bytes()
		fields := strings.Split(strings.TrimSpace(string(recbytes)), FIELD_SEP)
		fmt.Printf("reading record %s\n", fields[0])
		rec, err := NewRecord(fields[0])
		if err != nil {
			return nil, err
		}
		err = Unmarshal(rec, &fields)
		//fields, err = rec.Consume(fields)
		if err != nil {
			fmt.Println("bad line", strings.Join(fields, "|"))
			return nil, err
		}
		//fmt.Println("record =", rec)
		r.Records = append(r.Records, rec)
	}
	for _, rec := range r.Records {
		link, isa := rec.(Linkable)
		if isa {
			link.Link(r)
		}
	}
	rec := r.GetRecord("F3X")
	if rec != nil {
		f3x, isa := rec.(*F3X)
		if isa && f3x.ThisPeriod != nil && f3x.ThisPeriod.Receipts != nil && f3x.ThisPeriod.Receipts.Contributions != nil && f3x.ThisPeriod.Receipts.Contributions.Individuals != nil {
			scheda := []*ScheduleA{}
			for _, rec := range r.GetRecords("SA") {
				s, isa := rec.(*ScheduleA)
				if isa {
					scheda = append(scheda, s)
				}
			}
			f3x.ThisPeriod.Receipts.Contributions.Individuals.Itemized = scheda
		}
	}
	return r, nil
}

