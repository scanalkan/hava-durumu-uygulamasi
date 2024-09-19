package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Config struct {
	APIKey string `json:"api_key"`
}

type WeatherResponse struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func loadConfig() (Config, error) {
	var config Config
	file, err := os.Open("config.json")
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return config, err
}

func getWeatherData(city string, apiKey string) (*WeatherResponse, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("hava durumu bilgileri alınamadı: %v", err)
	}
	defer response.Body.Close()

	var weatherData WeatherResponse
	if err := json.NewDecoder(response.Body).Decode(&weatherData); err != nil {
		return nil, fmt.Errorf("JSON verileri çözümlenemedi: %v", err)
	}

	return &weatherData, nil
}

func main() {
	config, err := loadConfig()
	if err != nil {
		fmt.Println("Konfigurasyon yüklenemedi:", err)
		return
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hava Durumu Uygulaması",
		})
	})

	r.POST("/weather", func(c *gin.Context) {
		city := c.PostForm("city")
		weatherData, err := getWeatherData(city, config.APIKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		activity := suggestActivity(weatherData.Weather[0].Description)

		c.HTML(http.StatusOK, "weather.html", gin.H{
			"city":        city,
			"weather":     weatherData.Weather[0].Description,
			"temperature": fmt.Sprintf("%.2f", weatherData.Main.Temp),
			"activity":    activity,
		})
	})

	r.Run(":8080")
}

