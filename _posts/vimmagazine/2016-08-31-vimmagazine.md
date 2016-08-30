---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2016 年 8 月号

---

## 話題

*   VimConf 2016 (11/05開催予定) の発表者が募集されました。

## 今月の新機能

*   7.4.2137: `funcref()` が追加されました。
*   7.4.2170: `timer_info()` が追加されました。
*   7.4.2180: `timer_stopall()` と `timer_pause()` が追加されました。
*   7.4.2189: shellのプロセス置き換え(例: `"<(cmd)"`) でエンコーディングの検出が行われるようになりました。
*   7.4.2200: `getqflist()`, `getloclist()`, `setqflist()`, `setloclist()` の引数に `{what}` が追加されました。
*   7.4.2201: `'signcolumn'` が追加されました。
*   7.4.2204: `getbufinfo()`, `getwininfo()`, `gettabinfo()` が追加されました。
*   7.4.2213: ハイライトグループ `EndOfBuffer` が追加されました。

*   7.4.2215: `getwininfo()` で `'buftype'` が `"quickfix"` の場合の種別(`loclist` or `quickfix`)が判別できるようになりました。
*   7.4.2220: `printf()` が `%s` で文字列以外の引数も出力できるようになりました。
*   7.4.2221: `printf()` が `%b` と `%B` で2進数形式をサポートしました。
*   7.4.2230: `'tagcase'` の値に `followscs` と `smart` が追加されました。
*   7.4.2236: `'langremap'` が追加されました。(`'langnoremap'` は後方互換性のために残ります)
*   7.4.2237: `:tab` のカウント指定に `.` と `$` および `+2` 等の相対番号指定ができるようになりました。
*   7.4.2244: `:filter` が追加されました。
*   7.4.2251: diffモードの最大バッファ数が `4` から `8` に変更されました。
