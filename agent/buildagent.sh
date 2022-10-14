GOOS=windows go build -o agent.exe ./cmd 
zip agent.zip agent.exe config.json
rm agent.exe
