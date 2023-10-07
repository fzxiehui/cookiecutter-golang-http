package sid

import "github.com/sony/sonyflake"

type Sid struct {
	sf *sonyflake.Sonyflake
}

func NewSid() *Sid {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{})
	if sf == nil {
		panic("sonyflake not created")
	}
	return &Sid{sf}
}

func (s Sid) GenUint64() (uint64, error) {
	return s.sf.NextID()
}
