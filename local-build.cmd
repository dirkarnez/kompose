@REM SET GOOS=windows
@REM SET GOARCH=amd64
@REM go build

SET GOOS=js
SET GOARCH=wasm

@REM git clone -b v1.34.0 https://github.com/kubernetes/kompose.git &&^
@REM copy .\main_js.go .\kompose\

%USERPROFILE%\Downloads\go1.21.0.windows-amd64\go\bin\go build -o main.wasm main_js.go

@REM copy "%GOROOT%misc\wasm\wasm_exec.js" build-webassembly\
