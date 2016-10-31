---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2016 年 9 月号

---

## 話題

*   Vim 8.0 がリリースされました
*   VimConf 2016 のスピーカー募集が終了しました

## 今月の新機能

*   7.4.2298: `ch_close_in()` が追加されました。
*   7.4.2299: `:cexpr`, `:cbuffer` などにおいても `QuickFixCmdPre`, `QuickFixCmdPost` イベントがトリガーされるようになりました。
*   7.4.2302: Windows 向け公式インストーラーの対応する Perl と Python のバージョンが、それぞれ 5.24 と 3.5 になりました。
*   7.4.2304: timer callback 内で自身の timer を停止できるようになりました。
*   7.4.2306: `'langremap'` のデフォルト値がドキュメント通りの `enabled` になりました。
*   7.4.2311: Appveyor の 64 bit build の Python3 interface が Python 3.5 を使用するようになりました。
*   7.4.2314: 現在の augroup を `augroup!` で削除できなくなりました。(E936)
*   7.4.2319: $VIM/vimrc で `g:skip_defaults_vim` を定義すると `defaults.vim` の読み込みをスキップできるようになりました。
*   7.4.2340: Windows 向け公式インストーラーの対応 Ruby バージョンが 2.2 になりました。
*   8.0.0015: `ch_status()` に第 2 引数 `{options}` が追加されました。

## リリース情報

