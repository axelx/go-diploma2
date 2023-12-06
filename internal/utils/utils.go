package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/axelx/go-diploma2/internal/models"
	pb "github.com/axelx/go-diploma2/internal/proto"
)

func IntToStr(i int) string {
	s := strconv.Itoa(i)
	return s
}

func StrToInt(s string) int {
	if v, err := strconv.Atoi(s); err == nil {
		return v
	}
	return 0
}

func StringToTime(s string) time.Time {
	date, err := time.Parse("2006-01-02", s)

	if err != nil {
		fmt.Println(err)
	}
	return date
}

func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func ToPointer[K int64 | float64 | time.Time](val K) *K {
	return &val
}

func UnPointer[K int64 | float64](val *K) K {
	if val == nil {
		return 0
	}
	return *val
}

func UnPointerTime(val *time.Time) time.Time {
	if val == nil {
		return time.Time{}
	}
	return *val
}

func ProtoToEnt(entProto *pb.Entity) models.Entity {
	ent := models.Entity{
		UserID:             int(entProto.UserID),
		Text:               entProto.Text,
		BankCard:           IntToStr(int(entProto.BankCard)),
		CreatedAtTimestamp: int(entProto.CreatedAtTimestamp),
		CreatedAt:          ToPointer(StringToTime(entProto.CreatedAt)),
		UpdatedAt:          ToPointer(StringToTime(entProto.UpdatedAt)),
	}

	return ent
}

func EntToProto(en *models.Entity) *pb.Entity {
	entity := []pb.Entity{{
		ID:                 int32(en.ID),
		UserID:             int32(en.UserID),
		Text:               en.Text,
		BankCard:           int64(StrToInt(en.BankCard)),
		CreatedAtTimestamp: int64(en.CreatedAtTimestamp),
		CreatedAt:          TimeToString(UnPointerTime(en.CreatedAt)),
		UpdatedAt:          TimeToString(UnPointerTime(en.UpdatedAt)),
	}}

	return &entity[0]
}
