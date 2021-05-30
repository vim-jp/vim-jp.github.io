---
layout: docs
title: パッチの書き方
---

パッチを書いたことがない、または他の人のパッチを試したことがない皆さんにとって、パッチというのは敷居の高いことかもしれません。
しかし、一度Vimにパッチを書いてみたり、他人のパッチを試してみると、パッチを扱うことが空気のように思えてくるはずです。
Vimはバージョン管理ツールとして Git を採用しているおかげで、簡単にパッチを作成し管理することができます。
ここではパッチの作成方法について学びましょう。

最近、www.vim.org のページに[Vim development](http://www.vim.org/develop.php)という記事ができました。
こちらは英語ですが、手順がよくまとまっているため、英語を苦にしない人はこちらを参照するとよいでしょう。

## 準備

まず前提条件として、あなたのPCではコンパイル環境が整備されており、[Git](https://git-scm.com/)が使える状態になっているとします。
Linuxユーザーであれば[Linuxでのビルド方法](http://vim-jp.org/docs/build_linux.html)を、Windowsユーザーであれば[Visual Studioでのコンパイル方法](http://vim-jp.org/docs/build_windows_msvc.html)もしくは[MinGWでのコンパイル方法](http://vim-jp.org/docs/build_windows_mingw.html)も参照して下さい。

最新のソースコードを取得する為に、シェル(Windowsであればコマンドプロンプト)から以下を実行します。

    $ git clone https://github.com/vim/vim.git

## 修正

修正内容に対応するブランチを作成します。

    $ git switch -c add-new-func-trim

(git のバージョンが古い場合は `switch -c` の代わりに `checkout -b` を使います。)

ブランチを移動したらソースファイルを修正します。

    $ vim evalfunc.c

この際コーディングスタイルは既存のものに合わせるようにしましょう。
[`:help coding-style`](http://vim-jp.org/vimdoc-ja/develop.html#coding-style)にコーディング規約が書かれてあります。
また[`:help develop.txt`](http://vim-jp.org/vimdoc-ja/develop.html)にはVimの開発に関する事項が書かれています。一度目を通すといいでしょう。

すでに用意されているパッチをソースファイルに適用するなら、以下のようなコマンドを実行します。

    $ git apply something.patch

ソースファイルを修正したらコミットしましょう。

    $ git commit -m "added new function trim()"

もし、修正中に元のソースコードが変更された場合は、次のコマンドでリポジトリを同期させます。

    $ git fetch origin master:master
    $ git rebase master

## テスト

コードを修正したならテストも書きましょう。不具合修正であれば、同じ不具合を再発させないためにもテストは重要です。テストの書き方は以下を参照してください。

- [`:help testing`](http://vim-jp.org/vimdoc-ja/testing.html#testing)
- [src/testing/README.txt](https://github.com/vim/vim/blob/master/src/testing/README.txt)

### テストの実行方法

テストを実行するにはいくつかのやり方がありますが、`src/testing/` ディレクトリ内で make を実行するのが良いでしょう。

Unix 系の場合:

    全テストを実行
    $ make

    特定のファイルのテストを実行
    $ make test_channel

    特定のファイルの特定名称のテストを実行 (部分一致)
    $ make test_channel TEST_FILTER=Test_communicate

    GUI で実行
    $ make GUI_FLAG=-g

Windows の場合 (MSVC):

    全テストを実行
    > nmake -f Make_dos.mak

    特定のファイルのテストを実行
    > nmake -f Make_dos.mak test_channel

    特定のファイルの特定名称のテストを実行 (部分一致)
    > nmake -f Make_dos.mak test_channel TEST_FILTER=Test_communicate

    GUI で実行
    > nmake -f Make_dos.mak VIMPROG=..\gvim

## 送信

全ての修正が完了したら pull-request を送信しましょう。GitHub 上で vim のリポジトリから Fork ボタンをクリックするか、[hub](https://github.com/github/hub) コマンドをお持ちであれば `hub fork` でも可能です。fork が出来たら自分のリポジトリに対して push します。

    $ hub fork
    $ git push [ユーザーID] HEAD -u

push すると以下のような表示が出ます。

    remote:
    remote: Create a pull request for '[ブランチ名]' on GitHub by visiting:
    remote:      https://github.com/[ユーザーID]/vim/pull/new/[ブランチ名]
    remote:


表示された URL をブラウザで開くと pull-request の作成画面が出るので指示に従って pull-request を作成して下さい。その際、海外の方も開発に参加していますので本文は英語で書いて頂く必要があります。

pull-request の冒頭で

- 発生している現象
- 再現手順

を説明します。

パッチを作成してメーリングリスト[vim\_dev](https://groups.google.com/forum/#!forum/vim_dev)に送信する方法もありますが、CI が流れないので推奨はしません。(vim\_dev初回投稿時はBram氏の承認が必要です。投稿が表示されないからと何度も同じ投稿をしないように注意してください。)

あとは誰かがレビューしてくれるはずです。パッチを投げる前に、vim-jpの[Issues](http://github.com/vim-jp/issues/issues)で相談すると、歴戦の強者達がアドバイスを送ってくれることでしょう。
なおパッチファイルにはパッチ作成者の名前が記されますが、基本的には実名が相応しいと以前Bram氏が言っていました。

## マージ

レビューが完了しBram氏からのOKが出たら、あとはパッチが取り込まれたVimがリリースされるまで待ちます。
いつパッチが取り込まれるかは不定期なので分かりません。すぐに取り込まれる場合もありますし、[todo.txt](https://github.com/vim/vim/blob/master/runtime/doc/todo.txt)に追加されて数ヶ月あるいは数年待たされる場合もあります。(必要性をアピールしたり)パッチをメンテナンスしつつ、ひたすら待ちましょう。
Vimでは基本的にバグ修正が優先して取り込まれ、新機能の追加は後回しにされる傾向が強いです。
特に巨大な変更の場合、パッチの取り込みがマイナーリリース時点まで保留されることもあります。

パッチが取り込まれたらローカルのブランチを削除しても構いません。

    $ git checkout master
    $ git branch -D add-new-func-trim

で削除できます。あとは最新版のVimを Git で取ってきてコンパイルするだけです。

さぁ君もCONTRIBUTE AUTHORになろう！

## 参考資料

- [Vim development](http://www.vim.org/develop.php)
