package ethash

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"hash"
	"math/big"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"unicode/utf8"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/sha3"
)

var mathFuncs *MathFuncs
var targetBits uint64 = 24

var (
	// InitialDifficulty represents the starting diff
	InitialDifficulty uint = 4
	// BaselinePrecisionDigits represents the starting precision level
	BaselinePrecisionDigits uint = 128
	// MaximumPrecisionDigits represents the max precision level
	MaximumPrecisionDigits uint = 1000
	// CurrentPrecision represents the current precision level
	CurrentPrecision uint = 128
	// DiffAdjustInBlocks represents the blocks that need to pass to recalculate diff
	DiffAdjustInBlocks uint = 100
	// MaxDiffChange represents the max diff change per block
	MaxDiffChange uint = 1
	// BlockDurationTarget represents the block time target in seconds
	BlockDurationTarget uint = 12
)

// init is called once on startup and creates a new list of math funcs
func init() {
	mathFuncs = NewFuncList()
	messagePrint("plhash -> created new mathfunc list with lenght : ", len(mathFuncs.FuncList))
}

func messagePrint(args ...interface{}) {
	fmt.Println("------------------------------------------------------------")
	fmt.Println(args...)
	fmt.Println("------------------------------------------------------------")
}

// Hash32 takes a hash type and returns the byte slice of the hash
func Hash32(h hash.Hash, data []byte) [32]byte {
	var out [32]byte
	h.Write(data)
	h.Sum(out[:0])
	h.Reset()
	return out
}

// Hash64  takes a hash type and returns the byte slice of the hash
func Hash64(h hash.Hash, data []byte) [64]byte {
	var out [64]byte
	h.Write(data)
	h.Sum(out[:0])
	h.Reset()
	return out
}

// HashPassA represents the SHA256 pass of the algo and returns a big int
// it first hashes the data byte slice with the given algo and then creates
// a 64 byte slice of it by reversing the 32 bytes of the original and appending it
func HashPassA(data []byte, index int) (*big.Int, error) {
	h := sha256.New()
	hash := Hash32(h, data)
	hash64 := ReverseBytes32To64(hash)
	fbn, err := calcPass(hash64[:], index)
	if err != nil {
		return nil, err
	}
	return fbn, nil
}

// HashPassB represents the first SHA3.512 pass of the algo and returns a big int
func HashPassB(data []byte, index int) (*big.Int, error) {
	h := sha3.New512()
	hash64 := Hash64(h, data)
	fbn, err := calcPass(hash64[:], index)
	if err != nil {
		return nil, err
	}
	return fbn, nil
}

// HashPassC represents the second SHA3.512 pass of the algo and returns a big int
func HashPassC(data []byte, index int) (*big.Int, error) {
	h := sha3.New512()
	hash64 := Hash64(h, data)
	fbn, err := calcPass(hash64[:], index)
	if err != nil {
		return nil, err
	}
	return fbn, nil
}

// HashPassD represents the BLAKE2S pass of the algo and returns a big int
func HashPassD(data []byte, index int) (*big.Int, error) {
	h, _ := blake2s.New256(nil)
	hash := Hash32(h, data)
	hash64 := ReverseBytes32To64(hash)
	fbn, err := calcPass(hash64[:], index)
	if err != nil {
		return nil, err
	}
	return fbn, nil
}

// HashPassE represents the BLAKE2B pass of the algo and returns a big int
func HashPassE(data []byte, index int) (*big.Int, error) {
	h, _ := blake2b.New512(nil)
	hash64 := Hash64(h, data)
	fbn, err := calcPass(hash64[:], index)
	if err != nil {
		return nil, err
	}
	return fbn, nil
}

// calcPass takes a hash and passes it via the math funcs, returns a big int
func calcPass(in []byte, index int) (*big.Int, error) {
	bn := ByteToBigInt(in)
	// fmt.Println("this is the byte to bigInt : ", bn)
	fx := BigIntToBigFloat(bn, BaselinePrecisionDigits)
	// fmt.Println("CALCPASS -> this the big float conversion of the number : ", fx)
	fnc := mathFuncs.FuncList[index]
	fn := fnc(fx)
	// fmt.Println("CALCPASS -> this is the outcome via the math func : ", fn)
	fs, err := RemoveDecFromFloat(fn)
	if err != nil {
		return nil, err
	}
	fbn, errN := SelectLastDigits(fs, 1*5) // depth is multiply to diff
	if errN != nil {
		return nil, errN
	}
	return fbn, nil
}

