package models

// Define the request struct
type OCRRequest struct {
	MimeType      string   `json:"mimeType"`
	LanguageCodes []string `json:"languageCodes"`
	Model         string   `json:"model"`
	Content       string   `json:"content"`
}

type OCRAsyncResponse struct {
	Id          string      `json:"id"`
	Description string      `json:"description"`
	CreatedAt   string      `json:"createdAt"`
	CreatedBy   string      `json:"createdBy"`
	ModifiedAt  string      `json:"modifiedAt"`
	Done        bool        `json:"done"`
	Metadata    interface{} `json:"metadata,omitempty"`
}

type GetRecognitionResultResponseError struct {
	Error struct {
		GrpcCode   int           `json:"grpcCode"`
		HttpCode   int           `json:"httpCode"`
		Message    string        `json:"message"`
		HttpStatus string        `json:"httpStatus"`
		Details    []interface{} `json:"details"`
	} `json:"error,omitempty"`
}

type OCRResponse struct {
	Result struct {
		TextAnnotation TextAnnotation `json:"textAnnotation"`
		Page           int64          `json:"page,string"`
	} `json:"result"`
}

type TextAnnotation struct {
	Width    int64    `json:"width,string"`
	Height   int64    `json:"height,string"`
	Blocks   []Block  `json:"blocks"`
	Entities []Entity `json:"entities"`
	Tables   []Table  `json:"tables"`
	FullText string   `json:"fullText"`
	Rotate   string   `json:"rotate"`
}

type Block struct {
	BoundingBox  BoundingBox   `json:"boundingBox"`
	Lines        []Line        `json:"lines"`
	Languages    []Language    `json:"languages"`
	TextSegments []TextSegment `json:"textSegments"`
}

type BoundingBox struct {
	Vertices []Vertex `json:"vertices"`
}

type Vertex struct {
	X int64 `json:"x,string"`
	Y int64 `json:"y,string"`
}

type Line struct {
	BoundingBox  BoundingBox   `json:"boundingBox"`
	Text         string        `json:"text"`
	Words        []Word        `json:"words"`
	TextSegments []TextSegment `json:"textSegments"`
	Orientation  string        `json:"orientation"`
}

type Word struct {
	BoundingBox  BoundingBox   `json:"boundingBox"`
	Text         string        `json:"text"`
	EntityIndex  int64         `json:"entityIndex,string"`
	TextSegments []TextSegment `json:"textSegments"`
}

type TextSegment struct {
	StartIndex int64 `json:"startIndex,string"`
	Length     int64 `json:"length,string"`
}

type Language struct {
	LanguageCode string `json:"languageCode"`
}

type Entity struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

type Table struct {
	BoundingBox BoundingBox `json:"boundingBox"`
	RowCount    int64       `json:"rowCount,string"`
	ColumnCount int64       `json:"columnCount,string"`
	Cells       []Cell      `json:"cells"`
}

type Cell struct {
	BoundingBox  BoundingBox   `json:"boundingBox"`
	RowIndex     int64         `json:"rowIndex,string"`
	ColumnIndex  int64         `json:"columnIndex,string"`
	ColumnSpan   int64         `json:"columnSpan,string"`
	RowSpan      int64         `json:"rowSpan,string"`
	Text         string        `json:"text"`
	TextSegments []TextSegment `json:"textSegments"`
}
