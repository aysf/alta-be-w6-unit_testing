# alta-be-w6-unit_testing

## Testing with output format html 
##### `go test ./controllers -coverprofile=cover.out && go tool cover -html=cover.out`

## Testing with percentage report
##### `go test -v -coverpkg=./controllers -coverprofile=profile.cov ./controllers && go tool cover -func profile.cov`
