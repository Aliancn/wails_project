package controller

import (
	"changeme/Dao"
	"changeme/utils"
	"time"
)

type CountDown struct {
	Duration time.Duration `json:"duration"`
	Content  string        `json:"content"`
	DeadTime int64         `json:"deadTime"`
	Active   bool          `json:"active"`
	Title    string        `json:"title"`
	ID       uint          `json:"id"`
}

func (a *App) GetCountDown() *utils.Resp {
	var result []CountDown

	dao := Dao.CountDownDao{}
	var tmp, err = dao.GetAll(DB)
	if err != nil {
		return utils.Error("get count down error")
	}
	for i := range tmp {
		var ttmp = CountDown{
			Content:  tmp[i].Content,
			DeadTime: tmp[i].DeadTime.Unix(),
			Active:   tmp[i].Active,
			Title:    tmp[i].Title,
			ID:       tmp[i].ID,
		}
		ttmp.Duration = tmp[i].DeadTime.Sub(time.Now())
		if ttmp.Duration < 0 {
			err := tmp[i].Delete(DB)
			if err != nil {
				return utils.Error("delete count down error")
			}
			continue
		}
		// 转换为hao秒为单位
		ttmp.Duration = ttmp.Duration / time.Millisecond
		result = append(result, ttmp)
	}

	return utils.Success(result)
}

func (a *App) CreateCountDown(cd CountDown) *utils.Resp {

	var ttmp = Dao.CountDownDao{
		Content:  cd.Content,
		DeadTime: time.Unix(cd.DeadTime/1000, 0),
		Active:   cd.Active,
		Title:    cd.Title,
	}

	err := ttmp.Create(DB)
	if err != nil {
		return utils.Error("create count down error")
	}
	return utils.Success(nil)
}

func (a *App) DeleteCountDown(id uint) *utils.Resp {
	var ttmp = Dao.CountDownDao{}
	ttmp.ID = id
	err := ttmp.Delete(DB)
	if err != nil {
		return utils.Error("delete count down error")
	}
	return utils.Success(nil)
}
