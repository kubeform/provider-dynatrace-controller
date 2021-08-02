package dcrum_decoder

// Value The value to compare to.
type Value string

func (v Value) Ref() *Value {
	return &v
}

func (v *Value) String() string {
	return string(*v)
}

// Values offers the known enum values
var Values = struct {
	AllOther         Value
	CitrixAppFlow    Value
	CitrixIca        Value
	CitrixIcaOverSSL Value
	DB2Drda          Value
	HTTP             Value
	HTTPS            Value
	HTTPExpress      Value
	Informix         Value
	MySQL            Value
	Oracle           Value
	SAPGUI           Value
	SAPGUIOverHTTP   Value
	SAPGUIOverHTTPS  Value
	SAPHanaDB        Value
	SAPRfc           Value
	SSL              Value
	TDS              Value
}{
	"ALL_OTHER",
	"CITRIX_APPFLOW",
	"CITRIX_ICA",
	"CITRIX_ICA_OVER_SSL",
	"DB2_DRDA",
	"HTTP",
	"HTTPS",
	"HTTP_EXPRESS",
	"INFORMIX",
	"MYSQL",
	"ORACLE",
	"SAP_GUI",
	"SAP_GUI_OVER_HTTP",
	"SAP_GUI_OVER_HTTPS",
	"SAP_HANA_DB",
	"SAP_RFC",
	"SSL",
	"TDS",
}
