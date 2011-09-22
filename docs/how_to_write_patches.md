---
layout: docs
title: パッチの書き方
---

*追記*
> mercurialでテンポラリなブランチを作るのであれば、branchよりもmercurial queueもしくはbookmarkを使った方が良いとのご意見を頂きました。

ここでは既に、コンパイル環境や[Mercurial](http://mercurial.selenic.com/)が使える状態になっている事を前提とします。
なお、Windowsユーザであれば[Visual Studioでのコンパイル方法](http://vim-jp.org/docs/build_windows_msvc.html)もしくは[MinGWでのコンパイル方法](http://vim-jp.org/docs/build_windows_mingw.html)も参照して下さい。

最新のソースコードを取得する為に、シェル(Windowsであればコマンドプロンプト)から以下を実行します。

    hg clone https://code.google.com/p/vim/

次に修正内容に対応するブランチを作成します。

    hg branch fix-quickfix-on-windows

ここで気を付けておいた方がいいのが、`hg branch`する際のブランチは常に`default`であった方が良いという事です。
これはvimに限った話ですが、authorであるBram Moolenaar氏は1つのパッチに対して1チェンジセットでcommitを作りますが、パッチファイルとして提供する際にはvim\_devメーリングリストへ投げたパッチに対して少し変更を少し変えてパッチリリースする事があるからです。その変更の中にはvimのパッチ番号でもある`7.3.322`の`322`をソースコードに埋め込む作業も行われます。つまりパッチが取り入れられて`default`ブランチが更新された場合でマージが正常に行く事があまりないのです。例えば作成したブランチから作った派生ブランチが壊れてしまう可能性もあります。
よってブランチは1つの修正に対して1回限りの物と思って良いでしょう。

ブランチが作成されたらソースファイルを修正します。

    vim os_win32.c

修正したらcommitしましょう。

    hg commit

全ての修正が完了したらパッチを作りましょう。最終的には全修正を一連のパッチファイルとしてメーリングリストに投げます。

    hg diff -r default:fix-quickfix-on-windows >
	  fix-quickfix-on-windows.diff

*※実際は1行*

これでブランチ間の差分ファイルが作成出来るので、メーリングリスト[vim\_dev](https://groups.google.com/forum/#!forum/vim_dev)に添付ファイルを付けるか本文に貼り付けてメールを投げます。
メールの冒頭で

- 発生している現象
- 再現手順

を説明します。
あとは誰かがレビューしてくれるはずです。
なお、パッチファイルには名前が記されますが基本的には実名が相応しいと以前Bram氏が言っていました。

さぁ君もCONTRIBUTE AUTHORになろう！
