
//line atoi/lk.rl:1
package atoi

import (
  "fmt"
)


//line atoi/lk.go:11
const atoi_start int = 1
const atoi_first_final int = 13
const atoi_error int = 0

const atoi_en_main int = 1


//line atoi/lk.rl:10


func Atoi(data string) (int64, error) {
  cs, p, pe := 0, 0, len(data)
  var neg bool
  var val 
  
//line atoi/lk.go:27
	{
	cs = atoi_start
	}

//line atoi/lk.go:32
	{

	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	switch cs {
	case 1:
		if data[p] == 104 {
			goto tr0;
		}
		goto tr1;
	case 0:
		goto _out
	case 2:
		if data[p] == 101 {
			goto tr2;
		}
		goto tr1;
	case 3:
		if data[p] == 108 {
			goto tr3;
		}
		goto tr1;
	case 4:
		if data[p] == 108 {
			goto tr4;
		}
		goto tr1;
	case 5:
		if data[p] == 111 {
			goto tr5;
		}
		goto tr1;
	case 6:
		if data[p] == 97 {
			goto tr6;
		}
		goto tr1;
	case 7:
		if data[p] == -26 {
			goto tr7;
		}
		goto tr1;
	case 8:
		if data[p] == -104 {
			goto tr8;
		}
		goto tr1;
	case 9:
		if data[p] == -81 {
			goto tr9;
		}
		goto tr1;
	case 10:
		if data[p] == -25 {
			goto tr10;
		}
		goto tr1;
	case 11:
		if data[p] == -102 {
			goto tr11;
		}
		goto tr1;
	case 12:
		if data[p] == -124 {
			goto tr12;
		}
		goto tr1;
	case 13:
		goto tr1;
	}

	tr1: cs = 0; goto _again
	tr0: cs = 2; goto _again
	tr2: cs = 3; goto _again
	tr3: cs = 4; goto _again
	tr4: cs = 5; goto _again
	tr5: cs = 6; goto _again
	tr6: cs = 7; goto _again
	tr7: cs = 8; goto _again
	tr8: cs = 9; goto _again
	tr9: cs = 10; goto _again
	tr10: cs = 11; goto _again
	tr11: cs = 12; goto _again
	tr12: cs = 13; goto _again

_again:
	if cs == 0 {
		goto _out
	}
	if p++; p != pe {
		goto _resume
	}
	_test_eof: {}
	_out: {}
	}

//line atoi/lk.rl:24


  if neg {
    val = -val
  }

  if cs < atoi_first_final {
    return 0, fmt.Errorf("atoi parse error: %s", data)
  }

  return val, nil
}
