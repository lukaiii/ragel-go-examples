package time_parser

import (
  "fmt"
)


%%{
  machine time_parser;
  write data;
}%%

func ParseTime(data string) (year, month, day string) {
  cs, p, pe, eof := 0, 0, len(data), len(data)
  year, month, day = "", "", ""
  year_begin_index, year_end_index := -1, -1
  month_begin_index, month_end_index := -1, -1
  day_begin_index, day_end_index := -1, -1
  %%{
    action see_neg { neg = true }
    action add_digit { val = val * 10 + (int64(fc) - '0') }

    action begin_year {
        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    }
    action end_year {
        // fmt.Println("end year")
        year_end_index = p
        // fmt.Println(year_end_index)
    }
    action set_year {
        fmt.Println(year_begin_index, year_end_index)
        if year_begin_index < year_end_index {
            year = data[year_begin_index: year_end_index]
        }
    }


    action begin_month {
        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    }
    action end_month {
        // fmt.Println("end month")
        month_end_index = p
        // fmt.Println(month_end_index)
    }
    action set_month {
        if month_begin_index < month_end_index {
            month = data[month_begin_index: month_end_index]
        }
    }

    action begin_day {
        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    }
    action end_day {
        // fmt.Println("end day")
        day_end_index = p
        // fmt.Println(day_end_index)
    }
    action set_day {
        if day_begin_index < day_end_index {
            day = data[day_begin_index: day_end_index]
        }
    }

    Digit = [0-9] | "一" | "二" | "三" | "四" | "五" | "六" | "七" | "八" | "九" | "十" | "百" | "千";

    year_description =
        (
            ("今" | "去" | "前" | "大前") >begin_year %end_year "年" @set_year
        )
        |
        (
            Digit+ >begin_year %end_year "年" @set_year
        )
    ;

    month_description =
        (
            ("上个" | "这个" | "上上个" | "前个" | "上上上个") >begin_month %end_month "月" @set_month
        )
        |
        (
            Digit+ >begin_month %end_month ("月" "份"?) @set_month
        )
    ;

    month_unclear_day =
        month_description
        (
            "初" | "中旬" | "末" | "半" | "上旬" | "下旬"
        ) >begin_day %end_day %eof end_day %set_day %eof set_day
        any*
    ;

    day_description =
        (
            Digit+ >begin_day %end_day ("日" | "号") @set_day
        )
    ;

    only_day_description =
        ("今" | "昨" | "前" | "大前") >begin_day %end_day "天" @set_day
        # |
        # Digit+ >begin_day %end_day ("天前" | "天以前" | "天之前")
    ;

    main :=
        any*
        (
            year_description
            |
            month_description
            |
            year_description "的"? month_description
            |
            month_unclear_day
            |
            year_description "的"? month_unclear_day
            |
            day_description
            |
            month_description "的"? day_description
            |
            year_description "的"? month_description "的"? day_description
            |
            only_day_description
        )
        any*
    ;

    write init;
    write exec;
  }%%

    if cs < time_parser_first_final {
        fmt.Println("time parse error")
    } else {
    }
    fmt.Println(year_begin_index, year_end_index)
    // fmt.Println(day_begin_index, day_end_index)
    fmt.Println(year+"年 "+month + "月 " + day + "日")

    return year, month, day
}
