package fec

import (
	"strings"
)

type Software struct {
	Name string `json:"name"`
	Version string `json:"version,omitempty"`
}

type HDR struct {
	HDR string `json:"-"`
	Type string `json:"type"`
	Version string `json:"version"`
	Software *Software `json:"software,omitempty"`
	ReportId string `json:"report_id,omitempty"`
	ReportNumber *int `json:"report_number,omitempty"`
	Comment string `json:"comment,omitempty" fec:"optional"`
}

func (h *HDR) MarshalFECField(data *[]string) error {
	Marshal(h.HDR, data)
	Marshal(h.Type, data)
	Marshal(h.Version, data)
	Marshal(h.Software, data)
	Marshal(h.ReportId, data)
	Marshal(h.ReportNumber, data)
	if h.Comment != "" {
		Marshal(h.Comment, data)
	}
	return nil
}

func (h *HDR) Record() []byte {
	fields := []string{}
	Marshal(h, &fields)
	return []byte(strings.Join(fields, FIELD_SEP) + RECORD_SEP)
}

func (h *HDR) Name() string {
	return "HDR"
}

func NewHeader() *HDR {
	return &HDR{
		Type: "FEC",
		Version: "8.2",
		Software: &Software{
			Name: "FoothillDemsFEC",
			Version: "0.1",
		},
	}
}

