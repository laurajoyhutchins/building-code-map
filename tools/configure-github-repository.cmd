@echo off
setlocal
pwsh -NoLogo -NoProfile -ExecutionPolicy Bypass -File "%~dp0configure-github-repository.ps1" %*
exit /b %ERRORLEVEL%
