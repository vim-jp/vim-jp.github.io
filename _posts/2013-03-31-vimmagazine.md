---
layout: post
category: vimmagazine
title: Vim Magazine 3 月号
---

## 話題

- [Vim Advent Calendar 2012](http://atnd.org/events/33746)

  Vim Advent Calendar 2012 まだやってます。


## 今月の新機能

- 7.3.859: CUI 端末で 'ambiwidth' が自動設定されるようになりました。
- 7.3.867: 自動コマンドに TextChanged イベントと TextChangedI イベントが追加されました。
- 7.3.872: 'fileignorecase' オプションが追加されました。
- 7.3.873: 置換で簡単にタイトルケースに変換できるようになりました。`:echo substitute('abc', '.*', '\L\u\0', '') => Abc`

## リリース情報

- [7.3.844 : enum is not indented correctly with "public" etc](http://code.google.com/p/vim/source/detail?r=eccc0e8a82fb632efdfa600db3f65f4ed3dff01b)
- [7.3.845 : (after 7.3.844) enum indenting is not tested](http://code.google.com/p/vim/source/detail?r=171234e02b1e74a2482003d56e7110d37fb4a62b)
- [7.3.846 : missing proto file, listing a non-existing proto file](http://code.google.com/p/vim/source/detail?r=b06bb07d906be3be689f90399effd2b3b2ac2ee3)
- [7.3.847 : test 55 fails when messages are translated](http://code.google.com/p/vim/source/detail?r=4dd43d6b3df705a513be75db8a1b50a83f72088d)
- [7.3.848 : can't build with Ruby 2.0 when using MinGW x64 or MSVC10](http://code.google.com/p/vim/source/detail?r=e49b80f267dc3cde78de78026368dec3a20520f2)
- [7.3.849 : ":g//" gives "Pattern not found error" with E486](http://code.google.com/p/vim/source/detail?r=dc77c2a14a05810bda7e53373d76b7f41341f44d)
- [7.3.850 : ":vimgrep //" matches everywhere](http://code.google.com/p/vim/source/detail?r=07fef68eb0188adb3d750b92b78059cf41fc8c7c)
- [7.3.851 : empty pattern with :sort incorrect without previous pattern](http://code.google.com/p/vim/source/detail?r=dd669f296568a7525e02ab54b615bde20d37c466)
- [7.3.852 : system() breaks clipboard text](http://code.google.com/p/vim/source/detail?r=3fd805ca2a6a6ef231f111c6a7b3cd512abaf537)
- [7.3.853 : "ra" does not replace all chars on multi-byte chars > 1 line](http://code.google.com/p/vim/source/detail?r=eea32254940fb4cd61339650e10a808f4e702e62)
- [7.3.854 : after BS in insert mode completion current entry is wrong](http://code.google.com/p/vim/source/detail?r=14cdb6e4d5b88aaaf7e33eaee87efec5b1c1ebf5)
- [7.3.855 : compiler warnings](http://code.google.com/p/vim/source/detail?r=6604853dd8227c7983dc1b10f804f6aa4b28578f)
- [7.3.856 : system() garbles multi-byte clipboard contents](http://code.google.com/p/vim/source/detail?r=fb533fb6c015f40a3b1301c0c876c35665b0d6ef)
- [7.3.857 : QuitPre autocommand event does not trigger for :qa and :wq](http://code.google.com/p/vim/source/detail?r=4905e038b1057e133d0e72c48cf88580cdedf26b)
- [7.3.858 : "gv" selects the wrong area after some operators](http://code.google.com/p/vim/source/detail?r=0a3fab86f34d6bab27e0be69db4f39ffb664557a)
- [7.3.859 : 'ambiwidth' must be set by the user](http://code.google.com/p/vim/source/detail?r=ecf21be84def6033137d168e0a66500c29aeffec)
- [7.3.860 : in --remote-expr try/catch does not work](http://code.google.com/p/vim/source/detail?r=6ce32844727d8d04e8dc91310d0a4118d81aa161)
- [7.3.861 : ":setlocal number" clears global value of 'relativenumber'](http://code.google.com/p/vim/source/detail?r=6fe728bf165e9b6974ef47e86d98b0eb14ce21f3)
- [7.3.862 : dragging the status line can be slow](http://code.google.com/p/vim/source/detail?r=dd4f3ceb5d86f8367d4cfd6fdc898e493c6d0aad)
- [7.3.863 : (after 7.3.859) 'ambiwidth' detection for ANSI terminal](http://code.google.com/p/vim/source/detail?r=5c0652f455c9f2ae72790a782fe5b5fd1a0e857c)
- [7.3.864 : (after 7.3.862) can't build without the mouse feature](http://code.google.com/p/vim/source/detail?r=a82ef37d38c09523bab204253b52e177dbbc0581)
- [7.3.865 : (after 7.3.862) mouse position can be wrong](http://code.google.com/p/vim/source/detail?r=94342b0605fb9111b7fd4a45a8bbc72992c4c255)
- [7.3.866 : not serving the X selection during system() isn't nice](http://code.google.com/p/vim/source/detail?r=90d72df431e54d677f03171cbe558434029c027e)
- [7.3.867 : matchparen does not update match when using auto-indenting](http://code.google.com/p/vim/source/detail?r=0fcb050fd79d84ed190452e6387e42d320678fec)
- [7.3.868 : empty line at hit-return prompt with few lines of text](http://code.google.com/p/vim/source/detail?r=b31989e2b05832d432557a32ed8901dce237dec2)
- [7.3.869 : bufwinnr() matches buffers in other tabs](http://code.google.com/p/vim/source/detail?r=23ce9a61bdc2a35892ec4e4fa22e8964d60dc6cd)
- [7.3.870 : compiler warnings when using MingW 4.5.3](http://code.google.com/p/vim/source/detail?r=26e59a39fdd965ffe83bd654b705540551c8e0c8)
- [7.3.871 : search('\^$', 'c') does not use empty match under the cursor](http://code.google.com/p/vim/source/detail?r=d52c45b35fb0c3b77bbace6cec3deda8100fbb7a)
- [7.3.872 : on some systems file name case is always ignored](http://code.google.com/p/vim/source/detail?r=edd0bc1f26bd019cfd328a0dc2c9a854607e4a41)
- [7.3.873 : cannot easily use :s to make title case](http://code.google.com/p/vim/source/detail?r=7faeece39228fde0ac5c26a362a223d0dbe9bba5)
- [7.3.874 : comparing file names does not handle multi-byte chars properly](http://code.google.com/p/vim/source/detail?r=4d0b39bcec207130e483ffa66b6f51f999d8e95b)
- [7.3.875 : (after 7.3.866) build problem with a combination of features](http://code.google.com/p/vim/source/detail?r=beab15a2ef147b880bf7b5add269142b9af58317)

## 新着スクリプト

- [textobj-gitgutter : text object for a gitgutter hunk](http://www.vim.org/scripts/script.php?script_id=4458)
- [neosnippet : The neo-snippet plugin contains the neocomplcache snippet completion](http://www.vim.org/scripts/script.php?script_id=4459)
- [RelativeNumberCurrentWindow : Only show relative numbers in the currently active window.](http://www.vim.org/scripts/script.php?script_id=4461)
- [ProportionalResize : Adapt the window proportions after Vim is resized.](http://www.vim.org/scripts/script.php?script_id=4462)
- [vim-easy-submode : using vim-submode easily](http://www.vim.org/scripts/script.php?script_id=4463)
- [vim-tayra : Syntax highlighting for tayra scripts.](http://www.vim.org/scripts/script.php?script_id=4464)
- [LineJugglerCommands : Commands to duplicate and move around lines.](http://www.vim.org/scripts/script.php?script_id=4465)
- [vim-transmitty : Transmit.app integration for VIM (Mac only)](http://www.vim.org/scripts/script.php?script_id=4466)
- [vim-multiedit : Multi-selection editing for vim, inspired by Sublime Text](http://www.vim.org/scripts/script.php?script_id=4467)
- [dubs-scheme : a color scheme mainly for programming](http://www.vim.org/scripts/script.php?script_id=4468)
- [vim-forrestgump : Run code inline and on-the-fly in vim. Inspired by TextMate's ⌘R.](http://www.vim.org/scripts/script.php?script_id=4469)
- [hemisu.vim : A color scheme with light and dark variants](http://www.vim.org/scripts/script.php?script_id=4470)
- [noctu.vim : 16-color terminal color scheme](http://www.vim.org/scripts/script.php?script_id=4471)
- [obsession.vim : Continuously updated session files](http://www.vim.org/scripts/script.php?script_id=4472)
- [ghcmod : Happy Haskell programming on Vim, powered by ghc-mod](http://www.vim.org/scripts/script.php?script_id=4473)
- [interactive-replace : InteractiveReplace lets you choose replacement text for your search](http://www.vim.org/scripts/script.php?script_id=4474)
- [voogle : search google from vim easily with any modern browser](http://www.vim.org/scripts/script.php?script_id=4475)
- [vim-scroll-position : Displays the current scroll position like a scrollbar thumb](http://www.vim.org/scripts/script.php?script_id=4476)
- [vim-php-namespace : imports php namespaces](http://www.vim.org/scripts/script.php?script_id=4477)
- [vim-run-live : Run or Preview current buffer in highly configurable ways](http://www.vim.org/scripts/script.php?script_id=4478)
- [iwilldiffer.vim : Show git/hg/bzr add/mode/del signs in buffer](http://www.vim.org/scripts/script.php?script_id=4479)
- [composer.vim : Composer plugin for Vim](http://www.vim.org/scripts/script.php?script_id=4480)
- [joy : \<tab> becomes \<esc>, just like Bill Joy!](http://www.vim.org/scripts/script.php?script_id=4481)
- [vim-tmuxify : Handling tmux panes like a boss.](http://www.vim.org/scripts/script.php?script_id=4482)
- [zarniwoop.vim : A colourful, dark colour scheme for Vim.](http://www.vim.org/scripts/script.php?script_id=4483)
- [assistEncloseManip.vim : Assist manipulation on brackets, quotes](http://www.vim.org/scripts/script.php?script_id=4484)
- [vim-git-log : View your git log interactively in Vim.  Side by side diffs!](http://www.vim.org/scripts/script.php?script_id=4485)
- [vim-rst-sections : Plugin to work with sections in ReST](http://www.vim.org/scripts/script.php?script_id=4486)
- [vim-signify : Advanced plugin for showing VCS diffs in the SignColumn](http://www.vim.org/scripts/script.php?script_id=4487)
- [tbone.vim : tmux basics](http://www.vim.org/scripts/script.php?script_id=4488)
- [last256 : A dark vim color-scheme, based on vim-hybrid](http://www.vim.org/scripts/script.php?script_id=4489)
- [QFGrep.vim : A plugin to filter Quickfix entries](http://www.vim.org/scripts/script.php?script_id=4490)
- [venom : vim python interface extension](http://www.vim.org/scripts/script.php?script_id=4491)
- [mercury : execute any fragment of code directly from vim](http://www.vim.org/scripts/script.php?script_id=4492)
- [vim-rusmode : Bind pressing russian key in normal mode to toggle layout and repeat same key.](http://www.vim.org/scripts/script.php?script_id=4493)
- [Brolink.vim : Live browser editing for Vim](http://www.vim.org/scripts/script.php?script_id=4494)
- [molokai\_nancle : Modify the VertSplit color of the original one](http://www.vim.org/scripts/script.php?script_id=4495)
- [colorful-statusline : A colorful statusline for vim.](http://www.vim.org/scripts/script.php?script_id=4496)
- [Git Support : Inspect the state of a repository and execute Git commands without leaving Vim.](http://www.vim.org/scripts/script.php?script_id=4497)
- [Code Upstairs : A plugin for browsing source code at functions level.](http://www.vim.org/scripts/script.php?script_id=4498)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (3579)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (2484)
3. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (1450)
4. [pathogen.vim : Easy manipulation of 'runtimepath', 'path', 'tags', etc](http://www.vim.org/scripts/script.php?script_id=2332) (1268)
5. [snipMate : TextMate-style snippets for Vim](http://www.vim.org/scripts/script.php?script_id=2540) (1107)
6. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (1075)
7. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (1011)
8. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (922)
9. [SuperTab : Do all your insert-mode completion with Tab.](http://www.vim.org/scripts/script.php?script_id=1643) (892)
10. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (878)

## vim-jp issues

Open : 153 (+8) | Closed : 190 (+14)

- [Issue #322 : バグ報告テンプレ欲しい](https://github.com/vim-jp/issues/issues/322)
- [Issue #323 : syntax include でエラーがでる。](https://github.com/vim-jp/issues/issues/323)
- [Issue #324 : fileencodingsオプションの解析が遅い](https://github.com/vim-jp/issues/issues/324)
- [Issue #325 : /doc のテキストエンコードがバラバラ](https://github.com/vim-jp/issues/issues/325)
- [Issue #326 : search() で stoppos を利用するとカーソルのある位置が検索対象から外れる](https://github.com/vim-jp/issues/issues/326)
- [Issue #327 : Cygwinでtest55が失敗する](https://github.com/vim-jp/issues/issues/327)
- [Issue #328 : perlの正規表現中のextra commentがsyntax highlightされない。](https://github.com/vim-jp/issues/issues/328)
- [Issue #329 : fFtT ハイライト](https://github.com/vim-jp/issues/issues/329)
- [Issue #330 : 設定のハマりどころをなくす (was \*\_example.vim をデフォルトで読むかどうか)](https://github.com/vim-jp/issues/issues/330)
- [Issue #331 : Kaoriya版のシステムのvimrcの副作用が大きすぎる](https://github.com/vim-jp/issues/issues/331)
- [Issue #332 : 7.3.859 が Mac で問題が出てるらしい](https://github.com/vim-jp/issues/issues/332)
- [Issue #333 : Patch 7.3.872で入った 'fileignorecase' オプションの挙動が変かも？](https://github.com/vim-jp/issues/issues/333)
- [Issue #334 : macvim-kaoriyaのみ？で filetype=vim の時だけ正規表現がマッチしない](https://github.com/vim-jp/issues/issues/334)
- [Issue #335 : python3インターフェイスの vim.eval() で日本語のPython文字列を扱えない](https://github.com/vim-jp/issues/issues/335)
- [Issue #336 : poファイル修正](https://github.com/vim-jp/issues/issues/336)
- [Issue #337 : DirectWriteパッチ、完成させるよ。](https://github.com/vim-jp/issues/issues/337)
- [Issue #338 : py3 で Vim が落ちる](https://github.com/vim-jp/issues/issues/338)
- [Issue #339 : 64bit かつ ruby 2.0.0 の環境で、--enable-rubyinterp=dynamic でビルドできない](https://github.com/vim-jp/issues/issues/339)
- [Issue #340 : Lua トランスレーター作ろう](https://github.com/vim-jp/issues/issues/340)
- [Issue #341 : ifdef のインデント修正する?](https://github.com/vim-jp/issues/issues/341)
- [Issue #342 : バッファ系のコマンド(:ls, :bnext, :brewind, etc...)がタブページ使用時を考慮していない](https://github.com/vim-jp/issues/issues/342)
- [Issue #343 : help原文の「viにはない」の表記ゆれの統一を検討する](https://github.com/vim-jp/issues/issues/343)

