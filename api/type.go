package api

import (
	"math/big"

	"github.com/X-mob/mob-watcher/db"
)

type ShortMob struct {
	Address  string `json:"address"`
	ImageUrl string `json:"image_url"`
	Status   uint8  `json:"status"`
}

func DBToAPIShortMob(m db.Mob) ShortMob {
	var mob ShortMob
	mob = ShortMob{
		Address: m.Address.Hex(),
		// todo:	ImageUrl,
		Status: uint8(m.Status),
	}
	return mob
}

type Mob struct {
	Address       string   `json:"address"`
	Token         string   `json:"token"`
	TokenId       string   `json:"token_id"`
	TargetMode    uint8    `json:"target_mode"`
	RaiseTarget   *big.Int `json:"raise_target"`
	RaiseDeadline *big.Int `json:"raise_deadline"`
	Deadline      *big.Int `json:"deadline"`
	ImageUrl      string   `json:"image_url"`
	Creator       string   `json:"creator"`
	RaiseAmount   *big.Int `json:"raise_amount"`
	Status        uint8    `json:"status"`
}

func DBToAPIMob(m db.Mob) Mob {
	var mob Mob
	mob = Mob{
		Address:       m.Address.Hex(),
		Token:         m.Token.Hex(),
		TokenId:       m.TokenId.String(),
		TargetMode:    m.TargetMode,
		RaiseTarget:   m.RaiseTarget,
		RaiseDeadline: m.RaiseDeadline,
		Deadline:      m.Deadline,
		// todo:	ImageUrl,
		Creator:     m.Creator.Hex(),
		RaiseAmount: m.RaisedAmount,
		Status:      uint8(m.Status),
	}
	return mob
}
