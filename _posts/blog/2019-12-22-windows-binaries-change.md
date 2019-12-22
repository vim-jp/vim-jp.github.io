---
layout: blog
category: blog
title: Windows バイナリのリンク先変更のお知らせ
author: '<a href="https://twitter.com/kaoriya">KoRoN a.k.a. kaoriya</a>'

---

vim-jp のサイトにおいて右に表示されているVimのバイナリーへのリンクのうち
Windows版のものを [vim-win32-installer][vwi] に変更しました。

[vim-w32-installer][vwi] はVim本体に更新があれば [CI][ci] によって自動的に毎日決まった時刻(日本の朝の時間帯)にビルドされ、
テストが完走すれば新たな版としてリリースされます。
そのため常に最新版のVimを利用できます。
極稀にテストをパスしたけれども何かしらの不具合を抱えている場合もありますが、
過去にリリースされたものを [Releases][rels] からダウンロードできます。

リンク先はZIPアーカイブとなっていますが、
EXEインストーラーも [Releases][rels] より入手可能です。
必要な方はそちらも参照してみてください。

なおこれまでのリンク先だった+kaoriya版については [vim-kaoriya][kaoriya] にて引き続き配布しています。

[vwi]:https://github.com/vim/vim-win32-installer
[ci]:https://ci.appveyor.com/project/chrisbra/vim-win32-installer
[rels]:https://github.com/vim/vim-win32-installer/releases
[kaoriya]:https://github.com/koron/vim-kaoriya/releases
[koron]:https://twitter.com/kaoriya
