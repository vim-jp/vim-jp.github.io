---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2016 年 4 月号

---

## 話題

- [Software Design: コード編集の高速化からGitHub連携までVim［実戦］投入](http://gihyo.jp/magazine/SD/archive/2016/201605)

    mattn さん、rhysd さん、thinca さん、tyru さん、ujihisa さんによる Vim 特集記事。

- [New version of Vim will be reLeased](http://www.slideshare.net/cohama/2016-0415-nlnagoyavim8)

    [cohama](https://twitter.com/c0hama/status/721248879446917120) さんによる [NLNagoya](http://llnagoya.connpass.com/event/25541/) での発表資料 on 2015/04/15

## 今月の新機能

- 7.4.1703: `assert_notmatch()`、`assert_notequal()` 関数が追加されました。
- 7.4.1707: 辞書のキーに空文字列を使えるようになりました。
- 7.4.1727: `garbagecollect_for_testing()` 関数が追加されました。
- 7.4.1730: `strgetchar()`、`strcharpart()` 関数が追加されました。
- 7.4.1735: `:messages` コマンドに count と clear 引数が追加されました。
- 7.4.1758: `feedkeys()` に "!" フラグが追加されました。
- 7.4.1770: `'guicolors'` オプションが追加されました。
- 7.4.1799: `'guicolors'` は `'termguicolors'` オプションへ名称が変更になりました。

## リリース情報

- [7.4.1690 : can't compile with the conceal feature but without multi-byte](https://github.com/vim/vim/commit/42356150badca33b4e42dc8172abbceff2e86cbe)
- [7.4.1691 : ml&#x5f;get error when going to new buffer and automd sets syntax](https://github.com/vim/vim/commit/b681be175b6991cdc2b8ddd49b0e97e3fe2b201e)
- [7.4.1692 : feedkeys('i', 'x') gets stuck, waits for typing a character](https://github.com/vim/vim/commit/9bd547aca41799605c3a3f83444f6725c2d6eda9)
- [7.4.1693 : building the Perl interface gives compiler warnings](https://github.com/vim/vim/commit/864733ad92e30cd603314604af73f25106db4c90)
- [7.4.1694 : Win32 gvim doesn't work with "dvorakj" input method](https://github.com/vim/vim/commit/f28d87146544e3b5d70aaa6a2513019f6de043ad)
- [7.4.1695 : ":syn reset" clears the effect ":syn iskeyword"](https://github.com/vim/vim/commit/8bc189e81aa98ba4aebb03a9dc9527a210fce816)
- [7.4.1696 : using :stopinsert in silent mapping doesn't clear "INSERT" ](https://github.com/vim/vim/commit/fd773e9e88add7d1ffef890fb9f3a00d613b4326)
- [7.4.1697 : display problems when char width doesn't match terminal](https://github.com/vim/vim/commit/cb0700844c1274fe8bc0ceaffaee0ad21c406f30)
- [7.4.1698 : two tests fail when running tests with MinGW](https://github.com/vim/vim/commit/298c65971e884666d57c32bff6b730d517d9dc30)
- [7.4.1699 : :packadd does not work the same when used early or late](https://github.com/vim/vim/commit/71fb0c146bef08dc276fc5793bd47366e6e0f32a)
- [7.4.1700 : equivalence classes are not properly tested](https://github.com/vim/vim/commit/22e421549d54147d003f6444de007cb1d73f1d27)
- [7.4.1701 : equivalence classes still tested in old style tests](https://github.com/vim/vim/commit/f9f22dbe4f90673ecce601a9dee4bb750ce3cd8f)
- [7.4.1702 : using freed memory when parsing 'printoptions' fails](https://github.com/vim/vim/commit/4afc7c5d4a73340831077a02bfe1f74935e7f4a1)
- [7.4.1703 : can't assert for not equal and not matching](https://github.com/vim/vim/commit/b50e5f56861deb867478997397f7c784a7043233)
- [7.4.1704 : using freed memory with "wincmd p"](https://github.com/vim/vim/commit/3dda7db4e1f7c4a8110a1f83001ec36b46693d27)
- [7.4.1705 : the 'guifont' option does not allow for a quality setting](https://github.com/vim/vim/commit/7c1c6dbb6817640fd3956a0d5417da23fde336d8)
- [7.4.1706 : old style function declaration breaks build](https://github.com/vim/vim/commit/e185c1efba3cb2611ac303c39a08e908497cbac4)
- [7.4.1707 : cannot use empty dictionary key, even though it can be useful](https://github.com/vim/vim/commit/0921ecff1c5a74541bad6c073e8ade32247403d8)
- [7.4.1708 : new regexp engine does not work properly with EBCDIC](https://github.com/vim/vim/commit/2a6fa564a3b5061c14ff63b8b0f12801df0b0ac2)
- [7.4.1709 : mistake in #ifdef](https://github.com/vim/vim/commit/4c9ce053d9f2a94cd704342dd4c25670a5995cbd)
- [7.4.1710 : not all output of an external command is read](https://github.com/vim/vim/commit/17fe5e1aecbeff5ca4b2a821ede9badd5dddae59)
- [7.4.1711 : eror caught with try/catch in 'statusline' handled as error](https://github.com/vim/vim/commit/a742e084b677f76c67e9e52c4f9fb9ab24002e20)
- [7.4.1712 : for plugins in packages need to take care of all dependencies](https://github.com/vim/vim/commit/49b27326447d0827c59c6cd201d58f65c1163086)
- [7.4.1713 : GTK GUI doesn't work on Wayland](https://github.com/vim/vim/commit/f80663f17b2f2499b45eb4467088704c8298c385)
- [7.4.1714 : non-GUI specific settings in the gvimrc&#x5f;example file](https://github.com/vim/vim/commit/54f1b7abf8c48b1dd997202258d1d0673ed4bd29)
- [7.4.1715 : double free when a partial is in a cycle with a list or dict](https://github.com/vim/vim/commit/ddecc25947dbdd689d5bcaed32f298a08abdd497)
- [7.4.1716 : 'autochdir' doesn't work for the first file](https://github.com/vim/vim/commit/baec5c1768098f9dd867b465aaabfdfb294f10c2)
- [7.4.1717 : leaking memory when opening a channel fails](https://github.com/vim/vim/commit/0e4c1de5560c7f8b4cae539ec8cff0949daba3fc)
- [7.4.1718 : Coverity: not using return value of set&#x5f;ref&#x5f;in&#x5f;item()](https://github.com/vim/vim/commit/d56374e25df0b317b01423a01f158157faa647fa)
- [7.4.1719 : leaking memory with a cycle involving a job and a partial](https://github.com/vim/vim/commit/107e1eef1df3b786ad3ad49fbdb9e058649303b5)
- [7.4.1720 : tests fail without the job feature](https://github.com/vim/vim/commit/9e4043757cd2fc18d071fdd98abf297d74878943)
- [7.4.1721 : the vimtbar files are unused.](https://github.com/vim/vim/commit/a4d13de8363abf4c12fe99a52b4e49e763de92fe)
- [7.4.1722 : crash when calling garbagecollect() after starting a job](https://github.com/vim/vim/commit/0239acb11fe4bfe9b525ea90b782759da5eb7704)
- [7.4.1723 : eror caught with try/catch in 'tabline' handled as error ](https://github.com/vim/vim/commit/f73d3bc253fa79ad220f52f04b93e782e95a9d43)
- [7.4.1724 : (after 7.4.1723) tabline test fails in GUI](https://github.com/vim/vim/commit/73cd8fb3e87e4b29dfc489f58e56dee1839c18e5)
- [7.4.1725 : compiler errors for non-ANSI compilers](https://github.com/vim/vim/commit/3780bb923a688e0051a9a23474eeb38a8acb695a)
- [7.4.1726 : ANSI compiler complains about string length](https://github.com/vim/vim/commit/700eefe5a4385fd128f5496e3ca384869752376a)
- [7.4.1727 : cannot detect crash in tests when caused by garbagecollect()](https://github.com/vim/vim/commit/ebf7dfa6f121c82f97d2adca3d45fbaba9ad8f7e)
- [7.4.1728 : the help for functions require a space after the "("](https://github.com/vim/vim/commit/81edd171a9465cf99cede4fa4a7b7bca3d538b0f)
- [7.4.1729 : Perl: cannot use 'print' for writing directly in standard IO](https://github.com/vim/vim/commit/6244a0fc29163ba1c734f92b55a89e01e6cf2a67)
- [7.4.1730 : it is not easy to get a character out of a string](https://github.com/vim/vim/commit/58de0e2dcc1f2d251b74892a06d71a14973f3187)
- [7.4.1731 : Python: turns partial into simple funcref](https://github.com/vim/vim/commit/8110a091bc749d8748a20807a724a3af3ca6d509)
- [7.4.1732 : folds may close when using autocomplete](https://github.com/vim/vim/commit/429fcfbf9a9275367fe9441a50a3dcd773497d84)
- [7.4.1733 : "make install" doesn't know about cross-compiling](https://github.com/vim/vim/commit/839e954aaa72ef62f65416d177f829e681c15466)
- [7.4.1734 : (after 7.4.1730) test fails when not using utf-8](https://github.com/vim/vim/commit/0f518a8f4d4be4cac10389680f6bd5e3781f94b0)
- [7.4.1735 : it is not possible to clear or see part of message history](https://github.com/vim/vim/commit/451f849fd6282a4facd4f0f58af62837443fb5a6)
- [7.4.1736 : (after 7.4.1731) unused variable](https://github.com/vim/vim/commit/baa9fcaf4042a6dbe01e64ce63cb54941ee00f09)
- [7.4.1737 : argument marked as unused is used](https://github.com/vim/vim/commit/52196b2dbe3b64b5054e1df3d3aa8fc65e30addc)
- [7.4.1738 : count for ":messages" depends on number of lines](https://github.com/vim/vim/commit/5d91646599a7b74310845e30a2a52ff197dc0ed7)
- [7.4.1739 : messages test fails on MS-Windows](https://github.com/vim/vim/commit/bea1ede1c59a11ca5bf9d91cd30b7b2937b9fb41)
- [7.4.1740 : syn-cchar defined with matchadd() does not always appear](https://github.com/vim/vim/commit/4d585022023b96f6507e8cae5ed8fc8d926f5140)
- [7.4.1741 : not testing utf-8 characters](https://github.com/vim/vim/commit/b22bd46b9681d73d095f2eadff8163d3a6cf416b)
- [7.4.1742 : strgetchar() does not work correctly](https://github.com/vim/vim/commit/5d18e0eca59ffbba22c7f7c91c9f99d672095728)
- [7.4.1743 : Clang warns for uninitialzed variable](https://github.com/vim/vim/commit/75be2339d877bbd38df91c0181e1e0e388852df6)
- [7.4.1744 : Python: Converting a sequence may leak memory](https://github.com/vim/vim/commit/66210042892389d36e3d37203ec77f61467bfb1c)
- [7.4.1745 : README file is not clear about where to get Vim](https://github.com/vim/vim/commit/5d98c9d93278d6961bfee59151666b8a8bcd23c3)
- [7.4.1746 : memory leak in Perl](https://github.com/vim/vim/commit/95509e18f8806046eeee27482c77666bbec515da)
- [7.4.1747 : Coverity: missing check for NULL pointer](https://github.com/vim/vim/commit/fe4b18640656ddea41f60cf7a76956c9cc5494d6)
- [7.4.1748 : "gD" does not find match in first column of first line](https://github.com/vim/vim/commit/1538fc34fae3fae39773ca43f6ff52401fce61d8)
- [7.4.1749 : when using GTK 3.20 there are a few warnings](https://github.com/vim/vim/commit/30e12d259ee78272359f9da2655d0593a4f6a626)
- [7.4.1750 : messed up screen if buffer updated while in command line mode](https://github.com/vim/vim/commit/7f7c3325d3f1baba32263a3457cfc4d90ecd5ff1)
- [7.4.1751 : crash when 'tagstack' is off](https://github.com/vim/vim/commit/def5abe0a2727041ecee69afdccfca405333bd24)
- [7.4.1752 : when adding to the quickfix list current position is reset](https://github.com/vim/vim/commit/c1808d5822ed9534ef7f0fe509b15bee92a5cc28)
- [7.4.1753 : "noinsert" in 'completeopt' is sometimes ignored](https://github.com/vim/vim/commit/c020042083b9c0a4e932b562c3bef97c76328e18)
- [7.4.1754 : syntax not set when using ":buf" and 'filetype' set manually](https://github.com/vim/vim/commit/c3691332f72169c486066200c0f3935418364900)
- [7.4.1755 : using getreg() on non-existing register NULL list is returned](https://github.com/vim/vim/commit/517ffbee0d5b7b46320726faaa330b61f54e867c)
- [7.4.1756 : "dll" options are not expanded](https://github.com/vim/vim/commit/a6e42501424f6670fa864c739d2dc2eb764900b9)
- [7.4.1757 : 'modified' may be set by complete() when nothing changed](https://github.com/vim/vim/commit/8e42ae5069d4985869e46eaa56900ed19e30f504)
- [7.4.1758 : triggering CursorHoldI when in CTRL-X mode causes problems](https://github.com/vim/vim/commit/245c41070c7f37d52be43cce0cb140bd3ade6c7e)
- [7.4.1759 : using feedkeys() in a timer does not break wait loop](https://github.com/vim/vim/commit/40b1b5443c88fab77f1f7c6f9e801f7ffdb7e0a8)
- [7.4.1760 : (after 7.4.1759) compiler warning for unused variable](https://github.com/vim/vim/commit/4445f7ee708f1a1304526a5979c9dd9883a92a0a)
- [7.4.1761 : Coverity complains about ignoring return value](https://github.com/vim/vim/commit/8ed43916dbea4ccedcc84d271e292abbf658b9c5)
- [7.4.1762 : Coverity: useless assignments](https://github.com/vim/vim/commit/268a06ce901d2c780304e0395028e3c2f60ec755)
- [7.4.1763 : Coverity: useless assignment](https://github.com/vim/vim/commit/ba53435144f46eaaa53c63a62e748b3feee9742c)
- 7.4.1764 : C++ style comment (included 7.4.1765)
- [7.4.1765 : undo options are not together in the options window](https://github.com/vim/vim/commit/4694a17d1ec08382f996990a7fac1ac60197ec81)
- [7.4.1766 : building instructions for MS-Windows are outdated](https://github.com/vim/vim/commit/02cfac85b4e4b038bb2df6962699fa93a42c7eb1)
- [7.4.1767 : when installing on a GTK system the icon cache is not updated](https://github.com/vim/vim/commit/4adfaabfe7e07da9546b45130cad0d266ba48611)
- [7.4.1768 : arguments of setqflist() are not checked properly](https://github.com/vim/vim/commit/d106e5ba7f10f0d2a14eaefe5d78405044416cb9)
- [7.4.1769 : no "closed"/"errors"/"encoding" attribute on Python output](https://github.com/vim/vim/commit/6d4431e7b675ba7a0194c0b8eb84b7d92e4e7953)
- [7.4.1770 : cannot use true color in the terminal](https://github.com/vim/vim/commit/8a633e3427b47286869aa4b96f2bfc1fe65b25cd)
- [7.4.1771 : (after 7.4.1768) warning for unused variable](https://github.com/vim/vim/commit/21decdd3e6d5ce3469a3c2743de1f4aada388d8c)
- [7.4.1772 : (after 7.4.1767) installation fails when env var is empty](https://github.com/vim/vim/commit/54a38415caa31de8b4ac46a9234ae339b1fd6255)
- [7.4.1773 : (after 7.4.1770) compiler warnings](https://github.com/vim/vim/commit/380130f1e18da92a44372728fe044f56db58585b)
- [7.4.1774 : (after 7.4.1770) cterm true color feature has warnings](https://github.com/vim/vim/commit/902647d2dfb42dce8449dfbbc22dab27a528744d)
- [7.4.1775 : the rgb.txt file is not installed](https://github.com/vim/vim/commit/763b684373bf5954445d8d8d99da1e39a4fc5105)
- [7.4.1776 : using wrong buffer length](https://github.com/vim/vim/commit/a1c487eef71d1673e57511453009de9cb4c9af51)
- [7.4.1777 : newly added features can escape the sandbox](https://github.com/vim/vim/commit/3849992b16011e36a5cb5be4b127f843389b96fd)
- [7.4.1778 : with the term truecolor feature, t&#x5f;8f/t&#x5f;8b are not set](https://github.com/vim/vim/commit/b2fa54a84078e2b8dc3c7c7bfbccf6b75c0788d0)
- [7.4.1779 : using negative index in strcharpart()](https://github.com/vim/vim/commit/73dfe917ba6357413aaf98a021c91add5ac6e9bc)
- [7.4.1780 : warnings reported by cppcheck](https://github.com/vim/vim/commit/dc633cf82758f67f656cda7fa8ccc30414ee53f8)
- [7.4.1781 : synIDattr() does not respect 'guicolors'](https://github.com/vim/vim/commit/da5b3dcf06a3af5b398450258be32b0416451a9b)
- [7.4.1782 : strcharpart() has problems with some multi-byte characters](https://github.com/vim/vim/commit/fca66003053f8c0da5161d1fe4b75b3a389934b5)
- [7.4.1783 : old regexp engine doesn't handle character classes correctly](https://github.com/vim/vim/commit/af98a49dd0ef1661b4998f118151fddbf6e4df75)
- [7.4.1784 : the termtruecolor feature is enabled differently](https://github.com/vim/vim/commit/8e9eb3a6a1e589949f1c878c839528d917fdfc4a)
- [7.4.1785 : (after 7.4.1783) regexp test fails on windows](https://github.com/vim/vim/commit/490465bda6ab66f78041709cc02f48a25486a3e5)
- [7.4.1786 : compiled-in colors do not match rgb.txt](https://github.com/vim/vim/commit/d10abe52019d10403eb559ea0a424bbd310b738f)
- [7.4.1787 : channel close callback is invoked before other callbacks](https://github.com/vim/vim/commit/b2658a1ab02cd0ba848164f70c7c464fdc398162)
- [7.4.1788 : NSIS script is missing packages](https://github.com/vim/vim/commit/c7baa43fdb3f5b001ba3e6eb05bf6e199698eeea)
- [7.4.1789 : cannot use ch&#x5f;read() in the close callback](https://github.com/vim/vim/commit/437905c25d4cedfa16d0f87392e4a000d22362b7)
- [7.4.1790 : leading white space in a job command matters](https://github.com/vim/vim/commit/6231cb8b5b208becf088531816027001acc754e5)
- [7.4.1791 : channel could be garbage collected too early](https://github.com/vim/vim/commit/674127e1801fd02ff07dddf0dc3bf0d8cce68997)
- [7.4.1792 : color name decoding is implemented several times](https://github.com/vim/vim/commit/ab3022196ea4f1496e79b8ee85996e31c45d02f1)
- [7.4.1793 : character classes differ between systems, on OS/X test fails](https://github.com/vim/vim/commit/e8aee7dcf9b12becff86e8ce1783a86801c5f9f6)
- [7.4.1794 : (after 7.4.1792) can't build on MS-Windows](https://github.com/vim/vim/commit/c285fe7c3ffdb3ec4eff20a1d1d5accfc80f1a86)
- [7.4.1795 : compiler warning for redefining RGB](https://github.com/vim/vim/commit/5487544fa5d46af9f8f98aa4f5c59074a6e0639a)
- [7.4.1796 : (after 7.4.1795) colors are wrong on MS-Windows](https://github.com/vim/vim/commit/283ee8b3a07b9da18f6c73f35cf465b83f96406a)
- [7.4.1797 : warning from Windows 64 bit compiler](https://github.com/vim/vim/commit/c61348e83fd4eed8b999fe5b7cdfb4caf90bf3fd)
- [7.4.1798 : still compiler warning for unused return value](https://github.com/vim/vim/commit/bb82762907ba024717ad9af3b229c2fa6405cd36)
- [7.4.1799 : 'guicolors' is a confusing option name](https://github.com/vim/vim/commit/61be73bb0f965a895bfb064ea3e55476ac175162)
- [7.4.1800 : unnecessary #ifdef](https://github.com/vim/vim/commit/631225627d448b591e529a5d3e7ac74ef44b8459)
- [7.4.1801 : make uninstall leaves file behind](https://github.com/vim/vim/commit/113ce084569893a3897c8cc4d7d0f301fef7978a)

## 新着スクリプト

- [vim-shebang : Better filetype detection by the she·bang line at file](http://www.vim.org/scripts/script.php?script_id=5362)
- [you-debug-it :  Vim plugin that helps to control debug code](http://www.vim.org/scripts/script.php?script_id=5363)
- [vim-alias : versatile vim-command line aliases](http://www.vim.org/scripts/script.php?script_id=5364)
- [cpp-devel : Vim C++ development; This vim script (adopted from kde-devel-vim).](http://www.vim.org/scripts/script.php?script_id=5365)
- [apthogen : Easily manage your vim's plugins](http://www.vim.org/scripts/script.php?script_id=5366)
- [Outlaw : The wanted outliner!](http://www.vim.org/scripts/script.php?script_id=5367)
- [Vim-scavenger : Clean up multiple blank lines and trailing spaces.](http://www.vim.org/scripts/script.php?script_id=5370)
- [extra-instant-markdown : Instant Markdown previews from VIM!](http://www.vim.org/scripts/script.php?script_id=5371)
- [Google&#x5f;Translate&#x5f;Py : Translate text directly in vim from the google translate page.](http://www.vim.org/scripts/script.php?script_id=5373)
- [droWMark : Post to WordPress using Pandoc Flavoured Markdown](http://www.vim.org/scripts/script.php?script_id=5374)
- [vim-diction : Write better with vim-diction - find frequently misused, bad or wordy diction](http://www.vim.org/scripts/script.php?script_id=5375)
- [Syntax match : Saving syntax match rules to be loaded next time](http://www.vim.org/scripts/script.php?script_id=5376)
- [vim-grimmjow : Theme for Gvim, MacVim and Neovim](http://www.vim.org/scripts/script.php?script_id=5377)
- [Mesonic : A Vim plugin for Meson build system](http://www.vim.org/scripts/script.php?script_id=5378)
- [splat-syntax : Syntax for Splat](http://www.vim.org/scripts/script.php?script_id=5380)
- [Scalpel : Fast within-file word replacement for Vim](http://www.vim.org/scripts/script.php?script_id=5381)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (1623)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1437)
3. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (625)
4. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (566)
5. [winmanager : A windows style IDE for Vim 6.0](http://www.vim.org/scripts/script.php?script_id=95) (528)
6. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (485)
7. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (471)
8. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (466)
9. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (456)
10. [nginx.vim : initial version](http://www.vim.org/scripts/script.php?script_id=1886) (421)

## vim-jp issues

Open : 244 (-3) | Closed : 665 (+18)

- [Issue #895 : MacVim 最新版で :Exp でファイルを開いたとき新規で開かれない](https://github.com/vim-jp/issues/issues/895)
- [Issue #896 : wildmodeでlistとfullを合わせた時はlist内がハイライトされて欲しい。](https://github.com/vim-jp/issues/issues/896)
- [Issue #897 : CTRL-W gfでE325が発生した時、終了・中止した時に新しいタブを作るべきか](https://github.com/vim-jp/issues/issues/897)
- [Issue #898 : matchadd() の conceal がうまく適用されないケースがある](https://github.com/vim-jp/issues/issues/898)
- [Issue #899 : job&#x5f;startで起動したjobが変数に保存されている時にGCが走るとSEGV](https://github.com/vim-jp/issues/issues/899)
- [Issue #900 : 不適切に called&#x5f;emsg を使っているところを did&#x5f;esmg に変える](https://github.com/vim-jp/issues/issues/900)
- [Issue #901 : go言語関数定義移動するとE926が発生します](https://github.com/vim-jp/issues/issues/901)
- [Issue #902 : DLL を指定するオプション (&#x60;pythondll&#x60; 等) が環境変数や tilde を展開しない](https://github.com/vim-jp/issues/issues/902)
- [Issue #903 : 起動後最初の挿入モードで&#x60;complete()&#x60;関数での補完時に&#x60;noinsert&#x60;設定が無視される](https://github.com/vim-jp/issues/issues/903)
- [Issue #904 : Vim 8.0 に lambda をねじ込む](https://github.com/vim-jp/issues/issues/904)
- [Issue #905 : Vim 8.0 に lambda をねじ込む phase 3](https://github.com/vim-jp/issues/issues/905)
- [Issue #906 : 制限モードで job が実行できる](https://github.com/vim-jp/issues/issues/906)
- [Issue #907 : +termtruecolor の際に E254 が発生する](https://github.com/vim-jp/issues/issues/907)
- [Issue #908 : strcharpart()がマルチバイトで正しく動作しない](https://github.com/vim-jp/issues/issues/908)
- [Issue #909 : WindowsのgVimにおいて、グーグルIME使用時のカーソル表示がおかしくなる](https://github.com/vim-jp/issues/issues/909)

