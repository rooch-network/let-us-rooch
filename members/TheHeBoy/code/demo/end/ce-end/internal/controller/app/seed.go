package app

import (
	"github.com/gin-gonic/gin"
	"gohub/internal/request/app"
	"gohub/internal/request/validators"
	"gohub/internal/service"
	"gohub/pkg/logger"
	"gohub/pkg/response"
	"net/http"
)

type SeedController struct {
}

var seedService = service.Seed

func (ctl *SeedController) RandomUsableSeed(c *gin.Context) {
	req := app.AddressReq{}
	if ok := validators.Validate(c, &req); !ok {
		return
	}

	seed, id, err := seedService.RandomUsableSeed(req.Address)

	if err != nil {
		logger.Errorv(err)
		response.ErrorStr(c, "生成 NFT 随机种子失败")
	} else {
		response.SuccessData(c, gin.H{
			"num":   id + 1,
			"hSeed": seed,
		})
	}
}

func (ctl *SeedController) UsedTempSeed(c *gin.Context) {
	req := app.AddressReq{}
	if ok := validators.Validate(c, &req); !ok {
		return
	}

	seed := seedService.UsedTempSeed(req.Address)

	response.SuccessData(c, gin.H{
		"hSeed": seed,
	})
}

func (ctl *SeedController) GetSeedsByAddress(c *gin.Context) {
	req := app.AddressReq{}
	if ok := validators.Validate(c, &req); !ok {
		return
	}

	hSeeds, err := seedService.GetSeedsByAddress(req.Address)
	if err != nil {
		logger.Errorv(err)
		response.ErrorStr(c, "查询seed失败")
	} else {
		response.SuccessData(c, gin.H{
			"seeds": hSeeds,
		})
	}
}

func (ctl *SeedController) SeedHtml(c *gin.Context) {
	hSeed := c.Param("hSeed")
	template, err := service.Seed.FillTemplate("storage/template/seed_nft.html", hSeed)
	if err != nil {
		logger.Errorv(err)
		response.ErrorStr(c, "系统错误")
	}
	// 设置响应头
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, string(template))
}
