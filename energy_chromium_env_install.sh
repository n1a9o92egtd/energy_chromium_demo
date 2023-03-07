go get -u github.com/energye/energy
go install
energy install .
cd $(cd $(dirname $0);pwd)/energy_chromium_demo
go mod init energy_chromium_demo
go mod tidy
go build -ldflags "-s -w"
