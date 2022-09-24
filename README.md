# ScreenShotGo

Go 语言编写的全屏截图软件，主要用途是电脑自动截图。使用了 [screenshot](https://github.com/kbinani/screenshot) 库。

### 参数

```sh
-d string
        Directory name. (default "./")
-f string
        File name. (default "NoFileName")
```

### PowerShell

```PowerShell
$data = Get-Date -Format 'yyyy-MM-dd'
$time = Get-Date -Format 'HH-mm'
ScreenShotGo.exe -d "D:\TEMP\PicLog\$data\" -f $time
```