// CalcFinalHash returns the final hash of the hash passes
func CalcFinalHash(a, b, c, d, e *big.Int) *big.Int {
	fhp := sha3.New512()
	sumAB := ZeroBigInt()
	mulCD := ZeroBigInt()
	subABCD := ZeroBigInt()
	divE := ZeroBigInt()
	sumAB.Add(a, b)
	mulCD.Mul(c, d)
	subABCD.Sub(sumAB, mulCD)
	divE.Mul(subABCD, e)
	fh := Hash64(fhp, divE.Bytes())
	finalBigNumber := ZeroBigInt()
	finalBigNumber.SetBytes(fh[:])
	return finalBigNumber
}

// CompletePass represents all the pass in serial plus the final hash arithmetic function
func CompletePass(data []byte, indexes []int) (*big.Int, error) {
	passA, errA := HashPassA(data, indexes[0])
	if errA != nil {
		fmt.Println(errA)
		return nil, errA
	}

	passB, errB := HashPassB(data, indexes[1])
	if errB != nil {
		fmt.Println(errB)
		return nil, errB
	}

	passC, errC := HashPassC(data, indexes[2])
	if errC != nil {
		fmt.Println(errC)
		return nil, errC
	}

	passD, errD := HashPassD(data, indexes[3])
	if errD != nil {
		fmt.Println(errD)
		return nil, errD
	}

	passE, errE := HashPassE(data, indexes[4])
	if errE != nil {
		fmt.Println(errE)
		return nil, errE
	}
	// fmt.Println("")
	// fmt.Println("printing passes values")
	// fmt.Println("pass a : ", passA)
	// fmt.Println("pass b : ", passB)
	// fmt.Println("pass c : ", passC)
	// fmt.Println("pass d : ", passD)
	// fmt.Println("pass e : ", passE)
	finalHash := CalcFinalHash(passA, passB, passC, passD, passE)
	// fmt.Println("final big number : ", finalHash)
	return finalHash, nil
}

// ConstractDiffString constracts a new diff string
func ConstractDiffString(difflevel int) string {
	var diffString string
	var data []string
	for i := 1; i < difflevel+1; i++ {
		e := i % 10
		d := strconv.Itoa(e)

		data = append(data, d)
	}
	diffString = strings.Join(data, "")
	return diffString
}

// GetLeadingMatchingChars returns the number of leading matching chars in a string
func GetLeadingMatchingChars(in, match string) int {
	var numberOfLeadingChars int
	if len(match) > len(in) {
		result := strings.Replace(match, in, "", -1)
		numberOfLeadingChars = len(match) - len(result)
	} else {
		numberOfLeadingChars = 0
		for i, c := range []byte(match) {
			currentInputChar := in[i]
			if currentInputChar == c {
				numberOfLeadingChars = numberOfLeadingChars + 1
			} else {
				break
			}
		}
	}
	return numberOfLeadingChars
}

// Pow represents the pow function
func Pow(a *big.Float, e uint) *big.Float {
	e = uint(e)
	result := ZeroBigFloat().Copy(a)
	for i := uint(0); i < e-1; i++ {
		result = Mul(result, a)
	}
	return result
}

// Pow3 represents pow(x, 3)
func Pow3(in *big.Float) *big.Float {
	return Pow(in, 3)
}

// Pow4 represents pow(x, 4)
func Pow4(in *big.Float) *big.Float {
	return Pow(in, 4)
}

// Pow5 represents pow(x, 5)
func Pow5(in *big.Float) *big.Float {
	return Pow(in, 5)
}

// Pow6 represents pow(x, 6)
func Pow6(in *big.Float) *big.Float {
	return Pow(in, 6)
}

// Pow7 represents pow(x, 7)
func Pow7(in *big.Float) *big.Float {
	return Pow(in, 7)
}

// Pow8 represents pow(x, 8)
func Pow8(in *big.Float) *big.Float {
	return Pow(in, 8)
}

