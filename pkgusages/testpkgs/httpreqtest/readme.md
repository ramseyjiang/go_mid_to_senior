1. test current pkg test coverage % go test -v ./... -coverpkg ./...

% go test -v ./... -coverpkg ./...                        
=== RUN TestRealCallSuccess --- PASS: TestRealCallSuccess (0.58s)
=== RUN TestRealCallFail --- PASS: TestRealCallFail (0.04s)
=== RUN TestFakeCallSuccess --- PASS: TestFakeCallSuccess (0.00s)
=== RUN TestFakeCallFail --- PASS: TestFakeCallFail (0.00s)
=== RUN TestApp === RUN TestApp/GET_endpoint_to_get_a_sum === RUN TestApp/POST_endpoint_to_multiply,_wrong_header ===
RUN TestApp/POST_endpoint_to_multiply --- PASS: TestApp (0.00s)
--- PASS: TestApp/GET_endpoint_to_get_a_sum (0.00s)
--- PASS: TestApp/POST_endpoint_to_multiply,_wrong_header (0.00s)
--- PASS: TestApp/POST_endpoint_to_multiply (0.00s)
PASS coverage: 76.3% of statements in ./... ok github.com/ramseyjiang/go_mid_to_senior/pkgusages/testpkgs/httpreqtest
1.045s coverage: 76.3% of statements in ./...

2. test current pkg test coverage and output to a "cover.out" file % go test -v ./... -coverpkg ./... -coverprofile
   cover.out

% go test -v ./... -coverpkg ./... -coverprofile cover.out === RUN TestRealCallSuccess --- PASS: TestRealCallSuccess (
0.51s)
=== RUN TestRealCallFail --- PASS: TestRealCallFail (0.04s)
=== RUN TestFakeCallSuccess --- PASS: TestFakeCallSuccess (0.00s)
=== RUN TestFakeCallFail --- PASS: TestFakeCallFail (0.00s)
=== RUN TestApp === RUN TestApp/GET_endpoint_to_get_a_sum === RUN TestApp/POST_endpoint_to_multiply,_wrong_header ===
RUN TestApp/POST_endpoint_to_multiply --- PASS: TestApp (0.00s)
--- PASS: TestApp/GET_endpoint_to_get_a_sum (0.00s)
--- PASS: TestApp/POST_endpoint_to_multiply,_wrong_header (0.00s)
--- PASS: TestApp/POST_endpoint_to_multiply (0.00s)
PASS coverage: 76.3% of statements in ./... ok github.com/ramseyjiang/go_mid_to_senior/pkgusages/testpkgs/httpreqtest
0.999s coverage: 76.3% of statements in ./...