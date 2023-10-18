package o2common

type IdReq struct {
	Id int64 `json:"id" binding:"required,gte=1"`
}

type IdListReq struct {
	IdList []int64 `json:"idList" binding:"required,lt=0"`
}

type StringReq struct {
	Value string `json:"value" binding:"required"`
}

type StringListReq struct {
	Values string `json:"values" binding:"required,lt=0"`
}

type Pager struct {
	PageNo   int32 `json:"pageNo" binding:"required"`
	PageSize int32 `json:"pageSize" binding:"required,lt=0"`
}
