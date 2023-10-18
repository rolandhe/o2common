package o2common

type IdReq struct {
	Id int64 `json:"id" binding:"required,gte=1,lte=120"`
}

type IdListReq struct {
	IdList []int64 `json:"idList" binding:"required,lte>0"`
}

type StringReq struct {
	Value string `json:"value" binding:"required"`
}

type StringListReq struct {
	Values string `json:"values" binding:"required,lte>0"`
}
