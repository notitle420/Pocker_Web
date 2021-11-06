package check_hand

import (
	"fmt"
	"sort"
)

type Card struct {
	Mark   int
	Number int
}

var All_cards = []Card{}

//トランプを初期化、ダイヤ:1 ハート:2 スペード:3 クローバー:4
func All_card_init() {
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 13; j++ {
			tmp_card := Card{Mark: i, Number: j}
			All_cards = append(All_cards, tmp_card)
		}
	}
}

//役があるかを調べる関数
func Check_hand(hand1 Card, hand2 Card, frop1 Card, frop2 Card, frop3 Card, turn Card, river Card) {
	hand := []Card{hand1, hand2, frop1, frop2, frop3, turn, river}
	hand_numbers := []int{hand1.Number, hand2.Number, frop1.Number, frop2.Number, frop3.Number, turn.Number, river.Number}
	uniq_num := uniq_num(hand_numbers)
	same_num := same_num_count(hand_numbers, uniq_num)
	fmt.Println("hand", hand)
	for _, v := range hand {
		switch v.Mark {
		case 1:
			fmt.Print("♦︎")
		case 2:
			fmt.Print("❤︎")
		case 3:
			fmt.Print("♠︎")
		case 4:
			fmt.Print("♣︎")
		}
		fmt.Print(v.Number, ",")
	}
	fmt.Println("")
	fmt.Println("uniq_num", uniq_num)
	fmt.Println("same_num", same_num)
	//ソートするためのLISTを作る
	same_num_count_list := List{}
	for k, v := range same_num {
		e := CountSameNum{k, v}
		same_num_count_list = append(same_num_count_list, e)
	}
	//重複した数リストをソート
	sort.Sort(same_num_count_list)
	fmt.Println("countsort", same_num_count_list)

	//手札をソート
	sorted_hand := SortedHand{}
	sorted_hand = hand
	sort.Sort(sorted_hand)

	result_hand := []Card{}
	result := ""
	tmp_hand := []Card{}
	tmp2_hand := []Card{}

	straight_count := 0
	straight_end_num := 0
	straight_numbers := []int{}
	ace_straigt := []int{13, 12, 11, 10, 1}
	ace_straigt_count := 0

	//フラッシュチェック
	flush_check := []int{0, 0, 0, 0}
	for i := 1; i <= 4; i++ {
		for _, v := range hand {
			if i == v.Mark {
				flush_check[i-1] += 1
			}
		}
	}

	//flushのマークを確認する
	flush_Mark := 0
	for i, v := range flush_check {
		if v >= 5 {
			flush_Mark = i + 1
		}
	}

	if flush_Mark >= 1 {
		straight_count = 0
		ace_straigt_count = 0
		straight_end_num = 0
		straight_numbers = []int{}
		result = "flush"
		result_hand = []Card{}
		tmp_hand = []Card{}
		fmt.Println("flush_Mark", flush_Mark)
		fmt.Println("sorted hand", sorted_hand)
		//フラッシュマークを全部いったん手札にいれる
		for _, v := range sorted_hand {
			if flush_Mark == v.Mark {
				tmp_hand = append(tmp_hand, v)
			}
		}

		//ストふらチェック
		for i := 0; i < len(tmp_hand)-1; i++ {
			if tmp_hand[i].Number-1 == tmp_hand[i+1].Number {
				straight_count += 1
			} else {
				straight_count = 0
			}
			if straight_count >= 4 {
				straight_end_num = tmp_hand[i+1].Number
				break
			}
		}

		//ロイヤルストレートフラッシュチェック
		for _, v := range ace_straigt {
			for _, z := range tmp_hand {
				if v == z.Number {
					ace_straigt_count += 1
				}
			}
		}

		//ストレートの数を作る
		if ace_straigt_count >= 4 {
			straight_numbers = ace_straigt
			result = "rolyal_straight_flush"
		}
		if straight_count >= 4 {
			for i := straight_end_num + 4; i >= straight_end_num; i-- {
				straight_numbers = append(straight_numbers, i)
			}
			result = "straight_flush"
		}

		if ace_straigt_count < 4 && straight_count < 4 {
			for _, v := range tmp_hand {
				if v.Number == 1 {
					result_hand = append(result_hand, v)
				}
			}
			for _, v := range tmp_hand {
				result_hand = append(result_hand, v)
				if len(result_hand) > 4 {
					break
				}
			}
		}

		//ストレートかロイヤルストレート
		if straight_count >= 4 || ace_straigt_count >= 4 {
			tmp2_hand = []Card{}
			result_hand = []Card{}
			for _, v := range straight_numbers {
				tmp_z := Card{0, 0}
				for _, z := range tmp_hand {
					if v == z.Number {
						//66 5 4 3 2 とかの対策　前のループの数字じゃない時
						if z.Number != tmp_z.Number {
							tmp2_hand = append(tmp2_hand, z)
						}
						tmp_z = z
					}
				}
			}
			//1をがあったら一番左にする
			for _, v := range tmp2_hand {
				if v.Number == 1 {
					result_hand = append(result_hand, v)
				}
			}
			for _, v := range tmp2_hand {
				result_hand = append(result_hand, v)
				if len(result_hand) > 4 {
					break
				}
			}
		}
	}

	if flush_Mark <= 0 {
		//ストレートチェック
		//重複した数字を抜いたリストが５個以上の時
		if len(uniq_num) >= 5 {
			fmt.Println((uniq_num))
			straight_count = 0
			//数字が連続してるか確認
			for i := 0; i < len(uniq_num)-1; i++ {
				if uniq_num[i]-1 == uniq_num[i+1] {
					straight_count += 1
				} else {
					straight_count = 0
				}
				if straight_count >= 4 {
					straight_end_num = uniq_num[i+1]
					break
				}
			}

			//1 13 12 11 10 のストレートをチェック
			for _, v := range ace_straigt {
				for _, z := range uniq_num {
					if v == z {
						ace_straigt_count += 1
					}
				}
			}
			if ace_straigt_count >= 4 {
				straight_numbers = ace_straigt
			} else {
				for i := straight_end_num + 4; i >= straight_end_num; i-- {
					straight_numbers = append(straight_numbers, i)
				}
			}
			fmt.Println(straight_numbers)
		}

		if straight_count >= 4 || ace_straigt_count >= 4 {
			tmp_hand := []Card{}
			result_hand = []Card{}
			for _, v := range straight_numbers {
				tmp_z := Card{0, 0}
				for _, z := range sorted_hand {
					if v == z.Number {
						//66 5 4 3 2 とかの対策　前のループの数字じゃない時
						if z.Number != tmp_z.Number {
							tmp_hand = append(tmp_hand, z)
						}
						tmp_z = z
					}
				}
			}
			//1をがあったら一番左にする
			for _, v := range tmp_hand {
				if v.Number == 1 {
					result_hand = append(result_hand, v)
				}
			}
			for _, v := range tmp_hand {
				result_hand = append(result_hand, v)
				if len(result_hand) > 4 {
					break
				}
			}
			result = "straight"
		}
	}

	if flush_Mark == 0 && straight_count < 4 && ace_straigt_count < 4 {
		//同じ数何枚あるかで
		switch same_num_count_list[0].count {
		case 1: //ペアなし
			result_hand = []Card{}
			fmt.Println("ペアなし")
			//そーとした後にエースがある場所を確認
			ace_place_hand := check_ace_hand(sorted_hand)
			//エースがあった場合、エースを一番最初に入れてその後に高い順でハンドに追加
			if ace_place_hand != 10 {
				result_hand = append(result_hand, sorted_hand[ace_place_hand])
				for i := 1; i <= 4; i++ {
					result_hand = append(result_hand, sorted_hand[i])
				}
			} else {
				result_hand = sorted_hand[0:4]
			}
			result = "None"
		case 2: //１ペアor２ペア
			//カウントリストの０個目だけじゃなくて１個目も２個なら2ペア
			if same_num_count_list[1].count == 2 {
				result_hand = []Card{}
				result = "2pair"
				ace_pair_place := -1
				//配列にいれるためのペアの数字
				pair_number1 := 0
				pair_number2 := 0
				five_number := 0
				//エースのペアがあるか確認
				for i, v := range same_num_count_list {
					if v.same_num == 1 && v.count == 2 {
						ace_pair_place = i
					}
				}
				//エースのペアがあるか確認
				if ace_pair_place >= 0 {
					pair_number1 = same_num_count_list[ace_pair_place].same_num
					pair_number2 = same_num_count_list[0].same_num
					//ペアが3枚ある時
					if same_num_count_list[2].count == 2 {
						//ペアが3つあったらソートされてるのでエースが3つ目 3つ目のペアとペアじゃないやつでどっちのほうがでかいか、確認する
						if same_num_count_list[1].same_num > same_num_count_list[3].same_num {
							five_number = same_num_count_list[1].same_num
						} else {
							five_number = same_num_count_list[3].same_num
						}
					} else { //ペアが2枚ならそのままソートされてる数字が5枚目
						five_number = same_num_count_list[2].same_num
					}
				} else { //エースのペアがない場合
					pair_number1 = same_num_count_list[0].same_num
					pair_number2 = same_num_count_list[1].same_num
					//重複削除リストでエースの場所取得
					ace_place_hand := check_ace_SameNum(same_num_count_list)
					if ace_place_hand != 10 {
						five_number = 1 //エースがあったら5枚目のハンドはエース確定
					} else { //エースなかったら通常処理
						//ペアが3枚あるかチェック
						if same_num_count_list[2].count == 2 {
							//ペアが3つあったら3つ目と、1枚の数字で比較
							if same_num_count_list[2].same_num > same_num_count_list[3].same_num {
								five_number = same_num_count_list[2].same_num
							} else {
								five_number = same_num_count_list[3].same_num
							}
						} else { //ペアが2枚ならそのままソートされてる数字が5枚目
							five_number = same_num_count_list[2].same_num
						}
					}
				}

				//上記で取得した数をハンドの中からチョイス
				for _, v := range hand {
					if pair_number1 == v.Number {
						result_hand = append(result_hand, v)
					}
				}
				for _, v := range hand {
					if pair_number2 == v.Number {
						result_hand = append(result_hand, v)
					}
				}
				for _, v := range hand {
					if five_number == v.Number {
						result_hand = append(result_hand, v)
					}
					if len(result_hand) == 5 { //手札5枚になったらブレーク
						break
					}
				}

				//１ペア
			} else {
				result_hand = []Card{}
				result = "ipair"
				//ペアの数字を取得
				pair_number := same_num_count_list[0].same_num
				//ペア、高い順の格納するための配列
				tmp_hand = []Card{}
				//ペアの数字を入れる
				for _, v := range sorted_hand {
					if pair_number == v.Number {
						tmp_hand = append(tmp_hand, v)
					}
				}
				//ペアじゃない数を入れる。　ペア、高い順の配列がtmp_handになる
				for _, v := range sorted_hand {
					if pair_number != v.Number {
						tmp_hand = append(tmp_hand, v)
					}
				}
				//結果ハンドにペアを格納
				for i := 0; i < 2; i++ {
					result_hand = append(result_hand, tmp_hand[i])
				}
				//エースがある場所を確認
				ace_place_hand := check_ace_hand(tmp_hand)
				//ペアがエースじゃなく、手札にエースがあった場合、エースを三番目にいれる
				if pair_number != 1 {
					if ace_place_hand != 10 {
						result_hand = append(result_hand, tmp_hand[ace_place_hand])
						//4枚目と5枚目はソートされた値
						for i := 2; i < 4; i++ {
							result_hand = append(result_hand, tmp_hand[i])
						}
					} else {
						//エースない場合、そのままソートされたものを追加
						for i := 2; i < 5; i++ {
							result_hand = append(result_hand, tmp_hand[i])
						}
					}
				} else {
					//ペアがエースだったら、そのままペア以外を追加する
					for i := 2; i < 5; i++ {
						result_hand = append(result_hand, tmp_hand[i])
					}
				}
			}

		case 3: //3カード or フルハウス
			//3カード
			if same_num_count_list[1].count == 1 {
				result_hand = []Card{}
				result = "3card"
				three_card_number := same_num_count_list[0].same_num
				four_number := 0
				five_number := 0
				if check_ace_SameNum(same_num_count_list) != 10 {
					//エースがあったら4枚目はエース確定
					four_number = 1
					five_number = same_num_count_list[1].same_num
				} else {
					four_number = same_num_count_list[1].same_num
					five_number = same_num_count_list[2].same_num
				}
				for _, v := range hand {
					if three_card_number == v.Number {
						result_hand = append(result_hand, v)
					}
				}
				for _, v := range hand {
					if four_number == v.Number {
						result_hand = append(result_hand, v)
					}
				}
				for _, v := range hand {
					if five_number == v.Number {
						result_hand = append(result_hand, v)
					}
				}
				//フルハウス
			} else if same_num_count_list[1].count >= 2 {
				result_hand = []Card{}
				result = "fullhouse"
				three_card_number := 0
				two_card_number := 0
				//エースチェック
				if check_ace_SameNum(same_num_count_list) != 10 {
					//エースが一枚だったら　関係ないのでそのまま使う
					if same_num_count_list[check_ace_SameNum(same_num_count_list)].count == 1 {
						three_card_number = same_num_count_list[0].same_num
						two_card_number = same_num_count_list[1].same_num
					}
					//エースが2枚だったら
					if same_num_count_list[check_ace_SameNum(same_num_count_list)].count == 2 {
						three_card_number = same_num_count_list[0].same_num
						two_card_number = 1
					} else if same_num_count_list[check_ace_SameNum(same_num_count_list)].count == 3 { //エースが3枚だったら
						if (same_num_count_list[0].same_num) == 1 {
							//エース3枚　違う数2枚の時
							three_card_number = 1
							two_card_number = same_num_count_list[1].same_num
						} else {
							//違う数3枚　エース3枚の時
							three_card_number = 1
							two_card_number = same_num_count_list[0].same_num
						}
					}
					//エースがなければ
				} else {
					three_card_number = same_num_count_list[0].same_num
					two_card_number = same_num_count_list[1].same_num
				}

				for _, v := range hand {
					if three_card_number == v.Number {
						result_hand = append(result_hand, v)
					}
				}
				for _, v := range hand {
					if two_card_number == v.Number {
						result_hand = append(result_hand, v)
						if len(result_hand) == 5 {
							break
						}
					}
				}
			}
		case 4: //4カード
			result = "4card"
			four_card_number := same_num_count_list[0].same_num
			five_number := same_num_count_list[1].same_num
			//エースがあったら5枚目はエースにか着替え
			if check_ace_SameNum(same_num_count_list) != 10 {
				five_number = same_num_count_list[1].same_num
			}
			for _, v := range hand {
				if four_card_number == v.Number {
					result_hand = append(result_hand, v)
				}
			}
			for _, v := range hand {
				if five_number == v.Number {
					result_hand = append(result_hand, v)
					if len(result_hand) == 5 {
						break
					}
				}
			}

		}
	}
	//ダイヤ:1 ハート:2 スペード:3 クローバー:4
	for _, v := range result_hand {
		switch v.Mark {
		case 1:
			fmt.Print("♦︎")
		case 2:
			fmt.Print("❤︎")
		case 3:
			fmt.Print("♠︎")
		case 4:
			fmt.Print("♣︎")
		}
		fmt.Print(v.Number, ",")
	}
	fmt.Println("")
	fmt.Println(result)
}

