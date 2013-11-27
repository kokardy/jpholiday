package jpholiday

// 春分の日リスト。予測を含む。init()で初期化。
var SHUNBUN_LIST = map[Date]struct{}{}

// 春分の日判定関数
func ShunbunCheker(d Date) bool {
	if _, ok := SHUNBUN_LIST[d]; ok {
		return true
	}
	return false
}

var SHUNBUN_DAYS = []int{
	20, //2000年
	20,
	21,
	21,
	20,
	20,
	21,
	21,
	20,
	20,
	21,
	21,
	20,
	20,
	21,
	21,
	20,
	20,
	21,
	21,
	20,
	20,
	21,
	21,
	20,
	20,
	20,
	21,
	20,
	20,
	20, //2030年
}
