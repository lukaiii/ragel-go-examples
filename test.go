package main

import (
    "github.com/lukaiii/ragel-go-examples/time_parser"
)


func main() {
    time_parser.ParseTime("外婆去年的12月份")
    time_parser.ParseTime("去年的12月份")
    time_parser.ParseTime("前年8月份")
    time_parser.ParseTime("今年1月")
    time_parser.ParseTime("上个月")
    time_parser.ParseTime("八月")
    time_parser.ParseTime("1月")
    time_parser.ParseTime("大前年")
    time_parser.ParseTime("2018年")
    time_parser.ParseTime("9月")

    time_parser.ParseTime("9月2号开始交房租的")
    time_parser.ParseTime("去年的12月份中旬买的空调")
    time_parser.ParseTime("去年的12月份中旬")
    time_parser.ParseTime("前年8月份上映的电影")
    time_parser.ParseTime("2019年01月03日")

    time_parser.ParseTime("我昨天生病的")
    time_parser.ParseTime("9月2号开始交房租的")
    time_parser.ParseTime("去年的12月份中旬买的空调")
    time_parser.ParseTime("前年8月份上映的电影")
    time_parser.ParseTime("今年1月初上映的电影")
    time_parser.ParseTime("上个月末我生病的")
    time_parser.ParseTime("八月二十九日我生病的")
    time_parser.ParseTime("1月3号开始高烧")
    time_parser.ParseTime("2号订的机票")
}
