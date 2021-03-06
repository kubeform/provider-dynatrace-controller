package cos

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"hash/crc64"
	"io"
	"net/http"
	"os"
	"strconv"
)

type CIService service

type PicOperations struct {
	IsPicInfo int                  `json:"is_pic_info,omitempty"`
	Rules     []PicOperationsRules `json:"rules,omitemtpy"`
}
type PicOperationsRules struct {
	Bucket string `json:"bucket,omitempty"`
	FileId string `json:"fileid"`
	Rule   string `json:"rule"`
}

func EncodePicOperations(pic *PicOperations) string {
	if pic == nil {
		return ""
	}
	bs, err := json.Marshal(pic)
	if err != nil {
		return ""
	}
	return string(bs)
}

type ImageProcessResult struct {
	XMLName        xml.Name          `xml:"UploadResult"`
	OriginalInfo   *PicOriginalInfo  `xml:"OriginalInfo,omitempty"`
	ProcessResults *PicProcessObject `xml:"ProcessResults>Object,omitempty"`
}
type PicOriginalInfo struct {
	Key       string        `xml:"Key,omitempty"`
	Location  string        `xml:"Location,omitempty"`
	ImageInfo *PicImageInfo `xml:"ImageInfo,omitempty"`
	ETag      string        `xml:"ETag,omitempty"`
}
type PicImageInfo struct {
	Format      string `xml:"Format,omitempty"`
	Width       int    `xml:"Width,omitempty"`
	Height      int    `xml:"Height,omitempty"`
	Quality     int    `xml:"Quality,omitempty"`
	Ave         string `xml:"Ave,omitempty"`
	Orientation int    `xml:"Orientation,omitempty"`
}
type PicProcessObject struct {
	Key             string       `xml:"Key,omitempty"`
	Location        string       `xml:"Location,omitempty"`
	Format          string       `xml:"Format,omitempty"`
	Width           int          `xml:"Width,omitempty"`
	Height          int          `xml:"Height,omitempty"`
	Size            int          `xml:"Size,omitempty"`
	Quality         int          `xml:"Quality,omitempty"`
	ETag            string       `xml:"ETag,omitempty"`
	WatermarkStatus int          `xml:"WatermarkStatus,omitempty"`
	CodeStatus      int          `xml:"CodeStatus,omitempty"`
	QRcodeInfo      []QRcodeInfo `xml:"QRcodeInfo,omitempty"`
}
type QRcodeInfo struct {
	CodeUrl      string        `xml:"CodeUrl,omitempty"`
	CodeLocation *CodeLocation `xml:"CodeLocation,omitempty"`
}
type CodeLocation struct {
	Point []string `xml:"Point,omitempty"`
}

type picOperationsHeader struct {
	PicOperations string `header:"Pic-Operations" xml:"-" url:"-"`
}

type ImageProcessOptions = PicOperations

