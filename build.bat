@echo off
where go >nul 2>nul
if errorlevel 1 (
    echo Go is not installed. Please install Go: https://go.dev/dl/
    exit /b 1
)

if not exist build (
    mkdir build
)

echo Building Go Proximu...
go build -o build\Proximu.exe

if %errorlevel% equ 0 (
    echo Build successful: build/Proximu.exe
) else (
    echo Build failed.
    exit /b 1
)
