package main

const testinput06 = "3,4,3,1,2"

// func TestNewLanternfish(t *testing.T) {
// 	want := []Lanternfish{
// 		{3}, {4}, {3}, {1}, {2},
// 	}

// 	got := NewLanternfish(testinput06)

// 	for i, f := range got {
// 		if want[i] != *f {
// 			t.Errorf("want %+v, got %+v", want[i], *f)
// 		}
// 	}
// }

// func TestLiveDays(t *testing.T) {
// 	var tests = []struct {
// 		days int
// 		want []Lanternfish
// 	}{
// 		{1, []Lanternfish{{2}, {3}, {2}, {0}, {1}}},
// 		{2, []Lanternfish{{1}, {2}, {1}, {6}, {0}, {8}}},
// 		{3, []Lanternfish{{0}, {1}, {0}, {5}, {6}, {7}, {8}}},
// 		{4, []Lanternfish{{6}, {0}, {6}, {4}, {5}, {6}, {7}, {8}, {8}}},
// 		{5, []Lanternfish{{5}, {6}, {5}, {3}, {4}, {5}, {6}, {7}, {7}, {8}}},
// 		{6, []Lanternfish{{4}, {5}, {4}, {2}, {3}, {4}, {5}, {6}, {6}, {7}}},
// 		{7, []Lanternfish{{3}, {4}, {3}, {1}, {2}, {3}, {4}, {5}, {5}, {6}}},
// 		{8, []Lanternfish{{2}, {3}, {2}, {0}, {1}, {2}, {3}, {4}, {4}, {5}}},
// 		{9, []Lanternfish{{1}, {2}, {1}, {6}, {0}, {1}, {2}, {3}, {3}, {4}, {8}}},
// 		{10, []Lanternfish{{0}, {1}, {0}, {5}, {6}, {0}, {1}, {2}, {2}, {3}, {7}, {8}}},
// 		{11, []Lanternfish{{6}, {0}, {6}, {4}, {5}, {6}, {0}, {1}, {1}, {2}, {6}, {7}, {8}, {8}, {8}}},
// 		{12, []Lanternfish{{5}, {6}, {5}, {3}, {4}, {5}, {6}, {0}, {0}, {1}, {5}, {6}, {7}, {7}, {7}, {8}, {8}}},
// 		{13, []Lanternfish{{4}, {5}, {4}, {2}, {3}, {4}, {5}, {6}, {6}, {0}, {4}, {5}, {6}, {6}, {6}, {7}, {7}, {8}, {8}}},
// 		{14, []Lanternfish{{3}, {4}, {3}, {1}, {2}, {3}, {4}, {5}, {5}, {6}, {3}, {4}, {5}, {5}, {5}, {6}, {6}, {7}, {7}, {8}}},
// 		{15, []Lanternfish{{2}, {3}, {2}, {0}, {1}, {2}, {3}, {4}, {4}, {5}, {2}, {3}, {4}, {4}, {4}, {5}, {5}, {6}, {6}, {7}}},
// 		{16, []Lanternfish{{1}, {2}, {1}, {6}, {0}, {1}, {2}, {3}, {3}, {4}, {1}, {2}, {3}, {3}, {3}, {4}, {4}, {5}, {5}, {6}, {8}}},
// 		{17, []Lanternfish{{0}, {1}, {0}, {5}, {6}, {0}, {1}, {2}, {2}, {3}, {0}, {1}, {2}, {2}, {2}, {3}, {3}, {4}, {4}, {5}, {7}, {8}}},
// 		{18, []Lanternfish{{6}, {0}, {6}, {4}, {5}, {6}, {0}, {1}, {1}, {2}, {6}, {0}, {1}, {1}, {1}, {2}, {2}, {3}, {3}, {4}, {6}, {7}, {8}, {8}, {8}, {8}}},
// 	}

// 	for _, tt := range tests {
// 		t.Run("", func(t *testing.T) {
// 			fish := NewLanternfish(testinput06)
// 			for d := 1; d <= tt.days; d++ {
// 				for _, f := range fish {
// 					maybeFish := f.LiveADay()
// 					if maybeFish != nil {
// 						fish = append(fish, maybeFish)
// 					}
// 				}
// 			}
// 			if len(fish) != len(tt.want) {
// 				t.Errorf("got %d, want %d", len(fish), len(tt.want))
// 			}
// 		})
// 	}
// }
