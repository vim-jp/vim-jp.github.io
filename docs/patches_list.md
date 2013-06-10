---
layout: docs
title: vim-jpで開発中のパッチを紹介
---

以下はvim-jpで現在開発中のパッチです。興味のあるパッチがあれば、ぜひテストや議論に協力をよろしくおねがいします。

### 開発中のパッチ
[開発中のパッチ一覧](https://github.com/vim-jp/issues/issues?labels=status%2Fpatched&page=1&state=open)


* UTF-8環境での印刷文字化け解消
  - 作者：k-takata
  - [パッチ URL](https://bitbucket.org/k\_takata/vim-win32-mq/src/a11f5ea38acf24f121fbb55f88181ecfebd2591a/fix-windows-printing.patch?at=default)
  - [Issue](https://github.com/vim-jp/issues/issues/408)
  - 現在の状況：開発中


* UNIX 環境における LuaJIT 対応
  - 作者：shirosaki
  - [パッチ URL](https://gist.github.com/shirosaki/5663617)
  - [Issue](https://github.com/vim-jp/issues/issues/348)
  - 現在の状況：ほぼ完成、vim\_devへの投稿予定


* Make\_cyg.mak/Make\_ming.mak 統合
  - 作者：k-takata
  - [パッチ URL](https://gist.github.com/k-takata/5018169)
  - [Issue](https://github.com/vim-jp/issues/issues/314)
  - 現在の状況：開発中、修正案を議論している段階


* make config 修正
  - 作者：shirosaki
  - [パッチ URL](https://bitbucket.org/k_takata/vim-cygwin-linux-mq/src/e09a74f02fcef2177998f6cee829fa99112cd89f/fix-make-config.patch)
  - [Issue](https://github.com/vim-jp/issues/issues/280)
  - 現在の状況：開発中


* switchbuf 挙動修正
  - 作者：h-east
  - [パッチ URL](https://github.com/vim-jp/issues/issues/162#issuecomment-5534773)
  - [Issue](https://github.com/vim-jp/issues/issues/162)
  - 現在の状況：開発中、まだ懸念が残っている


* lasteol オプション追加
  - 作者：koron
  - [パッチ URL](https://gist.github.com/3260998)
  - [Issue](https://github.com/vim-jp/issues/issues/152)
  - 現在の状況：開発中、修正案を議論している段階


* Dictionary 関数仕様変更
  - 作者：thinca
  - [パッチ URL](https://github.com/vim-jp/issues/issues/63#issuecomment-11980700)
  - [Issue](https://github.com/vim-jp/issues/issues/63)
  - 現在の状況：開発中？　議論が途中で止まっている


### 完成したパッチ
[完成したパッチ一覧](https://github.com/vim-jp/issues/issues?labels=status%2Fsent+patch&page=1&state=open)

* ポップアップの絞り込み時の再描画改善
  - 作者：Shougo `->` mattn
  - [パッチ URL](https://gist.github.com/Shougo/5743202)
  - [Issue](https://github.com/vim-jp/issues/issues/401)
  - [投稿](https://groups.google.com/d/msg/vim_dev/kFmrNmdnpSs/GMCMHodLflEJ)
  - 現在の状況：vim\_devへの投稿済み


