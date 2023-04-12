package closed_range_test

import (
	"tdd_course/closed_range"
	"testing"
)

func TestNewClosedRange(t *testing.T) {
	t.Run("上端点が下端点より大きい場合の[3,8]だったらClosedRangeとnilが返る", func(t *testing.T) {
		cr, err := closed_range.NewClosedRange(3, 8)
		if err != nil && cr == nil {
			t.Fatalf("actual=%v, expect=%v", err, nil)
		}
	})

	t.Run("上端点が下端点より小さい場合の[8,3]だったらnilとerrorが返る", func(t *testing.T) {
		cr, err := closed_range.NewClosedRange(8, 3)
		if err == nil && cr != nil {
			t.Fatalf("actual=%v, expect=<error>", err)
		}
	})

	t.Run("上端点と下端点が等しい場合の[3,3]だったらnilとerrorが返る", func(t *testing.T) {
		cr, err := closed_range.NewClosedRange(3, 3)
		if err == nil && cr != nil {
			t.Fatalf("actual=%v, expect=<error>", err)
		}
	})
}

func TestString(t *testing.T) {
	tests := []struct {
		name   string
		lower  int
		upper  int
		expect string
	}{
		{
			name:   "下端点3,上端点8を渡すと文字列'[3,8]'が返ってくる",
			lower:  3,
			upper:  8,
			expect: "[3,8]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr, _ := closed_range.NewClosedRange(tt.lower, tt.upper)
			actual := cr.String()
			if actual != tt.expect {
				t.Fatalf("actual=%s, expect=%s", actual, tt.expect)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name   string
		lower  int
		upper  int
		target int
		expect bool
	}{
		{
			name:   "下端点3,上端点8を持っている時、整数5を指定するとtrueが返ってくる",
			lower:  3,
			upper:  8,
			target: 5,
			expect: true,
		},
		{
			name:   "下端点3,上端点8を持っている時、整数2を指定するとfalseが返ってくる",
			lower:  3,
			upper:  8,
			target: 2,
			expect: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr, _ := closed_range.NewClosedRange(tt.lower, tt.upper)
			actual := cr.Contains(tt.target)
			if actual != tt.expect {
				t.Fatalf("actual=%v, expect=%v", actual, tt.expect)
			}
		})
	}
}

func TestEquals(t *testing.T) {
	tests := []struct {
		name        string
		lower       int
		upper       int
		targetLower int
		targetUpper int
		expect      bool
	}{
		{
			name:        "下端点3,上端点8を持っている時、下端点3,上端点8を指定するとtrueが返ってくる",
			lower:       3,
			upper:       8,
			targetLower: 3,
			targetUpper: 8,
			expect:      true,
		},
		{
			name:        "下端点3,上端点8を持っている時、下端点2,上端点9を指定するとfalseが返ってくる",
			lower:       3,
			upper:       8,
			targetLower: 2,
			targetUpper: 9,
			expect:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr, _ := closed_range.NewClosedRange(tt.lower, tt.upper)
			target, _ := closed_range.NewClosedRange(tt.targetLower, tt.targetUpper)
			actual := cr.Equals(*target)
			if actual != tt.expect {
				t.Fatalf("actual=%v, expect=%v", actual, tt.expect)
			}
		})
	}
}

func TestRangeContains(t *testing.T) {

	tests := []struct {
		name        string
		lower       int
		upper       int
		targetLower int
		targetUpper int
		expect      bool
	}{
		{
			name:        "下端点3,上端点8を持っている時、下端点4,上端点7を指定するとtrueが返ってくる",
			lower:       3,
			upper:       8,
			targetLower: 4,
			targetUpper: 7,
			expect:      true,
		},
		{
			name:        "下端点3,上端点8を持っている時、下端点4,上端点9を指定するとfalseが返ってくる",
			lower:       3,
			upper:       8,
			targetLower: 4,
			targetUpper: 9,
			expect:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			target, _ := closed_range.NewClosedRange(tt.targetLower, tt.targetUpper)
			cr, _ := closed_range.NewClosedRange(tt.lower, tt.upper)
			actual := cr.RangeContains(*target)
			if actual != tt.expect {
				t.Fatalf("actual=%v, expect=%v", actual, tt.expect)
			}
		})
	}
}
