package jpholiday

// 秋分の日リスト。予測を含む。init()で初期化
var SHUBUN_LIST = map[Date]struct{}{}

// 秋分の日判定関数
func ShubunCheker(d Date) bool {
	if _, ok := SHUBUN_LIST[d]; ok {
		return true
	}
	return false
}

var SHUBUN_DAYS = []int{
	23, //2000年
	23,
	23,
	23,
	23,
	23,
	23,
	23,
	23,
	23,
	23,
	23,
	22,
	23,
	23,
	23,
	22,
	23,
	23,
	23,
	22,
	23,
	23,
	23,
	22,
	23,
	23,
	23,
	22,
	23,
	23, //2030年
}
