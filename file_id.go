package hedera

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hashgraph/hedera-sdk-go/proto"
)

type FileID struct {
	Shard uint64
	Realm uint64
	File  uint64
}

func FileIDFromString(s string) (FileID, error) {
	values := strings.SplitN(s, ".", 3)
	if len(values) != 3 {
		// Was not three values separated by periods
		return FileID{}, fmt.Errorf("expected {shard}.{realm}.{num}")
	}

	shard, err := strconv.Atoi(values[0])
	if err != nil {
		return FileID{}, err
	}

	realm, err := strconv.Atoi(values[1])
	if err != nil {
		return FileID{}, err
	}

	num, err := strconv.Atoi(values[2])
	if err != nil {
		return FileID{}, err
	}

	return FileID{
		Shard: uint64(shard),
		Realm: uint64(realm),
		File:  uint64(num),
	}, nil
}

func (id FileID) String() string {
	return fmt.Sprintf("%d.%d.%d", id.Shard, id.Realm, id.File)
}

func (id FileID) toProto() *proto.FileID {
	return &proto.FileID{
		ShardNum: int64(id.Shard),
		RealmNum: int64(id.Realm),
		FileNum:  int64(id.File),
	}
}

func fileIDFromProto(pb *proto.FileID) FileID {
	return FileID{
		Shard: uint64(pb.ShardNum),
		Realm: uint64(pb.RealmNum),
		File:  uint64(pb.FileNum),
	}
}