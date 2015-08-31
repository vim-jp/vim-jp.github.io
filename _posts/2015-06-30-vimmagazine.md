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

- [7.4.730 : when setting the crypt key text in swap file may be corrupted](http://code.google.com/p/vim/source/detail?r=cd47def2214a1a68780729a0033c4cbde8529461)
- [7.4.731 : the tab menu shows "Close tab" even when it doesn't work](http://code.google.com/p/vim/source/detail?r=5f24d6d5133386cdc5e4dcf781e4e49d78d75483)
- [7.4.732 : the cursor line is not always updated for the "O" command](http://code.google.com/p/vim/source/detail?r=cd7a7975f92cc43405307baa7b980b249c2c0ff1)
- [7.4.733 : test&#x5f;listchars breaks on MS-Windows](http://code.google.com/p/vim/source/detail?r=095b538d58a673521eb4dbe0b819c63538ceef0c)
- [7.4.734 : ml&#x5f;get error when using "p" in Visual selection in last line](http://code.google.com/p/vim/source/detail?r=bf3e6012dfbd3b404ae1d0e41ff90a9628380062)
- [7.4.735 : wrong argument for sizeof()](http://code.google.com/p/vim/source/detail?r=734c89ddf7b37a2860b376d14404d7e1ca56a228)
- [7.4.736 : invalid memory access](http://code.google.com/p/vim/source/detail?r=70fdcdaf96a65db5fd2ef5d87ffa213f2397cb9b)
- [7.4.737 : on MS-Windows vimgrep over arglist doesn't work (Issue 361)](http://code.google.com/p/vim/source/detail?r=1ccf9612841097ca6310564e713779bb1d2d1e91)
- [7.4.738 : (after 7.4.732) can't compile without the syntax HL feature](http://code.google.com/p/vim/source/detail?r=8730db0c187bee6127daebab4467fafc65198f25)
- [7.4.739 : in a string "\\U" only takes 4 digits, should be eight](http://code.google.com/p/vim/source/detail?r=541abe8d5f557fbb7fcdb4844d20b1b0634ea72c)
- [7.4.740 : ":1quit" works like ":.quit"](http://code.google.com/p/vim/source/detail?r=0229cfffc5605b0cdbd87803aeb6d76053c14abb)
- [7.4.741 : when using += with ":set" a trailing comma is not recognized](http://code.google.com/p/vim/source/detail?r=0f1cc87630d8b91ea91825d53b64f48a759b58a1)
- [7.4.742 : no vertical split when loading buffer for a quickfix command](http://code.google.com/p/vim/source/detail?r=50cdebea094fced0f3d05f5fbf93f47f71d6a715)
- [7.4.743 : "p" in Visual mode causes an unexpected line split](http://code.google.com/p/vim/source/detail?r=ac13f3533571f689231f0d9d0b0672601eee9cad)
- [7.4.744 : no tests for Ruby and Perl](http://code.google.com/p/vim/source/detail?r=18ec2c5a5ce7af451d37617062309f2b3960f73c)
- [7.4.745 : entries returned by getmatches() dont work with setmatches()](http://code.google.com/p/vim/source/detail?r=f4e5704c5c43ca62a642c84bcec3f1f17a41fff6)
- [7.4.746 : ":\[count\]tag" is not always working](http://code.google.com/p/vim/source/detail?r=b782813c73a61e98d881b481af7fb671585e404b)
- [7.4.747 : ":cnext" may jump to wrong column when 'virtualedit' is "all"](http://code.google.com/p/vim/source/detail?r=d1a87b307a5033bdb3e79cad32fca3db3ae43224)
- [7.4.748 : (after 7.4.745) buffer overflow](http://code.google.com/p/vim/source/detail?r=566e193323cebc0beb02fc2bdda1b8209f30eb62)
- [7.4.749 : (after 7.4.741) two consecutive commas are OK for some options](http://code.google.com/p/vim/source/detail?r=e73d5ac8b16b59f5d26195308ab1822e3ddee5e4)
- [7.4.750 : cannot build with clang 3.5 on Cygwin with perl enabled](http://code.google.com/p/vim/source/detail?r=809e3aa59c6f4c4ecff866b9c0d66bc765289604)
- [7.4.751 : it is not obvious how to enable the address sanitizer](http://code.google.com/p/vim/source/detail?r=db53fd1a7d5c6898bd86bfe203aac8b8d89d9307)
- [7.4.752 : Unicode 8.0 not supported](http://code.google.com/p/vim/source/detail?r=c9a5d51c9161826fb57da3b6c8e5fcb03c45e1cf)
- [7.4.753 : appending in Visual mode with 'linebreak' set is wrong](http://code.google.com/p/vim/source/detail?r=8ebcd7f1ff5b64764b747831c5b903bf1c721774)
- [7.4.754 : using CTRL-A in Visual mode does not work well](http://code.google.com/p/vim/source/detail?r=9798a98a1583fec87694567028d969a5abdd7695)
- [7.4.755 : it is not easy to count the number of characters](http://code.google.com/p/vim/source/detail?r=66ab6ec256d159860a004b29ea4e6074740e5ea4)
- [7.4.756 : can't use strawberry Perl 5.22 x64 on MS-Windows](http://code.google.com/p/vim/source/detail?r=64ff14cbb665a696a4e7eb66c47a9ffaec1d57bd)
- [7.4.757 : cannot detect the background color of a terminal](http://code.google.com/p/vim/source/detail?r=a89cff862dda7668be7915bb9ca2d52c2f752aa5)
- [7.4.758 : 'conceallevel' in command-line window hides first character](http://code.google.com/p/vim/source/detail?r=4ece46045a5aa6167aef99fb4ea7c02f4e74e9d1)
- [7.4.759 : building with Lua 5.3 doesn't work, symbols have changed](http://code.google.com/p/vim/source/detail?r=02466bf84babd8917e717390906aa0ed20c3700b)
- [7.4.760 : spelling mistakes are not displayed after ":syn spell"](http://code.google.com/p/vim/source/detail?r=bcc0909b379e89c7b45ac1eb9be4018bfe0577c9)
- [7.4.761 : (after 7.4.757) incomplete request-background implementation](http://code.google.com/p/vim/source/detail?r=7bd926109eea40d1366ac604b9b29b2c4ba84935)
- [7.4.762 : (after 7.4.757) comment for may&#x5f;req&#x5f;bg&#x5f;color() is wrong](http://code.google.com/p/vim/source/detail?r=ce209c5103a3dd3656076a926f8cf19be3019048)
- [7.4.763 : (after 7.4.759) building with Lua 5.1 doesn't work](http://code.google.com/p/vim/source/detail?r=2b7e36955265a2fc5bed4edf0faead3ae19218f5)
- [7.4.764 : (after 7.4.754) test&#x5f;increment fails on MS-Windows](http://code.google.com/p/vim/source/detail?r=9d41289b55126e08bcea87daf0ed5f1699524b57)

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

