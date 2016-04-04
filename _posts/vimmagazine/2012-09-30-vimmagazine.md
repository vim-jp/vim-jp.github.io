---
layout: post
category: vimmagazine
title: Vim Magazine 9 月号
---

## 話題

- [マンガで分かる！Vimの使い方 1 「Vimって何？」 - ナレッジエース](http://n.blueblack.net/articles/2012-08-29_01_vim_comic_01_what_is_vim/)

- [マンガで分かる！Vimの使い方 2 「VimのインストールとTutorial」 - ナレッジエース](http://n.blueblack.net/articles/2012-09-01_01_vim_comic_02_vim_install/)

- [マンガで分かる！Vimの使い方 3 「Vimの設定」 - ナレッジエース](http://n.blueblack.net/articles/2012-09-15_01_vim_comic_03_vim_setting/)

- [watchdogs.vim つくりました - C++でゲームプログラミング](http://d.hatena.ne.jp/osyo-manga/20120924/1348473304)

  Vim でシンタックスチェックを行うためのプラグイン。非同期処理対応で Vim
  のフォーカスが奪われることがなく、自動シンタックスチェックを行うことが
  出来る。

- [TokyoVim#11](http://partake.in/events/da3735cd-6d0e-48b6-8010-18cc031702c4)

  Vimに関する何かをやる会 東京 10月6日(土)


## リリース情報

- [7.3.647 : "gnd" doesn't work correctly in Visual mode](https://github.com/vim/vim/commit/57c0ea86927c2368f980e7989565c7a708d96526)
- [7.3.648 : crash when using a very long file name](https://github.com/vim/vim/commit/2c66669c33e7aab70db089c393bedc0f21ca33c7)
- [7.3.649 : with 'clipboard' set to "unnamed" small deletes go to "1](https://github.com/vim/vim/commit/7c82130a76898999bcebf28c71f0897644822fab)
- [7.3.650 : Completing ":help \\{-" gives error and messes up command line](https://github.com/vim/vim/commit/b4a80cdd913605eefb26d6f091a386a4af0ac048)
- [7.3.651 : Completing ":help \\{-" gives error message](https://github.com/vim/vim/commit/7c0a86b1e871a9546387252b5a33594fc17487aa)
- [7.3.652 : workaround for Python crash isn't perfect](https://github.com/vim/vim/commit/ac0ddc15e928bb2c9baf256e0cd073770512092f)
- [7.3.653 : lacking XPM support for MingW, 32 and 64 objects mixed up](https://github.com/vim/vim/commit/e2db4361d2236b1bf56fb436d274c8b9b361196c)
- [7.3.654 : creating Vim dictionary from Python an empty key might be used](https://github.com/vim/vim/commit/231e1a17232a6d396266194efad53259e7d05ff1)
- [7.3.655 : 64 bit MingW xpm .a file is missing](https://github.com/vim/vim/commit/e48a2dd3af920d712305e4225454cb2940fb3312)
- [7.3.656 : internal error in :pyeval](https://github.com/vim/vim/commit/77fceb890851e8c97c9d4bbe521dff165455d371)
- [7.3.657 : Python bindings silently truncate string values containing NUL](https://github.com/vim/vim/commit/cdab9051368a6437b7e40df673594147425df751)
- [7.3.658 : NUL bytes truncate strings when converted from Python](https://github.com/vim/vim/commit/afa6b9af865c18dd865ccf7b804a4645570027bd)
- [7.3.659 : recent Python changes are not tested](https://github.com/vim/vim/commit/c11073c9aa2b668a9c1faa8a2419d7984d8367ca)
- [7.3.660 : ":help !" jumps to help for ":!"](https://github.com/vim/vim/commit/fca93c093e0ec86331171ca4e169f73aa81ebfef)
- [7.3.661 : (after 7.3.652) SEGV in Python code](https://github.com/vim/vim/commit/e8cdcef87575cacbc9dd977de076af76e4405410)
- [7.3.662 : can't build Ruby interface with Ruby 1.9.3](https://github.com/vim/vim/commit/2623b4f412a77f6426ab997246d755fdcbddea7b)
- [7.3.663 : end of color scheme name not clear in E185](https://github.com/vim/vim/commit/be1e9e9fc1f31e3d0b82eb3febd51923bda2a1e4)
- [7.3.664 : buffer overflow in unescaping multi-byte characters](https://github.com/vim/vim/commit/4fabd7dd4a2f77bc741eda58126d529f26f16a8d)
- [7.3.665 : MSVC 11 is not supported](https://github.com/vim/vim/commit/2b017fae8ffbf77fa59a8e9d49d9479ac5009a25)
- [7.3.666 : Win32.mak can't be found](https://github.com/vim/vim/commit/6f586de755cf0b814d0d8723fada8dd0bb58b94d)
- [7.3.667 : unused variables in Perl interface](https://github.com/vim/vim/commit/f5fe79a5f38e0fef6077aa858b60989e94076910)
- [7.3.668 : loading Perl dynamically still uses static library](https://github.com/vim/vim/commit/01c10524d25e36e72c7e45048a80cdb410165560)
- [7.3.669 : when building with Cygwin loading Python dynamically fails](https://github.com/vim/vim/commit/f94a13c286ffcf1ba2ce7bf10d686c757c836fe0)
- [7.3.670 : tiny memory leak when throwing exception in Python](https://github.com/vim/vim/commit/696c2116dc5e5403644d7f7dcaea49b83cf5a8f5)
- [7.3.671 : more Python code can be shared between Python 2 and 3](https://github.com/vim/vim/commit/770456589e5e53507fcd32013a190b520f4fc688)
- [7.3.672 : not possible to lock/unlock lists in Python interface](https://github.com/vim/vim/commit/66b7985ec9e4924abc370636077ed33114d0327d)

## 新着スクリプト

- [tinykeymap : Tiny, temporary keymaps](http://www.vim.org/scripts/script.php?script_id=4199)
- [transliterate : Text transliteration](http://www.vim.org/scripts/script.php?script_id=4200)
- [tabpagecd : Use one tab page per project](http://www.vim.org/scripts/script.php?script_id=4201)
- [altr : Switch to the missing file without interaction](http://www.vim.org/scripts/script.php?script_id=4202)
- [tropikos : Tropikos color scheme](http://www.vim.org/scripts/script.php?script_id=4203)
- [My Code Style Toggle : help you toggle the style of the code blocks easily !](http://www.vim.org/scripts/script.php?script_id=4204)
- [cppSyntaxCheck : A syntax check script for cpp](http://www.vim.org/scripts/script.php?script_id=4205)
- [vim-java-sql : SQL syntax-highlighting within .java files.](http://www.vim.org/scripts/script.php?script_id=4206)
- [vim-clojure-sql : SQL syntax-highlighting within Clojure files](http://www.vim.org/scripts/script.php?script_id=4207)
- [nant 0.91 errorformat : nant 0.91 errorformat](http://www.vim.org/scripts/script.php?script_id=4208)
- [Vimarok : Plugin for control amarok via dbus-python](http://www.vim.org/scripts/script.php?script_id=4209)
- [easybuffer.vim : vim plugin to quickly switch between buffers](http://www.vim.org/scripts/script.php?script_id=4210)
- [PinyinSearch : Allow you to search Chinese by the first letter of Chinese pinyin](http://www.vim.org/scripts/script.php?script_id=4211)
- [RelOps : Relative line numbers for pending operations](http://www.vim.org/scripts/script.php?script_id=4212)
- [jinja : jinja syntax and indent plugin](http://www.vim.org/scripts/script.php?script_id=4213)
- [gvimt : Open new files in gvim in new tabs, splits or vertical splits](http://www.vim.org/scripts/script.php?script_id=4214)
- [R-MacVim : R plugin for MacVim (R Vim Mac, R MacVim)](http://www.vim.org/scripts/script.php?script_id=4215)
- [fwk\_ctags :   Manage your ctags tags configurations simple and efficent](http://www.vim.org/scripts/script.php?script_id=4216)
- [layoutManager : Create layouts from tabs, templates and windows and control them.](http://www.vim.org/scripts/script.php?script_id=4217)
- [vim2grace : Grace/Common Music integration](http://www.vim.org/scripts/script.php?script_id=4218)
- [mkvimball.vim : Create a vimball from a list of input files](http://www.vim.org/scripts/script.php?script_id=4219)
- [cospe.vim : Script to send your buffer or selection to http://cos.pe and returns a url](http://www.vim.org/scripts/script.php?script_id=4220)
- [localbundle.vim :   faster loading and working speed with vundle](http://www.vim.org/scripts/script.php?script_id=4221)
- [conque-repl.vim : Use vim as a command editor for REPL programs](http://www.vim.org/scripts/script.php?script_id=4222)
- [votl : A Vim Outliner](http://www.vim.org/scripts/script.php?script_id=4223)
- [System :  A way to use ":call system('cmd')" via ":! cmd" (note the space)](http://www.vim.org/scripts/script.php?script_id=4224)
- [Taboo : Customize and rename tabs](http://www.vim.org/scripts/script.php?script_id=4237)
- [RepeatableYank : Repeatable appending yank to a named register.](http://www.vim.org/scripts/script.php?script_id=4238)
- [Freya : A dark color scheme with warm colors.](http://www.vim.org/scripts/script.php?script_id=4239)
- [colorschemer : Vim colorschemer](http://www.vim.org/scripts/script.php?script_id=4240)
- [betterSearch.vim : Provide better search functionality by capturing all matched line](http://www.vim.org/scripts/script.php?script_id=4241)
- [SameFiletypeComplete : Insert mode completion from buffers with the same filetype.](http://www.vim.org/scripts/script.php?script_id=4242)
- [IndentTab : Use tabs for indent at the beginning, spaces for alignment in the rest of a line](http://www.vim.org/scripts/script.php?script_id=4243)
- [yacom : easily toggle comments in line(s)](http://www.vim.org/scripts/script.php?script_id=4244)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (3168)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (2750)
3. [snipMate : TextMate-style snippets for Vim](http://www.vim.org/scripts/script.php?script_id=2540) (1471)
4. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (1316)
5. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (1267)
6. [bufexplorer.zip : Buffer Explorer / Browser](http://www.vim.org/scripts/script.php?script_id=42) (1251)
7. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (1172)
8. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (1139)
9. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (1126)
10. [pathogen.vim : Easy manipulation of 'runtimepath', 'path', 'tags', etc](http://www.vim.org/scripts/script.php?script_id=2332) (1058)

## vim-jp issues

Open : 130 (+1) | Closed : 131 (+15)

- [Issue #246 : iconv.dllがノートンにウィルスと誤認&削除された](https://github.com/vim-jp/issues/issues/246)
- [Issue #247 : ヘルプ補完でE64が発生する](https://github.com/vim-jp/issues/issues/247)
- [Issue #248 : JIS第3水準(JIS X 0213)の一部の漢字を、iconv()でISO-2022-JP-3に変換しようとするとkaoriya版gvimが落ちる](https://github.com/vim-jp/issues/issues/248)
- [Issue #249 : cfileしたときにfencsが効かない](https://github.com/vim-jp/issues/issues/249)
- [Issue #250 : CP932とMS932、どっちにしよう。](https://github.com/vim-jp/issues/issues/250)
- [Issue #251 : i(nore)?mapが発動するとIMEが勝手にOFFになってしまう](https://github.com/vim-jp/issues/issues/251)
- [Issue #252 : Cygwinビルドで +perl/dyn, +python/dyn が動かない](https://github.com/vim-jp/issues/issues/252)
- [Issue #253 : vimLoadLibの修正について検討する](https://github.com/vim-jp/issues/issues/253)
- [Issue #254 : Windows で ActivePerl 5.16 を使うと落ちる](https://github.com/vim-jp/issues/issues/254)
- [Issue #255 : NetrwSettingsを実行するとエラーが出る。](https://github.com/vim-jp/issues/issues/255)
- [Issue #256 : NetrwSettingsを実行するとエラーが出る。](https://github.com/vim-jp/issues/issues/256)
- [Issue #257 : Windowsでenc=utf-8の時、ファイル名に日本語を含むファイルを編集すると、スワップファイルが削除されない](https://github.com/vim-jp/issues/issues/257)
- [Issue #258 : :cwindow で quickfix ウィンドウが開かない](https://github.com/vim-jp/issues/issues/258)
- [Issue #259 :  IMEに設定するフォント](https://github.com/vim-jp/issues/issues/259)
- [Issue #260 : sxq のデフォルト値が変わって動かないスクリプトがある](https://github.com/vim-jp/issues/issues/260)
- [Issue #261 : windowやtabpageを閉じた／開いたことを捕捉するautocmdがほしい](https://github.com/vim-jp/issues/issues/261)

