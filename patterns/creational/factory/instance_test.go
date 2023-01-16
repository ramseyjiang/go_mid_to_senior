package factory

import (
	"reflect"
	"testing"
)

func TestEntry(t *testing.T) {
	var tests []struct {
		name string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Entry()
		})
	}
}

func TestGetFootWear(t *testing.T) {
	type args struct {
		category string
		size     int64
		gender   string
		price    float32
	}
	var tests []struct {
		name    string
		args    args
		want    iFootwear
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFootWear(tt.args.category, tt.args.size, tt.args.gender, tt.args.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFootWear() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFootWear() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewShoe(t *testing.T) {
	type args struct {
		size     int64
		category string
		price    float32
		gender   string
		discount float32
	}
	var tests []struct {
		name string
		args args
		want iFootwear
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newShoe(tt.args.size, tt.args.category, tt.args.price, tt.args.gender, tt.args.discount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newShoe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShoeGetCategory(t *testing.T) {
	type fields struct {
		size     int64
		price    float32
		category string
		gender   string
	}
	var tests []struct {
		name   string
		fields fields
		want   string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &shoe{
				size:     tt.fields.size,
				price:    tt.fields.price,
				category: tt.fields.category,
				gender:   tt.fields.gender,
			}
			if got := s.getCategory(); got != tt.want {
				t.Errorf("getCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShoeGetGender(t *testing.T) {
	type fields struct {
		size     int64
		price    float32
		category string
		gender   string
	}
	var tests []struct {
		name   string
		fields fields
		want   string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &shoe{
				size:     tt.fields.size,
				price:    tt.fields.price,
				category: tt.fields.category,
				gender:   tt.fields.gender,
			}
			if got := s.getGender(); got != tt.want {
				t.Errorf("getGender() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShoeGetPrice(t *testing.T) {
	type fields struct {
		size     int64
		price    float32
		category string
		gender   string
	}
	var tests []struct {
		name   string
		fields fields
		want   float32
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &shoe{
				size:     tt.fields.size,
				price:    tt.fields.price,
				category: tt.fields.category,
				gender:   tt.fields.gender,
			}
			if got := s.getPrice(); got != tt.want {
				t.Errorf("getPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShoeSetCategory(t *testing.T) {
	type fields struct {
		size     int64
		price    float32
		category string
		gender   string
	}
	type args struct {
		name string
	}
	var tests []struct {
		name   string
		fields fields
		args   args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &shoe{
				size:     tt.fields.size,
				price:    tt.fields.price,
				category: tt.fields.category,
				gender:   tt.fields.gender,
			}
			s.setCategory(tt.args.name)
		})
	}
}

func TestShoeSetGender(t *testing.T) {
	type fields struct {
		size     int64
		price    float32
		category string
		gender   string
	}
	type args struct {
		gender string
	}
	var tests []struct {
		name   string
		fields fields
		args   args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &shoe{
				size:     tt.fields.size,
				price:    tt.fields.price,
				category: tt.fields.category,
				gender:   tt.fields.gender,
			}
			s.setGender(tt.args.gender)
		})
	}
}

func TestShoeSetPrice(t *testing.T) {
	type fields struct {
		size     int64
		price    float32
		category string
		gender   string
	}
	type args struct {
		discount float32
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   float32
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &shoe{
				size:     tt.fields.size,
				price:    tt.fields.price,
				category: tt.fields.category,
				gender:   tt.fields.gender,
			}
			if got := s.setPrice(tt.args.discount); got != tt.want {
				t.Errorf("setPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
