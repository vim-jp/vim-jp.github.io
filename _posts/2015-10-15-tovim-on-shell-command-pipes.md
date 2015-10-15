---
layout: post
category: blog
title: シェルのコマンド入出力パイプをvimで編集
---

reddit のトピック「[What's your best Vim related shell script?](https://www.reddit.com/r/vim/comments/3oo156/whats_your_best_vim_related_shell_script/)」より

```sh
#!/bin/bash
TOVIMTMP=~/.tovim_tmp_`date +%Y-%m-%d_%H-%M-%S.txt`
cat > $TOVIMTMP 
vim $TOVIMTMP < /dev/tty > /dev/tty
cat $TOVIMTMP 
rm $TOVIMTMP 
```

これを tovim というスクリプトファイルで保存し実行権限を付与します。シェルのコマンド入出力パイプに入れる事でストリームを編集出来る様になります。sed や grep でやっても良いのですが、編集中に undo が出来るという点でこちらにもメリットがありそうです。

```sh
$ firstcommand | tovim | othercommand
```

ただし reddit のトピック上に掲載されたスクリプトでは vim を異常終了させた場合、たとえば `:cquit` させた場合にコマンドパイプを中断出来ない問題があります。`set -e` を付け、かつ `trap` によりテンポラリファイルを削除する必要があるでしょう。

```sh
#!/bin/bash
set -e

TOVIMTMP=~/.tovim_tmp_`date +%Y-%m-%d_%H-%M-%S.txt`
trap 'rm $TOVIMTMP' ERR

cat > $TOVIMTMP
vim $TOVIMTMP < /dev/tty > /dev/tty
cat $TOVIMTMP
rm $TOVIMTMP
```
