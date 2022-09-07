; autohotkey
; 这是适用于autohotkey的脚本文件，将它搭配ffdec使用
; 作用是复制目录结构并将swf文件解析后导出为文件夹
; 发生卡死和错误不要惊慌，请去查看日志文件那里记录着相关信息

; 以下三项参数务必进行调整：
; 资源文件夹的根目录
global SrcDir := "D:\Code\PvzCapture\run\Resource" 
; 复制存放的根目录
global DstDir := "D:\Code\PvzCapture\run\Resource_" 
; ffdec程序启动路径
global FFdecPath ="ff.lnk"

; 日志文件存放位置
global log := DstDir "/log_.txt" 
global ffdecWindow:="JPEXS Free Flash Decompiler v.15.1.0"
global loadWindow:="FFDec v.15.1.0"

; 由于ffdec的缺陷，当打开文件过多后会被卡死，在这列确定打开多少次后重启
global taskCount=30
global winid

; 设置窗口为活动状态，如果窗口不存在则会重新启用
activateFFdec(){
    if not WinExist(ffdecWindow){
        Run, % FFdecPath
        WinWait, % ffdecWindow
        WinWaitClose, % loadWindow
    }
    WinActivate, % ffdecWindow
    WinWaitActive, % ffdecWindow
    ; 固定窗口宽高
    WinMove, % ffdecWindow, , , , 800, 600
    Loop {
        MouseClick, Left , 235, 123
        PixelGetColor, color, 67, 503, RGB
        if (color==0xB1DDFF)
            Break
        Sleep, 200
    }
    WinGet, winid, PId, %ffdecWindow%
}

global i:=0
; 打开一个src路径的文件文件，并导出到dst路径
Out(src, dstDir) {
    FileAppend, do %src% -> %dstDir%`n, %log%
    i++
    if (i = taskCount){
        i=0
        Process, Close, %winid%
        Process, WaitClose, %winid%
    }
    activateFFdec()
    ; 填写路径打开文件

    MouseClick, Left, 33, 77 
    WinWaitActive, 打开
    Send % src
    Send {Enter}

    ; 等待导出完毕，期间检测是否出现错误窗口，直到出现白底选择框
    Loop {
        if WinExist("消息") or WinExist("错误") {
            FileAppend, Error!!!`n`n, %log%
            FileRemoveDir, %dstDir%, 1
            FileCopy, %src%, %dstDir%

            Send, {Enter}
            Process, Close, %winid%
            Process, WaitClose, %winid%
            Return
        }
        PixelGetColor, color, 67, 503, RGB
        if (color==0xFFFFFF)
            Break
        Sleep 200
    }

    ; 全部导出，填写路径
    MouseClick, Left, 39, 169
    MouseClick, Left, 407, 96 
    WinWaitActive, 导出...
    Send {Enter}
    Send % dstDir
    Send {Enter}
    WinWaitClose, 导出...

    ; 等待导出完毕，期间检测是否出现错误窗口
    ; 导出完毕后点击全部关闭按钮直到白底选择框关闭
    Loop 
    {
        if WinExist("消息") or WinExist("错误") {
            FileAppend, Error!!!`n`n, %log%
            FileRemoveDir, %dstDir%, 1
            FileCopy, %src%, %dstDir%

            Send, {Enter}
            Process, Close, %winid%
            Process, WaitClose, %winid%
            Return
        }
        MouseClick, Left , 235, 123
        PixelGetColor, color, 67, 503, RGB
        if (color==0xB1DDFF)
            Break
        Sleep 200
    }
}

; 将源文件夹下的所有文件复制目标文件夹下
; 如果源文件是.swf文件则会被解析后在写入
Copy(srcDir,dstDir){
    FileCreateDir, %dstDir%
    Loop, Files, %srcDir%/* , FD 
    {
        if InStr(A_LoopFileAttrib,"D")
        {
            Copy(srcDir "/" A_LoopFileName , dstDir "/" A_LoopFileName)
        } else If (A_LoopFileExt == "swf") {
            FileCreateDir, %dstDir%/%A_LoopFileName%
            Out(A_LoopFileFullPath, dstDir "/" A_LoopFileName)
        } else {
            FileCopy, %A_LoopFileFullPath%, %dstDir%/%A_LoopFileName%
        }
    }
}

Copy(SrcDir,DstDir)
MsgBox, 任务完成！！！
^`::Reload
^+`::Pause