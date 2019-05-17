package model

import "net"


type Category struct {
	Id uint `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Title     string `gorm:"NOT NULL; UNIQUE" json:"title"`
	Slug	string	`gorm:"NOT NULL UNIQUE" json:"slug"`
	Parent     int `gorm:"NOT NULL" json:"parent"`
}

type ipRange struct {
    start net.IP
    end net.IP
}