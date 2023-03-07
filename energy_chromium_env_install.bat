rem https://github.com/energye/energy
go get -u github.com/energye/energy
go install
energy install %cd%
cd %cd%\energy_chromium_demo
go mod init energy_chromium_demo
go mod tidy
go build -ldflags "-H windowsgui -s -w"
