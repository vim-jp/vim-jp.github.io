---
layout: post
category: vimmagazine
title: Vim Magazine 1 月号
---

# Vim Magazine 1 月号

## 話題

- [Vim Advent Calendar 2012](http://atnd.org/events/33746)

  Vim Advent Calendar 2012 まだやってます。

- [Yokohama.vim.reverse #4](http://atnd.org/events/36424)

  座談形式でvimに関する何かに取り組む会 神奈川 2月17日(日)


## 今月の新機能

- 7.3.780 : char2nr() と nr2char() に utf8 引数が追加されました。'encoding' が utf8 でなくてもユニコードの変換ができます。
- 7.3.787 : 'relativenumber' 設定時、カーソル行に現在の行番号が表示されるように変更されました。

## リリース情報

- [7.3.763 : jumping to a mark does not open fold if it is in the same line](http://code.google.com/p/vim/source/detail?r=b6148e4a9c397b5dd1716cdac2f96730bd0f353d)
- [7.3.764 : not all message translation files are installed](http://code.google.com/p/vim/source/detail?r=b29e55d0ef01c506ec48c0c5782e669bb1d80f14)
- [7.3.765 : segfault when doing "cclose" on BufUnload in a python function](http://code.google.com/p/vim/source/detail?r=915fb3d2dd8ffc322a5cc9e294d9bb7dfa8c5a83)
- [7.3.766 : ":help cpo-*" jumps to the wrong place](http://code.google.com/p/vim/source/detail?r=b5103b7d9e1284ef9ffbc7d1170d9105c0769c4d)
- [7.3.767 : (Win32) The _errno used for iconv may be the wrong one](http://code.google.com/p/vim/source/detail?r=be1cffa1e477c4a96bb54a3097d3369b23f32825)
- [7.3.768 : settabvar() and setwinvar() may move the cursor](http://code.google.com/p/vim/source/detail?r=9ea81cd9b1f50054629547ce2eccfac1fa00d096)
- [7.3.769 : 'matchpairs' does not work with multi-byte characters](http://code.google.com/p/vim/source/detail?r=d179a8eff9d7b0adc561f4a4dcfb0b7612c1f301)
- [7.3.770 : vim.h indentation is inconsistent](http://code.google.com/p/vim/source/detail?r=da7a7ba36ed2a51f28e4d3f7e1fd550303e3ece8)
- [7.3.771 : (after 7.3.769) uninitialized variable](http://code.google.com/p/vim/source/detail?r=3bd2c055319e086665af5e449498a4d5bf40f25e)
- [7.3.772 : cursor at wrong location after doing %s/x/y/c](http://code.google.com/p/vim/source/detail?r=d828cab6964f4249db54bff6df0fe02d0f290387)
- [7.3.773 : (after 7.3.767) crash when OriginalFirstThunk is zero](http://code.google.com/p/vim/source/detail?r=3db9aee957f7eea6729e5bfd294e3a24d41c79e4)
- [7.3.774 : tiny GUI version misses console dialog feature](http://code.google.com/p/vim/source/detail?r=a8f5876e498125f4916c5f7168823885c27dad8c)
- [7.3.775 : Cygwin and Mingw builds miss dependency on gui_w48.c](http://code.google.com/p/vim/source/detail?r=a6e60677aea727622dde15d1306c61588b805500)
- [7.3.776 : ml_get error when searching, caused by curwin/curbuf mismatch](http://code.google.com/p/vim/source/detail?r=80b041b994d156360992666110964e74f7473d3e)
- [7.3.777 : when building with Gnome locale gets reset](http://code.google.com/p/vim/source/detail?r=9a1a4c486e6bb367c639c29c6120d5fb1fadcc38)
- [7.3.778 : compiler error for adding up two pointers](http://code.google.com/p/vim/source/detail?r=562bbee33aa4f9896cb23ed6de8fbf60cac323e0)
- [7.3.779 : backwards search lands in wrong place when on a multibyte char](http://code.google.com/p/vim/source/detail?r=2067ed78d37c331685ef8f293d6a23c372ca4b78)
- [7.3.780 : char2nr() and nr2char() always use 'encoding'](http://code.google.com/p/vim/source/detail?r=42edc5dac33c92876d93f00f8ab8c1f6bb25d3fc)
- [7.3.781 : drawing with 'guifontwide' can be slow](http://code.google.com/p/vim/source/detail?r=29f29e86602e4a7f83d4572d2fe75b9b0de2752f)
- [7.3.782 : Windows: IME composition may use a wrong font](http://code.google.com/p/vim/source/detail?r=0de969850c0687d2018e025f31e5cef9e9ce1444)
- [7.3.783 : crash when mark is not set](http://code.google.com/p/vim/source/detail?r=a80af62d5e3b3614ff9e058f42647c99bffe7481)
- [7.3.784 : (after 7.3.781) error when 'guifontwide' has a comma](http://code.google.com/p/vim/source/detail?r=fbadf0f6987790d98c3b3fd935fea0494743830a)
- [7.3.785 : (after 7.3.776) crash with specific use of search pattern](http://code.google.com/p/vim/source/detail?r=be4baed0c933e1f1543ea175e2e36551ff6caa18)
- [7.3.786 : Python threads don't run in the background (issue 103)](http://code.google.com/p/vim/source/detail?r=8b3e88bab70215a18768fca5a90c6f07ac1aa5eb)
- [7.3.787 : with 'relativenumber' can't see the absolute line number](http://code.google.com/p/vim/source/detail?r=e1b98967a985f4d1ad889c8e69b822d9126411ab)
- [7.3.788 : when only using patches build fails on missing nl.po](http://code.google.com/p/vim/source/detail?r=386fc822dd002958fdade4329845013cbfaa3dab)
- [7.3.789 : (after 7.3.776) "\\k" in regexp does not work in other window](http://code.google.com/p/vim/source/detail?r=43c15135926fb95fa82b01424677a9c37f4d0e46)
- [7.3.790 : after reloading a buffer the modelines are not processed](http://code.google.com/p/vim/source/detail?r=8819ea6dcb72dcf61fd0d20f8ea1904a7546917f)
- [7.3.791 : MzScheme interface doesn't work propely](http://code.google.com/p/vim/source/detail?r=16e8a09e8ab08d2226dc84b12b7e434bda6b4fa9)
- [7.3.792 : ":substitute" works differently without confirmation](http://code.google.com/p/vim/source/detail?r=7061704c2014e4563efa032672264a47b4b7b0d9)
- [7.3.793 : (after 7.3.792) new :substutite behavior is not tested](http://code.google.com/p/vim/source/detail?r=fdbded463fdcc861630b3ea61b9654374847231a)
- [7.3.794 : tiny build fails](http://code.google.com/p/vim/source/detail?r=9867f92c9b9fb082ee6daca816f859b53ed8e58b)
- [7.3.795 : MzScheme does not build with tiny features](http://code.google.com/p/vim/source/detail?r=17a312ba38a0b606cfd64d82ed2ebc844828085f)
- [7.3.796 : "/\[^\\n\]" does match at a line break](http://code.google.com/p/vim/source/detail?r=2c12c4e2fae8db60fbdf58b34c8b73eacdc32665)
- [7.3.797 : (after 7.3.792) compiler warning for size_t to int conversion](http://code.google.com/p/vim/source/detail?r=0cbea05f19c73c01dbd0f04378a65ab7de6719e7)

## 新着スクリプト

- [smartpairs : Fantastic selections for Vim](http://www.vim.org/scripts/script.php?script_id=4378)
- [paren-crosshairs : Displays crosshairs on the cursor when moving over parens ](http://www.vim.org/scripts/script.php?script_id=4379)
- [basicXmlParser : Create object trees and save them as xml, or create trees from xml files.](http://www.vim.org/scripts/script.php?script_id=4380)
- [RepeatLast.vim : Provides \<count>\\. to repeat the last group of actions you performed.](http://www.vim.org/scripts/script.php?script_id=4381)
- [asm2d-vim : Vim plugin for the asm2d assembly language](http://www.vim.org/scripts/script.php?script_id=4382)
- [configurationUtility.vim : Save and load configuration files in a dictionary format with sections and items](http://www.vim.org/scripts/script.php?script_id=4383)
- [File Pirate : A very fast file picker plug-in](http://www.vim.org/scripts/script.php?script_id=4384)
- [Bufstop : Fast and efficient buffer switching for Vim](http://www.vim.org/scripts/script.php?script_id=4385)
- [Smartgf : 'goto file' on steroids!](http://www.vim.org/scripts/script.php?script_id=4386)
- [projectroot : Automatic guessing of projectroots](http://www.vim.org/scripts/script.php?script_id=4387)
- [CoqIDE : Emulate Coq IDE in VIM](http://www.vim.org/scripts/script.php?script_id=4388)
- [ljyat : base on molokai, add my change](http://www.vim.org/scripts/script.php?script_id=4389)
- [ferm.vim : Syntax highlighting for ferm; "For Easy Rule Making", a frontend for iptables.](http://www.vim.org/scripts/script.php?script_id=4390)
- [sensible.vim : Defaults everyone can agree on](http://www.vim.org/scripts/script.php?script_id=4391)
- [boshiamy-cue.vim : Utilities for learning Boshiamy input method in traditional Chinese on vim](http://www.vim.org/scripts/script.php?script_id=4392)
- [boshiamy input method keymap file : This is the keymap file for Boshiamy chinese input method](http://www.vim.org/scripts/script.php?script_id=4393)
- [scriptease.vim : A Vim plugin for Vim plugins](http://www.vim.org/scripts/script.php?script_id=4394)
- [jk-jumps : Handle movements with j/k as if they were real jumps](http://www.vim.org/scripts/script.php?script_id=4395)
- [RocketJump : Jump to a visible line using alphabetic signs instead of line numbers](http://www.vim.org/scripts/script.php?script_id=4396)
- [CF Utils : ColdFusion (CF) syntax highlighting and utilities for Vim](http://www.vim.org/scripts/script.php?script_id=4397)
- [(color scheme) Faded-Black : A dark color scheme that looks as though it's been through the wash.](http://www.vim.org/scripts/script.php?script_id=4398)
- [(color scheme) Faded-White : A light color scheme that looks as though it's been through the wash.](http://www.vim.org/scripts/script.php?script_id=4399)
- [Tube : MacVim and iTerm or Terminal interaction without leaving MacVim](http://www.vim.org/scripts/script.php?script_id=4400)
- [vim-sunshine : A vim colorscheme](http://www.vim.org/scripts/script.php?script_id=4401)
- [vim-darksea : A vim colorscheme](http://www.vim.org/scripts/script.php?script_id=4402)
- [togglecursor : Changes cursor shape in supported terminals.](http://www.vim.org/scripts/script.php?script_id=4403)
- [vim-line-numbers : adds line-numbers & ruler to most edit files](http://www.vim.org/scripts/script.php?script_id=4404)
- [latex-support.vim : LaTeX-IDE for Vim/gVim](http://www.vim.org/scripts/script.php?script_id=4405)
- [lithochromatic : A clarity oriented, consistent vim color theme.](http://www.vim.org/scripts/script.php?script_id=4406)
- [haskell.vim : Another Haskell indent plugin.](http://www.vim.org/scripts/script.php?script_id=4407)
- [accelerated-jk : accelerate 'j' and 'k' moving](http://www.vim.org/scripts/script.php?script_id=4408)
- [Vimpanel : A modern side panel for Vim](http://www.vim.org/scripts/script.php?script_id=4409)
- [characterize.vim : Unicode character metadata](http://www.vim.org/scripts/script.php?script_id=4410)
- [metaweblog : Submit your RST article in VIM  to Blog with metaweblog interface.](http://www.vim.org/scripts/script.php?script_id=4411)
- [cljfold.vim : Automagic folding of λ clojure code based on 'foldwords'](http://www.vim.org/scripts/script.php?script_id=4412)
- [darkgit.vim : dark theme inspired in github.vim colorscheme](http://www.vim.org/scripts/script.php?script_id=4413)
- [GlobalOptions : Turn global options into buffer- or window-local ones.](http://www.vim.org/scripts/script.php?script_id=4414)
- [awk-support.vim : awk-support.vim : AWK-IDE for Vim/gVim](http://www.vim.org/scripts/script.php?script_id=4415)
- [vim-seek : Seek makes navigating long lines effortless, acting like f but taking two chars](http://www.vim.org/scripts/script.php?script_id=4416)
- [holokai : A slightly dark version of molokai](http://www.vim.org/scripts/script.php?script_id=4417)
- [Bck : Enhanced searching for Vim](http://www.vim.org/scripts/script.php?script_id=4418)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (702)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (375)
3. [snipMate : TextMate-style snippets for Vim](http://www.vim.org/scripts/script.php?script_id=2540) (225)
4. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (209)
5. [bufexplorer.zip : Buffer Explorer / Browser](http://www.vim.org/scripts/script.php?script_id=42) (192)
6. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (182)
7. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (159)
8. [pathogen.vim : Easy manipulation of 'runtimepath', 'path', 'tags', etc](http://www.vim.org/scripts/script.php?script_id=2332) (157)
9. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (156)
10. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (155)

## vim-jp issues

Open : 132 (-5) | Closed : 169 (+17)

- [Issue #290 : タグジャンプがSEGVで落ちる事がある](https://github.com/vim-jp/issues/issues/290)
- [Issue #291 : W11発生時のファイル再読み込みでモードラインが無視される](https://github.com/vim-jp/issues/issues/291)
- [Issue #292 : Kaoriya 版の Vim(GUI) で :set ft すると 「E218: autocommandの入れ子が深過ぎます 」が発生する](https://github.com/vim-jp/issues/issues/292)
- [Issue #293 : gfw 設定時の描画処理が潜在的に遅い](https://github.com/vim-jp/issues/issues/293)
- [Issue #294 : gfw の該当文字に斜体や太字が効かない](https://github.com/vim-jp/issues/issues/294)
- [Issue #295 : netrwが重かったのでパッチを書きました。](https://github.com/vim-jp/issues/issues/295)
- [Issue #296 : netrwがヤバい](https://github.com/vim-jp/issues/issues/296)
- [Issue #297 : if_mruby の可能性について](https://github.com/vim-jp/issues/issues/297)
- [Issue #298 : 別バッファに移動するとハイライトがおかしくなる](https://github.com/vim-jp/issues/issues/298)
- [Issue #299 : GVim で set visualbell set t_vb='' したら描画がおかしくなる ](https://github.com/vim-jp/issues/issues/299)
- [Issue #300 : if_ruby 複数バージョン対応化](https://github.com/vim-jp/issues/issues/300)
- [Issue #301 : Debian系のLinuxでPython 2.xと3.xが同時利用できない問題の原因と対策](https://github.com/vim-jp/issues/issues/301)

