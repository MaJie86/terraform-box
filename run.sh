go mod init github.com/MaJie86/terraform-box
go mod init github.com/MaJie86/terraform-box/api
go mod init github.com/MaJie86/terraform-box/cmd

go mod edit -replace="github.com/MaJie86/terraform-box/api"="./api"
go mod edit -replace="github.com/MaJie86/terraform-box/cmd"="./cmd"
go mod tidy

go list -m -json all

go get github.com/MaJie86/terraform-box/api
go get github.com/MaJie86/terraform-box/cmd

git rm --cache -r .idea