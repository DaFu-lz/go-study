package dao

import (
	"Demo/goweb/model"
	"reflect"
	"testing"
)

func TestAddPayRecode(t *testing.T) {
	type args struct {
		data model.ConsumptionRecode
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddPayRecode(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("AddPayRecode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChangeBalance(t *testing.T) {
	type args struct {
		pay_type_id int
		amount      float64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestCheckLogin(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckLogin(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckLogin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckRecode(t *testing.T) {
	type args struct {
		trade_time string
		price      float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckRecode(tt.args.trade_time, tt.args.price); got != tt.want {
				t.Errorf("CheckRecode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckUsername(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckUsername() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPayMethod(t *testing.T) {
	tests := []struct {
		name  string
		want  error
		want1 []model.PayType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetPayMethod()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPayMethod() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetPayMethod() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPayMonth(t *testing.T) {
	tests := []struct {
		name  string
		want  error
		want1 []model.ConsumptionRecode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := PayMonth()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PayMonth() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PayMonth() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPayToday(t *testing.T) {
	tests := []struct {
		name  string
		want  error
		want1 []model.ConsumptionRecode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := PayToday()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PayToday() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PayToday() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPayWeek(t *testing.T) {
	tests := []struct {
		name  string
		want  error
		want1 []model.ConsumptionRecode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := PayWeek()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PayWeek() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PayWeek() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestQueryBalance(t *testing.T) {
	tests := []struct {
		name  string
		want  []model.PayType
		want1 float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := QueryBalance()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryBalance() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("QueryBalance() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSaveUser(t *testing.T) {
	type args struct {
		username string
		password string
		email    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveUser(tt.args.username, tt.args.password, tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("SaveUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
