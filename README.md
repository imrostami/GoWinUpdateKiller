# Go Win Update Killer

A tool built in golang to prevent automatic Windows update, just copy the file windowsUpdateKiller.exe to the startup of your operating system and get rid of arbitrary execution of Windows update.


If you want to compile the source yourself, pay attention to the fact that if you want to hide the console so that it is displayed in the background of the system during execution, use the command 

```
go build -ldflags="-H windowsgui" -o output.exe

```
