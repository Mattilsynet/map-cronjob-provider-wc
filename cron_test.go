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
		{
			"CronExpression 5 1 * is 65 seconds",
			args{"5 1 *"},
			65,
			false,
		},
		{
			"CronExpression 5 1 1 is 3665 seconds",
			args{"5 1 1"},
			3665,
			false,
		},
		{
			"CronExpression * * 2 is 7200 seconds",
			args{"* * 2"},
			7200,
			false,
		},
		{
			"CronExpression /5 * * is not allowed",
			args{"/5 * *"},
			-1,
			true,
		},
		{
			"CronExpression 5 A * is not allowed",
			args{"5 A *"},
			-1,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cronParts, err := ParseCronExpression(tt.args.cronExpr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCronExpressionInSeconds() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else {
				seconds := ConvertCronPartsToSeconds(cronParts)

				if seconds != tt.want {
					t.Errorf("GetCronExpressionInSeconds() = %v, want %v", seconds, tt.want)
				}
			}
		})
	}
}
