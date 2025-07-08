# namaz-vakti
Go dilini öğrenmek için yaptığım bir CLI uygulaması.

## Açıklama
Basitçe aynı dizinde bulunan `vakitler.json` dosyasından aldığı veriler ile namaz vakitlerini gösterir.

## Kullanımı
- Go dilini bilgisayarınıza yükleyin.
- Depoyu bilgisayarınıza kopyalayın.
- `/namaz-vakti` dizinin içinde terminal açın.
- `$ go build` komutunu kullanın.
- Artık `$ ./namaz-vakti` komutu ile kullanabilirsiniz.

## Özellikleri
1) Düz `$ ./namaz-vakti` şeklinde kullanırsanız sistem saatinize göre günün vakitlerini gösterir.
2) `$ ./namaz-vakti --yarin` şeklinde kullanırsanız sistem saatinize göre yarının vakitlerini gösterir.
3) `$ ./namaz-vakti 5` şeklinde kullanırsanız sistem saatinize göre bu gün dahil 5 günün vakitlerini gösterir.

## `vakitler.json` verisini kendi konumum için nasıl bulurum?
Eee... Bunu açıklamak zaman alacak. İlk önce veriyi diyanet sitesinden alıyorum.
- İlk önce [Diyanet sitesi](https://namazvakitleri.diyanet.gov.tr/)'ne gidin.
- Kendi konumunuzu seçin.
- Aşağıya indiğinizde bir namaz vakitleri tablosu buluyorsunuz.
- Tablonun üst kısmında `Haftalık`, `Aylık`, `Yıllık` şeklinde namaz vakitleri var.
- `Yıllık` seçeneğini seçince `Kayıt sayısı ∇`, `PDF` ve `Excel` tuşları çıkıyor.
- `Excel` tuşunu seçiyoruz. İnen veriyi depodaki `/parse-data` klasörüne atıyorum.
- İnen veri dosyasının ismini `vakitler.xlsx` olarak sayalım.
- Python ile yazdığım `xlstocsv.py` betiği ile `$ python ./xlstocsv.py ./vakitler.xlsx` komutu ile Excel verisini csv verisine çeviriyorum.
- Daha sonra Bun (`javascript`) ile yazdığım `csvtojson.js` betiği ile `$ bun ./csvtojson.js ./vakitler.csv` komutu ile CSV verisini JSON verisine çeviriyorum.
- Sonra vakitler.json dosyasını derlenmiş program dosyasının oldu klasöre atarak `$ ./namaz-vakti` komutu ile programı kullanabilirsiniz.

### Ne diye antin kuntin diller kullandın?
Hızlıca bildiğim bir dil kullandım işte.