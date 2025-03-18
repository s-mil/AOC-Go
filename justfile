default:
  @just --list
run year day part:
  go run {{year}}/day{{day}}/main.go --part {{part}}

create year day:
  go run template/template.go --year {{year}} --day {{day}}
