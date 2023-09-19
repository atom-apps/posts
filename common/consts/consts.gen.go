// Code generated by go-enum DO NOT EDIT.
// Version: -
// Revision: -
// Build Date: -
// Built By: -

package consts

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	// ArticleFormatHtml is a ArticleFormat of type Html.
	ArticleFormatHtml ArticleFormat = iota
	// ArticleFormatMarkdown is a ArticleFormat of type Markdown.
	ArticleFormatMarkdown
)

var ErrInvalidArticleFormat = fmt.Errorf("not a valid ArticleFormat, try [%s]", strings.Join(_ArticleFormatNames, ", "))

const _ArticleFormatName = "HtmlMarkdown"

var _ArticleFormatNames = []string{
	_ArticleFormatName[0:4],
	_ArticleFormatName[4:12],
}

// ArticleFormatNames returns a list of possible string values of ArticleFormat.
func ArticleFormatNames() []string {
	tmp := make([]string, len(_ArticleFormatNames))
	copy(tmp, _ArticleFormatNames)
	return tmp
}

// ArticleFormatValues returns a list of the values for ArticleFormat
func ArticleFormatValues() []ArticleFormat {
	return []ArticleFormat{
		ArticleFormatHtml,
		ArticleFormatMarkdown,
	}
}

var _ArticleFormatMap = map[ArticleFormat]string{
	ArticleFormatHtml:     _ArticleFormatName[0:4],
	ArticleFormatMarkdown: _ArticleFormatName[4:12],
}

// String implements the Stringer interface.
func (x ArticleFormat) String() string {
	if str, ok := _ArticleFormatMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ArticleFormat(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x ArticleFormat) IsValid() bool {
	_, ok := _ArticleFormatMap[x]
	return ok
}

var _ArticleFormatValue = map[string]ArticleFormat{
	_ArticleFormatName[0:4]:  ArticleFormatHtml,
	_ArticleFormatName[4:12]: ArticleFormatMarkdown,
}

// ParseArticleFormat attempts to convert a string to a ArticleFormat.
func ParseArticleFormat(name string) (ArticleFormat, error) {
	if x, ok := _ArticleFormatValue[name]; ok {
		return x, nil
	}
	return ArticleFormat(0), fmt.Errorf("%s is %w", name, ErrInvalidArticleFormat)
}

var errArticleFormatNilPtr = errors.New("value pointer is nil") // one per type for package clashes

// Scan implements the Scanner interface.
func (x *ArticleFormat) Scan(value interface{}) (err error) {
	if value == nil {
		*x = ArticleFormat(0)
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case int64:
		*x = ArticleFormat(v)
	case string:
		*x, err = ParseArticleFormat(v)
		if err != nil {
			// try parsing the integer value as a string
			if val, verr := strconv.Atoi(v); verr == nil {
				*x, err = ArticleFormat(val), nil
			}
		}
	case []byte:
		*x, err = ParseArticleFormat(string(v))
		if err != nil {
			// try parsing the integer value as a string
			if val, verr := strconv.Atoi(string(v)); verr == nil {
				*x, err = ArticleFormat(val), nil
			}
		}
	case ArticleFormat:
		*x = v
	case int:
		*x = ArticleFormat(v)
	case *ArticleFormat:
		if v == nil {
			return errArticleFormatNilPtr
		}
		*x = *v
	case uint:
		*x = ArticleFormat(v)
	case uint64:
		*x = ArticleFormat(v)
	case *int:
		if v == nil {
			return errArticleFormatNilPtr
		}
		*x = ArticleFormat(*v)
	case *int64:
		if v == nil {
			return errArticleFormatNilPtr
		}
		*x = ArticleFormat(*v)
	case float64: // json marshals everything as a float64 if it's a number
		*x = ArticleFormat(v)
	case *float64: // json marshals everything as a float64 if it's a number
		if v == nil {
			return errArticleFormatNilPtr
		}
		*x = ArticleFormat(*v)
	case *uint:
		if v == nil {
			return errArticleFormatNilPtr
		}
		*x = ArticleFormat(*v)
	case *uint64:
		if v == nil {
			return errArticleFormatNilPtr
		}
		*x = ArticleFormat(*v)
	case *string:
		if v == nil {
			return errArticleFormatNilPtr
		}
		*x, err = ParseArticleFormat(*v)
		if err != nil {
			// try parsing the integer value as a string
			if val, verr := strconv.Atoi(*v); verr == nil {
				*x, err = ArticleFormat(val), nil
			}
		}
	}

	return
}

// Value implements the driver Valuer interface.
func (x ArticleFormat) Value() (driver.Value, error) {
	return int64(x), nil
}

// Set implements the Golang flag.Value interface func.
func (x *ArticleFormat) Set(val string) error {
	v, err := ParseArticleFormat(val)
	*x = v
	return err
}

// Get implements the Golang flag.Getter interface func.
func (x *ArticleFormat) Get() interface{} {
	return *x
}

// Type implements the github.com/spf13/pFlag Value interface.
func (x *ArticleFormat) Type() string {
	return "ArticleFormat"
}

type NullArticleFormat struct {
	ArticleFormat ArticleFormat
	Valid         bool
}

func NewNullArticleFormat(val interface{}) (x NullArticleFormat) {
	x.Scan(val) // yes, we ignore this error, it will just be an invalid value.
	return
}

// Scan implements the Scanner interface.
func (x *NullArticleFormat) Scan(value interface{}) (err error) {
	if value == nil {
		x.ArticleFormat, x.Valid = ArticleFormat(0), false
		return
	}

	err = x.ArticleFormat.Scan(value)
	x.Valid = (err == nil)
	return
}

// Value implements the driver Valuer interface.
func (x NullArticleFormat) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	// driver.Value accepts int64 for int values.
	return int64(x.ArticleFormat), nil
}

