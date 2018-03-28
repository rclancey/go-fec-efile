package fec

type Linkable interface {
	Link(*Report)
	Unlink(*Report)
}

type EntityType string
var EntityType_CAN = EntityType("CAN")
var EntityType_CCM = EntityType("CCM")
var EntityType_COM = EntityType("COM")
var EntityType_IND = EntityType("IND")
var EntityType_ORG = EntityType("ORG")
var EntityType_PAC = EntityType("PAC")
var EntityType_PTY = EntityType("PTY")

type CounterParty struct {
	Type EntityType `json:"type,omitempty"`
	Organization string `json:"organization,omitempty"`
	Name *Name `json:"name,omitempty"`
	Address *Address `json:"address,omitempty"`
}

type Committee struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Candidate struct {
	Id string `json:"id,omitempty"`
	Name *Name `json:"name,omitempty"`
	Office string `json:"office,omitempty"`
	State string `json:"state,omitempty"`
	District *int `json:"district,omitempty"`
}

type SwapCandidate struct {
	Id string `json:"id,omitempty"`
	Name *Name `json:"name,omitempty"`
	Office string `json:"office,omitempty"`
	District *int `json:"district,omitempty"`
	State string `json:"state,omitempty"`
}

type Conduit struct {
	Name string `json:"name,omitempty"`
	Address *Address `json:"address,omitempty"`
}

type DirectFECEntity struct {
	Committee *Committee `json:"committee,omitempty"`
	Candidate *Candidate `json:"candidate,omitempty"`
}

type FECEntity struct {
	DirectFECEntity
	Conduit *Conduit `json:"conduit,omitempty"`
}

type Memo struct {
	Code bool `json:"code,omitempty"`
	Text string `json:"text,omitempty"`
}

type SimpleElection struct {
	Code string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

