# nit

* git の拡張コマンド。
* いまのところ特定branchへのpushを禁止するだけ。
* [pre-push](https://git-scm.com/book/ja/v2/Git-%E3%81%AE%E3%82%AB%E3%82%B9%E3%82%BF%E3%83%9E%E3%82%A4%E3%82%BA-Git-%E3%83%95%E3%83%83%E3%82%AF)で、本来十分。
* 機能追加していくと便利になるかも。

nit is fobidden direct push gitcommand to develop, staging and master.

## Usage

```bash
$nit push origin master
You cant not push master to master.
```

* .nit.yml ( need .git same dir)

```yml
hooks:
  -
    prepush:
      forbidden:
        - master
        - develop
        - development
        - staging
```


## Installation

```
go get github.com/withnic/nit
```

# DONE
* push forbidden
* config by yaml

## TODO
