package nosql

import (
	"context"
	"narcissus/usecase"
	"strconv"
	"time"
)

func createInitData(ctx context.Context) {
	logger.Info("create init data in firestore")
	tags := []usecase.Tag{
		{Id: 11, Name: "春"},
		{Id: 12, Name: "夏"},
		{Id: 13, Name: "秋"},
		{Id: 14, Name: "冬"},
		{Id: 15, Name: "梅雨"},
		{Id: 16, Name: "通年"},
		{Id: 10001, Name: "木"},
		{Id: 10002, Name: "草"},
		{Id: 10003, Name: "花"},
		{Id: 10004, Name: "水草"},
		{Id: 10005, Name: "海藻"},
		{Id: 20001, Name: "コケ植物"},
		{Id: 20002, Name: "多肉植物"},
		{Id: 20003, Name: "つる植物"},
		{Id: 20004, Name: "シダ植物"},
		{Id: 20005, Name: "種子植物"},
		{Id: 20006, Name: "寄生植物"},
		{Id: 20007, Name: "観葉植物"},
		{Id: 20008, Name: "食虫植物"},
		{Id: 20009, Name: "裸子植物"},
		{Id: 20010, Name: "被子植物"},
		{Id: 20011, Name: "単子葉類"},
		{Id: 20012, Name: "双子葉類"},
		{Id: 30001, Name: "水辺"},
		{Id: 30002, Name: "海浜"},
		{Id: 30003, Name: "渓流"},
		{Id: 30004, Name: "高山"},
		{Id: 30005, Name: "熱帯雨林"},
		{Id: 30006, Name: "マングローブ"},
		{Id: 40001, Name: "野菜"},
		{Id: 40002, Name: "果物"},
		{Id: 40003, Name: "キノコ"},
		{Id: 40004, Name: "穀物"},
		{Id: 40005, Name: "農作物"},
		{Id: 50001, Name: "有毒"},
		{Id: 50002, Name: "美味しい"},
		{Id: 50003, Name: "不味い"},
		{Id: 50004, Name: "きれい"},
		{Id: 50005, Name: "大きい"},
		{Id: 50006, Name: "小さい"},
		{Id: 60001, Name: "赤"},
		{Id: 60002, Name: "青"},
		{Id: 60003, Name: "緑"},
		{Id: 60004, Name: "黄"},
		{Id: 60005, Name: "紫"},
		{Id: 60006, Name: "白"},
		{Id: 60007, Name: "黒"},
		{Id: 70001, Name: "アオイ科"},
		{Id: 70002, Name: "アオギリ科"},
		{Id: 70003, Name: "アカザ科"},
		{Id: 70004, Name: "アカネ科"},
		{Id: 70005, Name: "アカバナ科"},
		{Id: 70006, Name: "アケビ科"},
		{Id: 70007, Name: "アサ科"},
		{Id: 70008, Name: "アジサイ科"},
		{Id: 70009, Name: "アブラナ科"},
		{Id: 70010, Name: "アマモ科"},
		{Id: 70011, Name: "アヤメ科"},
		{Id: 70012, Name: "アリノトウグサ科"},
		{Id: 70013, Name: "アルストロメリア科"},
		{Id: 70014, Name: "アロエ科"},
		{Id: 70015, Name: "イイギリ科"},
		{Id: 70016, Name: "イキシオリリオン科"},
		{Id: 70017, Name: "イグサ科"},
		{Id: 70018, Name: "イソマツ科"},
		{Id: 70019, Name: "イチイ科"},
		{Id: 70020, Name: "イヌガヤ科"},
		{Id: 70021, Name: "イヌサフラン科"},
		{Id: 70022, Name: "イネ科"},
		{Id: 70023, Name: "イラクサ科"},
		{Id: 70024, Name: "イワウメ科"},
		{Id: 70025, Name: "イワタバコ科"},
		{Id: 70026, Name: "ウォキシア科"},
		{Id: 70027, Name: "ウキクサ科"},
		{Id: 70028, Name: "ウコギ科"},
		{Id: 70029, Name: "ウマノスズクサ科"},
		{Id: 70030, Name: "ウリ科"},
		{Id: 70031, Name: "ウルシ科"},
		{Id: 70032, Name: "エウポマティア科"},
		{Id: 70033, Name: "エゴノキ科"},
		{Id: 70034, Name: "オウムバナ科"},
		{Id: 70035, Name: "オオバコ科"},
		{Id: 70036, Name: "オシロイバナ科"},
		{Id: 70037, Name: "オトギリソウ科"},
		{Id: 70038, Name: "オミナエシ科"},
		{Id: 70039, Name: "オモダカ科"},
		{Id: 70040, Name: "カエデ科"},
		{Id: 70041, Name: "ガガイモ科"},
		{Id: 70042, Name: "カキノキ科"},
		{Id: 70043, Name: "カタバミ科"},
		{Id: 70044, Name: "カツラ科"},
		{Id: 70045, Name: "カネラ科"},
		{Id: 70046, Name: "カバノキ科"},
		{Id: 70047, Name: "ガマ科"},
		{Id: 70048, Name: "カヤツリグサ科"},
		{Id: 70049, Name: "カワゴケソウ科"},
		{Id: 70050, Name: "カンナ科"},
		{Id: 70051, Name: "キキョウ科"},
		{Id: 70052, Name: "キク科"},
		{Id: 70053, Name: "キツネノマゴ科"},
		{Id: 70054, Name: "キブシ科"},
		{Id: 70055, Name: "キョウチクトウ科"},
		{Id: 70056, Name: "ギョリュウ科"},
		{Id: 70057, Name: "キントラノオ科"},
		{Id: 70058, Name: "キンバイザサ科"},
		{Id: 70059, Name: "キンポウゲ科"},
		{Id: 70060, Name: "クサトベラ科"},
		{Id: 70061, Name: "クスノキ科"},
		{Id: 70062, Name: "クノニア科"},
		{Id: 70063, Name: "クマツヅラ科"},
		{Id: 70064, Name: "グミ科"},
		{Id: 70065, Name: "クルミ科"},
		{Id: 70066, Name: "クロウメモドキ科"},
		{Id: 70067, Name: "クロタキカズラ科"},
		{Id: 70068, Name: "クワ科"},
		{Id: 70069, Name: "グンネラ科"},
		{Id: 70070, Name: "ケシ科"},
		{Id: 70071, Name: "ケマンソウ科"},
		{Id: 70072, Name: "ゴクラクチョウカ科"},
		{Id: 70073, Name: "コショウ科"},
		{Id: 70074, Name: "ゴマノハグサ科"},
		{Id: 70075, Name: "ゴマ科"},
		{Id: 70076, Name: "サクラソウ科"},
		{Id: 70077, Name: "サトイモ科"},
		{Id: 70078, Name: "サボテン科"},
		{Id: 70079, Name: "サラセニア科"},
		{Id: 70080, Name: "サルトリイバラ科"},
		{Id: 70081, Name: "シキミ科"},
		{Id: 70082, Name: "シソ科"},
		{Id: 70083, Name: "シナノキ科"},
		{Id: 70084, Name: "シモンジア科"},
		{Id: 70085, Name: "シャクジョウソウ科"},
		{Id: 70086, Name: "ジャケツイバラ科"},
		{Id: 70087, Name: "シュウカイドウ科"},
		{Id: 70088, Name: "ショウガ科"},
		{Id: 70089, Name: "ショウブ科"},
		{Id: 70090, Name: "シレンゲ科"},
		{Id: 70091, Name: "ジンチョウゲ科"},
		{Id: 70092, Name: "スイカズラ科"},
		{Id: 70093, Name: "スイレン科"},
		{Id: 70094, Name: "スギ科"},
		{Id: 70095, Name: "スグリ科"},
		{Id: 70096, Name: "スズカケノキ科"},
		{Id: 70097, Name: "スズラン亜科"},
		{Id: 70098, Name: "スベリヒユ科"},
		{Id: 70099, Name: "スミレ科"},
		{Id: 70100, Name: "セリ科"},
		{Id: 70101, Name: "センダン科"},
		{Id: 70102, Name: "センリョウ科"},
		{Id: 70103, Name: "タコノキ科"},
		{Id: 70104, Name: "タシロイモ科"},
		{Id: 70105, Name: "タデ科"},
		{Id: 70106, Name: "タヌキアヤメ科"},
		{Id: 70107, Name: "タヌキモ科"},
		{Id: 70108, Name: "ツゲ科"},
		{Id: 70109, Name: "ツチトリモチ科"},
		{Id: 70110, Name: "ツツジ科"},
		{Id: 70111, Name: "ツヅラフジ科"},
		{Id: 70112, Name: "ツバキ科"},
		{Id: 70113, Name: "ツユクサ科"},
		{Id: 70114, Name: "ツリフネソウ科"},
		{Id: 70115, Name: "ツルボラン科"},
		{Id: 70116, Name: "ツルムラサキ科"},
		{Id: 70117, Name: "ディディエレア科"},
		{Id: 70118, Name: "テミス科"},
		{Id: 70119, Name: "トウダイグサ科"},
		{Id: 70120, Name: "ドクウツギ科"},
		{Id: 70121, Name: "ドクダミ科"},
		{Id: 70122, Name: "トケイソウ科"},
		{Id: 70123, Name: "トチカガミ科"},
		{Id: 70124, Name: "トチノキ科"},
		{Id: 70125, Name: "トベラ科"},
		{Id: 70126, Name: "ナス科"},
		{Id: 70127, Name: "ナデシコ科"},
		{Id: 70128, Name: "ナンヨウスギ科"},
		{Id: 70129, Name: "ニガキ科"},
		{Id: 70130, Name: "ニクズク科"},
		{Id: 70131, Name: "ニシキギ科"},
		{Id: 70132, Name: "ニレ科"},
		{Id: 70133, Name: "ヌマミズキ科"},
		{Id: 70134, Name: "ネギ科"},
		{Id: 70135, Name: "ネムノキ科"},
		{Id: 70136, Name: "ノウゼンカズラ科"},
		{Id: 70137, Name: "ノウゼンハレン科"},
		{Id: 70138, Name: "ノボタン科"},
		{Id: 70139, Name: "パイナップル科"},
		{Id: 70140, Name: "ハイノキ科"},
		{Id: 70141, Name: "ハエモドルム科"},
		{Id: 70142, Name: "ハゴロモモ科"},
		{Id: 70143, Name: "バショウ科"},
		{Id: 70144, Name: "ハスノハギリ科"},
		{Id: 70145, Name: "ハス科"},
		{Id: 70146, Name: "ハゼリソウ科"},
		{Id: 70147, Name: "ハナイ科"},
		{Id: 70148, Name: "ハナシノブ科"},
		{Id: 70149, Name: "パナマソウ科"},
		{Id: 70150, Name: "パパイア科"},
		{Id: 70151, Name: "ハマウツボ科"},
		{Id: 70152, Name: "ハマビシ科"},
		{Id: 70153, Name: "ハマミズナ科"},
		{Id: 70154, Name: "バラ科"},
		{Id: 70155, Name: "ハンニチバナ科"},
		{Id: 70156, Name: "パンヤ科"},
		{Id: 70157, Name: "バンレイシ科"},
		{Id: 70158, Name: "ヒガンバナ科"},
		{Id: 70159, Name: "ヒシ科"},
		{Id: 70160, Name: "ヒダテラ科"},
		{Id: 70161, Name: "ヒドロスタキス科"},
		{Id: 70162, Name: "ヒナノシャクジョウ科"},
		{Id: 70163, Name: "ヒノキ科"},
		{Id: 70164, Name: "ヒマンタンドラ科"},
		{Id: 70165, Name: "ヒメハギ科"},
		{Id: 70166, Name: "ビャクダン科"},
		{Id: 70167, Name: "ビャクブ科"},
		{Id: 70168, Name: "ヒユ科"},
		{Id: 70169, Name: "ヒルガオ科"},
		{Id: 70170, Name: "ヒルギ科"},
		{Id: 70171, Name: "ビワモドキ科"},
		{Id: 70172, Name: "フウチョウソウ科"},
		{Id: 70173, Name: "フウロソウ科"},
		{Id: 70174, Name: "フサザクラ科"},
		{Id: 70175, Name: "フジウツギ科"},
		{Id: 70176, Name: "ブドウ科"},
		{Id: 70177, Name: "フトモモ科"},
		{Id: 70178, Name: "ブナ科"},
		{Id: 70179, Name: "ベニノキ科"},
		{Id: 70180, Name: "ベンケイソウ科"},
		{Id: 70181, Name: "ホシクサ科"},
		{Id: 70182, Name: "ボタン科"},
		{Id: 70183, Name: "ホルトノキ科"},
		{Id: 70184, Name: "ボロボロノキ科"},
		{Id: 70185, Name: "マキ科"},
		{Id: 70186, Name: "マタタビ科"},
		{Id: 70187, Name: "マチン科"},
		{Id: 70188, Name: "マツブサ科"},
		{Id: 70189, Name: "マツムシソウ科"},
		{Id: 70190, Name: "マツモ科"},
		{Id: 70191, Name: "マツ科"},
		{Id: 70192, Name: "マメ科"},
		{Id: 70193, Name: "マンサク科"},
		{Id: 70194, Name: "ミカン科"},
		{Id: 70195, Name: "ミズアオイ科"},
		{Id: 70196, Name: "ミズキ科"},
		{Id: 70197, Name: "ミソハギ科"},
		{Id: 70198, Name: "ミツガシワ科"},
		{Id: 70199, Name: "ミツバウツギ科"},
		{Id: 70200, Name: "ムクロジ科"},
		{Id: 70201, Name: "ムラサキ科"},
		{Id: 70202, Name: "メギ科"},
		{Id: 70203, Name: "モウセンゴケ科"},
		{Id: 70204, Name: "モクセイ科"},
		{Id: 70205, Name: "モクレン科"},
		{Id: 70206, Name: "モチノキ科"},
		{Id: 70207, Name: "ヤマゴボウ科"},
		{Id: 70208, Name: "ヤマノイモ科"},
		{Id: 70209, Name: "ユキノシタ科"},
		{Id: 70210, Name: "ユズリハ科"},
		{Id: 70211, Name: "ユリズイセン科"},
		{Id: 70212, Name: "ユリ科"},
		{Id: 70213, Name: "ラン科"},
		{Id: 70214, Name: "リムナンテス科"},
		{Id: 70215, Name: "リュウゼツラン科"},
		{Id: 70216, Name: "リンドウ科"},
		{Id: 70217, Name: "ロウバイ科"},
		{Id: 70271, Name: "ミズアオイ属"},
		{Id: 80001, Name: "アウストロバイレヤ目"},
		{Id: 80002, Name: "アオイ目"},
		{Id: 80003, Name: "アカネ目"},
		{Id: 80004, Name: "アブラナ目"},
		{Id: 80005, Name: "アムボレラ目"},
		{Id: 80006, Name: "アリノトウグサ目"},
		{Id: 80007, Name: "アワゴケ目"},
		{Id: 80008, Name: "イグサ目"},
		{Id: 80009, Name: "イネ目"},
		{Id: 80010, Name: "イバラモ目"},
		{Id: 80011, Name: "イラクサ目"},
		{Id: 80012, Name: "ウツボカズラ目"},
		{Id: 80013, Name: "ウマノスズクサ目"},
		{Id: 80014, Name: "オオバコ目"},
		{Id: 80015, Name: "オモダカ目"},
		{Id: 80016, Name: "カキノキ目"},
		{Id: 80017, Name: "カネラ目"},
		{Id: 80018, Name: "カヤツリグサ目"},
		{Id: 80019, Name: "ガリア目"},
		{Id: 80020, Name: "キキョウ目"},
		{Id: 80021, Name: "キク目"},
		{Id: 80022, Name: "キントラノオ目"},
		{Id: 80023, Name: "キンポウゲ目"},
		{Id: 80024, Name: "キジカクシ目"},
		{Id: 80025, Name: "クスノキ目"},
		{Id: 80026, Name: "クロウメモドキ目"},
		{Id: 80027, Name: "クロッソソマ目"},
		{Id: 80028, Name: "グンネラ目"},
		{Id: 80029, Name: "ケシ目"},
		{Id: 80030, Name: "コショウ目"},
		{Id: 80031, Name: "ゴマノハグサ目"},
		{Id: 80032, Name: "サクラソウ目"},
		{Id: 80033, Name: "サトイモ目"},
		{Id: 80034, Name: "シキミ目"},
		{Id: 80035, Name: "シソ目"},
		{Id: 80036, Name: "ショウガ目"},
		{Id: 80037, Name: "ショウブ目"},
		{Id: 80038, Name: "スイレン目"},
		{Id: 80039, Name: "スミレ目"},
		{Id: 80040, Name: "セリ目"},
		{Id: 80041, Name: "タコノキ目"},
		{Id: 80042, Name: "ツツジ目"},
		{Id: 80043, Name: "ツバキ目"},
		{Id: 80044, Name: "ツユクサ目"},
		{Id: 80045, Name: "トウダイグサ目"},
		{Id: 80046, Name: "ナス目"},
		{Id: 80047, Name: "ナデシコ目"},
		{Id: 80048, Name: "ニシキギ目"},
		{Id: 80049, Name: "ハマビシ目"},
		{Id: 80050, Name: "バラ目"},
		{Id: 80051, Name: "ヒメハギ目"},
		{Id: 80052, Name: "ビャクダン目"},
		{Id: 80053, Name: "ビワモドキ目"},
		{Id: 80054, Name: "フウチョウソウ目"},
		{Id: 80055, Name: "フウロソウ目"},
		{Id: 80056, Name: "フトモモ目"},
		{Id: 80057, Name: "マツムシソウ目"},
		{Id: 80058, Name: "マメ目"},
		{Id: 80059, Name: "マンサク目"},
		{Id: 80060, Name: "ミズキ目"},
		{Id: 80061, Name: "ムクロジ目"},
		{Id: 80062, Name: "モクレン目"},
		{Id: 80063, Name: "モチノキ目"},
		{Id: 80064, Name: "ヤマグルマ目"},
		{Id: 80065, Name: "ヤマモガシ目"},
		{Id: 80066, Name: "ユキノシタ目"},
		{Id: 80067, Name: "ユリ目"},
		{Id: 80068, Name: "ラン目"},
		{Id: 80069, Name: "リンドウ目"},
		{Id: 90002, Name: "イチゴ科"},
		{Id: 90003, Name: "ウメ科"},
		{Id: 90004, Name: "キク亜綱"},
		{Id: 90005, Name: "グネツム綱"},
		{Id: 90006, Name: "紅藻"},
		{Id: 90008, Name: "米"},
		{Id: 90009, Name: "桜"},
		{Id: 90010, Name: "シダ類"},
		{Id: 90011, Name: "ショウガ亜綱"},
		{Id: 90012, Name: "スプリング・エフェメラル"},
		{Id: 90013, Name: "ハーブ"},
		{Id: 90014, Name: "バラ亜綱"},
		{Id: 90015, Name: "ビャクシン属"},
		{Id: 90016, Name: "ビワモドキ亜綱"},
		{Id: 90017, Name: "マンサク亜綱"},
		{Id: 90018, Name: "メロン"},
		{Id: 90019, Name: "モクレン亜綱"},
	}

	plants := []Plant{
		{Id: 3, Name: "もみじ", Detail: "紅葉狩りの「狩り」という言葉は「草花を眺めること」の意味をさし、平安時代には実際に紅葉した木の枝を手折り（狩り）、手のひらにのせて鑑賞する、という鑑賞方法があった。実際に枝を折り取って持ち帰る行為は森林窃盗罪となる。", Rarity: 0},
		{Id: 4, Name: "竹", Detail: "竹が草の一種か木の一種かは意見が分かれている。現在のところ、多年草の一種として扱う学説が多い。", Rarity: 0},
		{Id: 5, Name: "ドクダミ", Detail: "ドクダミの学名である Houttuynia cordata のうち、属名の Houttuynia はオランダの博物学者であるマールテン・ホッタイン (Maarten Houttuyn, 1720–1798) への献名であり、種小名の cordata はラテン語でハート形の葉の形を示している", Rarity: 0},
		{Id: 6, Name: "シロツメクサ", Detail: "濃厚な蜂蜜が得られる。また、若葉は食用になる。橋本郁三によると、塩茹でして葉柄が柔らかくなったら冷水で手早く冷まし、胡麻和え・辛子和え・甘酢などでいただくのが良い。花はフライ・てんぷらにする", Rarity: 0},
		{Id: 7, Name: "ヒガンバナ", Detail: "通常よく見られる赤色種のラジアータに加え、アルビノ種のように稀に色素形成異常で白みがかった個体もある", Rarity: 0},
		{Id: 8, Name: "タンポポ", Detail: "花のつくりは非常に進化していて、植物進化の系統では、頂点に立つグループの一つ。花は朝に開き、夕方に閉じる。雨が降らなければ、花は3日連続して、規則正しく開閉する。", Rarity: 0},
		{Id: 9, Name: "キンモクセイ", Detail: "樹皮が動物のサイ（犀）の足に似ていることから中国で「木犀」と名付けられ、ギンモクセイの白い花色に対して、橙黄色の花を金色に見立ててキンモクセイという名で呼ばれるようになった。", Rarity: 0},
		{Id: 10, Name: "イチョウ", Detail: "世界で最古の現生樹種の一つ。イチョウ類はペルム紀に出現し、大半が新生代に入ると絶滅したため、現在残っていたのはイチョウ類で唯一生き残っている種である。絶滅危惧種に指定されている。", Rarity: 0},
		{Id: 11, Name: "エノコログサ", Detail: "夏から秋にかけてつける花穂が、犬の尾に似ていることから、犬っころ草（いぬっころくさ）が転じてエノコログサという呼称になったとされ、漢字でも「狗（犬）の尾の草」と表記する。", Rarity: 0},
		{Id: 12, Name: "ジャスミン", Detail: "ほとんどの種が観賞用として栽培されている。栽培の歴史は古く、古代エジプトですでに行われていたといわれている。", Rarity: 0},
		{Id: 13, Name: "ゲッケイジュ", Detail: "庭木、公園樹としての利用のほか、ハーブとして、葉は香辛料（スパイス）として煮込み料理の香味づけに、葉や実は薬用として利用される。", Rarity: 0},
		{Id: 14, Name: "ラベンダー", Detail: "伝統的にハーブとして古代エジプト、ギリシャ、ローマ、アラビア、ヨーロッパなどで薬や調理に利用され、芳香植物としてその香りが活用されてきた。", Rarity: 0},
		{Id: 15, Name: "ミント", Detail: "変種が出来やすく600種を超えると言われるほど多種多様な種がある。", Rarity: 0},
		{Id: 16, Name: "ヨモギ", Detail: "ヨモギが持っている独特の香りは、害虫や雑菌から身を守るために抗菌化物質などの化学物質を発展させてきたものに由来する。多くの薬効があることからハーブの女王の異名がある。", Rarity: 0},
		{Id: 17, Name: "スイセン", Detail: "英語名はnarcissus。スイセンは日本の気候と相性が良いので、植え放しでも勝手に増える。", Rarity: 0},
		{Id: 18, Name: "ススキ", Detail: "枯れすすき（枯薄、花も穂も枯れたススキ）には枯れ尾花/枯尾花（かれおばな）という呼称（古名）もあり、現代でも「幽霊の正体見たり枯尾花」という諺はよく知られている。", Rarity: 0},
		{Id: 19, Name: "パンジー", Detail: "花が人間の顔に似て、8月には深く思索にふけるかのように前に傾くところからフランス語の「思想」を意味する単語パンセ（pensée）にちなんでパンジーと名づけられた。このその由来のために、パンジーは長い間自由思想のシンボルだった。", Rarity: 0},
		{Id: 20, Name: "プラタナス", Detail: "ニレ、ボダイジュ、マロニエとともに世界四大街路樹の一つに数えられる。樹皮が剥がれるとまるで迷彩柄のような見た目になる。", Rarity: 0},
		{Id: 21, Name: "梅", Detail: "日本では6月6日が「梅の日」とされている。天文14年4月17日（旧暦、1545年6月6日）、賀茂神社の例祭に梅が献上された故事に由来する。", Rarity: 0},
		{Id: 22, Name: "バラ", Detail: "近代以前、日本はバラの自生地として世界的に知られており、品種改良に使用された原種のうち3種類（ノイバラ、テリハノイバラ、ハマナス）は日本原産である。ちなみに、南半球にはバラは自生しない。", Rarity: 0},
		{Id: 23, Name: "チューリップ", Detail: "食用に適する一部の品種は、球根の糖度が極めて高くでん粉に富むため、オランダでは食用としての栽培も盛んで、主に製菓材料として用いられる。日本でもシロップ漬にした球根を使った和菓子やパイが富山県砺波市で販売されている。", Rarity: 0},
		{Id: 24, Name: "スミレ", Detail: "2021年（令和3年）4月13日発売の94円普通切手の意匠である。", Rarity: 0},
		{Id: 25, Name: "フジ", Detail: "種子散布に関しては、乾燥すると鞘が二つに裂開し、それぞれがよじれることで種子を飛ばすが、この際の種子の飛ぶ力は大変なもので、当たって怪我をした人が実在するという。", Rarity: 0},
	}

	plant_tags := map[int][]int{
		3:  {13, 10001, 20010, 20012, 50002, 50004, 60001, 60003, 70200, 80061},
		4:  {16, 10001, 10002, 20010, 20011, 40001, 50002, 60003, 70022, 80009},
		5:  {12, 10002, 20010, 50002, 50006, 60003, 60006, 80030, 70121},
		6:  {11, 12, 10002, 10003, 20010, 20012, 50004, 50006, 60003, 60006, 70266, 70192},
		7:  {13, 10002, 10003, 20010, 20011, 50001, 50004, 60001, 80024, 70158},
		8:  {11, 10002, 10003, 20005, 20010, 20012, 50004, 50006, 70052},
		9:  {13, 10001, 10003, 20005, 20010, 20012, 50002, 50004, 80035, 70204},
		10: {13, 10001, 20005, 20009, 50002, 50004, 50005, 60004},
		11: {12, 13, 10002, 20005, 20010, 20011, 50006, 60003, 70022, 80009},
		12: {11, 10002, 10003, 20005, 20010, 20012, 50002, 50004, 50006, 60006, 80035, 70204},
		13: {11, 10002, 20005, 20010, 20012, 50002, 50006, 60003, 80025, 70061},
		14: {12, 10002, 10003, 20005, 20010, 20012, 50002, 50004, 50006, 60005, 80035, 70082},
		15: {11, 12, 10002, 20005, 20010, 20012, 50002, 50006, 60003, 80035, 70082},
		16: {11, 10002, 20005, 20010, 20012, 50002, 50006, 60003, 70052, 80021},
		17: {11, 13, 14, 10003, 20005, 20010, 20011, 50001, 50004, 50006, 60004, 60006, 80024, 70158},
		18: {13, 10002, 20005, 20010, 20011, 50005, 60004, 80009, 70022},
		19: {11, 10003, 20005, 20010, 20012, 50001, 50002, 50003, 50006, 60002, 60004, 60005, 60006, 70057, 70099},
		20: {13, 16, 10001, 20005, 20010, 20012, 50005, 60003, 80065, 70096},
		21: {11, 10001, 10003, 20005, 20010, 20012, 50002, 50004, 60001, 60006, 80050, 70154, 90003},
		22: {11, 10003, 20005, 20010, 20012, 50002, 50004, 60001, 60002, 60003, 60004, 60005, 60006, 60007, 80050, 70154},
		23: {11, 10003, 20010, 20011, 50002, 50001, 50004, 50006, 60001, 60002, 60003, 60004, 60005, 60006, 70212, 80067},
		24: {11, 10002, 10003, 20005, 20010, 20012, 50001, 50002, 50004, 50006, 60001, 60002, 60004, 60005, 60006, 80022, 70099},
		25: {11, 10003, 20005, 20010, 20012, 50001, 50002, 50004, 60005, 80058, 70192},
	}

	uploads := map[int][]UploadPost{
		3: {
			{Id: 1, Latitude: 35.03206328632516, Longitude: 135.7859364806136, Hash: "autumn_leaf", UploadTime: time.Date(2021, 10, 1, 0, 0, 0, 0, time.Local)},
			{Id: 2, Latitude: 35.03396253383616, Longitude: 135.8005059253635, Hash: "arashiyama_autumn_leaf", UploadTime: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local)},
		},
		4: {
			{Id: 1, Latitude: 35.02980558218886, Longitude: 135.79593438200007, Hash: "bamboo", UploadTime: time.Date(2021, 10, 3, 0, 0, 0, 0, time.Local)},
		},
		5: {
			{Id: 1, Latitude: 35.031406520084325, Longitude: 135.78922952850533, Hash: "chameleon_plant", UploadTime: time.Date(2021, 10, 5, 0, 0, 0, 0, time.Local)},
		},
		6: {
			{Id: 1, Latitude: 35.03107306953346, Longitude: 135.79120499575424, Hash: "clover", UploadTime: time.Date(2021, 10, 8, 0, 0, 0, 0, time.Local)},
		},
		7: {
			{Id: 1, Latitude: 35.033123637980054, Longitude: 135.79609896518846, Hash: "cluster_amaryllis", UploadTime: time.Date(2021, 10, 10, 0, 0, 0, 0, time.Local)},
		},
		8: {
			{Id: 1, Latitude: 35.02956321614069, Longitude: 135.7910732767183, Hash: "dandelion", UploadTime: time.Date(2021, 10, 12, 0, 0, 0, 0, time.Local)},
		},
		9: {
			{Id: 1, Latitude: 35.02829345500896, Longitude: 135.7906561557475, Hash: "fragrant_olive", UploadTime: time.Date(2021, 10, 15, 0, 0, 0, 0, time.Local)},
		},
		10: {
			{Id: 1, Latitude: 35.02893246681226, Longitude: 135.79145921857554, Hash: "ginkgo", UploadTime: time.Date(2021, 10, 18, 0, 0, 0, 0, time.Local)},
		},
		11: {
			{Id: 1, Latitude: 35.029634180926735, Longitude: 135.7926468733469, Hash: "green_foxtail", UploadTime: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local)},
		},
		12: {
			{Id: 1, Latitude: 35.042769875750096, Longitude: 135.7874045716391, Hash: "jasmine", UploadTime: time.Date(2021, 10, 22, 0, 0, 0, 0, time.Local)},
		},
		13: {
			{Id: 1, Latitude: 35.044763041413454, Longitude: 135.7858223756318, Hash: "laurel", UploadTime: time.Date(2021, 10, 25, 0, 0, 0, 0, time.Local)},
		},
		14: {
			{Id: 1, Latitude: 35.04435991838187, Longitude: 135.789197896487, Hash: "lavender", UploadTime: time.Date(2021, 10, 27, 0, 0, 0, 0, time.Local)},
		},
		15: {
			{Id: 1, Latitude: 35.04515925461211, Longitude: 135.78726916545824, Hash: "mint", UploadTime: time.Date(2021, 10, 29, 0, 0, 0, 0, time.Local)},
		},
		16: {
			{Id: 1, Latitude: 35.03186070766323, Longitude: 135.78607848837888, Hash: "mugwort", UploadTime: time.Date(2021, 11, 1, 0, 0, 0, 0, time.Local)},
		},
		17: {
			{Id: 1, Latitude: 35.03613639322238, Longitude: 135.7886611787627, Hash: "narcissus", UploadTime: time.Date(2021, 11, 3, 0, 0, 0, 0, time.Local)},
		},
		18: {
			{Id: 1, Latitude: 35.03170063903726, Longitude: 135.77273809003546, Hash: "pampas_grass", UploadTime: time.Date(2021, 11, 5, 0, 0, 0, 0, time.Local)},
		},
		19: {
			{Id: 1, Latitude: 35.03528714118098, Longitude: 135.79254644806633, Hash: "pansy", UploadTime: time.Date(2021, 11, 7, 0, 0, 0, 0, time.Local)},
		},
		20: {
			{Id: 1, Latitude: 35.04674347083965, Longitude: 135.788973382866, Hash: "plane_tree", UploadTime: time.Date(2021, 11, 9, 0, 0, 0, 0, time.Local)},
		},
		21: {
			{Id: 1, Latitude: 35.042255154464804, Longitude: 135.79576852585632, Hash: "rose_arch", UploadTime: time.Date(2021, 11, 11, 0, 0, 0, 0, time.Local)},
		},
		22: {
			{Id: 1, Latitude: 35.036607297147526, Longitude: 135.78466777707337, Hash: "rose", UploadTime: time.Date(2021, 11, 13, 0, 0, 0, 0, time.Local)},
			{Id: 2, Latitude: 35.02800570100121, Longitude: 135.79274522090034, Hash: "rose", UploadTime: time.Date(2021, 11, 15, 0, 0, 0, 0, time.Local)},
		},
		23: {
			{Id: 1, Latitude: 35.02414961144393, Longitude: 135.79038449356278, Hash: "tulip", UploadTime: time.Date(2021, 11, 17, 0, 0, 0, 0, time.Local)},
		},
		24: {
			{Id: 1, Latitude: 35.034854902690974, Longitude: 135.78150180377762, Hash: "violet", UploadTime: time.Date(2021, 11, 19, 0, 0, 0, 0, time.Local)},
		},
		25: {
			{Id: 1, Latitude: 35.02238181435338, Longitude: 135.76647226352986, Hash: "wisteria", UploadTime: time.Date(2021, 11, 21, 0, 0, 0, 0, time.Local)},
		},
	}

	bulk := client.BulkWriter(ctx)
	for _, plant := range plants {
		ref := client.Collection("plant").Doc(strconv.Itoa(plant.Id))
		ups, ok := uploads[plant.Id]
		if ok && len(ups) > 0 {
			plant.Hash = ups[len(ups)-1].Hash
		}
		_, err := bulk.Set(ref, plant)
		if err != nil {
			logger.Error("failed bulk.Set plant: %v", err)
		}
	}
	for _, tag := range tags {
		ref := client.Collection("tag").Doc(strconv.Itoa(tag.Id))
		_, err := bulk.Set(ref, tag)
		if err != nil {
			logger.Error("failed bulk.Set tag: %v", err)
		}
	}
	for plant_id, tag_ids := range plant_tags {
		for _, tag_id := range tag_ids {
			ref := client.Collection("plant").Doc(strconv.Itoa(plant_id)).Collection("tags").Doc(strconv.Itoa(tag_id))
			_, err := bulk.Set(ref, PlantTag{Id: tag_id})
			if err != nil {
				logger.Error("failed bulk.Set plant_tag: %v", err)
			}
		}
	}
	for plant_id, uploads := range uploads {
		for _, upload := range uploads {
			ref := client.Collection("plant").Doc(strconv.Itoa(plant_id)).Collection("uploads").Doc(strconv.Itoa(upload.Id))
			_, err := bulk.Set(ref, upload)
			if err != nil {
				logger.Error("failed bulk.Set upload: %v", err)
			}
		}
	}
	ref := client.Collection("init_finished").Doc("hoge")
	bulk.Set(ref, map[string]bool{"init_finished": true})

	bulk.End()
}