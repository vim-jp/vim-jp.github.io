---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2017 年 9 月号

---

## 話題

*   [Fablic.vim #3](https://fablicvim.connpass.com/event/65469/) が 09/07 (木）に開催されました
*   [Osaka.vim #11](https://osaka-vim.connpass.com/event/64719/) が 09/30 (土) に開催されました
*   [Meguro.vim #6](https://megurovim.connpass.com/event/67608/) が 10/21 (土) に開催予定です
*   [VimConf 2017](http://vim-jp.org/blog/2017/08/04/vimconf2017-venue-and-date-ja.html) が 11/04 (土) に開催予定です

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。


*   8.0.1031: `getqflist()` 関数の引数 "text" が "lines" に変更されました
*   8.0.1034: `term_start()` 関数に `eof_chars` オプションが追加されました
*   8.0.1035: `:terminal` コマンドに `++eof` 引数が追加されました
*   8.0.1038: 打消し線がサポートされました (GUI 版のみ)
*   8.0.1039: `setbufline()` 関数が追加されました
*   8.0.1040: `getqflist()` 関数に `efm` 引数が追加されました
*   8.0.1050: huge ビルド時に端末機能がデフォルトで含まれるようになりました (`--enable-terminal` オプションを configure に指定する必要はありません)
*   8.0.1051: `:terminal` コマンドに半角スペースを含んだ引数を渡せるようになりました (バックスラッシュでスペースをエスケープします)
*   8.0.1052: `term_start()` 関数に `in_io`, `out_io`, `err_io` を指定可能になりました
*   8.0.1074: `:terminal NONE` が Windows でも動くようになりました
*   8.0.1075: Windows でもマウスが Terminal ウィンドウ内で動くようになりました
*   8.0.1076: `term_start()` 関数に `callback`, `out_cb`, `err_cb` をサポートしました
*   8.0.1077: 端末ウィンドウのデバッグを可能にする termdebug プラグインが追加されました (:help terminal-debug)
*   8.0.1085: ブレークポイントを設定・削除する `:Break`, `:Delete` コマンドが termdebug プラグインに追加されました (ソースコードを表示しているウィンドウ内でのみ使用可能です)
*   8.0.1089: `:command` コマンドに `<range>` が追加されました
*   8.0.1090: `v:beval_text` のようにカーソルの下のテキストを取得する `<cexpr>` が追加されました
*   8.0.1092: 式を評価する `:Evaluate`, `K` コマンド等が termdebug プラグインに追加されました (ソースコードを表示しているウィンドウ内でのみ使用可能です)
*   8.0.1108: Terminal ウィンドウ内でマッピングを指定できる `:tmap` コマンドが追加されました (関連パッチ: 8.0.1120)
*   8.0.1112: quickfix list の項目に `idx` と `size` が追加されました
*   8.0.1114: `'iminsert'`, `'imsearch'` オプションのデフォルト値が変更されました
    *   iminsert: インプットメソッドがサポートされていれば 2、されていなければ0 →常に 0
    *   imsearch: インプットメソッドがサポートされていれば 2、されていなければ0 →常に -1
*   8.0.1123: ウィンドウにツールバーが表示できるようになりました (:help window-toolbar)

## Vimに関する脆弱性

特筆すべき脆弱性の報告はありませんでした。

## リリース情報

- [8.0.1031 : "text" argument for getqflist() is confusing](https://github.com/vim/vim/commit/2c809b7c7d2bb5e4b7fd09c3d312cadecf0c1ff0)
- [8.0.1032 : "make tags" doesn't work well on MS-Windows](https://github.com/vim/vim/commit/9ac9dfa9e2b20659e5806982f027fa4fd637c3c7)
- [8.0.1033 : detecting background color does not work in screen](https://github.com/vim/vim/commit/995e4afcfe60aa3c214ba680e7b027a4e05cf62b)
- [8.0.1034 : sending buffer lines to terminal doesn't work on MS-Windows](https://github.com/vim/vim/commit/3346cc4ffb459ecddb97a8c19bcc5834afa4dead)
- [8.0.1035 : sending buffer lines to terminal doesn't work on MS-Windows](https://github.com/vim/vim/commit/ef68e4fa528165f8dd63156feeffc1af629b8d8a)
- [8.0.1036 : ++eof argument for terminal only available on MS-Windows](https://github.com/vim/vim/commit/dada6d2a8e887309e88cb126f1251d81f91b4b9d)
- [8.0.1037 : "icase" of 'diffopt' is not used for highlighting](https://github.com/vim/vim/commit/da22b8cc8b1b96fabd5a4c35c57b04a351340fb1)
- [8.0.1038 : strike-through text not supported](https://github.com/vim/vim/commit/cf4b00c856ef714482d8d060332ac9a4d74e6b88)
- [8.0.1039 : cannot change a line in not current buffer](https://github.com/vim/vim/commit/b31cf2bb0be95d106bd8eef93cc07550591c1d0d)
- [8.0.1040 : cannot use another error format in getqflist()](https://github.com/vim/vim/commit/3653822546fb0f1005c32bb5b70dc9bfacdfc954)
- [8.0.1041 : bogus characters when indenting during visual-block append](https://github.com/vim/vim/commit/e2e69e48134cbfdedea7802810932f8592705024)
- [8.0.1042 : without the syntax feature highlighting doesn't work](https://github.com/vim/vim/commit/9f45ccd1bf2d2de979d252fdd7e5c9eb3fa6cceb)
- [8.0.1043 : warning for uninitialized variable](https://github.com/vim/vim/commit/4ec86ddd77a89766e42bd0a6cfcf10af4c3d03b2)
- [8.0.1044 : warning for uninitialized variable](https://github.com/vim/vim/commit/2254a8ad0cb596a0a0863b1af92353f018c0b355)
- [8.0.1045 : running tests may pollute shell history](https://github.com/vim/vim/commit/6a8691d483914606213a24356a9124fa41c93b69)
- [8.0.1046 : code duplication in diff mode](https://github.com/vim/vim/commit/ae96b8d058cffd9d07b78cb7a9ccd382185b9dd6)
- [8.0.1047 : buffer overflow in Ruby](https://github.com/vim/vim/commit/00ccf54630dc68a9b8aedb92b268f3b697081f68)
- [8.0.1048 : no test for what 8.0.1020 fixes](https://github.com/vim/vim/commit/5e80de3f3e572805fe734b66bc42c13303ad9bdb)
- [8.0.1049 : shell on Mac can't handle long text](https://github.com/vim/vim/commit/d8d85bfb7d650d2fb584ea48aaaebdbc8a2e3dce)
- [8.0.1050 : terminal window feature not included by default](https://github.com/vim/vim/commit/595a40226ea1285689b622042534fd8442402db3)
- [8.0.1051 : cannot run terminal with spaces in argument](https://github.com/vim/vim/commit/9d654a8d8cd3421307445f111785fb303a38c2a0)
- [8.0.1052 : term&#x5f;start() does not allow in&#x5f;io, out&#x5f;io and err&#x5f;io options](https://github.com/vim/vim/commit/e88fc7a574263fd399c6815378bcd8fd228d8b54)
- [8.0.1053 : setline() does not work on startup](https://github.com/vim/vim/commit/9d954207e2cc807b475bb04f8b59ef5bb3772d99)
- [8.0.1054 : terminal test fails on MS-Windows](https://github.com/vim/vim/commit/178333783fac3a5edbc86f2e9c57a21c41f05697)
- [8.0.1055 : bufline test hangs on MS-Windows](https://github.com/vim/vim/commit/11aa62f8f949bb590b4d7792a334885fba5e4137)
- [8.0.1056 : cannot build with +diff but without +multi&#x5f;byte](https://github.com/vim/vim/commit/6e32f615eb824048a98666d1da4416bbeb7b2bed)
- [8.0.1057 : terminal scrape test waits too long](https://github.com/vim/vim/commit/1bfdc07a4996b5905b4971e18a6487a18e6461a4)
- [8.0.1058 : terminal redirection test is flaky](https://github.com/vim/vim/commit/8b53b79ea5304c7714cbb873210c82aa3fc8772f)
- [8.0.1059 : older Gnome terminal returns smaller version number](https://github.com/vim/vim/commit/3d8d2c7ca5dda5bfd62c6d5df6659cf62e960a3d)
- [8.0.1060 : when imstyle is one, mapping \<Left> breaks preediting](https://github.com/vim/vim/commit/3971905bac0fe7e7519b35b1e558b4f1de55708d)
- [8.0.1061 : Coverity: no check for NULL command](https://github.com/vim/vim/commit/6756c7037f07e1fb54e1b9d8e92567ac388617c4)
- [8.0.1062 : Coverity warnings in libvterm](https://github.com/vim/vim/commit/c3f81394eff2b3edc7ea08405743f0d32048374a)
- [8.0.1063 : Coverity warns for NULL check and array use](https://github.com/vim/vim/commit/28550b74bb4373417eb6fbf132bd4211b7b92afa)
- [8.0.1064 : Coverity warns for leaking resource](https://github.com/vim/vim/commit/1b9f9d315f137cc883c7aecd39f3df329b7f9085)
- [8.0.1065 : not all macro examples are included](https://github.com/vim/vim/commit/c35e4cb4fe3ce524d15ac9fdc9a8aa923761bdfd)
- [8.0.1066 : some terminals can't handle requesting cursor mode](https://github.com/vim/vim/commit/2e49b6b20c3d7229edf41ac03931ce32ce0f7bd8)
- [8.0.1067 : try/catch in timer does not prevent it from being stopped](https://github.com/vim/vim/commit/e723c42836d971180d1bf9f98916966c5543fff1)
- [8.0.1068 : vandyke SecureCRT terminal can't handle cursor mode request](https://github.com/vim/vim/commit/618d6d277e7b64ea03fcefdd75d673e8973b46c4)
- [8.0.1069 : still get CTRL-X sometimes](https://github.com/vim/vim/commit/37b53270e1a65d1d56187a3d68439e666e8b75fd)
- [8.0.1070 : terminal test is flaky on Mac](https://github.com/vim/vim/commit/75a60f74ce1000b0ffed2e6a203a05f3cd160c6a)
- [8.0.1071 : putty-color and cygwin-color are not recognized](https://github.com/vim/vim/commit/c6da01a5b8efe9ca4931074c0cf2189d357707a2)
- [8.0.1072 : :highlight command causes a redraw even when nothing changed](https://github.com/vim/vim/commit/99433291b135094d9592c41f96d3ccd60073e2c1)
- [8.0.1073 : may get an endless loop if 'statusline' changes a highlight](https://github.com/vim/vim/commit/ba2929b6afd2fc20479912a8dec789be26a38244)
- [8.0.1074 : ":term NONE" does not work on MS-Windows](https://github.com/vim/vim/commit/2dc9d26c14e410c09e538cccfa90da19ae344ba4)
- [8.0.1075 : MS-Windows: mouse does not work in terminal](https://github.com/vim/vim/commit/1a735d6c9bf4204492b9f9aeba280737fab46ed3)
- [8.0.1076 : term&#x5f;start() does not take callbacks](https://github.com/vim/vim/commit/3c518400d1a51929572dd9fcf77dba94d78d7545)
- [8.0.1077 : no debugger making use of the terminal window](https://github.com/vim/vim/commit/fe386641b0c56c5de2bca8e1f4cd5e2a1f1aea7e)
- [8.0.1078 : using freed memory with ":hi Normal"](https://github.com/vim/vim/commit/b4ea1914b8ca7c368253bd96e6b3cb9e3392da1c)
- [8.0.1079 : memory leak when remote&#x5f;foreground() fails](https://github.com/vim/vim/commit/09d6c3818d7451a659f4d12a4e974237fe823f7d)
- [8.0.1080 : memory leak for eof&#x5f;chars terminal option and buffer name](https://github.com/vim/vim/commit/74121231be50e245d18c64281fdef08e7ec1ed5b)
- [8.0.1081 : memory leak for the channel write queue](https://github.com/vim/vim/commit/aba680a8513124e9556956115db4df35bd4a0e56)
- [8.0.1082 : tests fail when run under valgrind](https://github.com/vim/vim/commit/9d18961323a2a5c3b609c98ce0d78613c71f3532)
- [8.0.1083 : leaking memory in input part of channel](https://github.com/vim/vim/commit/5b5adf5b9cd226d775643872e65c62674d6ada95)
- [8.0.1084 : GTK build has compiler warnings](https://github.com/vim/vim/commit/7be9b50fd7e238722c9ba5c0ef1d2a7e7e52b9e3)
- [8.0.1085 : terminal debugger can't set breakpoints](https://github.com/vim/vim/commit/e09ba7bae5c867f6d3abc184709dd27488318e97)
- [8.0.1086 : can't build with GTK 3](https://github.com/vim/vim/commit/8696bbaf52cca7c1fcbd5000c83406e462b7dbf4)
- [8.0.1087 : Test&#x5f;terminal&#x5f;cwd is flaky](https://github.com/vim/vim/commit/e9f6fd27d0e2dcae3f4aa40c459d5e6a3b3dd102)
- [8.0.1088 : occasional memory use after free](https://github.com/vim/vim/commit/414168d97fad45387a3d7dd16449d15b27079ad8)
- [8.0.1089 : cannot get range count in user command](https://github.com/vim/vim/commit/c168bd4bd3a9b856fc410fc4515dcca1d10d7461)
- [8.0.1090 : cannot get the text under the cursor like v:beval&#x5f;text](https://github.com/vim/vim/commit/65f084749b260746d7f186af4f080298be2df55b)
- [8.0.1091 : test for \<cexpr> fails without +balloon&#x5f;eval feature](https://github.com/vim/vim/commit/95c83c64be8d4af74bfda1f283a5bcf3f110719e)
- [8.0.1092 : terminal debugger can't evaluate expressions](https://github.com/vim/vim/commit/45d5f26d11d9aac2383453d2c1a8582cad1c8a3d)
- [8.0.1093 : various small quickfix issues](https://github.com/vim/vim/commit/b4d5fbabc99917a8069ba32a60c2d73d4f60e128)
- [8.0.1094 : using ssh from Terminal.app runs into xterm incompatibility](https://github.com/vim/vim/commit/c212798333cab2239acd387725987f27c6d8a3f7)
- [8.0.1095 : terminal multibyte escrape test is flaky](https://github.com/vim/vim/commit/a038cb5eab24a9c395861694d9b8f610ec0ccb61)
- [8.0.1096 : terminal window in Normal mode has wrong background](https://github.com/vim/vim/commit/238d43b32859d1b4e6b7072d552289a748cbfee1)
- [8.0.1097 : background color wrong if job changes background color](https://github.com/vim/vim/commit/49a613f84aaa7cccd1d12a2ffc4ebb1956f973aa)
- [8.0.1098 : build failure if libvterm installed](https://github.com/vim/vim/commit/b984b80cd835a483bdca1928bc50eb09b071f373)
- [8.0.1099 : warnings for GDK calls](https://github.com/vim/vim/commit/518d6996998b35ccbc408b1b323f8869c0c1b4d8)
- [8.0.1100 : stuck in redraw loop when 'lazyredraw' is set](https://github.com/vim/vim/commit/072412ed45aa20a67aaa7e387d7f6bf59c7d3fbe)
- [8.0.1101 : channel write fails if writing to log fails](https://github.com/vim/vim/commit/71eeb74a608ddae2b52fb43766616672d0cdda18)
- [8.0.1102 : terminal window does not use Normal colors](https://github.com/vim/vim/commit/c71053c61180ed1ad0c98ba6063b769757832932)
- [8.0.1103 : converting cterm color fails for grey ramp](https://github.com/vim/vim/commit/478e3a78c37de214aadf7399fac77227b7316f40)
- [8.0.1104 : the qf&#x5f;jump() function is too long](https://github.com/vim/vim/commit/ef6b8de42f439271edcf5cf22b2450f2cf302c2e)
- [8.0.1105 : match() and matchend() are not tested](https://github.com/vim/vim/commit/1190cf68e27a123cf9f6fb57897782a3b9f7b810)
- [8.0.1106 : terminal colors wrong on an MS-Windows console](https://github.com/vim/vim/commit/69e8aa75112a0b12811a4f4d11061adceffbbabc)
- [8.0.1107 : terminal debugger jumps to non-existing file](https://github.com/vim/vim/commit/38baa3e63427112d389de5e5942243414d9b1336)
- [8.0.1108 : cannot specify mappings for the terminal window](https://github.com/vim/vim/commit/69fbc9e1dab176f345719436cd89d854df0a2abd)
- [8.0.1109 : timer causes error on exit from Ex mode](https://github.com/vim/vim/commit/f5291f301e9322545f0621b2157e93050d1d4fb3)
- [8.0.1110 : FORTIFY&#x5f;SOURCE from Perl causes problems](https://github.com/vim/vim/commit/e8ff56b28ed67bff4ee1dbe7b4bd620c1c04a535)
- [8.0.1111 : syntax error in configure when using Perl](https://github.com/vim/vim/commit/b5e79ef5a9e85f5bb70eb4cc6e12cbeec2a820ca)
- [8.0.1112 : can't get size or current index from quickfix list](https://github.com/vim/vim/commit/fc2b270cfd36230166df486aae4d96d9d1f32755)
- [8.0.1113 : can go to Insert mode from Terminal-Normal mode](https://github.com/vim/vim/commit/eef9adddab4c7fc7bb1c54267e76849caf565edf)
- [8.0.1114 : default for 'iminsert' is annoying](https://github.com/vim/vim/commit/4cf56bbc85f77846aeb378cfb071677336dfad6d)
- [8.0.1115 : crash when using foldtextresult() recursively](https://github.com/vim/vim/commit/495b7dd213e096361e6f15e7aed313c1d63d9d3e)
- [8.0.1116 : terminal test fails on MS-Windows](https://github.com/vim/vim/commit/1514e8f4271d0039740b968a9f7d0d7f81a262a2)
- [8.0.1117 : Test&#x5f;terminal&#x5f;no&#x5f;cmd hangs on MS-Windows with GUI](https://github.com/vim/vim/commit/e738a1a033cd31cd2568ba99a9e2dca1e65b45ea)
- [8.0.1118 : FEAT&#x5f;WINDOWS adds a lot of #ifdefs](https://github.com/vim/vim/commit/4033c55eca575777718c0701e26635a0cc47d907)
- [8.0.1119 : quitting a split terminal window kills the job](https://github.com/vim/vim/commit/8adb0d03ca2694922da915356d7ede05e31c5a5c)
- [8.0.1120 : :tm means :tmap instead of :tmenu](https://github.com/vim/vim/commit/63c4e8a1986796094e6f15b893f2deccdf482617)
- [8.0.1121 : can uncheck executables in MS-Windows installer](https://github.com/vim/vim/commit/49150a41f86ff5c50de6d679168532682f49c672)
- [8.0.1122 : vimtutor.bat doesn't work well with vim.bat](https://github.com/vim/vim/commit/dde403c2d8f3dabe6fefa7b526958b49a8f2e6e9)
- [8.0.1123 : cannot define a toolbar for a window](https://github.com/vim/vim/commit/1b9645de3c05f37b5c30e78f999351b0cf486ade)
- [8.0.1124 : use of MZSCHEME&#x5f;VER is unclear](https://github.com/vim/vim/commit/f16c71b66927a28196b25a42795d083784ab1b9a)
- [8.0.1125 : wrong window height when splitting window with window toolbar](https://github.com/vim/vim/commit/d326ad6e93604222c64bae51f47521fe9114d1ea)
- [8.0.1126 : endless resize when terminal showing in two buffers](https://github.com/vim/vim/commit/71b29640669872b0209985fc5a3b447de62a4cf0)
- [8.0.1127 : Test&#x5f;peek&#x5f;and&#x5f;get&#x5f;char fails on 32 bit system](https://github.com/vim/vim/commit/a8e93d63e308020ee06eb5782e766941ac417a5a)
- [8.0.1128 : old xterm sends CTRL-X in response to t&#x5f;RS](https://github.com/vim/vim/commit/e22bbf65088a01b465918919782c19eb291f9cf0)
- [8.0.1129 : window toolbar missing a part of the patch](https://github.com/vim/vim/commit/a5a2be26febee62da480265ed9a52f782d5b4388)
- [8.0.1130 : the qf&#x5f;jump() function is still too long](https://github.com/vim/vim/commit/9cb03716c963338f9a98d2ebc7aa3ac8b9c1eea6)
- [8.0.1131 : not easy to trigger an autocommand for new terminal window](https://github.com/vim/vim/commit/b00fdf6eed5fec589a86655249a851c2d9ba3bb8)
- [8.0.1132 : #if condition is not portable](https://github.com/vim/vim/commit/fc7649f8b82efbb4c7066567dd69192d97a2749f)
- [8.0.1133 : syntax timeout not used correctly](https://github.com/vim/vim/commit/f3d769a585040ac47f7054057758809024ef6377)
- [8.0.1134 : superfluous call to syn&#x5f;get&#x5f;final&#x5f;id()](https://github.com/vim/vim/commit/76301956f0c079d893cfd927b11456328bed4f9b)
- [8.0.1135 : W&#x5f;WINCOL() is always the same](https://github.com/vim/vim/commit/53f8174eaeb93b381cf74c58863f8fe82748a22a)
- [8.0.1136 : W&#x5f;WIDTH() is always the same](https://github.com/vim/vim/commit/0263146b5dbbb6c120ce2e7720256503b864425d)
- [8.0.1137 : (after 8.0.1136) cannot build with Ruby](https://github.com/vim/vim/commit/e745d75c3e0d976e73bd17c395e9cce98f671692)
- [8.0.1138 : click in window toolbar starts Visual mode](https://github.com/vim/vim/commit/eb163d73b11c10b461a2839530173a33d7915a33)
- [8.0.1139 : using window toolbar changes state](https://github.com/vim/vim/commit/a21a6a9ade7bec3a07992d4d900d4ce82eeb8a29)
- [8.0.1140 : still old style tests](https://github.com/vim/vim/commit/bb3e6416f157f68799ccb6070fa2f91df6780e79)
- [8.0.1141 : MS-Windows build dependencies are incomplete](https://github.com/vim/vim/commit/96e7a6ee4569526c71b1e89f641526775831ac63)
- [8.0.1142 : window toolbar menu gets a tear-off item](https://github.com/vim/vim/commit/378daf87d380b9f3c4f822786dfbfdcad9ca2db9)
- [8.0.1143 : macros always expand to the same thing](https://github.com/vim/vim/commit/e0de17d84e1612dbc9362989194de5674f221de5)
- [8.0.1144 : using wrong #ifdef for computing length](https://github.com/vim/vim/commit/0b05e491b473dbf39cd9f519030bf6363c272300)
- [8.0.1145 : warning when compiling with Perl](https://github.com/vim/vim/commit/0c6a32963d1c0b5401ffd7f25e6fc7bf86895744)
- [8.0.1146 : redraw when highlight is set with same names](https://github.com/vim/vim/commit/452030e530aad9b08fcfa71737d098b33c752b85)
- [8.0.1147 : fail to build with tiny features](https://github.com/vim/vim/commit/7c456a4511dff70bb2fc23ebc3dc3d9201dbaacb)
- [8.0.1148 : gN doesn't work on last match with 'wrapscan' off](https://github.com/vim/vim/commit/22ab547dc281092d6a9d397db712a11733c38e97)
- [8.0.1149 : libvterm colors differ from xterm](https://github.com/vim/vim/commit/a8fc0d3817f69b3599a1b56d74a2ddbd3139f40c)
- [8.0.1150 : MS-Windows GUI: dialog font size is incorrect](https://github.com/vim/vim/commit/6edeaf387c7218a84a5521a4e82b4000496b523b)
- [8.0.1151 : "vim -c startinsert!" doesn't append](https://github.com/vim/vim/commit/09ca932f8e7d63a83b39baa7c03d4c6145e3baab)
- [8.0.1152 : encoding of error message wrong in Cygwin terminal](https://github.com/vim/vim/commit/2a02745709127bd56ccdbac8c568b3c25f3072a7)
- [8.0.1153 : no tests for diff&#x5f;hlID() and diff&#x5f;filler()](https://github.com/vim/vim/commit/97fbc404fc56f76df12b2d2658b1d6efda28d5dd)
- [8.0.1154 : 'indentkeys' does not work properly](https://github.com/vim/vim/commit/1b38344e00af65df12946fffda7f3201621c35ef)
- [8.0.1155 : Ruby command triggers a warning](https://github.com/vim/vim/commit/d1bc96ce2466ac1b4af4a2c36de455a1df318d1d)
- [8.0.1156 : trouble from removing one -W argument from Perl CFLAGS](https://github.com/vim/vim/commit/1ec96c926974208bd47710384c8202c47d0752d0)
- [8.0.1157 : compiler warning on MS-Windows](https://github.com/vim/vim/commit/9cf39cc57f63f57b4a335cd25371eb456824f049)
- [8.0.1158 : still old style tests](https://github.com/vim/vim/commit/db51007108a6ab0671e7f7b4844557cbe647185f)
- [8.0.1159 : typo in #ifdef](https://github.com/vim/vim/commit/d371bbe0ab4b07dd3aa8d0f77905d222f6c5fd7e)
- [8.0.1160 : getting tab-local variable fails after closing window](https://github.com/vim/vim/commit/816968defc8ae79eb7e2319e991e74661be8d750)
- [8.0.1161 : popup menu drawing problem when resizing terminal](https://github.com/vim/vim/commit/a5e6621aadadf78c7b344e93a4b328788076f14c)
- [8.0.1162 : shared script for tests cannot be included twice](https://github.com/vim/vim/commit/660b85e39a48a9f2e3bd7e4d8b97d0a1b4842997)
- [8.0.1163 : popup test is flaky](https://github.com/vim/vim/commit/c79977a437d91306d576fb59e490601409503303)
- [8.0.1164 : changing StatusLine highlight does not always work](https://github.com/vim/vim/commit/65ed136844fbaffdd473903ed841c944600234dc)
- [8.0.1165 : popup test is still flaky](https://github.com/vim/vim/commit/f52c38315669f85bbcf3bd74c590148bf588f6c6)
- [8.0.1166 : :terminal doesn't work on Mac High Sierra](https://github.com/vim/vim/commit/c958b316386fa44116a02bfbd5859edc54a4a7e0)
- [8.0.1167 : Motif: typing in terminal window is slow](https://github.com/vim/vim/commit/3a497e1a414dc44b3df6a6fca60838ecd8ff0ca6)
- [8.0.1168 : wrong highlighting with combination of match and 'cursorline'](https://github.com/vim/vim/commit/0aa398f55a327282c70f56e0bac2dcb9521da378)

## 新着スクリプト

- [vim-scratchpad : toggle a scratchpad](https://vim.sourceforge.io/scripts/script.php?script_id=5604)
- [vim-Verdin : An omni-complete function for Vim script](https://vim.sourceforge.io/scripts/script.php?script_id=5605)
- [MarkdownCopy : Copy Markdown to clipboard as HTML](https://vim.sourceforge.io/scripts/script.php?script_id=5606)
- [purge&#x5f;undodir.vim : Purging useless undofile on exit](https://vim.sourceforge.io/scripts/script.php?script_id=5607)
- [vim.emacs : vim.emacs is a vim plugin to evaluate elisp expressions in Emacs.](https://vim.sourceforge.io/scripts/script.php?script_id=5608)
- [qcvim : Highlighter for quantum chemistry input files](https://vim.sourceforge.io/scripts/script.php?script_id=5609)
- [makery.vim : A plugin for managing makeprgs.](https://vim.sourceforge.io/scripts/script.php?script_id=5610)
- [verilog-instance.vim : Create SystemVerilog port instantiation from port declaration.](https://vim.sourceforge.io/scripts/script.php?script_id=5611)
- [highlight-groups.vim : Add words in highlight groups on the fly.](https://vim.sourceforge.io/scripts/script.php?script_id=5612)
- [WORDComplete : Insert mode completion that completes an entire sequence of non-blank characters](https://vim.sourceforge.io/scripts/script.php?script_id=5613)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](https://vim.sourceforge.io/scripts/script.php?script_id=273) (1138)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](https://vim.sourceforge.io/scripts/script.php?script_id=1658) (756)
3. [wombat256.vim : Wombat for 256 color xterms](https://vim.sourceforge.io/scripts/script.php?script_id=2465) (617)
4. [AutomaticLaTeXPlugin : Background compilation, completion, bib serch, toc and other nice features.](https://vim.sourceforge.io/scripts/script.php?script_id=2945) (454)
5. [OmniCppComplete : C/C++ omni-completion with ctags database](https://vim.sourceforge.io/scripts/script.php?script_id=1520) (372)
6. [molokai : A port of the monokai scheme for TextMate](https://vim.sourceforge.io/scripts/script.php?script_id=2340) (364)
7. [python.vim : Enhanced version of the python syntax highlighting script](https://vim.sourceforge.io/scripts/script.php?script_id=790) (347)
8. [c.vim : C/C++ IDE -- Write and run programs. Insert statements, idioms, comments etc.](https://vim.sourceforge.io/scripts/script.php?script_id=213) (307)
9. [nginx.vim : initial version](https://vim.sourceforge.io/scripts/script.php?script_id=1886) (298)
10. [winmanager : A windows style IDE for Vim 6.0](https://vim.sourceforge.io/scripts/script.php?script_id=95) (289)

## vim-jp issues

Open : 259 (+9) | Closed : 844 (+16)

- [Issue #1079 : Windows の system("echo 1 > test") で test が出力されない。](https://github.com/vim-jp/issues/issues/1079)
- [Issue #1080 : foldmethod=expr なバッファに setbufline() 使った場合の挙動がおかしい](https://github.com/vim-jp/issues/issues/1080)
- [Issue #1081 : vim-fzで端末Vimが落ちる（SEGV）](https://github.com/vim-jp/issues/issues/1081)
- [Issue #1082 : warnings with the over-the-spot patch](https://github.com/vim-jp/issues/issues/1082)
- [Issue #1083 : iminsert=0 imsearch=0 をデフォルトにしたい](https://github.com/vim-jp/issues/issues/1083)
- [Issue #1084 : 行の真ん中を示すモーション](https://github.com/vim-jp/issues/issues/1084)
- [Issue #1085 : 絵文字のカーソル位置がずれる+gvimだと表示が欠ける](https://github.com/vim-jp/issues/issues/1085)
- [Issue #1086 : 一部の絵文字幅が1を返す問題](https://github.com/vim-jp/issues/issues/1086)
- [Issue #1087 : &#x60;vim.exe -&#x60; で文字化け](https://github.com/vim-jp/issues/issues/1087)
- [Issue #1088 : \[doc\] empty() のドキュメントが不親切](https://github.com/vim-jp/issues/issues/1088)
- [Issue #1089 : \[レビュー依頼\] xtermとの互換性が低い端末でゴミが表示される問題へのパッチ](https://github.com/vim-jp/issues/issues/1089)
- [Issue #1090 : v8.0.1102以降で :terminal の地の文(?)の文字が意図しない色で表示される](https://github.com/vim-jp/issues/issues/1090)
- [Issue #1091 : --enable-perlinterpを指定してVimをビルドするとエラーでビルドに失敗する](https://github.com/vim-jp/issues/issues/1091)
- [Issue #1092 : encoding=utf-8 で balloonexpr が文字化けする。](https://github.com/vim-jp/issues/issues/1092)
- [Issue #1093 : 'statusline'内でhighlightを設定して :vsplit 後にスクロールするとウィンドウが更新されないことがある](https://github.com/vim-jp/issues/issues/1093)
- [Issue #1094 : gn コマンドは ignorecase オプションを考慮してほしい](https://github.com/vim-jp/issues/issues/1094)
- [Issue #1095 : FEAT&#x5f;WINDOWS 削除のおかげで既存パッチの多くがマージ出来なくなっている](https://github.com/vim-jp/issues/issues/1095)
- [Issue #1096 : MacでVimの端末を開くと背景色の表示が灰色になる](https://github.com/vim-jp/issues/issues/1096)
- [Issue #1097 : WinBar をクリックするとマウス選択になって使い辛い](https://github.com/vim-jp/issues/issues/1097)
- [Issue #1098 : :rubyfile を実行中に例外が発生すると、segmentation fault する](https://github.com/vim-jp/issues/issues/1098)
- [Issue #1099 : tmuxでWindowを切り替えると直前に入力したシングルバイト文字が消える](https://github.com/vim-jp/issues/issues/1099)
- [Issue #1100 : Windows の保存確認ダイアログのフォントが変](https://github.com/vim-jp/issues/issues/1100)
- [Issue #1101 : OSX 限定で if&#x5f;ruby な vim で :ruby p "a" した時落ちる](https://github.com/vim-jp/issues/issues/1101)
- [Issue #1102 : Windowsの:terminal上でclsして、\<C-w>Nでノーマルモードに移るとclsでクリアした内容が残っている。](https://github.com/vim-jp/issues/issues/1102)
- [Issue #1103 : :terminal 経由でインタラクティブな操作をするツールで enter が効かない](https://github.com/vim-jp/issues/issues/1103)
