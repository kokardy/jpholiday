package jpholiday

func init() {
	//春分の日のリストを初期化
	for i, day := range SHUNBUN_DAYS {
		year := i + 2000
		SHUNBUN_LIST[NewDate(year, 3, day)] = struct{}{}
	}

	//秋分の日のリストを初期化
	for i, day := range SHUBUN_DAYS {
		year := i + 2000
		SHUBUN_LIST[NewDate(year, 9, day)] = struct{}{}
	}

}
