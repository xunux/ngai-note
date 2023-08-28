@echo off

rem 无限使用 NavicatPremium 15 版本

::setlocal EnableExtensions DisableDelayedExpansion
setlocal enabledelayedexpansion

rem 删除注册信息
reg delete "HKEY_CURRENT_USER\Software\PremiumSoft\NavicatPremium\Registration15XCS" /f 2>nul

set regpath=HKEY_CURRENT_USER\Software\Classes\CLSID

::REG QUERY %regpath% /s | find /I "Info"
REG QUERY %regpath% /s /k /f "Info"
IF ERRORLEVEL 1 (
  echo "%regpath%\ 没有匹配子项 Info."
) ELSE (
  echo "%regpath% 存在匹配子项 Info."
)

rem 在注册表 HKEY_CURRENT_USER\Software\Classes\CLSID 项下查找子项为 Info 的项，并删除
for /f "delims=" %%I in ('REG QUERY %regpath% /s /k /f "Info" ^| find /I "Info"') do (
    set "regkey=%%I"
    set "parent_regkey=!regkey:~0,-5!"
    @%SystemRoot%\System32\reg.exe delete "!parent_regkey!" /f 
)2>nul
