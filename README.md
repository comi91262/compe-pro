## 概要

競プロで書いたコード、スニペット, 競技中にコードを書くファイルを保存する

## 解法マップ(典型90より) (引き出し)

### 考察テクニック

- 制約からアルゴリズムを推測 (2, 32, 51, 63, 86)
- 特殊ケースを考える (54,74,83)
  - 特殊な制約に着目する
  - N = 1など、特殊なケースを考える
  - いろいろなケースを手計算で実験してみる
  - {a, b, c} -> {0, 1, 2} 置き換えて考える
- 同じものをまとめて考えて探索 (25,46)
- 部分問題に分解する (82,86)
- 答えへの貢献度を考える (39,66)
- x,yごとに独立して考える (70,86)
- bitごとに独立して考える (70,86)
- 条件の言い換え (49)
- 不要な状態を削る (23)
- パリティを考える (24,68)
- 両側から考える (9,13,60)
- 中央から考える (9,13,60)
- 後ろから考える (62)
- 何通りかの感覚を掴む (23,74,85)
- 前計算、クエリ先読み (4, 68)
- ものをグラフとして考える (49,62,71,88)
- 上界と下界を見積もる (48)
- マンハッタン距離は45回転 (36)
- ソートして貪欲法, DP (11,14)
- 階差を考える (64)
- 見かけ上の変化をメモ (44)
- 不変量に着目する (74)
  - 必ず1減る/操作によって変わらない
- 操作順序によらない (79)
- 余事象に着目する (17,65,84)
- 周期性と鳩の巣原理 (58,88)

### 全探索

- for (16, 55)
- bit全探索 (2, 32, 63, 86)
- 複雑な構造->再帰関数 (25, 71, 72, 88)
- 工夫 (16, 63, 81, 85)
  - 一文字固定
  - 状態数が少ない変量を全探索

### 二分探索

- 半分全列挙 -> マージ(二分探索) (51)
- 座標圧縮 (29, 34)
- 答えで二分探索 (1, 87)
- しゃくとり法 (34, 76)
- 三分探索、黄金分割探索 (53)

### 動的計画法

- 基本 (50)
- 状態DP (8, 73)
- 木DP (39, 73)
- BitDP (23, 45)
- ナップザックDP (11,37,56)
- 区間DP (19)
- 累積和DP (89)
- その他, 累積的な計算 (6, 84)
- DP復元 (56)
- LIS (60)
- グランディ数は排他的論理和が0で後手必勝 (31)
- 累積和, 二次元累積和 (10, 81)
- いもす法, 二次元いもす法 (28)

### グラフ理論

- 基礎, 木と二部グラフの性質 (78, 3, 26)
  - 多重/単純, 連結/非連結, 有向/無向, 循環/非循環
  - 奇数長の閉路を含まない
  - 最大マッチングが多項式時間ですむ
  - 木は必ず二部グラフ
- DFS, BFS (3, 26, 43, 59, 73)
- MST (49)
- LCA (35)
- ダイクストラ (13,43,87)
- フロー (40, 77)
- SCC (21)
- DAG, トポロジカルソート (59, 71)
  - トポロジカルソート -> DP
  - SCCでDAGにすることができる O(N+M)
- Union-Find (12, 68)

### 数学

- 最大公約数, 約数列挙 (22,38,75,85)
  - 素因数列挙の計算量はO(NloglogN)
- エラトステネスの篩 (30)
- N進数 (67)
- 累乗と二項係数 (15,65,69)
- 幾何 (9,18,41)
- 期待値の和, 積 (52,66,86)
- 調和級数, 数列の和 (15,30,82,84)
- FFTによる畳込み (65)
- 包除原理 (4,80)
- 行列累乗, 掃き出し (5,57)
- 特殊な倍数の性質 (42)
- 数列の和 (82)
  - n(n+1)(2n+1)/6
  - 2^n - 1 = 2^(n-1) + ... + 2 + 1
  - 1 / (1-p) = 1 + p + ...  (0 <= p < 1)
- 部分集合の部分集合は3^N 通り

### 競技プログラミング

- 誤差 (20, 41)
- コーナーケース (33,67)
  - 解法をしっかり証明する
  - 問題文の条件をしっかり読む
  - 制約の下限の方にも注意する
- オーバーフロー (38,55,82,85)
  - 計算順序を変える  a*b/d -> a/d*b 
- 定数倍に注意する (31,55,59)

### データ構造など

- ローリングハッシュ (47)
- 平方分割 (83)
- map/set/deque/bitset (27,61,68,83)
- BIT, セグメント木 (17,29,37)

## TODO, 未整理

- lowlink 橋検出
- Thorupのアルゴリズム
- Moのアルゴリズム
- 2Dセグメントツリー
- 全方位木DP
- ローリングハッシュ
- SA-IS
- 高速きたまさ法
- int64 の最大値: 10^19 > 9_223_372_036_854_775_807 > 9.0 * 10^18
- 40C20 = 137846528820 (1.0 * 10^11)
- a/b (mod p) = a * b^(p-2) (mod p)
