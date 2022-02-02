# Mock with `gomock`
> GoMock is a mocking framework for the Go programming language.   

----  

## Topic:
- [] make testable code
- [] build contracts through Interfaces
- [] generate mocks
- [] write tests


## Generate mocks
```bash
mockgen -destination=$(pwd)/mocks/mock_presenter.go --package=mocks -source=$(pwd)/presenter/presenter.go
mockgen -destination=$(pwd)/mocks/mock_converter.go --package=mocks -source=$(pwd)/converter/converter.go
mockgen -destination=$(pwd)/mocks/mock_calculator.go --package=mocks -source=$(pwd)/calculator/calculator.go
```

----
### Links:
- https://github.com/golang/mock

----
### Good reading:
- SOLID [principle](https://gogiggersin.wordpress.com/2020/01/15/solid-design-principle-in-go/)
- (Basic GO)[https://go.dev/tour/list] knowlege