type NullArticleFormatStr struct {
	NullArticleFormat
}

func NewNullArticleFormatStr(val interface{}) (x NullArticleFormatStr) {
	x.Scan(val) // yes, we ignore this error, it will just be an invalid value.
	return
}

// Value implements the driver Valuer interface.
func (x NullArticleFormatStr) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	return x.ArticleFormat.String(), nil
}

const (
	// ArticlePriceTypeContent is a ArticlePriceType of type Content.
	ArticlePriceTypeContent ArticlePriceType = iota
	// ArticlePriceTypeAttachment is a ArticlePriceType of type Attachment.
	ArticlePriceTypeAttachment
)

var ErrInvalidArticlePriceType = fmt.Errorf("not a valid ArticlePriceType, try [%s]", strings.Join(_ArticlePriceTypeNames, ", "))

const _ArticlePriceTypeName = "ContentAttachment"

var _ArticlePriceTypeNames = []string{
	_ArticlePriceTypeName[0:7],
	_ArticlePriceTypeName[7:17],
}

// ArticlePriceTypeNames returns a list of possible string values of ArticlePriceType.
func ArticlePriceTypeNames() []string {
	tmp := make([]string, len(_ArticlePriceTypeNames))
	copy(tmp, _ArticlePriceTypeNames)
	return tmp
}

// ArticlePriceTypeValues returns a list of the values for ArticlePriceType
func ArticlePriceTypeValues() []ArticlePriceType {
	return []ArticlePriceType{
		ArticlePriceTypeContent,
		ArticlePriceTypeAttachment,
	}
}

var _ArticlePriceTypeMap = map[ArticlePriceType]string{
	ArticlePriceTypeContent:    _ArticlePriceTypeName[0:7],
	ArticlePriceTypeAttachment: _ArticlePriceTypeName[7:17],
}

