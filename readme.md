<h1>Bookings and Reservations</h1>

<h2>Notes</h2>
<h3>Commands</h3>

* ``` go run .\cmd\web\. ``` - to main
* ``` go run . ``` - to run multiple go files at the same time
* ``` go test -v .\cmd\web\. ``` - run tests in main
* ``` go test -cover .\cmd\web\. ``` - run tests coverage in main
* ``` go test -coverprofile=coverage.out && go tool cover -html=coverage.out .\cmd\web\. ``` - detailed coverage for html - might not work as it is on windows
``` go test -coverprofile=overage.out""coverage.out" ```
``` go tool cover -html="coverage.out" ```