---
layout: docs
title: Linuxでのビルド方法
---

Ubuntu 14.04 LTSを使った場合のビルド方法を説明します。

1.  必要なパッケージのインストール

    Terminalアプリを開き、以下を実行します。ビルドに必要なパッケージが全てインストールされますが、ビルド設定によっては不要なパッケージも大量に含まれています。（行頭の`$`はプロンプトを示しており、実際には入力不要です。）

        $ sudo apt-get build-dep vim

    パッケージを個別にインストールするには以下を実行します。

        $ sudo apt-get install git gettext libncurses5-dev
          libacl1-dev libgpm-dev

    ※実際は1行

    gvim (GTK2 GUI版)をビルドするには以下も追加で必要です。(GUI版は一般的にはGTK2またはGTK3を使うのがよいでしょう。)

        $ sudo apt-get install libxmu-dev libgtk2.0-dev libxpm-dev

    ※GTK3 GUI版の場合は`libgtk2.0-dev`の代わりに、`libgtk-3-dev`を指定。<br />
    ※GTK2-GNOME GUI版の場合は`libgtk2.0-dev`の代わりに、`libgnomeui-dev`を指定。<br />

    Perl, Python2,3, Ruby拡張を使うには以下も追加で必要です。

        $ sudo apt-get install libperl-dev python-dev python3-dev ruby-dev

    Lua拡張を使うには以下も追加で必要です。

        $ sudo apt-get install lua5.2 liblua5.2-dev

    LuaJITのLua拡張を使うには以下も追加で必要です。

        $ sudo apt-get install luajit libluajit-5.1

    ソースコードを修正する場合は、以下のパッケージも必要になることがあります。

        $ sudo apt-get install autoconf automake cproto

2.  ソース取得

    以下のコマンドを実行します。

        $ git clone https://github.com/vim/vim.git

    `git clone`を実行した後にソースが更新された場合は、以下のコマンドで最新のソースを取得できます。

        $ git pull

    特定のバージョンを指定して取得する場合は、以下のコマンドを実行します。

        $ git checkout v7.4.393

3.  コンパイル

    `vim/src`フォルダに移動し以下のコマンドを実行します。

        $ ./configure --with-features=huge --enable-gui=gtk2
          --enable-fail-if-missing
        $ make

    ※`./configure`の行は実際は1行<br />
    ※GTK3 GUI版の場合は`--enable-gui=gtk2`の代わりに、`--enable-gui=gtk3`を指定<br />
    ※GTK2-GNOME GUI版の場合は`--enable-gui=gtk2`の代わりに、`--enable-gui=gnome2`を指定<br />
    ※`--enable-fail-if-missing`は足りないパッケージがある場合にエラーとするためのオプション<br />

    もしPerl拡張やPython2/3拡張、Ruby拡張、Lua拡張を使う場合は以下の様に指定します。

        $ ./configure --with-features=huge --enable-gui=gtk2
          --enable-perlinterp --enable-pythoninterp
          --enable-python3interp --enable-rubyinterp
          --enable-luainterp
          --enable-fail-if-missing
        $ make

    もしLuaインタプリタとしてLuaJITを利用したい場合は以下の様に指定します。(上記に加えて`--with-luajit`を指定している点に注意)

        $ ./configure --with-features=huge --enable-gui=gtk2
          --enable-perlinterp --enable-pythoninterp
          --enable-python3interp --enable-rubyinterp
          --enable-luainterp --with-luajit
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