// String implements the Stringer interface.
func (x ArticlePriceType) String() string {
	if str, ok := _ArticlePriceTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ArticlePriceType(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x ArticlePriceType) IsValid() bool {
	_, ok := _ArticlePriceTypeMap[x]
	return ok
}

var _ArticlePriceTypeValue = map[string]ArticlePriceType{
	_ArticlePriceTypeName[0:7]:  ArticlePriceTypeContent,
	_ArticlePriceTypeName[7:17]: ArticlePriceTypeAttachment,
}

// ParseArticlePriceType attempts to convert a string to a ArticlePriceType.
func ParseArticlePriceType(name string) (ArticlePriceType, error) {
	if x, ok := _ArticlePriceTypeValue[name]; ok {
		return x, nil
	}
	return ArticlePriceType(0), fmt.Errorf("%s is %w", name, ErrInvalidArticlePriceType)
}

var errArticlePriceTypeNilPtr = errors.New("value pointer is nil") // one per type for package clashes

// Scan implements the Scanner interface.
func (x *ArticlePriceType) Scan(value interface{}) (err error) {
	if value == nil {
		*x = ArticlePriceType(0)
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case int64:
		*x = ArticlePriceType(v)
	case string:
		*x, err = ParseArticlePriceType(v)
		if err != nil {
			// try parsing the integer value as a string
			if val, verr := strconv.Atoi(v); verr == nil {
				*x, err = ArticlePriceType(val), nil
			}
		}
	case []byte:
		*x, err = ParseArticlePriceType(string(v))
		if err != nil {
			// try parsing the integer value as a string
			if val, verr := strconv.Atoi(string(v)); verr == nil {
				*x, err = ArticlePriceType(val), nil
			}
		}
	case ArticlePriceType:
		*x = v
	case int:
		*x = ArticlePriceType(v)
	case *ArticlePriceType:
		if v == nil {
			return errArticlePriceTypeNilPtr
		}
		*x = *v
	case uint:
		*x = ArticlePriceType(v)
	case uint64:
		*x = ArticlePriceType(v)
	case *int:
		if v == nil {
			return errArticlePriceTypeNilPtr
		}
		*x = ArticlePriceType(*v)
	case *int64:
		if v == nil {
			return errArticlePriceTypeNilPtr
		}
		*x = ArticlePriceType(*v)
	case float64: // json marshals everything as a float64 if it's a number
		*x = ArticlePriceType(v)
	case *float64: // json marshals everything as a float64 if it's a number
		if v == nil {
			return errArticlePriceTypeNilPtr
		}
		*x = ArticlePriceType(*v)
	case *uint:
		if v == nil {
			return errArticlePriceTypeNilPtr
		}
		*x = ArticlePriceType(*v)
	case *uint64:
		if v == nil {
			return errArticlePriceTypeNilPtr
		}
		*x = ArticlePriceType(*v)
	case *string:
		if v == nil {
			return errArticlePriceTypeNilPtr
		}
		*x, err = ParseArticlePriceType(*v)
		if err != nil {
			// try parsing the integer value as a string
			if val, verr := strconv.Atoi(*v); verr == nil {
				*x, err = ArticlePriceType(val), nil
			}
		}
	}

	return
}

// Value implements the driver Valuer interface.
func (x ArticlePriceType) Value() (driver.Value, error) {
	return int64(x), nil
}

// Set implements the Golang flag.Value interface func.
func (x *ArticlePriceType) Set(val string) error {
	v, err := ParseArticlePriceType(val)
	*x = v
	return err
}

// Get implements the Golang flag.Getter interface func.
func (x *ArticlePriceType) Get() interface{} {
	return *x
}

// Type implements the github.com/spf13/pFlag Value interface.
func (x *ArticlePriceType) Type() string {
	return "ArticlePriceType"
}

type NullArticlePriceType struct {
	ArticlePriceType ArticlePriceType
	Valid            bool
}

func NewNullArticlePriceType(val interface{}) (x NullArticlePriceType) {
	x.Scan(val) // yes, we ignore this error, it will just be an invalid value.
	return
}

// Scan implements the Scanner interface.
func (x *NullArticlePriceType) Scan(value interface{}) (err error) {
	if value == nil {
		x.ArticlePriceType, x.Valid = ArticlePriceType(0), false
		return
	}

	err = x.ArticlePriceType.Scan(value)
	x.Valid = (err == nil)
	return
}

// Value implements the driver Valuer interface.
func (x NullArticlePriceType) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	// driver.Value accepts int64 for int values.
	return int64(x.ArticlePriceType), nil
}

type NullArticlePriceTypeStr struct {
	NullArticlePriceType
}

func NewNullArticlePriceTypeStr(val interface{}) (x NullArticlePriceTypeStr) {
	x.Scan(val) // yes, we ignore this error, it will just be an invalid value.
	return
}

