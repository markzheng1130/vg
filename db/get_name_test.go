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
		Get(gomock.Eq("mark")).
		DoAndReturn(func(_ string) (int, error) {
			time.Sleep(1000)
			return 123, nil
		})

	db.EXPECT().
		Get(gomock.Eq("eric")).
		DoAndReturn(func(_ string) (int, error) {
			time.Sleep(1000)
			return 456, nil
		})

	tests := []struct {
		name string
		arg  string
		want int
	}{
		{
			"test-case-1",
			"mark",
			123,
		},
		{
			"test-case-2",
			"eric",
			456,
		},
	}

	for _, test := range tests {
		Convey(test.name, t, func() {
			value := GetFromDB(db, test.arg)
			So(value, ShouldEqual, test.want)
		})
	}
}
