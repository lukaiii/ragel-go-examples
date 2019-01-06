
//line time_parser/time_parser.rl:1
package time_parser

import (
  "fmt"
)



//line time_parser/time_parser.go:12
const time_parser_start int = 0
const time_parser_first_final int = 39
const time_parser_error int = -1

const time_parser_en_main int = 0


//line time_parser/time_parser.rl:11


func ParseTime(data string) (year, month, day string) {
  cs, p, pe, eof := 0, 0, len(data), len(data)
  year, month, day = "", "", ""
  year_begin_index, year_end_index := -1, -1
  month_begin_index, month_end_index := -1, -1
  day_begin_index, day_end_index := -1, -1
  
//line time_parser/time_parser.go:30
	{
	cs = time_parser_start
	}

//line time_parser/time_parser.go:35
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	case 86:
		goto st_case_86
	case 87:
		goto st_case_87
	case 88:
		goto st_case_88
	case 89:
		goto st_case_89
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
	case 93:
		goto st_case_93
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	}
	goto st_out
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		switch data[p] {
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
tr1:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line time_parser/time_parser.go:277
		switch data[p] {
		case 228:
			goto tr2
		case 229:
			goto tr7
		case 230:
			goto tr8
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
tr2:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line time_parser/time_parser.go:319
		switch data[p] {
		case 184:
			goto st3
		case 185:
			goto st30
		case 186:
			goto st31
		case 187:
			goto st32
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 128:
			goto st1
		case 131:
			goto st1
		case 137:
			goto st1
		case 138:
			goto st4
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 228:
			goto tr15
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
tr15:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line time_parser/time_parser.go:419
		switch data[p] {
		case 184:
			goto st6
		case 185:
			goto st30
		case 186:
			goto st31
		case 187:
			goto st32
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 128:
			goto st1
		case 131:
			goto st1
		case 137:
			goto st1
		case 138:
			goto st4
		case 170:
			goto st7
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr18
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
tr3:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line time_parser/time_parser.go:521
		switch data[p] {
		case 133:
			goto st9
		case 137:
			goto st14
		case 141:
			goto st17
		case 142:
			goto st22
		case 155:
			goto st25
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 171:
			goto st1
		case 173:
			goto st1
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
tr4:
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line time_parser/time_parser.go:586
		switch data[p] {
		case 152:
			goto st11
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 168:
			goto st12
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch data[p] {
		case 228:
			goto tr2
		case 229:
			goto tr26
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
tr26:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
//line time_parser/time_parser.rl:62

        // fmt.Println("end day")
        day_end_index = p
        // fmt.Println(day_end_index)
    
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line time_parser/time_parser.go:680
		switch data[p] {
		case 133:
			goto st9
		case 137:
			goto st14
		case 141:
			goto st17
		case 142:
			goto st22
		case 155:
			goto st25
		case 164:
			goto st27
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch data[p] {
		case 141:
			goto st15
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		switch data[p] {
		case 228:
			goto tr15
		case 229:
			goto tr29
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
tr29:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:28

        // fmt.Println("end year")
        year_end_index = p
        // fmt.Println(year_end_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
//line time_parser/time_parser.rl:62

        // fmt.Println("end day")
        day_end_index = p
        // fmt.Println(day_end_index)
    
	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line time_parser/time_parser.go:790
		switch data[p] {
		case 133:
			goto st9
		case 137:
			goto st14
		case 141:
			goto st17
		case 142:
			goto st22
		case 155:
			goto st25
		case 164:
			goto st27
		case 185:
			goto st26
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch data[p] {
		case 129:
			goto st1
		case 131:
			goto st1
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
tr5:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line time_parser/time_parser.go:871
		switch data[p] {
		case 153:
			goto st19
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		switch data[p] {
		case 190:
			goto st1
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
tr6:
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line time_parser/time_parser.go:926
		switch data[p] {
		case 191:
			goto st21
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		switch data[p] {
		case 153:
			goto st4
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch data[p] {
		case 187:
			goto st23
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		switch data[p] {
		case 228:
			goto tr2
		case 229:
			goto tr34
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
tr34:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:28

        // fmt.Println("end year")
        year_end_index = p
        // fmt.Println(year_end_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line time_parser/time_parser.go:1043
		switch data[p] {
		case 133:
			goto st9
		case 137:
			goto st14
		case 141:
			goto st17
		case 142:
			goto st22
		case 155:
			goto st25
		case 185:
			goto st26
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		switch data[p] {
		case 155:
			goto st1
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		switch data[p] {
		case 180:
			goto tr35
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
tr35:
//line time_parser/time_parser.rl:33

        fmt.Println(year_begin_index, year_end_index)
        if year_begin_index < year_end_index {
            year = data[year_begin_index: year_end_index]
        }
    
	goto st39
tr36:
//line time_parser/time_parser.rl:67

        if day_begin_index < day_end_index {
            day = data[day_begin_index: day_end_index]
        }
    
	goto st39
tr89:
//line time_parser/time_parser.rl:62

        // fmt.Println("end day")
        day_end_index = p
        // fmt.Println(day_end_index)
    
//line time_parser/time_parser.rl:67

        if day_begin_index < day_end_index {
            day = data[day_begin_index: day_end_index]
        }
    
	goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line time_parser/time_parser.go:1154
		switch data[p] {
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
tr45:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st40
tr90:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:62

        // fmt.Println("end day")
        day_end_index = p
        // fmt.Println(day_end_index)
    
//line time_parser/time_parser.rl:67

        if day_begin_index < day_end_index {
            day = data[day_begin_index: day_end_index]
        }
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//line time_parser/time_parser.go:1228
		switch data[p] {
		case 228:
			goto tr46
		case 229:
			goto tr51
		case 230:
			goto tr52
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
tr46:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st41
tr91:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:62

        // fmt.Println("end day")
        day_end_index = p
        // fmt.Println(day_end_index)
    
//line time_parser/time_parser.rl:67

        if day_begin_index < day_end_index {
            day = data[day_begin_index: day_end_index]
        }
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
//line time_parser/time_parser.go:1302
		switch data[p] {
		case 184:
			goto st42
		case 185:
			goto st77
		case 186:
			goto st78
		case 187:
			goto st79
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		switch data[p] {
		case 128:
			goto st40
		case 131:
			goto st40
		case 137:
			goto st40
		case 138:
			goto st43
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		switch data[p] {
		case 228:
			goto tr59
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
tr59:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line time_parser/time_parser.go:1402
		switch data[p] {
		case 184:
			goto st45
		case 185:
			goto st77
		case 186:
			goto st78
		case 187:
			goto st79
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		switch data[p] {
		case 128:
			goto st40
		case 131:
			goto st40
		case 137:
			goto st40
		case 138:
			goto st43
		case 170:
			goto st46
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		switch data[p] {
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr62
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
tr47:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st47
tr92:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:62

        // fmt.Println("end day")
        day_end_index = p
        // fmt.Println(day_end_index)
    
//line time_parser/time_parser.rl:67

        if day_begin_index < day_end_index {
            day = data[day_begin_index: day_end_index]
        }
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
//line time_parser/time_parser.go:1536
		switch data[p] {
		case 133:
			goto st48
		case 137:
			goto st53
		case 141:
			goto st56
		case 142:
			goto st61
		case 155:
			goto st64
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		switch data[p] {
		case 171:
			goto st40
		case 173:
			goto st40
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
tr48:
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st49
tr93:
//line time_parser/time_parser.rl:62

        // fmt.Println("end day")
        day_end_index = p
        // fmt.Println(day_end_index)
    
//line time_parser/time_parser.rl:67

        if day_begin_index < day_end_index {
            day = data[day_begin_index: day_end_index]
        }
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st49
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
//line time_parser/time_parser.go:1621
		switch data[p] {
		case 152:
			goto st50
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		switch data[p] {
		case 168:
			goto st51
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		switch data[p] {
		case 228:
			goto tr46
		case 229:
			goto tr70
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
tr70:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
//line time_parser/time_parser.rl:62

        // fmt.Println("end day")
        day_end_index = p
        // fmt.Println(day_end_index)
    
	goto st52
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
//line time_parser/time_parser.go:1715
		switch data[p] {
		case 133:
			goto st48
		case 137:
			goto st53
		case 141:
			goto st56
		case 142:
			goto st61
		case 155:
			goto st64
		case 164:
			goto st66
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		switch data[p] {
		case 141:
			goto st54
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		switch data[p] {
		case 228:
			goto tr59
		case 229:
			goto tr73
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
tr73:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:28

        // fmt.Println("end year")
        year_end_index = p
        // fmt.Println(year_end_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
//line time_parser/time_parser.rl:62

        // fmt.Println("end day")
        day_end_index = p
        // fmt.Println(day_end_index)
    
	goto st55
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
//line time_parser/time_parser.go:1825
		switch data[p] {
		case 133:
			goto st48
		case 137:
			goto st53
		case 141:
			goto st56
		case 142:
			goto st61
		case 155:
			goto st64
		case 164:
			goto st66
		case 185:
			goto st65
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		switch data[p] {
		case 129:
			goto st40
		case 131:
			goto st40
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
tr49:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st57
tr94:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:62

        // fmt.Println("end day")
        day_end_index = p
        // fmt.Println(day_end_index)
    
//line time_parser/time_parser.rl:67

        if day_begin_index < day_end_index {
            day = data[day_begin_index: day_end_index]
        }
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st57
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
//line time_parser/time_parser.go:1938
		switch data[p] {
		case 153:
			goto st58
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		switch data[p] {
		case 190:
			goto st40
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
tr50:
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
	goto st59
tr95:
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:62

        // fmt.Println("end day")
        day_end_index = p
        // fmt.Println(day_end_index)
    
//line time_parser/time_parser.rl:67

        if day_begin_index < day_end_index {
            day = data[day_begin_index: day_end_index]
        }
    
	goto st59
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
//line time_parser/time_parser.go:2013
		switch data[p] {
		case 191:
			goto st60
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		switch data[p] {
		case 153:
			goto st43
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		switch data[p] {
		case 187:
			goto st62
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		switch data[p] {
		case 228:
			goto tr46
		case 229:
			goto tr78
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
tr78:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:28

        // fmt.Println("end year")
        year_end_index = p
        // fmt.Println(year_end_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st63
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
//line time_parser/time_parser.go:2130
		switch data[p] {
		case 133:
			goto st48
		case 137:
			goto st53
		case 141:
			goto st56
		case 142:
			goto st61
		case 155:
			goto st64
		case 185:
			goto st65
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		switch data[p] {
		case 155:
			goto st40
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		switch data[p] {
		case 180:
			goto tr35
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		switch data[p] {
		case 169:
			goto tr36
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
tr62:
//line time_parser/time_parser.rl:46

        // fmt.Println("end month")
        month_end_index = p
        // fmt.Println(month_end_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st67
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
//line time_parser/time_parser.go:2247
		switch data[p] {
		case 152:
			goto st50
		case 156:
			goto st68
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		switch data[p] {
		case 136:
			goto tr38
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
tr38:
//line time_parser/time_parser.rl:51

        if month_begin_index < month_end_index {
            month = data[month_begin_index: month_end_index]
        }
    
	goto st69
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
//line time_parser/time_parser.go:2304
		switch data[p] {
		case 228:
			goto tr80
		case 229:
			goto tr81
		case 230:
			goto tr82
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
tr80:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st70
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
//line time_parser/time_parser.go:2346
		switch data[p] {
		case 184:
			goto st71
		case 185:
			goto st77
		case 186:
			goto st78
		case 187:
			goto st79
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		switch data[p] {
		case 128:
			goto st40
		case 131:
			goto st40
		case 137:
			goto st40
		case 138:
			goto st72
		case 139:
			goto st76
		case 173:
			goto st76
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		switch data[p] {
		case 228:
			goto tr59
		case 229:
			goto tr47
		case 230:
			goto tr86
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
tr86:
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st73
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
//line time_parser/time_parser.go:2438
		switch data[p] {
		case 151:
			goto st74
		case 152:
			goto st50
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		switch data[p] {
		case 172:
			goto st75
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		switch data[p] {
		case 228:
			goto tr91
		case 229:
			goto tr92
		case 230:
			goto tr93
		case 231:
			goto tr94
		case 232:
			goto tr95
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr90
		}
		goto tr89
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		switch data[p] {
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr86
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		switch data[p] {
		case 157:
			goto st40
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		switch data[p] {
		case 140:
			goto st40
		case 148:
			goto st40
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		switch data[p] {
		case 138:
			goto st80
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
		switch data[p] {
		case 228:
			goto tr46
		case 229:
			goto tr73
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
tr81:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st81
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
//line time_parser/time_parser.go:2641
		switch data[p] {
		case 133:
			goto st48
		case 136:
			goto st82
		case 137:
			goto st53
		case 141:
			goto st83
		case 142:
			goto st61
		case 155:
			goto st64
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		switch data[p] {
		case 157:
			goto st75
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		switch data[p] {
		case 129:
			goto st40
		case 131:
			goto st40
		case 138:
			goto st75
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
tr82:
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st84
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
//line time_parser/time_parser.go:2733
		switch data[p] {
		case 152:
			goto st50
		case 156:
			goto st85
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		switch data[p] {
		case 171:
			goto st75
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
tr51:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:28

        // fmt.Println("end year")
        year_end_index = p
        // fmt.Println(year_end_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
//line time_parser/time_parser.rl:62

        // fmt.Println("end day")
        day_end_index = p
        // fmt.Println(day_end_index)
    
	goto st86
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
//line time_parser/time_parser.go:2814
		switch data[p] {
		case 133:
			goto st48
		case 137:
			goto st53
		case 141:
			goto st56
		case 142:
			goto st61
		case 143:
			goto st87
		case 155:
			goto st64
		case 185:
			goto st65
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		switch data[p] {
		case 183:
			goto tr36
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
tr52:
//line time_parser/time_parser.rl:46

        // fmt.Println("end month")
        month_end_index = p
        // fmt.Println(month_end_index)
    
//line time_parser/time_parser.rl:62

        // fmt.Println("end day")
        day_end_index = p
        // fmt.Println(day_end_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st88
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
//line time_parser/time_parser.go:2893
		switch data[p] {
		case 151:
			goto st89
		case 152:
			goto st50
		case 156:
			goto st90
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		switch data[p] {
		case 165:
			goto tr36
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
		switch data[p] {
		case 136:
			goto tr43
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
tr43:
//line time_parser/time_parser.rl:51

        if month_begin_index < month_end_index {
            month = data[month_begin_index: month_end_index]
        }
    
	goto st91
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
//line time_parser/time_parser.go:2975
		switch data[p] {
		case 228:
			goto tr103
		case 229:
			goto tr81
		case 230:
			goto tr82
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
tr103:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st92
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
//line time_parser/time_parser.go:3017
		switch data[p] {
		case 184:
			goto st71
		case 185:
			goto st77
		case 186:
			goto st78
		case 187:
			goto st93
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		switch data[p] {
		case 138:
			goto st80
		case 189:
			goto tr38
		case 228:
			goto tr46
		case 229:
			goto tr47
		case 230:
			goto tr48
		case 231:
			goto tr49
		case 232:
			goto tr50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st39
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch data[p] {
		case 169:
			goto tr36
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
tr18:
//line time_parser/time_parser.rl:46

        // fmt.Println("end month")
        month_end_index = p
        // fmt.Println(month_end_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line time_parser/time_parser.go:3109
		switch data[p] {
		case 152:
			goto st11
		case 156:
			goto st29
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch data[p] {
		case 136:
			goto tr38
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch data[p] {
		case 157:
			goto st1
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch data[p] {
		case 140:
			goto st1
		case 148:
			goto st1
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		switch data[p] {
		case 138:
			goto st33
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch data[p] {
		case 228:
			goto tr2
		case 229:
			goto tr29
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
tr7:
//line time_parser/time_parser.rl:23

        fmt.Println("begin year")
        year_begin_index = p
        fmt.Println(year_begin_index)
    
//line time_parser/time_parser.rl:28

        // fmt.Println("end year")
        year_end_index = p
        // fmt.Println(year_end_index)
    
//line time_parser/time_parser.rl:41

        // fmt.Println("begin month")
        month_begin_index = p
        // fmt.Println(month_begin_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
//line time_parser/time_parser.rl:62

        // fmt.Println("end day")
        day_end_index = p
        // fmt.Println(day_end_index)
    
	goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line time_parser/time_parser.go:3282
		switch data[p] {
		case 133:
			goto st9
		case 137:
			goto st14
		case 141:
			goto st17
		case 142:
			goto st22
		case 143:
			goto st35
		case 155:
			goto st25
		case 185:
			goto st26
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch data[p] {
		case 183:
			goto tr36
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
tr8:
//line time_parser/time_parser.rl:46

        // fmt.Println("end month")
        month_end_index = p
        // fmt.Println(month_end_index)
    
//line time_parser/time_parser.rl:62

        // fmt.Println("end day")
        day_end_index = p
        // fmt.Println(day_end_index)
    
//line time_parser/time_parser.rl:57

        // fmt.Println("begin day")
        day_begin_index = p
        // fmt.Println(day_begin_index)
    
	goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line time_parser/time_parser.go:3361
		switch data[p] {
		case 151:
			goto st37
		case 152:
			goto st11
		case 156:
			goto st38
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		switch data[p] {
		case 165:
			goto tr36
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		switch data[p] {
		case 136:
			goto tr43
		case 228:
			goto tr2
		case 229:
			goto tr3
		case 230:
			goto tr4
		case 231:
			goto tr5
		case 232:
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st_out:
	_test_eof0: cs = 0; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof55: cs = 55; goto _test_eof
	_test_eof56: cs = 56; goto _test_eof
	_test_eof57: cs = 57; goto _test_eof
	_test_eof58: cs = 58; goto _test_eof
	_test_eof59: cs = 59; goto _test_eof
	_test_eof60: cs = 60; goto _test_eof
	_test_eof61: cs = 61; goto _test_eof
	_test_eof62: cs = 62; goto _test_eof
	_test_eof63: cs = 63; goto _test_eof
	_test_eof64: cs = 64; goto _test_eof
	_test_eof65: cs = 65; goto _test_eof
	_test_eof66: cs = 66; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof
	_test_eof70: cs = 70; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof77: cs = 77; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof79: cs = 79; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
	_test_eof81: cs = 81; goto _test_eof
	_test_eof82: cs = 82; goto _test_eof
	_test_eof83: cs = 83; goto _test_eof
	_test_eof84: cs = 84; goto _test_eof
	_test_eof85: cs = 85; goto _test_eof
	_test_eof86: cs = 86; goto _test_eof
	_test_eof87: cs = 87; goto _test_eof
	_test_eof88: cs = 88; goto _test_eof
	_test_eof89: cs = 89; goto _test_eof
	_test_eof90: cs = 90; goto _test_eof
	_test_eof91: cs = 91; goto _test_eof
	_test_eof92: cs = 92; goto _test_eof
	_test_eof93: cs = 93; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 75:
//line time_parser/time_parser.rl:62

        // fmt.Println("end day")
        day_end_index = p
        // fmt.Println(day_end_index)
    
//line time_parser/time_parser.rl:67

        if day_begin_index < day_end_index {
            day = data[day_begin_index: day_end_index]
        }
    
//line time_parser/time_parser.go:3542
		}
	}

	}

//line time_parser/time_parser.rl:141


    if cs < time_parser_first_final {
        fmt.Println("time parse error")
    } else {
    }
    fmt.Println(year_begin_index, year_end_index)
    // fmt.Println(day_begin_index, day_end_index)
    fmt.Println(year+"年 "+month + "月 " + day + "日")

    return year, month, day
}
