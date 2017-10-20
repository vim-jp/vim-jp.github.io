---
layout: docs
title: MinGWを使ってのビルド方法
---

ビルド手順はVimのバージョンによって一部変更されていることがあります。公式のビルド手順は、ソースファイル内の [src/INSTALLpc.txt](https://github.com/vim/vim/blob/master/src/INSTALLpc.txt) および、[MinGW用Makefile](https://github.com/vim/vim/blob/master/src/Make_ming.mak)や[Cygwin/MinGW共通Makefile](https://github.com/vim/vim/blob/master/src/Make_cyg_ming.mak)内のコメントを参照してください。

1.  msys2 のインストール

    以前は Windows 上の gcc コンパイラをインストールする為に [MinGW \| Minimalist GNU for Windows](http://www.mingw.org/) をダウンロードしましたが、最近は <a href="https://msys2.github.io/">msys2</a> を使うのが一般的です。以下は msys2 がインストール済みの前提で記述しています。

2.  GitやGCCのインストール

    スタートメニューから `MSYS2 MSYS` を実行し、以下のコマンドを実行してシステムの更新を行います。

        pacman -Suuy

    MSYS2を終了して再実行するように言われた場合は、それに従ってminttyのウィンドウを閉じ、スタートメニューから `MSYS2 MSYS` を実行し、再度 `pacman -Suuy` を実行します。
    システムの更新が終わったら、以下のコマンドを実行してGitやGCCなど必要なパッケージをインストールします。

        pacman -S git mingw-w64-i686-toolchain mingw-w64-x86_64-toolchain

    pacmanのwrapperコマンドであるpacboyを使えば、もう少し簡単に、以下のコマンドでインストールできます。

        pacboy -S git: toolchain:m

    pacboyの使い方は、当該コマンドを引数無しで実行すると表示されます。`:`でプレフィックス無し、`:x`で64bit用(`mingw-w64-x86_64-`)、`:i`で32bit用(`mingw-w64-i686-`)、`:m`で64/32bit用両方がインストールされます。

3.  ソース取得

    ビルドするターゲットに合わせて、スタートメニューから `MSYS2 MinGW 64-bit` または `MSYS2 MinGW 32-bit` のいずれかを実行し、そこから以下を実行します。

        git clone https://github.com/vim/vim.git

    `git clone`を実行した後にソースが更新された場合は、以下のコマンドで最新のソースを取得できます。

        git pull

    特定のバージョンを指定して取得する場合は、以下のコマンドを実行します。

        git checkout v7.4.393

4.  コンパイル

    `vim/src`フォルダに移動し以下のコマンドを実行します。

        mingw32-make -f Make_ming.mak GUI=yes IME=yes MBYTE=yes \
          ICONV=yes DEBUG=no

    もしPerl拡張やPython2/3拡張、Ruby拡張、Lua拡張を使う場合は以下の様に指定します。

        mingw32-make -f Make_ming.mak GUI=yes IME=yes MBYTE=yes ICONV=yes \
          PERL=c:/strawberry/perl DYNAMIC_PERL=yes PERL_VER=526 \
          PYTHON=c:/python27 DYNAMIC_PYTHON=yes PYTHON_VER=27 \
          PYTHON3=c:/python36 DYNAMIC_PYTHON3=yes PYTHON3_VER=36 \
          RUBY=c:/Ruby24 DYNAMIC_RUBY=yes RUBY_VER=24 RUBY_API_VER_LONG=2.4.0 \
          LUA=c:/lua53 DYNAMIC_LUA=yes LUA_VER=53 \
          CSCOPE=yes NETBEANS=yes STATIC_STDCPLUS=yes DEBUG=no

    msys2 付属の perl/python/ruby/lua を使う場合は以下の様に指定します。

        mingw32-make -f Make_ming.mak GUI=yes IME=yes MBYTE=yes ICONV=yes \
          PERL=c:/msys64/mingw64 DYNAMIC_PERL=yes PERL_VER=522 \
            PERLEXE=c:/msys64/mingw64/bin/perl.exe \
            PERLLIB=c:/msys64/mingw64/lib/perl5/core_perl \
          PYTHON=c:/msys64/mingw64 DYNAMIC_PYTHON=yes PYTHON_VER=27 \
            PYTHON_HOME=c:/msys64/mingw64 \
            PYTHONINC=-Ic:/msys64/mingw64/include/python2.7 \
            DYNAMIC_PYTHON_DLL=libpython2.7.dll \
          PYTHON3=c:/msys64/mingw64 DYNAMIC_PYTHON3=yes PYTHON3_VER=36 \
            PYTHON3_HOME=c:/msys64/mingw64/ \
            PYTHON3INC=-Ic:/msys64/mingw64/include/python3.6m \
            DYNAMIC_PYTHON3_DLL=libpython3.6m.dll \
          RUBY=c:/msys64/mingw64 DYNAMIC_RUBY=yes RUBY_VER=24 RUBY_API_VER_LONG=2.4.0 \
          LUA=c:/msys64/mingw64 DYNAMIC_LUA=yes LUA_VER=53 \
          CSCOPE=yes NETBEANS=yes STATIC_STDCPLUS=yes DEBUG=no

    注意1) 複数行に跨るのでシェルスクリプトに記載して実行して下さい。

    注意2) Vimのバージョンによっては、`ARCH=x86-64` (64bit) または `ARCH=i686` (32bit) を指定しなければならない場合があります。

    注意3) ターゲットが32bitの場合は、`c:/msys64/mingw64` を `c:/msys64/mingw32` に置き換えてください。

    Vim 7.4.393以降で使えるようになったDirectWriteを有効にするには、`DIRECTX=yes` を指定する必要があります。
    なお、DirectWriteを有効にすると、デフォルトでは実行時に libstdc++-6.dll, libgcc\_s\_sjlj-1.dll 等が必要となりますが、`STATIC_STDCPLUS=yes` を指定することでこれらのライブラリをスタティックリンクすることが出来ます。(VimやGCCのバージョンによって出来ないこともあり。)
