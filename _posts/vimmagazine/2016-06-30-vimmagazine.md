---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2016 年 6 月号

---

## 話題

*   [netupvim 1.0 リリース](http://www.kaoriya.net/blog/2016/06/06/) 最新版は [1.1](https://github.com/koron/netupvim/releases/tag/v1.1) であることに注意してください
*   [Software Design 2016年7月号 Vimの細道【9】カーソルを自在に操る](http://gihyo.jp/magazine/SD/archive/2016/201607)
*   [なぜAtomはVimを置き換えられないのか](https://yakst.com/ja/posts/3985)

## 今月の新機能

*   7.4.1892: `'balloonexpr'` に指定した関数内で、定義済変数 `v:beval_winid` によりウィンドウIDを参照できるようになりました。
*   7.4.1893: バッファのウィンドウIDを取得する関数 `bufwinid()` が追加されました。
*   7.4.1894: `getchar()` でマウスクリックイベントを取得したときに、定義済変数 `v:mouse_winid` によりウィンドウIDが参照できるようになりました。
*   7.4.1895: ウィンドウIDが1000始まりになり、ウィンドウ番号(1始まり)との混乱が解消されました。
*   7.4.1898: ユーザーコマンドにコマンド修飾子 `<mods>` が指定できるようになりました。
*   7.4.1903: viminfo ファイルのフォーマットに変更がありました。
*   7.4.1925: `:clearjumps` コマンドが追加されました。
*   7.4.1952: `:cscope find` のクエリタイプに `a` が追加されました。 `:help cscope-find`
*   7.4.1960: Unicode 9.0 に対応

## リリース情報

- [7.4.1863 : compiler warnings on Win64](https://github.com/vim/vim/commit/9b0ac229bcfc91acabd35fc576055a94c1687c32)
- [7.4.1864 : Python: encoding error with Python 2](https://github.com/vim/vim/commit/22081f4a3397704645841121d994058abd6cb481)
- [7.4.1865 : memory leaks in test49](https://github.com/vim/vim/commit/9ad73239c26467832a5b553b2a4b99d7ffbaa25e)
- [7.4.1866 : invalid memory access when exiting with EXITFREE defined](https://github.com/vim/vim/commit/a96732150cda2f242133228579b05437a39b8daa)
- [7.4.1867 : memory leak in test&#x5f;matchstrpos](https://github.com/vim/vim/commit/3c809343c72d9964475f421fd03bb892bc584a51)
- [7.4.1868 : setting really&#x5f;exiting causes memory leaks to be reported](https://github.com/vim/vim/commit/b89a25f17e274dc308c584ea69a129ffbb26bc3d)
- [7.4.1869 : can't build with old version of Perl](https://github.com/vim/vim/commit/cf190c6f017563de1bdbf854b3376522b8b2748f)
- [7.4.1870 : (after 7.4.1863) one more Win64 compiler warning](https://github.com/vim/vim/commit/dec6c7beda5491b28f67f430fc9d01467050c1ca)
- [7.4.1871 : appending to quickfix list is slow if quickfix window is open](https://github.com/vim/vim/commit/864293abb72d62604d8d6b458addfb43c14230c3)
- [7.4.1872 : still build problem with old version of Perl](https://github.com/vim/vim/commit/c4bc0e6542185b659d2a165b635f9561549071ea)
- [7.4.1873 : GUI: when a callback adds a timer it is not used until later](https://github.com/vim/vim/commit/4231da403e3c879dd6ac261e51f4ca60813935e3)
- [7.4.1874 : unused variable in Win32 code](https://github.com/vim/vim/commit/ae3f33040ba752e7f6b722d2fe2b2ebeeb8c56b4)
- [7.4.1875 : comparing functions and partials doesn't work well](https://github.com/vim/vim/commit/8e759ba8651428995b338b66c615367259f79766)
- [7.4.1876 : typing "k" at the hit-enter prompt has no effect](https://github.com/vim/vim/commit/a0055ad3a789b8eeb0c983d8a18d4bcaeaf456b8)
- [7.4.1877 : no test for invoking "close&#x5f;cb" when writing to a buffer](https://github.com/vim/vim/commit/01d46e41ba4967ee534db4b94ad642007634841e)
- [7.4.1878 : exited job isn't detected until a character is typed](https://github.com/vim/vim/commit/cf7c11a9479ba7ce775b86c7a846fae48321d260)
- [7.4.1879 : (after 7.4.1877) channel test is flaky](https://github.com/vim/vim/commit/e38a2f7ebda791bcecccd657919ae9f7e2f6438c)
- [7.4.1880 : MS-Windows console build defaults to not having +channel](https://github.com/vim/vim/commit/76929af43134b4222b33648b6c53754a34f24524)
- [7.4.1881 : appending to a long quickfix list is slow](https://github.com/vim/vim/commit/83e6d7ac6a1c2a0cb5ee6c8420a5dc792f1d5ffa)
- [7.4.1882 : check for line break at end of line wrong](https://github.com/vim/vim/commit/b37662a0fbb952838fca87aff4d26b596030b67b)
- [7.4.1883 : cppcheck found 2 incorrect printf formats](https://github.com/vim/vim/commit/ea0345901cc2af29f9c5dd0d9d8a818d5f96a63b)
- [7.4.1884 : updating marks in long quickfix list is very slow](https://github.com/vim/vim/commit/2f095a4bc4d786e0ac834f48dd18a94fe2d140e3)
- [7.4.1885 : MinGW console build defaults to not having +channel](https://github.com/vim/vim/commit/1aa07bdead2e93501c54591e31fe73b9b09c58b5)
- [7.4.1886 : mapping times out when interrupted by receiving data](https://github.com/vim/vim/commit/cda7764d8e65325d4524e5d6c3174121eeb12cad)
- [7.4.1887 : when receiving channel data 'updatetime' is not respected](https://github.com/vim/vim/commit/e30a3d01dc1b6b066748b1d61cced05cf92b3c80)
- [7.4.1888 : wrong computation of remaining wait time in RealWaitForChar()](https://github.com/vim/vim/commit/76b6dfe54ba9b85cd9d8e6539205c1679a187961)
- [7.4.1889 : when umask is set to 0177 Vim can't create temp files](https://github.com/vim/vim/commit/35d88f4e2ff5dcd9904f04612d5febede996137c)
- [7.4.1890 : GUI: cursor blinking interrupted when channel data received](https://github.com/vim/vim/commit/703a8044b5393d37d355b0b1054a9a5a13912a3f)
- [7.4.1891 : channel reading very long lines is slow](https://github.com/vim/vim/commit/9ed96efb3d47d46e9637da04656efff715102407)
- [7.4.1892 : balloon eval only gets the window number, not the ID](https://github.com/vim/vim/commit/c9721bdc63378cc6123e775ffe43e9cba30322b3)
- [7.4.1893 : cannot easily get the window ID for a buffer](https://github.com/vim/vim/commit/b3619a90eae2702553ff9494ecc4c9b20c13c224)
- [7.4.1894 : cannot get the window ID for a mouse click](https://github.com/vim/vim/commit/511972d810ea490955161ff5097ec2f57919ceaf)
- [7.4.1895 : cannot use a window ID where a window number is expected](https://github.com/vim/vim/commit/888ccac8902cee186fbd47e971881f6d9b19c068)
- [7.4.1896 : invoking mark&#x5f;adjust() when it is pointless](https://github.com/vim/vim/commit/82faa259cc42379f2a17d598a2a39d14048685b0)
- [7.4.1897 : various typos, long lines and style mistakes](https://github.com/vim/vim/commit/fd89d7ea81b18d32363456b16258174dc9e095dc)
- [7.4.1898 : user commands don't support modifiers](https://github.com/vim/vim/commit/63a60ded3fd584847a05dccf058026e682abad90)
- [7.4.1899 : GTK 3: cursor blinking doesn't work well](https://github.com/vim/vim/commit/8e31fd52ec8be6ef1dc600b637d5f099d55e1715)
- [7.4.1900 : using CTRL-\] in the help on "{address}." doesn't work](https://github.com/vim/vim/commit/28b942a064dd486cc241894b625ab72f5a5c6d1b)
- [7.4.1901 : Win32: the "Disabled" menu items would appear enabled](https://github.com/vim/vim/commit/762f1754370a1278167c8cba6c047ef319fc099c)
- [7.4.1902 : no test for collapsing channel buffers; some text is lost](https://github.com/vim/vim/commit/bbe8d91e695184771d7e45315258fab8eb3e6b07)
- [7.4.1903 : recent history items may be dropped when writing viminfo](https://github.com/vim/vim/commit/45d2eeaad66939348893b9254171067b0457cd9d)
- [7.4.1904 : (after 7.4.1903) build fails](https://github.com/vim/vim/commit/c60954b48ec225c75558abee56c4a9cc89f84b7e)
- [7.4.1905 : (after 7.4.1903) some compilers can't handle double semicolon](https://github.com/vim/vim/commit/fdd82fe365d0e287bafc71f98c039cb5af8ed827)
- [7.4.1906 : collapsing channel buffers does not work properly](https://github.com/vim/vim/commit/5f1032d2a55b9417a0a6fa225e35089c98a5a419)
- [7.4.1907 : warnings from 64 bit compiler](https://github.com/vim/vim/commit/1d5f1d07aedb6f149f5de145b1dfd6528a769c93)
- [7.4.1908 : Netbeans uses uninitialzed pointer and freed memory](https://github.com/vim/vim/commit/5ce4a0b96ab688b1ea2481c2516e2889ff6713bf)
- [7.4.1909 : doubled semicolons](https://github.com/vim/vim/commit/945ec093cd4ddefab930239990564b12eb232153)
- [7.4.1910 : tests using external command to delete directory](https://github.com/vim/vim/commit/abc70bbf363dbbe3f2bf714102f55648a512791e)
- [7.4.1911 : recent history lines may be lost when exiting Vim](https://github.com/vim/vim/commit/1fd99c1ca89a3d13bb53aff4a5a8f5ee740713e5)
- [7.4.1912 : no test for using setqflist() on an older quickfix list](https://github.com/vim/vim/commit/1cee693b310e1494115a1677fac064941092e1bb)
- [7.4.1913 : after ":doautocmd" modelines are always processed](https://github.com/vim/vim/commit/1610d052413e0ed664498853a47acc2d677a22d1)
- [7.4.1914 : crash when executing autocommands when on the signal stack](https://github.com/vim/vim/commit/e429e70f050cb2941f1f8427cf918b68444c904e)
- [7.4.1915 : effect of the PopupMenu autocommand isn't directly visible](https://github.com/vim/vim/commit/2a67ed83a885e8e65e25498660dae4f34d532c7b)
- [7.4.1916 : (after 7.4.1906) no proper test for what 7.4.1906 fixes](https://github.com/vim/vim/commit/b8aefa46adf5e825118716e142fab7ef32076475)
- [7.4.1917 : history lines read from viminfo are not converted](https://github.com/vim/vim/commit/012270936c3c7df3bba45ad2b48938c23a2fd43a)
- [7.4.1918 : not enough testing for parsing viminfo lines](https://github.com/vim/vim/commit/62f8b4e18014b259bcde4a2845c602b0a44a3714)
- [7.4.1919 : register contents is not merged when writing viminfo](https://github.com/vim/vim/commit/46bbb0c4ba27395859dfeaa26938483946bb4ec2)
- [7.4.1920 : (after 7.4.1919) missing test changes](https://github.com/vim/vim/commit/e80ff7448d886805b459250b090aaa20d96e8c2b)
- [7.4.1921 : (after 7.4.1919) vim&#x5f;time() not included when needed](https://github.com/vim/vim/commit/724f2edf75bf084665b3e2cdfa7ab66ccb1b5c0c)
- [7.4.1922 : Ruby 2.4.0 unifies Fixnum and Bignum into Integer](https://github.com/vim/vim/commit/06469e979fe524ac6cb8f705ed4221aa267de11d)
- [7.4.1923 : command line editing is not tested much](https://github.com/vim/vim/commit/ae3150ec8d9da4a244acffebea55416946ca23d3)
- [7.4.1924 : missing "void" for functions without argument](https://github.com/vim/vim/commit/cf089463492fab53b2a5d81517829d22f882f82e)
- [7.4.1925 : viminfo does not merge file marks properly](https://github.com/vim/vim/commit/2d35899721da0e9359a9fe1059554f8c4ea7f0c1)
- [7.4.1926 : possible crash with many history items](https://github.com/vim/vim/commit/a890f5e34887bff7616bdb4b9ee0bf98c8d2a8f0)
- [7.4.1927 : compiler warning for signed/unsigned](https://github.com/vim/vim/commit/5b7d177e8994c003ae77ddab5bd54f8cd1ee181b)
- [7.4.1928 : overwriting pointer argument](https://github.com/vim/vim/commit/76ae22fef3cb224ca7fbf97517f881e825d4d0c2)
- [7.4.1929 : inconsistent indenting and weird name](https://github.com/vim/vim/commit/2016ae586b12513d973aabc30ed758b543114cbe)
- [7.4.1930 : can't build without +spell but with +quickfix](https://github.com/vim/vim/commit/bf2c0ee0b2e8465847cf67aff22dd4bda7be81d1)
- [7.4.1931 : using both old and new style file mark lines from viminfo](https://github.com/vim/vim/commit/a641e1d4da3f9152c489318c06a93fcd1c746637)
- [7.4.1932 : jumplist is not merged with the one in the viminfo file](https://github.com/vim/vim/commit/ece74ab103eca15e17435efbe9cb21039787f1ea)
- [7.4.1933 : compiler warning about uninitialzed variable](https://github.com/vim/vim/commit/72e697d189616265ecefe0df4509d476df3bae40)
- [7.4.1934 : new style tests not executed with MinGW compiler](https://github.com/vim/vim/commit/9bd31860521b53ad657c4d6d320d670e30a049f8)
- [7.4.1935 : with GUI search/replace second match is skipped](https://github.com/vim/vim/commit/bee666f239eada035d288b77269aebc42f644ea6)
- [7.4.1936 : off-by-one error in bounds check](https://github.com/vim/vim/commit/67e3720a9ddd8a9d2e8344358c28fa1f4196db0d)
- [7.4.1937 : no test for directory stack in quickfix](https://github.com/vim/vim/commit/ab47c61f46f0797308217cd3c045472cb0ec3195)
- [7.4.1938 : when writing viminfo numbered marks were duplicated](https://github.com/vim/vim/commit/36f0f0686ca313ef7b76387378cd5dc7acea1924)
- [7.4.1939 : memory access error when reading viminfo](https://github.com/vim/vim/commit/28607ba2b82668503f8406bc13690d59af46deb3)
- [7.4.1940 : "gd" hangs in some situations](https://github.com/vim/vim/commit/23c60f21b07b04351d846e6fbf4f4abd9aa09345)
- [7.4.1941 : not all quickfix tests are also done with the location lists](https://github.com/vim/vim/commit/3ef5bf7d459d6b8a21aaefc80f65448f5a7aa59f)
- [7.4.1942 : background is not drawn properly when 'termguicolors' is set](https://github.com/vim/vim/commit/d18f672fc9477f3c0cb7cc4ce8d9237ed825c612)
- [7.4.1943 : Coverity warns for unreachable code](https://github.com/vim/vim/commit/678e480529bce7ba3c09e71233249e3b56531b24)
- [7.4.1944 : Win32: Cannot compile with XPM feature using VC2015](https://github.com/vim/vim/commit/b5b95750a688d1fabafc6dbc2f31df90d5b5a17e)
- [7.4.1945 : the Man plugin doesn't work that well](https://github.com/vim/vim/commit/ddf8d1c746ee081d15c9d7e0515f6ac43adbf363)
- [7.4.1946 : (after 7.4.1944) file list does not include new XPM libraries](https://github.com/vim/vim/commit/cf2d77987ca2af092dad5d4cbca9eda621fe5567)
- [7.4.1947 : viminfo continuation line with wrong length isn't skipped](https://github.com/vim/vim/commit/ecefe71704850b94df44f65fc756c1551ec68388)
- [7.4.1948 : Ctrl-A with double-byte encoding may result in garbled text](https://github.com/vim/vim/commit/ad5ca9bc1e7145474adb082775a805f1731e9e37)
- [7.4.1949 : minor problems with the quickfix code](https://github.com/vim/vim/commit/38df43bd13a2498cc96b3ddd9a20dd75126bd171)
- [7.4.1950 : quickfix long lines test not executed for buffer](https://github.com/vim/vim/commit/f50df3925b7c909d1cda4c868d8c7ba38189aabe)
- [7.4.1951 : Ruby test is old style](https://github.com/vim/vim/commit/85babd6db65afb0eb06a7a9a0778d692248c5c2b)
- [7.4.1952 : Cscope interface does not support finding assignments](https://github.com/vim/vim/commit/b12e7ef956e0b0344778b7ef93d41f4b4ed2a670)
- [7.4.1953 : not all parts of the quickfix code are tested](https://github.com/vim/vim/commit/049cba9e9760152b5695399a991dc61cea9ba143)
- [7.4.1954 : (after 7.4.1948) no test for what 7.4.1948 fixes](https://github.com/vim/vim/commit/d388d2ac8bf8c770bf97dc406e99a20ba5104855)
- [7.4.1955 : using 32-bit Perl with 64-bit time&#x5f;t causes memory corruption](https://github.com/vim/vim/commit/f4fba6dcd508cb369ffa6916d9cb3fcf3d7ed548)
- [7.4.1956 : when file doesn't load with CTRL-W f new window is not closed](https://github.com/vim/vim/commit/5d2ca0402954ff79b73d9c86cc16c8a6454b75a7)
- [7.4.1957 : Perl interface has obsolete workaround](https://github.com/vim/vim/commit/7b61bf187a318cb710be40da9ce4c29972324a71)
- [7.4.1958 : Perl interface preprocessor statements not nicely indented](https://github.com/vim/vim/commit/eeb50ab5228c5c09743a9c2b907c3634c0146e84)
- [7.4.1959 : crash when running test&#x5f;channel.vim on Windows](https://github.com/vim/vim/commit/aad30bbcde2b268e2fffb6fd1443f27ebda9d1ff)
- [7.4.1960 : Unicode standard 9 was released](https://github.com/vim/vim/commit/04e2b4b0c4866586ecce3d1567f9b0bdeeb31f15)
- [7.4.1961 : popup menu remains if 'insertmode' is reset during completion](https://github.com/vim/vim/commit/00672e1d3f59dbff91a18d418b2984be96f89ee5)
- [7.4.1962 : two test files for increment/decrement](https://github.com/vim/vim/commit/52df117df724a7ad20affdf15214d2df2e507dcc)
- [7.4.1963 : running Win32 Vim in mintty does not work](https://github.com/vim/vim/commit/97ff9b9cffd97219d888874b9b3811d55e99c78f)
- [7.4.1964 : the quickfix init function is too big](https://github.com/vim/vim/commit/688e3d1fd9b9129a5ba0e0d599ccfe6f4443daf3)
- [7.4.1965 : job in raw mode appends garbage characterss to buffer](https://github.com/vim/vim/commit/adb78a77ebb47627bcf73bd16ac1119d970e17c8)
- [7.4.1966 : Coverity reports a resource leak](https://github.com/vim/vim/commit/bcf7772a23624edc0942120e564f6b4ac95604ad)
- [7.4.1967 : falling back from NFA to old regexp engine has problems](https://github.com/vim/vim/commit/6747fabc7348bf5f41ccfe851e2be3e900ec8ee0)

## 新着スクリプト

- [Ranger.vim : Ranger integration in vim and neovim](http://www.vim.org/scripts/script.php?script_id=5396)
- [Closing Brackets : A dead-simple Vim plugin to automatically insert closing brackets](http://www.vim.org/scripts/script.php?script_id=5397)
- [verilog&#x5f;emacsauto.vim : Minor update based on vimscript #1875](http://www.vim.org/scripts/script.php?script_id=5398)
- [ember&#x5f;tools.vim : Tools for working with ember.js projects](http://www.vim.org/scripts/script.php?script_id=5399)
- [Cheat40 : A cheat sheet that makes sense, inside Vim!](http://www.vim.org/scripts/script.php?script_id=5400)
- [vim-gencode-cpp : auto generate c++ function definition or declaration](http://www.vim.org/scripts/script.php?script_id=5401)
- [columnMove.vim : move along vertical axis, to next/previous character-blocks end/start](http://www.vim.org/scripts/script.php?script_id=5402)
- [vimport : Allows you to easily add imports within java, groovy, grails](http://www.vim.org/scripts/script.php?script_id=5403)
- [searchant.vim : Vim plugin for improved search highlighting](http://www.vim.org/scripts/script.php?script_id=5404)
- [BTN: Better tab new : Adds a tab new command accepting file:lin&#x5f;num as a parameter](http://www.vim.org/scripts/script.php?script_id=5405)
- [Impl Switcher : Easily switch between impl and header files](http://www.vim.org/scripts/script.php?script_id=5406)
- [fluid.vim : Plugin for the Fluid Template Engine (TYPO3), indenting, and syntax plugin.](http://www.vim.org/scripts/script.php?script_id=5407)
- [vim-java-ununsed-imports : Allows showing and removing unused imports in a java file.](http://www.vim.org/scripts/script.php?script_id=5408)
- [radical.vim : Convert decimal, hex, octal, binary number representations](http://www.vim.org/scripts/script.php?script_id=5409)
- [Srcery : Vim color scheme @ https://github.com/roosta/srcery](http://www.vim.org/scripts/script.php?script_id=5411)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (1442)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1279)
3. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (556)
4. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (549)
5. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (503)
6. [winmanager : A windows style IDE for Vim 6.0](http://www.vim.org/scripts/script.php?script_id=95) (501)
7. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (436)
8. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (420)
9. [bufexplorer.zip : Buffer Explorer / Browser](http://www.vim.org/scripts/script.php?script_id=42) (408)
10. [nginx.vim : initial version](http://www.vim.org/scripts/script.php?script_id=1886) (389)

## vim-jp issues

Open : 243 (-2) | Closed : 679 (+10)

- [Issue #915 : 端末機能、作っちゃわない？](https://github.com/vim-jp/issues/issues/915)
- [Issue #916 : OS XでVim (not MacVim)のビルドが通らない](https://github.com/vim-jp/issues/issues/916)
- [Issue #917 : split() 関数の挙動について](https://github.com/vim-jp/issues/issues/917)
- [Issue #918 : Make&#x5f;cyg&#x5f;ming.mak で lua と luajit を分けたい](https://github.com/vim-jp/issues/issues/918)
- [Issue #919 : jobの出力をrawモードで取得すると変な出力が混入する時がある](https://github.com/vim-jp/issues/issues/919)
- [Issue #920 : Ctrl-A/Ctrl-Xが文字コードによっては文字化けが発生する](https://github.com/vim-jp/issues/issues/920)
- [Issue #921 : sleep 中の入力が getchar(0) で取れない (on Solaris)](https://github.com/vim-jp/issues/issues/921)
- [Issue #922 : :help // が意図しないタグになっている](https://github.com/vim-jp/issues/issues/922)

