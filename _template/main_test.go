package main

import (
	"math"
	"testing"
)

func TestGetAngle(t *testing.T) {
	tests := []struct {
		x, y float64
		want float64
	}{
		{1, 0, 0},      // x軸上の点
		{0, 1, 90},     // y軸上の点
		{-1, 0, 180},   // x軸の負の方向
		{0, -1, -90},   // y軸の負の方向
		{1, 1, 45},     // 第1象限の点
		{-1, 1, 135},   // 第2象限の点
		{-1, -1, -135}, // 第3象限の点
		{1, -1, -45},   // 第4象限の点
	}

	for _, tt := range tests {
		t.Run("TestGetAngle", func(t *testing.T) {
			got := getAngle(tt.x, tt.y)
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("getAngle(%f, %f) = %f; want %f", tt.x, tt.y, got, tt.want)
			}
		})
	}
}