// Value implements the driver Valuer interface.
func (x NullArticlePriceTypeStr) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	return x.ArticlePriceType.String(), nil
}

const (
	// ArticleTypeText is a ArticleType of type Text.
	ArticleTypeText ArticleType = iota
	// ArticleTypeImage is a ArticleType of type Image.
	ArticleTypeImage
	// ArticleTypeAudio is a ArticleType of type Audio.
	ArticleTypeAudio
	// ArticleTypeVideo is a ArticleType of type Video.
	ArticleTypeVideo
)

var ErrInvalidArticleType = fmt.Errorf("not a valid ArticleType, try [%s]", strings.Join(_ArticleTypeNames, ", "))

const _ArticleTypeName = "TextImageAudioVideo"

var _ArticleTypeNames = []string{
	_ArticleTypeName[0:4],
	_ArticleTypeName[4:9],
	_ArticleTypeName[9:14],
	_ArticleTypeName[14:19],
}

// ArticleTypeNames returns a list of possible string values of ArticleType.
func ArticleTypeNames() []string {
	tmp := make([]string, len(_ArticleTypeNames))
	copy(tmp, _ArticleTypeNames)
	return tmp
}

// ArticleTypeValues returns a list of the values for ArticleType
func ArticleTypeValues() []ArticleType {
	return []ArticleType{
		ArticleTypeText,
		ArticleTypeImage,
		ArticleTypeAudio,
		ArticleTypeVideo,
	}
}

var _ArticleTypeMap = map[ArticleType]string{
	ArticleTypeText:  _ArticleTypeName[0:4],
	ArticleTypeImage: _ArticleTypeName[4:9],
	ArticleTypeAudio: _ArticleTypeName[9:14],
	ArticleTypeVideo: _ArticleTypeName[14:19],
}

// String implements the Stringer interface.
func (x ArticleType) String() string {
	if str, ok := _ArticleTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ArticleType(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x ArticleType) IsValid() bool {
	_, ok := _ArticleTypeMap[x]
	return ok
}

var _ArticleTypeValue = map[string]ArticleType{
	_ArticleTypeName[0:4]:   ArticleTypeText,
	_ArticleTypeName[4:9]:   ArticleTypeImage,
	_ArticleTypeName[9:14]:  ArticleTypeAudio,
	_ArticleTypeName[14:19]: ArticleTypeVideo,
}

// ParseArticleType attempts to convert a string to a ArticleType.
func ParseArticleType(name string) (ArticleType, error) {
	if x, ok := _ArticleTypeValue[name]; ok {
		return x, nil
	}
	return ArticleType(0), fmt.Errorf("%s is %w", name, ErrInvalidArticleType)
}

var errArticleTypeNilPtr = errors.New("value pointer is nil") // one per type for package clashes

// Scan implements the Scanner interface.
func (x *ArticleType) Scan(value interface{}) (err error) {
	if value == nil {
		*x = ArticleType(0)
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case int64:
		*x = ArticleType(v)
	case string:
		*x, err = ParseArticleType(v)
		if err != nil {
			// try parsing the integer value as a string
			if val, verr := strconv.Atoi(v); verr == nil {
				*x, err = ArticleType(val), nil
			}
		}
	case []byte:
		*x, err = ParseArticleType(string(v))
		if err != nil {
			// try parsing the integer value as a string
			if val, verr := strconv.Atoi(string(v)); verr == nil {
				*x, err = ArticleType(val), nil
			}
		}
	case ArticleType:
		*x = v
	case int:
		*x = ArticleType(v)
	case *ArticleType:
		if v == nil {
			return errArticleTypeNilPtr
		}
		*x = *v
	case uint:
		*x = ArticleType(v)
	case uint64:
		*x = ArticleType(v)
	case *int:
		if v == nil {
			return errArticleTypeNilPtr
		}
		*x = ArticleType(*v)
	case *int64:
		if v == nil {
			return errArticleTypeNilPtr
		}
		*x = ArticleType(*v)
	case float64: // json marshals everything as a float64 if it's a number
		*x = ArticleType(v)
	case *float64: // json marshals everything as a float64 if it's a number
		if v == nil {
			return errArticleTypeNilPtr
		}
		*x = ArticleType(*v)
	case *uint:
		if v == nil {
			return errArticleTypeNilPtr
		}
		*x = ArticleType(*v)
	case *uint64:
		if v == nil {
			return errArticleTypeNilPtr
		}
		*x = ArticleType(*v)
	case *string:
		if v == nil {
			return errArticleTypeNilPtr
		}
		*x, err = ParseArticleType(*v)
		if err != nil {
			// try parsing the integer value as a string
			if val, verr := strconv.Atoi(*v); verr == nil {
				*x, err = ArticleType(val), nil
			}
		}
	}

	return
}

// Value implements the driver Valuer interface.
func (x ArticleType) Value() (driver.Value, error) {
	return int64(x), nil
}

// Set implements the Golang flag.Value interface func.
func (x *ArticleType) Set(val string) error {
	v, err := ParseArticleType(val)
	*x = v
	return err
}

// Get implements the Golang flag.Getter interface func.
func (x *ArticleType) Get() interface{} {
	return *x
}

// Type implements the github.com/spf13/pFlag Value interface.
func (x *ArticleType) Type() string {
	return "ArticleType"
}

type NullArticleType struct {
	ArticleType ArticleType
	Valid       bool
}

func NewNullArticleType(val interface{}) (x NullArticleType) {
	x.Scan(val) // yes, we ignore this error, it will just be an invalid value.
	return
}

// Scan implements the Scanner interface.
func (x *NullArticleType) Scan(value interface{}) (err error) {
	if value == nil {
		x.ArticleType, x.Valid = ArticleType(0), false
		return
	}

	err = x.ArticleType.Scan(value)
	x.Valid = (err == nil)
	return
}

// Value implements the driver Valuer interface.
func (x NullArticleType) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	// driver.Value accepts int64 for int values.
	return int64(x.ArticleType), nil
}

