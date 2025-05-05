# UPS Kargo Takip Uygulaması

Bu uygulama, UPS kargo şirketinin web sitesinden paket takip bilgilerini çekerek, terminal üzerinde görsel olarak sunan bir Go programıdır.

## Özellikler

- UPS kargo takip numarası ile paket durumunu sorgulama
- Şık ve renkli terminal arayüzü (lipgloss kütüphanesi ile)
- Otomatik yenileme (her dakika)
- Yeni durum güncellemeleri için masaüstü bildirimleri
- Türkçe arayüz

## Kurulum

### Gereksinimler

- Go 1.16 veya üzeri

### Derleme

Uygulamayı derlemek için:

```bash
make build
```

Bu komut, uygulamayı `upsTrack.exe` olarak derleyecektir.

## Kullanım

Uygulamayı çalıştırdığınızda, sizden bir UPS takip numarası girmeniz istenecektir. Takip numarasını girdikten sonra, uygulama kargo durumunu gösterecek ve her dakika otomatik olarak güncelleyecektir.

```bash
./upsTrack.exe
```

## Nasıl Çalışır

Uygulama, UPS Türkiye web sitesinden (ups.com.tr) web scraping yöntemiyle takip bilgilerini çeker. Colly kütüphanesi kullanılarak HTML içeriği analiz edilir ve kargo durumu bilgileri çıkarılır.

## Katkıda Bulunma

1. Bu depoyu fork edin
2. Yeni bir özellik dalı oluşturun (`git checkout -b yeni-ozellik`)
3. Değişikliklerinizi commit edin (`git commit -am 'Yeni özellik eklendi'`)
4. Dalınıza push yapın (`git push origin yeni-ozellik`)
5. Bir Pull Request oluşturun

## Lisans

Bu proje MIT lisansı altında lisanslanmıştır - detaylar için [LICENSE](LICENSE) dosyasına bakınız.

---

# UPS Package Tracking Application

This is a Go program that fetches package tracking information from the UPS courier website and presents it visually in the terminal.

## Features

- Query package status with UPS tracking number
- Elegant and colorful terminal interface (using lipgloss library)
- Automatic refresh (every minute)
- Desktop notifications for new status updates
- Turkish interface

## Installation

### Requirements

- Go 1.16 or higher

### Building

To build the application:

```bash
make build
```

This command will compile the application as `upsTrack.exe`.

## Usage

When you run the application, you will be asked to enter a UPS tracking number. After entering the tracking number, the application will show the package status and automatically update it every minute.

```bash
./upsTrack.exe
```

## How It Works

The application fetches tracking information from the UPS Turkey website (ups.com.tr) using web scraping. The HTML content is analyzed using the Colly library, and package status information is extracted.

## Contributing

1. Fork this repository
2. Create a new feature branch (`git checkout -b new-feature`)
3. Commit your changes (`git commit -am 'Added new feature'`)
4. Push to the branch (`git push origin new-feature`)
5. Create a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
