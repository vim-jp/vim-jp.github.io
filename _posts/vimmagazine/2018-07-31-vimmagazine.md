---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2018 年 07 月号

---

## 話題

*   [Asakusa.vim #2](https://asakusavim.connpass.com/event/93282/) が 2018-07-18 (水) に秋葉原で開催され、Vim使いたちの貴重な知見が共有されました
*   [Asakusa.vim #3](https://asakusavim.connpass.com/event/95401/) が 2018-08-03 (金) に二子玉川で開催予定です。
*   [VimConf 2018](https://vimconf.org/2018/) 関連
    *   VimConf 2018 が 2018-11-24 に秋葉原のアキバPLAZAで開催されます
    *   引き続き、[VimConf 2018の企業スポンサー公募](https://vim-jp.org/blog/2018/06/07/VimConf2018-sponsor-ja.html) 中です
    *   引き続き、[VimConf 2018のスピーカーの公募](https://vim-jp.org/blog/2018/06/28/VimConf2018-cfp-ja.html) 中です
*   [RailsDM 2018](https://techplay.jp/event/679666) というRuby on Railsのカンファレンスで、でVimのC実装に関する発表が行われました。
    *   [スライド](https://speakerdeck.com/ujihisa/vim-railsdm-2018-07)と[動画](https://www.youtube.com/watch?v=12CCKbGg1EA)
*   [meguro.vim #11](https://megurovim.connpass.com/event/94477/) が 2018-08-11 (土) に開催予定です。

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

*   8.1.0142: 組込み termcap の xterm, vt320 エントリにキーパッドのエスケープシーケンスが追加されました
*   8.1.0182: Unicode 11 に対応しました
*   8.1.0184: `wincol` と `winrow` の項目が `getwininfo()` の返り値に追加されました (関連: 8.1.0187)
*   8.1.0198: `'redrawtime'` で設定した時間を過ぎてシンタックスハイライトが無効になる場合はメッセージが出るようになりました
*   8.1.0214: `autochdir` が `has()` 関数と `:version` の出力結果で確認できるようになりました
*   8.1.0218: `matchadd()`, `matchaddpos()` 関数に現在のウィンドウ以外のウィンドウ番号やウィンドウIDを指定できる `window` の項目が追加されました
*   8.1.0225: インサートモードで `CTRL-O` コマンドを実行した際、`mode()` 関数が `niI`, `niR`, `niV` 等の文字列を返すようになりました

*   ランタイムの更新
    *   reStructuredText の ftplugin スクリプトで reStructuredText 標準ルールの設定を行うようになりました (`setlocal expandtab shiftwidth=3 softtabstop=3 tabstop=8`)
        *   `shiftwidth=3 softtabstop=3` についてはディレクティブ (..) と 順序付きリスト (1.) のより良いサポートのためですが、問題を起こす可能性もあるため改善の余地ありとのこと ([参考](https://github.com/vim/vim/blob/91f84f6e11cd879d43d651c0903d85bff95f0716/runtime/ftplugin/rst.vim#L21-L24))
        *   これを抑制するには `let g:rst_style = 0` と vimrc に書く必要があります
    *   reStructuredText の ftplugin スクリプトでセクションやサブセクションごとに折畳みを行うようになりました
    *   WebAssembly テキストフォーマット用ファイルタイプ `wast` の ftplugin, indent, syntax プラグインが追加されました

## Vimに関する脆弱性

特筆すべき脆弱性の報告はありませんでした。

## リリース情報

- [8.1.0128 : building with MinGW does not work out-of-the-box](https://github.com/vim/vim/commit/304925e75031751fc87b51e387eb079bf3ae2066)
- [8.1.0129 : still some xterm-like terminals get a stray "p"](https://github.com/vim/vim/commit/668324ef4fd59980dc57b2eefe57eddc57c7b920)
- [8.1.0130 : ":profdel func" does not work if func was called already](https://github.com/vim/vim/commit/ad6480961080f80a455b2394f27b02935a2ded52)
- [8.1.0131 : :profdel is not tested](https://github.com/vim/vim/commit/1fbfe7c48cb711f5a6deae535b3ec3bfe7952ce9)
- [8.1.0132 : lua tests are old style](https://github.com/vim/vim/commit/4ff4814b383bc85fbf5d8f62c8022f4379d7a490)
- [8.1.0133 : tagfiles() can have duplicate entries](https://github.com/vim/vim/commit/46577b5e5445c4aaa1e7ae1764373d11dae71663)
- [8.1.0134 : Lua interface does not support funcref](https://github.com/vim/vim/commit/ca06da92432a57e5dcf3e0eebd322fae80941a6b)
- [8.1.0135 : undo message delays screen update for CTRL-O u](https://github.com/vim/vim/commit/e0429681aed5448d1c3b8399d152a2b1a05ea6d7)
- [8.1.0136 : Lua tests don't cover new features](https://github.com/vim/vim/commit/2f362bf7f9acc9ec87799d1e41bf0ae7712d1f7a)
- [8.1.0137 : CI does not run with TCL](https://github.com/vim/vim/commit/ade55787978e15fe57c5cedf38c9f85bfe1d983c)
- [8.1.0138 : negative value of 'softtabstop' not used correctly](https://github.com/vim/vim/commit/33d5ab3795720b7d986f9f17f660ee9e448466e0)
- [8.1.0139 : Lua tests fail on some platforms](https://github.com/vim/vim/commit/a8a60d0c6b292216e55f005cf9637789a771d34b)
- [8.1.0140 : recording into a register has focus events](https://github.com/vim/vim/commit/972bfddc6b3f52ae0865ad8c0bf6089bc8a9883a)
- [8.1.0141 : :cexpr no longer jumps to the first error](https://github.com/vim/vim/commit/531b9a3a638cb1147db4f096bfbcb14d2ae32b04)
- [8.1.0142 : xterm and vt320 builtin termcap missing keypad keys](https://github.com/vim/vim/commit/e6882bdc44efbcea0af647f5759a8f48cf29e27c)
- [8.1.0143 : matchit and matchparen don't handle E363](https://github.com/vim/vim/commit/3d1d6475f9665660c80cc53a7da2d5450b8b8d08)
- [8.1.0144 : the :cd command does not have good test coverage](https://github.com/vim/vim/commit/b2e0c94a4d27e3e6222d26f13e0418a85cab21a2)
- [8.1.0145 : test with grep is failing on MS-Windows](https://github.com/vim/vim/commit/851332ea9cdabf24980a7f7c293da81fa990f864)
- [8.1.0146 : when $LANG is set the compiler test may fail](https://github.com/vim/vim/commit/f0447e89a52885630947510f2d1b55f665a1a20e)
- [8.1.0147 : compiler warning when building with Python 3.7](https://github.com/vim/vim/commit/3b48b11c07b85520f4ca3566e92c3f6486509b72)
- [8.1.0148 : memory leak when using :tcl expr command](https://github.com/vim/vim/commit/92959fa46d305ded1da433ad6d9a8d38db240c74)
- [8.1.0149 : session is wrong with multiple tabs when :lcd was used](https://github.com/vim/vim/commit/26d4b896a789e65df7ee0cf3e25056eabc523fda)
- [8.1.0150 : insufficient test coverage for Tcl](https://github.com/vim/vim/commit/fd34cebe9a3a179a5639355c43d2d4959182a279)
- [8.1.0151 : mksession test fails on MS-Windows](https://github.com/vim/vim/commit/81e2ac7bb68a0321529ad5c89c8674caab4bca6b)
- [8.1.0152 : cannot easily run individual tests on MS-Windows](https://github.com/vim/vim/commit/f3dc235576da7394fbe743aba732f43289f32c24)
- [8.1.0153 : (after 8.1.0152) build with SHADOWDIR fails](https://github.com/vim/vim/commit/e85ce6ea2e09fe4d61c94e34e9e437c3c1f20b6e)
- [8.1.0154 : crash with "set smarttab shiftwidth=0 softtabstop=-1"](https://github.com/vim/vim/commit/c9fe5ab3b093803b6e8d03358ba16aca6b6f0db1)
- [8.1.0155 : evim.man missing from the distribution](https://github.com/vim/vim/commit/ae728edc93fe5a2f8c39fa75b1ae9ed71526858c)
- [8.1.0156 : MS-Windows compiler warning](https://github.com/vim/vim/commit/5d2f40792e47a48fc56efafd30b68660a42184a7)
- [8.1.0157 : old iTerm2 is not recognized, resulting in stray output](https://github.com/vim/vim/commit/e330ef42f256e60017e5f8bf4be79a5308fe9e2f)
- [8.1.0158 : GUI: input() fails if CTRL-C was pressed before](https://github.com/vim/vim/commit/1ebff3dc93b6d022ccfe0613c1d1ee2d62fc7935)
- [8.1.0159 : completion for user names does not work for a prefix.](https://github.com/vim/vim/commit/6c5d1043022520512ee36aa1a29662b60af33c95)
- [8.1.0160 : no Danish manual translations](https://github.com/vim/vim/commit/cbbe4ab4b27a2060506fcd2427aed8d8fbff1a21)
- [8.1.0161 : buffer not updated with 'autoread' set if file was deleted](https://github.com/vim/vim/commit/386bc82a3f82f70bad75aaad74dba57a176b5840)
- [8.1.0162 : Danish and German man pages are not installed](https://github.com/vim/vim/commit/08505312fe1ad3b843f95282bff9b0fbbbea508d)
- [8.1.0163 : insufficient testing for Tcl](https://github.com/vim/vim/commit/2549acf794109731eab693bd396bb863d0e2cff4)
- [8.1.0164 : luaeval('vim.buffer().name') returns an error](https://github.com/vim/vim/commit/fe08df452af10db8a24dbeb1bd9ef09492a4bc66)
- [8.1.0165 : :clist output can be very long](https://github.com/vim/vim/commit/4cde86c2ef885e82fff3d925dee9fb5671c025cf)
- [8.1.0166 : using dict&#x5f;add&#x5f;nr&#x5f;str() is clumsy](https://github.com/vim/vim/commit/e0be167a805fd547c25ec1ec97fd4c7f13046236)
- [8.1.0167 : lock flag in new dictitem is reset in many places](https://github.com/vim/vim/commit/c89d4b35300b98cf68b14c89c8e1add51bd857e3)
- [8.1.0168 : output of :marks is too short with multi-byte chars](https://github.com/vim/vim/commit/9d5185bf9dfaef59e47c573a60044a21d5e29c0c)
- [8.1.0169 : calling message&#x5f;filtered() a bit too often](https://github.com/vim/vim/commit/a9defadb8f03ecd03f3297305d5482ba380774dc)
- [8.1.0170 : invalid memory use with complicated pattern](https://github.com/vim/vim/commit/2338c32b53d20dc18540b1a20845bcd8a6371bff)
- [8.1.0171 : typing CTRL-W n in a terminal window causes ml&#x5f;get error](https://github.com/vim/vim/commit/875cf8789426cc258d85358ea2c86744a5a87b16)
- [8.1.0172 : 'viminfofile' option does not behave like a file name](https://github.com/vim/vim/commit/c229e54a69468722ca2449e807e90445b7479659)
- [8.1.0173 : compiler warning on MS-Windows](https://github.com/vim/vim/commit/6259e5769dd50d8a3b5b99f553bab34ff5c8a6ce)
- [8.1.0174 : after paging up and down fold line is wrong](https://github.com/vim/vim/commit/907dad72ef9d29422352fb74ba156e7085a3fc71)
- [8.1.0175 : marks test fails in very wide window](https://github.com/vim/vim/commit/bde14d8e24f6b8ca409290733dbf11cb6fee5751)
- [8.1.0176 : overlapping string argument for strcpy()](https://github.com/vim/vim/commit/18085fae7482906f8e94ecc7386ecf6a02dc407d)
- [8.1.0177 : defining function in sandbox is inconsistent](https://github.com/vim/vim/commit/93343725b5fa1cf580a24302455980faacae8ee2)
- [8.1.0178 : warning for passing pointer to non-pointer argument](https://github.com/vim/vim/commit/e76c4b237d41d0e62ca5d4fc76d1dd163fe222c0)
- [8.1.0179 : redundant condition for boundary check](https://github.com/vim/vim/commit/efc81331e7bceafa2d7bc452a98d6ec1943b4292)
- [8.1.0180 : static analysis errors in Lua interface](https://github.com/vim/vim/commit/d6ef5f9b3d3df2d5dcc666c8741e99fcc77043f6)
- [8.1.0181 : memory leak with trailing characters in skip expression](https://github.com/vim/vim/commit/a43ebe9454386427ca38c75810e2d36991f17812)
- [8.1.0182 : Unicode standard was updated](https://github.com/vim/vim/commit/4fc85003c5f53d6ceb3dbea4b8fae5681615a20c)
- [8.1.0183 : Lua API changed, breaking the build](https://github.com/vim/vim/commit/174136713181a1d1460951d7a0392b16603f81bb)
- [8.1.0184 : not easy to figure out the window layout](https://github.com/vim/vim/commit/b6959a8e06cef6d2126b030b2f8acd49457a3582)
- [8.1.0185 : running tests writes lua.vim even though it is not used](https://github.com/vim/vim/commit/1ce9a1515be5255641d31ab5b0e9418d8f250d84)
- [8.1.0186 : test for getwininfo() fails in GUI](https://github.com/vim/vim/commit/44a693a1bdfe987edb649677c94c214f94b2aada)
- [8.1.0187 : getwininfo() and win&#x5f;screenpos() return different numbers](https://github.com/vim/vim/commit/7132ddc1014a4e1230f0080e418221e5843e1277)
- [8.1.0188 : no test for ":cscope add"](https://github.com/vim/vim/commit/2196bca7377ff245866cc3cee65b0adb48432ac3)
- [8.1.0189 : function defined in sandbox not tested](https://github.com/vim/vim/commit/d90a144eda047816acffc7a8f297b43a7120710e)
- [8.1.0190 : Perl refcounts are wrong](https://github.com/vim/vim/commit/18c4f1badbc96d39de5b348f268ac8d55c2b0b67)
- [8.1.0191 : Perl test fails in 24 line terminal](https://github.com/vim/vim/commit/3166afd945e57f04b231e71e054b9593cc29ff0b)
- [8.1.0192 : executing regexp recursively fails with a crash](https://github.com/vim/vim/commit/0270f38e1ae484c31a80c813a08691c47a207f1a)
- [8.1.0193 : terminal debugger buttons don't always work](https://github.com/vim/vim/commit/ca4cc018addbeb3ac5d0e05f18847015f91ff814)
- [8.1.0194 : possibly use of NULL pointer](https://github.com/vim/vim/commit/414998023fbff15cce20ef01a54d0366370ad8b6)
- [8.1.0195 : terminal debugger commands don't always work](https://github.com/vim/vim/commit/963c1ad5d072346d9e95d4c3be066b5e03c601d3)
- [8.1.0196 : terminal debugger error with .gdbinit file](https://github.com/vim/vim/commit/f63db65b2418140d1bdbc032511f530234bd2496)
- [8.1.0197 : Windows GUI: title for search/replace is wrong](https://github.com/vim/vim/commit/9e42c866484fe0199e8b17e7c44489386173acc9)
- [8.1.0198 : there is no hint that syntax is disabled for 'redrawtime'](https://github.com/vim/vim/commit/0a6efcd27d62935c465b4406c0c0db9be10a0ddb)
- [8.1.0199 : spellbadword() does not check for caps error](https://github.com/vim/vim/commit/66ab916935585391b2efaa8e39075e1ef94717b1)
- [8.1.0200 : spellbadword() not tested](https://github.com/vim/vim/commit/872e451e8c326d5dd3062ef621fcbf0a4c5bef78)
- [8.1.0201 : newer Python uses "importlib" instead of "imp"](https://github.com/vim/vim/commit/79a494d5e2f97c10e74f92ea529552623c314422)
- [8.1.0202 : :version always shows +packages](https://github.com/vim/vim/commit/6183ccbd679751ff5b138f23a34ead3d7bbc5c1b)
- [8.1.0203 : building with Perl 5.28 fails on Windows](https://github.com/vim/vim/commit/578333b2ecdfef3109cb07a82c3aa08ad3dc5664)
- [8.1.0204 : inputlist() is not tested](https://github.com/vim/vim/commit/947b39e761b8a95cc1bd37ad0c2c30552238809a)
- [8.1.0205 : invalid memory access with invalid modeline](https://github.com/vim/vim/commit/9cf4b5005f12ce1d6692266140bdda05d0312d79)
- [8.1.0206 : duplicate test function name](https://github.com/vim/vim/commit/cd96eef3a869557bd3d2d4497861d87cb525db06)
- [8.1.0207 : need many menu translation files to cover regions](https://github.com/vim/vim/commit/68f1b1b37fa7aba985d9f8727fd9f0f3eb0c19a9)
- [8.1.0208 : file left behind after running individual test](https://github.com/vim/vim/commit/7624af02948de881eb5294fe5d266b948404d489)
- [8.1.0209 : stderr output from Ruby messes up display](https://github.com/vim/vim/commit/b6c8cd8dc3fbbe9da1180ce1b75f2252df2c1318)
- [8.1.0210 : still a few K&R function declarations](https://github.com/vim/vim/commit/6f8d2ac6f1f8a4b971a4c9b27f9250288198f3bb)
- [8.1.0211 : expanding a file name "~" results in $HOME](https://github.com/vim/vim/commit/00136dc321586800986e8f743c2f108f5eecbf92)
- [8.1.0212 : preferred cursor column not set in interfaces](https://github.com/vim/vim/commit/53901442f37a59e5495165f91db5574c0b43ab04)
- [8.1.0213 : CTRL-W CR does not work properly in a quickfix window](https://github.com/vim/vim/commit/0a08c63da17dfd93ac2885e3f3f8a083a9b3131c)
- [8.1.0214 : +autochdir feature not reported by has() or :version](https://github.com/vim/vim/commit/83ec2a7f5fb481b30a5d556b6aad49a62585bccd)
- [8.1.0215 : no error if configure --with-x cannot configure X](https://github.com/vim/vim/commit/d2a054910bf8f0888fac81e46a64a22dd4a763b5)
- [8.1.0216 : part of file not indented properly](https://github.com/vim/vim/commit/8e85db0376604274bccf12bf5859d92b4192aca1)
- [8.1.0217 : compiler warning for variable set but not used](https://github.com/vim/vim/commit/fd249460fe600dba479bca03058e679bae6b5d52)
- [8.1.0218 : cannot add matches to another window](https://github.com/vim/vim/commit/95e51470f10e1ddcc4b2ce53e4f7ff7aa2e58417)
- [8.1.0219 : expanding ## fails to escape backtick](https://github.com/vim/vim/commit/2c8c681bfcd5138a0ec8ce018216dc2dc69a11a0)
- [8.1.0220 : Ruby converts v:true and v:false to a number](https://github.com/vim/vim/commit/d84b26a03b13cd816d80ff32b61e8de740d499ce)
- [8.1.0221 : not enough testing for the Ruby interface](https://github.com/vim/vim/commit/edd6aacb010ad2402fc98f19e7f6b7c29ba2656f)
- [8.1.0222 : errors are reported for "make install"](https://github.com/vim/vim/commit/73b4abae5d47fe7e8b5829aaa0abe5b1eac8a408)
- [8.1.0223 : completing shell command finds sub-directories in $PATH](https://github.com/vim/vim/commit/6ab9e429da18f4d784222a9f7dfafb7c0218b7eb)
- [8.1.0224 : hang in bracketed paste mode when t&#x5f;PE not encountered](https://github.com/vim/vim/commit/fdd7155fab3447b38280035c66178330f8f041e7)
- [8.1.0225 : mode() does not indicate using CTRL-O from Insert mode](https://github.com/vim/vim/commit/612cc3888b136e80485132d9f997ed457dbc5501)
- [8.1.0226 : too many #ifdefs](https://github.com/vim/vim/commit/1f0bfe561737cd445532f20d7607a81d1dadddee)
- [8.1.0227 : spaces instead of tabs in makefile](https://github.com/vim/vim/commit/fda95e75721fb221495c69e493ec2761b5d85123)
- [8.1.0228 : dropping files is ignored while Vim is busy](https://github.com/vim/vim/commit/92d147be959e689f8f58fd5d138a31835e160289)
- [8.1.0229 : crash when dumping profiling data](https://github.com/vim/vim/commit/79c2ad50b8069bbb3281bcabe412bd6641756724)

## 新着スクリプト

- [insert&#x5f;cancel.vim : Cancel an insert mode operation](https://www.vim.org/scripts/script.php?script_id=5717)
- [require.vim : require vim easy](https://www.vim.org/scripts/script.php?script_id=5718)
- [simpleterm.vim : simple terminal in vim](https://www.vim.org/scripts/script.php?script_id=5719)
- [perl&#x5f;version&#x5f;bump.vim : Increment Perl version numbers](https://www.vim.org/scripts/script.php?script_id=5720)
- [git-backups : uses git to back up all files](https://www.vim.org/scripts/script.php?script_id=5721)
- [jsonpath : Navigate JSON documents using dot.notation.paths](https://www.vim.org/scripts/script.php?script_id=5722)
- [insert&#x5f;timeout.vim : Leave insert mode after a period of inactivity](https://www.vim.org/scripts/script.php?script_id=5723)
- [digraph&#x5f;search.vim : Search for a digraph by name while inserting ](https://www.vim.org/scripts/script.php?script_id=5724)
- [undoquit.vim :  Undo a :quit -- reopen the last window you closed ](https://www.vim.org/scripts/script.php?script_id=5725)
- [vim-buffest : Edit registers and the quickfix/location list as buffers](https://www.vim.org/scripts/script.php?script_id=5728)

## 月間ダウンロードランキング

1. [jellybeans.vim : Colorful, dark color scheme](https://www.vim.org/scripts/script.php?script_id=2555) (7)
2. [verilog&#x5f;systemverilog.vim : Extending Verilog syntax highlighting for SystemVerilog](https://www.vim.org/scripts/script.php?script_id=1586) (5)
3. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](https://www.vim.org/scripts/script.php?script_id=273) (4)
4. [bufexplorer.zip : Buffer Explorer / Browser](https://www.vim.org/scripts/script.php?script_id=42) (4)
5. [systemverilog.vim : SystemVerilog syntax highlighting](https://www.vim.org/scripts/script.php?script_id=1573) (4)
6. [ctags.exe : Windows version of ctags for 4GL/4Js programs](https://www.vim.org/scripts/script.php?script_id=2288) (4)
7. [Emmet.vim : vim plugins for HTML and CSS hi-speed coding.](https://www.vim.org/scripts/script.php?script_id=2981) (4)
8. [The NERD tree : A tree explorer plugin for navigating the filesystem](https://www.vim.org/scripts/script.php?script_id=1658) (4)
9. [DirDiff.vim : A plugin to diff and merge two directories recursively.](https://www.vim.org/scripts/script.php?script_id=102) (3)
10. [tcl.vim : tcl syntax file](https://www.vim.org/scripts/script.php?script_id=1603) (3)

## vim-jp issues

Open : 262 (+2) | Closed : 916 (+2)

- [Issue #1175 : Python 3.7 に関する問題](https://github.com/vim-jp/issues/issues/1175)
- [Issue #1176 : C ドライブ直下での finddir() による上方検索の結果が想定した動作と異なる](https://github.com/vim-jp/issues/issues/1176)
- [Issue #1177 : フォント選択の「スタイル」が &#x60;guifont&#x60; の直接指定では反映されない](https://github.com/vim-jp/issues/issues/1177)
- [Issue #1178 : Vim script 実行時にファイル先頭からの行番号を取得したい](https://github.com/vim-jp/issues/issues/1178)
