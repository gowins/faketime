# Faketime

This lib is inspired by https://github.com/chaos-mesh/chaos-mesh, and uses some of the code.

See [blog](https://int64.me/2020/%e5%9c%a8%e5%ae%b9%e5%99%a8%e4%b8%ad%e8%ae%a9%e6%97%b6%e9%97%b4%e8%87%aa%e7%94%b1%e6%91%87%e6%91%86.html) for more details.

**NOTE: only works with Linux Amd64.**

### Build

```sh
[root@tencent-shit faketime]# go mod tidy && go build
```

### Magic

```sh
[root@tencent-shit faketime]# faketime --pid [pid] --sec [sec]
```

### Help

```sh
[root@tencent-shit faketime]# ./faketime -h
faketime is simulated time

Usage:
  faketime [flags]

Flags:
  -h, --help      help for faketime
      --pid int   process identification number
      --sec int   TV_SEC_DELTA
```
