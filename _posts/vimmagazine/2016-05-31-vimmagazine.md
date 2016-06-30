---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2016 年 5 月号

---

## 話題

* [Software Design 2016年6月号 Vimの細道【8】ファイル操作を柔軟にするCtrlP](http://gihyo.jp/magazine/SD/archive/2016/201606)
* [Uganda 2016](https://goo.gl/photos/ukDzUAtNWQXDxfzY7) : Bram によるウガンダ支援の近況の写真です

## 今月の新機能

* 7.4.1804: `MANPAGER` として使うための manpager.vim プラグインが追加されました。
* 7.4.1833: `'keywordprg'` に単体のExコマンドを指定できるようになりました。
* 7.4.1839: `get()` を partial から関数、辞書、引数をそれぞれ取り出せるようになりました。
* 7.4.1842: `get()` を funcref に適用できるようなりました。

## リリース情報

- [7.4.1802 : quickfix doesn't handle long lines well, they are split](https://github.com/vim/vim/commit/6be8c8e165204b8aa4eeb8a52be87a58d8b41b9e)
- [7.4.1803 : GTK3 doesn't handle menu separators properly](https://github.com/vim/vim/commit/0b6cf69c038b9af198542edc349ebe8e53a8f847)
- [7.4.1804 : can't use Vim as MANPAGER](https://github.com/vim/vim/commit/b20545f2a718d4f19c3f609fd11c0ca4eff450ce)
- [7.4.1805 : running tests in shadow dir fails](https://github.com/vim/vim/commit/4cc39a527fecc96ad6639f10c9389c66af828cf1)
- [7.4.1806 : 'termguicolors' option missing from the options window](https://github.com/vim/vim/commit/8e3d1b6326c103cc92f8d07b1161ee5172acf201)
- [7.4.1807 : Test&#x5f;out&#x5f;close&#x5f;cb sometimes fails](https://github.com/vim/vim/commit/d75263c020a5044fa73bc8cd57ce73ca5e23b360)
- [7.4.1808 : (after 7.4.1806) using wrong feature name](https://github.com/vim/vim/commit/8a24b794b89916c8074892e7b25121a21f1fa9c9)
- [7.4.1809 : (after 7.4.1808) using wrong short option name](https://github.com/vim/vim/commit/868cfc19bb079a16ca58884b551486566f35419b)
- [7.4.1810 : sending DETACH after a channel was closed isn't useful](https://github.com/vim/vim/commit/715d285d79a1d97928dd717278e25e6de45d1496)
- [7.4.1811 : netbeans channel gets garbage collected](https://github.com/vim/vim/commit/3266c85a44a637862b0ed6e531680c6ab2897ab5)
- [7.4.1812 : failure on startup with Athena and Motif](https://github.com/vim/vim/commit/187147aedd588070d0676664d5076d046644094e)
- [7.4.1813 : memory access error when running test&#x5f;quickfix](https://github.com/vim/vim/commit/9b4ebc692d77ca8ef90d72517347f74c2474dd3d)
- [7.4.1814 : channel may be garbage collected while being used by a job](https://github.com/vim/vim/commit/b8d4905592fc26fcd09180d7d6bfefd899f2f6c6)
- [7.4.1815 : compiler warnings for unused variables](https://github.com/vim/vim/commit/9a3b3311d26c990208150255ad65472bb4eefaa4)
- [7.4.1816 : looping over a null list throws an error](https://github.com/vim/vim/commit/d8585eded6359f1d7e1981e96ae775efd077c638)
- [7.4.1817 : screen not updated if callback invoked when closing channel](https://github.com/vim/vim/commit/cefe4f994853c2d4866e2aa4ea3e3f36ab2fea13)
- [7.4.1818 : help completion adds @en to all matches except the first one](https://github.com/vim/vim/commit/89c79b99328b66f77f1d12dc8c6701dfe2c57f15)
- [7.4.1819 : compiler warnings when sprintf() is a macro](https://github.com/vim/vim/commit/827b165b2aebad2cfe98cc6d5804c6c0fe8afd89)
- [7.4.1820 : removing language from help tags too often](https://github.com/vim/vim/commit/9ccaae04c6f263e6db14fc403bca2404a7871114)
- [7.4.1821 : (after 7.4.1820) test fails on MS-Windows](https://github.com/vim/vim/commit/4ed6b2e2d7fd5959fb9b9f608935d47305c4bbe4)
- [7.4.1822 : redirecting stdout of a channel to "null" doesn't work](https://github.com/vim/vim/commit/ea83bf06b92baeb6d68a10d8e8ffad289d31dae2)
- [7.4.1823 : warning from 64 bit compiler](https://github.com/vim/vim/commit/d9db8b448c214eb583e84c598bca0688b9202ba7)
- [7.4.1824 : process remains defunc if job is not referenced](https://github.com/vim/vim/commit/36e0f7da9bc4a6ee8a7b17df503542a339e034c8)
- [7.4.1825 : when job writes to buffer nothing is written](https://github.com/vim/vim/commit/5d96e3ae534ade8ed09a5de9ff8fd7519537ec28)
- [7.4.1826 : callbacks are invoked when it's not safe](https://github.com/vim/vim/commit/cf7ff70ca73218d618e7c00ab785bcf5f9120a94)
- [7.4.1827 : no error when invoking a callback when it's not safe](https://github.com/vim/vim/commit/fb6ffc732e65dbc459c89247ff78134402f1a18b)
- [7.4.1828 : may try to access buffer that's already freed](https://github.com/vim/vim/commit/e0f76d00979c972329f6c371463a20da61ccad65)
- [7.4.1829 : (after 7.4.1828) no channel log message when buffer was freed](https://github.com/vim/vim/commit/de7eb0a47b557eb4656c6b63d421c7e7bae1ef30)
- [7.4.1830 : non-antialiased misnamed](https://github.com/vim/vim/commit/73a733e08bb7853d2ac12c60756ae51e39abb4d9)
- [7.4.1831 : no proper error message for wrong argument to timer&#x5f;stop()](https://github.com/vim/vim/commit/e40d75f6ef73095286b66e546c70ff6f045e009d)
- [7.4.1832 : memory leak in debug commands](https://github.com/vim/vim/commit/dc303bce10c60a3314078ea168064552fadf01af)
- [7.4.1833 : cannot use an Ex command for 'keywordprg'](https://github.com/vim/vim/commit/2ff8b64679242e73248774a388d54931c9ce49bd)
- [7.4.1834 : possible crash when conceal is active](https://github.com/vim/vim/commit/908be438794619f10ef7fa25e24e5893b2ae0189)
- [7.4.1835 : when splitting and closing a window the status height changes](https://github.com/vim/vim/commit/991dea3ab185fb35e577ab0bdfd443cd4b43ccc6)
- [7.4.1836 : when using partial on dict it's always bound to that dict](https://github.com/vim/vim/commit/1d429610bf9e99a6252be8abbc910d6667e4d1da)
- [7.4.1837 : the BufUnload event may be triggered twice](https://github.com/vim/vim/commit/c67e89213476b5f4756d92208b57ce9ef4a4cf24)
- [7.4.1838 : functions specifically for testing do not sort together](https://github.com/vim/vim/commit/574860b5ee9da281c875dad07a607454e135eaee)
- [7.4.1839 : cannot get the items stored in a partial](https://github.com/vim/vim/commit/2bbf8eff6fab16d86e7bcfc0da1962d31bec7891)
- [7.4.1840 : when using packages an "after" directory cannot be used](https://github.com/vim/vim/commit/a57024453115592b8847af40ddd965a33898e390)
- [7.4.1841 : code to reallocate the buffer used for quickfix is repeated](https://github.com/vim/vim/commit/2b2b8ae5ab37b04584633c469265d85825166905)
- [7.4.1842 : (after 7.4.1839) get() works for Partial but not for Funcref](https://github.com/vim/vim/commit/03e19a04ac2ca55643663b97b6ab94043233dcbd)
- [7.4.1843 : tests involving Python are flaky](https://github.com/vim/vim/commit/2177f9fe18a927ef65ccebb0856722a28dc00252)
- [7.4.1844 : more functions should start with test&#x5f;](https://github.com/vim/vim/commit/8e8df251bf2505e5decf258397c6069fbe5e2e01)
- [7.4.1845 : mentioning NetBeans when reading from channel](https://github.com/vim/vim/commit/f8df45d84f1d17dced23e185d4835fc23406d73d)
- [7.4.1846 : ubsan detects a multiplication overflow](https://github.com/vim/vim/commit/54c10ccf9274880e83093a99690e7bfa9a2d2fa8)
- [7.4.1847 : using NULL dict or list crashes Vim](https://github.com/vim/vim/commit/13ddc5c35921efa69e980284117b6db6465e019c)
- [7.4.1848 : can't build with Strawberry Perl 5.24](https://github.com/vim/vim/commit/6727bf861776cfbb93c97dfea5f87a095cf9f364)
- [7.4.1849 : still trying to read from channel that is going to be closed](https://github.com/vim/vim/commit/5850a764eae74a4dae7238e4e76b4c24f062699b)
- [7.4.1850 : GUI freezes when using a job](https://github.com/vim/vim/commit/bf981eeb6b4ee63ae8543a7f9865ab700159a79c)
- [7.4.1851 : test&#x5f;syn&#x5f;attr fails when using the GUI](https://github.com/vim/vim/commit/180fc2d41812c49b60224a1ca89945a002a090f5)
- [7.4.1852 : Unix: Cannot run all tests with the GUI](https://github.com/vim/vim/commit/af6c12c27bcb553b2fb2c8a9dcfde626fb3670fe)
- [7.4.1853 : crash when job and channel in same dict while using partials](https://github.com/vim/vim/commit/28ae5773422c2cf61aaf8d9d2b9fae70642d6a33)
- [7.4.1854 : when setting 'termguicolors' Ignore highlighting doesn't work](https://github.com/vim/vim/commit/d80629cef03cd40b0bf06c402dfe0b720b3bf608)
- [7.4.1855 : valgrind reports memory leak for job that is not freed](https://github.com/vim/vim/commit/655da31a18ef3f888acf10e68b438e2a851f7b14)
- [7.4.1856 : failing job test fails on MS-Windows](https://github.com/vim/vim/commit/324a78f3b649e7b14741519ecf19c4aba178772d)
- [7.4.1857 : channel can append to a buffer that is 'nomodifiable'](https://github.com/vim/vim/commit/9f5842e63fc63d438cbffcec503e072a06f74dc2)
- [7.4.1858 : channel writing to buffer doesn't find it by the short name](https://github.com/vim/vim/commit/b127cfd75f59e82580df395b6e2c009774644b16)
- [7.4.1859 : cannot use a function reference for "exit&#x5f;cb"](https://github.com/vim/vim/commit/ef3abc6442260e9a0314970a532400b05571d3fe)
- [7.4.1860 : using a partial for timer&#x5f;start() may cause a crash](https://github.com/vim/vim/commit/e3188e261569ae512fb1ae2653b57fdd9e259ca3)
- [7.4.1861 : compiler warnings with 64 bit compiler](https://github.com/vim/vim/commit/b055066a1daf12c349d6c575aff22ae4d999a157)
- [7.4.1862 : eval() can't handle result from string() with repeated item](https://github.com/vim/vim/commit/18dfb4404a618c52ee7138630a2381aed4d66eaf)

## 新着スクリプト

- [ensime-vim : Ensime support for Scala and Java files](http://www.vim.org/scripts/script.php?script_id=5382)
- [caw.vim : Comment plugin: Operator mappings/Dot-repeatable/300+ filetypes](http://www.vim.org/scripts/script.php?script_id=5383)
- [vim-stackoverflow : A tiny plugin for searching Stackoverflow from VIM editor ](http://www.vim.org/scripts/script.php?script_id=5384)
- [vim-freemarker : Freemarker syntax and functions](http://www.vim.org/scripts/script.php?script_id=5385)
- [vim-logreview : plugin to help log reviews and navigation ](http://www.vim.org/scripts/script.php?script_id=5386)
- [definition.vim : Jump to the last place where \[command\|function\|mapping\] was set](http://www.vim.org/scripts/script.php?script_id=5387)
- [Solarized 8 : Optimized Solarized colorschemes. Best served with true-color terminals!](http://www.vim.org/scripts/script.php?script_id=5388)
- [RougherExp : Changes your search expression into rougher regexp.](http://www.vim.org/scripts/script.php?script_id=5389)
- [tortus.vim : Dark (grey on black) color scheme based on on torte.vim and motus.vim.](http://www.vim.org/scripts/script.php?script_id=5390)
- [easy-execute : Execute files easy in unix](http://www.vim.org/scripts/script.php?script_id=5391)
- [SIDECAR Assembly Syntax Highlighting Fil : SIDECAR Assembly Syntax Highlighting File](http://www.vim.org/scripts/script.php?script_id=5392)
- [vim eslint : eslint](http://www.vim.org/scripts/script.php?script_id=5393)
- [twig.vim : Twig syntax highlighting, indentation and neocomplete snippets](http://www.vim.org/scripts/script.php?script_id=5394)
- [tabulous : Lightweight Vim plugin for setting the tabline including the tab page labels.](http://www.vim.org/scripts/script.php?script_id=5395)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (1665)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1426)
3. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (681)
4. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (677)
5. [winmanager : A windows style IDE for Vim 6.0](http://www.vim.org/scripts/script.php?script_id=95) (552)
6. [bufexplorer.zip : Buffer Explorer / Browser](http://www.vim.org/scripts/script.php?script_id=42) (545)
7. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (485)
8. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (477)
9. [a.vim : Alternate Files quickly (.c --> .h etc)](http://www.vim.org/scripts/script.php?script_id=31) (471)
10. [SuperTab : Do all your insert-mode completion with Tab.](http://www.vim.org/scripts/script.php?script_id=1643) (466)

## vim-jp issues

Open : 245 (+1) | Closed : 669 (+4)

- [Issue #910 : tabonlyで1001個以降のタブが閉じられない](https://github.com/vim-jp/issues/issues/910)
- [Issue #911 : string(\[a,a,a\])が正しくない](https://github.com/vim-jp/issues/issues/911)
- [Issue #912 : bufhiddenを設定するとBufUnloadが2回実行される](https://github.com/vim-jp/issues/issues/912)
- [Issue #913 : ウィンドウ関連のコマンドを実行するとコマンドラインモードの高さが上に伸びる](https://github.com/vim-jp/issues/issues/913)
- [Issue #914 : 挿入モードに入っただけでTextChangedIが発火する](https://github.com/vim-jp/issues/issues/914)

