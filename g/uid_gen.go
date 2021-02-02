package g

import (
	"github.com/sony/sonyflake"
	"sync"
	"time"
)

type idGen struct {
	once sync.Once
	sf *sonyflake.Sonyflake
}

var genID idGen

func UIDGen() *sonyflake.Sonyflake {
	genID.once.Do(func() {
		t, _ := time.Parse("2006-01-02 15:04", "2020-05-20 13:14")
		settings := sonyflake.Settings{
			StartTime:      t,
		}
		genID.sf = sonyflake.NewSonyflake(settings)
	})
	return genID.sf
}
