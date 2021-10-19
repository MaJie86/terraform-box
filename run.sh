go mod init github.com/MaJie86/terraform-box
go mod init github.com/MaJie86/terraform-box/api

go mod edit -replace="github.com/MaJie86/terraform-box/api"="./api"
go mod tidy

git rm --cache -r .idea