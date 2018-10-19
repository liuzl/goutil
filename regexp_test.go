package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var str = `<td class="a-content"><p>发信人: huainan (南瓜), 信区: RealEstate <br/> 标  题: Re: 预测 天津房价 2年内下跌30%以上 <br/> 发信站: 水木社区 (Wed May 23 06:26:52 2018), 站内 <br/>  <br/> 你真逗，天津房价从1月份就会暖了，在出政策前大概回暖5-10%吧。 <br/>  <br/> 【 在 winn 的大作中提到: 】 <br/> <font class="f006">:  </font> <br/> <font class="f006">: 前提：房住不炒政策长期执行。天津限购政策不解除。  </font> <br/> <font class="f006">: 理由：天津在限购前的房价，是有很多外地人买房，结合开发商拉人冒充买房客，炒上去  </font> <br/> <font class="f006">: 的。  </font> <br/> <font class="f006">: 天津人民的收入支撑不了这种房价，所以从限购以来，房价一直在下跌。  </font> <br/>  <br/> #发自zSMTH@PAT-AL00 <br/> -- <br/>  <br/> <font class="f000"></font><font class="f001">※ 来源:·水木社区 <a target="_blank" href="http://www.newsmth.net">http://www.newsmth.net</a>·[FROM: 117.136.38.*]</font><font class="f000"> <br/> </font></p></td>`

var rp = []string{
	`发信人:(?P<author>.+?)\((?P<nick>.+?)\).*?信区:(?P<board>.+?)<br/>?`,
	`标  题:(?P<title>.+?)<br/>`,
	`发信站:(?P<site>.+?)\((?P<time>.+?)\)`,
}

func TestRe(t *testing.T) {
	for _, p := range rp {
		ret, err := RegexpExtract(str, p)
		if err != nil {
			t.Error(err)
		}
		t.Log(ret)
	}

	s := "https://music.163.com/artist?id=9489"
	r := "/artist\\?id=\\d+"
	assert.Equal(t, RegexpMatch(s, r), true, "should be equal")
}
