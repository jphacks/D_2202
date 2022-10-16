// このファイルを実行するとデモが行えます
// 実行コマンド（コピペ用）
// go run demos.go label_detector.go landmark_detector.go web_detector.go

package main

// import "io"

func main() {
	// ラベル抽出
	filename_label := "./img/platanus.jpg"
	println("################################################################")
	println("Now starting a label-detect demo .... ")
	println("################################################################")
	DetectLabels(filename_label)

	// ランドマーク抽出
	filename_landmark := "./img/farm_tomita.jpg"
	println("################################################################")
	println("Now starting a landmark-detect demo .... ")
	println("################################################################")
	detectLandmarks(filename_landmark)

	// ウェブ検索
	filename_web := "./img/fish_mint.jpg"
	println("################################################################")
	println("Now starting a web-detect demo .... ")
	println("################################################################")
	detectLandmarks(filename_web)
}
