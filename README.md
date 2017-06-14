# p2m-auto-convert

自動讀取 Pukiwiki 的內容，並在 Mediawiki 上建立條目。

※注：本程式在轉換過程有針對 Komica wiki 的部分模板進行處理。

## 安裝
本程式由 Golang 撰寫。安裝前請先[安裝 Go Lang](https://golang.org/doc/install) 以及 [Git](https://git-scm.com/downloads)

之後 clone 本專案

```
git clone https://github.com/shiyou0130011/p2m-auto-convert.git
```

Windows 使用者請執行 build.bat； linux 使用者請執行 build.sh

## 使用
首先先修改 wiki.json，將之改為你的 pukiwiki 和 mediawiki 的資料
wiki.json 各欄位如下
| 欄位 | 解說 |
| --- | --- |
| wiki | Mediawiki 的網址 |
| api | Mediawiki 的 api 頁的網址，可參考 Mediawiki 的特殊頁面 >版本 > 入口 URL |
| puki | Pukiwiki 的網址 |
| account | Mediawiki 的機器人的帳號 |
| password | Mediawiki 的機器人的密碼 |

### 執行
本程式的參數如下：
| 參數 | 類型  | 說明 |
| --- | --- | --- |
| --changetitle |boolean |  是否轉換標題，例如 角色/ABC 轉成 ABC 這樣 |
| --ci | boolean | 是否僅上傳圖片 |
| --cp | boolean | 是否僅建立條目 |
| --title / --t | string | 條目名稱 |
