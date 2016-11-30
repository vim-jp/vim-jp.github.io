---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2016 年 11 月号

---

## 話題

*   [VimConf 2016](http://vimconf.vim-jp.org/2016/) が開催されました
*   2016 年のアドベントカレンダー
    *   [Vim](http://qiita.com/advent-calendar/2016/vim)
    *   [Vim (その2)](http://qiita.com/advent-calendar/2016/vim2)
    *   [Vim 8.0](http://qiita.com/advent-calendar/2016/vim8)
*   [Meguro.vim #1](https://megurovim.connpass.com/event/46044/)

    Vimmer もくもく交流会です。 12/18(日) 目黒

## 今月の新機能

*   8.0.0056: `'filetype'` と `'syntax'` と `'keymap'` で正常値チェックをおこなうようになりました。

    [(脆弱性 CVE-2016-1248 の修正)](https://web.nvd.nist.gov/view/vuln/detail?vulnId=CVE-2016-1248)

*   8.0.0062: digraph に HORIZONTAL ELLIPSIS (` … `) が追加されました。
*   8.0.0077: Travis CI で GUI の test がおこなわれるようになりました。
*   8.0.0093: MSVC 10 のビルドがマルチプロセスに対応しました。
*   8.0.0096: 起動オプションに `--ttyfail` が追加されました。また、 `has("ttyin")` と `has("ttyout")` が追加されました。
*   8.0.0105: `ch_canread()` が追加されました。

## リリース情報

- [8.0.0056 : when setting 'filetype' there is no check for a valid name](https://github.com/vim/vim/commit/d0b5138ba4bccff8a744c99836041ef6322ed39a)
- [8.0.0057 : (after 8.0.0056) tests fail without the 'keymap' features](https://github.com/vim/vim/commit/9376f5f482a4d579436bf364778c2d8ab8e2f22d)
- [8.0.0058 : positioning of the popup menu is not good](https://github.com/vim/vim/commit/91e44a3305ef6bf2d43496c351dcff0a45c6bfb8)
- [8.0.0059 : Vim does not build on VMS systems](https://github.com/vim/vim/commit/de5e2c219b99895445fb75ae3541ee69282a5846)
- [8.0.0060 : when 'keywordprg' is an Ex command it is shell-escaped](https://github.com/vim/vim/commit/426f3754223c8ff8a1bc51d6ba1eba11e8982ebc)
- [8.0.0061 : (after 8.0.0058) compiler warning for unused variable](https://github.com/vim/vim/commit/aab3383e70456f054fe9d0963fe3eb45994aa5e7)
- [8.0.0062 : no digraph for HORIZONTAL ELLIPSIS](https://github.com/vim/vim/commit/81615517249bb78cba9c37c9834b787c1b265521)
- [8.0.0063 : compiler warning for comparing with unsigned](https://github.com/vim/vim/commit/55952d4dd490bb2f63bda5d7f6d8fb69f58c333c)
- [8.0.0064 : (after 8.0.0060) normal test fails on MS-Windows](https://github.com/vim/vim/commit/eb828d01d9c91791fa1fe217ba651cdc25746d1b)
- [8.0.0065 : (after 8.0.0056) compiler warning for unused function](https://github.com/vim/vim/commit/40d3f137e751c0e9f5e7e6a587b93a52dd833d04)
- [8.0.0066 : calling an operator function without 'linebreak' set](https://github.com/vim/vim/commit/4a08b0dc4dd70334056fc1bf069b5e938f2ed7d5)
- [8.0.0067 : VMS has a problem with infinity](https://github.com/vim/vim/commit/98500fdc6119eb5f02d7a52ab6ffcac3085181be)
- [8.0.0068 : checking did&#x5f;throw after executing autocommands is wrong](https://github.com/vim/vim/commit/21662be2211675824df1771c7f169948ede40c41)
- [8.0.0069 : compiler warning for self-comparison](https://github.com/vim/vim/commit/a1f4cb93ba50ea9e40cd4b1f5592b8a6d1398660)
- [8.0.0070 : tests referred in Makefile that no longer exist](https://github.com/vim/vim/commit/3f9ebf32a392a9cae1c3e4b6bf8cecad60e2a22a)
- [8.0.0071 : exit value from a shell command is wrong](https://github.com/vim/vim/commit/c4d4ac22f78fb1394c79eccc8a1e6812c0c7d8a7)
- [8.0.0072 : MS-Windows: Crash with long font name](https://github.com/vim/vim/commit/38bc49563782ee1cb91660e58acf1afe1a31020a)
- [8.0.0073 : (after 8.0.0069) more comparisons between firstwin & lastwin](https://github.com/vim/vim/commit/459ca563128f2edb7e3bb190090bbb755a56dd55)
- [8.0.0074 : cannot make Vim fail on an internal error](https://github.com/vim/vim/commit/95f096030ed1a8afea028f2ea295d6f6a70f466f)
- [8.0.0075 : using number for exception type lacks type checking](https://github.com/vim/vim/commit/8a5883b7488e492419dde7e1637cc72f2d566ba4)
- [8.0.0076 : channel log has double parens ()()](https://github.com/vim/vim/commit/79cbdcb1eeead00bad6fc64ab4a53e5860d5a663)
- [8.0.0077 : GUI is not tested on Travis CI](https://github.com/vim/vim/commit/e3505dfc858260a21f2bbc6056f5769dea2f273d)
- [8.0.0078 : accessing freed memory in quickfix](https://github.com/vim/vim/commit/63bed3d319b5d90765dbdae93a3579b6322d79fb)
- [8.0.0079 : accessing freed memory in quickfix](https://github.com/vim/vim/commit/2b946c9f9b0e0fd805fb8f3e4c16e0a68ae13129)
- [8.0.0080 : Travis CI fails for OS X](https://github.com/vim/vim/commit/b6be3ea45b3bb260d475ba63620252ce2963c560)
- [8.0.0081 : inconsistent function names](https://github.com/vim/vim/commit/d4db7719bdfbc54df396eac08d8cbb2389feacf4)
- [8.0.0082 : extension for configure should be ".ac"](https://github.com/vim/vim/commit/3f7d0907269558cb3ea184a3083640f9e20bb21e)
- [8.0.0083 : using freed memory with win&#x5f;getid()](https://github.com/vim/vim/commit/8e639052638a9bb8c7dd6e3e10776b1218cec1a3)
- [8.0.0084 : using freed memory when adding to a quickfix list](https://github.com/vim/vim/commit/7618e00d3b8bfe064cfc524640d754607361f9df)
- [8.0.0085 : using freed memory with recursive function call](https://github.com/vim/vim/commit/8a01f969c198eeb655ad2f96f2796a6f6f4a1924)
- [8.0.0086 : cannot add a comment after ":hide"](https://github.com/vim/vim/commit/2256c9947164229c0960803e2a2992b793c23298)
- [8.0.0087 : channel callback may get job info after it was deleted](https://github.com/vim/vim/commit/7df915d113ac1981792c50e8b000c9f5f784b78b)
- [8.0.0088 : when a test fails in Setup/Teardown problem is not reported](https://github.com/vim/vim/commit/cc28e2d05d05552d8b72a520be8a193f3d9822d4)
- [8.0.0089 : various problems with GTK 3.22.2](https://github.com/vim/vim/commit/a859f04b4db651860c07db3587f29906517c552b)
- [8.0.0090 : cursor moved after last character when using 'breakindent'](https://github.com/vim/vim/commit/6c896867c4f5d759616028ef7cbfce2a9ed32600)
- [8.0.0091 : test&#x5f;help&#x5f;complete sometimes fails in MS-Windows console](https://github.com/vim/vim/commit/9f0e423c2818c0cacd0810f9c3c67cbb6b80963d)
- [8.0.0092 : C indent does not support nested namespaces that C++ 17 has](https://github.com/vim/vim/commit/ca8b8d6956dd881de6446fc32c38e817a364a6cc)
- [8.0.0093 : not using multiprocess build feature](https://github.com/vim/vim/commit/dda39aeafc94484e7d209d7bdfd2fc403b7383f5)
- [8.0.0094 : if vimrun.exe not found error message is not properly encoded](https://github.com/vim/vim/commit/63e4344edc0cc1b4ed68a9d9c787265799602670)
- [8.0.0095 : problems with GTK 3.22.2 fixed in 3.22.4](https://github.com/vim/vim/commit/182707ac10d77359bf7a87c6b23ce4025d5b0ad4)
- [8.0.0096 : when the input or output is not a tty Vim appears to hang](https://github.com/vim/vim/commit/2cab0e191055a8145ccd46cd52869fbb9798b971)
- [8.0.0097 : if channel callback takes a lot of time Vim is unresponsive](https://github.com/vim/vim/commit/833eb1d752426689051bf2001083359899536939)
- [8.0.0098 : (after 8.0.0097) can't build on MS-Windows](https://github.com/vim/vim/commit/58c358753ef47e217e18c0bf8a2fca9d1ea23a8e)
- [8.0.0099 : popup menu appears above the cursor too often](https://github.com/vim/vim/commit/73095288da839f7c738a49baa109773e76106806)
- [8.0.0100 : file name options may contain non-filename characters](https://github.com/vim/vim/commit/319afe3804741db5a6c188bd69535fa7ed044c62)
- [8.0.0101 : some options are not strictly checked](https://github.com/vim/vim/commit/031cb743ae154cfb727a9b7787bdcb61202ff1c8)
- [8.0.0102 : (after 8.0.0101) cannot set 'dictionary' to a path](https://github.com/vim/vim/commit/7554da4033498c4da0af3cde542c3e87e9097b73)
- [8.0.0103 : may not process channel readahead](https://github.com/vim/vim/commit/8a8199e4a1814b10630a770165502abb1121cd1b)
- [8.0.0104 : value of 'thesaurus' option not checked properly](https://github.com/vim/vim/commit/f422bcc7f9615fe91fa69b059cfe4785093d3d4a)
- [8.0.0105 : can't tell difference between reading empty line or nothing](https://github.com/vim/vim/commit/4b785f69c0616dba5d3f38e8ce4b5398cec89407)
- [8.0.0106 : (after 8.0.0100) cannot use a semicolon in 'backupext'](https://github.com/vim/vim/commit/0945eaface83e78138fbd40f95cc590bab0e8c86)

## 新着スクリプト

- [vim-codequery : An asynchronous code searcher](http://www.vim.org/scripts/script.php?script_id=5477)
- [mythos : A dark, low contrast, Cthulhu-esque colour scheme](http://www.vim.org/scripts/script.php?script_id=5478)
- [vim-makejob : Minimal, asynchronous make and lmake for Vim 8.0](http://www.vim.org/scripts/script.php?script_id=5479)
- [StyleGuide : Code with google-syleguide. Readability, quality and unity well be better](http://www.vim.org/scripts/script.php?script_id=5480)
- [gnugo.vim : Play a game of Go in your text editor, using GnuGo](http://www.vim.org/scripts/script.php?script_id=5481)
- [vim-matchit : Extended matching with "%"](http://www.vim.org/scripts/script.php?script_id=5482)
- [Replay : Convenient re-run of last recorded Vim macro](http://www.vim.org/scripts/script.php?script_id=5483)
- [syn-preproc : C preprocessor syntax on top of other syntaxes.](http://www.vim.org/scripts/script.php?script_id=5484)
- [vim-shaderlab : Unity ShaderLab](http://www.vim.org/scripts/script.php?script_id=5485)
- [vim-yesod : Syntax and utility for the Haskell Yesod web framework](http://www.vim.org/scripts/script.php?script_id=5486)
- [ViewMaps : Lightweight plugin for viewing mappings from vimrc](http://www.vim.org/scripts/script.php?script_id=5487)
- [flatwhite : A Vim port of biletskyys flatwhite-syntax for Atom](http://www.vim.org/scripts/script.php?script_id=5488)
- [vim-guesslang : guess the spell-check language of the current buffer](http://www.vim.org/scripts/script.php?script_id=5489)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (1417)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1239)
3. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (584)
4. [c.vim : C/C++ IDE -- Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (505)
5. [wombat256.vim : Wombat for 256 color xterms](http://www.vim.org/scripts/script.php?script_id=2465) (430)
6. [nginx.vim : initial version](http://www.vim.org/scripts/script.php?script_id=1886) (429)
7. [winmanager : A windows style IDE for Vim 6.0](http://www.vim.org/scripts/script.php?script_id=95) (425)
8. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (419)
9. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (399)
10. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (393)

## vim-jp issues

Open : 249 (+12) | Closed : 742 (+5)

- [Issue #975 : デュアルモニタで「変更を "無題" に保存しますか？」ダイアログが常にメインモニタ上に現れてしまう](https://github.com/vim-jp/issues/issues/975)
- [Issue #976 : \[count\]tabnextの挙動が\[count\]tabcloseや\[count\]tabonlyなどのcountの解釈と異なる](https://github.com/vim-jp/issues/issues/976)
- [Issue #977 : マルチ選択領域機能の考察](https://github.com/vim-jp/issues/issues/977)
- [Issue #978 : EPIPE error](https://github.com/vim-jp/issues/issues/978)
- [Issue #979 : ユーザー定義のモーションでマクロを中断させる方法](https://github.com/vim-jp/issues/issues/979)
- [Issue #980 : キーマクロの実行中に意図しないキーマップの展開が起こる](https://github.com/vim-jp/issues/issues/980)
- [Issue #981 : :help :&& の項目が存在しない](https://github.com/vim-jp/issues/issues/981)
- [Issue #982 : PYTHONHOMEをif&#x5f;pythonとif&#x5f;python3に同時に設定できない](https://github.com/vim-jp/issues/issues/982)
- [Issue #983 : 乱数が欲しい](https://github.com/vim-jp/issues/issues/983)
- [Issue #984 : 補完候補が 1 つの際の noselect,noinsert の挙動](https://github.com/vim-jp/issues/issues/984)
- [Issue #985 : channel の IPv6 対応](https://github.com/vim-jp/issues/issues/985)
- [Issue #986 : テスト実行後に test29.in の改行コードが書き換わる](https://github.com/vim-jp/issues/issues/986)
- [Issue #987 : textobj と syntax の連携](https://github.com/vim-jp/issues/issues/987)
- [Issue #988 : v:register が更新されない場合がある](https://github.com/vim-jp/issues/issues/988)
- [Issue #989 : :filter! の help がない](https://github.com/vim-jp/issues/issues/989)
- [Issue #990 : 異なる引数に対して同じウィンドウ番号が返されるケース](https://github.com/vim-jp/issues/issues/990)
- [Issue #991 : :filter の {pat} が / で囲まれていないと {command} の補完が効かない](https://github.com/vim-jp/issues/issues/991)

