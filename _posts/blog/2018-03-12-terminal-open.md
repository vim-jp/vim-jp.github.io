---
layout: post
category: blog
title: :terminal コマンドを終了する方法
---
# terminal バッファ強制終了の動作

これまで GVim で `:terminal` コマンドを実行した状態で画面の × ボタンをクリックすると、エラーが発生して GVim を終了する事が出来ませんでした。これは Vim の設計方針としてバッファを閉じる際にジョブを強制終了するかどうかが明確に決まっていなかった事が原因です。

[Patch 8.0.1593](https://github.com/vim/vim/commit/25cdd9c33b21ddbd31321c075873bb225450d2d2) にて × ボタンを押した時の動作 `:qall` に変更が加えられました。

バッファを閉じる際にはジョブを自動で終了するべきではないというスタンスには変わりはありませんが、× ボタンを押した場合のみ、今回追加された `term_setkill` 関数により内部的にジョブを自動終了する為のフラグ付けが行われます。

更新されたバッファと :terminal が複数開かれている状態で、× ボタンをクリックすると、全てを破棄するかを確認されます。全てを破棄を選択すると :terminal も閉じられます。

# TerminalOpen イベント

また [Patch 8.0.1596](https://github.com/vim/vim/commit/b852c3e64d319d6ec47dd780c8654ae095e1d8c2) にて TerminalOpen イベントも導入されました。 :terminal が開かれた際に固有の設定を施したい方は以下の様に設定する事が出来る様になりました。

```vim

function! s:my_terminal()
  " :terminal バッファ固有の設定
endfunction

augroup MyTerminal
  au!
  au TerminalOpen * call s:my_terminal()
augroup END
```

# ExitPre イベント

Vim プラグイン開発者はこれまで Vim の終了を検知する為に QuitPre イベントを使ってきました。しかしこの QuitPre イベントは、バッファの終了を確認する手前で発火される為、例えば channel を使ったソケット接続を一つだけ保持し終了時に切断したいといったニーズに使用するには少し無理がありました。

[Patch 8.0.1595](https://github.com/vim/vim/commit/12a96de430779b88795fac87a2be666d9f661d1e) にて ExitPre イベントが導入され、バッファが全て閉じられた後に発火するイベントが追加されました。Vim プラグインから Vim が ExitPre イベントをサポートしているかどうかを確認するには以下の様に実行します。

```vim
:echo exists("##ExitPre")
```

