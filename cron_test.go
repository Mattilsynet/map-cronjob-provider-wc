package main

import "testing"

func TestGetCronExpressionInSeconds(t *testing.T) {
	type args struct {
		cronExpr string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"CronExpression 5 * * is 5 seconds",
			args{"5 * *"},
			5,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCronExpressionInSeconds(tt.args.cronExpr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCronExpressionInSeconds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetCronExpressionInSeconds() = %v, want %v", got, tt.want)
			}
		})
	}
}
