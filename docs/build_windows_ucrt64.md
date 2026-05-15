---
layout: docs
title: MSYS2 UCRT64を使ってのビルド方法
---

ビルド手順はVimのバージョンによって一部変更されていることがあります。公式のビルド手順は、ソースファイル内の [src/INSTALLpc.txt](https://github.com/vim/vim/blob/master/src/INSTALLpc.txt) および、[MinGW用Makefile](https://github.com/vim/vim/blob/master/src/Make_ming.mak)や[Cygwin/MinGW共通Makefile](https://github.com/vim/vim/blob/master/src/Make_cyg_ming.mak)内のコメントを参照してください。

UCRT64 環境は、従来の MSVCRT ではなく Universal C Runtime (UCRT) に対してリンクする MSYS2 のサブシステムです。Windows 10 以降では UCRT が標準で搭載されており、新規にビルドする場合は MinGW64 (MSVCRT) よりも UCRT64 を選択することが推奨されています。

1.  msys2 のインストール

    <a href="https://www.msys2.org/">msys2</a> をダウンロードしてインストールします。以下は msys2 がインストール済みの前提で記述しています。

2.  Git・GCCのインストール

    スタートメニューから `MSYS2 MSYS` を実行し、以下のコマンドを実行してシステムの更新を行います。

        pacman -Suuy

    MSYS2を終了して再実行するように言われた場合は、それに従ってminttyのウィンドウを閉じ、スタートメニューから再度 `MSYS2 MSYS` を実行し、`pacman -Suu` を実行します。
    システムの更新が終わったら、以下のコマンドを実行してGitやGCCなど必要なパッケージをインストールします。

        pacman -S git mingw-w64-ucrt-x86_64-toolchain

    pacmanのwrapperコマンドであるpacboyを使えば、もう少し簡単に、以下のコマンドでインストールできます。

        pacboy -S git toolchain:u

    pacboyの使い方は、`pacboy help`を実行すると表示されます。パッケージ名の後に`:u`でUCRT64用(`mingw-w64-ucrt-x86_64-`)、`:x`でMinGW64用(`mingw-w64-x86_64-`)、`:c`でClang64用(`mingw-w64-clang-x86_64-`)のパッケージがインストールされます。

3.  ソース取得

    スタートメニューから `MSYS2 UCRT64` を実行し、そこから以下を実行します。

        git clone https://github.com/vim/vim.git

    `git clone`を実行した後にソースが更新された場合は、以下のコマンドで最新のソースを取得できます。

        git pull

    特定のバージョンを指定して取得する場合は、以下のコマンドを実行します。

        git checkout v9.1.0000

4.  コンパイル

    `vim/src`フォルダに移動し以下のコマンドを実行します。

        mingw32-make -f Make_ming.mak GUI=yes IME=yes MBYTE=yes \
          ICONV=yes DEBUG=no

    Python3拡張やXPMサポートなどを有効にする場合は、以下の様に指定します(GUI版)。

        mingw32-make -f Make_ming.mak GUI=yes IME=yes MBYTE=yes ICONV=yes \
          PYTHON3=c:/msys64/ucrt64 DYNAMIC_PYTHON3=yes PYTHON3_VER=314 \
            PYTHON3_HOME=c:/msys64/ucrt64/ \
            PYTHON3INC=-Ic:/msys64/ucrt64/include/python3.14 \
            DYNAMIC_PYTHON3_DLL=libpython3.14.dll \
          CSCOPE=yes NETBEANS=yes STATIC_STDCPLUS=yes DEBUG=no \
          XPM=c:/msys64/ucrt64

    CUI版(コンソール版)を同時にビルドする場合は、`GUI=no` を指定してもう一度実行します。

        mingw32-make -f Make_ming.mak GUI=no IME=yes MBYTE=yes ICONV=yes \
          PYTHON3=c:/msys64/ucrt64 DYNAMIC_PYTHON3=yes PYTHON3_VER=314 \
            PYTHON3_HOME=c:/msys64/ucrt64/ \
            PYTHON3INC=-Ic:/msys64/ucrt64/include/python3.14 \
            DYNAMIC_PYTHON3_DLL=libpython3.14.dll \
          CSCOPE=yes NETBEANS=yes STATIC_STDCPLUS=yes DEBUG=no \
          XPM=c:/msys64/ucrt64

    Perl・Ruby・Lua拡張も含めて、msys2 付属のパッケージを利用する場合は以下の様に指定します。バージョン番号(`PERL_VER`等)はインストールされているパッケージに合わせて適宜変更してください。

        mingw32-make -f Make_ming.mak GUI=yes IME=yes MBYTE=yes ICONV=yes \
          PERL=c:/msys64/ucrt64 DYNAMIC_PERL=yes PERL_VER=540 \
            PERLEXE=c:/msys64/ucrt64/bin/perl.exe \
            PERLLIB=c:/msys64/ucrt64/lib/perl5/core_perl \
          PYTHON3=c:/msys64/ucrt64 DYNAMIC_PYTHON3=yes PYTHON3_VER=314 \
            PYTHON3_HOME=c:/msys64/ucrt64/ \
            PYTHON3INC=-Ic:/msys64/ucrt64/include/python3.14 \
            DYNAMIC_PYTHON3_DLL=libpython3.14.dll \
          RUBY=c:/msys64/ucrt64 DYNAMIC_RUBY=yes RUBY_VER=34 RUBY_API_VER_LONG=3.4.0 \
          LUA=c:/msys64/ucrt64 DYNAMIC_LUA=yes LUA_VER=54 \
          CSCOPE=yes NETBEANS=yes STATIC_STDCPLUS=yes DEBUG=no \
          XPM=c:/msys64/ucrt64

    注意1) 複数行に跨るのでシェルスクリプトに記載して実行して下さい。Windowsのバッチファイルで実行する場合は行末のバックスラッシュ(`\`)をキャレット(`^`)に置き換えてください。

    注意2) `PYTHON3_VER`、`PYTHON3INC`、`DYNAMIC_PYTHON3_DLL` 等の数字は、インストールされているPythonのバージョンに合わせてください。バージョンは `pacman -Qi mingw-w64-ucrt-x86_64-python` で確認できます。Perl・Ruby・Lua も同様です。

    注意3) `XPM=c:/msys64/ucrt64` を指定する場合は、事前に `pacboy -S xpm-nox:u` で XPM ライブラリをインストールしておく必要があります。

    Vim 7.4.393以降で使えるようになったDirectWriteを有効にするには、`DIRECTX=yes` を指定する必要があります。
    なお、DirectWriteを有効にすると、デフォルトでは実行時に libstdc++-6.dll, libgcc\_s\_seh-1.dll 等が必要となりますが、`STATIC_STDCPLUS=yes` を指定することでこれらのライブラリをスタティックリンクすることが出来ます。(VimやGCCのバージョンによって出来ないこともあり。)
