package db

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := NewMockDB(ctrl)

	db.EXPECT().
		GetNameByIndex(gomock.Eq(1)).
		DoAndReturn(func(_ int) string {
			time.Sleep(1000)
			return "cat"
		})

	db.EXPECT().
		GetNameByIndex(gomock.Eq(2)).
		Return("dog")

	tests := []struct {
		name string
		arg  int
		want string
	}{
		{
			"test-case-1",
			1,
			"cat",
		},
		{
			"test-case-2",
			2,
			"dog",
		},
	}

	for _, test := range tests {
		Convey(test.name, t, func() {
			name := GetName(db, test.arg)
			So(name, ShouldEqual, test.want)
		})
	}
}