- [7.4.2296 : no tests for :undolist and "U" command](https://github.com/vim/vim/commit/c628fdcd46e93c308f742efdf54248695960e290)
- [7.4.2297 : job that reads from a buffer when reaching the end it hangs](https://github.com/vim/vim/commit/d8b554904d18fe19bd9fa79dbda880845cb017d2)
- [7.4.2298 : it is not possible to close the "in" part of a channel](https://github.com/vim/vim/commit/0874a83e9be1b39fdb217f02b427bf1d6133a4d8)
- [7.4.2299 : quickfix autocommands are not always triggered](https://github.com/vim/vim/commit/04c4ce650f9e533cd35b2aa6803f4d354d3ec7aa)
- [7.4.2300 : warning for deleting autocmd group when autocmd just deleted](https://github.com/vim/vim/commit/5c80908ced601be6db7554a147cdb0f98ac8daa1)
- [7.4.2301 : MS-Windows: some files remain after testing](https://github.com/vim/vim/commit/641ad6c7ac7367f95fd927b8efa4bf74ddb9ccf3)
- [7.4.2302 : default interface versions for MS-Windows are outdated](https://github.com/vim/vim/commit/f1b512a0b62d0fa47ee9b2612556f375e1f22fc2)
- [7.4.2303 : when using "is" the mode isn't always updated](https://github.com/vim/vim/commit/779f2fc3a7468e273897d2fd0672315812a2e3da)
- [7.4.2304 : in timer callback the timer itself can't be found or stopped](https://github.com/vim/vim/commit/417ccd7138d4d230d328de8b0d3892dd82ff1bee)
- [7.4.2305 : marks, writefile and nested function tests are old style](https://github.com/vim/vim/commit/19a1669ffc796e30a83c5600f82f12ebf63a2261)
- [7.4.2306 : default value for 'langremap' is wrong](https://github.com/vim/vim/commit/da9ce2cde11ddd0e16cdfbab6d4ac4e8110218e1)
- [7.4.2307 : several tests are old style](https://github.com/vim/vim/commit/cd055da370114f66c960be9c8b1eb0f33a9e0a85)
- [7.4.2308 : (after 7.4.2307) old charsearch test still listed in Makefile](https://github.com/vim/vim/commit/44f660c42191062214552f2bdb2e9bcc4e23e6fc)
- [7.4.2309 : crash when doing tabnext in a BufUnload autocmd](https://github.com/vim/vim/commit/11fbc2866ccc11b4dd1726abdaf582a78ef3f743)
- [7.4.2310 : (after 7.4.2304) accessing freed memory if timer not repeats](https://github.com/vim/vim/commit/a5d4121eed942bbbc672a7b00bf20866c17dc858)
- [7.4.2311 : Appveyor 64 bit build still using Python 3.4](https://github.com/vim/vim/commit/dd905a2ae14bf2ee59a068f1c3acbca1ff0b7067)
- [7.4.2312 : crash when autocommand moves to another tab](https://github.com/vim/vim/commit/5a49789a9b1f6447aeafbbbdd5b235dd10c471d5)
- [7.4.2313 : crash when deleting an augroup and listing an autocommand](https://github.com/vim/vim/commit/b62cc36a600e2e1e5a1d1d484fef89898c847c4c)
- [7.4.2314 : no error when deleting an augroup while it's the current one](https://github.com/vim/vim/commit/de653f08805dde14424d417502a0480a6ad292f8)
- [7.4.2315 : insufficient testing for Normal mode commands](https://github.com/vim/vim/commit/87bc3f74598ae8c648957e5755000cc6cdbc89ce)
- [7.4.2316 : channel sort test is flaky](https://github.com/vim/vim/commit/f7f3e32fcc4984d87f26337316379c2e2d90f7f1)
- [7.4.2317 : (after 7.4.2315) normal mode tests fail on MS-Windows](https://github.com/vim/vim/commit/0913a1089a07ac7b17abc3a1343dfa7cd25613f4)
- [7.4.2318 : when 'incsearch' not set CTRL-T/CTRL-G not inserted as before](https://github.com/vim/vim/commit/349e7d94e6bbb253bb87adad9039f095128ab543)
- [7.4.2319 : no way for a system wide vimrc to stop loading defaults.vim](https://github.com/vim/vim/commit/b07a82b6d5f904ed3e623e775c3458adb1cc0a3e)
- [7.4.2320 : redraw problem when using 'incsearch'](https://github.com/vim/vim/commit/dda933d06c06c2792bd686d059f6ad19191ad30b)
- [7.4.2321 : when a test is commented out we forget about it](https://github.com/vim/vim/commit/dac1947bb366ef43cd6da95acc730554e76d8b84)
- [7.4.2322 : access memory beyond the end of the line](https://github.com/vim/vim/commit/bc54f3f3fed4dc3556df8c46cee6739d211b0eb2)
- [7.4.2323 : using freed memory when using 'formatexpr'](https://github.com/vim/vim/commit/d77f9d595eb5f301b39b4373f2900a13c0ca30e2)
- [7.4.2324 : crash when BufUnload autocommand wipes out the new buffer](https://github.com/vim/vim/commit/e0ab94e7123ca7855f45919114d948ef2bc1e8c3)
- [7.4.2325 : (after 7.4.2324) tiny build fails](https://github.com/vim/vim/commit/30180b8dad5c1478e7920e56a71352cb318fadb0)
- [7.4.2326 : illegal memory access when invalid Visual selection position](https://github.com/vim/vim/commit/d5824ce1b5491df7d2eb0b66189d366fa67b4585)
- [7.4.2327 : freeing a variable that is on the stack](https://github.com/vim/vim/commit/c8fe338d64cc6183c03d4c12b1e036a7745e2932)
- [7.4.2328 : crash when BufWinLeave autocmd goes to another tab page](https://github.com/vim/vim/commit/f9e687e0681a250e1549ab27b6c7ef2c500395e3)
- [7.4.2329 : error for min() and max() contains %s](https://github.com/vim/vim/commit/26b84339fd8766898bcf6a259cbc2e0c38689726)
- [7.4.2330 : coverity complains about not checking curwin to be NULL](https://github.com/vim/vim/commit/030cddc7ec0c3d2fe3969140cd1b92b2f18633c0)
- [7.4.2331 : C-X C-V in Insert mode fails after entering expression](https://github.com/vim/vim/commit/33a80eeb859a78ba93432da6fa585786cfd77249)
- [7.4.2332 : crash when stop&#x5f;timer() usd in callback of a callback](https://github.com/vim/vim/commit/75537a93e985ef32e6c267b06ce93629855dd983)
- [7.4.2333 : outdated comments in normal mode test](https://github.com/vim/vim/commit/31845093b7f1b33e0c7e9e592bef65528674a1f2)
- [7.4.2334 : on MS-Windows test&#x5f;getcwd leaves Xtopdir behind](https://github.com/vim/vim/commit/1b0c2fcf6e85c9b85c24757ba970061e1f3e4e80)
- [7.4.2335 : taglist() is slow](https://github.com/vim/vim/commit/72b4b870fcc445c14faf282e0595b5f9406b101d)
- [7.4.2336 : running normal mode tests leave a couple of files behind](https://github.com/vim/vim/commit/df0db16cf74281a83a9ea3388a2ef9aedccd013c)
- [7.4.2337 : taglist() is still slow](https://github.com/vim/vim/commit/77b642501ac4c84b4e6820d4e5636eb856faaea7)
- [7.4.2338 : (after 7.4.2335) can't build with small features](https://github.com/vim/vim/commit/7947312871e7d01cdba058199904c212ec32f1c0)
- [7.4.2339 : tab page test fails when run as fake root](https://github.com/vim/vim/commit/100f5c90f4d4fb40bc3aeabc35192db371f5988f)
- [7.4.2340 : MS-Windows: Building with Ruby uses old version](https://github.com/vim/vim/commit/6384c5db8dda70076c878d393ba19a1510695228)
- [7.4.2341 : tiny things; test doesn't clean up properly](https://github.com/vim/vim/commit/aeac9006d5d14910f214f09df52c026a5936e737)
- [7.4.2342 : typo in MS-Windows build script](https://github.com/vim/vim/commit/1b0c1d050e28f351650c0c84e234fd24f12e5e29)
- [7.4.2343 : too many old style tests  ](https://github.com/vim/vim/commit/53f1673cd909eb1c809c6a9086e3d104a0df9bed)
- [7.4.2344 : the "Reading from channel output..." message can be unwanted](https://github.com/vim/vim/commit/169ebb080454357279ad5ad21ac532deaec605e8)
- [7.4.2345 : (after 7.4.2340) using old versions in MS-Windows makefiles](https://github.com/vim/vim/commit/0eaadec6b275a8add49242e1940855fcd154ba64)
- [7.4.2346 : autocommand test fails when run directly](https://github.com/vim/vim/commit/f18c4dbbe545757ce93563b25380e6f010340b4e)
- [7.4.2347 : crash when closing a buffer while Visual mode is active](https://github.com/vim/vim/commit/c4a908e83690844b0d3a46124ba6af7d23485d69)
- [7.4.2348 : crash on exit when EXITFREE is defined](https://github.com/vim/vim/commit/9a27c7fde6d453d9892b6f6baa756bce4d6d419d)
- [7.4.2349 : valgrind reports using uninitialzed memory](https://github.com/vim/vim/commit/2321ca2a78286bc026fa7f407281ddbeb04114bb)
- [7.4.2350 : test 86 and 87 fail with some version of Python](https://github.com/vim/vim/commit/cf703fe9d66744ffd7d9c625ee1cbe048ee07740)
- [7.4.2351 : netbeans test fails when run from unpacked MS-Windows sources](https://github.com/vim/vim/commit/ff330ff2c4500dc5535631418d601194654e4536)
- [7.4.2352 : netbeans test fails in shadow directory](https://github.com/vim/vim/commit/922a443cfdb1a37ef8bdeadf84f8b821ec4b37bd)
- [7.4.2353 : not enough test coverage for Normal mode commands](https://github.com/vim/vim/commit/2931f2a5df0d962032d41060af84d9fd2cb35c9f)
- [7.4.2354 : nested backreferences do not work with new regexp engine](https://github.com/vim/vim/commit/d563883a1fb5ec6cf4a2758c5e36ac1ff4e9bb3d)
- [7.4.2355 : regexp fails to match when using "\\>\\)\\?"](https://github.com/vim/vim/commit/16b3578f355282846f2600ce77fb344950f0b9ce)
- [7.4.2356 : reading past end of line when using previous subst. pattern](https://github.com/vim/vim/commit/ea683da58cf9ecf3afab9d650d3d2da76e5298d3)
- [7.4.2357 : attempt to read history entry while not initialized](https://github.com/vim/vim/commit/46643713dc6bb04b4e84986b1763ef309e960161)
- [7.4.2358 : compiler warnings with Solaris Studio when using GTK3](https://github.com/vim/vim/commit/d47d83745ff450232328ca7a4b8b00b31bad22fc)
- [7.4.2359 : memory leak in timer&#x5f;start()](https://github.com/vim/vim/commit/26fe0d56912e42c2b16a61b2480e19ba569aee98)
- [7.4.2360 : invalid memory access when formatting](https://github.com/vim/vim/commit/80c3fd7c559c7d329d57afe10db9bfb0adf10e46)
- [7.4.2361 : checking for last&#x5f;timer&#x5f;id to overflow is not reliable](https://github.com/vim/vim/commit/ee39ef0b93d31763d05e54ba99801e3f1a254c0d)
- [7.4.2362 : illegal memory access with ":1@"](https://github.com/vim/vim/commit/4930a76a0357f76a829eafe4985d04cf3ce0e9e0)
- [7.4.2363 : superfluous function prototypes](https://github.com/vim/vim/commit/cbeab96f87b91ed4b096216c397552b7f6d270f9)
- [7.4.2364 : sort test sometimes fails](https://github.com/vim/vim/commit/e1c8c7a6742be6072290f9aa54ae358060d9c42f)
- [7.4.2365 : needless line break; confusing directory name](https://github.com/vim/vim/commit/1b010058235fb803c1d4f42a02d2883921be8ef4)
- [7.4.2366 : MS-Windows gvim.exe does not have DirectX support](https://github.com/vim/vim/commit/100f27721b2f901eae348ef3635897e8be5d0307)
- [7.4.2367 : (after 7.4.2364) test runner misses a comma](https://github.com/vim/vim/commit/edeb846c1f04a49466992077eaea3396838bf4fd)
- [8.0.0001 : intro screen still mentions version7](https://github.com/vim/vim/commit/6401024869dcd722ee1b98b4ba58e6caa215201b)
- [8.0.0002 : the netrw plugin does not work](https://github.com/vim/vim/commit/9e1e7afc1f29b67fe3721442b49af3b20644a3ee)
- [8.0.0003 : getwinvar() returns wrong Value of boolean and number options](https://github.com/vim/vim/commit/789a5c0e3d27f09456678f0cfb6c1bd2d8ab4a35)
- [8.0.0004 : error messagge of function() with wrong argument says NULL](https://github.com/vim/vim/commit/5582ef14384525e8cec86016876d97a6b32dd548)
- [8.0.0005 : netbeans test fails with Python 3](https://github.com/vim/vim/commit/3c4ebeba17399e4ef33d004c269e667f696f0eb6)
- [8.0.0006 : ":lb" is interpreted as ":lbottom" instead of ":lbuffer"](https://github.com/vim/vim/commit/ebdd90ac282909c0bfcd6e83e70505abbbd5a38d)
- [8.0.0007 : Vim 7.4 is still mentioned in a few places](https://github.com/vim/vim/commit/bbe917d9d465a66a29e38879c7f66df51b0b0dc3)
- [8.0.0008 : popup complete test is disabled](https://github.com/vim/vim/commit/9e02cfa226b2577ec867b544a1a450a428a19880)
- [8.0.0009 : unnecessary workaround for AppVeyor](https://github.com/vim/vim/commit/3f1c15b24a3939d193c80c509132db9ee4d8cceb)
- [8.0.0010 : crash when editing file that starts with crypt header](https://github.com/vim/vim/commit/680e015bfe19be6772d3bd754486fbd45c1a9d3b)
- [8.0.0011 : on OSX Test&#x5f;pipe&#x5f;through&#x5f;sort&#x5f;all() sometimes fails](https://github.com/vim/vim/commit/947000d6e481042d3869bd65b149d72511c31ab2)
- [8.0.0012 : typos in comments](https://github.com/vim/vim/commit/9af418427652562384744648d7d173a4bfebba95)
- [8.0.0013 : (after 8.0.0011) missing comma in list](https://github.com/vim/vim/commit/c79d6aa01d86cd98cffddc80825c1c600e833bbf)
- [8.0.0014 : crypt tests are old style](https://github.com/vim/vim/commit/1eceadaf481e34ed8155011534159775697ce884)
- [8.0.0015 : can't tell which part of a channel has "buffered" status](https://github.com/vim/vim/commit/7ef3810d28b7ab2edbfcafab3fe8ad8bc2c2f138)
- [8.0.0016 : (after 8.0.0015) build fails](https://github.com/vim/vim/commit/0e77b7691ee7c477facb4c9d9162c8603ada2d84)
- [8.0.0017 : cannot get the number of current quickfix or location list](https://github.com/vim/vim/commit/890680ca6364386fabb271c85e0755bcaa6a33c1)
- [8.0.0018 : when using ":sleep" channel input is not handled](https://github.com/vim/vim/commit/b9c31e71f5a4653a27e81c21226039bfa35b9131)
- [8.0.0019 : test&#x5f;command&#x5f;count is old style](https://github.com/vim/vim/commit/b3435b0a3a0967115658d0a8c0224a28969cfa02)

## 新着スクリプト

- [hindent : Automatic Haskell code reformatting](http://www.vim.org/scripts/script.php?script_id=5437)
- [SearchCompleteMotion : Updated SearchCompete plugin with support for motions](http://www.vim.org/scripts/script.php?script_id=5438)
- [vim-sml : Vim support for Standard ML programming](http://www.vim.org/scripts/script.php?script_id=5439)
- [WWDC16 : Colorful dark color scheme inspired by WWDC16 page](http://www.vim.org/scripts/script.php?script_id=5440)
- [stardict.vim : Quick access to stardict dicts](http://www.vim.org/scripts/script.php?script_id=5441)
- [wordbyword.vim : Word by word search in dictionaries over cursor.](http://www.vim.org/scripts/script.php?script_id=5442)
- [rimas.vim : Diccionario de rimas para Vim ](http://www.vim.org/scripts/script.php?script_id=5443)
- [git-switcher.vim : Session management based on git branch.](http://www.vim.org/scripts/script.php?script_id=5446)
- [pgn.vim : syntax highlighting for .pgn (chess) files ](http://www.vim.org/scripts/script.php?script_id=5447)
- [vim-zim : Zim desktop wiki syntax and helper functions](http://www.vim.org/scripts/script.php?script_id=5448)
- [ALE - Asynchronous Lint Engine : Lint while you type in Vim](http://www.vim.org/scripts/script.php?script_id=5449)
- [unicon : uniform contrast light and dark colorscheme](http://www.vim.org/scripts/script.php?script_id=5450)
- [vim-seattle : An elegant yet efficient color scheme for coders](http://www.vim.org/scripts/script.php?script_id=5451)
- [Buffy.vim : Switch between buffers with a handy menu](http://www.vim.org/scripts/script.php?script_id=5452)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (1517)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1303)
3. [c.vim : C/C++ IDE -- Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (718)
4. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (632)
5. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (507)
6. [winmanager : A windows style IDE for Vim 6.0](http://www.vim.org/scripts/script.php?script_id=95) (476)
7. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (461)
8. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (451)
9. [SuperTab : Do all your insert-mode completion with Tab.](http://www.vim.org/scripts/script.php?script_id=1643) (446)
10. [wombat256.vim : Wombat for 256 color xterms](http://www.vim.org/scripts/script.php?script_id=2465) (430)

## vim-jp issues

Open : 237 (+0) | Closed : 721 (+8)

- [Issue #951 : 翻訳関連の問題](https://github.com/vim-jp/issues/issues/951)
- [Issue #952 : \[RFC\] Windows の &#x5f;vimrc や &#x5f;gvimrc は AppData\\Roadming の下も読むべきじゃないか](https://github.com/vim-jp/issues/issues/952)
- [Issue #953 : src/po の make check の改善](https://github.com/vim-jp/issues/issues/953)
- [Issue #954 : イベント発生時にtabeditすると現在のバッファが新規タブに開かれる](https://github.com/vim-jp/issues/issues/954)
- [Issue #955 : $VIM/&#x5f;vimrc が読み込まれない](https://github.com/vim-jp/issues/issues/955)
- [Issue #956 : コマンドの出力末尾が改行の場合の systemlist() の結果がおかしい](https://github.com/vim-jp/issues/issues/956)
- [Issue #957 : job&#x5f;start から起動するコマンドが直ぐに終了するとcallbackが呼ばれない](https://github.com/vim-jp/issues/issues/957)
- [Issue #958 : Inconsistency to handle stderr between &#x60;ch&#x5f;status()&#x60; and &#x60;ch&#x5f;read()&#x60;](https://github.com/vim-jp/issues/issues/958)
