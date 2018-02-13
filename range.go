package jpholiday

//Rangeは祝日の有効期間を表す
type Range struct {
	Start Date
	End   Date
}

func NewRange(d1, d2 Date) Range {
	return Range{d1, d2}
}

//Range.Containsは日付を引数にとって範囲にあればtrue otherwise false
func (r Range) Contains(d Date) (b bool) {
	start := r.Start
	end := r.End
	if start.Before(d.Time) && end.After(d.Time) {
		b = true
	} else {
		b = false
	}
	return
}