//手札を大きい順番にソートする〜〜ここから

type SortedHand []Card

func (l SortedHand) Len() int {
	return len(l)
}

func (l SortedHand) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l SortedHand) Less(i, j int) bool {
	return (l[i].Number > l[j].Number)
}

//〜〜ここまで　 数：重複数　= [5:1 7:1 8:1 9:1 11:2 13:1] →　[{11 2} {13 1} {9 1} {8 1} {7 1} {5 1}]

//重複数順にソートする〜〜ここから
type CountSameNum struct {
	same_num int
	count    int
}

type List []CountSameNum

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	if l[i].count == l[j].count {
		return (l[i].same_num > l[j].same_num)
	} else {
		return (l[i].count > l[j].count)
	}
}

//〜〜ここまで　 数：重複数　= [5:1 7:1 8:1 9:1 11:2 13:1] →　[{11 2} {13 1} {9 1} {8 1} {7 1} {5 1}]

//ユニークな数字を調べる関数
func uniq_num(slice []int) []int {
	m := make(map[int]bool)
	uniqSlice := make([]int, 0)

	for _, v := range slice {
		if _, ok := m[v]; !ok {
			m[v] = true
			uniqSlice = append(uniqSlice, v)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(uniqSlice)))
	return uniqSlice
}

//同じ数字をカウントする関数
func same_num_count(hand []int, uniq_num []int) map[int]int {
	result := make(map[int]int)
	for _, v := range uniq_num {
		tmp_count := 0
		for _, m := range hand {
			if v == m {
				tmp_count += 1
			}
			//数字をキーにして、重複数をバリューで保存
			result[v] = tmp_count
		}
	}
	return result
}

//重複削除したリストでエースの位置確認する関数
func check_ace_SameNum(hand []CountSameNum) int {
	for k, v := range hand {
		if v.same_num == 1 {
			return k
		}
	}
	return 10
}

//ハンドでエースの位置を確認する関数
func check_ace_hand(hand []Card) int {
	for k, v := range hand {
		if v.Number == 1 {
			return k
		}
	}
	return 10
}
