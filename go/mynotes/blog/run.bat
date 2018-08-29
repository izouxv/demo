@echo off
go build %-ldflags "-H windowsgui"%
blog.exe %-c=/data1/cfg.toml%
echo.