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

	ff := func(year int) func(month time.Month, day int) Date {
		f := func(month time.Month, day int) Date {
			return NewDate(year, month, day)
		}
		return f
	}
	var (
		f    func(time.Month, int) Date
		days []Date
	)

	//2011年の祝日チェック
	f = ff(2011)
	days = []Date{
		f(1, 1), f(1, 10), f(2, 11), f(3, 21), f(4, 29), f(5, 3), f(5, 4), f(5, 5),
		f(7, 18), f(9, 19), f(9, 23), f(10, 10), f(11, 3), f(11, 23), f(12, 23),
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
