---
layout: docs
title: VisualStudioを使ってのビルド方法
---

ビルド手順はVimのバージョンによって一部変更されていることがあります。公式のビルド手順は、ソースファイル内の [src/INSTALLpc.txt](https://github.com/vim/vim/blob/master/src/INSTALLpc.txt) および、[VC用Makefile](https://github.com/vim/vim/blob/master/src/Make_mvc.mak)内のコメントを参照してください。

ここではVisualStudio(2010以降)が既にインストールされている前提で説明します。

1.  Gitのインストール

    [Git](https://git-scm.com/)からGitをインストールします。

2.  ソース取得

    コマンドプロンプトから以下を実行します。

        git clone https://github.com/vim/vim.git

    `git clone`を実行した後にソースが更新された場合は、以下のコマンドで最新のソースを取得できます。

        git pull

    特定のバージョンを指定して取得する場合は、以下のコマンドを実行します。

        git checkout v7.4.393

3.  コンパイル

    必要に応じてビルド前に`vcvarsall.bat`や`vcvars32.bat`等を実行して、VCを動かすための環境変数をセットアップして下さい。`vim/src`フォルダ内にある`msvc****.bat`が使える場合もあります。

    `vim/src`フォルダに移動し以下のコマンドを実行します。

        nmake -f Make_mvc.mak GUI=yes IME=yes MBYTE=yes ^
          ICONV=yes DEBUG=no

    もしPerl拡張やPython2/3拡張などを使う場合は以下の様に指定します。

        nmake -f Make_mvc.mak GUI=yes IME=yes MBYTE=yes ICONV=yes ^
          PERL=C:\strawberry\perl DYNAMIC_PERL=yes PERL_VER=526 ^
          PYTHON=c:\python27 DYNAMIC_PYTHON=yes PYTHON_VER=27 ^
          PYTHON3=c:\python36 DYNAMIC_PYTHON3=yes PYTHON3_VER=36 ^
          CSCOPE=yes NETBEANS=yes DEBUG=no

    注意) 複数行に跨るのでバッチファイルに記載して実行して下さい。(行末の `^` は行継続を示しています。)

    Ruby拡張を使うには、少々面倒な手順が必要になります。通常はRubyとVimを同じコンパイラでビルドする必要があるのですが、現在Windowsで広く使われている[RubyInstaller](https://rubyinstaller.org/)はMinGWでビルドされているので、そのままではVCから使うことができないのです。
    詳細は、[src/INSTALLpc.txt](https://github.com/vim/vim/blob/master/src/INSTALLpc.txt)の "Building with Ruby support" の項を参照してください。

    Vim 7.4.393以降で使えるようになったDirectWriteを有効にするには、`DIRECTX=yes` を指定する必要があります。

    VC2012以降を使う場合は、`SDK_INCLUDE_DIR` で `Win32.mak` があるフォルダを指定しなければならない場合があります。例:

        "SDK_INCLUDE_DIR=C:\Program Files (x86)\Microsoft SDKs\Windows\v7.1A\Include"

    `SDK_INCLUDE_DIR` の指定が必要かどうかは、Vimのバージョンに依存します。
