package models

// DNSQuery Table Schema
type DNSQuery struct {
	ID        uint      `gorm:"primarykey" json:"queryID"`
	CreatedAt int64     `json:"created_time"`
	ClientIp  string    `json:"client_ip"`
	Domain    string    `json:"domain"`
	Addresses []Address `gorm:"foreignKey:DNSQueryID" json:"addresses"`
}

// Address Table Schema
type Address struct {
	ID         uint   `gorm:"primarykey" json:"queryID"`
	DNSQueryID uint   `json:"-"`
	IP         string `json:"ip"`
}
