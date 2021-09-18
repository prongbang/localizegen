android_gen_en:
	go run main.go -platform android -target ./example/android -locale en -document 1r91ECV-As0XtuqGKXU7dXnoY4og9XPBoCqwRcdio6EU -sheet 0

ios_gen_en:
	go run main.go -platform ios -target ./example/ios -locale en -document 1r91ECV-As0XtuqGKXU7dXnoY4og9XPBoCqwRcdio6EU -sheet 0

android_gen:
	go run main.go -platform android -target ./example/android -document 1r91ECV-As0XtuqGKXU7dXnoY4og9XPBoCqwRcdio6EU -sheet 0

ios_gen:
	go run main.go -platform ios -target ./example/ios -document 1r91ECV-As0XtuqGKXU7dXnoY4og9XPBoCqwRcdio6EU -sheet 0

build_linux:
	env GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o ./binary/linux/localizegen github.com/prongbang/localizegen

build_macos:
	env GOOS=darwin go build -ldflags "-w" -o ./localizegen github.com/prongbang/localizegen && chmod +x ./localizegen

build_window:
	env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./binary/windows/localizegen.exe github.com/prongbang/localizegen
