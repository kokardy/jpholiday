package jpholiday

import "time"

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
	YAMA                      = iota

	//振替休日
	FURIKAEKYUJITSU NamedHoliday = iota
	//国民の休日
	KOKUMINNOKYUJITSU NamedHoliday = iota
)

// 祝日の名前マップ
var (
	HOLIDAY_NAMES = map[NamedHoliday]string{
		GANTAN:            "元旦",
		SEIJIN:            "成人の日",
		KENKOKUKINEN:      "建国記念の日",
		SHUNBUN:           "春分の日",
		SHOWA:             "昭和の日",
		KENPOKINEN:        "憲法記念日",
		MIDORI:            "みどりの日",
		KODOMO:            "こどもの日",
		UMI:               "海の日",
		KEIRO:             "敬老の日",
		SHUBUN:            "秋分の日",
		TAIIKU:            "体育の日",
		BUNKA:             "文化の日",
		KINROKANSHA:       "勤労感謝の日",
		TENNOTANJOBI:      "天皇誕生日",
		FURIKAEKYUJITSU:   "振替休日",
		KOKUMINNOKYUJITSU: "国民の休日",
		YAMA:              "山の日",
	}
)

var (
	LOCATION_JP, _ = time.LoadLocation("Asia/Tokyo") // Japan Locale
	DHCF           = DynamicHolidayCheckerFactory    // Alias
	SHCF           = StaticHolidayCheckerFactory     // Alias
	HCF            = HolidayCheckerFactory           // Alias
	EVER           = NewDate(2999, 12, 31)
	LAWDAY         = NewDate(1948, 7, 20)
)
var (
	// 祝日チェッカー関数のmap
	NAMED_HOLIDAYS = map[NamedHoliday]func(Date) bool{
		GANTAN: SHCF(1, 1, NewRange(LAWDAY, EVER)),
		SEIJIN: HCF(
			SHCF(1, 15, NewRange(LAWDAY, NewDate(1999, 12, 31))),
			DHCF(1, 2, time.Monday, NewRange(NewDate(2000, 1, 1), EVER))),
		KENKOKUKINEN: SHCF(2, 11, NewRange(NewDate(1966, 1, 1), EVER)),
		SHUNBUN:      ShunbunCheker,
		SHOWA:        SHCF(4, 29, NewRange(NewDate(2007, 1, 1), EVER)),
		KENPOKINEN:   SHCF(5, 3, NewRange(LAWDAY, EVER)),
		MIDORI: HCF(
			SHCF(4, 29, NewRange(NewDate(1989, 1, 1), NewDate(2006, 12, 31))),
			SHCF(5, 4, NewRange(NewDate(2007, 1, 1), EVER))),
		KODOMO: SHCF(5, 5, NewRange(LAWDAY, EVER)),
		UMI: HCF(
			SHCF(7, 20, NewRange(NewDate(1996, 1, 1), NewDate(2002, 12, 31))),
			DHCF(7, 3, time.Monday, NewRange(NewDate(2003, 1, 1), EVER))),
		KEIRO: HCF(
			SHCF(9, 15, NewRange(LAWDAY, NewDate(2002, 12, 1))),
			DHCF(9, 3, time.Monday, NewRange(NewDate(2003, 1, 1), EVER))),
		SHUBUN: ShubunCheker,
		TAIIKU: HCF(
			SHCF(10, 10, NewRange(LAWDAY, NewDate(1999, 12, 31))),
			DHCF(10, 2, time.Monday, NewRange(NewDate(2000, 1, 1), EVER))),
		BUNKA:       SHCF(11, 3, NewRange(LAWDAY, EVER)),
		KINROKANSHA: SHCF(11, 23, NewRange(LAWDAY, EVER)),
		TENNOTANJOBI: HCF(
			SHCF(4, 29, NewRange(LAWDAY, NewDate(1989, 2, 16))),
			SHCF(12, 23, NewRange(NewDate(1989, 2, 17), NewDate(2018, 12, 31))),
			SHCF(2, 23, NewRange(NewDate(2020, 1, 1), EVER))),
		YAMA: SHCF(8, 11, NewRange(NewDate(2016, 1, 1), EVER)),
	}
)

func (h NamedHoliday) String() string {
	return HOLIDAY_NAMES[h]
}

//第X Y曜日で判定する関数を作成するファクトリ関数
//ハッピーマンデーなどに使用。
func DynamicHolidayCheckerFactory(month time.Month, nth int, weekday time.Weekday, period Range) (f func(Date) bool) {
	f = func(d Date) bool {
		if !period.Contains(d) {
			return false
		}
		if month == d.Month() && weekday == d.Weekday() && nth == d.NthWeekday() {
			return true
		}
		return false
	}
	return
}

// 何月何日で判定する関数を作成するファクトリ関数
// 毎年同じ日付の祝日に使用
func StaticHolidayCheckerFactory(month time.Month, day int, period Range) (f func(Date) bool) {
	f = func(d Date) bool {
		if !period.Contains(d) {
			return false
		}
		if month == d.Month() && day == d.Day() {
			return true
		}
		return false
	}
	return
}

// 祝日判定関数のWrapper
func HolidayCheckerFactory(funcs ...func(Date) bool) (f func(Date) bool) {
	f = func(d Date) bool {
		for _, ff := range funcs {
			if ff(d) {
				return true
			}
		}
		return false
	}
	return
}
