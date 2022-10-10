package rank

import (
	"context"

	"google.golang.org/protobuf/proto"
)

const (
	rankSnapshotKey = "cross:rank:snapshot" // 排行榜快照数据，用于对比昨日排名
)

type Rank struct {
	snapshotRank      map[string]*RankSnapshot // 排行榜快照 rankKey -> id -> score
	dirtySnapshotRank map[string]struct{}      // 脏标记
	db                *Dao                     // redis 数据存储
	// rpc
}

func NewRank() *Rank {
	r := &Rank{
		snapshotRank:      make(map[string]*RankSnapshot),
		dirtySnapshotRank: make(map[string]struct{}),
		db:                NewDao(),
	}
	return r
}

// Init 启动
func (r *Rank) Init(ctx context.Context) error {
	return r.loadSnapshotData()
}

// 加载昨日排行榜快照
func (r *Rank) loadSnapshotData() error {
	saveSnapshotRanks, err := r.db.LoadHashAllStringBytesSync(rankSnapshotKey)
	if err != nil {
		return err
	}

	for rankKey, data := range saveSnapshotRanks {
		snapshotData := new(RankSnapshot)
		err := proto.Unmarshal([]byte(data), snapshotData)
		if err != nil {
			return err
		}
		r.snapshotRank[rankKey] = snapshotData
	}
	return nil
}

// AddRankScore 更新单个排行榜，score 增加
func (r *Rank) AddRankScore(rankKey string, rankData map[string]int) {
	r.db.asyncRedis.ZIncrby(rankKey, rankData)
}

// UpdateRankScore 更新单个排行榜，score直接替换
func (r *Rank) UpdateRankScore(rankKey string, rankData map[string]float64) {
	r.db.asyncRedis.ZAdd(rankKey, rankData)
}

// DeleteRankData 删除排行榜中指定成员
func (r *Rank) DeleteRankData(rankKey string, mems ...string) {
	r.db.asyncRedis.ZRem(rankKey, mems...)
}

fun
