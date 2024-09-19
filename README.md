Hava Durumu Uygulaması
Bu uygulama, kullanıcıların belirli bir şehir için hava durumu bilgilerini görüntülemesine ve hava durumuna uygun aktivite önerileri almasına olanak tanır.
Özellikler

Şehir bazlı hava durumu bilgisi
Sıcaklık gösterimi
Hava durumuna uygun aktivite önerileri

Gereksinimler

Go 1.16 veya üzeri
OpenWeatherMap API anahtarı

Kurulum

Bu depoyu klonlayın:
Copygit clone https://github.com/scanalkan/hava-durumu-uygulamasi.git

Proje dizinine gidin:
Copy cd hava-durumu-uygulamasi

Gerekli bağımlılıkları yükleyin:
Copy go mod tidy

config.json dosyasını açın ve YOUR_API_KEY_HERE kısmını kendi OpenWeatherMap API anahtarınızla değiştirin.

Kullanım
Uygulamayı çalıştırmak için:
Copy go run main.go
Tarayıcınızda http://localhost:8080 adresine gidin ve uygulamayı kullanmaya başlayın.

