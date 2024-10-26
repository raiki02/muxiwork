

```golang
s := "hello,世界"
len(s) = 12 
```

* 每个中文字符占3个字节 

```
1.for i, _ := range s {
		b := string(s[i])
		fmt.Print(" ", b)
	}
```

输出结果：` h e l l o , ä ç` 

```
2.for i := 0; i < len; i++ {
		c := string(s[i])
		fmt.Print(" ", c)
	}
```

输出结果：` h e l l o , ä ¸  ç`

* 以UTF-8字符为单位遍历， 中文字符被截断



```
for _, i := range s {  
		d := string(i)
		fmt.Print(" ", d)
	}
```

输出结果：` h e l l o , 世 界`

* i以rune值对应字符，包括中文



