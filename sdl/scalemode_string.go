// Code generated by "stringer -type=ScaleMode"; DO NOT EDIT.

package sdl

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ScaleModeNearest-0]
	_ = x[ScaleModeLinear-1]
	_ = x[ScaleModeBest-2]
}

const _ScaleMode_name = "ScaleModeNearestScaleModeLinearScaleModeBest"

var _ScaleMode_index = [...]uint8{0, 16, 31, 44}

func (i ScaleMode) String() string {
	if i >= ScaleMode(len(_ScaleMode_index)-1) {
		return "ScaleMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ScaleMode_name[_ScaleMode_index[i]:_ScaleMode_index[i+1]]
}
