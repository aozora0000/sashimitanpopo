# 刺し身たんぽぽ
Terminalからパイプ入力されたストリームを、刺し身の上にたんぽぽを乗せるが如く人間の暖かみのある処理のストレスを軽減出来ます。

手処理が終わるとTerminalでキーを入力すると、次のストリームの処理が開始出来ます。

パイプで渡された文字列を、%sで置換したコマンドで実装出来ます

## 使用例
入力されたURL郡をキーを押す度にブラウザで表示する

```url.txt
https://www.google.co.jp
https://www.yahoo.co.jp
https://duckduckgo.com/
```

```shell
$ cat url.txt | sashimitanpopo "open %s"
```