jpholiday
=========

日本の祝日ライブラリ。ある日付が祝日かどうか判定し、

祝日なら何の祝日かを返す関数がついてます。

振替休日/国民の休日かどうかも判定します。

春分の日と秋分の日については、2000年〜2030年のみ対応。

この2つの祝日は本来、前年の官庁発表で決定するので2015年以降は天文台予測の春分日、秋分日です。

また、海の日と敬老の日の2002年以前(ハッピーマンデー化前)は未対応なので、

実質2003年~2030年が対応年です。

```go
package main

import (
	"fmt"
	jp "github.com/kokardy/jpholiday"
)

func main() {
	day := jp.NewDate(2013, 5, 4)
	isHoliday, holiday := day.Holiday()
	if isHoliday {
		fmt.Printf("%s is %s \n", day, holiday) //2013-05-04 is みどりの日
	}
}
```


