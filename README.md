## 概要

競プロで書いたコード、スニペット, 競技中にコードを書くファイルを保存する

## テクニック集 (引き出し)

- 001: 答えで二分探索
- 002: 小さい制約は全探索を考えよう
- 003: 木の直径は最短距離計算を２回やる
- 004: 扱いやすい形にして前計算しよう
- 006: 辞書順最小は前から貪欲法
- 007: 要素の探索はソートして二部探索
- 008: 状態DPによる高速化
- 010: 区間の総和は累積和
- 012: 連携判定はunion-find
- 013: 各頂点への最短経路はダイクストラ法
- 014: ソートして貪欲法
- 016: 工夫した全探索
- 018: 三角関数を使いこなそう
- 020: 整数で処理をして誤差をなくそう
- 021: 強連結成分分解をしよう
- 022: 最大公約数はユークリッド互除法
- 024: パリティで考える
- _026: 二部グラフ: ２色で塗り分けできるグラフ
  - 奇数長の閉路を含まない
  - 最大マッチングが多項式時間ですむ
  - 木は必ず二部グラフ
- 027: mapを使いこなそう
- 028: 領域加算は二次元いもす法 
- 029: 座標圧縮で効率化/区間に対する処理はセグメント木
- 030: 素因数列挙の計算量はO(NloglogN)
  - 素数の逆数の和はO(loglogN)
- 032: 小さい制約は順列全探索 
- 033: コーナーケースに気をつけよう
  - 解法をしっかり証明する
  - 問題文の条件をしっかり読む
  - 制約の下限の方にも注意する
- 034: 単調性を利用したしゃくとり法
- 036: マンハッタン距離は45度回転
- 037: DPをセグメント木で高速化
- 038: オーバーフローに注意, 計算順序を変える  a*b/d -> a/d*b 
- 039: 主客転倒: ある点からの経路の合計を求めるではなく、辺をまたぐ経路がいくつあるかを考える
- 042: ９の倍数の性質: あまりが９で割れる
- 043: 拡張BFS/ダイクストラ 各点を拡張して状態をもたせる
- 044: 見かけ上の変化をメモ
- 046: 同じ意味のものをまとめて考える
- 048: 上界/下界を見積もる
- 050: 漸化式を求めてDPをする
- 051: 半分全列挙をしよう
- 052: 因数分解をしよう
- 055: 「定数倍」を見積もる
- 056: DP復元
- 058: 周期性を考える
- 060: 両側から考える / 最長増加部分列
- 061: dequeを知っていますか？
- 063: 変な制約に着目する / 状態数が少ない変量を全探索
- 064: 階差を考えよう
- 066: 期待値の線形性
- 067: N進法展開を理解しよう
- 068: クエリ先読み / std:set
- 075: O(sqrt(n))で素因数分解
- 076: 円環を列にして二倍にする
- 078: グラフの基本を知ろう 隣接リスト/隣接行列
- 079: 操作順序によらない
- 081: データを二次元にプロット/二次元累積和
- 082: 部分問題に分解する/数列の和の公式
  - n(n+1)(2n+1)/6
  - 2^n - 1 = 2^(n-1) + ... + 2 + 1
  - 1 / (1-p) = 1 + p + ...  (0 <= p < 1)

- aをbごとにわけると何グループできるか  (a + b - 1) / b
- i & (1 << j) == 0 => iのjビット目(0<=j)が0
- i & (1 << j) != 0 => iのjビット目(0<=j)が1
- Ceil(x/n) = (x+1)/n,  Floor(x/n) = x/n
- int64 の最大値: 10^19 > 9_223_372_036_854_775_807 > 9.0 * 10^18
- 40C20 = 137846528820 (1.0 * 10^11)

## 基本用語

- 木: 連結で閉路を持たない無向グラフ
- 動的計画法: 複数の部分問題に分解して、漸化式を立てることで答えを出す手法
