---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2017 年 7 月号

---

## 話題

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

* Patch 8.0.0693: 試験的に :terminal コマンドが追加されました (まだ大部分が動きません) (関連パッチ 8.0.0694, 8.0.0695, 8.0.0712, 8.0.0713, 8.0.0717)
* Patch 8.0.0716: `--clean` コマンドラインオプション、`'viminfofile'` オプション、`-u DEFAULTS` コマンドラインオプションが追加されました。
  * `--clean`: `-u DEFAULTS -i NONE` と同じです
  * `-u DEFAULTS`: defaults.vim のみ読み込まれます (vimrc:no, plugin:no, defaults.vim:yes)
  * `'viminfofile'`: 空文字以外の時は viminfo ファイル名を上書きします。`NONE` の時は viminfo ファイルが読み込まれません
* Patch 8.0.0718: job の端末への出力が表示できるようになりました
* Patch 8.0.0721: `:argedit` が複数の引数に対応しました

## Vimに関する脆弱性

* 8.0.0703 <http://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2017-11109> (関連パッチ 8.0.0706)
