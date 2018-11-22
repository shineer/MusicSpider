package musicspider

import (
	"fmt"
	"testing"
)

const testAll4artist, index4artist = false, 4

func Test_artist(t *testing.T) {

	artistID := map[string]string{
		"netease": "3681",
		"tencent": "003Nz2So3XXYek",
		"xiami":   "23485",
		"kugou":   "d3520",
		"baidu":   "362"}

	if testAll4artist {
		for _, s := range sites {
			resp := artist(s, artistID[s])
			fmt.Println(resp)
		}
	} else {
		resp := artist(sites[index4artist], artistID[sites[index4artist]])
		fmt.Println(resp)
	}

}

/*
TODO:
	虾米接口待完善
	酷狗接口提示参数不合法
*/
