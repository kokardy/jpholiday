package jpholiday

import (
	"testing"
	"time"
)

// 祝日チェック
func TestDate(t *testing.T) {
	namedHolidays := []NamedHoliday{
		GANTAN,
		SEIJIN,
		KENKOKUKINEN,
		SHUNBUN,
		SHOWA,
		KENPOKINEN,
		MIDORI,
		KODOMO,
		UMI,
		KEIRO,
		SHUBUN,
		TAIIKU,
		BUNKA,
		KINROKANSHA,
		TENNOTANJOBI,
	}

	year_maker := func(year int) func(month time.Month, day int) Date {
		f := func(month time.Month, day int) Date {
			return NewDate(year, month, day)
		}
		return f
	}
	var (
		d    func(time.Month, int) Date
		days []Date
	)

	//2011年の祝日チェック
	d = year_maker(2011)
	days = []Date{
		d(1, 1), d(1, 10), d(2, 11), d(3, 21), d(4, 29), d(5, 3), d(5, 4), d(5, 5),
		d(7, 18), d(9, 19), d(9, 23), d(10, 10), d(11, 3), d(11, 23), d(12, 23),
	}
	for i, day := range days {
		_, holiday := day.Holiday()
		if namedHolidays[i] != holiday {
			t.Fatal("not equals", namedHolidays[i], holiday)
		}
	}
}

// 振替休日
func TestFurikae(t *testing.T) {
	d := NewDate
	days := []Date{
		d(2001, 2, 12), d(2005, 3, 21), d(2016, 3, 21),
		d(2008, 5, 6), d(2009, 5, 6),
		d(2014, 11, 24), d(2025, 11, 24),
	}
	for _, day := range days {
		_, holiday := day.Holiday()
		if holiday != FURIKAEKYUJITSU {
			t.Fatal("振替休日 CHECK FAIL", day)
		}
	}
}

// 国民の休日
func TestKokumin(t *testing.T) {
	d := NewDate
	days := []Date{
		d(2009, 9, 22), d(2015, 9, 22), d(2026, 9, 22),
		d(2004, 5, 4), d(2005, 5, 4), //みどりの日以前
	}
	for _, day := range days {
		_, holiday := day.Holiday()
		if holiday != KOKUMINNOKYUJITSU {
			t.Fatal("国民の休日 CHECK FAIL", day)
		}
	}
}

//山の日
func TestYamanohi(t *testing.T) {

	d := NewDate
	days := []Date{
		d(2016, 8, 11),
		d(2017, 8, 11),
	}
	for _, day := range days {
		if isholiday, holiday := day.Holiday(); !isholiday {
			t.Fatalf("day:%s must be 山の日 but %s", day, holiday)
		}
	}
}

//天皇誕生日
func TestTennotanjobi(t *testing.T) {

	d := NewDate
	days := []Date{
		//祝日法は1948/07/20なので1948は天皇誕生日ではなく天長節
		d(1949, 4, 29), //昭和
		d(1960, 4, 29),
		d(1988, 4, 29),
		d(1989, 12, 23), //平成
		d(2000, 12, 23),
		d(2001, 12, 23),
		d(2002, 12, 23),
		d(2006, 12, 23),
		d(2020, 2, 23), //2020年以降
		d(2023, 2, 23),
	}
	for _, day := range days {
		if isholiday, holiday := day.Holiday(); !isholiday {
			t.Fatalf("day:%s must be 天皇誕生日 but %s", day, holiday)
		}
	}

	days = []Date{
		d(2019, 12, 23),
		d(2019, 2, 23),
	}

	for _, day := range days {
		if isholiday, holiday := day.Holiday(); isholiday {
			t.Fatalf("day:%s must not be 天皇誕生日 but %s", day, holiday)
		}
	}

}
