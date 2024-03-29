netstat -ano |findstr "8080" | findstr LIST >.\pid.txt
for /f "tokens=5 delims= " %%i in (.\pid.txt) do @taskkill /pid %%i -t -f
del .\pid.txt