---
layout: post
category: vimmagazine
title: Vim Magazine 11 月号
---

## 話題

- [Vim Advent Calendar 2015](http://qiita.com/advent-calendar/2015/vim)

    伝説のVim Advent Calendarが、また今年もやってまいりました!

- [Vim script Advent Calendar 2015](http://qiita.com/advent-calendar/2015/vim-script)

    ランタイムのための Vim Advent Calendar とは違ってこっちは言語の方です。


## 今月の新機能

- 7.4.907 : 言語拡張の dll パスを指定するためのオプションが追加されました。
- 7.4.911 : `'t_Ce'` と `'t_Cs'` オプションが追加されました。
- 7.4.941 : `'tagcase'` オプションが追加されました。
- 7.4.944 : `assertEqual()`, `assertFalse()`, `assertTrue()` 関数が追加されました。

## リリース情報

- [7.4.903 : MS-Windows: expandinig wildcards may cause bad memory access](https://github.com/vim/vim/commit/7314efd87d8c4095229bdc2867a553c36c064918)
- [7.4.904 : Vim does not provide .desktop files](https://github.com/vim/vim/commit/6407b3e80d7d7f8f0797c13ae35cc06f96be46c9)
- [7.4.905 : Python interface misses some functions on vim.message](https://github.com/vim/vim/commit/d424747d5821c2873e24d25d3407d08b25ea3443)
- [7.4.906 : MS-Windows: viminfo file is (always) given hidden attribute](https://github.com/vim/vim/commit/8a52ba791893fd55c5bdf98825c5b3e8892eaa62)
- [7.4.907 : dynamically loaded libs can only be defined at compile time](https://github.com/vim/vim/commit/d94464ee294a351ce7b6ba18e8bd3f24f1bef920)
- [7.4.908 : (after 7.4.907) build error with MingW compiler](https://github.com/vim/vim/commit/a16f472edfa028e5574c7c145d02f3821edbc698)
- [7.4.909 : "make install" fails](https://github.com/vim/vim/commit/de59ba33aa3b94f2757dbf3451682d762c15ebcf)
- [7.4.910 : (after 7.4.905) compiler complains about type punned pointer](https://github.com/vim/vim/e7427f4b7e1af6a63600183be6b4c5724beb2f66/43345546ae63710441f066648b8485fb545b3801)
- [7.4.911 : t&#x5f;Ce and t&#x5f;Cs are documented but not supported](https://github.com/vim/vim/commit/450ca4335e467ac29c1560b7397225a974aee3bf)
- [7.4.912 : wrong indenting for C++ constructor](https://github.com/vim/vim/commit/e01f4f86cef7bed3cb99b26f9f57d86f6eb5fe1a)
- [7.4.913 : no utf-8 support for the hangul input feature](https://github.com/vim/vim/commit/72f4cc4a987d123c0ed909c85b9a05f65cef7202)
- [7.4.914 : new compiler warning: logical-not-parentheses](https://github.com/vim/vim/commit/98b30a473a58ae98c280e0383c8b1e08c0ebced5)
- [7.4.915 : comma goes missing when removing from 'path' and then adding](https://github.com/vim/vim/commit/174674743d9a2d7361c9cd89836f8dd8651edeeb)
- [7.4.916 : when running out of memory, memory may be freed twice](https://github.com/vim/vim/commit/8648357841065295e39831d2b559d87ca01a7a7c)
- [7.4.917 : compiler warning for comparing signed and unsigned](https://github.com/vim/vim/commit/1be2ed6c11671eabefa0fc8600fd2af6cd3963e8)
- [7.4.918 : (after 7.4.907) a digit in an option name has problems](https://github.com/vim/vim/commit/0796c0625fa4b9eb2f47fe8c976b78523924e1fb)
- [7.4.919 : the dll options are not in the options window](https://github.com/vim/vim/commit/d4ece23e2e602d820ab7367c383dc0d72dd87029)
- [7.4.920 : the rubydll option is not in the options window](https://github.com/vim/vim/commit/a93f975e8b39d7cfc8145dbe181cc4e5e4ec0bdf)
- [7.4.921 : (after 7.4.906) missing proto file update](https://github.com/vim/vim/commit/cd1c55f706c2f9b8310b8a9fc1f8226c7fd19372)
- [7.4.922 : leaking memory with ":helpt {dir-not-exists}"](https://github.com/vim/vim/commit/1c2836e268ce930bca9ea1287d0d83e92ce1b3ff)
- [7.4.923 : prototypes not always generated](https://github.com/vim/vim/commit/6a2697ffd7e894861853b351689b0ddec8901c96)
- [7.4.924 : DEVELOPER&#x5f;DIR gets reset by configure](https://github.com/vim/vim/commit/32d03b34ac8a34a962f57847fc431a2b4e14efea)
- [7.4.925 : user may yank or put using the register being recorded in](https://github.com/vim/vim/commit/a0ed84a26897c994512873a895b9fc54e90c6845)
- [7.4.926 : completing longest match doesn'thandle multi-byte chars](https://github.com/vim/vim/commit/4f8fa1633cdfbd09a41160c8480fe67c198067e9)
- [7.4.927 : Ruby crashes when there is a runtime error](https://github.com/vim/vim/commit/9b1067e038d371bd6c51e5da025383761f4921b4)
- [7.4.928 : a clientserver message interrupts handling keys of a mapping](https://github.com/vim/vim/commit/5d8afebb5bf7fb1e8ce06062451dc6a1f9a53ac0)
- [7.4.929 : "gv" after paste selects one char less if 'sel' is "exclusive"](https://github.com/vim/vim/commit/d29c6fea94947b3f4b54fbd5a6f832a7d744bf27)
- [7.4.930 : MS-Windows: Most users appear not to like the window border](https://github.com/vim/vim/commit/97b0b0ec764d3a247ef600d809b965d5ab37155d)
- [7.4.931 : (after 7.4.929) test 94 fails on some systems](https://github.com/vim/vim/commit/cfcd1ddd103129b309671cba5cff55e19a9908e4)
- [7.4.932 : (after 7.4.926) test&#x5f;utf8 has confusing dummy command](https://github.com/vim/vim/commit/8f08dab18df6dbf6c4b4973fd2d480e4bffb82d8)
- [7.4.933 : (after 7.4.926) crash using longest match in completion](https://github.com/vim/vim/commit/e4eda3bc7157932b0bf380fd3fdc1ba8f4438b60)
- [7.4.934 : Appveyor also builds on a tag push](https://github.com/vim/vim/commit/7487792ab14c1fb8dbdb37bdd74265d8b1d3be50)
- [7.4.935 : test&#x5f;utf8 fails on MS-Windows when executed with gvim](https://github.com/vim/vim/commit/6040256d8b0404564ac6f192296b12ea9d175e7d)
- [7.4.936 : crash when dragging with the mouse](https://github.com/vim/vim/commit/294a7e55b01149154807a23323038784549b8946)
- [7.4.937 : segfault reading unitialized memory](https://github.com/vim/vim/commit/5ad075c0735d3d8b97708d17c22de8facb15f997)
- [7.4.938 : X11 and GTK have moure mouse buttons than Vim supports](https://github.com/vim/vim/commit/88e484bf1b0afb5f2dec44f19335729578ace66a)
- [7.4.939 : memory leak when encountering a syntax error](https://github.com/vim/vim/commit/9a7d58e42ed54406437c2394e5a489ee6a9e4220)
- [7.4.940 : vt52 terminal codes are not correct](https://github.com/vim/vim/commit/2a1b474fd82aff922f18570593972b12feaa2073)
- [7.4.941 : there is no way to ignore case only for tag searches](https://github.com/vim/vim/commit/0f6562e9036f889185dff49a75c7fc5ffb28b307)
- [7.4.942 : (after 7.4.941) test&#x5f;tagcase breaks for small builds](https://github.com/vim/vim/commit/60422e68a3a555144f8c76c666f050e8d104c16b)
- [7.4.943 : test&#x5f;writefile tests are not run](https://github.com/vim/vim/commit/48a969b48898fb08dce636c6b918408c6fbd3ea0)
- [7.4.944 : Writing tests for Vim script is hard](https://github.com/vim/vim/commit/43345546ae63710441f066648b8485fb545b3801)

## 新着スクリプト

- [Never Optimal Wiki : a minimalistic pure vim personal wiki](http://www.vim.org/scripts/script.php?script_id=5263)
- [vim-abbrev-matcher : Abbreviation matcher plugin to be used with CtrlP and Unite](http://www.vim.org/scripts/script.php?script_id=5264)
- [wimproved.vim : Better fullscreen and more on Windows.](http://www.vim.org/scripts/script.php?script_id=5265)
- [vim-maildrop : Syntax highlighting for Courier maildrop filters](http://www.vim.org/scripts/script.php?script_id=5266)
- [simple&#x5f;playground.vim : make vim be language playgrounds](http://www.vim.org/scripts/script.php?script_id=5267)
- [fountainflow.vim : Convert Fountain screenplay files to HTML and LibreOffice](http://www.vim.org/scripts/script.php?script_id=5268)
- [Nette : Plugin for Nette Framework; syntax, indent, snipMate snippets, ...](http://www.vim.org/scripts/script.php?script_id=5269)
- [todo-vim : todo-vim is plugin for manage your todo notes](http://www.vim.org/scripts/script.php?script_id=5270)
- [EasyClipRing.vim : Autocomplete style vim pop-up menu for your registers.](http://www.vim.org/scripts/script.php?script_id=5271)
- [Buffer Hint : A handy buffer switcher](http://www.vim.org/scripts/script.php?script_id=5272)
- [Delphi : Executes your python code as you write it](http://www.vim.org/scripts/script.php?script_id=5275)
- [optcmd : plugin provides dialog for choicing command and process it by prefix option](http://www.vim.org/scripts/script.php?script_id=5276)

## 月間ダウンロードランキング

1. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1275)
2. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (1241)
3. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (656)
4. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (598)
5. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (517)
6. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (509)
7. [AutomaticLaTeXPlugin : Background compilation, completion, bib serch, toc and other nice features.](http://www.vim.org/scripts/script.php?script_id=2945) (401)
8. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (372)
9. [nginx.vim : initial version](http://www.vim.org/scripts/script.php?script_id=1886) (348)
10. [SuperTab : Do all your insert-mode completion with Tab.](http://www.vim.org/scripts/script.php?script_id=1643) (346)

## vim-jp issues

Open : 238 (-2) | Closed : 553 (+8)

- [Issue #786 : if&#x5f;pythonとif&#x5f;python3の同時使用](https://github.com/vim-jp/issues/issues/786)
- [Issue #787 : HomebrewでvimのHEADをインストールしようとすると失敗する](https://github.com/vim-jp/issues/issues/787)
- [Issue #788 : vim.exe gvim.exe起動するとノートンセキュリティに不正と判断されて削除される](https://github.com/vim-jp/issues/issues/788)
- [Issue #789 : Unix, Cygwin, Windows で fnamemodify の相対パス変換結果が違う](https://github.com/vim-jp/issues/issues/789)
- [Issue #790 : getcwd() でカレントウィンドウ以外のローカル cwd を取りたい](https://github.com/vim-jp/issues/issues/790)
- [Issue #791 : wildmode=longest 時にマルチバイト文字がバイト単位で補完される](https://github.com/vim-jp/issues/issues/791)

