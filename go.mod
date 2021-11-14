module github.com/majie86/terraform-box

go 1.16

require (
	github.com/majie86/terraform-box/api v0.0.0-00010101000000-000000000000
	github.com/majie86/terraform-box/cmd v0.0.0-00010101000000-000000000000
)

replace github.com/majie86/terraform-box/api => ./api

replace github.com/majie86/terraform-box/cmd => ./cmd
