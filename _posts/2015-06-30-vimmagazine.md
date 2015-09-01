---
layout: post
category: vimmagazine
title: Vim Magazine 6 月号
---

## 今月の新機能

- 7.4.739: 文字列の文字指定で `\U` が 8 桁の数字を取るようになりました。
- 7.4.742: 'switchbuf' オプションに "vsplit" が追加されました。
- 7.4.754: ビジュアルモードで `CTRL-A`, `CTRL-X`, `g CTRL-A`, `g CTRL-X` が追加されました。連番の生成、更新に便利です。
- 7.4.755: strchars() に skipcc 引数が追加されました。基底文字と結合文字をまとめて一つとしてカウントできます。
- 7.4.757: `'t_RB'` オプションが追加されました。`'background'` の自動設定ができます。

## リリース情報

- [7.4.730 : when setting the crypt key text in swap file may be corrupted](https://github.com/vim/vim/commit/bc56336bb4501884257352446abb60713cef6452)
- [7.4.731 : the tab menu shows "Close tab" even when it doesn't work](https://github.com/vim/vim/commit/7098ee5c37e35494ed36f3ad2d1c305341a41a45)
- [7.4.732 : the cursor line is not always updated for the "O" command](https://github.com/vim/vim/commit/d0d0fe09cf95bf88f643ec07d07939cbd41f8e63)
- [7.4.733 : test&#x5f;listchars breaks on MS-Windows](https://github.com/vim/vim/commit/d68f2219b57acb86ddedebdcc1476fee15c9c0c7)
- [7.4.734 : ml&#x5f;get error when using "p" in Visual selection in last line](https://github.com/vim/vim/commit/d009e8682686a56f7565e6e093a42cd0596e121f)
- [7.4.735 : wrong argument for sizeof()](https://github.com/vim/vim/commit/e2e4b98622f7c44910dc9660035048a4177d3777)
- [7.4.736 : invalid memory access](https://github.com/vim/vim/commit/b878bbbca67c384101c2ebb822e390f0feda9489)
- [7.4.737 : on MS-Windows vimgrep over arglist doesn't work (Issue 361)](https://github.com/vim/vim/commit/6e8d3b0d41bcef7ebe0ba2e887eb25781e1437fe)
- [7.4.738 : (after 7.4.732) can't compile without the syntax HL feature](https://github.com/vim/vim/commit/d710e0d2478e7f813562bfb084406d524bfecd19)
- [7.4.739 : in a string "\\U" only takes 4 digits, should be eight](https://github.com/vim/vim/commit/acc39888cddbc8082fb3b89c1a46646bdca55e42)
- [7.4.740 : ":1quit" works like ":.quit"](https://github.com/vim/vim/commit/c7a0d32c834ff321491d8fd7bc68bf6cabd11e44)
- [7.4.741 : when using += with ":set" a trailing comma is not recognized](https://github.com/vim/vim/commit/a7b7b1cef9b61700ba5076c1d889dce791392388)
- [7.4.742 : no vertical split when loading buffer for a quickfix command](https://github.com/vim/vim/commit/a594d77ffcccf2ac0e4079c41342ca55d4c9bb08)
- [7.4.743 : "p" in Visual mode causes an unexpected line split](https://github.com/vim/vim/commit/c004bc2726eafc7a56d1d9f8398a65a0a7dc8d6c)
- [7.4.744 : no tests for Ruby and Perl](https://github.com/vim/vim/commit/cf0213ef7e88d38c34ef71459d413cf5813e2a95)
- [7.4.745 : entries returned by getmatches() dont work with setmatches()](https://github.com/vim/vim/commit/0fce4257727f9d75e488963b73e407d31dd46546)
- [7.4.746 : ":\[count\]tag" is not always working](https://github.com/vim/vim/commit/01cf376da1726862afc8fa1d84cf5a773909fd0d)
- [7.4.747 : ":cnext" may jump to wrong column when 'virtualedit' is "all"](https://github.com/vim/vim/commit/b8c890035efd694daab5cdd71b5265c52c23fa81)
- [7.4.748 : (after 7.4.745) buffer overflow](https://github.com/vim/vim/commit/6a7e2a668b492b5b574e489790e349a9058e2a48)
- [7.4.749 : (after 7.4.741) two consecutive commas are OK for some options](https://github.com/vim/vim/commit/0e7c4b99c608d22fd1b8526e700420b248319b1f)
- [7.4.750 : cannot build with clang 3.5 on Cygwin with perl enabled](https://github.com/vim/vim/commit/280a868a246eda3904eb925588f111b69e304878)
- [7.4.751 : it is not obvious how to enable the address sanitizer](https://github.com/vim/vim/commit/70b5a5f5b9beb2c5d7ccb3518c474c432f31ad91)
- [7.4.752 : Unicode 8.0 not supported](https://github.com/vim/vim/commit/66312acf145f9d09fbc0dfa5caeafb715d5fd9ed)
- [7.4.753 : appending in Visual mode with 'linebreak' set is wrong](https://github.com/vim/vim/commit/74db34cc9162445e3d500ab2d61c5c19ce9af0ca)
- [7.4.754 : using CTRL-A in Visual mode does not work well](https://github.com/vim/vim/commit/3a304b23823b089e499063e8211c5695d049f3ba)
- [7.4.755 : it is not easy to count the number of characters](https://github.com/vim/vim/commit/641e48c2248ccb3c25a5cdaa3709f16152d8c77d)
- [7.4.756 : can't use strawberry Perl 5.22 x64 on MS-Windows](https://github.com/vim/vim/commit/367fbf17dad7bfb65cef8a4a18147989516218f2)
- [7.4.757 : cannot detect the background color of a terminal](https://github.com/vim/vim/commit/b5c3265521749fda81ae4b052de35a0d1209cf50)
- [7.4.758 : 'conceallevel' in command-line window hides first character](https://github.com/vim/vim/commit/fa67fbe6b84133207271e4ff582c3d589172efd9)
- [7.4.759 : building with Lua 5.3 doesn't work, symbols have changed](https://github.com/vim/vim/commit/9514b1f9a12f6d032f7d749dfb0e947332d1c5e1)
- [7.4.760 : spelling mistakes are not displayed after ":syn spell"](https://github.com/vim/vim/commit/5081d202475328a920c0bdcab990e8da84128c13)
- [7.4.761 : (after 7.4.757) incomplete request-background implementation](https://github.com/vim/vim/commit/e5401220b0def9ac380c7910e18f2e042ed5ddc5)
- [7.4.762 : (after 7.4.757) comment for may&#x5f;req&#x5f;bg&#x5f;color() is wrong](https://github.com/vim/vim/commit/4921c24b3183d6000fdc66684f0752dcf86f48a9)
- [7.4.763 : (after 7.4.759) building with Lua 5.1 doesn't work](https://github.com/vim/vim/commit/1f860d8590460b6f13b39bebf191b81572fa0fdf)
- [7.4.764 : (after 7.4.754) test&#x5f;increment fails on MS-Windows](https://github.com/vim/vim/commit/6a57cce04801d3e8fe30d0dcd31f9325b38de2fe)

## 新着スクリプト

- [ctrlsf.vim : an ack/ag powered global code search and view tool](http://www.vim.org/scripts/script.php?script_id=5189)
- [vim-tag-comment : property comments HTML/XML tags](http://www.vim.org/scripts/script.php?script_id=5190)
- [scilab.vim : Scilab syntax highlighting (Windows)](http://www.vim.org/scripts/script.php?script_id=5191)
- [cpp&#x5f;cppcheck.vim : Run Cppcheck on the current window.](http://www.vim.org/scripts/script.php?script_id=5192)
- [zftool.vim : PHP plugin for running ZFTool commands](http://www.vim.org/scripts/script.php?script_id=5193)
- [cpp&#x5f;doxygen : Add Doxygen templates and fold Doxygen comments.](http://www.vim.org/scripts/script.php?script_id=5194)
- [Indent Detector : detect indent automatically](http://www.vim.org/scripts/script.php?script_id=5195)
- [Dokumentary : Improve what K does.](http://www.vim.org/scripts/script.php?script_id=5197)
- [vim-phpfmt : Automatically formats PHP code (needs PHP 5.6 installed)](http://www.vim.org/scripts/script.php?script_id=5198)
- [flagship.vim : Configurable and extensible tab line and status line](http://www.vim.org/scripts/script.php?script_id=5199)
- [vim-js-arrow-function : changes a JavaScript anonymous function to an arrow function expression (ES2015)](http://www.vim.org/scripts/script.php?script_id=5200)
- [scilab.vim : Scilab syntax highlighting (Linux)](http://www.vim.org/scripts/script.php?script_id=5201)
- [supersearch : source view plugin](http://www.vim.org/scripts/script.php?script_id=5202)
- [asynccompl : A simple asynchronous complete plugin](http://www.vim.org/scripts/script.php?script_id=5203)
- [mirror.vim : Efficient way to edit remote files on multiple environments](http://www.vim.org/scripts/script.php?script_id=5204)
- [ShaderHighLight : 常用的shader文件语法高亮](http://www.vim.org/scripts/script.php?script_id=5205)
- [auto&#x5f;autoread : Automatically read files when they've changed. Does what 'autoread' promises.](http://www.vim.org/scripts/script.php?script_id=5206)
- [undofile&#x5f;warn.vim : Warn when using the undofile.](http://www.vim.org/scripts/script.php?script_id=5207)

## 月間ダウンロードランキング

1. [AutomaticLaTeXPlugin : Background compilation, completion, bib serch, toc and other nice features.](http://www.vim.org/scripts/script.php?script_id=2945) (1987)
2. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (1722)
3. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1479)
4. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (741)
5. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (629)
6. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (600)
7. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (591)
8. [winmanager : A windows style IDE for Vim 6.0](http://www.vim.org/scripts/script.php?script_id=95) (577)
9. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (538)
10. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (507)

## vim-jp issues

Open : 230 (-4) | Closed : 514 (+12)

- [Issue #737 : Windows で Ex モードで起動すると画面がクリアされる](https://github.com/vim-jp/issues/issues/737)
- [Issue #738 : コマンドラインに対して ':wincmd &#x5f;' が効いてしまう](https://github.com/vim-jp/issues/issues/738)
- [Issue #739 : Linux上で入力時にしかCursorIMのハイライトに変わらない](https://github.com/vim-jp/issues/issues/739)
- [Issue #740 : 07 インクリメント 010 ？](https://github.com/vim-jp/issues/issues/740)
- [Issue #741 : プロポーショナルフォントを使いたい](https://github.com/vim-jp/issues/issues/741)
- [Issue #742 : strchars()に合成文字スキップフラグを追加する](https://github.com/vim-jp/issues/issues/742)
- [Issue #743 : test&#x5f;increment FAILED](https://github.com/vim-jp/issues/issues/743)
- [Issue #744 : dos-standard-mappings の記述不足](https://github.com/vim-jp/issues/issues/744)

