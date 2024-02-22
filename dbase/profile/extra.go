package profile

import (
	"github.com/opentdp/wechat-rest/args"
)

func GetAiModel(wxid, roomid string) *args.LLModel {

	var llmc *args.LLModel

	p, _ := Fetch(&FetchParam{Wxid: wxid, Roomid: roomid})

	if p != nil {
		llmc = args.LLM.Models[p.AiModel]
	}

	if llmc == nil {
		llmc = args.LLM.Models[args.LLM.Default]
	}

	if llmc == nil {
		for _, v := range args.LLM.Models {
			return v
		}
	}

	return llmc

}

func SetAiModel(wxid, roomid, argot, model string) error {

	p, err := Fetch(&FetchParam{
		Wxid:   wxid,
		Roomid: roomid,
	})

	if err == nil && p.Rd > 0 {
		err = Update(&UpdateParam{
			Wxid:    wxid,
			Roomid:  roomid,
			AiArgot: argot,
			AiModel: model,
		})
	} else {
		_, err = Create(&CreateParam{
			Wxid:    wxid,
			Roomid:  roomid,
			AiArgot: argot,
			AiModel: model,
		})
	}

	return err

}
