package etc

import (
	"os"
	"path/filepath"
)

func PathWindowsYaBrowser() (path string) {
	return filepath.Join(os.Getenv("HOMEDRIVE")+os.Getenv("HOMEPATH"), "AppData", "Local", "Yandex", "YandexBrowser", "Application", "browser.exe")
	//filepath.Join("%LocalAppData%", "Yandex", "YandexBrowser", "Application", "browser.exe")
}
