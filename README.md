# cats

Reads files,Output multi-column


## Usage

#### $ cats \<FILE\> \<FILE\> \<FILE\> ... 

option:  
  -c int  :columns count (default 80)  
  -n	    :add filename  

example

```
$ echo ABCDEFGHIJKLMNOPQRSTUVWXYZ >test1.txt
$ echo いろはにほへとちりぬるをわかよたれそつねならむうゐのおくやまけふこえてあさきゆめみしゑひもせすん >test2.txt
$ echo 春眠不覺曉 處處聞啼鳥夜來風雨聲 花落知多少 >test3.txt

$ cats test1.txt test2.txt 
ABCDEFGHIJKLMNOPQRSTUVWXYZ             |いろはにほへとちりぬるをわかよたれそつ
                                       |ねならむうゐのおくやまけふこえてあさき
                                       |ゆめみしゑひもせすん
                                       |
$ cats test1.txt test2.txt test3.txt
ABCDEFGHIJKLMNOPQRSTUVWXY|いろはにほへとちりぬるを |春眠不覺曉 處處聞啼鳥夜來
Z                        |わかよたれそつねならむう |風雨聲 花落知多少
                         |ゐのおくやまけふこえてあ |
                         |さきゆめみしゑひもせすん |
                         |                         |
```



## Installation
[Download](https://github.com/kurehajime/cats/releases)

or 

Build yourself(Go lang)

```
$ go install github.com/kurehajime/cats@latest
```