// Pow9 represents pow(x, 9)
func Pow9(in *big.Float) *big.Float {
	return Pow(in, 9)
}

// Pow10 represents pow(x, 10)
func Pow10(in *big.Float) *big.Float {
	return Pow(in, 10)
}

// Pow11 represents pow(x, 11)
func Pow11(in *big.Float) *big.Float {
	return Pow(in, 11)
}

// Pow12 represents pow(x, 12)
func Pow12(in *big.Float) *big.Float {
	return Pow(in, 12)
}

// Root represets root(x, n)
func Root(a *big.Float, n uint64) *big.Float {
	limit := Pow(NewBigFloat(2), CurrentPrecision)
	n1 := n - 1
	n1f, rn := NewBigFloat(float64(n1)), Div(NewBigFloat(1.0), NewBigFloat(float64(n)))
	x, x0 := NewBigFloat(1.0), ZeroBigFloat()
	_ = x0
	for {
		potx, t2 := Div(NewBigFloat(1.0), x), a
		for b := n1; b > 0; b >>= 1 {
			if b&1 == 1 {
				t2 = Mul(t2, potx)
			}
			potx = Mul(potx, potx)
		}
		x0, x = x, Mul(rn, Add(Mul(n1f, x), t2))
		if Lesser(Mul(Abs(Sub(x, x0)), limit), x) {
			break
		}
	}
	return x
}

// Root3 represents root(x, 3)
func Root3(in *big.Float) *big.Float {
	return Root(in, 3)
}

// Root4 represents root(x, 4)
func Root4(in *big.Float) *big.Float {
	return Root(in, 4)
}

// Root5 represents root(x, 5)
func Root5(in *big.Float) *big.Float {
	return Root(in, 5)
}

// Root6 represents root(x, 6)
func Root6(in *big.Float) *big.Float {
	return Root(in, 6)
}

// Root7 represents root(x, 7)
func Root7(in *big.Float) *big.Float {
	return Root(in, 7)
}

// Root8 represents root(x, 8)
func Root8(in *big.Float) *big.Float {
	return Root(in, 8)
}

// Root9 represents root(x, 9)
func Root9(in *big.Float) *big.Float {
	return Root(in, 9)
}

// Root10 represents root(x, 10)
func Root10(in *big.Float) *big.Float {
	return Root(in, 10)
}

// Root11 represents root(x, 11)
func Root11(in *big.Float) *big.Float {
	return Root(in, 11)
}

// Root12 represents root(x, 12)
func Root12(in *big.Float) *big.Float {
	return Root(in, 12)
}

// Abs represents the abs of a big float
func Abs(a *big.Float) *big.Float {
	return ZeroBigFloat().Abs(a)
}

// NewBigFloat returns a new big float with current global precision
func NewBigFloat(f float64) *big.Float {
	r := big.NewFloat(f)
	r.SetPrec(CurrentPrecision)
	return r
}

// Div devides two big floats
func Div(a, b *big.Float) *big.Float {
	return ZeroBigFloat().Quo(a, b)
}

// ZeroBigFloat returns a big float with zero value
func ZeroBigFloat() *big.Float {
	r := big.NewFloat(0.0)
	r.SetPrec(CurrentPrecision)
	return r
}

// Mul multiplies two big floats
func Mul(a, b *big.Float) *big.Float {
	return ZeroBigFloat().Mul(a, b)
}

// Add adds two big floats
func Add(a, b *big.Float) *big.Float {
	return ZeroBigFloat().Add(a, b)
}

// Sub subs two big floats
func Sub(a, b *big.Float) *big.Float {
	return ZeroBigFloat().Sub(a, b)
}

// Lesser returns a bool
func Lesser(x, y *big.Float) bool {
	return x.Cmp(y) == -1
}

// ZeroBigInt returns a new big int with zero value
func ZeroBigInt() *big.Int {
	return big.NewInt(0)
}

// MathFuncs represent the collection of the math functions used by the chain hasher
type MathFuncs struct {
	FuncList []func(*big.Float) *big.Float
}

// NewFuncList returns a new list of math functions
func NewFuncList() *MathFuncs {
	return &MathFuncs{
		FuncList: createMathList(),
	}
}

