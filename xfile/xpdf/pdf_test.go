package xpdf

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWritePDF1(t *testing.T) {
	lines := []*Line{
		{
			Type: TypeDivider,
			Text: "项目概述",
		},
		{
			Type: TypeText,
			Text: "第一章 项目概述\n本项目旨在实现新一代智能文档处理系统，通过结合人工智能与自动化技术，提升企业文档处理效率。",
		},
		{
			Type: TypeImage,
			Text: "./image001_S.jpg",
		},
	}
	pdfpath := "/tmp/output.pdf"
	fmt.Println(pdfpath)
	assert.Nil(t, Write(pdfpath, lines))
}

func TestWritePDF2(t *testing.T) {
	s := `[
    {
        "Type": 2,
        "Text": "血源诅咒DLC"
    },
    {
        "Type": 0,
        "Text": "《血源》DLC老猎人中新增4位BOSS及16把不同种类武器，同时通关难度也大大增加，那么该如何通关呢？小编以最快速度带来老猎人全流程图文攻略，将推进流程分区讲解，并着重讲解各武器获得方法，希望对各位有帮助。"
    },
    {
        "Type": 0,
        "Text": "大教堂区域及DLC进入方法"
    },
    {
        "Type": 0,
        "Text": "虽然很多人已经知道进入DLC的方法，不过这边还是先写一下吧"
    },
    {
        "Type": 0,
        "Text": "首先流程得进行到打死教堂的白羊，然后来到梦境在人偶旁边得到进入DLC的道具"
    },
    {
        "Type": 1,
        "Text": "/var/folders/x4/39kj3h0529zfnr7_fj_qm9cw0000gn/T/handbook_201511_687935/image001_S.jpg"
    },
    {
        "Type": 0,
        "Text": "然后传送到大教堂区，从左边出口出去，故意被原先的那只亚米拉达用手抓一下就正式进入DLC地图了"
    },
    {
        "Type": 1,
        "Text": "/var/folders/x4/39kj3h0529zfnr7_fj_qm9cw0000gn/T/handbook_201511_687935/image002_S.jpg"
    },
    {
        "Type": 0,
        "Text": "大教堂区全部装备位置"
    },
    {
        "Type": 0,
        "Text": "由于地形复杂，口述实难表达，直接截图告知，为老猎人四件套以及一把新武器——蛇尾丸。PS：人形猎人敌人击倒后可以获得害虫，也就是说单机也能得到害虫了，但是否固定单次掉落未知。"
    },
    {
        "Type": 1,
        "Text": "/var/folders/x4/39kj3h0529zfnr7_fj_qm9cw0000gn/T/handbook_201511_687935/image003_S.jpg"
    },
    {
        "Type": 1,
        "Text": "/var/folders/x4/39kj3h0529zfnr7_fj_qm9cw0000gn/T/handbook_201511_687935/image004_S.jpg"
    },
    {
        "Type": 1,
        "Text": "/var/folders/x4/39kj3h0529zfnr7_fj_qm9cw0000gn/T/handbook_201511_687935/image005_S.jpg"
    },
    {
        "Type": 1,
        "Text": "/var/folders/x4/39kj3h0529zfnr7_fj_qm9cw0000gn/T/handbook_201511_687935/image006_S.jpg"
    },
    {
        "Type": 1,
        "Text": "/var/folders/x4/39kj3h0529zfnr7_fj_qm9cw0000gn/T/handbook_201511_687935/image007_S.jpg"
    },
    {
        "Type": 0,
        "Text": "搜刮完道具之后别忘记拿掉劳伦斯身前的一个关键道具——眼球坠子（稍后用来开启通往研究大楼的升降梯），然后沿着上来的方向右手边出去进到下一个区域"
    },
    {
        "Type": 1,
        "Text": "/var/folders/x4/39kj3h0529zfnr7_fj_qm9cw0000gn/T/handbook_201511_687935/image008_S.jpg"
    },
    {
        "Type": 1,
        "Text": "/var/folders/x4/39kj3h0529zfnr7_fj_qm9cw0000gn/T/handbook_201511_687935/image009_S.jpg"
    }
]`
	lines := make([]*Line, 0)
	if err := json.Unmarshal([]byte(s), &lines); err != nil {
		panic(err)
	}
	pdfpath := "/tmp/output.pdf"
	fmt.Println(pdfpath)
	assert.Nil(t, Write(pdfpath, lines))
}
