---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2017 年 11 月号

---

## 話題

*   [VimConf 2017](http://vim-jp.org/blog/2017/08/04/vimconf2017-venue-and-date-ja.html) が 11/04 (土) に開催されました
*   [Fablic.vim #4](https://fablicvim.connpass.com/event/71673/) が 12/04 (月) に東京で開催予定です
*   [Meguro.vim #7](https://megurovim.connpass.com/event/72311/) が 12/16 (土) に東京で開催予定です
*   [Yokohama.vim #10](https://yokohamavim.connpass.com/event/73077/) が 12/10 (日) に神奈川で開催予定です

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

*   8.0.1278: GUI ウィンドウのスクロールバー表示時にウィンドウのサイズを保持する `'guioptions'` オプションの `k` フラグが追加されました
    *   これによりスクロールバーが表示される時に GUI のウィンドウの最大化状態を解除されてしまう問題が修正されました
*   8.0.1279: メニューの初期化処理が高速化され、起動時間が速くなりました
*   8.0.1280: if\_python における `None` が Vim script の `v:none` に変換されるようになりました
*   8.0.1281: Vim script で実行されるファイルタイプの検知処理が高速化され、起動時間が速くなりました (関連パッチ: 8.0.1282, 8.0.1284)
*   8.0.1287: ウィンドウローカルのディレクトリを保持する `'viewoptions'` オプションの `curdir` が追加されました
*   8.0.1295: UNIX 系の端末版 Vim でも `v:servername` が常に生成される `--enable-autoservername` オプションが configure スクリプトに追加されました (関連パッチ: 8.0.1297)
*   8.0.1299: Bracketed paste mode が端末ウィンドウで動くようになりました
*   8.0.1305: `writefile()` の第3引数のフラグに `fsync(2)` を実行するか否かのフラグ `s`, `S` が追加されました (関連パッチ: 8.0.1335)
    *   `s` が指定された場合は実行し、`S` が指定された場合は実行せず、どちらも指定されなかった場合は `'fsync'` オプションに従います
    *   今までは常に `fsync(2)` を実行していませんでした
*   8.0.1309: `'balloonevalterm'` オプションが追加され、Unix 系の GUI 環境から起動された端末版 Vim でも ballooneval 機能が使用できるようになりました (関連パッチ: 8.0.1323, 8.0.1324)
*   8.0.1318: `balloon_split()` 関数が追加され、`balloon_show()` 関数が端末版 Vim では List を受け付けるようになりました (関連パッチ: 8.0.1346)
*   8.0.1330: Vim の端末機能から実行されたジョブ内で Windows の場合も `VIM_SERVERNAME` 環境変数が定義されるようになりました
*   8.0.1332: `qfSeparator` ハイライトグループが追加されました
*   8.0.1336: xim 機能付きでコンパイル済みでなくとも `'imactivatefunc'` オプションが使えるようになりました (関連パッチ: 8.0.1338)
*   8.0.1343: Windows で色付き絵文字を表示できるようになりました

## Vimに関する脆弱性

*   スワップファイルのパーミッションを問題とする脆弱性 [CVE-2017-1000382][1000382] ([JVNDB-2017-009929 ][009929_AD_1]) が指摘されました。

    基本的にはWebサーバーなどの設定および使い方の問題であると考えられています。
    しかしスワップファイルが編集中のファイルと同じディレクトリに作られるのを
    避けたいのであれば、 `'dir'` オプションで別のディレクトリに作るように
    設定すると良いでしょう。設定例:

    ```vimscript
    silent! call mkdir($HOME . '/.vim/swap', 'p', 0700)
    set directory=~/.vim/swap//
    ```

*   その他、軽微な脆弱性が 8.0.1345 までで 3点修正されています。 [参考資料][bramcomment]

## リリース情報

- [8.0.1241 : popup test is flaky](https://github.com/vim/vim/commit/89c394faca40d2f5d57705432a433173b295bf73)
- [8.0.1242 : function argument with only dash is seen as number zero](https://github.com/vim/vim/commit/ffd99f729bd806e09d9355ede9c17780b61057bf)
- [8.0.1243 : no test for what 8.0.1227 fixes](https://github.com/vim/vim/commit/f45938cc20ed6992e5215ffe41b73b528c78be9c)
- [8.0.1244 : search test does not work correctly on MS-Windows](https://github.com/vim/vim/commit/b94340c04f524482b408543daa1142f883831cab)
- [8.0.1245 : when WaitFor() has a wrong expression it just waits a second](https://github.com/vim/vim/commit/3e1c617d49224038a820ee2c552ebe31a84aaa89)
- [8.0.1246 : popup test has an arbitrary delay](https://github.com/vim/vim/commit/b315876efa7865486b9cc160d43f0ead47e58d6c)
- [8.0.1247 : not easy to find Debian build info](https://github.com/vim/vim/commit/86b21bb3e7bc687fc74e127d78a89b411929716b)
- [8.0.1248 : stray + in README file](https://github.com/vim/vim/commit/5130f316612dfc71c433c2743340dc45444eb17d)
- [8.0.1249 : no error when WaitFor() gets an invalid wrong expression](https://github.com/vim/vim/commit/c20e0d52071a3f6e12321ec3344024faa4695da9)
- [8.0.1250 : 'hlsearch' highlighting not removed after incsearch](https://github.com/vim/vim/commit/f8f8b2eadbaf3090fcfccbab560de5dbd501833d)
- [8.0.1251 : invalid expressin passed to WaitFor()](https://github.com/vim/vim/commit/d97fbf171ec0e63117813da045d2a1c51a9b6f62)
- [8.0.1252 : incomplete translations makefile for MinGW/Cygwin](https://github.com/vim/vim/commit/8889a5c305e69aa49fd08036e624e365097a5b7b)
- [8.0.1253 : still too many old style tests](https://github.com/vim/vim/commit/430dc5d360166ca5bb6a73f2c87ae53e09282ecb)
- [8.0.1254 : undefined left shift in gethexchrs()](https://github.com/vim/vim/commit/4c22a91d20cce4f28dd2852a13129b5a4cc691da)
- [8.0.1255 : duplicate badge README file](https://github.com/vim/vim/commit/ea84df8041dbbff95acb1db8532281679c5fbe5a)
- [8.0.1256 : typo in configure variable vim&#x5f;cv&#x5f;tgent](https://github.com/vim/vim/commit/a88254f70458e7d4c998de7cc41aa06176eb61dc)
- [8.0.1257 : no test for fix of undefined behavior](https://github.com/vim/vim/commit/2973daafe1732963b8924cb9df53c608804d66b3)
- [8.0.1258 : 'ttymouse' is set to "sgr" even though it's not supported](https://github.com/vim/vim/commit/52a2f0f1da4e554a81beb45211a9d09afffde595)
- [8.0.1259 : search test can be flaky](https://github.com/vim/vim/commit/13deab8d08145c1f6e2a3e82cb547bc7f87a3686)
- [8.0.1260 : using global variables for WaitFor()](https://github.com/vim/vim/commit/ab8b1c14a31e36ae87cc7e13c4a75318d513fc7b)
- [8.0.1261 : program in terminal window gets NL instead of CR](https://github.com/vim/vim/commit/1232624ae5b56c167c5982a2620f736c4bbc19ef)
- [8.0.1262 : terminal redir test is flaky](https://github.com/vim/vim/commit/7dd88c5133feda7e9ccfedcb38b6dfdba459d507)
- [8.0.1263 : others can read the swap file if a user is careless](https://github.com/vim/vim/commit/5a73e0ca54c77e067c3b12ea6f35e3e8681e8cf8)
- [8.0.1264 : terminal debugger gets stuck in small window](https://github.com/vim/vim/commit/c3632516303842244442f354734e54a403ec4b50)
- [8.0.1265 : swap test not skipped when there is one group](https://github.com/vim/vim/commit/ad7dac85c3c90893e78e5463ca44b874082b482f)
- [8.0.1266 : Test&#x5f;swap&#x5f;directory was commented out](https://github.com/vim/vim/commit/ffe010fa0363d1a04b15cc6af119a4af63ba9363)
- [8.0.1267 : Test&#x5f;swap&#x5f;group may leave file behind](https://github.com/vim/vim/commit/5842a748be8039fd6d267f5557fe391c6c95399d)
- [8.0.1268 : PC install instructions are incomplete](https://github.com/vim/vim/commit/3bf8c3c38ff0d24047e2f2979c07a349fb89e557)
- [8.0.1269 : effect of autocommands on marks is not tested](https://github.com/vim/vim/commit/aace21581345aa51c09d809ab3744a943a71c879)
- [8.0.1270 : mismatching file name with Filelist](https://github.com/vim/vim/commit/8fdb35a974d92c7d0267bc8988aa6af74ee6885d)
- [8.0.1271 : still too many old style tests](https://github.com/vim/vim/commit/fb094e14c19337de824d4e6710ca6a2617930ab0)
- [8.0.1272 : warnings for unused variables in tiny build](https://github.com/vim/vim/commit/53ec79537a062b453ed72a4efcb4b1cd284fe9a6)
- [8.0.1273 : old test file remaining](https://github.com/vim/vim/commit/4148be468f82eb4bf463be3610b00a132aee4186)
- [8.0.1274 : setbufline() fails when using folding](https://github.com/vim/vim/commit/0c4dc88a637a5027209aa00226996af84e248636)
- [8.0.1275 : CmdlineLeave autocmd prevents fold from opening](https://github.com/vim/vim/commit/c9e9c7140994154bb4a29764491d42c001b24596)
- [8.0.1276 : key lost when window closed in exit callback](https://github.com/vim/vim/commit/a3f7e58bf846b28b347a5d5b96c2d829871bd212)
- [8.0.1277 : terminal window CR-NL conversions may cause problems](https://github.com/vim/vim/commit/26d205dcd886b48713f22cbdbf2a8e55400083dc)
- [8.0.1278 : GUI window always resizes when adding scrollbar](https://github.com/vim/vim/commit/8ac441576fd219fb7227349e228d5b68520b204a)
- [8.0.1279 : initializing menus can be slow](https://github.com/vim/vim/commit/040c1feb212568cf04265e59b26182f8b76057aa)
- [8.0.1280 : Python None cannot be converted to a Vim type](https://github.com/vim/vim/commit/de323093e1be165a3eadd1ab3d02942358cd97bf)
- [8.0.1281 : loading file type detection slows down startup](https://github.com/vim/vim/commit/851ee6c3da5fd726d92e1e3300d7e5e2e8b907c5)
- [8.0.1282 : script-local variable defined in the wrong script](https://github.com/vim/vim/commit/cef7322d8a902b4655ed861489c4e798672074f0)
- [8.0.1283 : test 86 fails under ASAN](https://github.com/vim/vim/commit/9ed7d34af303190bb222624157e0c6cd78b7d120)
- [8.0.1284 : loading file type detection slows down startup](https://github.com/vim/vim/commit/462455ee8b81cb5709007d91248ac4a88308d6e9)
- [8.0.1285 : distributed autoload files may clash with user files](https://github.com/vim/vim/commit/d09a206ee94ccb653707ce9b6e536d4d58886e04)
- [8.0.1286 : occasional crash when using a channel](https://github.com/vim/vim/commit/5fd8b78214a52b561780eb5ba000b1a3f5ba1d3d)
- [8.0.1287 : temp file used for viminfo may have wrong permissions](https://github.com/vim/vim/commit/aeeb6888ca3ee5122793b37c5aacc6c960b02c37)
- [8.0.1288 : GUI: cannot drag the statusline of a terminal window](https://github.com/vim/vim/commit/ae147ab2d7649e7a0bcf11b9f3db1900c1a4d594)
- [8.0.1289 : mkview always includes the local directory](https://github.com/vim/vim/commit/13e904199c2af6aac9b289a7f520d8e16f6e478a)
- [8.0.1290 : seq&#x5f;cur of undotree() wrong after undo](https://github.com/vim/vim/commit/80eaddd3a0bc47cb14168964678420cfe03a2502)
- [8.0.1291 : C indent wrong when &#x2a; immediately follows comment](https://github.com/vim/vim/commit/f8c53d3d268fc67a29c8c1a4e76fae85762e11b5)
- [8.0.1292 : quick clicks in the WinBar start Visual mode](https://github.com/vim/vim/commit/d327b0c68fa9b32877123f77560f4521a9eb39db)
- [8.0.1293 : setting a breakpoint in the terminal debugger sometimes fails](https://github.com/vim/vim/commit/60e73f2acc890f07056fa1bea8dbe89d26918b3e)
- [8.0.1294 : GUI: get stuck when splitting a terminal window](https://github.com/vim/vim/commit/2f27aab8e68c48a1b8a51715e1009d572f8e13b5)
- [8.0.1295 : cannot automatically get a server name in a terminal](https://github.com/vim/vim/commit/e42a6d250907e278707753d7d1ba91ffc2471db0)
- [8.0.1296 : checking the same condition twice](https://github.com/vim/vim/commit/1dcada1933acdab93ac423951b5958125343833d)
- [8.0.1297 : +autoservername does not show enabled on MS-Windows](https://github.com/vim/vim/commit/23921432369e210029a26864cfb4f1537b3c3395)
- [8.0.1298 : missing test file](https://github.com/vim/vim/commit/58a8f1706faa9e7a17f0efdd0b9030d4c0e502b0)
- [8.0.1299 : bracketed paste does not work well in terminal window](https://github.com/vim/vim/commit/a42ad57e481d6ed9c34c80c99bd0512d4660515b)
- [8.0.1300 : file permissions may end up wrong when writing](https://github.com/vim/vim/commit/cd142e3369db8888163a511dbe9907bcd138829c)
- [8.0.1301 : generated license file for NSIS has a modeline](https://github.com/vim/vim/commit/7cb769a69f5b9d6951f4427c85e75f5b2909cf10)
- [8.0.1302 : still too many old style tests](https://github.com/vim/vim/commit/209d3874c19d73b209e7d223cf05ea22e183f408)
- [8.0.1303 : 'ttymouse' is not set to "sgr" for Terminal.app and Iterm2](https://github.com/vim/vim/commit/9c6ce0e62297294b41344e684429a91853bc2268)
- [8.0.1304 : CTRL-G/CTRL-T don't work with incsearch and empty pattern](https://github.com/vim/vim/commit/d0480097177369a6ed91d47aba189ae647afcd68)
- [8.0.1305 : writefile() never calls fsync()](https://github.com/vim/vim/commit/7567d0b115e332f61a9f390aaccdf7825b891227)
- [8.0.1306 : ASAN error stack trace is not useful](https://github.com/vim/vim/commit/9e27217c489c6eca2bd98d1c9fa89c1b652ed792)
- [8.0.1307 : compiler warning for ignoring return value](https://github.com/vim/vim/commit/ae1e108caaa9ce2ca03076f8c4c10275f0c66b9c)
- [8.0.1308 : the "Reading from stdin" message may be undesired](https://github.com/vim/vim/commit/234d16286a2733adedef56784c17415ae169b9ad)
- [8.0.1309 : cannot use 'balloonexpr' in a terminal](https://github.com/vim/vim/commit/51b0f3701ecb440aa72ab6017c1df6940c0e0f6f)
- [8.0.1310 : cproto generates errors because of missing type](https://github.com/vim/vim/commit/0f1e643138d47bfc94a7050c5bd25493c2153960)
- [8.0.1311 : no test for strpart()](https://github.com/vim/vim/commit/c7d16dce2f180c8ebfc8105ad090b0ea2deedcdc)
- [8.0.1312 : balloon&#x5f;show() only works in terminal when compiled with GUI](https://github.com/vim/vim/commit/c3719bd87beca9f72d2e9f11e36d561c2c3b57b0)
- [8.0.1313 : missing dependencies cause parallel make to fail](https://github.com/vim/vim/commit/d1c28346e194ecf2303fe744052fb7cedc28fcfa)
- [8.0.1314 : build fails on Mac](https://github.com/vim/vim/commit/4ab9d9e9a41f441d7d96f51fec2213d412ea64e7)
- [8.0.1315 : build still fails on Mac](https://github.com/vim/vim/commit/e86ee877c1497890e64190c148ff59e8aaf10936)
- [8.0.1316 : build still still fails on Mac](https://github.com/vim/vim/commit/44c2bffde74a98acdc60505e57c70cd7cc2afa1e)
- [8.0.1317 : accessing freed memory in term&#x5f;wait()](https://github.com/vim/vim/commit/e518226713784e628ae7ee077f1b66cb12b9ffd9)
- [8.0.1318 : terminal balloon only shows one line](https://github.com/vim/vim/commit/246fe03d154c09070d5b7365b7f61716c4e0ddd4)
- [8.0.1319 : can't build GUI on MS-Windows](https://github.com/vim/vim/commit/669a828cdc1f019e7abda0e5e9b223d7897f06a8)
- [8.0.1320 : popup test fails on GUI-only build](https://github.com/vim/vim/commit/7221fce8b395972b57aea5f36ba2844fe1d4c259)
- [8.0.1321 : can't build huge version with Athena](https://github.com/vim/vim/commit/40e280d949fa26e33f69c4bf80ccc0b70f2e13f6)
- [8.0.1322 : textformat test isn't run](https://github.com/vim/vim/commit/5bbef3194984f175eb12b6efaa04c49a42ab58f2)
- [8.0.1323 : mouse events in a terminal window may cause endless loop](https://github.com/vim/vim/commit/73675fbc4810470f8935f05a3c16c14e0d5e737f)
- [8.0.1324 : some xterm sends different mouse move codes](https://github.com/vim/vim/commit/bb160a188abb3c40258d3e5726524ce203333e84)
- [8.0.1325 : more tests are not run](https://github.com/vim/vim/commit/5df95ea9ef34b5a898141ddc7134e4a7de713ba5)
- [8.0.1326 : largefile test fails on CI, glob test on MS-Windows](https://github.com/vim/vim/commit/6e77df2d8555ade4470e566011603ae40f1f0f3a)
- [8.0.1327 : new proto file missing from distribution](https://github.com/vim/vim/commit/91ffda9852916004c9f318cd01b2f90751329bad)
- [8.0.1328 : trouble when using ":term ++close" with autocmd](https://github.com/vim/vim/commit/ff5467965e3871d3dc0288416fcc6b1e2ba4f822)
- [8.0.1329 : when a flaky test fails it also often fails the second time](https://github.com/vim/vim/commit/550586015232ecc4f68b3479fa8ba2cf81c76e7b)
- [8.0.1330 : MS-Windows: job in terminal can't get back to Vim](https://github.com/vim/vim/commit/52dbb5ea7fde4a77178bc59e2383ca40df503812)
- [8.0.1331 : possible crash when window can be zero lines high](https://github.com/vim/vim/commit/3679c17917d7ff22e836982c81e5816bd07451dd)
- [8.0.1332 : highlighting in quickfix window could be better](https://github.com/vim/vim/commit/93a32e2ec4b5a391c5006ca09f196e6510c8c2e9)
- [8.0.1333 : some tests are run twice](https://github.com/vim/vim/commit/2c997d76034d5294c92305108357daef6feabd95)
- [8.0.1334 : splitting a window with a WinBar damages window layout](https://github.com/vim/vim/commit/3167c3e7010ac7e2fd7d3aa198dd2d1dcff8ecc8)
- [8.0.1335 : writefile() using fsync() may give an error.](https://github.com/vim/vim/commit/291a9d15ed1eb1094edc8ad6dda00a6da3bd7072)
- [8.0.1336 : cannot use imactivatefunc() unless compiled with +xim](https://github.com/vim/vim/commit/6315a9ae924d2104b61d1667ea71af3ac7cf48fa)
- [8.0.1337 : typo in #ifdef](https://github.com/vim/vim/commit/50d43153a7e3e452cd166cde0d84ecacd316321f)
- [8.0.1338 : USE&#x5f;IM&#x5f;CONTROL is confusing and incomplete](https://github.com/vim/vim/commit/819edbe078c8579f3620d12dac830f12ccdc5a45)
- [8.0.1339 : no test for what 8.0.1335 fixes](https://github.com/vim/vim/commit/83799a7b7414048df4ff4e507293416e8438e225)
- [8.0.1340 : MS-Windows: cannot build GUI without IME](https://github.com/vim/vim/commit/281c93e7140ba87552c4edcbfdaaa5e25dc6a97a)
- [8.0.1341 : 'imactivatefunc' test fails on MS-Windows](https://github.com/vim/vim/commit/be5d998d0ea2226f1f7180c32b19df17ad674853)
- [8.0.1342 : cannot build with Motif and multi-byte](https://github.com/vim/vim/commit/fb1db0e3550355e10ac86c4a4af21209c5456a66)
- [8.0.1343 : MS-Windows: does not show colored emojis](https://github.com/vim/vim/commit/d7ccc4d81dbcfa3ac0352bacea6e294fc9e33fda)
- [8.0.1344 : using 'imactivatefunc' in the GUI does not work](https://github.com/vim/vim/commit/2877d334ad1321d1fcd5f903c0493bd0cdd787f8)
- [8.0.1345 : race condition between stat() and open() for viminfo](https://github.com/vim/vim/commit/c41838aa01ef99540e2737c42e9b1283e3da5e26)
- [8.0.1346 : crash when passing 50 char string to balloon&#x5f;split()](https://github.com/vim/vim/commit/a3571ebef5a2969d16e3e202ec4bca8858aca6ba)
- [8.0.1347 : MS-Windows: build broken by misplaced curly](https://github.com/vim/vim/commit/e0aa23f7e3cc389251ecb3266926ac7d83056b1b)
- [8.0.1348 : make testclean deletes script file on MS-Windows](https://github.com/vim/vim/commit/c8c75796a60e89210cc470f1b3864546fd4f319a)
- [8.0.1349 : options test fails when using Motif or GTK GUI](https://github.com/vim/vim/commit/17471e84a72b5137a6010a98c4adb780237750db)
- [8.0.1350 : cannot build with +eval and -multi&#x5f;byte](https://github.com/vim/vim/commit/1355aad2b94790217aeef9077d3fb9925461f137)
- [8.0.1351 : warning for unused variables building with MinGW](https://github.com/vim/vim/commit/bdb8139098d170ede2bc79dd4f62e4ed5e778d3e)
- [8.0.1352 : dead URLs in the help go unnoticed](https://github.com/vim/vim/commit/f38c86eb6b665d9fcd3d232820a48198eccac376)
- [8.0.1353 : QuickFixCmdPost is not used consistently](https://github.com/vim/vim/commit/1ed2276fd50f34e824eeae93d5a5ebfdf118be26)
- [8.0.1354 : Shift-Insert doesn't always work in MS-Windows console](https://github.com/vim/vim/commit/3660a10c73a6d716ae9fca21f7c5a3282ddc7c85)
- [8.0.1355 : cursor keys don't work in MS-Windows console](https://github.com/vim/vim/commit/a45ff6caba706191b4fe328cab4c742ce17fa779)
- [8.0.1356 : using simalt in a GUIEnter autocommand inserts characters](https://github.com/vim/vim/commit/c5aa55db7e5bc791f99fb15b0f4be0d5dd166f62)
- [8.0.1357 : startup test fails on OpenBSD](https://github.com/vim/vim/commit/1a2ab991a6d168699fbcc934b63cfc3579c3d34d)
- [8.0.1358 : undercurl is not used in the terminal](https://github.com/vim/vim/commit/8b9e20afb0da79a21155988a38704b5777cc2efe)
- [8.0.1359 : libvterm ANSI colors can not always be recognized](https://github.com/vim/vim/commit/46359e198f6d6884dc3d3c4a3e46625f1b2a2ad2)

## 新着スクリプト

- [files-complete : completefunc which looks to all files in the current directory](https://vim.sourceforge.io/scripts/script.php?script_id=5617)
- [AutoSave : Automatically persist a buffer frequently.](https://vim.sourceforge.io/scripts/script.php?script_id=5618)
- [SearchAsQuickJump : Quick search without affecting 'hlsearch', search pattern and history.](https://vim.sourceforge.io/scripts/script.php?script_id=5619)
- [mkdx.vim : Adds useful mappings for working with markdown files](https://vim.sourceforge.io/scripts/script.php?script_id=5620)
- [vim-lessmess : Removes trailing white-space, fix mixed-indent and remove empty lines at the EOF](https://vim.sourceforge.io/scripts/script.php?script_id=5622)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](https://vim.sourceforge.io/scripts/script.php?script_id=273) (1258)
2. [wombat256.vim : Wombat for 256 color xterms](https://vim.sourceforge.io/scripts/script.php?script_id=2465) (1068)
3. [The NERD tree : A tree explorer plugin for navigating the filesystem](https://vim.sourceforge.io/scripts/script.php?script_id=1658) (895)
4. [AutomaticLaTeXPlugin : Background compilation, completion, bib serch, toc and other nice features.](https://vim.sourceforge.io/scripts/script.php?script_id=2945) (620)
5. [nginx.vim : initial version](https://vim.sourceforge.io/scripts/script.php?script_id=1886) (458)
6. [molokai : A port of the monokai scheme for TextMate](https://vim.sourceforge.io/scripts/script.php?script_id=2340) (429)
7. [AutoComplPop : Automatically opens popup menu for completions](https://vim.sourceforge.io/scripts/script.php?script_id=1879) (391)
8. [OmniCppComplete : C/C++ omni-completion with ctags database](https://vim.sourceforge.io/scripts/script.php?script_id=1520) (381)
9. [python.vim : Enhanced version of the python syntax highlighting script](https://vim.sourceforge.io/scripts/script.php?script_id=790) (379)
10. [c.vim : C/C++ IDE -- Write and run programs. Insert statements, idioms, comments etc.](https://vim.sourceforge.io/scripts/script.php?script_id=213) (347)

## vim-jp issues

Open : 254 (-4) | Closed : 871 (+18)

- [Issue #1112 : 検索コマンドラインが空の状態での Ctrl-G/T の挙動があやしい](https://github.com/vim-jp/issues/issues/1112)
- [Issue #1113 : undotree() の返すツリー上の現在位置がおかしい場合がある](https://github.com/vim-jp/issues/issues/1113)
- [Issue #1114 : macvimの:terminalでファイル名の濁点の位置がおかしい](https://github.com/vim-jp/issues/issues/1114)
- [Issue #1115 : 位置表現に offset を指定できるようにしたい](https://github.com/vim-jp/issues/issues/1115)
- [Issue #1116 : 端末機能に対する要望](https://github.com/vim-jp/issues/issues/1116)
- [Issue #1117 : 端末機能で起動したzshにクリップボードから貼り付けると~が入ったり最後の文字が大文字になったりする](https://github.com/vim-jp/issues/issues/1117)
- [Issue #1118 : XIM/IME が使えない場合でも iminsert=2 は設定できるべき。](https://github.com/vim-jp/issues/issues/1118)
- [Issue #1119 : Windows の :terminal で環境変数 VIM&#x5f;SERVERNAME が設定されない。](https://github.com/vim-jp/issues/issues/1119)
- [Issue #1120 : autoload ファイルを手動で source するとautoload関数定義時に2重で追加のsourceが発生する](https://github.com/vim-jp/issues/issues/1120)
- [Issue #1121 : writefile(\[...\], '/dev/stderr) すると E667: Fsync failed の例外が飛ぶ](https://github.com/vim-jp/issues/issues/1121)
- [Issue #1122 : WindowsのVimにおいて、DLL関数の呼び出し規約はどうなっているのか](https://github.com/vim-jp/issues/issues/1122)
- [Issue #1123 : vim.exe で Shift-Insert でペーストできない](https://github.com/vim-jp/issues/issues/1123)
- [Issue #1124 : プロファイルログに invalid byte sequence が含まれることがある](https://github.com/vim-jp/issues/issues/1124)
- [Issue #1125 : ch&#x5f;readraw と NL モードの問題提起](https://github.com/vim-jp/issues/issues/1125)

[1000382]:https://oss.sios.com/security/editors-security-vulnerabiltiy-20171101
[009929_AD_1]:http://vrda.jpcert.or.jp/feed/ja/JVNiPedia_JVNDB-2017-009929_AD_1.html
[bramcomment]:http://www.openwall.com/lists/oss-security/2017/11/28/4
