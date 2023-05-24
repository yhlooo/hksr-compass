# hksr-compass

游戏“崩坏：星穹铁道”中“引航罗盘”的求解工具。

## 使用方式

求解一个引航罗盘：

```shell
hksr-compass solve COMPASS_EXPRESSION
```

其中 `COMPASS_EXPRESSION` 为罗盘表达式，其格式为 `{outerLocation}{outerSpeed},{middleLocation}{middleSpeed},{innerLocation}{innerSpeed}/{ringGroup1},{ringGroup2},{ringGroup3}`

其中：

- `{outerLocation}` `{middleLocation}` 和 `{innerLocation}`

  分别为外圈、中圈、内圈的初始位置，以正整数表示。指针从目标位置（即罗盘正左方向）沿顺时针方向旋转到当前位置所需旋转的角度处以 60 度。

  比如 `0` 表示目标位置， `3` 表示指针指向正右方向。因为一周是 360 度，因此有效范围是： 0-5

- `{outerSpeed}` `{middleSpeed}` 和 `{innerSpeed}` 分别为外圈、中圈、内圈的旋转速度（单次旋转的角度）

  以带符号的整数表示。单位为 60 度，符号表示旋转方向，正数表示顺时整旋转，负数表示逆时针旋转。

  比如 `-1` 表示每次逆时针旋转 60 度； `+2` 表示每次顺时针旋转 120 度。

- `{ringGroup1}` `{ringGroup2}` 和 `{ringGroup3}` 是三种旋转的圈的组合

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