func suggestActivity(weatherDescription string) string {
	switch weatherDescription {
	case "thunderstorm with light rain":
		return "Hafif yağmurlu gök gürültüsü, evde rahatlatıcı bir müzik dinleyebilirsiniz."
	case "thunderstorm with rain":
		return "Yağmurlu ve gök gürültülü hava, evde sıcak bir çay eşliğinde film izlemek iyi olur."
	case "thunderstorm with heavy rain":
		return "Şiddetli yağmur ve gök gürültüsü, evde kalıp kitap okumak en iyisi."
	case "light thunderstorm":
		return "Hafif gök gürültüsü, içeride bir hobi ile ilgilenmek güzel olabilir."
	case "thunderstorm":
		return "Gök gürültülü hava, evde dinlenip rahatlamak iyi bir tercih."
	case "heavy thunderstorm":
		return "Şiddetli gök gürültüsü, dışarı çıkmaktan kaçınıp güvenli bir yerde kalmak en iyisi."
	case "ragged thunderstorm":
		return "Parçalı gök gürültüsü, evde güvenli bir şekilde kalmak akıllıca olur."
	case "thunderstorm with light drizzle":
		return "Hafif çiseleyen yağmurla gök gürültüsü, içeride rahat bir aktivite yapmak uygun olur."
	case "thunderstorm with drizzle":
		return "Çiseleyen yağmur ve gök gürültüsü, iç mekanda vakit geçirmek en iyi seçenek."
	case "thunderstorm with heavy drizzle":
		return "Şiddetli çiseleme ve gök gürültüsü, güvenli bir şekilde içeride kalmak en iyisi."
	case "light intensity drizzle":
		return "Hafif çiseleyen yağmur, pencere kenarında kahve içmek için güzel bir fırsat."
	case "drizzle":
		return "Çiseleyen yağmur, içeride bir film veya dizi izlemek keyifli olabilir."
	case "heavy intensity drizzle":
		return "Şiddetli çiseleme, dışarı çıkmaktan kaçınıp içeride vakit geçirmek iyi bir fikir."
	case "light intensity drizzle rain":
		return "Hafif çiseleyen yağmur, dışarıda yürüyüş yerine evde bir kitap okumak keyifli olabilir."
	case "drizzle rain":
		return "Çiseleyen yağmur, sıcak bir içecek eşliğinde dinlenmek iyi olur."
	case "heavy intensity drizzle rain":
		return "Şiddetli çiseleyen yağmur, evde vakit geçirmek için harika bir zaman."
	case "shower rain and drizzle":
		return "Sağanak ve çiseleme, içeride kalıp bir film izlemek güzel olabilir."
	case "heavy shower rain and drizzle":
		return "Şiddetli sağanak yağmur ve çiseleme, evde vakit geçirmek en iyisi."
	case "shower drizzle":
		return "Çiseleyen yağmur, iç mekanda sessiz bir aktivite yapmak uygun olur."
	case "light rain":
		return "Hafif yağmur, pencere kenarında kitap okumak için güzel bir hava."
	case "moderate rain":
		return "Orta şiddetli yağmur, evde film izlemek keyifli olabilir."
	case "heavy intensity rain":
		return "Şiddetli yağmur, içeride vakit geçirmek en iyisi."
	case "very heavy rain":
		return "Çok şiddetli yağmur, dışarı çıkmaktan kaçınıp evde kalmak iyi olur."
	case "extreme rain":
		return "Aşırı yağmur, evde kalıp güvenli bir yerde olmak en iyisi."
	case "freezing rain":
		return "Donmuş yağmur, dışarı çıkmak tehlikeli olabilir, içeride kalmak en iyisi."
	case "light intensity shower rain":
		return "Hafif sağanak yağmur, pencere kenarında kahve içip rahatlayabilirsiniz."
	case "shower rain":
		return "Sağanak yağmur, içeride bir hobi ile ilgilenmek iyi olabilir."
	case "heavy intensity shower rain":
		return "Şiddetli sağanak yağmur, içeride kalıp bir kitap okumak keyifli olabilir."
	case "ragged shower rain":
		return "Parçalı sağanak yağmur, dışarı çıkmamak ve içeride vakit geçirmek iyi olur."
	case "light snow":
		return "Hafif kar, dışarıda yürüyüş yapıp karın tadını çıkarabilirsiniz."
	case "snow":
		return "Kar yağıyor, kar topu oynamak veya kardan adam yapmak harika olabilir!"
	case "heavy snow":
		return "Yoğun kar, dışarıda kar aktiviteleri yapmak eğlenceli olabilir."
	case "sleet":
		return "Sulu kar, evde sıcak bir içecek eşliğinde vakit geçirebilirsiniz."
	case "light shower sleet":
		return "Hafif sağanak sulu kar, içeride kalmak daha iyi olabilir."
	case "shower sleet":
		return "Sağanak sulu kar, içeride güvenli bir şekilde vakit geçirmek iyi olur."
	case "light rain and snow":
		return "Hafif yağmur ve kar, sıcak bir içecek ile evde vakit geçirmek iyi olur."
	case "rain and snow":
		return "Yağmur ve kar, içeride kalıp dinlenmek için güzel bir zaman."
	case "light shower snow":
		return "Hafif sağanak kar, kar yürüyüşü yapmak güzel olabilir."
	case "shower snow":
		return "Sağanak kar, dışarıda karla eğlenmek güzel olabilir."
	case "heavy shower snow":
		return "Şiddetli sağanak kar, dışarıda kar aktiviteleri yapmak harika olabilir."
	case "mist":
		return "Sisli hava, meditasyon yapmak veya hafif bir yürüyüş yapmak uygun olabilir."
	case "smoke":
		return "Dumanlı hava, dışarı çıkmaktan kaçınmak en iyisi."
	case "haze":
		return "Puslu hava, içeride kalıp bir hobi ile ilgilenebilirsiniz."
	case "sand/dust whirls":
		return "Kum veya toz fırtınaları, dışarı çıkmaktan kaçınıp güvenli bir yerde kalmak iyi olur."
	case "fog":
		return "Sisli hava, yürüyüş yapmaktansa içeride kalmak daha güvenli olabilir."
	case "sand":
		return "Kum fırtınası, dışarı çıkmaktan kaçınıp güvenli bir yerde kalmak en iyisi."
	case "dust":
		return "Tozlu hava, içeride kalmak sağlığınız için daha güvenli olur."
	case "volcanic ash":
		return "Volkanik kül var, evde kalıp pencereleri kapatmak en iyisi."
	case "squalls":
		return "Fırtına var, dışarı çıkmak tehlikeli olabilir, içeride kalmak iyi olur."
	case "tornado":
		return "Tornado var, acil durum planınızı uygulayın ve güvenli bir yerde kalın."
	default:
		return "Hava durumuna uygun bir aktivite önerisi bulunamadı."
	}
}
