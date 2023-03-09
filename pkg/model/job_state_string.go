// Code generated by "stringer -type=JobStateType --trimprefix=JobState --output job_state_string.go"; DO NOT EDIT.

package model

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[JobStateNew-0]
	_ = x[JobStateInProgress-1]
	_ = x[JobStateCancelled-2]
	_ = x[JobStateError-3]
	_ = x[JobStateCompleted-4]
	_ = x[JobStateQueued-5]
}

const _JobStateType_name = "NewInProgressCancelledErrorCompletedQueued"

var _JobStateType_index = [...]uint8{0, 3, 13, 22, 27, 36, 42}

func (i JobStateType) String() string {
	if i < 0 || i >= JobStateType(len(_JobStateType_index)-1) {
		return "JobStateType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _JobStateType_name[_JobStateType_index[i]:_JobStateType_index[i+1]]
}