type NullArticleTypeStr struct {
	NullArticleType
}

func NewNullArticleTypeStr(val interface{}) (x NullArticleTypeStr) {
	x.Scan(val) // yes, we ignore this error, it will just be an invalid value.
	return
}

// Value implements the driver Valuer interface.
func (x NullArticleTypeStr) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	return x.ArticleType.String(), nil
}

const (
	// AudioProviderUrl is a AudioProvider of type url.
	AudioProviderUrl AudioProvider = "url"
	// AudioProviderXimalaya is a AudioProvider of type ximalaya.
	AudioProviderXimalaya AudioProvider = "ximalaya"
)

var ErrInvalidAudioProvider = fmt.Errorf("not a valid AudioProvider, try [%s]", strings.Join(_AudioProviderNames, ", "))

var _AudioProviderNames = []string{
	string(AudioProviderUrl),
	string(AudioProviderXimalaya),
}

// AudioProviderNames returns a list of possible string values of AudioProvider.
func AudioProviderNames() []string {
	tmp := make([]string, len(_AudioProviderNames))
	copy(tmp, _AudioProviderNames)
	return tmp
}

// AudioProviderValues returns a list of the values for AudioProvider
func AudioProviderValues() []AudioProvider {
	return []AudioProvider{
		AudioProviderUrl,
		AudioProviderXimalaya,
	}
}

// String implements the Stringer interface.
func (x AudioProvider) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x AudioProvider) IsValid() bool {
	_, err := ParseAudioProvider(string(x))
	return err == nil
}

var _AudioProviderValue = map[string]AudioProvider{
	"url":      AudioProviderUrl,
	"ximalaya": AudioProviderXimalaya,
}

// ParseAudioProvider attempts to convert a string to a AudioProvider.
func ParseAudioProvider(name string) (AudioProvider, error) {
	if x, ok := _AudioProviderValue[name]; ok {
		return x, nil
	}
	return AudioProvider(""), fmt.Errorf("%s is %w", name, ErrInvalidAudioProvider)
}

var errAudioProviderNilPtr = errors.New("value pointer is nil") // one per type for package clashes

