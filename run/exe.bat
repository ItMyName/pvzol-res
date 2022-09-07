@REM 抓取文件
cd ..\src & go build -o ../run/main.exe
cd ..\run
echo %time%
main.exe -d200 del
echo %time%