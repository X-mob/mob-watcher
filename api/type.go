package api

import (
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
		Address:  m.Address.Hex(),
		ImageUrl: m.Asset.ImageUrl,
		Status:   uint8(m.Status),
	}
	return mob
}
