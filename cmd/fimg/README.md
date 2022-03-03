# fimg

> 终端图片处理工具


```shell
go install github.com/evercyan/brick/cmd/fimg@latest
fimg help
```

---

## splicing 拼接图片

```shell
目录下图片文件格式需要满足 xxx_数字.(png|jpg|jpeg), 程序会按 `数字` 顺序来依次拼接

Usage:
  fimg splicing [flags]

Flags:
      --col int        图片列数 (default 2)
      --color string   背景颜色 (default "ffffff")
  -d, --dir string     图片目录, 目录下图片文件格式需要满足 xxx_数字.(png|jpg|jpeg)
  -h, --help           help for splicing
      --padding int    图片边距 (default 20)
      --quality int    图片质量: 取值范围 1-100 (default 100)
      --row int        图片行数 (default 2)
      --space int      图片间距 (default 10)
      --watermark      图片水印
```

来一堆 🌰🌰🌰

```shell
fimg splicing -d /tmp --row=5 --col=4 --color="ffff00" --padding=40 --space=20 --quality=70
```

- --row=5 --col=4   图片共 5 行 4 列, 即取目录下前 20 张图片拼接
- --color="ffff00"  图片背景色为 "ffff00"
- --padding=40      图片边距为 40px
- --space=20        图片间距为 20px
- --quality=50      图片大小为原图平均值的 70/100, 即 70%

![splicing](https://cdn.jsdelivr.net/gh/evercyan/repository/resource/37/376751123f3047b6c0cf7884a063120f.png)