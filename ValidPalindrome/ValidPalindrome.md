1、题目名称

Valid Palindrome（回文字符串）

2、题目地址

https://leetcode.com/problems/valid-palindrome/

3、题目内容

英文：Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

中文：给出一个字符串，只考虑字母和数字，判断该字符串是否为回文字符串

例如：

"A man, a plan, a canal: Panama" 是一个回文字符串

"race a car" 不是一个回文字符串

4、题目分析

如果要判断一个字符串是否是回文字符串，可以两个变量，从两侧开始比较有效的字符是否相等。这里有效的字符是指 'a'-'z'、'A'-'Z'、'0'-'9' 这三个区间内的字符。