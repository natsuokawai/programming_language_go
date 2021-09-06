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
## 代入
## 型宣言
## パッケージとファイル
## スコープ

