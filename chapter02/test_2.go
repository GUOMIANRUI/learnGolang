package main

import "fmt"

func main() {
	// 统计一篇文章里面英文字母出现的次数
	// 字符（rune/byte） => int
	// for range 遍历  拿出来的类型是rune类型的   统计结果的key应该用rune类型
	// ascii  A - Z ... a - z
	// ch >= 'A' && ch <= 'Z' 是大写字母   ch >= 'a' && ch <= 'z' 是小写字母

	// var article string = ``
	article := `
	One Day
	Sometimes I lay under the moon
	And thank god I'm breathing
	Then I pray don't take me soon
	'Cause I am here for a reason
	Sometimes in my tears I drown
	But I never let it get me down
	So when negativity surrounds
	I know some day it'll all turn around
	Because
	All my life I've been waiting for
	I've been praying for
	For the people to say
	That we don't wanna fight no more
	There'll be no more war
	And our children will play
	One day one day one day
	One day one day one day
	It's not about win or lose
	Cause we all lose when they feed on the souls of the innocent
	Blood drenched pavement
	Keep on moving though the waters stay raging
	In this maze you can lose your ways your way
	It might drive you crazy but don't let it faze you no way
	No way
	Sometimes in my tears I drown
	I drown
	But I never let it get me down
	Get me down
	So when negativity surrounds
	Surrounds
	I know some day it'll all turn around
	Because
	All my life I've been waiting for
	I've been praying for
	For the people to say
	That we don't wanna fight no more
	There'll be no more war
	And our children will play
	One day one day one day
	One day one day one day
	One day this all will change
	Treat people the same
	Stop with the violence
	Down with the hate
	One day we'll all be free
	And proud to be under the same sun
	Singing songs of freedom like
	Why ohh
	One day one day
	Ohh why
	One day one day
	Ohh
	All my life I've been waiting for
	I've been praying for
	For the people to say
	That we don't wanna fight no more
	There'll be no more war
	And our children will play
	One day one day one day
	One day one day one day
	`
	article_stats := map[rune]int{}
	for _, ch := range article { // ch为ascii   rune是int32类型的
		if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' {
			article_stats[ch]++
		}
	}
	// fmt.Println(article_stats)
	for ch, count := range article_stats {
		fmt.Printf("%c:%d\n", ch, count) // 将int32类型转换成字符输出
	}
}
