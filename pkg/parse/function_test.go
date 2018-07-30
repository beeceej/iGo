package parse

import (
	"testing"
)

func Test_isFunction(t *testing.T) {
	var (
		expectSuccess func(error)
		expectFailure func(error)
		err           error
	)
	expectSuccess = func(err error) {
		if err != nil {
			t.Fail()
		}
	}
	expectFailure = func(err error) {
		if err == nil {
			t.Fail()
		}
	}

	tests := []func(){
		func() {
			_, err = NewFunction(
				`func a(){
	
			}`,
			)
			expectSuccess(err)
		},
		func() {
			_, err = NewFunction(
				`func a()
		
				}`,
			)
			expectFailure(err)
		},
		func() {
			_, err = NewFunction(
				`func a(){}`,
			)
			expectSuccess(err)
		},
		func() {
			_, err = NewFunction(
				`func a(a ...string){
					fmt.Println(a)
				}`,
			)
			expectSuccess(err)
		},
		func() {
			_, err = NewFunction(
				`func a(a ...string) (err error){
					return err
				}`,
			)
			expectSuccess(err)
		},
		func() {
			_, err = NewFunction(
				`func abcdef(a, b,c,d string) error{
					return nil
				}`,
			)
			expectSuccess(err)
		},
	}

	run(tests)
}

func Test_matchIdentifier(t *testing.T) {
	var (
		f *Function
	)
	tests := []func(){
		func() {
			f, _ = NewFunction(
				`func b() (err error, key string) {
					fmt.Println("str")
				}
				`,
			)

			if f.matchIdentifier() != "b" {
				t.Fail()
			}
		},
		func() {
			f, _ = NewFunction(
				`func a(s ...string) (err error) {
					return err
				}
				`,
			)
			if f.matchIdentifier() != "a" {
				t.Fail()
			}
		},
		func() {
			f, _ = NewFunction(
				`func c (a,b int) int {
					return a+b
				}`,
			)
			if f.matchIdentifier() != "c" {
				t.Fail()
			}
		},
		func() {
			f, _ = NewFunction(
				`func b () {}`,
			)
			if f.matchIdentifier() != "b" {
				t.Fail()
			}
		},
	}
	run(tests)
}

func run(funcs []func()) {
	for _, f := range funcs {
		f()
	}
}
