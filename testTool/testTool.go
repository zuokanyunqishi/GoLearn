package testTool

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/smartystreets/assertions/should"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/ysmood/got/lib/gop"
)

const (
	Equal int = iota
	NotEqual
	AlmostEqual
	NotAlmostEqual
	Resemble
	NotResemble
	PointTo
	NotPointTo
	BeNil
	NotBeNil
	BeTrue
	BeFalse
	BeZeroValue
	NotBeZeroValue

	BeGreaterThan
	BeGreaterThanOrEqualTo
	BeLessThan
	BeLessThanOrEqualTo
	BeBetween
	NotBeBetween
	BeBetweenOrEqual
	NotBeBetweenOrEqual
	Contain
	NotContain
	ContainKey
	NotContainKey
	BeIn
	NotBeIn
	BeEmpty
	NotBeEmpty
	HaveLength
	StartWith
	NotStartWith
	EndWith
	NotEndWith
	BeBlank
	NotBeBlank
	ContainSubstring
	NotContainSubstring
	Panic
	NotPanic
	PanicWith
	NotPanicWith
	HaveSameTypeAs
	NotHaveSameTypeAs
	Implement
	NotImplement
	HappenBefore
	HappenOnOrBefore
	HappenAfter
	HappenOnOrAfter
	HappenBetween
	HappenOnOrBetween
	NotHappenOnOrBetween
	HappenWithin
	NotHappenWithin
	BeChronological

	BeError
)

var assertionArray = [...]func(actual interface{}, expected ...interface{}) string{
	Equal: should.Equal,
	should.NotEqual,
	should.AlmostEqual,
	should.NotAlmostEqual,
	should.Resemble,
	should.NotResemble,
	should.PointTo,
	should.NotPointTo,
	should.BeNil,
	should.NotBeNil,
	should.BeTrue,
	should.BeFalse,
	should.BeZeroValue,
	should.NotBeZeroValue,

	should.BeGreaterThan,
	should.BeGreaterThanOrEqualTo,
	should.BeLessThan,
	should.BeLessThanOrEqualTo,
	should.BeBetween,
	should.NotBeBetween,
	should.BeBetweenOrEqual,
	should.NotBeBetweenOrEqual,
	should.Contain,
	should.NotContain,
	should.ContainKey,
	should.NotContainKey,
	should.BeIn,
	should.NotBeIn,
	should.BeEmpty,
	should.NotBeEmpty,
	should.HaveLength,
	should.StartWith,
	should.NotStartWith,
	should.EndWith,
	should.NotEndWith,
	should.BeBlank,
	should.NotBeBlank,
	should.ContainSubstring,
	should.NotContainSubstring,
	should.Panic,
	should.NotPanic,
	should.PanicWith,
	should.NotPanicWith,
	should.HaveSameTypeAs,
	should.NotHaveSameTypeAs,
	should.Implement,
	should.NotImplement,
	should.HappenBefore,
	should.HappenOnOrBefore,
	should.HappenAfter,
	should.HappenOnOrAfter,
	should.HappenBetween,
	should.HappenOnOrBetween,
	should.NotHappenOnOrBetween,
	should.HappenWithin,
	should.NotHappenWithin,
	should.BeChronological,

	BeError: should.BeError,
}

func Test(t *testing.T, actual interface{}, assertion int, expected ...interface{}) {
	Convey("test", t, func() {
		So(actual, assertionArray[assertion], expected...)
	})
}

func EqualTest(t *testing.T, actual interface{}, expected ...interface{}) {
	Convey("testEqual", t, func() {
		So(actual, assertionArray[Equal], expected...)
	})
}

func NotEqualTest(t *testing.T, actual interface{}, expected ...interface{}) {
	Convey("testNotEqual", t, func() {
		So(actual, assertionArray[NotEqual], expected...)
	})
}

func BeTrueTest(t *testing.T, actual interface{}, expected ...interface{}) {
	Convey("testBeTrue", t, func() {
		So(actual, assertionArray[BeTrue], expected...)
	})
}

func BeFalseTest(t *testing.T, actual interface{}, expected ...interface{}) {
	Convey("testBeFalse", t, func() {
		So(actual, assertionArray[BeFalse], expected...)
	})
}

func BeBetweenTest(t *testing.T, actual interface{}, expected ...interface{}) {
	Convey("testBeBetween", t, func() {
		So(actual, assertionArray[BeBetween], expected...)
	})
}

func BeLessThanTest(t *testing.T, actual interface{}, expected ...interface{}) {
	Convey("testBeLessThan", t, func() {
		So(actual, assertionArray[BeLessThan], expected...)
	})
}

// GoID go
func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// 得到id字符串
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func Dump(value ...interface{}) {
	spew.Dump(value...)

}

func Dd(value ...interface{}) {
	gop.P(value...)
}
