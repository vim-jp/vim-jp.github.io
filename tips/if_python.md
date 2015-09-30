---
layout: tips
title: Linuxでif_pythonを使う際のtips
---

Linux上のVimでPython2とPython3を両方使用する際に注意が必要です。

Debian系Linuxでapt-getなどを利用してインストールしたPython2,Python3を利用した場合、VimからPython2とPython3を同時に使用する事はできません。

    :py print(sys.version)
    :2.7.3 (...
    :py3 print(sys.version)
    E837: ...

たとえば「A」というプラグインがPython2を利用している「B」というプラグインがPython3を利用している場合`E836`もしくは`E837`が発生する可能性があります。

VimのヘルプのE836,E837にPython2とPython3を同時に利用できない理由が書いてあります。
[Vim documentation: if_pyth](http://vim-jp.org/vimdoc-ja/if_pyth.html#E836)

なぜDebian系のLinuxでは利用できないかは以下のリンク先にある通り、Pythonがstatic linkで作られるのと同時に.soも提供されているからです。

[Debian系のLinuxでPython 2.xと3.xが同時利用できない問題の原因と対策 · Issue #301 · vim-jp/issues](https://github.com/vim-jp/issues/issues/301)

Pythonを`--enable-shared `でビルドしVimに組み込むと利用できるようになりますが、`--enable-shared`あり/なしを比較した場合、Pythonが遅くなるようです。

- [Python 3.3.0 amd64 core2duo debian 7.0 --enable-shared ありと無しの比較](https://gist.github.com/4681131)
- [FreeBSD i386 結構、古いCPUの上のpybenchの結果、sharedとstaticを比べたもの。 sharedは 20% も遅くなる。](https://gist.github.com/4681320)


Linux上でVimからPython2とPython3を同時に使いたい場合

- Pythonを`--enable-shared`でビルドしVimに組み込む(ただしPythonが遅くなる可能性が高い)
- Python2を利用しているプラグインをPython3で動くように修正して貰う(あるいは修正する)

などの対応が考えられます。

なおKaoriya版Vimでは問題なくPython2とPython3が同時に利用できるようです。

- [Windowsについて](http://www.kaoriya.net/blog/2013/01/30)
- [MacVim Kaoriya](https://github.com/splhack/macvim-kaoriya/wiki/Readme#perl--python--python3--ruby--lua)
