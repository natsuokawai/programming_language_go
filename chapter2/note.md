# プログラム構造
## 名前
- 変数、定数、型、文のラベル、パッケージの名前
- 名前は文字（Unicode）かアンダースコアで始まり、その後に任意の数の文字、数字、アンダースコアを続けられる
- 大文字と小文字は区別される
- 予約語（keyword）は名前としては使えない
  ```
  break        default      func         interface    select
  case         defer        go           map          struct
  chan         else         goto         package      switch
  const        fallthrough  if           range        type
  continue     for          import       return       var
  ```
  - https://golang.org/ref/spec#Keywords
- predeclaredな名前があるがこれらは予約語ではないので名前に使うことができる
  - true, nil, int, make など30種
- 関数内で宣言された名前はローカル、外で宣言された名前は同じパッケージに属するすべてのファイルから見える
- キャメルケース
  - 頭字後（ASCII）や頭文字語（HTML）は常に同じケース
    - htmlEscape、HTMLEscape、escapeHTML

## 宣言
- プログラムのエンティティに対する名前付け、性質の規定
- var: 変数
- const: 定数
- type: 型
- func: 関数

## 変数
- var宣言は特定の型の変数を作成、名前とその変数を結びつけ、初期値（与えられない場合はセロ値）を設定
  ```go
  var name type = expression
  ```
- ゼロ値
  - 数値: 0
  - ブーリアン: false
  - 文字列: ""
  - インターフェースと参照型: nil
- 型を省略することで異なる型の複数の変数を宣言できる
  ```go
  var i, j, k int
  var b, f = true, 2.3
  ```
- 省略変数宣言（short variable declaration)
  ```go
  name := expression
  ```
  - :=は宣言、=は代入
  - 省略変数宣言は左辺の変数がどれかが宣言済みの場合、それらの変数に対しては代入のように働く
  - すべての変数が宣言済みの場合はコンパイルエラーとなる。その場合は代入で書き換える
- ポインタ
  - 変数のアドレス（値が格納されている場所）
  ```go
  x := 1
  p := &x
  *p = 2 // x == 2
  ```
- new関数
  - new(T)はT型の変数の領域を動的に確保し、Tのゼロ値で初期化、*T型の値を返す
  - 以下の関数は同じ
  ```go
  func newInt() *int {
    return new(int)
  }

  func newInt() *int {
    var dummy int
    return &dummy
  }
  ```
## 代入
## 型宣言
## パッケージとファイル
## スコープ

