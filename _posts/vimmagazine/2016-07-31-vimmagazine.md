---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2016 年 7 月号

---

## 話題

*   [Osaka.vim #7](http://osaka-vim.connpass.com/event/32702/) 7/2(土) 大阪 が開催されました
*   [Software Design 2016年8月号 Vimの細道【10】 Vimからgitを使い倒す by mattn][sd201608]
*   [Vim の地味だけどよく使う設定](http://note103.hateblo.jp/entry/2016/07/06/114746)
*   [vimでソースを追うときの小技](http://qiita.com/masaharu-suizu/items/2e2fed71bf4074d05713)
*   [sudoを付けずにVimでルート権限のファイルを編集してしまった時](http://qiita.com/shishamo_dev/items/ab722c53301b95ccfac1)
*   [Vim v.s. Emacsをフローチャートで表現](http://rcmdnk.github.io/blog/2016/07/21/computer-vim-emacs/) - 編者注:フローチャートというのは誤った用語だと思われます
*   [VimとGitの効率的なコラボレーションの仕方](http://myenigma.hatenablog.com/entry/2016/07/10/084048)
*   [Vim-Pokemon ～あなたが Vim で開いているファイルに潜んでいるポケモン～](http://d.hatena.ne.jp/heavenshell/20160715/1468606360)
*   [Vimのおすすめプラグイン23個とその設定](http://post.simplie.jp/posts/56)
*   [VimのmapleaderにMacのCommandを割り当てたい](http://ja.stackoverflow.com/q/27794/208)
*   [vimのヤンクがmacのクリップボードに保存されない](https://teratail.com/questions/41191)
*   [DockerとWebSocketを使って、vimの設定をブラウザで即体感できるサービスを作った](http://mosuke5.hateblo.jp/entry/2015/07/19/135844)
*   [tmux 起動でvimなどがおかしくなった時の対処法](http://kirin.hatenadiary.jp/entry/2016/07/16/150206)

## 今月の新機能

*   7.4.1971: `:clist` に `+{count}` が指定できるようになりました。
*   7.4.1975: Windows で 2GB を超えるファイルが正しく扱えるようになりました。
*   7.4.1976: Vim script の整数の内部表現が 64bit になりました。
*   7.4.1989: `filter()` と `map()` に関数への参照を指定できるようになりました。
*   7.4.1997: `:cbottom` が追加されました。
*   7.4.2008: `execute()` が追加されました。
*   7.4.2010: `:lbottom` が追加されました。
*   7.4.2011: `getcompletion()` が追加されました。
*   7.4.2015: `test_autochdir()` が追加されました。
*   7.4.2029: `printf()` が64bit整数に対応しました。
*   7.4.2034: MinGWでのビルド時に `ARCH` が自動設定されるようになりました。
*   7.4.2044: lambda が導入されました。
*   7.4.2049: `:chistory` と `:lhistory` が追加されました。
*   7.4.2071: 型の値をあらわす `v:t_` で始まる定義済変数が追加されました。
*   7.4.2072: `substitute()` に関数への参照を指定できるようになりました。
*   7.4.2075: `TabNew` イベントと `WinNew` イベントが追加されました。
*   7.4.2077: `TabClosed` イベントが追加されました。
*   7.4.2095: `assert_inrange()` が追加されました。
*   7.4.2109: `display` オプションに `truncate` が追加されました。
*   7.4.2111: `.vimrc` が見つからなかった時に `defaults.vim` が読み込まれるようになりました。
*   7.4.2119: lambda内で外側のスコープの変数をキャプチャできるようになりました。
*   7.4.2120: ユーザー関数に `closure` を指定できるようになりました。
*   7.4.2121: lambda と closure のサポートを `+lambda` で確認できるようになりました。
*   [パッチ番号なし][travis+osx]: Travis CI で OSX のテストを実行するようになりました。

[travis+osx]:https://github.com/vim/vim/commit/b20617b0b0d15523961618717d6cd2195674c363
[sd201608]:http://gihyo.jp/magazine/SD/archive/2016/201608

## リリース情報

- [7.4.1968 : invalid memory access with "\\\<C-">](https://github.com/vim/vim/commit/1d90a5a5af84250e226f8a9121e771f7b72aa894)
- [7.4.1969 : crash when the netbeans channel is closed](https://github.com/vim/vim/commit/24cf233ef9e92f947618c55d8bff423898a45c24)
- [7.4.1970 : using ":insert" in an empty buffer sets the jump mark](https://github.com/vim/vim/commit/70e136e1d86ea1d795774824c7b712245912946d)
- [7.4.1971 : not easy to see error lines below the current error position](https://github.com/vim/vim/commit/e8fea0728a2fa1fe78ef0ac90dee1a84bd7ef9fb)
- [7.4.1972 : on Solaris select() does not work as expected with typeahead ](https://github.com/vim/vim/commit/40de45664c20e7ca46a28a3f472202f90e47f8bf)
- [7.4.1973 : on MS-Windows package directory added to 'rtp' at the end](https://github.com/vim/vim/commit/4c5717ed8a81f5ae9dfe4f38b17a61fc8421054b)
- [7.4.1974 : GUI has a problem with some termcodes](https://github.com/vim/vim/commit/4a6c670b844a3ef9aec865a8216eaf363bab8721)
- [7.4.1975 : on MS-Windows large files (> 2Gbyte) cause problems](https://github.com/vim/vim/commit/8767f52fbfd4f053ce00a978227c95f1d7d323fe)
- [7.4.1976 : number variables are not 64 bits while they could be](https://github.com/vim/vim/commit/22fcfad29276bd5f317faf516637dcd491b96a12)
- [7.4.1977 : with 64 bit changes don't need three calls to sprintf()](https://github.com/vim/vim/commit/bde9810d6103ffe3a22a9330021cb21db1ed1792)
- [7.4.1978 : (after 7.4.1975) large file test does not delete its output](https://github.com/vim/vim/commit/c5af40ae646ceda817eff93b4f9ba274f031bea6)
- [7.4.1979 : (after 7.4.1976) getting value of binary option is wrong](https://github.com/vim/vim/commit/2acfbed9dbea990f129535de7ff3df360365130b)
- [7.4.1980 : 'errorformat' is parsed for every call to ":caddexpr"](https://github.com/vim/vim/commit/361c8f0e517e41f1f1d34dae328044406fde80ac)
- [7.4.1981 : no testing for Farsi code](https://github.com/vim/vim/commit/ee2615af64fdcee87d8e4b13b65356e77fbd969b)
- [7.4.1982 : viminfo file contains duplicate change marks](https://github.com/vim/vim/commit/951fa1819ad0e74d73e885d0ec04dd6c70f8d2f7)
- [7.4.1983 : farsi.c and arabic.c are included in a strange way](https://github.com/vim/vim/commit/75464dc434c43efac60e8bfd9bec2a8b736407e9)
- [7.4.1984 : not all quickfix features are tested](https://github.com/vim/vim/commit/0fcc7c6dd1902b71e0e7d0a35ddabafef6455a83)
- [7.4.1985 : (after 7.4.1983) missing changes in VMS build file](https://github.com/vim/vim/commit/3b98b830017b8bf5fd270db93dfe187434df1960)
- [7.4.1986 : compiler warns for loss of data](https://github.com/vim/vim/commit/fef524bbff9aa186838c35212b2f89f61d627cf8)
- [7.4.1987 : viminfo may have useless continuation lines](https://github.com/vim/vim/commit/dec85cf75044ed94f611c825a7a0b0050a2597b9)
- [7.4.1988 : when updating viminfo with file marks there is no time order](https://github.com/vim/vim/commit/ab9c89b68dcbdb3fbda8c5a50dd90caca64f1bfd)
- [7.4.1989 : filter() and map() only accept a string argument](https://github.com/vim/vim/commit/b33c7eb5b813cb631b2b0ca5c4029e1788a09bde)
- [7.4.1990 : (after 7.4.1952) cscope items are not sorted](https://github.com/vim/vim/commit/80632db65e8f5f775dadbbc10c5ba6c173ebb24f)
- [7.4.1991 : glob() does not add symbolic link when there are no wildcards](https://github.com/vim/vim/commit/00efded1064427ab3f84e4d57af62e0aab876fc6)
- [7.4.1992 : values for true and false can be confusing](https://github.com/vim/vim/commit/e381d3d5e098546854b008e01ca1d28ba1a4a057)
- [7.4.1993 : not all TRUE and FALSE arguments are tested](https://github.com/vim/vim/commit/6bb450145e96d7b182769fd9502a267da72667ec)
- [7.4.1994 : (after 7.4.1993) true-false test fails](https://github.com/vim/vim/commit/05e418d436410cd8bbf5a29ff81e8ad68408b1e8)
- [7.4.1995 : GUI: cursor drawn in wrong place when using a timer callback](https://github.com/vim/vim/commit/9d5d3c9c4468ad76f16b50eabd3d9e7eab2ed44d)
- [7.4.1996 : capturing the output of a command takes a few commands](https://github.com/vim/vim/commit/1e5e1231ac9e1ba9678812c96f9d554a078eeec4)
- [7.4.1997 : cannot easily scroll the quickfix window](https://github.com/vim/vim/commit/dcb170018642ec144cd87d9d9fe076575b8d1263)
- [7.4.1998 : no NL to NUL conversion when writing buffer lines to a job](https://github.com/vim/vim/commit/bf2cc5f36d5ffd5de445e6970602000c7869b65a)
- [7.4.1999 : evalcmd() doesn't work recursively](https://github.com/vim/vim/commit/bc5d6dd1dd1dc3a06e4e655fc9479529db288365)
- [7.4.2000 : (after 7.4.1999) evalcmd test fails](https://github.com/vim/vim/commit/245a7cb6d33e556415c9fcaa874b784a6d8f0ef3)
- [7.4.2001 : (after 7.4.2000) tiny build fails](https://github.com/vim/vim/commit/ba768495c2e6bae74e49d22d489fbf211ecad55e)
- [7.4.2002 : crash when passing number to filter() or map()](https://github.com/vim/vim/commit/a06ec8f345eabb66e5b7d7c0192cfebdde63115d)
- [7.4.2003 : still cursor flickering when a callback updates the screen](https://github.com/vim/vim/commit/144445d15f2369a58b94b6fee2fe162fc81c8abe)
- [7.4.2004 : GUI: cursor displayed in the wrong position](https://github.com/vim/vim/commit/65549bdef5805ba5a18a3d2d3909998dd7d9ed46)
- [7.4.2005 : after using evalcmd() message output is in the wrong position](https://github.com/vim/vim/commit/ee1deb4a00f39f133558321ec535354497f490c8)
- [7.4.2006 : crash when using tabnext in BufUnload autocmd](https://github.com/vim/vim/commit/30445cb6e94698d212ba866ef3e4022ac625540a)
- [7.4.2007 : running the tests leaves a viminfo file behind](https://github.com/vim/vim/commit/fc4ad616073a169badfb2b9906fee2844f76f730)
- [7.4.2008 : evalcmd() has a confusing name](https://github.com/vim/vim/commit/79815f1ec77406f2f21a618c053e5793b597db7a)
- [7.4.2009 : (after 7.4.2008) messages test fails](https://github.com/vim/vim/commit/ed59aa60d3905f935283727f4a7b33c81a00174b)
- [7.4.2010 : there is a :cbottom command but no :lbottom command](https://github.com/vim/vim/commit/537ef08408c50e0c4104d57f74993b3b0ed9560d)
- [7.4.2011 : it is not easy to get a list of command arguments](https://github.com/vim/vim/commit/aa4d73235bf4deee167aa5314b89ae3d3db334b7)
- [7.4.2012 : (after 7.4.2011) test for getcompletion() may fail](https://github.com/vim/vim/commit/0d3e24be5686c0710aa3c6671e4c626d6cb21a5f)
- [7.4.2013 : using "noinsert" in 'completeopt' breaks redo](https://github.com/vim/vim/commit/67081e50616ae9546621072c5eaaa59bd0a4bed7)
- [7.4.2014 : using "noinsert" in 'completeopt' does not insert match](https://github.com/vim/vim/commit/32b808a4bdf35b0dea63c735702a591e5869fecd)
- [7.4.2015 : when file gets a name when writing it 'acd' is not effective](https://github.com/vim/vim/commit/5c71994f4ee5f87d4cce990dbc9684c70b1e108b)
- [7.4.2016 : warning from MinGW about &#x5f;WIN32&#x5f;WINNT redefined](https://github.com/vim/vim/commit/2bc127f94016801250f8f24234f90a5182d77e73)
- [7.4.2017 : adding many errors to the quickfix list takes a long time](https://github.com/vim/vim/commit/8240433f48f7383c281ba2453cc55f10b8ec47d9)
- [7.4.2018 : buf&#x5f;valid() can be slow when there are many buffers](https://github.com/vim/vim/commit/b25f9a97e9aad3cbb4bc3fe87cdbd5700f8aa0c6)
- [7.4.2019 : when ignoring case utf&#x5f;fold() may consume a lot of time](https://github.com/vim/vim/commit/c4a927ca8dc383190d5df2cacd3f966698b6190c)
- [7.4.2020 : can't build without +autocmd feature](https://github.com/vim/vim/commit/453f37dbfd6f8304a36ea84e40a9965404206186)
- [7.4.2021 : still too many buf&#x5f;valid() calls](https://github.com/vim/vim/commit/19ff9bf454b7492be64dd87aaf0830fa7961871e)
- [7.4.2022 : warnings from 64 bit compiler](https://github.com/vim/vim/commit/25065ec375a8a55462f6c07c76dc1a72a770ac19)
- [7.4.2023 : buflist&#x5f;findname&#x5f;stat() may find a dummy buffer](https://github.com/vim/vim/commit/ea3f2e7be447a8f0c4436869620f908de5e8ef1e)
- [7.4.2024 : more buf&#x5f;valid() calls can be optimized](https://github.com/vim/vim/commit/7c0a2f367f2507669560b1a66423155c70d2e75b)
- [7.4.2025 : cursor blinking irregular when using channel or timer](https://github.com/vim/vim/commit/da3a77d9ec28407b8fa2aa014e76944d0a525662)
- [7.4.2026 : reference counting for callbacks isn't right](https://github.com/vim/vim/commit/1436d8d51cce114be56209924fc71376407e5bad)
- [7.4.2027 : can't build with +eval but without +menu](https://github.com/vim/vim/commit/4c06815c44dfeaafdad25dfcc40f60860096a900)
- [7.4.2028 : cppcheck warns for using index before limits check](https://github.com/vim/vim/commit/5498a41f5a62c3877fee0185adf3bf7245a9a547)
- [7.4.2029 : printf() does not work with 64 bit numbers](https://github.com/vim/vim/commit/38ee6b041e73ad31c3b6b99d56d20833b59b2b57)
- [7.4.2030 : ARCH must be set properly when using MinGW](https://github.com/vim/vim/commit/e048539195e1c0a060cf41b39fca118130bdc151)
- [7.4.2031 : list&#x5f;lbr&#x5f;utf8 test fails because of user setting](https://github.com/vim/vim/commit/89b10421ca1aea55acbafac1df93be4e20210ceb)
- [7.4.2032 : (after 7.4.2030) build fails with 64 bit MinGW](https://github.com/vim/vim/commit/3e946fdc7b3639f84c96cd8a5d5b4b4b010f56ea)
- [7.4.2033 : 'cscopequickfix' option does not accept new value "a"](https://github.com/vim/vim/commit/6d20e1754461b0f8d395f2e3464f0dc1060497f7)
- [7.4.2034 : build fails with some version of MinGW](https://github.com/vim/vim/commit/f62e797ffa6ab5d4c476b4c9e9a1f7f852910652)
- [7.4.2035 : on Solaris with ZFS the ACL may get removed](https://github.com/vim/vim/commit/da4127794aa333631bde879e73bcfce4aef42f85)
- [7.4.2036 : looking up a buffer by number is slow if there are many](https://github.com/vim/vim/commit/480778b805bd8bdc5d657560230e9c50feda1d0f)
- [7.4.2037 : (after 7.4.2036) small build fails](https://github.com/vim/vim/commit/7399895520b3248d5751c8dba8b9a4f27522d2dd)
- [7.4.2038 : (after 7.4.2036) small build still fails](https://github.com/vim/vim/commit/9280e3f95d065733f04fa22869e5ef071d531931)
- [7.4.2039 : the Netbeans integration is not tested](https://github.com/vim/vim/commit/321efdd77a7b9ac11ade90dd7634b5d37f4820fe)
- [7.4.2040 : new files missing from distribution](https://github.com/vim/vim/commit/a6aa78a3e3433db42b8ab644dab7cd327f6d9499)
- [7.4.2041 : netbeans file authentication not tested](https://github.com/vim/vim/commit/603d657219154ef074eeb70f6892b0b54e4fc19b)
- [7.4.2042 : GTK: display updating is not done properly and can be slow](https://github.com/vim/vim/commit/fdadad994a6e8f6cc8b11519082e23200b96d0ba)
- [7.4.2043 : setbufvar() causes a screen redraw](https://github.com/vim/vim/commit/93431df9eb02f7cf3d7f2142bb1bef24c5f325b2)
- [7.4.2044 : filter() and map() either require a string or a function](https://github.com/vim/vim/commit/069c1e7fa9f45a665064f7f2c17da84d6a48f544)
- [7.4.2045 : memory leak when using a function callback](https://github.com/vim/vim/commit/5ef2e762844e0eb1f5b213bd58e6bc38bdb69dac)
- [7.4.2046 : the qf&#x5f;init&#x5f;ext() function is too big](https://github.com/vim/vim/commit/e0d3797664c59afc9705808f86a7cf00fd6d874d)
- [7.4.2047 : compiler warning for initializing a struct](https://github.com/vim/vim/commit/bfafb4c4a01db3f8c508716daf689e0dfe92b649)
- [7.4.2048 : there is still code and help for unsupported systems](https://github.com/vim/vim/commit/a06ecab7a5159e744448ace731036f0dc5f87dd4)
- [7.4.2049 : there is no way to get a list of the error lists](https://github.com/vim/vim/commit/f6acffbe83e622542d9fdf3066f51933e46e4954)
- [7.4.2050 : when using ":vimgrep" may end up with duplicate buffers](https://github.com/vim/vim/commit/015102e91e978a0bb42a14461c132a85e8f7e1ea)
- [7.4.2051 : no proper testing of trunc&#x5f;string()](https://github.com/vim/vim/commit/502ae4ba63561c98ac69af26cd9883bfd18d225f)
- [7.4.2052 : coverage report is messed up by the unittests](https://github.com/vim/vim/commit/c3e539fa703b23d9a24c3520c96ec305ee932cb4)
- [7.4.2053 : can't run scripttests in the top directory](https://github.com/vim/vim/commit/e21d62435ef966b8adb40dee13ef9d6791316674)
- [7.4.2054 : (after 7.4.2048) wrong part of #ifdef removed](https://github.com/vim/vim/commit/b2b050ab16565c117f0e7e411ffef3700d99203b)
- [7.4.2055 : eval.c is too big](https://github.com/vim/vim/commit/cd52459c387785796713826c63174cdeed295dd4)
- [7.4.2056 : (after 7.4.2055) build fails](https://github.com/vim/vim/commit/1abc52c1599fede69b4e76ffde0e879d6fb5080a)
- [7.4.2057 : eval.c is too big](https://github.com/vim/vim/commit/da861d631d7e22654faee2789286c685ad548911)
- [7.4.2058 : eval.c is too big](https://github.com/vim/vim/commit/a9b579f3d7463720a316e11e77a7a9fbb9267986)
- [7.4.2059 : non-Unix builds fail](https://github.com/vim/vim/commit/6583c44857368f28c802dabe10ac7b7b0c266f50)
- [7.4.2060 : (after 7.4.2059) typo in build file](https://github.com/vim/vim/commit/a9093fe0946032b1bcaecaad82bfaf6763195aa4)
- [7.4.2061 : qf&#x5f;init&#x5f;ext() is too big](https://github.com/vim/vim/commit/e87e6dddc2b2a99572ec0db0833c052214c4fbd3)
- [7.4.2062 : using dummy variable to compute struct member offset](https://github.com/vim/vim/commit/840268400dc8fda62a14f8a084e8b1ea46619454)
- [7.4.2063 : eval.c is still too big](https://github.com/vim/vim/commit/73dad1e64cb42842d8259cb1a255a6fa59822f76)
- [7.4.2064 : Coverity warns for possible buffer overflow](https://github.com/vim/vim/commit/4f5c5f29806e53251b7a7b68ce7de86a21ff8015)
- [7.4.2065 : compiler warns for uninitialzed variable](https://github.com/vim/vim/commit/69aa099641616268bfcde9dc3aa313c677846b12)
- [7.4.2066 : getcompletion() not well tested](https://github.com/vim/vim/commit/c1fb763184c8ae82300357867fa2070aa94366e9)
- [7.4.2067 : compiler warning for char/char&#x5f;u conversion, inefficient code](https://github.com/vim/vim/commit/16ec3c9be3fcdc38530bddb12978bc5a7b98c0f6)
- [7.4.2068 : memory access error when running the message tests](https://github.com/vim/vim/commit/b9644433d2728e99fab874e5e33147ad95d23a31)
- [7.4.2069 : spell.c is too big](https://github.com/vim/vim/commit/9ccfebddc3ff2a3c2853cf706fd4c26f639bf381)
- [7.4.2070 : (after 7.4.2069) missing change to include file](https://github.com/vim/vim/commit/6cfdb2a3bad5a6049de22dcdd1da0f6666478398)
- [7.4.2071 : the return value of type() is difficult to use](https://github.com/vim/vim/commit/f562e72df726c6191fa305e1c0a113f1cfb87f76)
- [7.4.2072 : substitute() does not support a Funcref argument](https://github.com/vim/vim/commit/72ab729c3dcdea0fba44d8e676602c847e841bcd)
- [7.4.2073 : rgb.txt is read for every color name](https://github.com/vim/vim/commit/68015bbd846181d49842d6ef60246c4195d20b89)
- [7.4.2074 : one more place using a dummy variable](https://github.com/vim/vim/commit/0a0f641b9841189ba4180758109d04d0a26e50e3)
- [7.4.2075 : no autocommand event to initialize a window or tab page](https://github.com/vim/vim/commit/c917da4b3e8801a255dbefea8e4ed19c1c716dd8)
- [7.4.2076 : syntax error when dict has '>' key](https://github.com/vim/vim/commit/4f0383bc3fe5af0229fb66b53fe94329af783eff)
- [7.4.2077 : cannot update 'tabline' when a tab was closed](https://github.com/vim/vim/commit/12c11d553053f5a9eae9eb3c518279b12fa928c2)
- [7.4.2078 : running checks in po diretory fails](https://github.com/vim/vim/commit/ca8942c6e331a69ddd533dd78931f399f7dcaa79)
- [7.4.2079 : Netbeans test fails on non-Unix systems](https://github.com/vim/vim/commit/53bdec2ac9b994eee9c60f2ab9ad8a8afb2b9594)
- [7.4.2080 : PERROR() on some systems does not trigger assert&#x5f;fails()](https://github.com/vim/vim/commit/b869c0da31716ff14bbfd63346d140d0a1d68af7)
- [7.4.2081 : line numbers in the error list are not always adjusted](https://github.com/vim/vim/commit/c1542744e788d96fed24dd421f43009288092504)
- [7.4.2082 : not much test coverage for digraphs](https://github.com/vim/vim/commit/397eadbe25370394e0dc3cb18766533a15b8d1c0)
- [7.4.2083 : Coverity complains about not restoring a value](https://github.com/vim/vim/commit/19df5cc04de93519145901d9cbc78c44d0c47c5d)
- [7.4.2084 : (after 7.4.2081) new digraph test makes testing hang](https://github.com/vim/vim/commit/60084333816c585d5858bc085b2942f813102ae3)
- [7.4.2085 : digraph tests fails on some systems](https://github.com/vim/vim/commit/dfd63e30d13fff5603416b6c7e247cebeb003eb0)
- [7.4.2086 : using the system default encoding makes tests unpredictable](https://github.com/vim/vim/commit/ac105ed3c420660ddbddc501c97875c48220817e)
- [7.4.2087 : digraph code test coverage is still low](https://github.com/vim/vim/commit/e25bc5abb4621c6263b0248988a207d53b57fc54)
- [7.4.2088 : (after 7.4.2087) keymap test fails with normal features](https://github.com/vim/vim/commit/f36213597d737ab500771e87edcf121acd55e277)
- [7.4.2089 : color handling of X11 GUIs is too complicated](https://github.com/vim/vim/commit/36edf0685c8b55ee3ce709058d83ada8027fec1e)
- [7.4.2090 : using submatch() in lambda passed to substitute() is verbose](https://github.com/vim/vim/commit/df48fb456fb6bf63d94cad9b302ff01d8ee8d311)
- [7.4.2091 : Coverity reports a resource leak when out of memory](https://github.com/vim/vim/commit/2e45d21c819272051f7ef4436f483e4b2ecfb369)
- [7.4.2092 : GTK 3 build fails with older GTK version](https://github.com/vim/vim/commit/870b749ce1db1cec80f0f3766064a031688e6a46)
- [7.4.2093 : Netbeans test fails once in a while; leaving log file behind](https://github.com/vim/vim/commit/a58c58b7e1404bb91d59d14539b41390284e7026)
- [7.4.2094 : the color allocation in X11 is overly complicated](https://github.com/vim/vim/commit/4658228262f491fcb582d531d4e8e5754b0d5e83)
- [7.4.2095 : Man test fails when run with the GUI](https://github.com/vim/vim/commit/61c04493b00f85d0b97436260a9ef9ab82143b78)
- [7.4.2096 : lambda functions show up with completion](https://github.com/vim/vim/commit/b49edc11a1872fa99befa9a4a8ea6c8537868038)
- [7.4.2097 : warning from 64 bit compiler](https://github.com/vim/vim/commit/d4f31dc45482e8db527ca044b9c3436b1e750006)
- [7.4.2098 : text object tests are old style](https://github.com/vim/vim/commit/00b24be454800f544676aa8850fb4378a568901e)
- [7.4.2099 : when a keymap is active only "(lang)" is displayed](https://github.com/vim/vim/commit/73ac0c4281a3606651604a3cbcc334bfb3859a87)
- [7.4.2100 : "cgn" and "dgn" do not always work correctly](https://github.com/vim/vim/commit/6835dc61aebca2b602d85a9d63c449ace58683b4)
- [7.4.2101 : looping over windows, buffers and tab pages is inconsistant](https://github.com/vim/vim/commit/2932359000b2f918d5fade79ea4d124d5943cd07)
- [7.4.2102 : (after 7.4.2101) tiny build with GUI fails](https://github.com/vim/vim/commit/45a249513f5ed2de0a6f6b5cb6e32015e35644be)
- [7.4.2103 : can't have "augroup END" right after ":au!"](https://github.com/vim/vim/commit/e99e84497b89e5f91df519790802770920ecf4fe)
- [7.4.2104 : code duplication when unreferencing a function](https://github.com/vim/vim/commit/97baee80f0906ee2f651ee1215ec033e84f866ad)
- [7.4.2105 : configure reports "normal" default features while its "huge"](https://github.com/vim/vim/commit/eec2981bbee42411044800bc23731ebcc82b5b66)
- [7.4.2106 : clang warns about missing field in initializer](https://github.com/vim/vim/commit/ac77aec4daea8d73468fcf4690cb4ccab1d807ed)
- [7.4.2107 : (after 7.4.2106) misplaced equal sign](https://github.com/vim/vim/commit/d955a0971cb4207e466ad40e78138d71052e1758)
- [7.4.2108 : Netbeans test is flaky](https://github.com/vim/vim/commit/adfc5c2f9cf8032e496fcd4793d829768bd6d0a7)
- [7.4.2109 : 'display=lastline" is too drastic, without it not nice](https://github.com/vim/vim/commit/ad9c2a08f0509294269a2f11a59a438b944bdd5a)
- [7.4.2110 : with CmdUndefined errof for missing autocmd is E464](https://github.com/vim/vim/commit/eac784eced501c54d2c99e18a1af96cd996f3a6c)
- [7.4.2111 : defaults are very conservative](https://github.com/vim/vim/commit/8c08b5b569e2a9e9f63dea514591ecfa2d3bb392)
- [7.4.2112 : getcompletion(.., 'dir') returns trailing "&#x2a;" without matches](https://github.com/vim/vim/commit/b56195ed00a9a79aa6217cddbeedbc8cc7a5b6d8)
- [7.4.2113 : test for undo is flaky](https://github.com/vim/vim/commit/170b10b421f0c9fda08b7cfd3bb043c064f3659a)
- [7.4.2114 : (after 7.4.2113) tiny build fails](https://github.com/vim/vim/commit/d05b191b91c4e16d6887bf781832d135d2a8fae5)
- [7.4.2115 : loading defaults.vim with -C argument](https://github.com/vim/vim/commit/b9a46fec3e79d1fc8c406084a41733c647a5e535)
- [7.4.2116 : the default vimrc for Windows is very conservative](https://github.com/vim/vim/commit/c73e4474b1f1b5b18a8d504eec5305e0c77981f7)
- [7.4.2117 : no warning when deleting an augroup that still has autocmds](https://github.com/vim/vim/commit/f2c4c391192cab6e923b1a418d4af09106fba25f)
- [7.4.2118 : Mac: can't build with tiny features](https://github.com/vim/vim/commit/83a2a80d6f699ad9a236431170038698e355c025)
- [7.4.2119 : closures are not supported](https://github.com/vim/vim/commit/1e96d9bf98f9ab84d5af7f98d6a961d91b17364f)
- [7.4.2120 : user defined functions can't be a closure](https://github.com/vim/vim/commit/10ce39a0d52272a3dfff2feb8c631529f29e6740)
- [7.4.2121 : no easy way to check if lambda and closure are supported](https://github.com/vim/vim/commit/9532fe7fbe1b14531931e83bd9f8054efdcf7509)
- [7.4.2122 : (after 7.4.2118) Mac: don't get +clipboard in huge build](https://github.com/vim/vim/commit/b822cb0f93177bb045b221f607aee735f08ce428)
- [7.4.2123 : no new style test for diff mode](https://github.com/vim/vim/commit/42093c0ec52e6ff29e80aae65ac6a744c7de79bb)
- [7.4.2124 : diffmode test leaves files behind, breaking another test](https://github.com/vim/vim/commit/623cf88f9c5ad49cce8e846af29a1bb9346c7481)
- [7.4.2125 : compiler warning for loss of data](https://github.com/vim/vim/commit/a5c0cc11330157c721748e317e8ff54b649610ca)
- [7.4.2126 : no tests for :diffget and :diffput](https://github.com/vim/vim/commit/90d121fa3637b423169c64528efe84a1e67060c9)
- [7.4.2127 : short form of ":noswapfile" is ":noswap" instead of ":nos"](https://github.com/vim/vim/commit/3bcfca3ab4db415d0e750e00204dd25a91fcee77)
- [7.4.2128 : memory leak when saving for undo fails](https://github.com/vim/vim/commit/1e2258297bb31720bfbeb234f2dae4d1b3b04fbd)
- [7.4.2129 : memory leak when using timer&#x5f;start()](https://github.com/vim/vim/commit/3ab14355ed93958433eec58e96ec9bd9a472f2b2)
- [7.4.2130 : pending timers cause false memory leak reports](https://github.com/vim/vim/commit/623e263ffb998acacd6fd0de18c44e03af2a47d1)
- [7.4.2131 : more memory leaks when using partial, e.g. for "exit-cb"](https://github.com/vim/vim/commit/57e69ff2cca9edee3546373d04e0dde6810309e5)
- [7.4.2132 : test&#x5f;partial has memory leaks reported](https://github.com/vim/vim/commit/31440a1f2b6ae8a2f7ffdd4a8a3661eadc4df23c)
- [7.4.2133 : (after 7.4.2128) can't build with tiny features](https://github.com/vim/vim/commit/fc1f2015e833e69b89b994faf4b2c4fc1e09220f)

## 新着スクリプト

- [One Half Color Schemes : Clean and vibrant light/dark color schemes for terminal and gui Vim.](http://www.vim.org/scripts/script.php?script_id=5412)
- [vim-dispatch-vimproc : the plugin to add vimproc strategy to vim-dispatch](http://www.vim.org/scripts/script.php?script_id=5413)
- [vim-tf2 : TF2 .cfg syntax highlighting. https://github.com/au/vim-tf2](http://www.vim.org/scripts/script.php?script_id=5414)
- [vim-one : one dark and one light for atom adapted to Vim](http://www.vim.org/scripts/script.php?script_id=5415)
- [vim-two-firewatch : A blend between duotone light and firewatch for atom](http://www.vim.org/scripts/script.php?script_id=5416)
- [python-compilers.vim : Just a python compilers for VIM](http://www.vim.org/scripts/script.php?script_id=5417)
- [vimmake : Customize shell tools for vim (supports async jobs)](http://www.vim.org/scripts/script.php?script_id=5418)
- [gtm-plugin.vim : Simple, seamless, lightweight time tracking for all your git projects](http://www.vim.org/scripts/script.php?script_id=5419)
- [cpf&#x5f;upf&#x5f;sdc : Highlight files for CPF, UPF and SDC](http://www.vim.org/scripts/script.php?script_id=5420)
- [audionote.vim : Audio note - voice recorder ](http://www.vim.org/scripts/script.php?script_id=5421)
- [vseq.vim : Generate sequential number vertically](http://www.vim.org/scripts/script.php?script_id=5422)
- [ditto.vim : Highlight overused words](http://www.vim.org/scripts/script.php?script_id=5423)
- [dialect.vim : Project specific spell files](http://www.vim.org/scripts/script.php?script_id=5424)
- [highlightedyank : Make the yanked region apparent!](http://www.vim.org/scripts/script.php?script_id=5425)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (1492)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1290)
3. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (568)
4. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (557)
5. [winmanager : A windows style IDE for Vim 6.0](http://www.vim.org/scripts/script.php?script_id=95) (495)
6. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (477)
7. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (436)
8. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (405)
9. [bufexplorer.zip : Buffer Explorer / Browser](http://www.vim.org/scripts/script.php?script_id=42) (401)
10. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (391)

## vim-jp issues

Open : 239 (-4) | Closed : 699 (+20)

- [Issue #923 : タブページ関連のコマンドと自動コマンドの組み合わせでVimが落ちたりカレントウィンドウが閉じられない](https://github.com/vim-jp/issues/issues/923)
- [Issue #924 : filter() と map() を Funcref に対応させる](https://github.com/vim-jp/issues/issues/924)
- [Issue #925 : filter(), map() の第2引数が number だと SEGV する ](https://github.com/vim-jp/issues/issues/925)
- [Issue #926 : 異なるencodingを使用すると不正な&#x5f;viminfoが生成されて落ちる事がある](https://github.com/vim-jp/issues/issues/926)
- [Issue #927 : サポートを打切ったOS関連のキーワードをCソース、ドキュメント、同梱Vim scriptから削除する](https://github.com/vim-jp/issues/issues/927)
- [Issue #928 : printf() の %d が 64bit 整数に対応していない](https://github.com/vim-jp/issues/issues/928)
- [Issue #929 : printf() に2進数表示できる変換指示子 %b が欲しい](https://github.com/vim-jp/issues/issues/929)
- [Issue #930 : 起動して表示前に confirm() されると詰む](https://github.com/vim-jp/issues/issues/930)
- [Issue #931 : vim.exe 起動時の confirm() で文字化け](https://github.com/vim-jp/issues/issues/931)
- [Issue #932 : MinGWでのビルド時に&#x60;ARCH&#x60;を自動設定して欲しい](https://github.com/vim-jp/issues/issues/932)
- [Issue #933 : 自動コマンドからタブ操作とバッファ削除するとSEGVする](https://github.com/vim-jp/issues/issues/933)
- [Issue #934 : テストをcleanしたあとにファイルが残る](https://github.com/vim-jp/issues/issues/934)
- [Issue #935 : '>' を dict の key にすると E125 エラーになる](https://github.com/vim-jp/issues/issues/935)
- [Issue #936 : printf() に変換指示子 %v を追加する](https://github.com/vim-jp/issues/issues/936)
- [Issue #937 : スタックトレース情報を持つ変数を追加したい](https://github.com/vim-jp/issues/issues/937)
- [Issue #938 : php で一部のネストしたインデントが正しく行なわれない](https://github.com/vim-jp/issues/issues/938)
