package main

import (
	"fmt"
	"encoding/json"
)

type BlockAct struct {
	Id uint32 `json:"id,omitempty"`
	OpenDay uint32 `json:"open_day,omitempty"`
	ActIds []uint32 `json:"act_ids,omitempty"`
}

type BlockActList struct {
	BlockAct []BlockAct `json:"block_act,omitempty"`
	Dels []uint32 `json:"dels,omitempty"`
}

func main() {
	body := `
		{"block_act":[{"id":1,"open_day":2,"act_ids":[3,4]}],
		"dels":[]
	}
	`

	//var b BlockActList
	b := &BlockActList{}
	if err := json.Unmarshal([]byte(body),&b); err != nil {
		fmt.Println("Unmarshal error %v", err)
		return
	}

	fmt.Println(b)
	fmt.Println(b.BlockAct[0].OpenDay)
}

