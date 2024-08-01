package models

import (
	"github.com/Cleverse/go-utilities/nullable"
	"github.com/uptrace/bun"
)

type PrimitiveArtist struct {
	bun.BaseModel  `bun:"table:Author_Main"`
	UUID           int8   `bun:"uuid,pk,autoincrement"`
	Author         string `bun:"Author,notnull"`
	Twitter_link   string `bun:"Twitter_link"`
	Facebook_link  string `bun:"Facebook_link"`
	Instagram_link string `bun:"Instagram_link"`
	Plurk_link     string `bun:"Plurk_link"`
	Pixiv_link     string `bun:"Pixiv_link"`
	Baha_link      string `bun:"Baha_link"`
	Youtube_link   string `bun:"Youtube_link"`
	Twitch_link    string `bun:"Twitch_link"`
	Official_link  string `bun:"Official_link"`
	Store_link     string `bun:"Store_link"`
	Myacg_link     string `bun:"Myacg_link"`
	Photo          string `bun:"Photo"`
	Introduction   string `bun:"Introduction"`
	Tags           string `bun:"Tags"`
}

type Pagination struct {
	TotalPage    int                    `json:"total_page"`
	PreviousPage nullable.Nullable[int] `json:"prev_page"`
	NextPage     nullable.Nullable[int] `json:"next_page"`
	CurrentPage  int                    `json:"current_page"`
	TotalRecords int                    `json:"total_records"`
}
