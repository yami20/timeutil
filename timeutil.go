package timeutil

import (
	"time"
)

// NextMonth 同Locationの翌月1日00:00:00を返す
func NextMonth(t time.Time) time.Time {
	if int(t.Month()) == 12 {
		return time.Date(t.Year()+1, 1, 1, 0, 0, 0, 0, t.Location())
	}
	return time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth 同Locationの同月最終日23:59:59.999999999 を返す. 実際にはNextMonthを < で使うことのほうが多いが一応用意しておく
func EndOfMonth(t time.Time) time.Time {
	return NextMonth(t).Add(time.Nanosecond * -1) // 翌月の1sec前
}

// LastDayOfMonth 同Locationの同月最終日00:00:00 を返す
func LastDayOfMonth(t time.Time) time.Time {
	return NextMonth(t).AddDate(0, 0, -1)
}

// CorrespondingDate 同Locationでy年mヶ月後の応当日の00:00:00を返す. 応当日の定義から該当日が存在しない場合は当該月の月末日を返す
func CorrespondingDate(t time.Time, y int, m int) time.Time {
	// 12ヶ月を1年に換算.愚直にやっても良いが簡単・高速化のため12ヶ月飛ばしで計算する
	for m > 12 {
		y++
		m -= 12
	}

	target := time.Date(t.Year()+y, t.Month(), 1, 0, 0, 0, 0, t.Location())
	for m > 0 {
		target = NextMonth(target)
		m--
	}

	if LastDayOfMonth(target).Day() <= t.Day() {
		return LastDayOfMonth(target)
	}

	return target.AddDate(0, 0, t.Day()-1)

}
