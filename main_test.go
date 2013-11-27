package holiday

import (
	"log"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	log.Println("start")
	for day, d := 2, NewDate(2013, 1, 1); d.Year() == 2013; d, day = NewDate(2013, 1, day), day+1 {
		if isHoliday, holiday_name := d.Holiday(); isHoliday {
			log.Printf("%s is %s\n", d, holiday_name)
		}
	}

	t1 := time.Date(2013, 1, 1, 23, 0, 0, 0, time.UTC)
	d := TimeToDate(t1)
	log.Println("t:", t1)
	log.Println("d:", d)

	log.Println("end")
}
