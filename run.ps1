$env:TELEGRAM_TOKEN=""
$env:UNSPLASH_CLIENTID=""
cd cmd
Remove-Item ./cmd.exe -ErrorAction Ignore
go build -mod vendor #.\main.go
cd ..
./cmd/cmd.exe