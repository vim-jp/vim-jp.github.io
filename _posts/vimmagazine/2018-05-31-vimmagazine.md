---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2018 年 05 月号

---

## 話題

*   [vim-jp.org](https://vim-jp.org) が HTTPS になりました
*   Vim 8.1 が 05/18(金) に[リリース](https://vim-jp.org/blog/2018/05/18/Vim-8.1-released.html)されました
*   [asakusa.vim #0](https://asakusavim.connpass.com/event/88381/) が 05/24(木) に開催されました
*   ドキュメント翻訳の [8.1 へのキャッチアップタスクフォース](https://github.com/vim-jp/vimdoc-ja-working/issues/207) が開催中です

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

*   8.0.1782: `getqflist()` と `setqflist()` で取り扱う辞書型のキーに `"module"` が追加されました
*   8.0.1783: Windows のコンソールで256色が使えるようになりました
*   8.0.1786: コマンドラインにおいて `<C-R><C-L>` でバッファのカーソルのある行全体の内容を貼り付けられるようになりました
*   8.0.1792: オプション `-?` でもヘルプを表示できるようになりました
*   8.0.1832: 環境変数に対して `:unlet` ができるようになりました
*   8.1.0000: Vim 8.1 としてリリースされました! (関連パッチ: 8.1.0001 netrw.vim が修正されています)
*   8.1.0020: マクロ記録中と実行中のレジスタを知るための関数 `reg_recording()` と `reg_executing()` が追加されました

*   ランタイムの更新
    *   `debcopyright` ファイルタイプ及びシンタックスが追加されました
    *   [chrisbra/vim-xml-ftplugin](https://github.com/chrisbra/vim-xml-ftplugin) がバンドルされました
        *   フォーマッタについての詳細は `:help format-formatexpr` も参照してください

## Vimに関する脆弱性

特筆すべき脆弱性の報告はありませんでした。

## リリース情報

- [8.0.1781 : file names in quickfix window are not shortened](https://github.com/vim/vim/commit/a796d46f29e3cc235cc981696d7ee80faccb5000)
- [8.0.1782 : no simple way to label quickfix entries](https://github.com/vim/vim/commit/d76ce852668635d81778cedacc2d3f021ed4e475)
- [8.0.1783 : cannot use 256 colors in a MS-Windows console](https://github.com/vim/vim/commit/c5cd88554f1e0b2e9ff08d9a0748238dd8340ce1)
- [8.0.1784 : (after 8.0.1782) gvim test gets stuck in dialog](https://github.com/vim/vim/commit/bc7845da935c0707e119812077cecd6cfb5a65e2)
- [8.0.1785 : (after 8.0.1783) missing symbol in Win32 small build](https://github.com/vim/vim/commit/8a938af6ddefab4b4bc751d3f3545e1d95622c8a)
- [8.0.1786 : no test for 'termwinkey'](https://github.com/vim/vim/commit/b2ac14c0b5e23f8ab97c5c784bcd83e13ba8ded3)
- [8.0.1787 : cannot insert the whole cursor line](https://github.com/vim/vim/commit/e2c8d8392684a940cc5608acc73ff47486bd7b92)
- [8.0.1788 : tool to check a color scheme is not installed](https://github.com/vim/vim/commit/dddab90d22d5f8be9c6751505339397ca5742831)
- [8.0.1789 : BufWinEnter does not work well for a terminal window](https://github.com/vim/vim/commit/606cb8b08ed510962fcdc8ef1abcc1fe35fbffef)
- [8.0.1790 : 'winfixwidth' is not always respected by :close](https://github.com/vim/vim/commit/c136af29c0b1939076fbae7d36afd90dce740315)
- [8.0.1791 : using uint8&#x5f;t does not work everywhere](https://github.com/vim/vim/commit/9894e394b230723abefc05a1c9a336b05a63088b)
- [8.0.1792 : MS-Windows users expect -? to work like --help](https://github.com/vim/vim/commit/c3e81694fca484ebabd99f3637bd83b9ee2dd50a)
- [8.0.1793 : no test for "vim -g"](https://github.com/vim/vim/commit/248be5c5de723c4e2715c574fd920b8b1a1dfebb)
- [8.0.1794 : duplicate term options after renaming](https://github.com/vim/vim/commit/b833c1ef7be1ed216a967dd7262473ec97084fa2)
- [8.0.1795 : lose contact with jobs when :gui forks](https://github.com/vim/vim/commit/b0b98d523036d534755bf1cf79d1595e61c3f7ce)
- [8.0.1796 : GUI: click on tab fails when the focus is in a terminal](https://github.com/vim/vim/commit/231a2db1b372fe9645505e711ac53b9ee748c535)
- [8.0.1797 : terminal window is redrawn too often](https://github.com/vim/vim/commit/6eddadff13164b98fe8198153cc656897b2cbcb3)
- [8.0.1798 : MS-Windows: file considered read-only too often](https://github.com/vim/vim/commit/5aa9896b2e3330e32dc42a54731cc44ec904acca)
- [8.0.1799 : no test for :registers command](https://github.com/vim/vim/commit/7ce551f317a0bb92f8c0521e96325301e2d220ca)
- [8.0.1800 : X11: getting color is slow](https://github.com/vim/vim/commit/778df2a3cb8b58b07647952c708439acb0b06d17)
- [8.0.1801 : MS-Windows: redirecting terminal output does not work](https://github.com/vim/vim/commit/f25329cb94e481999e8b08d886cc0f0169e2020c)
- [8.0.1802 : (after 8.0.1802) MS-Windows: terminal test fails](https://github.com/vim/vim/commit/402c83921e4e7ac8e9b0631dbac845f32ec7245f)
- [8.0.1803 : warning for uninitialized variable](https://github.com/vim/vim/commit/9b50f36e405309be97e28b5e19d81ff0902c5224)
- [8.0.1804 : using :normal in terminal window causes problems](https://github.com/vim/vim/commit/655a82a8d32c2cd248a46c0ce365a4d50e3d56bd)
- [8.0.1805 : qf&#x5f;parse&#x5f;line() is too long](https://github.com/vim/vim/commit/18cebf44177542e6658251bacf6152aa9009ca58)
- [8.0.1806 : InsertCharPre causes problems for autocomplete](https://github.com/vim/vim/commit/39de95257714b76ccd845d081cff57830a79b488)
- [8.0.1807 : function to set terminal name is too long](https://github.com/vim/vim/commit/69e056915c4145b7b64c60963797692a5b293561)
- [8.0.1808 : can't build without TGETENT](https://github.com/vim/vim/commit/9289df5ca9546bda0e0046a45ceb66c6bb7a838f)
- [8.0.1809 : various typos](https://github.com/vim/vim/commit/b9464821901623f983528acaed9e4dc2cea7387b)
- [8.0.1810 : buffer of a terminal only updated in Terminal-Normal mode](https://github.com/vim/vim/commit/56bc8e299cdae1b4606897749ef54a0161381a4a)
- [8.0.1811 : no test for winrestcmd()](https://github.com/vim/vim/commit/72cf47a279f7261abf4ae6c9c3ee54024ee87a12)
- [8.0.1812 : the qf&#x5f;jump&#x5f;to&#x5f;usable&#x5f;window() function is too long](https://github.com/vim/vim/commit/7a2b0e55e9460493c4a949bda8be70950dbb8f85)
- [8.0.1813 : Windows installer doesn't install terminal debugger](https://github.com/vim/vim/commit/38623c8512ae3d1bd9556a7b99fc6f7b3a9b1cda)
- [8.0.1814 : crash with terminal window and with 'lazyredraw' set](https://github.com/vim/vim/commit/a10ae5e323f4ebe6620869f4f9b0812b78090915)
- [8.0.1815 : crash with terminal window and with 'lazyredraw' set](https://github.com/vim/vim/commit/0cb8ac71ae42f66d525ad855db01361ca38d935a)
- [8.0.1816 : no test for setcmdpos()](https://github.com/vim/vim/commit/ff3be4fe1e2e723de48b826cb992c798e296c41e)
- [8.0.1817 : a timer may change v:count unexpectedly](https://github.com/vim/vim/commit/b0f42ba60d9e6d101d103421ba0c351811615c15)
- [8.0.1818 : lines remove from wrong buffer when using terminal window](https://github.com/vim/vim/commit/3f1a53c4349a309c3f4acf358d3cf18f9c455050)
- [8.0.1819 : swap file warning for file with non-existing directory](https://github.com/vim/vim/commit/8c3169c58eef3e04f643fe9e045a97b81429e0cb)
- [8.0.1820 : terminal window redirecting stdout does not show stderr](https://github.com/vim/vim/commit/cd8fb449d6486a1a5a09c8c098ea3a38c19e8dc5)
- [8.0.1821 : cursor in terminal window moves when pressing CTRL-W](https://github.com/vim/vim/commit/2bc799579d06936968fa00e52b79ddd4eaca7678)
- [8.0.1822 : make uninstall does not remove colors/tools](https://github.com/vim/vim/commit/ee62f9fa748ea528eb2a06b3e22ef9a5aecb7cca)
- [8.0.1823 : test for terminal stdout redirection is flaky](https://github.com/vim/vim/commit/54c6bafa1609f2d870874d0047b394670602350e)
- [8.0.1824 : Coverity warns for variable that may be uninitialized](https://github.com/vim/vim/commit/73dd1bd54e3d358f156f45e956f6e56527238ac6)
- [8.0.1825 : might use NULL pointer when out of memory](https://github.com/vim/vim/commit/dd1d7be4e6691f447667a8217844a3b2561203be)
- [8.0.1826 : configure uses old compiler flag](https://github.com/vim/vim/commit/52ecaaa63d8d4c6d2ae9072e944a45d8d4e2ac76)
- [8.0.1827 : compiler warning for signed/unsigned char pointers](https://github.com/vim/vim/commit/b2ed680ba760ec5eb2c626363ecbda4dd2528794)
- [8.0.1828 : get no clue why :gui does not fork](https://github.com/vim/vim/commit/e65fffd9b4f02539d8c2717b36a09d0e3250e220)
- [8.0.1829 : MS-Windows: script for vimdiff can't handle ! chars](https://github.com/vim/vim/commit/12365ce49f671a62a1e647df205c9a020828ee1b)
- [8.0.1830 : switching to Terminal-Normal mode does not redraw](https://github.com/vim/vim/commit/05c4a471d235987b914a9cc3ca44b98c46abd157)
- [8.0.1831 : sometimes the quickfix title is incorrectly prefixed with ':'](https://github.com/vim/vim/commit/8b62e31003693fee4b288e7aea49170f032aeef3)
- [8.0.1832 : cannot use :unlet for an environment variable](https://github.com/vim/vim/commit/137374fd6538cf9dee0cb22907728d8fdecb5832)
- [8.0.1833 : X11: ":echo 3.14" gives E806](https://github.com/vim/vim/commit/889fe2c2b9ba5c92b7e862a4aafbd863585fd633)
- [8.0.1834 : GUI: find/replace dialog does not handle some chars](https://github.com/vim/vim/commit/518bc174ed34dc79303488914aaaa3c238a85080)
- [8.0.1835 : print document name does not support multi-byte](https://github.com/vim/vim/commit/2290b1f8aaafbcb38bd801d08e8bf86cb07abfa5)
- [8.0.1836 : buffer-local window options may not be recent](https://github.com/vim/vim/commit/25782a7ff4755daf16c2e1cb5e5f826b13b672ce)
- [8.0.1837 : one character cmdline abbreviation not triggered after '\<,'>](https://github.com/vim/vim/commit/5e3423d192bfa502c6704f731fa2ec6821f9a2f0)
- [8.0.1838 : cursor in wrong pos when switching to Terminal-Normal mode](https://github.com/vim/vim/commit/620020eb1d937e1fd9068ee97d5ef70a99913361)
- [8.0.1839 : script to check .po file doesn't check for plural header](https://github.com/vim/vim/commit/9cfc7d885cb3bf37362b11df19a38992893fb385)
- [8.0.1840 : getwinpos() is not tested](https://github.com/vim/vim/commit/027df2a7d9ef7a8346192b8a00efe6dfb7d7a8cd)
- [8.0.1841 : HP-UX does not have setenv()](https://github.com/vim/vim/commit/1af6a4b891abd30f6bd09b15fb609e5db450392e)
- [8.0.1842 : popup menu inside terminal window isn't cleared](https://github.com/vim/vim/commit/0b565e5da5c05890c2008b18bb5262638bf804f2)
- [8.0.1843 : entry for 'wrap' in options window is wrong](https://github.com/vim/vim/commit/8776889b5befd8eba66f4ad32282db36f85392a6)
- [8.0.1844 : superfluous quickfix code, missing examples](https://github.com/vim/vim/commit/78ddc06bdd5c59cffdbb61eed7dcb5dcc4a17f19)
- [8.0.1845 : various comment updates needed, missing white space](https://github.com/vim/vim/commit/259f26ac2d41ecfb28b82c651b2bfc1edc7c3e29)
- [8.0.1846 : Python interface is incompatible with lldb](https://github.com/vim/vim/commit/d4a8c98ef6b294d605c5f93309509a944dc97a79)
- [8.0.1847 : some build options don't have an example](https://github.com/vim/vim/commit/825ccf4039a1be1a6f27a8c447c57faebaf2dc83)
- [8.0.1848 : 'termwinscroll' does not work properly](https://github.com/vim/vim/commit/4d6cd291cec668b991f2b43d76c6feab8b2e7d98)
- [8.0.1849 : compiler warning for unused arguments, missing prototype](https://github.com/vim/vim/commit/0b39464f69ae8ca9acf6e5bff4435fd0693e22f6)
- [8.0.1850 : todo items in source code not visible for users](https://github.com/vim/vim/commit/03998f606a846e7995047b0356dbb9f515f5244c)

## 新着スクリプト

- [vim-preview : The missing preview window for vim](https://www.vim.org/scripts/script.php?script_id=5678)
- [vim-remembrall : Helps users remember their mappings](https://www.vim.org/scripts/script.php?script_id=5679)
- [yaTranslator :  Translation inside VIM](https://www.vim.org/scripts/script.php?script_id=5680)
- [vim-tmux-pilot : Unified navigation of splits and tabs in vim and tmux](https://www.vim.org/scripts/script.php?script_id=5681)
- [vpager : Pipe input from Vims internal terminal back into Vim](https://www.vim.org/scripts/script.php?script_id=5682)
- [evince-synctex : Python-based vim script for bidirectional synchronization between Vim and Evince](https://www.vim.org/scripts/script.php?script_id=5683)
- [uncrustify : uncrustify formatter support](https://www.vim.org/scripts/script.php?script_id=5684)
- [vipsql : A vim-frontend for interacting with psql ](https://www.vim.org/scripts/script.php?script_id=5685)
- [Exoskeleton : Skeleton - improved!](https://www.vim.org/scripts/script.php?script_id=5686)
- [vimteractive : Sending commands from vim to interactive programs](https://www.vim.org/scripts/script.php?script_id=5687)
- [juvenile.vim : Colorscheme with only grey comments and non-text](https://www.vim.org/scripts/script.php?script_id=5688)
- [nextag.vim : Move to next/previous XML/HTML tag](https://www.vim.org/scripts/script.php?script_id=5689)
- [insert&#x5f;suspend&#x5f;hlsearch.vim : Suspend search highlighting in insert mode](https://www.vim.org/scripts/script.php?script_id=5690)
- [fixed&#x5f;join.vim : Join lines without jumping around](https://www.vim.org/scripts/script.php?script_id=5691)
- [toggle&#x5f;option&#x5f;flags.vim : Toggle flags in option values](https://www.vim.org/scripts/script.php?script_id=5692)
- [copy&#x5f;linebreak.vim : Switch between readable and terminal-copy-friendly wrapping options](https://www.vim.org/scripts/script.php?script_id=5693)
- [sahara.vim : Tweaked desert256 colorscheme](https://www.vim.org/scripts/script.php?script_id=5694)
- [strip&#x5f;trailing&#x5f;whitespace.vim : Strip trailing whitespace exactly the way this author wants](https://www.vim.org/scripts/script.php?script_id=5695)
- [big&#x5f;file&#x5f;options.vim : Automatically set options to speed up working with big files](https://www.vim.org/scripts/script.php?script_id=5696)
- [uncap&#x5f;ex.vim : Correct accidentally capitalised ex commands](https://www.vim.org/scripts/script.php?script_id=5697)
- [auto&#x5f;cache&#x5f;dirs.vim : Create and set cache file directories easily](https://www.vim.org/scripts/script.php?script_id=5698)

## 月間ダウンロードランキング

1. [The NERD tree : A tree explorer plugin for navigating the filesystem](https://www.vim.org/scripts/script.php?script_id=1658) (12)
2. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](https://www.vim.org/scripts/script.php?script_id=273) (10)
3. [verilog&#x5f;systemverilog.vim : Extending Verilog syntax highlighting for SystemVerilog](https://www.vim.org/scripts/script.php?script_id=1586) (8)
4. [Align : Help folks to align text, eqns, declarations, tables, etc](https://www.vim.org/scripts/script.php?script_id=294) (5)
5. [jellybeans.vim : Colorful, dark color scheme](https://www.vim.org/scripts/script.php?script_id=2555) (4)
6. [python.vim : Enhanced version of the python syntax highlighting script](https://www.vim.org/scripts/script.php?script_id=790) (4)
7. [twilight : A clone of Textmates twilight scheme](https://www.vim.org/scripts/script.php?script_id=1677) (4)
8. [indent/python.vim : An alternative indentation script for python](https://www.vim.org/scripts/script.php?script_id=974) (4)
9. [Colortest : xterm 256 color test and visual colors list](https://www.vim.org/scripts/script.php?script_id=1349) (4)
10. [SuperTab : Do all your insert-mode completion with Tab.](https://www.vim.org/scripts/script.php?script_id=1643) (4)

## vim-jp issues

Open : 261 (+0) | Closed : 909 (+5)

- [Issue #1166 : Vimのterminal内のcursesを使ったプログラム内でcursesのKEY&#x5f;RESIZEイベントが取得できない](https://github.com/vim-jp/issues/issues/1166)
- [Issue #1167 : statusline上でTerminal windowの行数を表示するとエラーが出る](https://github.com/vim-jp/issues/issues/1167)
- [Issue #1168 : 'omnifunc' の第一引数 a:findstart の型に一貫性がない](https://github.com/vim-jp/issues/issues/1168)
- [Issue #1169 : :set guioptions+=! のとき v:shell&#x5f;error が反映されない](https://github.com/vim-jp/issues/issues/1169)
- [Issue #1170 : :set guioptions+=! のとき :! コマンドでリダイレクトが機能しない](https://github.com/vim-jp/issues/issues/1170)
