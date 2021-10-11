# gotosakura

## ざっくり
amazon で見ている商品に関する [サクラチェッカー](https://sakura-checker.jp/) のページを開く.

## 対応ブラウザ
- Safari
- Google Chrome
- Brave Browser

## 使い方
1. Alfred の検索窓に `sakura` と入力する.
2. amazon で見ている商品に関するサクラチェッカーのページが開かれる.

## インストール手順
1. https://github.com/qawatake/gotosakura/releases/latest から `mdlink.alfredworkflow` をダウンロード.
2. ダウンロードしたファイルを開けば, Alfred が自動的に workflow を追加してくれる (はず).
3. Google Chrome あるいは Brave Browser を使用する場合, ↓が必要.
    1. ブラウザを起動.
    2. ツールバー: [View] -> [Developer] -> [Allow JavaScript from Apple Events]

## workflow を構成するファイル
- `main.sh`: シュルスクリプト. AppleScript と Go の橋渡し的な処理を行う.
- `gotosakura`: バイナリファイル. この workflow の中心的な役割を持つ. `GOOS=darwin go build` によって生成する.
  - `main.go`
  - `entity.go`
  - `go.mod`
- `appscript/url_from_*.scpt`: AppleScript. ブラウザから現在開いているページの URL を取得し, 標準出力に吐き出す.
- `appscript/frontmost_appname.scpt`: AppleScript. 最前面にあるアプリ名を取得し, 標準出力に吐き出す.
- `info.plist`: XML ファイル. Alfred Workflow の全体構成を記述したもの. Alfred によって自動的に生成される.