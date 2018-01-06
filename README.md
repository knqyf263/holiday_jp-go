# holiday_jp-go
Japanese holiday (insipired by [emasaka/jpholidayp](https://github.com/emasaka/jpholidayp))

Is it holiday today in Japan?

Recommend: Use with cron

## Description

Check "Is it holiday today?" in Japan.

No arguments, and no output.
Reports as exit status:

* 0: holiday
* 1: not holiday
* 2: error

"Holiday" means:

* Sunday and Saturday
* In [holiday_jp](https://github.com/holiday-jp/holiday_jp)


## Examples

usages in crontab.

    holiday_jp-go && some-command

Run some-command in holiday.

    holiday_jp-go || some-command

Run some-command in weekday.

## Install
### RedHat, CentOS
```
$ sudo rpm -ivh https://github.com/knqyf263/holiday_jp-go/releases/download/v0.0.1/holiday_jp-go_0.0.1_linux_amd64.rpm
```

### Debian, Ubuntu
```
$ wget https://github.com/knqyf263/holiday_jp-go/releases/download/v0.0.1/holiday_jp-go_0.0.1_linux_amd64.deb
$ dpkg -i holiday_jp-go_0.0.1_linux_amd64.deb
```

### Other 
Download binary from https://github.com/knqyf263/holiday_jp-go/releases

## Requires
Nothing!


## Author

Teppei Fukuda