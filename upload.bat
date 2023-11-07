git add .
git commit -m "Last Commit"
git push
set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build main.go
del main.zip