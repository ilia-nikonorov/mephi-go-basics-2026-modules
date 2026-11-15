package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		wantLen  int
		checkPos bool
	}{
		{
			name:    "нулевой размер",
			size:    0,
			wantLen: 0,
		},
		{
			name:    "отрицательный размер",
			size:    -10,
			wantLen: 0,
		},
		{
			name:     "один элемент",
			size:     1,
			wantLen:  1,
			checkPos: true,
		},
		{
			name:     "небольшой срез",
			size:     100,
			wantLen:  100,
			checkPos: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if len(got) != tt.wantLen {
				t.Fatalf("ожидалась длина %d, получена %d", tt.wantLen, len(got))
			}

			if tt.checkPos {
				for i, val := range got {
					if val <= 0 {
						t.Fatalf("ожидалось положительное число на индексе %d, получено %d", i, val)
					}
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "nil срез",
			data: nil,
			want: 0,
		},
		{
			name: "пустой срез",
			data: []int{},
			want: 0,
		},
		{
			name: "один элемент",
			data: []int{42},
			want: 42,
		},
		{
			name: "один отрицательный элемент",
			data: []int{-7},
			want: -7,
		},
		{
			name: "одинаковые элементы",
			data: []int{5, 5, 5, 5},
			want: 5,
		},
		{
			name: "максимум в начале",
			data: []int{100, 20, 30, 5},
			want: 100,
		},
		{
			name: "максимум в середине",
			data: []int{10, 250, 30, 5},
			want: 250,
		},
		{
			name: "максимум в конце",
			data: []int{10, 20, 30, 500},
			want: 500,
		},
		{
			name: "отрицательные числа",
			data: []int{-10, -3, -50, -7},
			want: -3,
		},
		{
			name: "смешанные числа",
			data: []int{-10, 0, 15, -2, 8},
			want: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum() = %d, ожидалось %d", got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "nil срез",
			data: nil,
			want: 0,
		},
		{
			name: "пустой срез",
			data: []int{},
			want: 0,
		},
		{
			name: "один элемент",
			data: []int{42},
			want: 42,
		},
		{
			name: "меньше элементов чем потоков",
			data: []int{3, 9, 1, 7},
			want: 9,
		},
		{
			name: "ровно CHUNKS элементов",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8},
			want: 8,
		},
		{
			name: "не кратно CHUNKS и максимум в хвосте",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 99},
			want: 99,
		},
		{
			name: "отрицательные числа",
			data: []int{-20, -15, -30, -5, -40, -12, -8, -100, -3, -50},
			want: -3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks() = %d, ожидалось %d", got, tt.want)
			}
		})
	}

	t.Run("большой случайный срез", func(t *testing.T) {
		data := generateRandomElements(100_003)
		want := maximum(data)
		got := maxChunks(data)
		if got != want {
			t.Errorf("maxChunks() = %d, maximum() = %d", got, want)
		}
	})
}
