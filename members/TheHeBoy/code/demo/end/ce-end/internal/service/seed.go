package service

import (
	"bytes"
	"github.com/bwmarrin/snowflake"
	"github.com/pkg/errors"
	"gohub/internal/dao"
	"gohub/internal/model"
	"gohub/pkg/hashidsP"
	"gohub/pkg/snowflakeP"
	"gorm.io/gorm"
	"html/template"
	"math"
	"math/rand"
	"sync"
	"time"
)

type SeedService struct {
}

var seedDao = dao.Seed
var Seed = new(SeedService)
var mu sync.Mutex

// address => hSeed
var usedTempSeeds = make(map[string]string)

// 生成一个唯一的整数ID，支持JavaScript的整数长度
func (s *SeedService) generateUniqueInt() int64 {
	mu.Lock()
	defer mu.Unlock()
	// 获取当前时间的毫秒级时间戳
	timestamp := time.Now().UnixNano() / 1e6 // 毫秒级时间戳

	// 生成一个随机数部分
	// 随机数的范围需要根据时间戳的长度进行调整，以确保结果不超过2^53 - 1
	rand.Seed(time.Now().UnixNano())
	randomPart := rand.Int63n(1 << 20) // 生成一个随机数，20位足够大

	// 组合时间戳和随机数，确保结果不超过2^53 - 1
	uniqueID := (timestamp << 20) | randomPart

	return uniqueID
}

func (s *SeedService) generateShortID(node *snowflake.Node) int64 {
	snowflakeID := node.Generate().Int64()
	jsMaxInt := int64(math.Pow(2, 53) - 1)
	shortID := snowflakeID % jsMaxInt
	return shortID
}

func (s *SeedService) RandomUsableSeed(address string) (string, uint64, error) {
	address = dealAddress(address)
	hSeed, err := hashidsP.HashID.EncodeInt64([]int64{s.generateShortID(snowflakeP.Node)})
	if err != nil {
		return "", 0, errors.WithStack(err)
	}
	usedTempSeeds[address] = hSeed

	seedDO := model.SeedDO{}
	seedDao.Model().Last(&seedDO)

	return hSeed, seedDO.ID, nil
}

func (s *SeedService) useSeed(tx *gorm.DB, hSeed string, address string) error {
	address = dealAddress(address)

	if err := seedDao.Tx(tx).New().Create(&model.SeedDO{
		Address: address,
		HSeed:   hSeed,
	}).Error; err != nil {
		return errors.WithStack(err)
	}

	delete(usedTempSeeds, address)
	return nil
}

func (s *SeedService) UsedTempSeed(address string) string {
	address = dealAddress(address)
	return usedTempSeeds[address]
}

func (s *SeedService) GetSeedsByAddress(address string) ([]model.SeedDO, error) {
	address = dealAddress(address)
	hSeeds := make([]model.SeedDO, 0)
	if err := seedDao.Model().Where("address = ?", address).Find(&hSeeds).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return hSeeds, nil
}

func (s *SeedService) FillTemplate(filePath, hSeed string) ([]byte, error) {
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	data := struct {
		HSeed string
	}{
		HSeed: hSeed,
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, errors.WithStack(err)
	}
	return buf.Bytes(), nil
}
