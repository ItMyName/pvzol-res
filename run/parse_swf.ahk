; autohotkey
; ����������autohotkey�Ľű��ļ�����������ffdecʹ��
; �����Ǹ���Ŀ¼�ṹ����swf�ļ������󵼳�Ϊ�ļ���
; ���������ʹ���Ҫ���ţ���ȥ�鿴��־�ļ������¼�������Ϣ

; �������������ؽ��е�����
; ��Դ�ļ��еĸ�Ŀ¼
global SrcDir := "D:\Code\PvzCapture\run\Resource" 
; ���ƴ�ŵĸ�Ŀ¼
global DstDir := "D:\Code\PvzCapture\run\Resource_" 
; ffdec��������·��
global FFdecPath ="ff.lnk"

; ��־�ļ����λ��
global log := DstDir "/log_.txt" 
global ffdecWindow:="JPEXS Free Flash Decompiler v.15.1.0"
global loadWindow:="FFDec v.15.1.0"

; ����ffdec��ȱ�ݣ������ļ������ᱻ������������ȷ���򿪶��ٴκ�����
global taskCount=30
global winid

; ���ô���Ϊ�״̬��������ڲ����������������
activateFFdec(){
    if not WinExist(ffdecWindow){
        Run, % FFdecPath
        WinWait, % ffdecWindow
        WinWaitClose, % loadWindow
    }
    WinActivate, % ffdecWindow
    WinWaitActive, % ffdecWindow
    ; �̶����ڿ��
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
; ��һ��src·�����ļ��ļ�����������dst·��
Out(src, dstDir) {
    FileAppend, do %src% -> %dstDir%`n, %log%
    i++
    if (i = taskCount){
        i=0
        Process, Close, %winid%
        Process, WaitClose, %winid%
    }
    activateFFdec()
    ; ��д·�����ļ�

    MouseClick, Left, 33, 77 
    WinWaitActive, ��
    Send % src
    Send {Enter}

    ; �ȴ�������ϣ��ڼ����Ƿ���ִ��󴰿ڣ�ֱ�����ְ׵�ѡ���
    Loop {
        if WinExist("��Ϣ") or WinExist("����") {
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

    ; ȫ����������д·��
    MouseClick, Left, 39, 169
    MouseClick, Left, 407, 96 
    WinWaitActive, ����...
    Send {Enter}
    Send % dstDir
    Send {Enter}
    WinWaitClose, ����...

    ; �ȴ�������ϣ��ڼ����Ƿ���ִ��󴰿�
    ; ������Ϻ���ȫ���رհ�ťֱ���׵�ѡ���ر�
    Loop 
    {
        if WinExist("��Ϣ") or WinExist("����") {
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

; ��Դ�ļ����µ������ļ�����Ŀ���ļ�����
; ���Դ�ļ���.swf�ļ���ᱻ��������д��
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
MsgBox, ������ɣ�����
^`::Reload
^+`::Pause