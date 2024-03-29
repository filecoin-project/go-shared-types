package types

import (
	"fmt"
	"github.com/filecoin-project/go-shared-types/pkg/params"
	"math/big"
	"strings"
)

type FIL BigInt

func (f FIL) String() string {
	r := new(big.Rat).SetFrac(f.Int, big.NewInt(params.FilecoinPrecision))
	if r.Sign() == 0 {
		return "0"
	}
	return strings.TrimRight(strings.TrimRight(r.FloatString(18), "0"), ".")
}

func (f FIL) Format(s fmt.State, ch rune) {
	switch ch {
	case 's', 'v':
		fmt.Fprint(s, f.String())
	default:
		f.Int.Format(s, ch)
	}
}

func ParseFIL(s string) (FIL, error) {
	r, ok := new(big.Rat).SetString(s)
	if !ok {
		return FIL{}, fmt.Errorf("failed to parse %q as a decimal number", s)
	}

	r = r.Mul(r, big.NewRat(params.FilecoinPrecision, 1))
	if !r.IsInt() {
		return FIL{}, fmt.Errorf("invalid FIL value: %q", s)
	}

	return FIL{r.Num()}, nil
}
