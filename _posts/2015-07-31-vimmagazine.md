---
layout: post
category: vimmagazine
title: Vim Magazine 7 月号
---

## 話題

- [Osaka.vim #5 at Kyoto](http://osaka-vim.connpass.com/event/18194/)

  そうだ 京都、行こう。 8月9日(土) 京都

- [TokyoVim#25](http://tokyovim.connpass.com/event/18101/)

  Vimに関する何かをやる会です。 8月22日(土) 東京

- [Yokohama.vim.reboot #6](https://atnd.org/events/68683)

  またまた約1年ぶりの開催です。 8月29日(土) 横浜

## 今月の新機能

- 7.4.774 : CompleteDone イベント発生時に使える v:completed\_items 変数が追加されました。
- 7.4.775 : 'completeopt' オプションに "noinsert", "noselect" が追加されました。
- 7.4.785 : 'fixendofline' オプションが追加されました。
- 7.4.786 : OptionSet イベントが追加されました。
- 7.4.791 : :ls コマンドにバッファの種類を指定する引数が追加されました。
- 7.4.792 : matchadd() 関数で conceal の定義ができるようになりました。
- 7.4.793 : 'belloff' オプションが追加されました。

## リリース情報

- [7.4.765 : (after 7.4.754) CTRL-A/CTRL-X in Visual mode do not work well](http://code.google.com/p/vim/source/detail?r=f9876721bedc4f794d425bd015e4b0a65cb4d703)
- [7.4.766 : (after 7.4.757) bg color check does not work on Tera Term](http://code.google.com/p/vim/source/detail?r=31a14cb3a8e35da6f2bfeca5ced12e6420895a17)
- [7.4.767 : --remote-tab-silent can fail on MS-Windows](http://code.google.com/p/vim/source/detail?r=63ac0cdf3b2822cc114dd201c4b5ef59d6395e86)
- [7.4.768 : :diffoff only works properly once](http://code.google.com/p/vim/source/detail?r=d3a23af4bba127e0cc240bb967484781e646f412)
- [7.4.769 : (after 7.4 768) behavior of :diffoff is not tested](http://code.google.com/p/vim/source/detail?r=cfd9a2befd67b56c5a618ee21affa88ed05c34ca)
- [7.4.770 : (after 7.4.766) different bg color response is not ignored](http://code.google.com/p/vim/source/detail?r=2ce6b77a49de5b997f4b2264b3582983058438c9)
- [7.4.771 : search wrong when multi-byte character at the start position](http://code.google.com/p/vim/source/detail?r=dd923806ae3bebc810b87a0dd38e48aa30a50da1)
- [7.4.772 : Racket 6.2 is not supported on MS-Windows](http://code.google.com/p/vim/source/detail?r=d973e23c3bf81e4126c6ff3450d2847b1c4a89ee)
- [7.4.773 : 'langmap' is used in cmdline mode when checking for mappings](http://code.google.com/p/vim/source/detail?r=5b1eefbf9a532f32a66fa13abbd671488aaafd5c)
- [7.4.774 : it's difficult to get info about the completed item](http://code.google.com/p/vim/source/detail?r=676906c337681c90de19aa2ed71bffff5893f9c9)
- [7.4.775 : it is not possible to avoid using the first item of completion](http://code.google.com/p/vim/source/detail?r=e09a7ff8804653c1b784a41fc38e185e5f8ec1f7)
- [7.4.776 : equivalence class for 'd' does not work correctly](http://code.google.com/p/vim/source/detail?r=94b17958585153641ad1e813fa144cad57c7170f)
- [7.4.777 : the README file doesn't look nice on github](http://code.google.com/p/vim/source/detail?r=228dc3f3f7114887bf5c870c75446681fa3a2bee)
- [7.4.778 : Coverity warns for uninitialized variable](http://code.google.com/p/vim/source/detail?r=dd094497f9fa237f3b1d7e673146d8128e3ee2d0)
- [7.4.779 : cursor moves when using CTRL-A in line without a number](http://code.google.com/p/vim/source/detail?r=1f78058351a6fcf079783b1f28b5b77d7d98e4d8)
- [7.4.780 : compiler complains about uninitialized and clobbered variables](http://code.google.com/p/vim/source/detail?r=9a13e3910e3d45c768a1e3d87a15a62b4cee439c)
- [7.4.781 : line2byte() returns one less when 'bin' and 'noeol' are set](http://code.google.com/p/vim/source/detail?r=30c9301f33ed76d0b76b9e22362f5e78e26786da)
- [7.4.782 : still a few problems with CTRL-A and CTRL-X in Visual mode](http://code.google.com/p/vim/source/detail?r=58d9f967ae1a4dc16be4ca92f921976fbc3225ed)
- [7.4.783 : copy&#x5f;chars() and copy&#x5f;spaces() are inefficient](http://code.google.com/p/vim/source/detail?r=e55929fca0cf6078c2c9791a7fcbd3a8b96884d4)
- [7.4.784 : using both "noinsert" and "noselect" in 'completeopt' fails](http://code.google.com/p/vim/source/detail?r=f587f685701094fc9a90ec6b8b946b4a6abc013f)
- [7.4.785 : automatically adding the missing EOL may cause problems](http://code.google.com/p/vim/source/detail?r=62ba356c2d4e5836c86908d2ec65e53a6e335f27)
- [7.4.786 : it is not possible for a plugin to adjust to a changed setting](http://code.google.com/p/vim/source/detail?r=4db70c94226b522a0beb6d54ae0a992bd47405ae)
- [7.4.787 : (after 7.4.786) snprintf() isn't available everywhere](http://code.google.com/p/vim/source/detail?r=c31bd9ca75d237eee7a3cfe92e464fc58b2d534b)
- [7.4.788 : (after 7.4.786) can't build without the crypt feature](http://code.google.com/p/vim/source/detail?r=ca0990adc2e3b10fe06a066144976adc8a769576)
- [7.4.789 : (after 7.4.788) using freed memory and crash](http://code.google.com/p/vim/source/detail?r=b63f09edb493ab555e0ed78f1b045e365fd1f9c8)
- [7.4.790 : (after 7.4.786) test fails without the autochdir feature](http://code.google.com/p/vim/source/detail?r=d37e168dee2030187d77492b0a88c2467c87bf26)
- [7.4.791 : the buffer list can be too long to list](http://code.google.com/p/vim/source/detail?r=957a3bb312fa105b9476ae029742eeb4b5f9b490)
- [7.4.792 : can only conceal text by defining syntax items](http://code.google.com/p/vim/source/detail?r=1efa7c2b9368b7d550a382dbe8b5502bb14416d1)
- [7.4.793 : can't specify when not to ring the bell](http://code.google.com/p/vim/source/detail?r=1e621b31948bed78579a3312a907be76263f1204)
- [7.4.794 : Visual Studio 2015 is not recognized](http://code.google.com/p/vim/source/detail?r=98d1d53edc668ec036ea39d9ac6392263d0268f1)
- [7.4.795 : the 'fixeol' option is not copied to a new window](http://code.google.com/p/vim/source/detail?r=a958ac497a8103f18065b7ae04c0a98c4c82edc6)
- [7.4.796 : warning from 64 bit compiler](http://code.google.com/p/vim/source/detail?r=d19e79d29bf99fe0483339e2793c7b81b3e55ed5)
- [7.4.797 : crash when using more command line lines than 'maxcombine'](http://code.google.com/p/vim/source/detail?r=87cb71d54e8ae31700862bea78ceb89b50741161)
- [7.4.798 : repeating a change in Visual mode does not work as expected](http://code.google.com/p/vim/source/detail?r=c2b132add309bd2fca3b19432a2ceb0fcf6d237b)
- [7.4.799 : accessing memory before an allocated block](http://code.google.com/p/vim/source/detail?r=4bf9538ff25e208503d313ef99e661124394fbd2)
- [7.4.800 : using freed memory when triggering CmdUndefined autocommands](http://code.google.com/p/vim/source/detail?r=86d43639de31ba99947cf9bfe0a3fd171ca7019e)
- [7.4.801 : (after 7.4.769) ":diffoff" test could catch more problems](http://code.google.com/p/vim/source/detail?r=20ed78fab428bee80800fe3492f5f589571352f6)
- [7.4.802 : using "A" in Visual mode with 'linebreak' set is not tested](http://code.google.com/p/vim/source/detail?r=417201526af3e5cca3f3b4a4bf3aac736521c803)
- [7.4.803 : C indent does not support C11 raw strings](http://code.google.com/p/vim/source/detail?r=e859731ea1cd074cc1e2d12c6047978cd6c1947c)

## 新着スクリプト

- [complete-email : Allow completion of email addresses](http://www.vim.org/scripts/script.php?script_id=5208)
- [confirm-quit : Ask for confirmation before quitting Vim](http://www.vim.org/scripts/script.php?script_id=5209)
- [autoSolarize : Automatically switches between light and dark background for solarized](http://www.vim.org/scripts/script.php?script_id=5210)
- [vim-qf : Tame the quickfix window.](http://www.vim.org/scripts/script.php?script_id=5211)
- [vim-on-write : Easily register commands executed when the current buffer or a given files match](http://www.vim.org/scripts/script.php?script_id=5212)
- [vim-AHKcomplete : Omnicomplete for autohotkey script.](http://www.vim.org/scripts/script.php?script_id=5213)
- [vim-qlist : Persist the results of :ilist and related commands via the quickfix list.](http://www.vim.org/scripts/script.php?script_id=5214)
- [Loupe : Enhanced in-file search for Vim](http://www.vim.org/scripts/script.php?script_id=5215)
- [Terminus : Enhanced terminal integration for Vim](http://www.vim.org/scripts/script.php?script_id=5216)
- [vim-clipper : Clipper integration for Vim](http://www.vim.org/scripts/script.php?script_id=5217)
- [vim-gurl : Grabs link to Github / Gitlab remote repository for current file & line range](http://www.vim.org/scripts/script.php?script_id=5218)
- [airline-colornum : Cursor line number color = Airline mode color](http://www.vim.org/scripts/script.php?script_id=5219)
- [Ferret : Enhanced multi-file search for Vim](http://www.vim.org/scripts/script.php?script_id=5220)
- [ctrlp-locate : locate and open it via ctrlp.vim!](http://www.vim.org/scripts/script.php?script_id=5221)
- [vim-js-pretty-template : highlights JavaScript's Template Strings in other FileType syntax rule](http://www.vim.org/scripts/script.php?script_id=5222)
- [vim3diff : 3 way diff](http://www.vim.org/scripts/script.php?script_id=5223)
- [progress.vim : Update - taking over as maintainer](http://www.vim.org/scripts/script.php?script_id=5224)
- [pipe-mysql.vim : MySQL Client for Vim (easy-to-use / hassle-free configuration)](http://www.vim.org/scripts/script.php?script_id=5225)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (1480)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1272)
3. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (745)
4. [AutomaticLaTeXPlugin : Background compilation, completion, bib serch, toc and other nice features.](http://www.vim.org/scripts/script.php?script_id=2945) (707)
5. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (572)
6. [winmanager : A windows style IDE for Vim 6.0](http://www.vim.org/scripts/script.php?script_id=95) (547)
7. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (542)
8. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (518)
9. [SuperTab : Do all your insert-mode completion with Tab.](http://www.vim.org/scripts/script.php?script_id=1643) (438)
10. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (436)

## vim-jp issues

Open : 234 (+4) | Closed : 517 (+3)

- [Issue #745 : Visual mode で \<C-x> や \<C-a> を押すと落ちる場合がある (Patch 7.4.765)](https://github.com/vim-jp/issues/issues/745)
- [Issue #746 : gui で :sh 時にエスケープシーケンスを除去出来ないか](https://github.com/vim-jp/issues/issues/746)
- [Issue #747 : 特定の引数の組み合わせで search()/searchpos() 関数がカーソル直下の Unicode文字を検索できない](https://github.com/vim-jp/issues/issues/747)
- [Issue #748 : UTF-8 以外で DirectWrite を有効にできないか](https://github.com/vim-jp/issues/issues/748)
- [Issue #749 : utf-8時に外部コマンドの出力が文字化けする](https://github.com/vim-jp/issues/issues/749)
- [Issue #750 : fixendofline がバッファローカル変数に引き継がれない](https://github.com/vim-jp/issues/issues/750)
- [Issue #751 : Windows10 の vim.exe で IME エコーバックが機能しない](https://github.com/vim-jp/issues/issues/751)

