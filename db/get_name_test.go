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

	m := NewMockDB(ctrl)

	m.EXPECT().
		GetNameByIndex(gomock.Eq(1)).
		DoAndReturn(func(_ int) string {
			time.Sleep(1000)
			return "cat"
		})

	m.EXPECT().
		GetNameByIndex(gomock.Eq(2)).
		Return("dog")

}
