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

- [7.3.844 : enum is not indented correctly with "public" etc](https://github.com/vim/vim/commit/7534221ebaba7244f4daaf54bb91206cf45e8a5e)
- [7.3.845 : (after 7.3.844) enum indenting is not tested](https://github.com/vim/vim/commit/d999db2f1e5ec51042883d35ca44f0d751324b19)
- [7.3.846 : missing proto file, listing a non-existing proto file](https://github.com/vim/vim/commit/84e80b51d12f11a95aa76be51abed6a9d0a8c6c5)
- [7.3.847 : test 55 fails when messages are translated](https://github.com/vim/vim/commit/4e3c70d241842d3129d4c4d480c60910029d56af)
- [7.3.848 : can't build with Ruby 2.0 when using MinGW x64 or MSVC10](https://github.com/vim/vim/commit/db3fbe5228025b7f87702e74cfbbb5acf01c7c39)
- [7.3.849 : ":g//" gives "Pattern not found error" with E486](https://github.com/vim/vim/commit/c389fd3a49c01674960fe2153af2f6e42b316c61)
- [7.3.850 : ":vimgrep //" matches everywhere](https://github.com/vim/vim/commit/60abe75379f8b7c2076856c5f12ce2e7650110f6)
- [7.3.851 : empty pattern with :sort incorrect without previous pattern](https://github.com/vim/vim/commit/210f3704f795870c513a136725718e10ac2bdbfc)
- [7.3.852 : system() breaks clipboard text](https://github.com/vim/vim/commit/e2e663f67d9f44ae876659bd492f734d48bd2970)
- [7.3.853 : "ra" does not replace all chars on multi-byte chars > 1 line](https://github.com/vim/vim/commit/db813951c4fc1f91aaaeb4e72d37942e13cb1f0e)
- [7.3.854 : after BS in insert mode completion current entry is wrong](https://github.com/vim/vim/commit/3978e08fbee9a1ef6b988cae73e95cb18467733b)
- [7.3.855 : compiler warnings](https://github.com/vim/vim/commit/b3cb98216296dc166f3d4544dd95d661cefe73b1)
- [7.3.856 : system() garbles multi-byte clipboard contents](https://github.com/vim/vim/commit/1a0316ca2a617c5becc187337c4b9e3f08578d3b)
- [7.3.857 : QuitPre autocommand event does not trigger for :qa and :wq](https://github.com/vim/vim/commit/4f8301f6415e86631dadbc19066ba0bc8550df49)
- [7.3.858 : "gv" selects the wrong area after some operators](https://github.com/vim/vim/commit/a390bb6ad307115c5893788f8f7e1ba5576800e0)
- [7.3.859 : 'ambiwidth' must be set by the user](https://github.com/vim/vim/commit/9584b31aae31d8b031243eab75e26fbaa1755a1f)
- [7.3.860 : in --remote-expr try/catch does not work](https://github.com/vim/vim/commit/1e284f515581e0516e3f3dea568b9b9084bbcab1)
- [7.3.861 : ":setlocal number" clears global value of 'relativenumber'](https://github.com/vim/vim/commit/20754027b3b8c29dfc5ee0b5dfa6a5459ea6b903)
- [7.3.862 : dragging the status line can be slow](https://github.com/vim/vim/commit/2526ef276b97b1a5275fc0039fb45ca7aa7b6fac)
- [7.3.863 : (after 7.3.859) 'ambiwidth' detection for ANSI terminal](https://github.com/vim/vim/commit/e533bbe094c105884696d69e433604fd1641f988)
- [7.3.864 : (after 7.3.862) can't build without the mouse feature](https://github.com/vim/vim/commit/a5be25e1dbb6d129cb7d1efb5c73141a51836e06)
- [7.3.865 : (after 7.3.862) mouse position can be wrong](https://github.com/vim/vim/commit/b897871ce9a4a6bd033bfe127a83eb6348361739)
- [7.3.866 : not serving the X selection during system() isn't nice](https://github.com/vim/vim/commit/090cfc1b02d5dcd98ba414ccf56e69cc38f272bd)
- [7.3.867 : matchparen does not update match when using auto-indenting](https://github.com/vim/vim/commit/186628f67153726edf130015e2a8043a5d309c42)
- [7.3.868 : empty line at hit-return prompt with few lines of text](https://github.com/vim/vim/commit/b59494cab15310c8e2aaf59d48b270282c3e2017)
- [7.3.869 : bufwinnr() matches buffers in other tabs](https://github.com/vim/vim/commit/0c279bbb9c2b9fce1c837a35ace2d4644eced0b8)
- [7.3.870 : compiler warnings when using MingW 4.5.3](https://github.com/vim/vim/commit/af62ff3696514a4f1b595629bcfa584748bc29ce)
- [7.3.871 : search('\^$', 'c') does not use empty match under the cursor](https://github.com/vim/vim/commit/db333a5b8d0c72b7342d6d65ad2895a19a1c29d2)
- [7.3.872 : on some systems file name case is always ignored](https://github.com/vim/vim/commit/71afbfe6cd697de30a9e0d57e5a6434cf4bb0f13)
- [7.3.873 : cannot easily use :s to make title case](https://github.com/vim/vim/commit/c2c355df6f094cdb9e599fd395a78c14486ec697)
- [7.3.874 : comparing file names does not handle multi-byte chars properly](https://github.com/vim/vim/commit/d0e2d94589d11cad514d1d66ade76d497e9d7567)
- [7.3.875 : (after 7.3.866) build problem with a combination of features](https://github.com/vim/vim/commit/95a513511835b017479a513664eeed6d8899fa0d)

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