func createMathList() []func(*big.Float) *big.Float {
	var collection []func(*big.Float) *big.Float
	collection = append(collection, Pow3)
	collection = append(collection, Pow4)
	collection = append(collection, Pow5)
	collection = append(collection, Pow6)
	collection = append(collection, Pow7)
	collection = append(collection, Pow8)
	collection = append(collection, Pow9)
	collection = append(collection, Pow10)
	collection = append(collection, Pow11)
	collection = append(collection, Pow12)
	collection = append(collection, Root3)
	collection = append(collection, Root4)
	collection = append(collection, Root5)
	collection = append(collection, Root6)
	collection = append(collection, Root7)
	collection = append(collection, Root8)
	collection = append(collection, Root9)
	collection = append(collection, Root10)
	collection = append(collection, Root11)
	collection = append(collection, Root12)
	collection = append(collection, Abs)

	// Apend more here
	// dummys
	collection = append(collection, Pow3)
	collection = append(collection, Pow4)
	collection = append(collection, Pow5)
	collection = append(collection, Pow6)
	collection = append(collection, Pow7)
	collection = append(collection, Pow8)
	collection = append(collection, Pow9)
	collection = append(collection, Pow10)
	collection = append(collection, Pow11)
	collection = append(collection, Pow12)
	collection = append(collection, Root3)
	collection = append(collection, Root4)
	collection = append(collection, Root5)
	collection = append(collection, Root6)
	collection = append(collection, Root7)
	collection = append(collection, Root8)
	collection = append(collection, Root9)
	collection = append(collection, Root10)
	collection = append(collection, Root11)
	collection = append(collection, Root12)
	collection = append(collection, Abs)
	collection = append(collection, Pow3)
	collection = append(collection, Pow4)
	collection = append(collection, Pow5)
	collection = append(collection, Pow6)
	collection = append(collection, Pow7)
	collection = append(collection, Pow8)
	collection = append(collection, Pow9)
	collection = append(collection, Pow10)
	collection = append(collection, Pow11)
	collection = append(collection, Pow12)
	collection = append(collection, Root3)
	collection = append(collection, Root4)
	collection = append(collection, Root5)
	collection = append(collection, Root6)
	collection = append(collection, Root7)
	collection = append(collection, Root8)
	collection = append(collection, Root9)
	collection = append(collection, Root10)
	collection = append(collection, Root11)
	collection = append(collection, Root12)
	collection = append(collection, Abs)
	collection = append(collection, Pow3)
	collection = append(collection, Pow4)
	collection = append(collection, Pow5)
	collection = append(collection, Pow6)
	collection = append(collection, Pow7)
	collection = append(collection, Pow8)
	collection = append(collection, Pow9)
	collection = append(collection, Pow10)
	collection = append(collection, Pow11)
	collection = append(collection, Pow12)
	collection = append(collection, Root3)
	collection = append(collection, Root4)
	collection = append(collection, Root5)
	collection = append(collection, Root6)
	collection = append(collection, Root7)
	collection = append(collection, Root8)
	collection = append(collection, Root9)
	collection = append(collection, Root10)
	collection = append(collection, Root11)
	collection = append(collection, Root12)
	collection = append(collection, Abs)
	collection = append(collection, Pow3)
	collection = append(collection, Pow4)
	collection = append(collection, Pow5)
	collection = append(collection, Pow6)
	collection = append(collection, Pow7)
	collection = append(collection, Pow8)
	collection = append(collection, Pow9)
	collection = append(collection, Pow10)
	collection = append(collection, Pow11)
	collection = append(collection, Pow12)
	collection = append(collection, Root3)
	collection = append(collection, Root4)
	collection = append(collection, Root5)
	collection = append(collection, Root6)
	collection = append(collection, Root7)
	collection = append(collection, Root8)
	collection = append(collection, Root9)
	collection = append(collection, Root10)
	collection = append(collection, Root11)
	collection = append(collection, Root12)
	collection = append(collection, Abs)
	return collection
}

