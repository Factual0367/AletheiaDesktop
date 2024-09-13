# Makefile

build: linux windows

linux:
	go build -o fyne-cross/bin/linux-amd64/Aletheia-linux-amd64

windows:
	fyne-cross windows -arch=amd64 --app-id io.Aletheia.Aletheia
	mv fyne-cross/bin/windows-amd64/AletheiaDesktop.exe fyne-cross/bin/windows-amd64/Aletheia-windows-amd64.exe
