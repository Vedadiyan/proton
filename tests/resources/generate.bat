@echo off
set /p file_path=".Proto File Path: "
set /p out_dir="Output Directory: "
@echo on
protoc --plugin=protoc-gen-protogenic=./protogenic.exe --go_out=./ --csharp_out=./  %file_path% --protogenic_out=%out_dir%
pause