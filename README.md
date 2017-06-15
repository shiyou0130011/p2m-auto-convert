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
<table>
	<thead>
		<tr>
			<th>欄位</th>
			<th>解說</th>
		</tr>
	</thead>
	<tbody><tr>
		<td>wiki</td>
		<td>Mediawiki 的網址</td>
		</tr>
		<tr>
			<td>api</td>
			<td>Mediawiki 的 api 頁的網址，可參考 Mediawiki 的特殊頁面 &gt;版本 &gt; 入口 URL</td>
		</tr>
		<tr>
			<td>puki</td>
			<td>Pukiwiki 的網址</td>
		</tr>
		<tr>
			<td>account</td>
			<td>Mediawiki 的機器人的帳號</td>
		</tr>
		<tr>
			<td>password</td>
			<td>Mediawiki 的機器人的密碼</td>
		</tr>
	</tbody>
</table>

### 執行
本程式的參數如下：
<table>
	<thead>
		<tr>
			<th>參數</th>
			<th>類型</th>
			<th>說明</th>
		</tr>
	</thead>
	<tbody><tr>
		<td>–changetitle</td>
		<td>boolean</td>
		<td>
			是否轉換標題，例如 角色/ABC 轉成 ABC 這樣<br />
			默認是不轉換
		</td>
		</tr>
		<tr>
			<td>–ci</td>
			<td>boolean</td>
			<td>是否僅上傳圖片</td>
		</tr>
		<tr>
			<td>–cp</td>
			<td>boolean</td>
			<td>是否僅建立條目</td>
		</tr>
		<tr>
			<td>–title / –t</td>
			<td>string</td>
			<td>條目名稱</td>
		</tr>
	</tbody>
</table>

## 範例
本程式沒有GUI（圖形使用者介面），所以需以終端機執行

複製條目《Important Article》並同時上傳檔案
``` 
p2m --title "Important Article"
```

僅上傳條目《Important Article》裡的圖片
``` 
p2m --title "Important Article" --ci
```

僅複製條目《Important Article》，不上傳圖片
``` 
p2m --title "Important Article" --cp
```

複製條目《角色/John Smith》

此範例會建立條目《John Smith》和重定向條目《角色/John Smith》
```
p2m --title "角色/John Smith" --changetitle
```


複製條目《角色/John Smith》，但不建立重定向
```
p2m --title "角色/John Smith" 
```