---
layout: post
category: vimmagazine
title: Vim Magazine 8 月号
---

## 話題

- [Vim 7.4 released!](https://groups.google.com/d/msg/vim_announce/knOQ_t_H5to/STMi8d25ii8J)

  Vim 7.4 がリリースされました。

- [「実践Vim」 が出版されました](http://ascii.asciimw.jp/books/books/detail/978-4-04-891659-2.shtml)

  Vimのコア機能を徹底解説

  古くて新しいUnixのエディタVim--このVimのエディタとしてのコア機能をマスターするためのレシピ集。

- [Snake Vim Trainer](http://www.vimsnake.com/)

  Vim の hjkl を練習できるミニゲーム。

- [nagoya.vim #1](http://connpass.com/event/3179/)

  今あつい国、なごやで Vim の勉強をします! 9月7日(土) 名古屋

- <del>[2013 Emacs/Vim温泉（希望調査）](http://connpass.com/event/2773/)</del>
  [2013 Emacs/Vim温泉（本登録）](http://connpass.com/event/3233/)

  Emacs/Vimなど、エディタを肴にリゾートしてみます。
  何かを作ってみたり、発表してみたり、ただ単に観光したり温泉入るなど、何でも良し。
  九重・湯布院周辺 9月21日(土)-9月23日(月) ※日程調整中

- [Vim Advent Calendar 2012](http://atnd.org/events/33746)

  Vim Advent Calendar 2012 まだやってます。

## リリース情報

- [7.4a.047 : some comments are not so nice](https://github.com/vim/vim/commit/1a4a75c5dcc4a255bbaa60ea6aade5ffdd32e4d9)
- [7.4b.000 : Update files for the 7.4b BETA release.](https://github.com/vim/vim/commit/1a42b4befb5307d20dc6ff4685116e085ece3b1f)
- [7.4b.001 : Win32: dialog may extend off-screen](https://github.com/vim/vim/commit/1b1b09449fffbbfc7598101f515819407a14d920)
- [7.4b.002 : crash searching for \\%()\*](https://github.com/vim/vim/commit/398d53decf7e02550b45a6f5c46fa02e08bb99fe)
- [7.4b.003 : regexp code is not nicely aligned](https://github.com/vim/vim/commit/6dbe68cd9ec2e8904ecf5da9f2e729835bfd0329)
- [7.4b.004 : regexp crash on pattern "@\\%\[\\w\\-\]\*"](https://github.com/vim/vim/commit/eec3e1e72a3d4e31c9b9546cd62fcff69c12ce29)
- [7.4b.005 : finding %s in shellpipe and shellredir does not ignore %%s](https://github.com/vim/vim/commit/542805a59df19db785dc631f174a15eec40911af)
- [7.4b.006 : Using \\{n,m} in an autocommand pattern no longer works](https://github.com/vim/vim/commit/a946afe075ce041a638194b6b3bcb100ba2b3689)
- [7.4b.007 : on 32 bit MS-Windows :perldo does not work](https://github.com/vim/vim/commit/aee1f4a5b5065da6602fb30e01529644e85eb036)
- [7.4b.008 : 'autochdir' causes setbufvar() to change current directory](https://github.com/vim/vim/commit/4008f4fb08f969c87e8f13bd9e019790757e08ad)
- [7.4b.009 : manually defined Visual area doesn't work in new buffer](https://github.com/vim/vim/commit/bc88a275375c65538c2eaadb55d4ddc26346f695)
- [7.4b.010 : Win32: Tcl library load does not use standard mechanism](https://github.com/vim/vim/commit/364ab2ffd3ecc88f9547c28d51c138d53fead488)
- [7.4b.011 : ":he \\%(\\)" does not work](https://github.com/vim/vim/commit/c528b1d1edb5ab87242704c82422ae363647603c)
- [7.4b.012 : output from a shell command is truncated at a NUL](https://github.com/vim/vim/commit/fb332a2b33ddbbaa172b32148640b9b3c92a8dc0)
- [7.4b.013 : install dir for JP man pages is wrong](https://github.com/vim/vim/commit/48483dbaf78b4e76e1d61c8c2bd14cb533e4a73b)
- [7.4b.014 : (after 7.4b.012) stupid mistake](https://github.com/vim/vim/commit/f40f4ab8df0bb52a2d5f684d961b7c932c30e5df)
- [7.4b.015 : (after 7.4b.008) can't compile without the 'acd' feature](https://github.com/vim/vim/commit/01c458e4383002519b542c95371153eb4b83789e)
- [7.4b.016 : Ruby detection fails on Fedora 19](https://github.com/vim/vim/commit/026a445c4034ce3f489d582c45e2cf49bd77322e)
- [7.4b.017 : ":he \\\^x" gives a strange error message](https://github.com/vim/vim/commit/d82db6046ceab1294c54f93fbe0ea619d87d80be)
- [7.4b.018 : Win32: Dialog can still be too big](https://github.com/vim/vim/commit/a95d8235549a91f1ae3381b82c81d6a75f425642)
- [7.4b.019 : tabline is not updated properly when closing a tab on Win32](https://github.com/vim/vim/commit/86f931e1c3252798dbd60a5d4e07e3c64ae8ee74)
- [7.4b.020 : "g~ap" changes first character of next paragraph](https://github.com/vim/vim/commit/f17968b54a00380d8e955bc13eb0a2ffd30ff8ea)
- [7.4b.021 : multiple press-enter messages when pressing "u"](https://github.com/vim/vim/commit/b09129684b15de58bd2eea8328a8d91896f3e7d9)
- [7.4b.022 : not waiting for a character when the tick count overflows](https://github.com/vim/vim/commit/b7512b79ce62b5dba75eb5768ec5f0d0fddb49ee)
- [7.4 : release version 7.4](https://github.com/vim/vim/commit/3b1db36689ee9b3a49de1362c1678372263a01fe)
- [7.4.001 : 'ic' doesn't work for patterns such as \[a-z\]](https://github.com/vim/vim/commit/1cfad52a03dc391435d32ad49b2ebb9eeb0df841)
- [7.4.002 : pattern with two alternative look-behind matches doesn't match](https://github.com/vim/vim/commit/3f0df06c2860788c6591b03e38b7dcd5bcfd27a6)
- [7.4.003 : memory access error in Ruby syntax highlighting](https://github.com/vim/vim/commit/ebefd997bbc2ec79fd33145fd88126cfd1b87fa3)
- [7.4.004 : when closing a window fails ":bwipe" may hang](https://github.com/vim/vim/commit/c93df6b0752cf1d11b0c880748821f881c1c4d08)
- [7.4.005 : "vaB" while 'virtualedit' is set selects the wrong area](https://github.com/vim/vim/commit/c56c459075260b4e7fb870ab4687b22222daec08)
- [7.4.006 : mkdir("foo/bar/", "p") gives an error message](https://github.com/vim/vim/commit/67652c2935ce5a4313bb071f06b9fa476e10ef78)
- [7.4.007 : creating a preview window on startup messes up the screen](https://github.com/vim/vim/commit/74cd624fdeef12526e1f69b9ddaee037fcef8f3f)
- [7.4.008 : new regexp engine can't be interrupted](https://github.com/vim/vim/commit/41f1205fef7eb24656abb3eeb1dea9ff5cdf1b62)
- [7.4.009 : too easy to write a file was not decrypted (yet)](https://github.com/vim/vim/commit/cf81aefd9c33e2619965ce9854c79105dc7b8bea)
- [7.4.010 : (after 7.4.006) crash with invalid argument to mkdir()](https://github.com/vim/vim/commit/195ef0ccb26b3edfd0f31b30d80108290b2e9b1d)
- [7.4.011 : cannot find out if "acl" and "xpm" features are supported](https://github.com/vim/vim/commit/b5ef5e1c217931023e46ce4a6563d80d259c8fc8)
- [7.4.012 : MS-Windows: resolving multi-bye shortcut does not work](https://github.com/vim/vim/commit/604729e75b14d3a7c5620c43c102e1f4ed156915)
- [7.4.013 : MS-Windows: File name buffer too small for utf-8](https://github.com/vim/vim/commit/d2a203bb58cc546ffb81becf05380c52cb8bae31)
- [7.4.014 : MS-Windows: check for writing to device does not work](https://github.com/vim/vim/commit/4e4f529792c27442b7ad7222cff329907ec389fe)
- [7.4.015 : MS-Windows: Detecting node type fails for multi-byte chars](https://github.com/vim/vim/commit/4dee1bb0dbcfd4c8aa22533f99603dd337b73ab0)
- [7.4.016 : MS-Windows: File name completion has problem with Chinese](https://github.com/vim/vim/commit/65f04f67f15fa5df5b67e038484ba0ea50418931)

## 新着スクリプト

- [better-header : add file header automatically and intelligently](http://www.vim.org/scripts/script.php?script_id=4676)
- [unstack : Open files from stack traces in vsplits](http://www.vim.org/scripts/script.php?script_id=4677)
- [neat.vim : A simple VIM plugin to help you pretty print your data.](http://www.vim.org/scripts/script.php?script_id=4678)
- [Vim-Ultimate-Colorscheme-Utility : Utility for managing multiple favorite colorschemes and easy color browsing.](http://www.vim.org/scripts/script.php?script_id=4679)
- [Visual Studio : A color scheme that tries to imitate Visual Studio 2010.](http://www.vim.org/scripts/script.php?script_id=4680)
- [Pretty Args : Pretty print for :args command](http://www.vim.org/scripts/script.php?script_id=4681)
- [jshint2.vim : Lightweight, customizable and functional Vim plugin for JSHint integration.](http://www.vim.org/scripts/script.php?script_id=4682)
- [playroom : A light colorscheme based on the FlatUI palette](http://www.vim.org/scripts/script.php?script_id=4683)
- [EasyClip : Simplified clipboard functionality for Vim](http://www.vim.org/scripts/script.php?script_id=4684)
- [logpad-plus-duration : Emulate Windows Notepad logging feature, and log duration of each entry](http://www.vim.org/scripts/script.php?script_id=4685)
- [Crunch : An easier way to perform calculations inside Vim](http://www.vim.org/scripts/script.php?script_id=4686)
- [move : Move lines and selected text up and down](http://www.vim.org/scripts/script.php?script_id=4687)
- [simple&#x5f;highlighting : Easily highlight multiple words in all buffers](http://www.vim.org/scripts/script.php?script_id=4688)
- [vim-kitchen : initial version](http://www.vim.org/scripts/script.php?script_id=4689)
- [oops&#x5f;trace.vim : oops&#x5f;trace.vim - Use tag-style lookups on Linux kernel oops backtraces.](http://www.vim.org/scripts/script.php?script_id=4690)
- [movealong.vim : Context-aware motion commands](http://www.vim.org/scripts/script.php?script_id=4691)
- [desertink.vim : Colorscheme based on desert](http://www.vim.org/scripts/script.php?script_id=4692)
- [fixkey : Use non-ASCII keys in console Vim](http://www.vim.org/scripts/script.php?script_id=4693)
- [WSearch : Web Search Plugin](http://www.vim.org/scripts/script.php?script_id=4694)
- [0x7A69&#x5f;dark.vim : 0x7A69](http://www.vim.org/scripts/script.php?script_id=4695)
- [liquidfold.vim : liquidfold.vim provides Vim syntax folding for Liquid templates. It supports any](http://www.vim.org/scripts/script.php?script_id=4696)
- [perl-nextmethod : \]m, \]M, \[m, \[M next/previous perl method motions](http://www.vim.org/scripts/script.php?script_id=4697)
- [nextfile : vim script to switch between test files using \\n ](http://www.vim.org/scripts/script.php?script_id=4698)
- [webkitgtk : syntax for webkitgtk-3.0](http://www.vim.org/scripts/script.php?script_id=4699)
- [libxml2 : syntax for libxml2](http://www.vim.org/scripts/script.php?script_id=4700)
- [mlp.vim : My Little Pony-Inspired Color Scheme](http://www.vim.org/scripts/script.php?script_id=4701)
- [tex-fold : Easy TeX folding](http://www.vim.org/scripts/script.php?script_id=4702)
- [tag-surfer : Fuzzy tag searches](http://www.vim.org/scripts/script.php?script_id=4703)
- [vim-vsnumber : vim script to make the number sort in selection area](http://www.vim.org/scripts/script.php?script_id=4704)
- [AtelierDune : Part of a set of colorschemes that come with a light and dark background ](http://www.vim.org/scripts/script.php?script_id=4705)
- [excess-lines : Highlight surplus characters of long lines](http://www.vim.org/scripts/script.php?script_id=4706)
- [clever-f.vim : Extended f/F/t/T mappings for more convenience.](http://www.vim.org/scripts/script.php?script_id=4707)
- [open-pdf.vim : Convert pdf file to plain text, cache it and open it quickly in vim using pdftot](http://www.vim.org/scripts/script.php?script_id=4708)
- [unite-n3337 : A vim plugin for unite.vim to look in N3337 quickly.](http://www.vim.org/scripts/script.php?script_id=4709)
- [tweetvim-advanced-filter : A high-configurable timeline filter for TweetVim.](http://www.vim.org/scripts/script.php?script_id=4710)

## 月間ダウンロードランキング

1. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (2122)
2. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (2096)
3. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (1462)
4. [pathogen.vim : Easy manipulation of 'runtimepath', 'path', 'tags', etc](http://www.vim.org/scripts/script.php?script_id=2332) (1152)
5. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (954)
6. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (861)
7. [snipMate : TextMate-style snippets for Vim](http://www.vim.org/scripts/script.php?script_id=2540) (857)
8. [github theme : A gvim colorscheme based on github's syntax highlighting](http://www.vim.org/scripts/script.php?script_id=2855) (725)
9. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (721)
10. [bufexplorer.zip : Buffer Explorer / Browser](http://www.vim.org/scripts/script.php?script_id=42) (667)

## vim-jp issues

Open : 175 (+4) | Closed : 290 (+12)

- [Issue #450 : GUI版で画面にノイズが残ることがある (Windows)](https://github.com/vim-jp/issues/issues/450)
- [Issue #451 : :setlocal nowrap と :setlocal wrap nowrap で動作が変わる](https://github.com/vim-jp/issues/issues/451)
- [Issue #452 : 除算で0で割ったときエラーが出ない](https://github.com/vim-jp/issues/issues/452)
- [Issue #453 : enc=utf-8 時に外部コマンドに渡されるマルチバイト文字がutf-8である為、正しく実行されない。](https://github.com/vim-jp/issues/issues/453)
- [Issue #454 : Vim 7.4.14b BETA で GUI タブが表示されない場合がある](https://github.com/vim-jp/issues/issues/454)
- [Issue #455 : enc=utf-8で正しく動作しない処理 (Windows)](https://github.com/vim-jp/issues/issues/455)
- [Issue #456 : filetype=systemverilogの追加について](https://github.com/vim-jp/issues/issues/456)
- [Issue #457 : スペルチェックのポップアップメニューが日本語化されていない](https://github.com/vim-jp/issues/issues/457)
- [Issue #458 : \\ze の後に \\zs がマッチすると内部エラーになる](https://github.com/vim-jp/issues/issues/458)
- [Issue #459 : 新正規表現エンジンが暴走しCTRL-Cでもインタラプトできない](https://github.com/vim-jp/issues/issues/459)
- [Issue #460 : オートロード変数（型は辞書）の中身の関数を、変数未定義の状態でcallするとオートロードしない](https://github.com/vim-jp/issues/issues/460)
- [Issue #461 : autoload関数（オートロードスクリプト）の初回読み込みが遅い](https://github.com/vim-jp/issues/issues/461)
- [Issue #462 : autoload関数のテストを追加する](https://github.com/vim-jp/issues/issues/462)
- [Issue #463 : コマンドラインモードでのレジスタ内容挿入コマンドのバリエーションを増やそうぜ！](https://github.com/vim-jp/issues/issues/463)
- [Issue #464 : FileType イベントでpeditを開くと、閉じるときに落ちる](https://github.com/vim-jp/issues/issues/464)
- [Issue #465 : ユーザー定義コマンドの引数がレジスタ扱いされたsyntaxになっている](https://github.com/vim-jp/issues/issues/465)