// Scan implements the Scanner interface.
func (x *AudioProvider) Scan(value interface{}) (err error) {
	if value == nil {
		*x = AudioProvider("")
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case string:
		*x, err = ParseAudioProvider(v)
	case []byte:
		*x, err = ParseAudioProvider(string(v))
	case AudioProvider:
		*x = v
	case *AudioProvider:
		if v == nil {
			return errAudioProviderNilPtr
		}
		*x = *v
	case *string:
		if v == nil {
			return errAudioProviderNilPtr
		}
		*x, err = ParseAudioProvider(*v)
	default:
		return errors.New("invalid type for AudioProvider")
	}

	return
}

// Value implements the driver Valuer interface.
func (x AudioProvider) Value() (driver.Value, error) {
	return x.String(), nil
}

// Set implements the Golang flag.Value interface func.
func (x *AudioProvider) Set(val string) error {
	v, err := ParseAudioProvider(val)
	*x = v
	return err
}

// Get implements the Golang flag.Getter interface func.
func (x *AudioProvider) Get() interface{} {
	return *x
}

// Type implements the github.com/spf13/pFlag Value interface.
func (x *AudioProvider) Type() string {
	return "AudioProvider"
}

type NullAudioProvider struct {
	AudioProvider AudioProvider
	Valid         bool
}

func NewNullAudioProvider(val interface{}) (x NullAudioProvider) {
	err := x.Scan(val) // yes, we ignore this error, it will just be an invalid value.
	_ = err            // make any errcheck linters happy
	return
}

// Scan implements the Scanner interface.
func (x *NullAudioProvider) Scan(value interface{}) (err error) {
	if value == nil {
		x.AudioProvider, x.Valid = AudioProvider(""), false
		return
	}

	err = x.AudioProvider.Scan(value)
	x.Valid = (err == nil)
	return
}

// Value implements the driver Valuer interface.
func (x NullAudioProvider) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	// driver.Value accepts int64 for int values.
	return string(x.AudioProvider), nil
}

type NullAudioProviderStr struct {
	NullAudioProvider
}

func NewNullAudioProviderStr(val interface{}) (x NullAudioProviderStr) {
	x.Scan(val) // yes, we ignore this error, it will just be an invalid value.
	return
}

// Value implements the driver Valuer interface.
func (x NullAudioProviderStr) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	return x.AudioProvider.String(), nil
}

const (
	// VideoProviderUrl is a VideoProvider of type url.
	VideoProviderUrl VideoProvider = "url"
	// VideoProviderBilibili is a VideoProvider of type bilibili.
	VideoProviderBilibili VideoProvider = "bilibili"
	// VideoProviderTencent is a VideoProvider of type tencent.
	VideoProviderTencent VideoProvider = "tencent"
	// VideoProviderYouku is a VideoProvider of type youku.
	VideoProviderYouku VideoProvider = "youku"
	// VideoProviderYoutube is a VideoProvider of type youtube.
	VideoProviderYoutube VideoProvider = "youtube"
	// VideoProviderIqiyi is a VideoProvider of type iqiyi.
	VideoProviderIqiyi VideoProvider = "iqiyi"
	// VideoProviderSohu is a VideoProvider of type sohu.
	VideoProviderSohu VideoProvider = "sohu"
	// VideoProviderXigua is a VideoProvider of type xigua.
	VideoProviderXigua VideoProvider = "xigua"
	// VideoProviderDouyin is a VideoProvider of type douyin.
	VideoProviderDouyin VideoProvider = "douyin"
	// VideoProviderWeishi is a VideoProvider of type weishi.
	VideoProviderWeishi VideoProvider = "weishi"
	// VideoProviderWeibo is a VideoProvider of type weibo.
	VideoProviderWeibo VideoProvider = "weibo"
)

var ErrInvalidVideoProvider = fmt.Errorf("not a valid VideoProvider, try [%s]", strings.Join(_VideoProviderNames, ", "))

var _VideoProviderNames = []string{
	string(VideoProviderUrl),
	string(VideoProviderBilibili),
	string(VideoProviderTencent),
	string(VideoProviderYouku),
	string(VideoProviderYoutube),
	string(VideoProviderIqiyi),
	string(VideoProviderSohu),
	string(VideoProviderXigua),
	string(VideoProviderDouyin),
	string(VideoProviderWeishi),
	string(VideoProviderWeibo),
}

