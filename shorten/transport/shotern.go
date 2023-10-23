package transport

import (
	"fmt"
	"net/http"
	"strings"

	"crypto/md5"
	"encoding/hex"

	"github.com/gin-gonic/gin"
	"github.com/quandoan/shorten_url/db"
	"github.com/quandoan/shorten_url/modules/shorten/biz"
	"github.com/quandoan/shorten_url/modules/shorten/storage"
)

type md5Hash struct {}

func NewMD5Hash() *md5Hash {
	return &md5Hash{}
}

func (h *md5Hash) Hash(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
} 

func CreateVirtualLink(db *db.MemDb, r *gin.Engine) gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		var url string
		if err := ctx.ShouldBind(&url); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid request",
			})
			return
		}

		store := storage.NewShortenStore(*db)
		biz := biz.NewShortenBiz(store, &md5Hash{})

		link, err := biz.CreateVirtualLink(ctx, url)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Can not create",
			})
			return
		}

		fmt.Println("====MemDb: ", db.Store)

		ctx.JSON(http.StatusOK, gin.H{
			"data": link,
		})

		linkArr := strings.Split(link, "/")
		code := linkArr[len(linkArr) - 1]
		r.GET(fmt.Sprintf("/%s", code), func(ctx *gin.Context) {
			ctx.Redirect(http.StatusFound, db.Store[code])
		})
	}
}