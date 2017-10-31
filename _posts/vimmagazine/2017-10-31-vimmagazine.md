---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2017 年 10 月号

---

## 話題

*   [Meguro.vim #6](https://megurovim.connpass.com/event/67608/) が 10/21 (土) に開催されました
*   [VimConf 2017](http://vim-jp.org/blog/2017/08/04/vimconf2017-venue-and-date-ja.html) が 11/04 (土) に開催予定です

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

*   8.0.1182: 'mzschemedll' オプションと 'mzschemegcdll' オプションが追加されました (関連パッチ: 8.0.1213)
*   8.0.1192: Windows で端末機能がデフォルトで有効になりました
*   8.0.1201: `yL` 等が 'scrolloff` の影響を受けないようになりました
*   8.0.1206: コマンドラインに入る時・抜ける時に実行される `CmdlineEnter` イベントと `CmdlineLeave` イベントが追加されました
*   8.0.1211: GTK と Windows でタブページをドラッグドロップで並べ替えられるようになりました
*   8.0.1232: Windows 版のインストーラで、Windows ライクなキーマッピングを行うかどうかの選択肢が出るようになりました
*   8.0.1234: Win32 GUI 版で合成文字や異体字が正しく表示されるようになりました
*   8.0.1236: `has('osxdarwin')` で `--enable-darwin` でコンパイルされたのかチェックできるようになりました
*   8.0.1238: 'incsearch' と 'hlsearch' がオンの場合、インクリメンタルサーチ中にも全ての候補をハイライトするようになりました
*   8.0.1239: `searchpair()` 関数の skip パラメーターに lambda を渡せるようになりました
*   8.0.1240: Windows において `term_start()` に環境変数が渡されるようになりました

## Vimに関する脆弱性

特筆すべき脆弱性の報告はありませんでした。

## リリース情報

- [8.0.1169 : highlignting one char too many with 'list' and 'cul'](https://github.com/vim/vim/commit/5ece3e359c4b96d0f9bf1ce690223545f40f34b8)
- [8.0.1170 : using termdebug results in 100% CPU time](https://github.com/vim/vim/commit/f33606112a59de6965692840b47d9cb4b924fc97)
- [8.0.1171 : popup test is still a bit flaky](https://github.com/vim/vim/commit/712549e04eddd6687c4b7654ec9af6da6c365603)
- [8.0.1172 : when E734 is given option is still set](https://github.com/vim/vim/commit/2a6a6c3014e728cd01c750b0f60484d4eaf22a8c)
- [8.0.1173 : terminal window is not redrawn after CTRL-L](https://github.com/vim/vim/commit/19a3d68b2cd4beb1ee1d97a84ad4e860ffe0cbb2)
- [8.0.1174 : Mac Terminal.app has wrong color for white](https://github.com/vim/vim/commit/a0a6f2776eabfbd4d9174fabdf3921beb7348eda)
- [8.0.1175 : build failure without +termresponse](https://github.com/vim/vim/commit/c902609f69b5662484f78dbf509715940d4134e4)
- [8.0.1176 : job&#x5f;start() does not handle quote and backslash correctly](https://github.com/vim/vim/commit/d78f03f86045184dfd191f00359baa61e2e79d1f)
- [8.0.1177 : in a terminal window the popup menu is not cleared](https://github.com/vim/vim/commit/54e5dbf58e16d0d757a315480c9e63ac7f8a74c7)
- [8.0.1178 : using old compiler on MS-Windows](https://github.com/vim/vim/commit/73f4439ca690d8224df8f88dc71a43fafdcc89bf)
- [8.0.1179 : Test&#x5f;popup&#x5f;and&#x5f;window&#x5f;resize() does not always pass](https://github.com/vim/vim/commit/631820536e4084e01bf990f9314ec385b60b21d7)
- [8.0.1180 : MS-Windows testclean target deletes the color script](https://github.com/vim/vim/commit/75f69e5a183e00cb5aa05d43c1ecf3b7dead7a98)
- [8.0.1181 : tests using Vim command fail on MS-Windows](https://github.com/vim/vim/commit/4635e11c6bd63d8be90a2d9455802c6e7bde28f5)
- [8.0.1182 : cannot see or change mzscheme dll name](https://github.com/vim/vim/commit/0ab35b279f9d64f32e99424166b047f90cb042d8)
- [8.0.1183 : MS-Windows build instructions are outdated](https://github.com/vim/vim/commit/18cfa940e2bbf657e8f90e6486812876198ef963)
- [8.0.1184 : the :marks command is not tested](https://github.com/vim/vim/commit/9b69f22e66d51d764e9ade87ae8a57ac13ab7348)
- [8.0.1185 : Ruby library includes minor version number](https://github.com/vim/vim/commit/9202162c5c069c925b6b9bf84e546fbd362cdf46)
- [8.0.1186 : still quite a few old style tests](https://github.com/vim/vim/commit/4a6fcf8047de13c7949ab2f27f7774acaec4ae4d)
- [8.0.1187 : building with lua fails for OSX on Travis](https://github.com/vim/vim/commit/8065cf2bfba93a8aeed29f591152edeb0841bff6)
- [8.0.1188 : autocmd test fails on MS-Windows](https://github.com/vim/vim/commit/1d68d9b2bd60d848552c08763e590edde16056c9)
- [8.0.1189 : E172 is not actually useful](https://github.com/vim/vim/commit/6047e2c722347ae9443c2632c6e061918ccef9c6)
- [8.0.1190 : unusable after opening new window in BufWritePre event](https://github.com/vim/vim/commit/2c33d7bb69c4c2c5b0e39b03cc4b0c04cfdfbb0b)
- [8.0.1191 : MS-Windows: missing 32 and 64 bit files in installer](https://github.com/vim/vim/commit/6199d43f4b59a9bb1c87d408c5b33fa19a23ebcd)
- [8.0.1192 : MS-Windows: terminal feature not enabled by default](https://github.com/vim/vim/commit/ac8069bb63c5954c787fe93b7a9265de3c0c6853)
- [8.0.1193 : crash when wiping out a buffer after using getbufinfo()](https://github.com/vim/vim/commit/b2c8750c4e95f64a8dff912af81e13318b3f6ed6)
- [8.0.1194 : actual fg and bg colors of terminal are unknown](https://github.com/vim/vim/commit/65e4c4f6868882a380c319632a1728a5e7d274ad)
- [8.0.1195 : can't build on MS-Windows](https://github.com/vim/vim/commit/9377df3ecd0bd3acb5d46cb8af7fe60867f247f2)
- [8.0.1196 : crash when t&#x5f;RF is not set](https://github.com/vim/vim/commit/a20f83df1d646a2d296af835ba1154f09986f102)
- [8.0.1197 : MS-Windows build instructions are not up to date](https://github.com/vim/vim/commit/81b07b527e8b9c35d12fd45b63885d6938564723)
- [8.0.1198 : older compilers don't know uint8&#x5f;t](https://github.com/vim/vim/commit/00ce63dc230b6c8f55997b09396aee6d6d722734)
- [8.0.1199 : when 'clipboard' is "autoselectplus" star register is set](https://github.com/vim/vim/commit/8bfe07b70878c83c838320cc6a93d62cd78b953f)
- [8.0.1200 : tests switch the bell off twice](https://github.com/vim/vim/commit/67418d97b457d2e27fe342472d3c9fd342ffc47f)
- [8.0.1201 : "yL" is affected by 'scrolloff'](https://github.com/vim/vim/commit/44cc4cf72fdd12cd9a779031d443aac4254d51ae)
- [8.0.1202 : :wall gives an errof for a terminal window](https://github.com/vim/vim/commit/059db5c29ffef283a4b90bab9228708fa32e3dd2)
- [8.0.1203 : terminal window mistreats composing characters ](https://github.com/vim/vim/commit/6daeef1933be68055aabe1d55f8467d46a707753)
- [8.0.1204 : a QuitPre autocommand may get the wrong file name](https://github.com/vim/vim/commit/87ffb5c1a3aa506a1be07af4e794b3753f839dc3)
- [8.0.1205 : it is possible to unload a changed buffer](https://github.com/vim/vim/commit/ff930cad8a9100eeb04256aab1a14de993c1d7e9)
- [8.0.1206 : no autocmd for entering or leaving the command line](https://github.com/vim/vim/commit/fafcf0dd59fd9c4ef743bb333ae40d1d322b6079)
- [8.0.1207 : profiling skips the first and last script line](https://github.com/vim/vim/commit/67435d9983965c5c77fc74f0559779ce4554dacb)
- [8.0.1208 : 'statusline' drops empty group with highlight change ](https://github.com/vim/vim/commit/6b89dbb55f84c485310c8c9e094dbafe3ecbace6)
- [8.0.1209 : still too many old style tests](https://github.com/vim/vim/commit/53f0c962394dc6bc66d5b0762af9434e672d1b25)
- [8.0.1210 : CTRL-G and CTRL-T are ignored with typeahead](https://github.com/vim/vim/commit/f8e8c0643b1cd97db11912bc4f773e1328a0da02)
- [8.0.1211 : cannot reorder tab pages with drag & drop](https://github.com/vim/vim/commit/ca05aa24af88836f8aa792360b3780589f294981)
- [8.0.1212 : MS-Windows: tear-off menu does not work on 64 bit](https://github.com/vim/vim/commit/66857f410426ca335f4771a58a32b2d14a7e52b9)
- [8.0.1213 : setting 'mzschemedll' has no effect](https://github.com/vim/vim/commit/2e4cb3b042faba5efc9b5a8b4f86942a417494ce)
- [8.0.1214 : accessing freed memory when EXITFREE is set](https://github.com/vim/vim/commit/4f1982800f0aff28df6875e718a786f6c4b11ad9)
- [8.0.1215 : newer gcc warns for implicit fallthrough](https://github.com/vim/vim/commit/2f40d129bf45cd35976e4120336ae6d504f5a5dd)
- [8.0.1216 : tabline is not always updated for :file command](https://github.com/vim/vim/commit/6ce650480844bfaa5410874416b4a2e15f40b870)
- [8.0.1217 : can't use remote eval to inspect vars in debug mode](https://github.com/vim/vim/commit/d99388ba8535a6fecf7d0bf7b982832c0b816062)
- [8.0.1218 : writing to freed memory in autocmd](https://github.com/vim/vim/commit/8d84ff1a3c8cfe59399d3f675ec080066582fdb6)
- [8.0.1219 : terminal test is flaky](https://github.com/vim/vim/commit/f204e05ae9f6bc5d922d14d457e16b590efbf3e5)
- [8.0.1220 : skipping empty statusline groups is not correct](https://github.com/vim/vim/commit/235dddf1f4afe3a40047dbf2aca1bd177b7be18b)
- [8.0.1221 : still too many old style tests](https://github.com/vim/vim/commit/15993ce9210e8b8d4bc11e1d640f6447b18d3e6c)
- [8.0.1222 : test functions interfere with each other](https://github.com/vim/vim/commit/ce11de87e26e1420703242f8e07b4fd69c4032ba)
- [8.0.1223 : crash when using autocomplete and tab pages](https://github.com/vim/vim/commit/9ad89c6c4f89cd710d8244d8010b8b0ae30ba79d)
- [8.0.1224 : still interference between test functions](https://github.com/vim/vim/commit/cf1ba35fc2ebd41b9a7738bbd1f026f5311560aa)
- [8.0.1225 : no check for spell region being zero](https://github.com/vim/vim/commit/ee03b941241eae1d36bc29b84eec09116cefe7cd)
- [8.0.1226 : edit and popup tests failing](https://github.com/vim/vim/commit/2a45d64d0a7ab28d77eee277244e76dbbf8c2db8)
- [8.0.1227 : undefined left shift in readfile() ](https://github.com/vim/vim/commit/dc1c98129484e7879bc6dbf38e523beb730988b6)
- [8.0.1228 : invalid memory access in GUI test](https://github.com/vim/vim/commit/0e19fc07e73214f94441cb3a495504a1de21eb07)
- [8.0.1229 : condition in vim&#x5f;str2nr() is always true](https://github.com/vim/vim/commit/9a91c7a1f9134f799b8672a4e3844781263e8cf3)
- [8.0.1230 : CTRL-A in Visual mode uses character after selection](https://github.com/vim/vim/commit/ce1577502693940f1f9285cd803a075b738a2330)
- [8.0.1231 : expanding file name drops dash](https://github.com/vim/vim/commit/c312b8b87a589ed8452dbf0f555f05ff86d04692)
- [8.0.1232 : MS-Windows users are confused about default mappings](https://github.com/vim/vim/commit/c3fdf7f80b2febdd8a8f7a1310631567d257d66a)
- [8.0.1233 : typo in dos installer](https://github.com/vim/vim/commit/b9fce6cbf7ed0a2766582c5db797ccf99a838a13)
- [8.0.1234 : MS-Windows: composing chars are not shown properly](https://github.com/vim/vim/commit/a6ce1ccf5c10baa5c2a25897c46961d751a21dda)
- [8.0.1235 : cannot disable the terminal feature in a huge build](https://github.com/vim/vim/commit/ef83956e1e67736b4c6b886d897b74f022622a74)
- [8.0.1236 : Mac features are confusing](https://github.com/vim/vim/commit/d057301b1f28736f094affa17b190244ad56e8d9)
- [8.0.1237 : ":set scroll&" often gives an error](https://github.com/vim/vim/commit/af2d20c6285c1d2973e3d9b5e8f727e3ed180493)
- [8.0.1238 : incremental search only shows one match](https://github.com/vim/vim/commit/2e51d9a0972080b087d566608472928d5b7b35d7)
- [8.0.1239 : cannot use a lambda for the skip argument to searchpair()](https://github.com/vim/vim/commit/48570488f17e397183ea7d5c7ca67d6e4ffb013d)
- [8.0.1240 : MS-Windows: term&#x5f;start() does not support environment](https://github.com/vim/vim/commit/ba6febd380c931b92361a189e85b19ed467c9c64)

## 新着スクリプト

- [Colortemplate : A tentative proposal for an official colorscheme template](https://vim.sourceforge.io/scripts/script.php?script_id=5614)
- [man.vim : open man pages within vim using :e man:page or :e man:section:page](https://vim.sourceforge.io/scripts/script.php?script_id=5615)
- [Archery : Vim colorscheme inspired by Arch Linux colors](https://vim.sourceforge.io/scripts/script.php?script_id=5616)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](https://vim.sourceforge.io/scripts/script.php?script_id=273) (1097)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](https://vim.sourceforge.io/scripts/script.php?script_id=1658) (781)
3. [AutomaticLaTeXPlugin : Background compilation, completion, bib serch, toc and other nice features.](https://vim.sourceforge.io/scripts/script.php?script_id=2945) (655)
4. [wombat256.vim : Wombat for 256 color xterms](https://vim.sourceforge.io/scripts/script.php?script_id=2465) (626)
5. [python.vim : Enhanced version of the python syntax highlighting script](https://vim.sourceforge.io/scripts/script.php?script_id=790) (434)
6. [nginx.vim : initial version](https://vim.sourceforge.io/scripts/script.php?script_id=1886) (425)
7. [c.vim : C/C++ IDE -- Write and run programs. Insert statements, idioms, comments etc.](https://vim.sourceforge.io/scripts/script.php?script_id=213) (375)
8. [molokai : A port of the monokai scheme for TextMate](https://vim.sourceforge.io/scripts/script.php?script_id=2340) (371)
9. [OmniCppComplete : C/C++ omni-completion with ctags database](https://vim.sourceforge.io/scripts/script.php?script_id=1520) (339)
10. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](https://vim.sourceforge.io/scripts/script.php?script_id=159) (285)

## vim-jp issues

Open : 258 (-1) | Closed : 853 (+9)

- [Issue #1104 : インクリメンタルサーチ中にすべてのマッチをハイライトしたい](https://github.com/vim-jp/issues/issues/1104)
- [Issue #1105 : LUAオプション付きビルド時のmingwでのパス指定について](https://github.com/vim-jp/issues/issues/1105)
- [Issue #1106 : Win32: 異体字が正しく表示されない](https://github.com/vim-jp/issues/issues/1106)
- [Issue #1107 : :terminal + GitBashで256color.pl実行後に文字化け](https://github.com/vim-jp/issues/issues/1107)
- [Issue #1108 : newやvnewで作成したウィンドウにてsyntax regionで定義した色が付かない](https://github.com/vim-jp/issues/issues/1108)
- [Issue #1109 : 編集中の Window が閉じられてしまう](https://github.com/vim-jp/issues/issues/1109)
- [Issue #1110 : :fileでファイル名を変更した時にタブページのラベルの名前が更新されない](https://github.com/vim-jp/issues/issues/1110)
- [Issue #1111 : menu.vimの読み込みを高速化する試み](https://github.com/vim-jp/issues/issues/1111)
