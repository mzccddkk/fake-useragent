# fake-useragent

a simple useragent faker

## Installation

```go
go get github.com/mzccddkk/fake-useragent
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/mzccddkk/fake-useragent"
)

func main() {
	fmt.Printf("Random: %s\n", useragent.Random())
	fmt.Printf("Chrome: %s\n", useragent.Chrome())
	fmt.Printf("Edge: %s\n", useragent.Edge())
	fmt.Printf("IE: %s\n", useragent.IE())
	fmt.Printf("Firefox: %s\n", useragent.Firefox())
	fmt.Printf("Safari: %s\n", useragent.Safari())
	fmt.Printf("UC: %s\n", useragent.UC())
	fmt.Printf("Opera: %s\n", useragent.Opera())
	fmt.Printf("Android: %s\n", useragent.Android())
	fmt.Printf("IOS: %s\n", useragent.IOS())
	fmt.Printf("Windows: %s\n", useragent.Windows())
	fmt.Printf("Linux: %s\n", useragent.Linux())
	fmt.Printf("MacOS: %s\n", useragent.MacOS())
	fmt.Printf("MacOSX: %s\n", useragent.MacOSX())
	fmt.Printf("IPhone: %s\n", useragent.IPhone())
	fmt.Printf("IPad: %s\n", useragent.IPad())
	fmt.Printf("Mobile: %s\n", useragent.Mobile())
	fmt.Printf("Computer: %s\n", useragent.Computer())
}
```

## Output

```
Random: MobileSafari/604.1 CFNetwork/897.15 Darwin/17.5.0
Chrome: Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36
Edge: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.134 Safari/537.36 Edg/102.0.1245.50
IE: Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 4.0)
Firefox: Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.7) Gecko/20070914 Firefox/2.0.0.7
Safari: Mozilla/5.0 (iPhone; CPU iPhone OS 14_8_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.2 Mobile/15E148 Safari/604.1
UC: Mozilla/5.0 (Linux; U; Android 8.1.0; zh-CN; EML-AL00 Build/HUAWEIEML-AL00) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/57.0.2987.108 UCBrowser/11.9.4.974 UWS/2.13.1.48 Mobile Safari/537.36 AliApp(DingTalk/4.5.11) com.alibaba.android.rimet/10487439 Channel/227200 language/zh-CN
Opera: Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36 OPR/55.0.2994.37
Android: Mozilla/5.0 (Linux; Android 6.0.1; SM-G532M Build/MMB29T; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/55.0.2883.91 Mobile Safari/537.36
IOS: MobileSafari/604.1 CFNetwork/902.2 Darwin/17.7.0
Windows: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36
Linux: Mozilla/5.0 (X11; Linux i686; rv:2.0.1) Gecko/20100101 Firefox/4.0.1
MacOS: Safari/13605.2.8 CFNetwork/901.1 Darwin/17.6.0 (x86_64)
MacOSX: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/600.3.18 (KHTML, like Gecko) Version/7.1.3 Safari/537.85.12
IPhone: Mozilla/5.0 (iPhone; CPU iPhone OS 13_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/83.0.4103.88 Mobile/15E148 Safari/604.1
IPad: Mozilla/5.0 (iPad; CPU OS 5_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko ) Version/5.1 Mobile/9B176 Safari/7534.48.3
Mobile: Mozilla/5.0 (Linux; Android 7.1.1; Moto G (5S) Build/NPPS26.102-49-4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.109 Mobile Safari/537.36
Computer: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36
```

