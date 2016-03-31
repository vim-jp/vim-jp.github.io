---
layout: post
category: vimmagazine
title: Vim Magazine 1 月号
---

## 話題

- [Vim Advent Calendar 2012](http://atnd.org/events/33746)

  Vim Advent Calendar 2012 まだやってます。

- [Yokohama.vim.reverse #4](http://atnd.org/events/36424)

  座談形式でvimに関する何かに取り組む会 神奈川 2月17日(日)


## 今月の新機能

- 7.3.780 : char2nr() と nr2char() に utf8 引数が追加されました。'encoding' が utf8 でなくてもユニコードの変換ができます。
- 7.3.787 : 'relativenumber' 設定時、カーソル行に現在の行番号が表示されるように変更されました。

## リリース情報

- [7.3.763 : jumping to a mark does not open fold if it is in the same line](https://github.com/vim/vim/commit/8754deb3b622186ff5efa0e5547a392aaed6a1c7)
- [7.3.764 : not all message translation files are installed](https://github.com/vim/vim/commit/c4fba6f8811d8f205f5c9772d4797691e615c6a2)
- [7.3.765 : segfault when doing "cclose" on BufUnload in a python function](https://github.com/vim/vim/commit/802418d5eb5c98dcbe53390d1eceba92dd097aa4)
- [7.3.766 : ":help cpo-\*" jumps to the wrong place](https://github.com/vim/vim/commit/117f2c4b91fb3675d56be36c1d972f0fb4fda13a)
- [7.3.767 : (Win32) The \_errno used for iconv may be the wrong one](https://github.com/vim/vim/commit/8fae8e658ffd8f5177a22e41ece18ea32edaa4a2)
- [7.3.768 : settabvar() and setwinvar() may move the cursor](https://github.com/vim/vim/commit/3e37fd0950081e277fac44a0bd8e60815898945c)
- [7.3.769 : 'matchpairs' does not work with multi-byte characters](https://github.com/vim/vim/commit/8c7694a8649cbd5ce91c150737b1f52b20631cb4)
- [7.3.770 : vim.h indentation is inconsistent](https://github.com/vim/vim/commit/08ed30eca7d64e623a5363a6a0eac79af48ec1c0)
- [7.3.771 : (after 7.3.769) uninitialized variable](https://github.com/vim/vim/commit/09365025386ac6af4710517a89fe0c4ed2bc86e4)
- [7.3.772 : cursor at wrong location after doing %s/x/y/c](https://github.com/vim/vim/commit/8c4fbd1a153245fd3c496507693140c0e8e2f260)
- [7.3.773 : (after 7.3.767) crash when OriginalFirstThunk is zero](https://github.com/vim/vim/commit/b5f7bf6ed9c59b8841efcd41c2ed1687f080f4cf)
- [7.3.774 : tiny GUI version misses console dialog feature](https://github.com/vim/vim/commit/f939c4e6b176d89626eff5fe9a0b01a5bbb40ef5)
- [7.3.775 : Cygwin and Mingw builds miss dependency on gui\_w48.c](https://github.com/vim/vim/commit/4da70dbc4ddef9b2a2815cc0f6f09a283b65e671)
- [7.3.776 : ml\_get error when searching, caused by curwin/curbuf mismatch](https://github.com/vim/vim/commit/9d182dd0a666f7cb706cd54dd78df3f59a160fff)
- [7.3.777 : when building with Gnome locale gets reset](https://github.com/vim/vim/commit/3be71ce28d832553a1350b15d9323c30014382c0)
- [7.3.778 : compiler error for adding up two pointers](https://github.com/vim/vim/commit/94a8adfdd890f0550f81755b247ca7cc5336ad20)
- [7.3.779 : backwards search lands in wrong place when on a multibyte char](https://github.com/vim/vim/commit/55b7b7eeb50d4fab541ba151e0f653eb6e3852b8)
- [7.3.780 : char2nr() and nr2char() always use 'encoding'](https://github.com/vim/vim/commit/d35d784e9106a9d398b444ed9557664bbcb6a65c)
- [7.3.781 : drawing with 'guifontwide' can be slow](https://github.com/vim/vim/commit/459339662566bcf5ae9892c7dabb9518e7f70f83)
- [7.3.782 : Windows: IME composition may use a wrong font](https://github.com/vim/vim/commit/0f2721223e5b97ac250d729b2359471241a3ab7d)
- [7.3.783 : crash when mark is not set](https://github.com/vim/vim/commit/15364d74ca33d8d722483c7b8ebede17f344ab8c)
- [7.3.784 : (after 7.3.781) error when 'guifontwide' has a comma](https://github.com/vim/vim/commit/826763f87b4a15fb125b3b5cc3fc308a5976bfaf)
- [7.3.785 : (after 7.3.776) crash with specific use of search pattern](https://github.com/vim/vim/commit/2f315ab7bec2713d2fb986c5d863f15ab9949166)
- [7.3.786 : Python threads don't run in the background (issue 103)](https://github.com/vim/vim/commit/02366255c9bd16120c5025106f7ac5d5a06774d7)
- [7.3.787 : with 'relativenumber' can't see the absolute line number](https://github.com/vim/vim/commit/700e7345ded4e2af25e63409b81c1ce516410a28)
- [7.3.788 : when only using patches build fails on missing nl.po](https://github.com/vim/vim/commit/8d616172389a666e74c47c4e41f5cd9a17a38268)
- [7.3.789 : (after 7.3.776) "\\k" in regexp does not work in other window](https://github.com/vim/vim/commit/f813a1883f767ad9c4a7a44b5f42cee312d53bc0)
- [7.3.790 : after reloading a buffer the modelines are not processed](https://github.com/vim/vim/commit/52f85b730e379c4519b4a9ef1f008417150713f7)
- [7.3.791 : MzScheme interface doesn't work propely](https://github.com/vim/vim/commit/7567646f13beb8c7aad79340f30ad2b5d5c8225f)
- [7.3.792 : ":substitute" works differently without confirmation](https://github.com/vim/vim/commit/4bc8cf0e7c4c1d2c9d99ac8b3c080b50f239406d)
- [7.3.793 : (after 7.3.792) new :substutite behavior is not tested](https://github.com/vim/vim/commit/8ceb1735a2e845c7987f01a6319775b39c4b0177)
- [7.3.794 : tiny build fails](https://github.com/vim/vim/commit/a50e5866b019616eacdc4498fd94ee3b0795b80d)
- [7.3.795 : MzScheme does not build with tiny features](https://github.com/vim/vim/commit/d2142213754df16fe11a2674716a7c33a012fa07)
- [7.3.796 : "/\[&#94;\\n\]" does match at a line break](https://github.com/vim/vim/commit/e337e5f634dcb0e74999d67286125a85423e38a9)
- [7.3.797 : (after 7.3.792) compiler warning for size\_t to int conversion](https://github.com/vim/vim/commit/04e5b5a22823fdf30e35c383aea8843f7c747eec)

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
- [Issue #297 : if\_mruby の可能性について](https://github.com/vim-jp/issues/issues/297)
- [Issue #298 : 別バッファに移動するとハイライトがおかしくなる](https://github.com/vim-jp/issues/issues/298)
- [Issue #299 : GVim で set visualbell set t\_vb='' したら描画がおかしくなる ](https://github.com/vim-jp/issues/issues/299)
- [Issue #300 : if\_ruby 複数バージョン対応化](https://github.com/vim-jp/issues/issues/300)
- [Issue #301 : Debian系のLinuxでPython 2.xと3.xが同時利用できない問題の原因と対策](https://github.com/vim-jp/issues/issues/301)

