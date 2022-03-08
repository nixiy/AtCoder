# AtCoder

[![CodeQL](https://github.com/nixiy/AtCoder/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/nixiy/AtCoder/actions/workflows/codeql-analysis.yml)

---

# Golang

## スキャナのバッファ

`2*10^6`文字くらいの文字列が1行で飛んでくるケースが稀によくあるため、デフォルトバッファサイズを拡張しておかないと謎のWAが発生する
https://github.com/nixiy/AtCoder/blob/3ec27d52c1cb6ebe668d2a4501053ced8a35d237/go_template/template_mini/template_mini.go#L25-L29

ちなみにデフォルトのバッファサイズは以下のように65536バイトになっている
https://github.com/golang/go/blob/93200b98c75500b80a2bf7cc31c2a72deff2741c/src/bufio/scan.go#L76-L84

## Mathライブラリ

goのMathライブラリは基本的にfloat64で扱われるようになっており、 min, max, abs等利用頻度の高い関数をintを引数に取りintで返す関数を自作している方も少なくない。

しかし一度float64を経由する事でintでは扱えていた大きい数値がおかしくなる事がある(あった)ため、簡単な実装に置き換えている

## 実行時間計測

多言語だと単純にSystemCurrentTimeMillisを2点で変数に入れて引き算して出しがちだが、timeパッケージにSinceメソッドが用意されている。

```go
now := time.Now()

// 何らかの処理

fmt.Printf("経過: %vms\n", time.Since(now).Milliseconds())
```

## MapとSlice

非常に大きなmapに入力を取り、その後sliceに変換したらランタイムエラーが出てしまった。 冷静に考えると単純にKey,
Valueをペアで持ちたい場合、構造体配列で事足りるため、Mapが必要でない箇所で用いすぎないほうが良いかもしれない。
https://atcoder.jp/contests/abc229/submissions/29960138