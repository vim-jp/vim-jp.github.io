---
layout: post
category: vimmagazine
title: Vim Magazine 12 月号
---

## 今月の新機能

- 7.4.122 : Windowsで `enc=utf-8` の時、マルチバイト文字が外部コマンドに正しく渡されるようになりました。

## リリース情報

- [7.4.111 : memory leak in Python OptionsAssItem](https://github.com/vim/vim/commit/c2401d62e0f8de45598c8ed7d9e9b09227f6682d)
- [7.4.112 : MS-Windows: defaults for 'dir' and 'bdir' do not include $TEMP](https://github.com/vim/vim/commit/b050c6418647c5358564e4608e8bb6ac7caa93b8)
- [7.4.113 : MSVC static analysis gives warnings](https://github.com/vim/vim/commit/73c61633b6c0387333f208afb11a0a00bed09ef0)
- [7.4.114 : new GNU make directory change messages are different](https://github.com/vim/vim/commit/5d7f9df8fa12e076c233069bb0a98d490749a63a)
- [7.4.115 : Zsh: expanding ~abc fails when the result contains a space](https://github.com/vim/vim/commit/b011af969652f6be0775ba8dcd54aac4dcb57cdb)
- [7.4.116 : 'showcmd' does not show a typed space](https://github.com/vim/vim/commit/7ba0741451d8eea497fba07ca2bbf6fae107a6a6)
- [7.4.117 : can't build with Cygwin/MingW and Perl 5.18](https://github.com/vim/vim/commit/6b10721fa2d883540c1c106ef26fff8abed0ac43)
- [7.4.118 : redrawing status lines may causes recursive call](https://github.com/vim/vim/commit/1d633413e5961589c2ae81300c96197443eee0c8)
- [7.4.119 : Vim doesn't work well on OpenVMS](https://github.com/vim/vim/commit/4ffa07081f97db6b49d6d087ce46ff0b3a4c8a5c)
- [7.4.120 : (after 7.4.117) can't build with Perl 5.18 on Linux](https://github.com/vim/vim/commit/fcf94d399015984742c4306c7e6676f07d7ed5fe)
- [7.4.121 : completion doesn't work for ":py3d" and ":py3f"](https://github.com/vim/vim/commit/893b2d73d16870f08d7d442488f3790156e11105)
- [7.4.122 : Win32: :grep doesn't work when 'encoding' and ACP differ](https://github.com/vim/vim/commit/910cffbb5ea2df88688301168eb04926c9e68990)
- [7.4.123 : Win32: Getting user name does not use wide function](https://github.com/vim/vim/commit/c8020ee825b9d9196b1329c0e097424576fc9b3a)
- [7.4.124 : Win32: Getting host name does not use wide function](https://github.com/vim/vim/commit/2cc8738f2f87ffa3e584b8e606f11fbb188c46b4)
- [7.4.125 : Win32: Dealing with messages may not work for multi-byte chars](https://github.com/vim/vim/commit/175d0704dfb1247be7f414bfbcf9f5cfeced2794)
- [7.4.126 : compiler warnings for "const" and incompatible types](https://github.com/vim/vim/commit/36c85b2c4639c06fd8e4398f1f7d29f5d4e0405f)
- [7.4.127 : Perl 5.18 on Unix doesn't work](https://github.com/vim/vim/commit/7c0daf088e880e7cd4031c673876a2750323c79c)
- [7.4.128 : Perl 5.18 for MSVC doesn't work](https://github.com/vim/vim/commit/207fd75cac7bc658da039f551620d51c422def44)
- [7.4.129 : getline(-1) returns zero](https://github.com/vim/vim/commit/959a143f2e17d81c108b60493815419c7772aa63)
- [7.4.130 : relative line numbers mix up windows when using folds](https://github.com/vim/vim/commit/0bd7b3f4ce0e70ff762689dc0fd9f41b54692694)
- [7.4.131 : syncbind causes E315 errors in some situations](https://github.com/vim/vim/commit/dedd1b0aba7e7125081fd3e71b629058d48dac1d)

## 新着スクリプト

- [goyo.vim : Distraction-free writing in Vim](http://www.vim.org/scripts/script.php?script_id=4785)
- [update-time : Automatic update Last Change time as you like!](http://www.vim.org/scripts/script.php?script_id=4786)
- [vim-svngutter : A plugin to put SVN diff info into the gutter ](http://www.vim.org/scripts/script.php?script_id=4787)
- [Vim-Kompleter : A smart, fast, simple, and reliable keyword completion replacement for Vim](http://www.vim.org/scripts/script.php?script_id=4788)
- [Vim-CtrlSpace : Vim Workspace Controller](http://www.vim.org/scripts/script.php?script_id=4789)
- [gsl.vim : Resaltador de sintaxis para GNU SL](http://www.vim.org/scripts/script.php?script_id=4790)
- [SwissCalc : A powerful interactive calculator for geeks](http://www.vim.org/scripts/script.php?script_id=4791)
- [netgrep : vim plugin for grepping and finding on a remote server](http://www.vim.org/scripts/script.php?script_id=4792)
- [easydir.vim : An easy way to create directories and files at the same time!](http://www.vim.org/scripts/script.php?script_id=4793)
- [AddCppClass&#x5f;Plus : this plugin is to help you add a c++ class conveniently](http://www.vim.org/scripts/script.php?script_id=4794)
- [ExtractMatches : Yank matches from range into a register.](http://www.vim.org/scripts/script.php?script_id=4795)
- [vim-nfo : ftplugin for .nfo files](http://www.vim.org/scripts/script.php?script_id=4796)
- [CtrlP-SmartTabs : Vim CtrlP plugin to switch between opened tabs](http://www.vim.org/scripts/script.php?script_id=4797)
- [swamplight : light color scheme](http://www.vim.org/scripts/script.php?script_id=4798)
- [Include Fixer for C/C++ : Categorizes and Alphabetizes C/C++ #include directives](http://www.vim.org/scripts/script.php?script_id=4799)
- [Investigate.vim : A plugin for viewing documentation](http://www.vim.org/scripts/script.php?script_id=4800)
- [vim-skuld : A pomodoro timer that runs inside Vim.](http://www.vim.org/scripts/script.php?script_id=4801)
- [vim-javacompleteex : better java omni-completion.](http://www.vim.org/scripts/script.php?script_id=4802)
- [ft-cmake : dict and snippet for cmake files (e.g. CMakeLists.txt)](http://www.vim.org/scripts/script.php?script_id=4803)
- [InnerFragmentComplete : Insert mode completion based on fragments inside words.](http://www.vim.org/scripts/script.php?script_id=4804)
- [MultiWordComplete : Insert mode completion that completes a sequence of words based on anchors.](http://www.vim.org/scripts/script.php?script_id=4805)
- [vim-systemverilog : system verilog indent, syntax, and ftdetect](http://www.vim.org/scripts/script.php?script_id=4806)
- [Accurev diff tool : Tool to view the changes that you make in accurev workspace](http://www.vim.org/scripts/script.php?script_id=4807)
- [auto-pairs-gentle : Gentle version of Auto Pairs plugin: Insert or delete brackets, parens in pair](http://www.vim.org/scripts/script.php?script_id=4808)
- [vim-sneak : motion improved](http://www.vim.org/scripts/script.php?script_id=4809)
- [LEF&#x5f;DEF&#x5f;syntax : Design/Library Exchange Formats syntax files](http://www.vim.org/scripts/script.php?script_id=4810)
- [vim-quotable : Extending Vim to better support typographic (curly) quote characters](http://www.vim.org/scripts/script.php?script_id=4811)
- [vim-sexp : Precision Editing for S-expressions](http://www.vim.org/scripts/script.php?script_id=4812)
- [candyVirus.vim : An elegant dark 256 colors colorscheme for cterm](http://www.vim.org/scripts/script.php?script_id=4813)
- [vim-howdoi : Howdoi VIM Plugin](http://www.vim.org/scripts/script.php?script_id=4814)
- [choowsewin.vim : land to window you choose like tmux's 'display-pane'.](http://www.vim.org/scripts/script.php?script_id=4815)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (2605)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (2273)
3. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (1383)
4. [pathogen.vim : Easy manipulation of 'runtimepath', 'path', 'tags', etc](http://www.vim.org/scripts/script.php?script_id=2332) (1198)
5. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (1022)
6. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (916)
7. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (888)
8. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (821)
9. [snipMate : TextMate-style snippets for Vim](http://www.vim.org/scripts/script.php?script_id=2540) (804)
10. [winmanager : A windows style IDE for Vim 6.0](http://www.vim.org/scripts/script.php?script_id=95) (802)

## vim-jp issues

Open : 187 (+10) | Closed : 320 (+6)

- [Issue #492 : `:h mode()@en`内にtypo。（表記が統一されていない。）](https://github.com/vim-jp/issues/issues/492)
- [Issue #493 : maparg()が実際に返すDictionaryとhelpに違いがある](https://github.com/vim-jp/issues/issues/493)
- [Issue #494 : Windowsコンソール版でCTRL-&#x5e;が効かない。](https://github.com/vim-jp/issues/issues/494)
- [Issue #495 : pastetoggleした時にstatuslineが評価されない](https://github.com/vim-jp/issues/issues/495)
- [Issue #496 : removing bevel (inset) borders.](https://github.com/vim-jp/issues/issues/496)
- [Issue #497 : test100 is failed on Windows](https://github.com/vim-jp/issues/issues/497)
- [Issue #498 : v:oldfilesの状態がおかしい。](https://github.com/vim-jp/issues/issues/498)
- [Issue #499 : quickfixウィンドウだけになったときにbwipeout!する設定すると、Vimが死んでしまう。](https://github.com/vim-jp/issues/issues/499)
- [Issue #500 : :keeppatterns で @/ も変更しないようにして欲しい](https://github.com/vim-jp/issues/issues/500)
- [Issue #501 : enc=utf-8にするとフォント名にマルチバイト文字が使えない](https://github.com/vim-jp/issues/issues/501)
- [Issue #502 : virtualedit=block で矩形選択中に表示行移動がおかしい場合がある](https://github.com/vim-jp/issues/issues/502)
- [Issue #503 : substitute('string', '.\\zs', 'a', 'g') の結果が 7.4.045 以前と以降で異なる](https://github.com/vim-jp/issues/issues/503)
- [Issue #504 : runtime/menu.vim内のtypo. Se&T Compiler](https://github.com/vim-jp/issues/issues/504)
- [Issue #505 : getregtype()の挙動がヘルプファイルに書かれているものと違う](https://github.com/vim-jp/issues/issues/505)
- [Issue #506 : mingw4.6.1でビルドしたlua5x.dllとneocomplete](https://github.com/vim-jp/issues/issues/506)
- [Issue #507 : autoload 変数に関数参照が設定されている場合に autoload 関数を呼ぶと内部エラーが発生する](https://github.com/vim-jp/issues/issues/507)

