---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2016 年 6 月号

---

## 話題

*   [netupvim 1.0 リリース](http://www.kaoriya.net/blog/2016/06/06/) 最新版は [1.1](https://github.com/koron/netupvim/releases/tag/v1.1) であることに注意してください
*   [Software Design 2016年7月号 Vimの細道【9】カーソルを自在に操る](http://gihyo.jp/magazine/SD/archive/2016/201607)
*   [なぜAtomはVimを置き換えられないのか](https://yakst.com/ja/posts/3985)

## 今月の新機能

*   7.4.1892: `'balloonexpr'` に指定した関数内で、定義済変数 `v:beval_winid` によりウィンドウIDを参照できるようになりました。
*   7.4.1893: バッファのウィンドウIDを取得する関数 `bufwinid()` が追加されました。
*   7.4.1894: `getchar()` でマウスクリックイベントを取得したときに、定義済変数 `v:mouse_winid` によりウィンドウIDが参照できるようになりました。
*   7.4.1895: ウィンドウIDが1000始まりになり、ウィンドウ番号(1始まり)との混乱が解消されました。
*   7.4.1898: ユーザーコマンドにコマンド修飾子 `<mods>` が指定できるようになりました。
*   7.4.1903: viminfo ファイルのフォーマットに変更がありました。
*   7.4.1925: `:clearjumps` コマンドが追加されました。
*   7.4.1952: `:cscope find` のクエリタイプに `a` が追加されました。 `:help cscope-find`
*   7.4.1960: Unicode 9.0 に対応