// ?????????????????? https://cloud.tencent.com/document/product/460/18147
func (s *CIService) ImageProcess(ctx context.Context, name string, opt *ImageProcessOptions) (*ImageProcessResult, *Response, error) {
	header := &picOperationsHeader{
		PicOperations: EncodePicOperations(opt),
	}
	var res ImageProcessResult
	sendOpt := sendOptions{
		baseURL:   s.client.BaseURL.BucketURL,
		uri:       "/" + encodeURIComponent(name) + "?image_process",
		method:    http.MethodPost,
		optHeader: header,
		result:    &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

type ImageRecognitionOptions struct {
	CIProcess  string `url:"ci-process,omitempty"`
	DetectType string `url:"detect-type,omitempty"`
}

type ImageRecognitionResult struct {
	XMLName       xml.Name         `xml:"RecognitionResult"`
	PornInfo      *RecognitionInfo `xml:"PornInfo,omitempty"`
	TerroristInfo *RecognitionInfo `xml:"TerroristInfo,omitempty"`
	PoliticsInfo  *RecognitionInfo `xml:"PoliticsInfo,omitempty"`
	AdsInfo       *RecognitionInfo `xml:"AdsInfo,omitempty"`
}
type RecognitionInfo struct {
	Code    int    `xml:"Code,omitempty"`
	Msg     string `xml:"Msg,omitempty"`
	HitFlag int    `xml:"HitFlag,omitempty"`
	Score   int    `xml:"Score,omitempty"`
	Label   string `xml:"Label,omitempty"`
	Count   int    `xml:"Count,omitempty"`
}

// ???????????? https://cloud.tencent.com/document/product/460/37318
func (s *CIService) ImageRecognition(ctx context.Context, name string, DetectType string) (*ImageRecognitionResult, *Response, error) {
	opt := &ImageRecognitionOptions{
		CIProcess:  "sensitive-content-recognition",
		DetectType: DetectType,
	}
	var res ImageRecognitionResult
	sendOpt := sendOptions{
		baseURL:  s.client.BaseURL.BucketURL,
		uri:      "/" + encodeURIComponent(name),
		method:   http.MethodGet,
		optQuery: opt,
		result:   &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

type PutVideoAuditingJobOptions struct {
	XMLName     xml.Name              `xml:"Request"`
	InputObject string                `xml:"Input>Object"`
	Conf        *VideoAuditingJobConf `xml:"Conf"`
}
type VideoAuditingJobConf struct {
	DetectType string                       `xml:",omitempty"`
	Snapshot   *PutVideoAuditingJobSnapshot `xml:",omitempty"`
	Callback   string                       `xml:",omitempty"`
}
type PutVideoAuditingJobSnapshot struct {
	Mode         string  `xml:",omitempty"`
	Count        int     `xml:",omitempty"`
	TimeInterval float32 `xml:",omitempty"`
	Start        float32 `xml:",omitempty"`
}

type PutVideoAuditingJobResult struct {
	XMLName    xml.Name `xml:"Response"`
	JobsDetail struct {
		JobId        string `xml:"JobId,omitempty"`
		State        string `xml:"State,omitempty"`
		CreationTime string `xml:"CreationTime,omitempty"`
		Object       string `xml:"Object,omitempty"`
	} `xml:"JobsDetail,omitempty"`
}

// ????????????-???????????? https://cloud.tencent.com/document/product/460/46427
func (s *CIService) PutVideoAuditingJob(ctx context.Context, opt *PutVideoAuditingJobOptions) (*PutVideoAuditingJobResult, *Response, error) {
	var res PutVideoAuditingJobResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/video/auditing",
		method:  http.MethodPost,
		body:    opt,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

type GetVideoAuditingJobResult struct {
	XMLName        xml.Name           `xml:"Response"`
	JobsDetail     *AuditingJobDetail `xml:",omitempty"`
	NonExistJobIds string             `xml:",omitempty"`
}
type AuditingJobDetail struct {
	Code          string                       `xml:",omitempty"`
	Message       string                       `xml:",omitempty"`
	JobId         string                       `xml:",omitempty"`
	State         string                       `xml:",omitempty"`
	CreationTime  string                       `xml:",omitempty"`
	Object        string                       `xml:",omitempty"`
	SnapshotCount string                       `xml:",omitempty"`
	Result        int                          `xml:",omitempty"`
	PornInfo      *RecognitionInfo             `xml:",omitempty"`
	TerrorismInfo *RecognitionInfo             `xml:",omitempty"`
	PoliticsInfo  *RecognitionInfo             `xml:",omitempty"`
	AdsInfo       *RecognitionInfo             `xml:",omitempty"`
	Snapshot      *GetVideoAuditingJobSnapshot `xml:",omitempty"`
}
type GetVideoAuditingJobSnapshot struct {
	Url           string           `xml:",omitempty"`
	PornInfo      *RecognitionInfo `xml:",omitempty"`
	TerrorismInfo *RecognitionInfo `xml:",omitempty"`
	PoliticsInfo  *RecognitionInfo `xml:",omitempty"`
	AdsInfo       *RecognitionInfo `xml:",omitempty"`
}

// ????????????-???????????? https://cloud.tencent.com/document/product/460/46926
func (s *CIService) GetVideoAuditingJob(ctx context.Context, jobid string) (*GetVideoAuditingJobResult, *Response, error) {
	var res GetVideoAuditingJobResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/video/auditing/" + jobid,
		method:  http.MethodGet,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

type PutAudioAuditingJobOptions struct {
	XMLName     xml.Name              `xml:"Request"`
	InputObject string                `xml:"Input>Object"`
	Conf        *AudioAuditingJobConf `xml:"Conf"`
}
type AudioAuditingJobConf struct {
	DetectType string `xml:",omitempty"`
	Callback   string `xml:",omitempty"`
}
type PutAudioAuditingJobResult PutVideoAuditingJobResult
type GetAudioAuditingJobResult GetVideoAuditingJobResult

// ????????????-???????????? https://cloud.tencent.com/document/product/460/53395
func (s *CIService) PutAudioAuditingJob(ctx context.Context, opt *PutAudioAuditingJobOptions) (*PutAudioAuditingJobResult, *Response, error) {
	var res PutAudioAuditingJobResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/audio/auditing",
		method:  http.MethodPost,
		body:    opt,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// ????????????-???????????? https://cloud.tencent.com/document/product/460/53396
func (s *CIService) GetAudioAuditingJob(ctx context.Context, jobid string) (*GetAudioAuditingJobResult, *Response, error) {
	var res GetAudioAuditingJobResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/audio/auditing/" + jobid,
		method:  http.MethodGet,
		result:  &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// ?????????????????????-??????????????? https://cloud.tencent.com/document/product/460/18147
// ?????????-??????????????? https://cloud.tencent.com/document/product/460/19017
// ???????????????-??????????????? https://cloud.tencent.com/document/product/460/37513
func (s *CIService) Put(ctx context.Context, name string, r io.Reader, uopt *ObjectPutOptions) (*ImageProcessResult, *Response, error) {
	if r == nil {
		return nil, nil, fmt.Errorf("reader is nil")
	}
	if err := CheckReaderLen(r); err != nil {
		return nil, nil, err
	}
	opt := CloneObjectPutOptions(uopt)
	totalBytes, err := GetReaderLen(r)
	if err != nil && opt != nil && opt.Listener != nil {
		if opt.ContentLength == 0 {
			return nil, nil, err
		}
		totalBytes = opt.ContentLength
	}
	if err == nil {
		// ??? go http ????????????, ???bytes.Buffer/bytes.Reader/strings.Reader???????????????ContentLength, ????????? Chunk ??????
		if opt != nil && opt.ContentLength == 0 && IsLenReader(r) {
			opt.ContentLength = totalBytes
		}
	}
	reader := TeeReader(r, nil, totalBytes, nil)
	if s.client.Conf.EnableCRC {
		reader.writer = crc64.New(crc64.MakeTable(crc64.ECMA))
	}
	if opt != nil && opt.Listener != nil {
		reader.listener = opt.Listener
	}

	var res ImageProcessResult
	sendOpt := sendOptions{
		baseURL:   s.client.BaseURL.BucketURL,
		uri:       "/" + encodeURIComponent(name),
		method:    http.MethodPut,
		body:      reader,
		optHeader: opt,
		result:    &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)

	return &res, resp, err
}

// ci put object from local file
func (s *CIService) PutFromFile(ctx context.Context, name string, filePath string, opt *ObjectPutOptions) (*ImageProcessResult, *Response, error) {
	fd, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer fd.Close()

	return s.Put(ctx, name, fd, opt)
}

// ?????????????????? https://cloud.tencent.com/document/product/460/36540
// ?????????-??????????????? https://cloud.tencent.com/document/product/460/19017
func (s *CIService) Get(ctx context.Context, name string, operation string, opt *ObjectGetOptions, id ...string) (*Response, error) {
	var u string
	if len(id) == 1 {
		u = fmt.Sprintf("/%s?versionId=%s&%s", encodeURIComponent(name), id[0], operation)
	} else if len(id) == 0 {
		u = fmt.Sprintf("/%s?%s", encodeURIComponent(name), operation)
	} else {
		return nil, errors.New("wrong params")
	}

	sendOpt := sendOptions{
		baseURL:          s.client.BaseURL.BucketURL,
		uri:              u,
		method:           http.MethodGet,
		optQuery:         opt,
		optHeader:        opt,
		disableCloseBody: true,
	}
	resp, err := s.client.send(ctx, &sendOpt)

	if opt != nil && opt.Listener != nil {
		if err == nil && resp != nil {
			if totalBytes, e := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64); e == nil {
				resp.Body = TeeReader(resp.Body, nil, totalBytes, opt.Listener)
			}
		}
	}
	return resp, err
}

func (s *CIService) GetToFile(ctx context.Context, name, localpath, operation string, opt *ObjectGetOptions, id ...string) (*Response, error) {
	resp, err := s.Get(ctx, name, operation, opt, id...)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	// If file exist, overwrite it
	fd, err := os.OpenFile(localpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		return resp, err
	}

	_, err = io.Copy(fd, resp.Body)
	fd.Close()
	if err != nil {
		return resp, err
	}

	return resp, nil
}

type GetQRcodeResult struct {
	XMLName     xml.Name    `xml:"Response"`
	CodeStatus  int         `xml:"CodeStatus,omitempty"`
	QRcodeInfo  *QRcodeInfo `xml:"QRcodeInfo,omitempty"`
	ResultImage string      `xml:"ResultImage,omitempty"`
}

// ???????????????-??????????????? https://cloud.tencent.com/document/product/436/54070
func (s *CIService) GetQRcode(ctx context.Context, name string, cover int, opt *ObjectGetOptions, id ...string) (*GetQRcodeResult, *Response, error) {
	var u string
	if len(id) == 1 {
		u = fmt.Sprintf("/%s?versionId=%s&ci-process=QRcode&cover=%v", encodeURIComponent(name), id[0], cover)
	} else if len(id) == 0 {
		u = fmt.Sprintf("/%s?ci-process=QRcode&cover=%v", encodeURIComponent(name), cover)
	} else {
		return nil, nil, errors.New("wrong params")
	}

	var res GetQRcodeResult
	sendOpt := sendOptions{
		baseURL:   s.client.BaseURL.BucketURL,
		uri:       u,
		method:    http.MethodGet,
		optQuery:  opt,
		optHeader: opt,
		result:    &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

type GenerateQRcodeOptions struct {
	QRcodeContent string `url:"qrcode-content,omitempty"`
	Mode          int    `url:"mode,omitempty"`
	Width         int    `url:"width,omitempty"`
}
type GenerateQRcodeResult struct {
	XMLName     xml.Name `xml:"Response"`
	ResultImage string   `xml:"ResultImage,omitempty"`
}

// ??????????????? https://cloud.tencent.com/document/product/436/54071
func (s *CIService) GenerateQRcode(ctx context.Context, opt *GenerateQRcodeOptions) (*GenerateQRcodeResult, *Response, error) {
	var res GenerateQRcodeResult
	sendOpt := &sendOptions{
		baseURL:  s.client.BaseURL.BucketURL,
		uri:      "/?ci-process=qrcode-generate",
		method:   http.MethodGet,
		optQuery: opt,
		result:   &res,
	}
	resp, err := s.client.send(ctx, sendOpt)
	return &res, resp, err
}

func (s *CIService) GenerateQRcodeToFile(ctx context.Context, filePath string, opt *GenerateQRcodeOptions) (*GenerateQRcodeResult, *Response, error) {
	res, resp, err := s.GenerateQRcode(ctx, opt)
	if err != nil {
		return res, resp, err
	}

	// If file exist, overwrite it
	fd, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		return res, resp, err
	}
	defer fd.Close()

	bs, err := base64.StdEncoding.DecodeString(res.ResultImage)
	if err != nil {
		return res, resp, err
	}
	fb := bytes.NewReader(bs)
	_, err = io.Copy(fd, fb)

	return res, resp, err
}

// ?????? Guetzli ?????? https://cloud.tencent.com/document/product/460/30112
func (s *CIService) PutGuetzli(ctx context.Context) (*Response, error) {
	sendOpt := &sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/?guetzli",
		method:  http.MethodPut,
	}
	resp, err := s.client.send(ctx, sendOpt)
	return resp, err
}

type GetGuetzliResult struct {
	XMLName       xml.Name `xml:"GuetzliStatus"`
	GuetzliStatus string   `xml:",chardata"`
}

// ?????? Guetzli ?????? https://cloud.tencent.com/document/product/460/30111
func (s *CIService) GetGuetzli(ctx context.Context) (*GetGuetzliResult, *Response, error) {
	var res GetGuetzliResult
	sendOpt := &sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/?guetzli",
		method:  http.MethodGet,
		result:  &res,
	}
	resp, err := s.client.send(ctx, sendOpt)
	return &res, resp, err
}

// ?????? Guetzli ?????? https://cloud.tencent.com/document/product/460/30113
func (s *CIService) DeleteGuetzli(ctx context.Context) (*Response, error) {
	sendOpt := &sendOptions{
		baseURL: s.client.BaseURL.CIURL,
		uri:     "/?guetzli",
		method:  http.MethodDelete,
	}
	resp, err := s.client.send(ctx, sendOpt)
	return resp, err
}
