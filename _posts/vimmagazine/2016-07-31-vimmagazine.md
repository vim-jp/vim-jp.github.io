---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2016 年 7 月号

---

## 話題

*   [Osaka.vim #7](http://osaka-vim.connpass.com/event/32702/) 7/2(土) 大阪 が開催されました
*   [Software Design 2016年8月号 Vimの細道【10】 Vimからgitを使い倒す by mattn][sd201608]
*   [Vim の地味だけどよく使う設定](http://note103.hateblo.jp/entry/2016/07/06/114746)
*   [vimでソースを追うときの小技](http://qiita.com/masaharu-suizu/items/2e2fed71bf4074d05713)
*   [sudoを付けずにVimでルート権限のファイルを編集してしまった時](http://qiita.com/shishamo_dev/items/ab722c53301b95ccfac1)
*   [Vim v.s. Emacsをフローチャートで表現](http://rcmdnk.github.io/blog/2016/07/21/computer-vim-emacs/) - 編者注:フローチャートというのは誤った用語だと思われます
*   [VimとGitの効率的なコラボレーションの仕方](http://myenigma.hatenablog.com/entry/2016/07/10/084048)
*   [Vim-Pokemon ～あなたが Vim で開いているファイルに潜んでいるポケモン～](http://d.hatena.ne.jp/heavenshell/20160715/1468606360)
*   [Vimのおすすめプラグイン23個とその設定](http://post.simplie.jp/posts/56)
*   [VimのmapleaderにMacのCommandを割り当てたい](http://ja.stackoverflow.com/q/27794/208)
*   [vimのヤンクがmacのクリップボードに保存されない](https://teratail.com/questions/41191)
*   [DockerとWebSocketを使って、vimの設定をブラウザで即体感できるサービスを作った](http://mosuke5.hateblo.jp/entry/2015/07/19/135844)
*   [tmux 起動でvimなどがおかしくなった時の対処法](http://kirin.hatenadiary.jp/entry/2016/07/16/150206)

## 今月の新機能

*   7.4.1971: `:clist` に `+{count}` が指定できるようになりました。
*   7.4.1975: Windows で 2GB を超えるファイルが正しく扱えるようになりました。
*   7.4.1976: Vim script の整数の内部表現が 64bit になりました。
*   7.4.1989: `filter()` と `map()` に関数への参照を指定できるようになりました。
*   7.4.1997: `:cbottom` が追加されました。
*   7.4.2008: `execute()` が追加されました。
*   7.4.2010: `:lbottom` が追加されました。
*   7.4.2011: `getcompletion()` が追加されました。
*   7.4.2015: `test_autochdir()` が追加されました。
*   7.4.2029: `printf()` が64bit整数に対応しました。
*   7.4.2034: MinGWでのビルド時に `ARCH` が自動設定されるようになりました。
*   7.4.2044: lambda が導入されました。
*   7.4.2049: `:chistory` と `:lhistory` が追加されました。
*   7.4.2071: 型の値をあらわす `v:t_` で始まる定義済変数が追加されました。
*   7.4.2072: `substitute()` に関数への参照を指定できるようになりました。
*   7.4.2075: `TabNew` イベントと `WinNew` イベントが追加されました。
*   7.4.2077: `TabClosed` イベントが追加されました。
*   7.4.2095: `assert_inrange()` が追加されました。
*   7.4.2109: `display` オプションに `truncate` が追加されました。
*   7.4.2111: `.vimrc` が見つからなかった時に `defaults.vim` が読み込まれるようになりました。
*   7.4.2119: lambda内で外側のスコープの変数をキャプチャできるようになりました。
*   7.4.2120: ユーザー関数に `closure` を指定できるようになりました。
*   7.4.2121: lambda と closure のサポートを `+lambda` で確認できるようになりました。
*   [パッチ番号なし](travis+osx): Travis CI で OSX のテストを実行するようになりました。

[travis+osx]:https://github.com/vim/vim/commit/b20617b0b0d15523961618717d6cd2195674c363
[sd201608]:http://gihyo.jp/magazine/SD/archive/2016/201608
