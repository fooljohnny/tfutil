package logistic_test

import (
	"path"
	"testing"

	"github.com/flyingyizi/tfutil/df"
	. "github.com/flyingyizi/tfutil/supervisedLearning/logistic"

	"gonum.org/v1/gonum/mat"
)

func TestComputeCost(t *testing.T) {
	filename := "ex2data1.txt"
	//load traning data
	X, Y := func() (x *mat.Dense, y *mat.VecDense) {
		orig := mat.NewDense(df.CsvToArray(path.Join("testdata", filename)))
		or, oc := orig.Dims()
		// assign Y
		var Y mat.VecDense
		Y.CloneVec(orig.ColView(oc - 1))
		// assign X
		ones := mat.NewVecDense(or, df.Ones(or))
		X := df.HorizJoinDense(ones, orig.Slice(0, or, 0, oc-1)) //X shape is: 'or by (oc)'
		return X, &Y
	}()

	_, rc := X.Dims()
	theta1 := mat.NewVecDense(rc, nil)

	type args struct {
		X     *mat.Dense
		y     *mat.VecDense
		theta mat.Vector
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{name: "11",
			args: args{X: X, y: Y, theta: theta1},
			want: 0.6931471805599458},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeCost(LogicHyphothesis, tt.args.X, tt.args.y, tt.args.theta); got != tt.want {
				t.Errorf("ComputeCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
