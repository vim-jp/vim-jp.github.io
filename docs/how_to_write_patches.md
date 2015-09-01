---
layout: docs
title: パッチの書き方
---

パッチを書いたことがない、または他の人のパッチを試したことがない皆さんにとって、パッチというのは敷居の高いことかもしれません。
しかし、一度Vimにパッチを書いてみたり、他人のパッチを試してみると、パッチを扱うことが空気のように思えてくるはずです。
Vimはバージョン管理ツールとしてMercurialを採用しているおかげで、簡単にパッチを作成し管理することができます。
ここではパッチの作成方法について学びましょう。

最近、www.vim.org のページに[Vim development](http://www.vim.org/develop.php)という記事ができました。
こちらは英語ですが、手順がよくまとまっているため、英語を苦にしない人はこちらを参照するとよいでしょう。

まず前提条件として、あなたのPCではコンパイル環境が整備されており、[Mercurial](http://mercurial.selenic.com/)が使える状態になっているとします。
Windowsユーザーであれば[Visual Studioでのコンパイル方法](http://vim-jp.org/docs/build_windows_msvc.html)もしくは[MinGWでのコンパイル方法](http://vim-jp.org/docs/build_windows_mingw.html)も参照して下さい。

最新のソースコードを取得する為に、シェル(Windowsであればコマンドプロンプト)から以下を実行します。

    $ hg clone https://bitbucket.org/vim-mirror/vim

今回の手順では [Mercurial Queue](http://mercurial.selenic.com/wiki/JapaneseMqExtension)を利用するので、あらかじめ機能を有効化しておきましょう。

    $ cat ~/.hgrc
    [extensions]
    mq=

次に修正内容に対応するパッチを作成します。
今回はパッチ名を`010-fix-quickfix-on-windows.patch`としました。

    $ hg qnew 010-fix-quickfix-on-windows.patch

パッチが作成されたらソースファイルを修正します。

    $ vim os_win32.c

この際コーディングスタイルは既存のものに合わせるようにしましょう。
[`:help coding-style`](http://vim-jp.org/vimdoc-ja/develop.html#coding-style)にコーディング規約が書かれてあります。
また[`:help develop.txt`](http://vim-jp.org/vimdoc-ja/develop.html)にはVimの開発に関する事項が書かれています。一度目を通すといいでしょう。

すでに用意されているパッチをソースファイルに適用するなら、以下のようなコマンドを実行します。

    $ patch -p1 < something.patch

ソースファイルを修正したらパッチを更新しましょう。

    $ hg qrefresh

もし、修正中に元のソースコードが変更された場合は、次のコマンドでリポジトリを同期させます。

    $ hg qpop -a; hg pull -u; hg qpush -a

全ての修正が完了したらパッチファイルを作りましょう。最終的には全修正を一連のパッチファイルとしてメーリングリストに投げます。

    $ hg qdiff > fix-quickfix-on-windows.diff

これで作成したパッチの差分ファイルが作成出来るので、メーリングリスト[vim\_dev](https://groups.google.com/forum/#!forum/vim_dev)に添付ファイルを付けるか本文に貼り付けてメールを投げます。(vim\_dev初回投稿時はBram氏の承認が必要なため、すぐには表示されません。)
メールの冒頭で

- 発生している現象
- 再現手順

を説明します。
あとは誰かがレビューしてくれるはずです。パッチを投げる前に、vim-jpの[Issues](http://github.com/vim-jp/issues/issues)で相談すると、歴戦の強者達がアドバイスを送ってくれることでしょう。
なおパッチファイルにはパッチ作成者の名前が記されますが、基本的には実名が相応しいと以前Bram氏が言っていました。

レビューが完了しBram氏からのOKが出たら、あとはパッチが取り込まれたVimがリリースされるまで待ちます。
いつパッチが取り込まれるかは不定期なので分かりません。すぐに取り込まれる場合もありますし、[todo.txt](https://github.com/vim/vim/blob/master/runtime/doc/todo.txt)に追加されて数ヶ月あるいは数年待たされる場合もあります。ひたすら待ちましょう。
Vimでは基本的にバグ修正が優先して取り込まれ、新機能の追加は後回しにされる傾向が強いです。
特に巨大な変更の場合、パッチの取り込みがマイナーリリース時点まで保留されることもあります。

パッチが取り込まれ、パッチが不要になったら

    $ hg qdelete

で削除できます。あとは最新版のVimをMercurialで取ってきてコンパイルするだけです。

さぁ君もCONTRIBUTE AUTHORになろう！

※ここで紹介しているMercurial Queueは[quilt](http://savannah.nongnu.org/projects/quilt)と呼ばれるパッチ管理ソフトウェアのコマンド体系をベースにしています。
Mercurial Queueの使い方を勉強するときには、まずquiltについて学んでおくと理解が早いでしょう。

参考資料：

- [Mercurial Queues エクステンション](http://mercurial.selenic.com/wiki/JapaneseMqExtension)
- [Mercurial MQ について](http://d.hatena.ne.jp/dayflower/20090520/1242794877)
- [12 Managing change with Mercurial Queues](http://foozy.bitbucket.org/hgbook-ja/d6ca1334a19d/hgbookch12.html#x105-26500012)
- [13 Advanced uses of Mercurial Queues](http://foozy.bitbucket.org/hgbook-ja/d6ca1334a19d/hgbookch13.html#x108-30000013)
- [Vim development](http://www.vim.org/develop.php)
