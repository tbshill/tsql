package tsql

import . "github.com/tbshill/goparsec"

const (
	BIGINTLiteral           = "BIGINT"
	BITLiteral              = "BIT"
	DECIMALLiteral          = "DECIMAL"
	INTLiteral              = "INT"
	MONEYLiteral            = "MONEY"
	NUMERICLiteral          = "NUMERIC"
	SMALLINTLiteral         = "SMALLINT"
	SMALLMONEYLiteral       = "SMALLMONEY"
	TINYINTLiteral          = "TINYINT"
	FLOATLiteral            = "FLOAT"
	REALLiteral             = "REAL"
	DATELiteral             = "DATE"
	DATETIME2Literal        = "DATETIME2"
	DATETIMELiteral         = "DATETIME"
	DATETIMEOFFSETLiteral   = "DATETIMEOFFSET"
	SMALLDATETIMELiteral    = "SMALLDATETIME"
	TIMELiteral             = "TIME"
	CHARLiteral             = "CHAR"
	VARCHARLiteral          = "VARCHAR"
	TEXTLiteral             = "TEXT"
	NCHARLiteral            = "NCHAR"
	NTEXTLiteral            = "NTEXT"
	NVARCHARLiteral         = "NVARCHAR"
	BINARYLiteral           = "BINARY"
	IMAGELiteral            = "IMAGE"
	VARBINARYLiteral        = "VARBINARY"
	ROWVERSIONLiteral       = "ROWVERSION"
	HIERARCHYIDLiteral      = "HIERARCHYID"
	UNIQUEIDENTIFIERLiteral = "UNIQUEIDENTIFIER"
	SQL_VARIANTLiteral      = "SQL_VARIANT"
	XMLLiteral              = "XML"
)

var (
	ExpectBIGINTLiteral           = ExpectCaseInsensitiveString(BIGINTLiteral)
	ExpectBITLiteral              = ExpectCaseInsensitiveString(BITLiteral)
	ExpectDECIMALLiteral          = ExpectCaseInsensitiveString(DECIMALLiteral)
	ExpectINTLiteral              = ExpectCaseInsensitiveString(INTLiteral)
	ExpectMONEYLiteral            = ExpectCaseInsensitiveString(MONEYLiteral)
	ExpectNUMERICLiteral          = ExpectCaseInsensitiveString(NUMERICLiteral)
	ExpectSMALLINTLiteral         = ExpectCaseInsensitiveString(SMALLINTLiteral)
	ExpectSMALLMONEYLiteral       = ExpectCaseInsensitiveString(SMALLMONEYLiteral)
	ExpectTINYINTLiteral          = ExpectCaseInsensitiveString(TINYINTLiteral)
	ExpectFLOATLiteral            = ExpectCaseInsensitiveString(FLOATLiteral)
	ExpectREALLiteral             = ExpectCaseInsensitiveString(REALLiteral)
	ExpectDATELiteral             = ExpectCaseInsensitiveString(DATELiteral)
	ExpectDATETIME2Literal        = ExpectCaseInsensitiveString(DATETIME2Literal)
	ExpectDATETIMELiteral         = ExpectCaseInsensitiveString(DATETIMELiteral)
	ExpectDATETIMEOFFSETLiteral   = ExpectCaseInsensitiveString(DATETIMEOFFSETLiteral)
	ExpectSMALLDATETIMELiteral    = ExpectCaseInsensitiveString(SMALLDATETIMELiteral)
	ExpectTIMELiteral             = ExpectCaseInsensitiveString(TIMELiteral)
	ExpectCHARLiteral             = ExpectCaseInsensitiveString(CHARLiteral)
	ExpectVARCHARLiteral          = ExpectCaseInsensitiveString(VARCHARLiteral)
	ExpectTEXTLiteral             = ExpectCaseInsensitiveString(TEXTLiteral)
	ExpectNCHARLiteral            = ExpectCaseInsensitiveString(NCHARLiteral)
	ExpectNTEXTLiteral            = ExpectCaseInsensitiveString(NTEXTLiteral)
	ExpectNVARCHARLiteral         = ExpectCaseInsensitiveString(NVARCHARLiteral)
	ExpectBINARYLiteral           = ExpectCaseInsensitiveString(BINARYLiteral)
	ExpectIMAGELiteral            = ExpectCaseInsensitiveString(IMAGELiteral)
	ExpectVARBINARYLiteral        = ExpectCaseInsensitiveString(VARBINARYLiteral)
	ExpectROWVERSIONLiteral       = ExpectCaseInsensitiveString(ROWVERSIONLiteral)
	ExpectHIERARCHYIDLiteral      = ExpectCaseInsensitiveString(HIERARCHYIDLiteral)
	ExpectUNIQUEIDENTIFIERLiteral = ExpectCaseInsensitiveString(UNIQUEIDENTIFIERLiteral)
	ExpectSQL_VARIANTLiteral      = ExpectCaseInsensitiveString(SQL_VARIANTLiteral)
	ExpectXMLLiteral              = ExpectCaseInsensitiveString(XMLLiteral)
)
