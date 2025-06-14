if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    Write-Host "Go is not installed. Please install Go: https://go.dev/dl/"
    exit 1
}

if (-not (Test-Path -Path "build")) {
    New-Item -ItemType Directory -Path "build" | Out-Null
}

Write-Host "Building Go Proximu..."
go build -o ./build/Proximu.exe

if ($LASTEXITCODE -eq 0) {
    Write-Host "Build successful: build/Proximu.exe"
} else {
    Write-Host "Build failed."
    exit 1
}