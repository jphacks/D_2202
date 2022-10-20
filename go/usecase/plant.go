package usecase

import "narcissus/errors"

type Plant struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Url  string `json:"url" db:"hash`
}

func ListPlant() ([]Plant, error) {
	plants, err := DbPlant.ListPlant()
	//TODO 画像の保存先とか拡張子が決まったら変更する
	if err != nil {
		return nil, errors.ErrorWrap(err)
	}
	for i, v := range plants {
		plants[i].Url = "http://localhost:8080/figure/" + v.Url + ".png"
	}
	return plants, nil
}

// tagを指定すると、該当する植物を返してくれる関数
// tagは必須で含むべきものと、任意で含むべきものの2種類で指定できる
func SearchPlant(necessary_tags []int, optional_tags []int) ([]Plant, error) {
	plants, err := DbPlant.SearchPlant(necessary_tags, optional_tags)
	if err != nil {
		return nil, errors.ErrorWrap(err)
	}
	return plants, nil
}

// Plant型(idは適当で良い)を渡すと新たに追加してくれる
// 返り値は (存在していたか, 登録後のid, error)
func InsertPlant(plant Plant) (bool, int, error) {
	isNew, newId, err := DbPlant.InsertPlant(plant)
	if err != nil {
		return false, -1, errors.ErrorWrap(err)
	}
	return isNew, newId, nil
}

// 植物idとタグ名のスライスを渡すと、該当する植物にタグを追加する
// isAddTagがtrueなら、存在しないタグを新たにtagテーブルに追加する
func SetTagsToPlant(id int, tags []string) error {
	err := DbPlant.SetTagsToPlant(id, tags)
	if err != nil {
		return errors.ErrorWrap(err)
	}
	return nil
}
