# hksr-compass

游戏“崩坏：星穹铁道”中“引航罗盘”的求解工具。

## 使用方式

1. 获取 hksr-compass 二进制

   - 从源码构建

     ```shell
     git clone https://github.com/keybrl/hksr-compass.git
     cd hksr-compass
     go build github.com/keybrl/chatgpt-cli
     ```

   - 或下载源码并构建安装到 `${GOPATH}/bin`

     ```shell
     go install github.com/keybrl/chatgpt-cli@latest
     ```

   - 或从 [Releases](https://github.com/keybrl/hksr-compass/releases) 下载对应平台二进制

2. 运行以下命令以求解一个引航罗盘（ `hksr-compass` 替换为步骤 1 获取的二进制实际路径或 `${PATH}` 下的二进制名）：

   ```shell
   hksr-compass solve COMPASS_EXPRESSION
   ```

其中 `COMPASS_EXPRESSION` 为罗盘表达式，其格式为 `{oLoc}{oSpeed},{mLoc}{mSpeed},{iLoc}{iSpeed}/{rg1},{rg2},{rg3}`

其中：

- `{oLoc}` `{mLoc}` 和 `{iLoc}` 分别为外圈、中圈、内圈的初始位置

  以正整数表示。指针从目标位置（即罗盘正左方向）沿顺时针方向旋转到当前位置所需旋转的角度处以 60 度。

  比如 `0` 表示目标位置， `3` 表示指针指向正右方向。因为一周是 360 度，因此有效范围是： 0-5

- `{oSpeed}` `{mSpeed}` 和 `{iSpeed}` 分别为外圈、中圈、内圈的旋转速度（单次旋转的角度）

  以带符号的整数表示。单位为 60 度，符号表示旋转方向，正数表示顺时整旋转，负数表示逆时针旋转。

  比如 `-1` 表示每次逆时针旋转 60 度； `+2` 表示每次顺时针旋转 120 度。

- `{rg1}` `{rg2}` 和 `{rg3}` 是三种旋转的圈的组合

  可选值如下：

  - `o` 外圈单独转
  - `m` 中圈单独转
  - `i` 内圈单独转
  - `om` 或 `mo` 外圈和中圈一起转
  - `oi` 或 `io` 外圈和内圈一起转
  - `mi` 或 `im` 中圈和内圈一起转

比如

```shell
hksr-compass solve '0+1,4-4,0+2/oi,om,mi'
```

其输出结果为：

```
Compass:  0+1,4-4,0+2/mi,oi,om
Solution: mi2,oi4,om2
```

`mi2,oi4,om2` 即为罗盘问题的解决步骤。步骤间以 `,` 分割，每个步骤包含旋转的圈组合和旋转次数。

圈组合的可能值有：

- `o` 外圈单独转
- `m` 中圈单独转
- `i` 内圈单独转
- `om` 外圈和中圈一起转
- `oi` 外圈和内圈一起转
- `mi` 中圈和内圈一起转

比如 `mi2` 表示旋转中圈和内圈 2 次

## 致谢

1. [Maozu](https://github.com/maozu) 提供了解题思路
