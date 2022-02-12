$env:TELEGRAM_TOKEN="YOUR_TOKEN"
$env:UNSPLASH_CLIENTID="YOUR_CLIENT_ID"
cd cmd
Remove-Item ./cmd.exe -ErrorAction Ignore
go build -mod vendor #.\main.go
cd ..
./cmd/cmd.exe