---
layout: post
category: blog
title: 端末機能に関する変更
---
# 自動クローズ

これまで Vim の `:terminal` はシェルが終了した後もウィンドウを残す動作でしたが、[Patch 8.0.1612](https://github.com/vim/vim/commit/1dd98334d6daee8abefcd640291d4b777d9f0f96) にてデフォルトで閉じる様に変更されました。ただし `:terminal make` の様にコマンドライン引数を与えた場合はこれまで通りウィンドウが残る動作となります。さらにこれまで通り

```
:terminal ++close make
```

`++close` フラグを付ける事でコマンドライン引数を指定したとしても自動で閉じる動作もこれまで通りです。

# GUI 版の :shell が :terminal と同等に

`guioptions` に `!` が追加されました。このオプションが指定されている時には GUI 版での `:shell` が `:terminal` と同等になります。ただし `:terminal` と異なり画面いっぱいに表示されます。デフォルトは `!` が付いていませんのでもしこの動作がお好みであれば gvimrc で

```
set guioptions+=!
```

を追記して下さい。
