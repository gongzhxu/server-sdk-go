package sdk

import (
	"encoding/json"
	"errors"

	"github.com/astaxie/beego/httplib"
)

// ListWordFilterResult listWordFilter返回结果
type ListWordFilterResult struct {
	Words []SensitiveWord `json:"words"`
}

// SensitiveWord 敏感词
type SensitiveWord struct {
	Type        string `json:"type"`
	Word        string `json:"word"`
	ReplaceWord string `json:"replaceWord"`
}

// SensitiveAdd 添加敏感词
/*
*@param  word:敏感词，最长不超过 32 个字符。（必传）
*
*@return error
 */
func (rc *RongCloud) SensitiveAdd(keyword, replace string, sensitiveType int) error {
	if keyword == "" {
		return errors.New("20005 Paramer 'keyword' is required")
	}
	if replace == "" {
		return errors.New("20005 Paramer 'replace' is required")
	}
	req := httplib.Post(rc.RongCloudURI + "/sensitiveword/add." + ReqType)
	rc.FillHeader(req)
	req.Param("word", keyword)
	switch sensitiveType {
	case 0:
		req.Param("replaceWord", replace)
	case 1:

	default:
		return errors.New("20005 Paramer 'replace' is required")
	}
	byteData, err := req.Bytes()
	if err != nil {
		return err
	}
	var code CodeResult
	if err := json.Unmarshal(byteData, &code); err != nil {
		return err
	}
	if code.Code != 200 {
		return RCErrorNew(code.Code, code.ErrorMessage)
	}

	return nil
}

// SensitiveGetList 查询敏感词列表方法
/*
*@return ListWordFilterResult
 */
func (rc *RongCloud) SensitiveGetList() (ListWordFilterResult, error) {

	req := httplib.Post(rc.RongCloudURI + "/sensitiveword/list." + ReqType)
	rc.FillHeader(req)
	byteData, err := req.Bytes()
	if err != nil {
		return ListWordFilterResult{}, err
	}

	var ret ListWordFilterResult
	if err := json.Unmarshal(byteData, &ret); err != nil {
		return ListWordFilterResult{}, err
	}
	var code CodeResult
	if err := json.Unmarshal(byteData, &code); err != nil {
		return ListWordFilterResult{}, err
	}
	if code.Code != 200 {
		return ListWordFilterResult{}, RCErrorNew(code.Code, code.ErrorMessage)
	}
	return ret, err

}

// SensitiveRemove 移除敏感词方法（从敏感词列表中，移除某一敏感词。）
/*
*@param  word:敏感词，最长不超过 32 个字符。（必传）
*
*@return error
 */
func (rc *RongCloud) SensitiveRemove(keywords []string) error {
	if len(keywords) == 0 {
		return errors.New("20005 Paramer 'keywords' is required")
	}

	req := httplib.Post(rc.RongCloudURI + "/sensitiveword/batch/delete." + ReqType)
	rc.FillHeader(req)
	for _, v := range keywords {
		req.Param("words", v)
	}
	byteData, err := req.Bytes()
	if err != nil {
		return err
	}

	var code CodeResult
	if err := json.Unmarshal(byteData, &code); err != nil {
		return err
	}
	if code.Code != 200 {
		return errors.New(code.ErrorMessage)
	}
	return nil

}
