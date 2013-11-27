package holiday

import (
	"fmt"
	"time"
)

// 祝日表現
type NamedHoliday int

const (
	GANTAN       NamedHoliday = iota
	SEIJIN                    = iota
	KENKOKUKINEN              = iota
	SHUNBUN                   = iota
	SHOWA                     = iota
	KENPOKINEN                = iota
	MIDORI                    = iota
	KODOMO                    = iota
	UMI                       = iota
	KEIRO                     = iota
	SHUBUN                    = iota
	TAIIKU                    = iota
	BUNKA                     = iota
	KINROKANSHA               = iota
	TENNOTANJOBI              = iota
)

// 祝日の名前マップ
var (
	HOLIDAY_NAMES = map[NamedHoliday]string{
		GANTAN:       "元旦",
		SEIJIN:       "成人の日",
		KENKOKUKINEN: "建国記念の日",
		SHUNBUN:      "春分の日",
		SHOWA:        "昭和の日",
		KENPOKINEN:   "憲法記念日",
		MIDORI:       "みどりの日",
		KODOMO:       "こどもの日",
		UMI:          "海の日",
		KEIRO:        "敬老の日",
		SHUBUN:       "秋分の日",
		TAIIKU:       "体育の日",
		BUNKA:        "文化の日",
		KINROKANSHA:  "勤労感謝の日",
		TENNOTANJOBI: "天皇誕生日",
	}
)

func (h NamedHoliday) String() string {
	return HOLIDAY_NAMES[h]
}

//第○×曜日で判定する関数を作成するファクトリ関数
//ハッピーマンデーなどに使用。
func DynamicHolidayCheckerFactory(month time.Month, nth int, weekday time.Weekday) (f func(Date) bool) {
	f = func(d Date) bool {
		if month == d.Month() && weekday == d.Weekday() && nth == d.NthWeekday() {
			return true
		}
		return false
	}
	return
}

// 何月何日で判定する関数を作成するファクトリ関数
// 毎年同じ日付の祝日に使用
func StaticHolidayCheckerFactory(month time.Month, day int) (f func(Date) bool) {
	f = func(d Date) bool {
		if month == d.Month() && day == d.Day() {
			return true
		}
		return false
	}
	return
}

var (
	LOCATION_JP, _ = time.LoadLocation("Asia/Tokyo") // Japan Locale
	DHCF           = DynamicHolidayCheckerFactory    // Alias
	SHCF           = StaticHolidayCheckerFactory     // Alias
)

var (
	// 祝日チェッカー関数のmap
	NAMED_HOLIDAYS = map[NamedHoliday]func(Date) bool{
		GANTAN:       SHCF(1, 1),
		SEIJIN:       DHCF(1, 2, time.Monday),
		KENKOKUKINEN: SHCF(2, 11),
		SHUNBUN:      ShunbunCheker,
		SHOWA:        SHCF(4, 29),
		KENPOKINEN:   SHCF(5, 3),
		MIDORI:       SHCF(5, 4),
		KODOMO:       SHCF(5, 5),
		UMI:          DHCF(7, 3, time.Monday),
		KEIRO:        DHCF(9, 3, time.Monday),
		SHUBUN:       ShubunCheker,
		TAIIKU:       DHCF(10, 2, time.Monday),
		BUNKA:        SHCF(11, 3),
		KINROKANSHA:  SHCF(11, 23),
		TENNOTANJOBI: SHCF(12, 23),
	}
)

func init() {

	for i, day := range SHUNBUN_DAYS {
		year := i + 2000
		SHUNBUN_LIST[NewDate(year, 3, day)] = struct{}{}
	}

	for i, day := range SHUBUN_DAYS {
		year := i + 2000
		SHUBUN_LIST[NewDate(year, 9, day)] = struct{}{}
	}

}

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
func (d Date) Time() time.Time {
	return time.Date(d.Year(), d.Month(), d.Day, 0, 0, 0, LOLOCATION_JP)
}

// その曜日がその月で何回目かを返す
func (d Date) NthWeekday() (nth int) {
	day := d.Day()
	nth = ((day - 1) / 7) + 1
	return
}

// 国民の祝日ならtrueと祝日名を返す。振替休日はチェックしない。
func (d Date) RealHoliday() (isHoliday bool, name string) {
	for holiday_name, f := range NAMED_HOLIDAYS {
		if f(d) {
			isHoliday = true
			name = holiday_name.String()
			return
		}
	}
	return false, ""
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

// 祝日ならtrueと祝日名を返す。振替休日の祝日名は"振替休日"
func (d Date) Holiday() (isHoliday bool, name string) {
	if isHoliday, name = d.RealHoliday(); isHoliday {
		return
	} else {
		if d.AlternativeHoliday() {
			return true, "振替休日"
		}
	}
	return false, ""
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