// VideoProviderNames returns a list of possible string values of VideoProvider.
func VideoProviderNames() []string {
	tmp := make([]string, len(_VideoProviderNames))
	copy(tmp, _VideoProviderNames)
	return tmp
}

// VideoProviderValues returns a list of the values for VideoProvider
func VideoProviderValues() []VideoProvider {
	return []VideoProvider{
		VideoProviderUrl,
		VideoProviderBilibili,
		VideoProviderTencent,
		VideoProviderYouku,
		VideoProviderYoutube,
		VideoProviderIqiyi,
		VideoProviderSohu,
		VideoProviderXigua,
		VideoProviderDouyin,
		VideoProviderWeishi,
		VideoProviderWeibo,
	}
}

// String implements the Stringer interface.
func (x VideoProvider) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x VideoProvider) IsValid() bool {
	_, err := ParseVideoProvider(string(x))
	return err == nil
}

var _VideoProviderValue = map[string]VideoProvider{
	"url":      VideoProviderUrl,
	"bilibili": VideoProviderBilibili,
	"tencent":  VideoProviderTencent,
	"youku":    VideoProviderYouku,
	"youtube":  VideoProviderYoutube,
	"iqiyi":    VideoProviderIqiyi,
	"sohu":     VideoProviderSohu,
	"xigua":    VideoProviderXigua,
	"douyin":   VideoProviderDouyin,
	"weishi":   VideoProviderWeishi,
	"weibo":    VideoProviderWeibo,
}

// ParseVideoProvider attempts to convert a string to a VideoProvider.
func ParseVideoProvider(name string) (VideoProvider, error) {
	if x, ok := _VideoProviderValue[name]; ok {
		return x, nil
	}
	return VideoProvider(""), fmt.Errorf("%s is %w", name, ErrInvalidVideoProvider)
}

var errVideoProviderNilPtr = errors.New("value pointer is nil") // one per type for package clashes

// Scan implements the Scanner interface.
func (x *VideoProvider) Scan(value interface{}) (err error) {
	if value == nil {
		*x = VideoProvider("")
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case string:
		*x, err = ParseVideoProvider(v)
	case []byte:
		*x, err = ParseVideoProvider(string(v))
	case VideoProvider:
		*x = v
	case *VideoProvider:
		if v == nil {
			return errVideoProviderNilPtr
		}
		*x = *v
	case *string:
		if v == nil {
			return errVideoProviderNilPtr
		}
		*x, err = ParseVideoProvider(*v)
	default:
		return errors.New("invalid type for VideoProvider")
	}

	return
}

// Value implements the driver Valuer interface.
func (x VideoProvider) Value() (driver.Value, error) {
	return x.String(), nil
}

// Set implements the Golang flag.Value interface func.
func (x *VideoProvider) Set(val string) error {
	v, err := ParseVideoProvider(val)
	*x = v
	return err
}

// Get implements the Golang flag.Getter interface func.
func (x *VideoProvider) Get() interface{} {
	return *x
}

// Type implements the github.com/spf13/pFlag Value interface.
func (x *VideoProvider) Type() string {
	return "VideoProvider"
}

type NullVideoProvider struct {
	VideoProvider VideoProvider
	Valid         bool
}

func NewNullVideoProvider(val interface{}) (x NullVideoProvider) {
	err := x.Scan(val) // yes, we ignore this error, it will just be an invalid value.
	_ = err            // make any errcheck linters happy
	return
}

// Scan implements the Scanner interface.
func (x *NullVideoProvider) Scan(value interface{}) (err error) {
	if value == nil {
		x.VideoProvider, x.Valid = VideoProvider(""), false
		return
	}

	err = x.VideoProvider.Scan(value)
	x.Valid = (err == nil)
	return
}

// Value implements the driver Valuer interface.
func (x NullVideoProvider) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	// driver.Value accepts int64 for int values.
	return string(x.VideoProvider), nil
}

type NullVideoProviderStr struct {
	NullVideoProvider
}

func NewNullVideoProviderStr(val interface{}) (x NullVideoProviderStr) {
	x.Scan(val) // yes, we ignore this error, it will just be an invalid value.
	return
}

// Value implements the driver Valuer interface.
func (x NullVideoProviderStr) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	return x.VideoProvider.String(), nil
}
