package main

import (
	"testing"
)

func Test_passport_validByr(t *testing.T) {
	type fields struct {
		byr string
		iyr string
		eyr string
		hgt string
		hcl string
		ecl string
		pid string
		cid string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "byr valid:   2002",
			fields: fields{
				byr: "2002",
			},
			want: true,
		},
		{
			name: "byr invalid:   2003",
			fields: fields{
				byr: "2003",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &passport{
				byr: tt.fields.byr,
				iyr: tt.fields.iyr,
				eyr: tt.fields.eyr,
				hgt: tt.fields.hgt,
				hcl: tt.fields.hcl,
				ecl: tt.fields.ecl,
				pid: tt.fields.pid,
				cid: tt.fields.cid,
			}
			if got := p.validByr(); got != tt.want {
				t.Errorf("passport.validByr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_passport_validIyr(t *testing.T) {
	type fields struct {
		byr string
		iyr string
		eyr string
		hgt string
		hcl string
		ecl string
		pid string
		cid string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "iyr valid:   2010",
			fields: fields{
				iyr: "2010",
			},
			want: true,
		},
		{
			name: "iyr invalid:   2009",
			fields: fields{
				iyr: "2009",
			},
			want: false,
		},
		{
			name: "iyr valid:   2020",
			fields: fields{
				iyr: "2020",
			},
			want: true,
		},
		{
			name: "iyr invalid:   2021",
			fields: fields{
				iyr: "2021",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &passport{
				byr: tt.fields.byr,
				iyr: tt.fields.iyr,
				eyr: tt.fields.eyr,
				hgt: tt.fields.hgt,
				hcl: tt.fields.hcl,
				ecl: tt.fields.ecl,
				pid: tt.fields.pid,
				cid: tt.fields.cid,
			}
			if got := p.validIyr(); got != tt.want {
				t.Errorf("passport.validIyr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_passport_validHgt(t *testing.T) {
	type fields struct {
		byr string
		iyr string
		eyr string
		hgt string
		hcl string
		ecl string
		pid string
		cid string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "hgt valid:   60in",
			fields: fields{
				hgt: "60in",
			},
			want: true,
		},
		{
			name: "hgt valid:   190cm",
			fields: fields{
				hgt: "190cm",
			},
			want: true,
		},
		{
			name: "hgt invalid:   190in",
			fields: fields{
				hgt: "190in",
			},
			want: false,
		},
		{
			name: "hgt invalid:   190",
			fields: fields{
				hgt: "190",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &passport{
				byr: tt.fields.byr,
				iyr: tt.fields.iyr,
				eyr: tt.fields.eyr,
				hgt: tt.fields.hgt,
				hcl: tt.fields.hcl,
				ecl: tt.fields.ecl,
				pid: tt.fields.pid,
				cid: tt.fields.cid,
			}
			if got := p.validHgt(); got != tt.want {
				t.Errorf("passport.validHgt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_passport_validHcl(t *testing.T) {
	type fields struct {
		byr string
		iyr string
		eyr string
		hgt string
		hcl string
		ecl string
		pid string
		cid string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "hcl valid:   #123abc",
			fields: fields{
				hcl: "#123abc",
			},
			want: true,
		},
		{
			name: "hcl invalid:   #123abz",
			fields: fields{
				hcl: "#123abz",
			},
			want: false,
		},
		{
			name: "hcl invalid:   123abc",
			fields: fields{
				hcl: "123abc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &passport{
				byr: tt.fields.byr,
				iyr: tt.fields.iyr,
				eyr: tt.fields.eyr,
				hgt: tt.fields.hgt,
				hcl: tt.fields.hcl,
				ecl: tt.fields.ecl,
				pid: tt.fields.pid,
				cid: tt.fields.cid,
			}
			if got := p.validHcl(); got != tt.want {
				t.Errorf("passport.validHcl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_passport_validEyr(t *testing.T) {
	type fields struct {
		byr string
		iyr string
		eyr string
		hgt string
		hcl string
		ecl string
		pid string
		cid string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &passport{
				byr: tt.fields.byr,
				iyr: tt.fields.iyr,
				eyr: tt.fields.eyr,
				hgt: tt.fields.hgt,
				hcl: tt.fields.hcl,
				ecl: tt.fields.ecl,
				pid: tt.fields.pid,
				cid: tt.fields.cid,
			}
			if got := p.validEyr(); got != tt.want {
				t.Errorf("passport.validEyr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_passport_validPid(t *testing.T) {
	type fields struct {
		byr string
		iyr string
		eyr string
		hgt string
		hcl string
		ecl string
		pid string
		cid string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "pid valid:   000000001",
			fields: fields{
				pid: "000000001",
			},
			want: true,
		},
		{
			name: "pid invalid:   0123456789",
			fields: fields{
				pid: "0123456789",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &passport{
				byr: tt.fields.byr,
				iyr: tt.fields.iyr,
				eyr: tt.fields.eyr,
				hgt: tt.fields.hgt,
				hcl: tt.fields.hcl,
				ecl: tt.fields.ecl,
				pid: tt.fields.pid,
				cid: tt.fields.cid,
			}
			if got := p.validPid(); got != tt.want {
				t.Errorf("passport.validPid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_passport_validEcl(t *testing.T) {
	type fields struct {
		byr string
		iyr string
		eyr string
		hgt string
		hcl string
		ecl string
		pid string
		cid string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &passport{
				byr: tt.fields.byr,
				iyr: tt.fields.iyr,
				eyr: tt.fields.eyr,
				hgt: tt.fields.hgt,
				hcl: tt.fields.hcl,
				ecl: tt.fields.ecl,
				pid: tt.fields.pid,
				cid: tt.fields.cid,
			}
			if got := p.validEcl(); got != tt.want {
				t.Errorf("passport.validEcl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
