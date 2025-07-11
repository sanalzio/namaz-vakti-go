package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	// terminal renkleri için bir modül
	"namaz-vakti/colorama"
)

// biliyom aşırı karmaşık ama sadece renkli bir tablo oluşturmaya yarıyor.
var output string = "\n          " +
	colorama.Fore.BrightBlue +
	"Sonraki         Sonraki vakte" +
	colorama.Reset +
	"\n           " +
	colorama.Fore.BrightBlue +
	"vakit              kalan" +
	colorama.Reset +
	"\n\n         " +
	"%s" +
	colorama.Style.Underline +
	colorama.Fore.Yellow +
	"%s" +
	colorama.Reset +
	"             " +
	colorama.Style.Underline +
	colorama.Fore.Yellow +
	"%02d:%02d:%02d" +
	colorama.Reset +
	"\n\n  " +
	colorama.Fore.Magenta +
	"İmsak   Güneş   Öğle    İkindi   Akşam   Yatsı" +
	colorama.Reset + "\n  " +
	colorama.Fore.BrightMagenta +
	"%s   " +
	"%s   " +
	"%s   " +
	"%s    " +
	"%s   " +
	"%s" +
	colorama.Reset + "\n\n"

// bu da sonraki vakit kısmının atılmış hali (--yarin opsiyonu için)
var OutputTomorrow string = "\n  " +
	colorama.Fore.Magenta +
	"İmsak   Güneş   Öğle    İkindi   Akşam   Yatsı" +
	colorama.Reset + "\n  " +
	colorama.Fore.BrightMagenta +
	"%s   " +
	"%s   " +
	"%s   " +
	"%s    " +
	"%s   " +
	"%s" +
	colorama.Reset + "\n\n"

func writeTable(number *int, data *map[string]map[string]string) {
	// bunlarda renkli bir tablo için
	var intermediateLine string = colorama.Fore.Yellow + "├────────────┼───────┼───────┼───────┼────────┼───────┼───────┤" + colorama.Fore.Reset
	var replaceStr string = colorama.Fore.BrightMagenta + " %s " + colorama.Fore.Yellow
	var tableChar string = colorama.Fore.Yellow + "│" + colorama.Fore.Reset
	var tableLine string = tableChar + replaceStr + tableChar + replaceStr + tableChar + replaceStr + tableChar + replaceStr + tableChar + replaceStr + " " + tableChar + replaceStr + tableChar + replaceStr + tableChar + "\n"

	fmt.Println(colorama.Fore.Yellow + "┌────────────┬───────┬───────┬───────┬────────┬───────┬───────┐" + colorama.Fore.Reset + "\n" +
		colorama.Fore.Yellow + "│   Tarih    │       │       │       │        │       │       │" + colorama.Fore.Reset + "\n" +
		colorama.Fore.Yellow + "│ gün/ay/yıl │ İmsak │ Güneş │ Öğle  │ İkindi │ Akşam │ Yatsı │" + colorama.Fore.Reset)

	now := time.Now()

	for i := 0; i < *number; i++ {
		if i > 0 {
			now = now.AddDate(0, 0, 1)
		}

		// yıl-ay-gün  (Daha fazla bilgi: https://gosamples.dev/date-time-format-cheatsheet/)
		todayStr := now.Format("2006-01-02")

		vakitler, ok := (*data)[todayStr]
		if !ok {
			break
		}

		fmt.Println(intermediateLine)
		fmt.Printf(tableLine, now.Format("02/01/2006"), vakitler["İmsak"], vakitler["Güneş"], vakitler["Öğle"], vakitler["İkindi"], vakitler["Akşam"], vakitler["Yatsı"])
	}

	fmt.Println(colorama.Fore.Yellow + "└────────────┴───────┴───────┴───────┴────────┴───────┴───────┘" + colorama.Fore.Reset)
}

func main() {
	dirname, err := os.UserHomeDir()
        if err != nil {
		panic(err)
	}

	file, err := os.ReadFile(dirname + "/db/vakitler.json")
	if err != nil {
		panic(err)
	}

	// json'ı go verisine dönüştürüyor
	var data map[string]map[string]string
	err = json.Unmarshal(file, &data)
	if err != nil {
		panic(err)
	}

	now := time.Now()
	tomorrow := false

	// $ namaz-vakti argüman1 argüman2 şeklinde bir argüman var mı diye bakıyor
	if len(os.Args) > 1 {
		// $ namaz-vakti --yarin
		if os.Args[1] == "--yarin" || os.Args[1] == "--yarın" || os.Args[1] == "-y" || os.Args[1] == "--tomorrow" || os.Args[1] == "-t" {
			now = now.AddDate(0, 0, 1)
			tomorrow = true
		} else {
			number, err := strconv.Atoi(os.Args[1])
			if err != nil {
				fmt.Println("Hatalı giriş: Bir tam sayı yada --yarın kullanın.")
				return
			}

			writeTable(&number, &data)
			return
		}
	}

	// yıl-ay-gün  (Daha fazla bilgi: https://gosamples.dev/date-time-format-cheatsheet/)
	todayStr := now.Format("2006-01-02")
	// yıl-ay-gün saat:dakika (saat 24 saat oluyor yani 01 olarak değil 13 olarak gösteriliyor.)
	layout := "2006-01-02 15:04"

	vakitler, ok := data[todayStr]
	if !ok {
		fmt.Println("Bugüne ait vakit bilgisi yok.")
		return
	}

	if tomorrow {
		fmt.Printf(OutputTomorrow, vakitler["İmsak"], vakitler["Güneş"], vakitler["Öğle"], vakitler["İkindi"], vakitler["Akşam"], vakitler["Yatsı"])
		return;
	}

	var nextVakitName string
	//// var nextVakitTime time.Time
	var minFark time.Duration = time.Hour * 24
	found := false

	for vakitAdi, saatStr := range vakitler {
		fullStr := todayStr + " " + saatStr

		vakitTime, err := time.ParseInLocation(layout, fullStr, time.Local)
		if err != nil {
			fmt.Printf("Saat parse edilemedi: %s\n", fullStr)
			continue
		}

		if vakitTime.After(now) {
			fark := vakitTime.Sub(now)
			if fark < minFark {
				minFark = fark
				nextVakitName = vakitAdi
				//// nextVakitTime = vakitTime
				found = true
			}
		}
	}

	if found {
		solPadding := (8 - len([]rune(nextVakitName))) / 2
		solBosluk := strings.Repeat(" ", solPadding)

		//// fmt.Printf("Sıradaki vakit: %s (%s)\n", nextVakitName, nextVakitTime.Format("15:04"))

		saat := int(minFark.Hours())
		dakika := int(minFark.Minutes()) % 60
		saniye := int(minFark.Seconds()) % 60

		//// fmt.Printf(colorama.Fore.Green+"Kalan süre"+colorama.Fore.Reset+": %02d saat %02d dakika %02d saniye\n", saat, dakika, saniye)

		fmt.Printf(output, solBosluk, nextVakitName, saat, dakika, saniye, vakitler["İmsak"], vakitler["Güneş"], vakitler["Öğle"], vakitler["İkindi"], vakitler["Akşam"], vakitler["Yatsı"])
	} else {
		fmt.Println("Bugün için kalan bir vakit yok.")
	}
}
