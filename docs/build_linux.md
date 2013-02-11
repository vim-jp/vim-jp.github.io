---
layout: docs
title: Linuxでのビルド方法
---

Ubuntu 12.04 LTSを使った場合のビルド方法を説明します。

1.  必要なパッケージのインストール

    Terminalアプリを開き、以下を実行します。ビルドに必要なパッケージが全てインストールされますが、ビルド設定によっては不要なパッケージも大量に含まれています。（行頭の`$`はプロンプトを示しており、実際には入力不要です。）

        $ sudo apt-get build-dep vim

    パッケージを個別にインストールするには以下を実行します。

        $ sudo apt-get install mercurial gettext libncurses5-dev

    GVim (GTK2-GNOME GUI版)をビルドするには以下も追加で必要です。

        $ sudo apt-get install libxmu-dev libgnomeui-dev

    ※GTK2 GUI版の場合は`libgnomeui-dev`の代わりに、`libgtk2.0-dev`を指定。

    Perl, Python2,3, Ruby拡張を使うには以下も追加で必要です。

        $ sudo apt-get install libperl-dev python-dev python3-dev ruby-dev

    ソースコードを修正する場合は、以下のパッケージも必要になることがあります。

        $ sudo apt-get install autoconf automake cproto

2.  ソース取得

    以下のコマンドを実行します。

        $ hg clone https://vim.googlecode.com/hg/ vim

    `hg clone`を実行した後にソースが更新された場合は、以下のコマンドで最新のソースを取得できます。

        $ hg pull
        $ hg update

    あるいは2つをまとめて、以下のコマンドでもOKです。

        $ hg pull -u

3.  コンパイル

    `vim/src`フォルダに移動し以下のコマンドを実行します。

        $ ./configure --with-features=huge --enable-gui=gnome2
          --enable-fail-if-missing
        $ make

    ※`./configure`の行は実際は1行  
    ※GTK2 GUI版の場合は`--enable-gui=gnome2`の代わりに、`--enable-gui=gtk2`を指定  
    ※`--enable-fail-if-missing`は足りないパッケージがある場合にエラーとするためのオプション  

    もしPerl拡張やRuby拡張、Python拡張を使う場合は以下の様に指定します。

        $ ./configure --with-features=huge --enable-gui=gnome2
          --enable-perlinterp --enable-pythoninterp
          --enable-python3interp --enable-rubyinterp
          --enable-fail-if-missing
        $ make

    ※`./configure`の行は実際は1行

    `./configure`のオプションの詳細は以下のコマンドで確認できます。

        $ ./configure --help

    これらの`./configure`のオプションは、`vim/src/Makefile`を編集することでも設定可能です。この場合は、単純に`make`コマンドを実行することでビルドができます。オプションを変更後ビルドし直す場合は、以下のコマンドを実行します。

        $ make reconfig

4.  インストール

    以下のコマンドを実行すると、`./configure`の`--prefix`オプションで指定した先にインストールされます。（無指定の場合は、`/usr/local`など）

        $ sudo make install
