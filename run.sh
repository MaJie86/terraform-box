go mod init github.com/majie86/terraform-box
go mod init github.com/majie86/terraform-box/api
go mod init github.com/majie86/terraform-box/cmd
go mod init github.com/majie86/terraform-box/pool

go mod edit -replace="github.com/majie86/terraform-box/api"="./api"
go mod edit -replace="github.com/majie86/terraform-box/cmd"="./cmd"
go mod edit -replace="github.com/majie86/terraform-box/pool"="./pool"
go mod tidy

go list -m -json all

go get github.com/majie86/terraform-box/api
go get github.com/majie86/terraform-box/cmd

git rm --cache -r .idea