// GetFuncIndexes retuns the indexes for the math funcs chosen
func GetFuncIndexes(lastBlockHash []byte) []int {
	if len(lastBlockHash) == 0 {
		return []int{1, 2, 3, 4, 5}
	}
	var funcIndexes = make([]int, 5)
	a := big.NewInt(0)
	a.SetBytes(lastBlockHash)
	// Trim ending zeros
	s := a.String()
	st := strings.TrimRight(s, "0")
	sm := st[len(st)-10:]
	// Seperate values and store them
	sm1, _ := strconv.Atoi(sm[:2])
	sm2, _ := strconv.Atoi(sm[2:4])
	sm3, _ := strconv.Atoi(sm[4:6])
	sm4, _ := strconv.Atoi(sm[6:8])
	sm5, _ := strconv.Atoi(sm[8:10])
	funcIndexes[0] = sm1
	funcIndexes[1] = sm2
	funcIndexes[2] = sm3
	funcIndexes[3] = sm4
	funcIndexes[4] = sm5
	return funcIndexes
}

// GetFunctionName returns the name of the function
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// IntToHex converts an int64 to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		return []byte{}
	}

	return buff.Bytes()
}

// ByteToBigInt takes a byte slice as input and returns a big.Int
func ByteToBigInt(in []byte) *big.Int {
	out := new(big.Int)
	out.SetBytes(in)
	return out
}

// BigIntToBigFloat takes a big int as input and returns a float64 with precision given
func BigIntToBigFloat(in *big.Int, prec uint) *big.Float {
	f := new(big.Float)
	f.SetInt(in)
	f.SetPrec(prec)

	return f
}

// ReverseString returns the reverse of the string input
func ReverseString(input string) string {
	size := len(input)
	buf := make([]byte, size)
	for i := 0; i < size; {
		r, n := utf8.DecodeRuneInString(input[i:])
		i += n
		utf8.EncodeRune(buf[size-i:], r)
	}
	return string(buf)
}

// ReverseBytes returns the reversed byte slice
func ReverseBytes(in []byte) []byte {
	revIn := in
	for i := len(revIn)/2 - 1; i >= 0; i-- {
		opp := len(revIn) - 1 - i
		revIn[i], revIn[opp] = revIn[opp], revIn[i]
	}
	return revIn
}

// BytesTo64Bytes returns a filled 64 byte slice from the input
func BytesTo64Bytes(in []byte) ([64]byte, error) {
	var b [64]byte
	i := len(in)
	if i == 0 {
		return [64]byte{}, fmt.Errorf("%s", "size of input is zero")
	}
	for index := 0; index < i; index++ {
		b[index] = in[index]
	}
	return b, nil
}

// RemoveDecFromFloat returns a big.Int from a input float after the deciminal removal
func RemoveDecFromFloat(in *big.Float) (*big.Int, error) {
	// fmt.Println("REMOVEDECFROMFLOAT -> this is the incoming float", in)
	s := fmt.Sprintf("%.128f", in)
	// fmt.Println("REMOVEDECFROMFLOAT -> this is the string", s)
	if len(s) == 0 {
		return nil, fmt.Errorf("%s", "string has 0 zero length")
	}
	x := strings.Replace(s, ".", "", -1)
	// fmt.Println("REMOVEDECFROMFLOAT -> this is the string result : ", x)
	a := big.NewInt(0)
	a.SetString(x, 0)
	// fmt.Println("REMOVEDECFROMFLOAT -> this is the final big int : ", a)
	return a, nil
}

// ReverseBytes32To64 returns a 64 byte slice filled with a 32 byte and its reverse version
func ReverseBytes32To64(in [32]byte) [64]byte {
	var b [64]byte
	revIn := in
	for i := len(revIn)/2 - 1; i >= 0; i-- {
		opp := len(revIn) - 1 - i
		revIn[i], revIn[opp] = revIn[opp], revIn[i]
	}
	for i := 0; i < 32; i++ {
		b[i] = in[i]
		b[i+32] = revIn[i]
	}
	return b
}

// SelectLastDigits select the number of digits from a big number and returns the outcome
// depth is always a multiply of diff
func SelectLastDigits(in *big.Int, depth uint64) (*big.Int, error) {
	a := ZeroBigInt()
	s := in.String()
	if s == "0" {
		return nil, fmt.Errorf("%s", "string has value of 0? something is not ok")
	}
	fs := strings.TrimRight(s, "0")
	index := int(depth)
	// fmt.Println(index, fs)
	fx := fs[len(fs)-index:]
	a.SetString(fx, 0)

	return a, nil
}
