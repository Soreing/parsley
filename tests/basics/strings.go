package basics

import "time"

type StringsColl struct {
	SDat string      `json:"sdat"`
	SSlc []string    `json:"sslc"`
	SPtr *string     `json:"sptr"`
	TDat time.Time   `json:"tdat"`
	TSlc []time.Time `json:"tslc"`
	TPtr *time.Time  `json:"tptr"`
}
