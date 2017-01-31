---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2017 年 1 月号

---

## 話題

* [Meguro.vim #2](https://megurovim.connpass.com/event/49507/) が開催されます。

## 今月の新機能

*   8.0.0151: `system()` と `systemlist()` の第 2 引数にバッファ番号を渡せるようになりました。
*   8.0.0157: `:syntax spell` と `:syntax sync` で補間が効くようになりました。
*   8.0.0171: `js_decode()` がシングルクォート(')で囲まれた文字列をサポートしました。
*   8.0.0175: オフィシャルで頒布している [Windows 版 Gvim](https://github.com/vim/vim-win32-installer/releases)の言語設定が反映されない件を修正しました。
*   8.0.0179: `'formatprg'` がグローバルからグローバル or バッファローカルになりました。
*   8.0.0190: タグファイルからの重複タグの検出方法が線形検索からハッシュテーブル方式に変更になり、より高速になりました。
*   8.0.0210: [bracketed paste mode](http://www.xfree86.org/current/ctlseqs.html#Bracketed%20Paste%20Mode)をサポートしました。
*   8.0.0218: `:cexpr`, `:cgetexpr`, `:caddexpr` 等でコマンドライン補完が効くようになりました。
*   8.0.0251: Python 2 と Python 3 の両方で動作するスクリプトを書きやすくするための仕組みとして `'pyxversion'` や `:pyx` 等が追加されました。
*   8.0.0254: `assert_equal()` 等の assert 関数に `msg` を指定した場合でも未指定時のメッセージが連結されるようになりました。
*   8.0.0259: `:tabnext`, `:tabclose`, `:tabonly` 等が後方カウント指定をサポートしました。
