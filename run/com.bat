@echo off
@REM parse_swf.ahk??????????????
setlocal enabledelayedexpansion
for /r ./Resource %%i in (*) do (
    set a=%%i
    if not exist %a:Resource=Resource_% ( echo %%i )
)
echo Íê³É