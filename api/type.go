package api

import (
	"math/big"

	"github.com/X-mob/mob-watcher/db"
)

type ShortMob struct {
	Address       string   `json:"address"`
	ImageUrl      string   `json:"image_url"`
	RaiseTarget   *big.Int `json:"raise_target"`
	RaiseDeadline *big.Int `json:"raise_deadline"`
	MemberCount   uint8    `json:"member_count"`
	Status        uint8    `json:"status"`
}

func DBToAPIShortMob(m db.Mob) ShortMob {
	var mob ShortMob
	mob = ShortMob{
		Address:       m.Address.Hex(),
		ImageUrl:      m.Asset.ImageUrl,
		RaiseTarget:   m.RaiseTarget,
		RaiseDeadline: m.RaiseDeadline,
		MemberCount:   0,
		Status:        uint8(m.Status),
	}
	return mob
}

// todo: members
type Mob struct {
	Info    db.Mob   `json:"info"`
	Members []string `json:"members"`
}
