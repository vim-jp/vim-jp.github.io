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
