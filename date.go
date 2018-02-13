package jpholiday

import (
	"fmt"
	"time"
)

// 祝日判定用日付構造体。
// 時間以下のデータは無視する。
// NewDate, TimeToDateでオブジェクト作成することで時間以下のデータをzero-fillして作る。
// Date{time.Time}で作成しないこと。
type Date struct {
	time.Time
}

// 年、月、日からDateを生成する
func NewDate(year int, month time.Month, day int) (d Date) {
	d = Date{time.Date(year, month, day, 0, 0, 0, 0, LOCATION_JP)}
	return
}

// time.TimeからDateを生成する
func TimeToDate(t time.Time) (d Date) {
	tmp := t.In(LOCATION_JP)
	d = NewDate(tmp.Year(), tmp.Month(), tmp.Day())
	return
}

// time.Timeに変換
func (d Date) ToTime() time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, LOCATION_JP)
}

// その曜日がその月で何回目かを返す
func (d Date) NthWeekday() (nth int) {
	day := d.Day()
	nth = ((day - 1) / 7) + 1
	return
}

// 国民の祝日ならtrueと祝日を返す。
// 振替休日と国民の休日はチェックしない。
func (d Date) RealHoliday() (isHoliday bool, holiday NamedHoliday) {
	for holiday, f := range NAMED_HOLIDAYS {
		if f(d) {
			//2019以降は退位のため天皇誕生日変更
			//さらに2019は天皇誕生日は祝日ではない
			if holiday == TENNOTANJOBI && d.Year() > 2019 {
				return false, -1
			}
			if holiday == TENNOTANJOBI2 && d.Year() < 2019 {
				return false, -1
			}

			//緑の日は2008以降
			if holiday == MIDORI && d.Year() < 2007 {
				return false, -1
			}
			return true, holiday
		}
	}
	return false, -1
}

// 振り替え休日ならtrue otherwise false
func (d Date) AlternativeHoliday() (isHoliday bool) {
	yesterday := d.Yesterday()
	for {
		y, _ := yesterday.RealHoliday()
		if y {
			if yesterday.Weekday() == time.Sunday {
				return true
			}
		} else {
			return false
		}
		yesterday = yesterday.Yesterday()
	}
	return false
}

// 国民の休日ならtrue otheriwise false
func (d Date) IsSandwitched() (isHoliday bool) {
	if dok, _ := d.RealHoliday(); dok {
		return false
	}
	yesterday := d.Yesterday()
	tommorow := d.Tommorow()
	yok, _ := yesterday.RealHoliday()
	tok, _ := tommorow.RealHoliday()
	if yok && tok {
		isHoliday = true
	} else {
		isHoliday = false
	}
	return
}

// 祝日ならtrueと祝日を返す。
// 振替休日と国民の休日もチェックする
func (d Date) Holiday() (isHoliday bool, holiday NamedHoliday) {
	if isHoliday, holiday = d.RealHoliday(); isHoliday {
		return
	} else {
		if d.AlternativeHoliday() {
			return true, FURIKAEKYUJITSU
		} else if d.IsSandwitched() {
			return true, KOKUMINNOKYUJITSU
		}
	}
	return false, -1
}

// 1日前のDateを返す
func (d Date) Yesterday() Date {
	return NewDate(d.Year(), d.Month(), d.Day()-1)
}

// 1日後のDateを返す
func (d Date) Tommorow() Date {
	return NewDate(d.Year(), d.Month(), d.Day()+1)
}

// Dateの同一性判定
func (d Date) Equal(another Date) bool {
	if d.Year() == another.Year() && d.Month() == another.Month() && d.Day() == another.Day() {
		return true
	}
	return false
}

func (d Date) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year(), d.Month(), d.Day())
}
