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

- [7.4.765 : (after 7.4.754) CTRL-A/CTRL-X in Visual mode do not work well](https://github.com/vim/vim/commit/9bb1930af908338b68588e988be9601ad144af07)
- [7.4.766 : (after 7.4.757) bg color check does not work on Tera Term](https://github.com/vim/vim/commit/92abe859ed8dc8b9a49b9cc2692d1a2569ebd31e)
- [7.4.767 : --remote-tab-silent can fail on MS-Windows](https://github.com/vim/vim/commit/fafeee6b9e4b1a295001ed2139f9c5bc81eab836)
- [7.4.768 : :diffoff only works properly once](https://github.com/vim/vim/commit/4392996461f56803d5c4a4919ba6dc509434bef5)
- [7.4.769 : (after 7.4 768) behavior of :diffoff is not tested](https://github.com/vim/vim/commit/2693ca21cee8a729d74682fd86a4818f2b050228)
- [7.4.770 : (after 7.4.766) different bg color response is not ignored](https://github.com/vim/vim/commit/46fd4df719b1c3f3d84474c71da8b10e17601ffc)
- [7.4.771 : search wrong when multi-byte character at the start position](https://github.com/vim/vim/commit/5f1e68b7bc241118e5dd8fc781147fdda881ada8)
- [7.4.772 : Racket 6.2 is not supported on MS-Windows](https://github.com/vim/vim/commit/5c5c9801b1e28050a40cef80aeba1ecdea6ea411)
- [7.4.773 : 'langmap' is used in cmdline mode when checking for mappings](https://github.com/vim/vim/commit/ae94b7718ab5ecc820ce487c9ea0178c04a9853c)
- [7.4.774 : it's difficult to get info about the completed item](https://github.com/vim/vim/commit/42a4512f62322fcc12c0071037161d2079d98ba5)
- [7.4.775 : it is not possible to avoid using the first item of completion](https://github.com/vim/vim/commit/b6be1e2390485ff314022d2b7ce0307e42d7e2a6)
- [7.4.776 : equivalence class for 'd' does not work correctly](https://github.com/vim/vim/commit/2c61ec65e8dddcaf9bb5751f6dead97f2a794593)
- [7.4.777 : the README file doesn't look nice on github](https://github.com/vim/vim/commit/694cac7a352ea66827cd94ed1ccbc2fc5655e9cc)
- [7.4.778 : Coverity warns for uninitialized variable](https://github.com/vim/vim/commit/ae2fe73abc954b8fd3dbd7994daa8e31e6690bd0)
- [7.4.779 : cursor moves when using CTRL-A in line without a number](https://github.com/vim/vim/commit/3ec326198029d5a59413b3b8b33dbc9c06c4f28b)
- [7.4.780 : compiler complains about uninitialized and clobbered variables](https://github.com/vim/vim/commit/1db43b1145fe466b5d41af6744a08083983de3a9)
- [7.4.781 : line2byte() returns one less when 'bin' and 'noeol' are set](https://github.com/vim/vim/commit/fe6f18687761403470979feef8fe8ce582ae3b60)
- [7.4.782 : still a few problems with CTRL-A and CTRL-X in Visual mode](https://github.com/vim/vim/commit/5d1bc78a2b9fbe3e3112afcde7c80eb19d5989f4)
- [7.4.783 : copy&#x5f;chars() and copy&#x5f;spaces() are inefficient](https://github.com/vim/vim/commit/2536d4f396680f079e2cb96f9787a434846b5ce7)
- [7.4.784 : using both "noinsert" and "noselect" in 'completeopt' fails](https://github.com/vim/vim/commit/2cf6938706ebcd20e4b4a6723f5f70568b26b1c6)
- [7.4.785 : automatically adding the missing EOL may cause problems](https://github.com/vim/vim/commit/34d72d4b6c1a2b04a214d8a49b7d22c97bc7a8bc)
- [7.4.786 : it is not possible for a plugin to adjust to a changed setting](https://github.com/vim/vim/commit/537443018d41918639695a442c91b34ccec69fc3)
- [7.4.787 : (after 7.4.786) snprintf() isn't available everywhere](https://github.com/vim/vim/commit/fb9bc4829a1442fc8e93f078c9f923c9d382dbd2)
- [7.4.788 : (after 7.4.786) can't build without the crypt feature](https://github.com/vim/vim/commit/5cbb8dbc329ab33ce56ea7ce4695778630c4e9ee)
- [7.4.789 : (after 7.4.788) using freed memory and crash](https://github.com/vim/vim/commit/9cac424d05c0e79cd621f1b9f2f01a9f459fcbe6)
- [7.4.790 : (after 7.4.786) test fails without the autochdir feature](https://github.com/vim/vim/commit/d113a80c77101e9f5723abb075882012753cf611)
- [7.4.791 : the buffer list can be too long to list](https://github.com/vim/vim/commit/d51cb706a4e3ae99555bc214a64c83603c701139)
- [7.4.792 : can only conceal text by defining syntax items](https://github.com/vim/vim/commit/6561d52ecbe6e24d1b90403078cc8b76c53c42fc)
- [7.4.793 : can't specify when not to ring the bell](https://github.com/vim/vim/commit/165bc69d1b7f70ca9d5b657f35d0584ecb7b5183)
- [7.4.794 : Visual Studio 2015 is not recognized](https://github.com/vim/vim/commit/d4a95c3bc5a70f43db0abd020a15d8a1a65db961)
- [7.4.795 : the 'fixeol' option is not copied to a new window](https://github.com/vim/vim/commit/b388be0abf04bba744f390587e8293bfbb1708a4)
- [7.4.796 : warning from 64 bit compiler](https://github.com/vim/vim/commit/33c3a694fca7978cdd28c5e41b498473a08bc05e)
- [7.4.797 : crash when using more command line lines than 'maxcombine'](https://github.com/vim/vim/commit/5f95f288a2d303be1571e818655fd90e399ee58e)
- [7.4.798 : repeating a change in Visual mode does not work as expected](https://github.com/vim/vim/commit/31b259bf9571cae6408be3ef75d9485e24029be5)
- [7.4.799 : accessing memory before an allocated block](https://github.com/vim/vim/commit/6b9be1b6f29884536d99f7347fb9faa08e642e61)
- [7.4.800 : using freed memory when triggering CmdUndefined autocommands](https://github.com/vim/vim/commit/829aef1eb48b17445b1f0a801948c1b826f507f8)
- [7.4.801 : (after 7.4.769) ":diffoff" test could catch more problems](https://github.com/vim/vim/commit/278236000f6449c79041be8d6f4faabb1d295260)
- [7.4.802 : using "A" in Visual mode with 'linebreak' set is not tested](https://github.com/vim/vim/commit/52e21cf057e313a867b0f529c8e4d52ad5ae740d)
- [7.4.803 : C indent does not support C11 raw strings](https://github.com/vim/vim/commit/f7bb86dc593913d055e4cce16cec43f6271adda3)

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

