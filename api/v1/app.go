package v1

type CreateAppRequest struct {
	AppName string `json:"app_name" binding:"required"`
}

type EditAppRequest struct {
	AppId   uint   `json:"app_id"`
	AppName string `json:"app_name" binding:"required"`
	Status  int64  `json:"status"`
}

type DeleteAppRequest struct {
	AppId uint `json:"app_id" binding:"required"`
}

type GetAppRequest struct {
	AppId uint `json:"app_id" binding:"required"`
}

type GetAppsRequest struct {
	Status  int64  `json:"status"`
	Keyword string `json:"keyword"`
}

type GetAppResponse struct {
	GetAppResponseData
}

type GetAppsResponse struct {
	List []GetAppResponseData `json:"list"`
	Page int64                `json:"page"`
}

type GetAppResponseData struct {
	AppId   uint   `json:"app_id"`
	AppName string `json:"app_name"`
	Status  int64  `json:"status"`
}
