---
layout: post
category: vimmagazine
title: Vim Magazine 6 月号
---

## 話題

- [Vim Advent Calendar 2012](http://atnd.org/events/33746)

  Vim Advent Calendar 2012 まだやってます。

- [2013 Emacs/Vim温泉（希望調査）](http://connpass.com/event/2773/)

  Emacs/Vimなど、エディタを肴にリゾートしてみます。
  何かを作ってみたり、発表してみたり、ただ単に観光したり温泉入るなど、何でも良し。
  九重・湯布院周辺 9月か10月頃

## 今月の新機能

- 7.3.1129: `:syntime` が追加されました。

- 7.3.1164: `screenattr({row}, {col})` と `screenchar({row}, {col})` が追加されました。

- 7.3.1248: `'imactivatefunc'`、`'imstatusfunc'` オプションが追加されました。

- 7.3.1261: `:map-nowait` が追加されました。

## リリース情報

- [7.3.1079 : test 87 fails](https://github.com/vim/vim/commit/96c7dfd806c29bb2be225b368499911700e89f71)
- [7.3.1080 : test 86 fails](https://github.com/vim/vim/commit/ca982c8d4527262569abbbcecc352204d3b0ab15)
- [7.3.1081 : compiler warnings on 64-bit Windows](https://github.com/vim/vim/commit/99dc19d1cc4d68694f168ccee468bbab908b9c58)
- [7.3.1082 : new regexp engine: Problem with \\@= matching](https://github.com/vim/vim/commit/14f55c6f67b55b693d0689beb66af90f53589a77)
- [7.3.1083 : new regexp engine: Does not support \\%^ and \\%$](https://github.com/vim/vim/commit/4b78063368eca61c8c8991d3c9afc18d5a474b82)
- [7.3.1084 : new regexp engine: only accepts up to \\{,10}](https://github.com/vim/vim/commit/54dafde1707b253036f42e6e3478c2e9d61bf4b9)
- [7.3.1085 : new regexp engine: Non-greedy multi doesn't work](https://github.com/vim/vim/commit/36b3a011d38dc96a85ec4451ae6a3ed9cb1ac9e5)
- [7.3.1086 : old regexp engine accepts illegal range, new one doesn't](https://github.com/vim/vim/commit/75d7a069206e8f3805d30f4df4655cb18d022fa4)
- [7.3.1087 : a leading star is not seen as a normal char when \\{} follows](https://github.com/vim/vim/commit/3737fc1c30a559c82c57b5ce6aedae76cbeabaef)
- [7.3.1089 : tests 86 and 87 fail on MS-Windows](https://github.com/vim/vim/commit/0418811869a64add42b0f2a37eff911dd550eb28)
- [7.3.1090 : new regexp engine does not support \\z1 .. \\z9 and \\z(](https://github.com/vim/vim/commit/efb23f26e8da3edb9fe8c980f7a0f27906f98bf2)
- [7.3.1091 : new regexp engine: no error for \\z1 or \\z( when not working](https://github.com/vim/vim/commit/5de820b916c38d95f6af7f4315f827764d587497)
- [7.3.1092 : regexp debugging: can't build, debug output has wrong pattern](https://github.com/vim/vim/commit/69afb7bf0aef4e9b39acbdb7849dd9c44b1345c5)
- [7.3.1093 : new regexp engine: empty \\1 skips a character](https://github.com/vim/vim/commit/b122e9798e058645ab3936eb12e02bd61403ac33)
- [7.3.1094 : new regexp engine: Attempts to match "\^" at every character](https://github.com/vim/vim/commit/307aa16a6958e5aa76c7c9f6e3f1b47faf8c50d2)
- [7.3.1095 : compiler warnings for shadowed variables](https://github.com/vim/vim/commit/525666f28201c313825065d49b98606c8cade457)
- [7.3.1096 : Python: popitem() was not defined in a standard way](https://github.com/vim/vim/commit/de71b5658f780e1fac29b1ef8b58af3dd42e65ad)
- [7.3.1097 : Python: a few recently added items are not documented](https://github.com/vim/vim/commit/bc41196bdc94f5025a5d0bee2dcbb3c63aef15b8)
- [7.3.1098 : Python: Possible memory leaks](https://github.com/vim/vim/commit/4f2109d782f91f7341a51644e3629ad1cbaea991)
- [7.3.1099 : Python: changing directory causes problems](https://github.com/vim/vim/commit/f4258308e288c82c30d35cb2d0e045e60bfdb21c)
- [7.3.1100 : Python: a few more memory problems](https://github.com/vim/vim/commit/78b595743931457c054156f1c86aa1fae8ba42f5)
- [7.3.1101 : configure doesn't find Python 3 on Ubuntu 13.04](https://github.com/vim/vim/commit/ac499e311677a58b90eadc1f5558e72da5eb1c5f)
- [7.3.1102 : completion of ":py3do" and ":py3file" fails after ":py3"](https://github.com/vim/vim/commit/9f5d60072301711ac1175fa679676d0f7e40e60a)
- [7.3.1103 : new regexp engine: overhead in saving and restoring](https://github.com/vim/vim/commit/f6de032afe1c4583cd4df585a7d46f93740357dc)
- [7.3.1104 : new regexp engine does not handle "~"](https://github.com/vim/vim/commit/f18fb7af7504a47e59ce85ddd1e50b56651a80d3)
- [7.3.1105 : new regexp engine: too much code in one function; dead code](https://github.com/vim/vim/commit/f46da70603be5280555da3bc8b5519b807e93afd)
- [7.3.1106 : new regexp engine: save/restore lastlist takes a lot of time](https://github.com/vim/vim/commit/dd2ccdf6eaf3de0a5e9a4bb3b9e08dd55168c629)
- [7.3.1107 : compiler warnings for unused variables](https://github.com/vim/vim/commit/01d89dddfa710f8a014f9bb85ffff4ad35451dc9)
- [7.3.1108 : error message for os.fchdir()](https://github.com/vim/vim/commit/e9056b121d6e9751bc0f802c09d4697a6ce4bdd5)
- [7.3.1109 : building on MS-Windows doesn't see changes in if\_py\_both.h](https://github.com/vim/vim/commit/dd9a4a44a9f73fea7305b4ff8ac0ccfbaea3f9aa)
- [7.3.1110 : new regexp matching: Using \\@= and the like can be slow](https://github.com/vim/vim/commit/a2d951086166fb0ece62fb1c6f9071237bf27f24)
- [7.3.1111 : nfa\_recognize\_char\_class() implementation is inefficient](https://github.com/vim/vim/commit/f811509c036306595be026f5e49dde170abb216c)
- [7.3.1112 : new regexp engine: \\%V not supported](https://github.com/vim/vim/commit/dacd7de3fd5edd8f3a41ba3c9c389b653b6a9e5d)
- [7.3.1113 : new regexp engine: \\%'m not supported](https://github.com/vim/vim/commit/044aa290c5f20afc871629acec9aae0dae367a01)
- [7.3.1114 : (after 7.3.1110) can't build without the syntax feature](https://github.com/vim/vim/commit/b76591ef4afd62431c7975d29ac22687d5a50ca2)
- [7.3.1115 : users don't like cursor line number with 'relativenumber' set](https://github.com/vim/vim/commit/5ebc09b4505b38c0b928c68769e10cda1226a2e4)
- [7.3.1116 : can't build without Visual mode](https://github.com/vim/vim/commit/78eae9aaf7d73ab676d7adecb77765834dccc460)
- [7.3.1117 : new regexp engine: \\%\[abc\] not supported](https://github.com/vim/vim/commit/d75799ab7f712738795919005046b15d6ac356f8)
- [7.3.1118 : match failure rate is not very specific](https://github.com/vim/vim/commit/e2b8cb3b15c47107c1b05d698ce267ed1ea05079)
- [7.3.1119 : flags in 'cpo' are searched for several times](https://github.com/vim/vim/commit/1cd3f2c4507ad242a6ed50270009f19fa49e5a7b)
- [7.3.1120 : crash when regexp logging is enabled](https://github.com/vim/vim/commit/5b84ddcceae72c14c3f608006d305f9a8e00fefc)
- [7.3.1121 : new regexp engine: adding states that are not used](https://github.com/vim/vim/commit/67604aed755c71f51114c0125eb49ba7bcfcf8c7)
- [7.3.1122 : new regexp engine: \\@> not supported](https://github.com/vim/vim/commit/8795374bd31a98e488e1ce293604f17eec33067b)
- [7.3.1123 : can't build tiny Vim on MS-Windows](https://github.com/vim/vim/commit/e60c2e5e1c8df565076c392c3052c34078fda777)
- [7.3.1124 : Python: Crash on MS-Windows when os.fchdir() is not available](https://github.com/vim/vim/commit/d4209d203ac77e8c9debcc152557d80b4357aaae)
- [7.3.1125 : error for using \\%V in a pattern in tiny Vim](https://github.com/vim/vim/commit/973fced7632fcbc2856aa842d5d872b6472f523c)
- [7.3.1126 : compiler warning for uninitialized variable](https://github.com/vim/vim/commit/c19b4b50a800c9866df4de0d30a25dc9d9af4939)
- [7.3.1127 : no error for using empty \\%\[\]](https://github.com/vim/vim/commit/2976c028ca65631825581dcd4e36b56ddc1ae82d)
- [7.3.1128 : every failure in the NFA engine is a syntax error](https://github.com/vim/vim/commit/cd2d8bb6ea55179d69aaf559942133ed8e93341e)
- [7.3.1129 : can't see what pattern in syntax highlighting is slow](https://github.com/vim/vim/commit/8a7f5a2d4379bdc16502c01456bb4dc5051ed965)
- [7.3.1130 : (after 7.3.1129) can't build with anything but huge features](https://github.com/vim/vim/commit/f75125528325b396556eee58e4b2cb9622da4f14)
- [7.3.1131 : new regexp engine is a bit slow](https://github.com/vim/vim/commit/188c57bcd180ea88911e3083920246c8b28b1ce7)
- [7.3.1132 : crash when debugging regexp](https://github.com/vim/vim/commit/6d3a5d755a71f6df472c82ed4e619a8497a75f14)
- [7.3.1133 : new regexp engine is a bit slow](https://github.com/vim/vim/commit/d89616ebb8c0f7c4b96c96f971e2bf9ac944dd44)
- [7.3.1134 : running test 49 takes a long time](https://github.com/vim/vim/commit/cc908adac8b11d6ae67aac3ccf8a8b27482f4676)
- [7.3.1135 : compiler warning for unused argument](https://github.com/vim/vim/commit/4e31296ffa9d6b0138a59d8abe0d2de6db93cb32)
- [7.3.1136 : ":func Foo" does not show attributes](https://github.com/vim/vim/commit/4cd92d5aae72098baed6e630f548c92dbe54c448)
- [7.3.1137 : new regexp engine: collections are slow](https://github.com/vim/vim/commit/417bad227ae3c6c930dc088d598f65205cce13d9)
- [7.3.1138 : new regexp engine: neglist no longer used](https://github.com/vim/vim/commit/8aca2e9befea1a0bea1c59b7cfda976e22c87862)
- [7.3.1139 : new regexp engine: negated flag is hardly used](https://github.com/vim/vim/commit/decd9540fd09465c9e39e1609e94676b7b962cea)
- [7.3.1140 : new regexp engine: can skip expensive match](https://github.com/vim/vim/commit/43e029841607bbc1e1f2cc14bec578eab326eef4)
- [7.3.1141 : Win32: Check for available memory is not reliable](https://github.com/vim/vim/commit/a238431302ff50f7db401248c4e66b8b298222c2)
- [7.3.1142 : memory leak in ":syntime report"](https://github.com/vim/vim/commit/45fc539fe048e037a8fc8dc8c043a781c30c9c73)
- [7.3.1143 : when mapping NUL it is displayed as an X](https://github.com/vim/vim/commit/3798519b9eb1f3c8c614c2d21743660a2a2cf947)
- [7.3.1144 : "RO" is not translated everywhere](https://github.com/vim/vim/commit/235840336312c73d72d30f38b2e628f98b62f353)
- [7.3.1145 : new regexp engine: addstate() is called very often](https://github.com/vim/vim/commit/f96d109827062fb9a361aa592f207d6c0a7f2ed1)
- [7.3.1146 : new regexp engine: look-behind followed by zero-width fails](https://github.com/vim/vim/commit/b1b284fd5d2dc11eec0b627663aa063236c466cd)
- [7.3.1147 : new regexp engine: regstart is only used to find first match](https://github.com/vim/vim/commit/87f764a891ff717cd84f361afa5e2220820431e3)
- [7.3.1148 : no command line completion for ":syntime"](https://github.com/vim/vim/commit/cd9c46265e4a12cf716187bc8188c7399797f806)
- [7.3.1149 : new regexp engine: Matching plain text could be faster](https://github.com/vim/vim/commit/473de61b0409f8f8c86585733f099f882122b280)
- [7.3.1150 : new regexpengine: Slow for look-behind match without a width](https://github.com/vim/vim/commit/e7766eeaa5d92cf727649923cfd3e42b1696df03)
- [7.3.1151 : new regexp engine: Slow for look-behind plus zero-width match](https://github.com/vim/vim/commit/1e02e6620be6f6ded23879dec344badf71a34ca8)
- [7.3.1152 : In tiny build ireg\_icombine is undefined](https://github.com/vim/vim/commit/a940aa6fc44c18e012bf66e1f2569c046132d509)
- [7.3.1153 : new regexp engine: look-behind matches can be very expensive](https://github.com/vim/vim/commit/2a4e98ac1ef8c86232f51f1c7079a192eecdea82)
- [7.3.1154 : new regexp\_nfa engine: Uneccessary code](https://github.com/vim/vim/commit/927d4a1fb58dcac856cb0959c895db899f692703)
- [7.3.1155 : MS-DOS: "make test" uses external rmdir command](https://github.com/vim/vim/commit/11354759e2c35616ea7a4daef0b05337846adf76)
- [7.3.1156 : compiler warnings](https://github.com/vim/vim/commit/4380d1ea239fe7f2546b7cad4ad0c424f0f5979a)
- [7.3.1157 : new regexp engine fails on "\\(\\\<command\\)\\@\<=.*"](https://github.com/vim/vim/commit/bcf4d178abef9336709b53516fbf0164ce5ebe09)
- [7.3.1158 : crash when running test 86](https://github.com/vim/vim/commit/0bdda37fb4b8118ec4bd797eddca9eaf5d3b381d)
- [7.3.1159 : the round() function is not always available](https://github.com/vim/vim/commit/a2e14fc764c306cc33fe174707756be70f883319)
- [7.3.1160 : mixing long and pointer doesn't always work](https://github.com/vim/vim/commit/570064cf21091a211d9de50dbfd23304cdf048ee)
- [7.3.1161 : Python: PyList\_SetItem() is inefficient](https://github.com/vim/vim/commit/c1ba10c7f63d65cb7ec3eb11932feb0ee636a033)
- [7.3.1162 : Python: Memory leaks](https://github.com/vim/vim/commit/f9c9b32bd124235136980749bee754cae99e164a)
- [7.3.1163 : not easy to load Python modules](https://github.com/vim/vim/commit/c09a6d6c0c9d0e9056816f518171864158c72076)
- [7.3.1164 : can't test what is actually displayed on screen](https://github.com/vim/vim/commit/9a773488a7113cd028151e71ab6c9cd43bf56972)
- [7.3.1165 : HP-UX compiler can't handle zero size array](https://github.com/vim/vim/commit/16619a235bae49c11d28ac462dc2d42fcfea44ba)
- [7.3.1166 : loading Python modules is not tested](https://github.com/vim/vim/commit/a9f2220176cc318ed076472ac41d00e70934bc83)
- [7.3.1167 : no good Python 3 configure check; may need -pthreads](https://github.com/vim/vim/commit/3c7ad013fbe9a2ff58e2b5e65c710397e3a0f25c)
- [7.3.1168 : Python "sane" configure checks give a warning message](https://github.com/vim/vim/commit/780c3e9b043dac4cbbe1ac900022ea891f7e2a3f)
- [7.3.1169 : new regexp engine: doing work while executing a pattern](https://github.com/vim/vim/commit/a2947e2bea0e6084a76d5ba7a5a1203da39e0f4b)
- [7.3.1170 : (after 7.3.1058) can'tuse a function reference as a string](https://github.com/vim/vim/commit/60bf1f58d0edc45582fa065c3e0f68ae0de637ee)
- [7.3.1171 : check for digits and ascii letters can be faster](https://github.com/vim/vim/commit/0ea4a6b94b6f7a13ef5027b43c36bda0836b51af)
- [7.3.1172 : Python 2: loading modules doesn't work well](https://github.com/vim/vim/commit/9f3685a527c9293b05c8ad7a6a4a7c2fa8b3b8b8)
- [7.3.1173 : Python 2 tests don't have the same output everywhere](https://github.com/vim/vim/commit/27610ed76c500cf680fdbac000d269e30dcba54c)
- [7.3.1174 : Python 2 and 3 use different ways to load modules](https://github.com/vim/vim/commit/81c40c507c69ab0c3aede3ee14a2ba76c20c4595)
- [7.3.1175 : using isalpha() and isalnum() can be slow](https://github.com/vim/vim/commit/2d473ab932b5c46a77715ea507de3990b70d96c1)
- [7.3.1176 : compiler warnings on 64 bit system](https://github.com/vim/vim/commit/edb07a2d889481525396eb76cd5bff4be592784e)
- [7.3.1177 : wasting memory on padding](https://github.com/vim/vim/commit/6395af8018c523da2339ecf60a457ba70494a9c4)
- [7.3.1178 : can't put all Vim config files together in one directory](https://github.com/vim/vim/commit/22971aaa4530d6e8cb2c9c0c19b4f291d8442d61)
- [7.3.1179 : waiting for character after buffer-local map match](https://github.com/vim/vim/commit/0825c00f14bb60202bcd0391bc3f0eba4d188aac)
- [7.3.1180 : path from cscope may not be valid after changing directory](https://github.com/vim/vim/commit/cab465a6d7a7d158c99b04ddc81650b468d82227)
- [7.3.1181 : wrong error message for 1.0\[0\]](https://github.com/vim/vim/commit/2a876e40cebe45c773302b5b9518d07100888d6b)
- [7.3.1182 : 'backupcopy' default on MS-Windows does not work for links](https://github.com/vim/vim/commit/12b559e7c3cba38d53261b925b441f2dffff2368)
- [7.3.1183 : Python tests 86 and 87 fail](https://github.com/vim/vim/commit/82850dd9cbb1dc30dea02695bdda1384cd47185b)
- [7.3.1184 : highlighting is sometimes wrong](https://github.com/vim/vim/commit/484d241a4abb4a7e892fbb1e3afb269890526e5a)
- [7.3.1185 : new regexp engine: no match with ^ after \\n](https://github.com/vim/vim/commit/b62bcd1e05a00bdacde11c52762cd48f6f0325f0)
- [7.3.1186 : Python 3: test 87 may crash](https://github.com/vim/vim/commit/5f87b23229b1d5cc5bad4a263e513e72a91fe8ae)
- [7.3.1187 : (after 7.3.1170) "s:" is recognized but "\<SID>" is not](https://github.com/vim/vim/commit/0c6633a7ea50873ce8f0d744ea5758a1eb494010)
- [7.3.1188 : newline characters messing up error message](https://github.com/vim/vim/commit/9d438d3b3862e6cdb4a7b02702f1f82d49bc177d)
- [7.3.1189 : (after 7.3.1185) highlighting is still wrong sometimes](https://github.com/vim/vim/commit/4c46b5e00116420edcd5a75ca97449a6f29bb0fb)
- [7.3.1190 : compiler warning for parentheses](https://github.com/vim/vim/commit/6e70736cbc3ed8bf7178ffa884e022d916d0c1bf)
- [7.3.1191 : Backreference to previous line doesn't work](https://github.com/vim/vim/commit/580abea48a0de604b40aca734b22f3ea7c0e7531)
- [7.3.1192 : valgrind reports errors when using backreferences](https://github.com/vim/vim/commit/c5089bb8ff4d9f15c11632d7486b9068e09ff202)
- [7.3.1193 : fail\_if\_missing not used for Python 3](https://github.com/vim/vim/commit/1612b1abe77c0cfff7d9d8f6e8b1b96b6f412546)
- [7.3.1194 : Yaml highlighting is slow](https://github.com/vim/vim/commit/44c71db7711e3faa7d94ed4b5f32985563e7e494)
- [7.3.1195 : compiler warning for unitialized variable](https://github.com/vim/vim/commit/4cff8faf0567209ab95b39b9cf7f3d555162392e)
- [7.3.1196 : old regexp engine does not match pattern with backref right](https://github.com/vim/vim/commit/141f6bb34124872d8676066e021bf899d4023c23)
- [7.3.1197 : ":wviminfo!" does not write history read from a viminfo file](https://github.com/vim/vim/commit/ff1806f8da8830ca2a528f2eaa39b3e85489da6d)
- [7.3.1198 : build error when using Perl 5.18.0 and dynamic loading](https://github.com/vim/vim/commit/9be6e21575e52f5d22839f660f8ca1032097b2f5)
- [7.3.1199 : error from evaluating 'foldexpr' is silently ignored](https://github.com/vim/vim/commit/dab38d516effcab23c8381545d4b1e6be38bce5b)
- [7.3.1200 : calling setline() from Insert mode messes up undo](https://github.com/vim/vim/commit/df9259abcee07f217e36f85f4fd77d059fb1bbe7)
- [7.3.1201 : on startup preview window may be used as the current window](https://github.com/vim/vim/commit/e66f06d819c2bbb09205a5c28e7ccf481f3ea9e7)
- [7.3.1202 : tags are not found in case-folded tags file](https://github.com/vim/vim/commit/335167994849a66a6635429fb42461ad79574b2f)
- [7.3.1203 : matches from matchadd() might be highlighted incorrectly](https://github.com/vim/vim/commit/4ce239b0b167cbbce1fbbaeced2133bcd69ee90e)
- [7.3.1204 : calling gettabwinvar() in 'tabline' cancels Visual mode](https://github.com/vim/vim/commit/d69497413f878fbe6db1cdae45171127281acabd)
- [7.3.1205 : logtalk.dict is not removed on uninstall](https://github.com/vim/vim/commit/926b5d3b302d0bcff58ede9b9aba6c21c8d2a13b)
- [7.3.1206 : inconsistent function argument declarations](https://github.com/vim/vim/commit/5ab9d98b9a5d3ccc0353a685d9a115058e53f6cf)
- [7.3.1207 : new regexp engine: no match found on "#if FOO"](https://github.com/vim/vim/commit/a50d02d98350a5d876ff0521ad87f3930cf63b5c)
- [7.3.1208 : compiler warnings on MS-Windows](https://github.com/vim/vim/commit/427d51c23f0c83a1e3bc40da072f2ad87f033370)
- [7.3.1209 : no completion for ":tabdo"](https://github.com/vim/vim/commit/70baa405f2b738f3827390aae11c8be1712e039d)
- [7.3.1210 : (after 7.3.1182) 'backupcopy' default on MS-Windows is wrong](https://github.com/vim/vim/commit/03e114b0f2c994d58cd744f9be1b498245bd8c08)
- [7.3.1211 : MS-Windows: ":hardcopy" does not work properly](https://github.com/vim/vim/commit/5246cd7a9bb514fd5c44f8deaf0b3ed889e1602f)
- [7.3.1212 : "make test" on MS-Windows does not report like Unix does](https://github.com/vim/vim/commit/6949eea7cc79ef35921f82411bf8439a079ca424)
- [7.3.1213 : can't build with Python and small features. ](https://github.com/vim/vim/commit/8c0e322f18c11f4cc2fb6ee2525d47374044a706)
- [7.3.1214 : missing declaration for init\_users() and realloc\_post\_list()](https://github.com/vim/vim/commit/01b626c26e2eadc32e15f43d3ef51423621adfdc)
- [7.3.1215 : compiler warning for function not defined](https://github.com/vim/vim/commit/06ae70d07c31debed5fa1b2224805e6276194a1b)
- [7.3.1216 : configure can't find Motif on Ubuntu](https://github.com/vim/vim/commit/6324c3b920ce9efa57884768fc6d9cb208f8e983)
- [7.3.1217 : new regexp engine: Can't handle \\%\[\[ao\]\]](https://github.com/vim/vim/commit/d79862599d8cdd2186480dda1cf990d4ca0134fe)
- [7.3.1218 : "make test" on MS-Windows does not clean all temporary files](https://github.com/vim/vim/commit/555d2a83402a4dd21f0b353c757b3cde4a4e47de)
- [7.3.1219 : no test for using \[\]\] inside \\%\[\]](https://github.com/vim/vim/commit/a0c85c7ae1be04bcf5850fc88e5e42ab96b518a2)
- [7.3.1220 : MS-Windows: When using wide font italic/bold are not included](https://github.com/vim/vim/commit/db250526bbeac519665605127c18aad3ce2853ed)
- [7.3.1221 : when build flags change "make distclean" causes an error](https://github.com/vim/vim/commit/a6cc03101e30d55d4039d539ed732bc02ffd909b)
- [7.3.1222 : cannot execute some tests from the src directly](https://github.com/vim/vim/commit/22f93c700f79577a395d28b186a900bae20fd5b9)
- [7.3.1223 : tests fail on MS-Windows](https://github.com/vim/vim/commit/d73895ec00b56e22bf4f95857aa88dd129e11d06)
- [7.3.1224 : Clang gives warnings on xxd](https://github.com/vim/vim/commit/fe70acb3767ef10badd2eeb9928988c39ebebae8)
- [7.3.1225 : compiler warnings when building with Motif](https://github.com/vim/vim/commit/3dbcd0c7ad93533f46981ebafe6710c1e03b1be2)
- [7.3.1226 : Python: duplicate code](https://github.com/vim/vim/commit/6c85e7f3be87fa4a669e8f080d527a3982e8e96c)
- [7.3.1227 : Python: inconsistent string conversion](https://github.com/vim/vim/commit/389a1793f40b149ed8553fb75e7ee40bc684d8c5)
- [7.3.1228 : Python: various inconsistencies and problems](https://github.com/vim/vim/commit/808c2bc8bf435917b792e172a22dc5f8c0e7066a)
- [7.3.1229 : Python: not so easy to delete/restore translating](https://github.com/vim/vim/commit/0bd80ccd853ddfa4cc2385b6ceb3b2610245a9ed)
- [7.3.1230 : Python: Exception messages are not clear](https://github.com/vim/vim/commit/c476e52fcaf9e1ad779add6368e4f46b2c0d881c)
- [7.3.1231 : Python: use of numbers not consistent](https://github.com/vim/vim/commit/141be8a585145d2217fb425739ccc121aed4a8d4)
- [7.3.1232 : Python: inconsistencies in variable names](https://github.com/vim/vim/commit/c4b99e0be76e0fbb9b43efb16f2d56607ce45de0)
- [7.3.1233 : various Python problems](https://github.com/vim/vim/commit/841fbd290790376561eaaaccaad9fc52f45ded89)
- [7.3.1234 : (after 7.3.1229) Python: Strings are not marked for translate](https://github.com/vim/vim/commit/6f1404f8a8d00f7aa9accf4fc8a5a48857454a4a)
- [7.3.1235 : inserting CTRL-\] differs between insert and command-line mode](https://github.com/vim/vim/commit/ede3e6383d0bc86c13f039e9013ff72e307937d2)
- [7.3.1236 : Python: WindowSetattr() missing support for NUMBER\_UNSIGNED](https://github.com/vim/vim/commit/dee2e315d786cbe9d5bba2d388fb72d96ad1a846)
- [7.3.1237 : Python: non-import errors not handled correctly](https://github.com/vim/vim/commit/7e85d3d432e799fc56126f6e3867b52539978343)
- [7.3.1238 : crash in Python interface on 64 bit machines](https://github.com/vim/vim/commit/d5e376eb3b11dbeca5478724731399d6c23c74a6)
- [7.3.1239 : can't build with Python and MSVC10](https://github.com/vim/vim/commit/c1c3d68da9d5a089949e8687b8783daf7ee2d416)
- [7.3.1240 : memory leak in findfile()](https://github.com/vim/vim/commit/df2bc27b52944fbb75b8203a22f3ebde172e9c79)
- [7.3.1241 : (after 7.3.1236) some test files missing from distribution](https://github.com/vim/vim/commit/c8c13a0307bd6a6b5d27875de03d01a3e3ff311c)
- [7.3.1242 : no failure when trying to use a number as a string](https://github.com/vim/vim/commit/daa270289b0eeb8467b1fea47e72d4727f2f7d28)
- [7.3.1243 : back references in look-behind match don't work](https://github.com/vim/vim/commit/f86c0b01a74fa3e5e37033ac037ef3af41d81a8d)
- [7.3.1244 : MS-Windows: confirm() dialog text may not fit](https://github.com/vim/vim/commit/f05d81102c263b746c344a71cd3aafdec96d9f38)
- [7.3.1245 : MS-Windows: confirm() dialog text may still not fit](https://github.com/vim/vim/commit/b052fe0e2c428fa1168b45a8f425a15bf25d4706)
- [7.3.1246 : wrong window layout when resizing with 'winfixheight' set](https://github.com/vim/vim/commit/b893ac2098e2b75d3138f6e91289892c1ef4b340)
- [7.3.1247 : pattern '\[ \]\\@!\\p\\%(\[ \]\\@!\\p\\)*:' does not always match](https://github.com/vim/vim/commit/a0169128d98669c753c899353127263e1ea289f9)
- [7.3.1248 : still have old hacking code for Input Method](https://github.com/vim/vim/commit/abab85a4994fec3e95ab79c5cdf94a7f893c1aeb)
- [7.3.1249 : modeline not recognized when using "Vim" instead of "vim"](https://github.com/vim/vim/commit/c14621eafb1e14cd41a77c69893b417abcf60e1a)
- [7.3.1250 : Python tests fail on MS-Windows](https://github.com/vim/vim/commit/877aa0081151a11c7a7f783038e949a8a4e0b2bb)
- [7.3.1251 : test 61 messes up viminf.](https://github.com/vim/vim/commit/9b846cb8ac0f1db4ffecdd2dabb0b99758228f74)
- [7.3.1252 : Win32: Gvim does not always find the toolbar bitmap file](https://github.com/vim/vim/commit/c9dcb427042eaecd95231682d559408bd5c7c46c)
- [7.3.1253 : still undo problem after using CTRL-R = setline()](https://github.com/vim/vim/commit/5737ca2350ffcc2afc1495f27ed746f6163bab47)
- [7.3.1254 : (after 7.3.1252) can't build without the multi-lang feature](https://github.com/vim/vim/commit/a5f5c8be8f0f5f9e6cbe7b9f3fed234f32cf3467)
- [7.3.1255 : Clang warnings when building with Athena](https://github.com/vim/vim/commit/6470de83f2be942dfbd05071a99a6f154e9f1d05)
- [7.3.1256 : can't build without eval or autocmd feature](https://github.com/vim/vim/commit/f0327f67c778f46cdf19d793c013f024f01ea818)
- [7.3.1257 : GNU gettext() ignores :lang language setting](https://github.com/vim/vim/commit/e3a0b53c2e8985e51174f01ee58f8746ea2c45a9)
- [7.3.1258 : using submatch() may crash Vim](https://github.com/vim/vim/commit/4d9ae21a342a3c81913e9f41b362778989b0c564)
- [7.3.1259 : no test for patch 7.3.1258](https://github.com/vim/vim/commit/560c52d32b3e509710e5e7fd620feb751e97689b)
- [7.3.1260 : completion gets incomplete command line in cmdline window](https://github.com/vim/vim/commit/b751546636c71f206733af1bc30415aaab9a0b42)
- [7.3.1261 : (after 7.3.1179) keymap causes global mapping to stop working](https://github.com/vim/vim/commit/72179e1bd0bbb1d1c3083c85bd9fb0a8d6928f7f)
- [7.3.1262 : crash and compilation warnings with Cygwin](https://github.com/vim/vim/commit/fe17e7640d242e717e9498158570fe3eea2a769b)
- [7.3.1263 : typo in short option name](https://github.com/vim/vim/commit/4a4607036553448cd03abcabc8fdc14e9e2f4687)
- [7.3.1264 : (after 7.3.1261) missing m\_nowait](https://github.com/vim/vim/commit/f3fe39276039f78beb7ce5a2ea21c445bee70545)
- [7.3.1265 : (after 7.3.1249) errors for allowing "Vim:" in a modeline](https://github.com/vim/vim/commit/630a730f9b91511f8da556dd5dc0d15f0a23f46a)
- [7.3.1266 : QNX: GUI fails to start](https://github.com/vim/vim/commit/a8e3580be61f582aad05d3154896e73b0c414efb)
- [7.3.1267 : MS-Windows ACL support doesn't work well](https://github.com/vim/vim/commit/27515921196e61034c18718812b6417f65d0337f)
- [7.3.1268 : ACL support doesn't work when when compiled with MingW](https://github.com/vim/vim/commit/39efa89afb10ec967bea92f02aa58c4109628e5f)
- [7.3.1269 : insert completion keeps wrong entry selected](https://github.com/vim/vim/commit/6b487147acc7d76754d7f041e009afbd1eeffb58)
- [7.3.1270 : using "Vp" in an empty buffer can't be undone](https://github.com/vim/vim/commit/10315b1201716a10d39477fc09e3eb89e944191f)

## 新着スクリプト

- [seoul256.vim : A low-contrast Vim color scheme based on Seoul Colors](http://www.vim.org/scripts/script.php?script_id=4605)
- [move\_until\_char\_changes.vim : Moves up or down the file until a different character appears under the cursor](http://www.vim.org/scripts/script.php?script_id=4606)
- [autoscroll.vim : Automatically scrolls the window slowly so you can read a file hands-free](http://www.vim.org/scripts/script.php?script_id=4607)
- [darknight256.vim : Just another dark color scheme](http://www.vim.org/scripts/script.php?script_id=4608)
- [Psearch : Vim plugin for parallel searching the pattern string](http://www.vim.org/scripts/script.php?script_id=4609)
- [vim-c-cr : Finish braces, brackets and parentheses with ctrl+enter in gvim](http://www.vim.org/scripts/script.php?script_id=4610)
- [racoon.vim : racoon config syntax](http://www.vim.org/scripts/script.php?script_id=4611)
- [aPen : like the Brackets](http://www.vim.org/scripts/script.php?script_id=4612)
- [vim-origami : A script to align, insert and delete foldmarkers conveniently](http://www.vim.org/scripts/script.php?script_id=4613)
- [vim-autopep8 : vim-autopep8 is a Vim plugin that runs the currently open file through autopep8.](http://www.vim.org/scripts/script.php?script_id=4614)
- [thumbnail.vim : A thumbnail-style buffer selector for Vim](http://www.vim.org/scripts/script.php?script_id=4615)
- [shootingstar : Like the magic * but start matching at the cursor position](http://www.vim.org/scripts/script.php?script_id=4616)
- [vim-cpp-enhanced-highlight : Improved C++ syntax highlight](http://www.vim.org/scripts/script.php?script_id=4617)
- [kink.vim : Vim Plugin for Kink Programming Language](http://www.vim.org/scripts/script.php?script_id=4618)
- [K.vim : K.vim features VIM with key-value instant search from external dictionary](http://www.vim.org/scripts/script.php?script_id=4619)
- [Luna : Simple colorscheme created for fun and night usage](http://www.vim.org/scripts/script.php?script_id=4620)
- [wir\_black : heavily modified ir\_black theme for 255 terminals](http://www.vim.org/scripts/script.php?script_id=4621)
- [muzzl.vim : Tango colors with a torte scheme](http://www.vim.org/scripts/script.php?script_id=4622)
- [cisco.vim : Script to highlight cisco configuration files](http://www.vim.org/scripts/script.php?script_id=4624)
- [tex.vim : Every symbol defined by the unicode-math package](http://www.vim.org/scripts/script.php?script_id=4626)
- [auto tags gen : Plugin for creating tags automatically (C/C++)](http://www.vim.org/scripts/script.php?script_id=4627)
- [bats.vim : Syntax highlighting for Bats - Bash Automated Testing System](http://www.vim.org/scripts/script.php?script_id=4628)
- [LGS : Laravel 4 Generators interface.](http://www.vim.org/scripts/script.php?script_id=4629)
- [vim-addon-errorformats : error-formats only](http://www.vim.org/scripts/script.php?script_id=4630)
- [vim-github-comment : Comment on GitHub straight from within Vim.](http://www.vim.org/scripts/script.php?script_id=4631)
- [tjs.vim : syntax file for kirikiri2 tjs script](http://www.vim.org/scripts/script.php?script_id=4632)
- [robo : A vim plugin for Android development.](http://www.vim.org/scripts/script.php?script_id=4633)
- [NumUtils : calculator with regex support](http://www.vim.org/scripts/script.php?script_id=4634)
- [vim-github-dashboard : Browse GitHub events in Vim](http://www.vim.org/scripts/script.php?script_id=4635)
- [TempKeyword : Highlight keyword temporarily.](http://www.vim.org/scripts/script.php?script_id=4636)
- [symfind : find symbol or file in your project](http://www.vim.org/scripts/script.php?script_id=4637)
- [newnoise.vim : Black and red, low contrast scheme](http://www.vim.org/scripts/script.php?script_id=4638)
- [vim-wakatime : Automatic time tracker using wakati.me](http://www.vim.org/scripts/script.php?script_id=4639)

## 月間ダウンロードランキング

1. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1951)
2. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (1894)
3. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (1376)
4. [pathogen.vim : Easy manipulation of 'runtimepath', 'path', 'tags', etc](http://www.vim.org/scripts/script.php?script_id=2332) (980)
5. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (859)
6. [bufexplorer.zip : Buffer Explorer / Browser](http://www.vim.org/scripts/script.php?script_id=42) (830)
7. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (731)
8. [snipMate : TextMate-style snippets for Vim](http://www.vim.org/scripts/script.php?script_id=2540) (725)
9. [github theme : A gvim colorscheme based on github's syntax highlighting](http://www.vim.org/scripts/script.php?script_id=2855) (720)
10. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (583)

## vim-jp issues

Open : 170 (+13) | Closed : 255 (+16)

- [Issue #397 : sort の比較関数を書きたくない。](https://github.com/vim-jp/issues/issues/397)
- [Issue #398 : 'scrolloff'オプションのカーソルジャンプ系コマンド専用版(例えば 'jumpoff') とかあってもいいかも](https://github.com/vim-jp/issues/issues/398)
- [Issue #399 : 7.3.919以降、*.ja が開けない](https://github.com/vim-jp/issues/issues/399)
- [Issue #400 : 新たな補完関数のリフレッシュ機能の提案](https://github.com/vim-jp/issues/issues/400)
- [Issue #401 : 補完の絞り込み時のチラツキを改善する(Vim標準の絞り込み)](https://github.com/vim-jp/issues/issues/401)
- [Issue #402 : if\_lua を使用するとVimがSEGV？](https://github.com/vim-jp/issues/issues/402)
- [Issue #403 : 特定のプラグイン使用時にVimが暴走する](https://github.com/vim-jp/issues/issues/403)
- [Issue #404 : プラグインの切り分け支援](https://github.com/vim-jp/issues/issues/404)
- [Issue #405 : select modeで文字挿入すると暗黙のうちにyankされる](https://github.com/vim-jp/issues/issues/405)
- [Issue #406 : :mapにおいて\<Nul>がXと表示される](https://github.com/vim-jp/issues/issues/406)
- [Issue #407 : :function コマンドに abort dict range が表示されない。](https://github.com/vim-jp/issues/issues/407)
- [Issue #408 : Windowsでenc=utf-8の時に印刷が文字化けする](https://github.com/vim-jp/issues/issues/408)
- [Issue #409 : Windowsでtest17が失敗する](https://github.com/vim-jp/issues/issues/409)
- [Issue #410 : Travis CIでVimのビルドテストをする](https://github.com/vim-jp/issues/issues/410)
- [Issue #411 : タブが閉じられるかで BufWinLeave の挙動が違う](https://github.com/vim-jp/issues/issues/411)
- [Issue #412 : ノーマルモードでのマルチバイト文字列の入力はそのまま挿入してほしい](https://github.com/vim-jp/issues/issues/412)
- [Issue #413 : re=1 でマルチバイト文字に対するPOSIXブラケットの結果がおかしい](https://github.com/vim-jp/issues/issues/413)
- [Issue #414 : make test を使い勝手良く](https://github.com/vim-jp/issues/issues/414)
- [Issue #415 : tablineオプションに指定した関数内でgettabwinvar()を使用しているとその後に作成したタブでビジュアルモードに入れなくなる](https://github.com/vim-jp/issues/issues/415)
- [Issue #416 : Python が片方だとテストに失敗する](https://github.com/vim-jp/issues/issues/416)
- [Issue #417 : msvc で失敗するテスト ](https://github.com/vim-jp/issues/issues/417)
- [Issue #418 : 相対行番号表示時に関数内のredrawがうまく動いていない？](https://github.com/vim-jp/issues/issues/418)
- [Issue #419 : コマンドラインモードで \<C-\]> がそのまま入力される](https://github.com/vim-jp/issues/issues/419)
- [Issue #420 : syntax/vim.vim](https://github.com/vim-jp/issues/issues/420)
- [Issue #421 : 7.3.1200以降でexpressionレジスタを使用するとundo履歴がおかしくなる](https://github.com/vim-jp/issues/issues/421)
- [Issue #422 : Kaoriya版で配布しているVimをhugeバージョンにして欲しい](https://github.com/vim-jp/issues/issues/422)
- [Issue #423 : confirm() でマルチバイト文字があるとウィンドウの幅がおかしい](https://github.com/vim-jp/issues/issues/423)
- [Issue #424 : 7.3.1236 以降、test87 が失敗する](https://github.com/vim-jp/issues/issues/424)
- [Issue #425 : Vim 7.3.1260以降でコマンドラインのユーザー定義補完が動作しない](https://github.com/vim-jp/issues/issues/425)

