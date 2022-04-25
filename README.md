# go_mid_to_senior
It is a personal learn golang lab. All basic knowledge has been moved to https://github.com/ramseyjiang/go_junior_to_mid. 

The whole repo is following golangci-lint run.

Run the following command in your command line. If something is not the same with golangci-lint, it will have warnings and errors.

% golangci-lint run

All pkgs in this repo, if you want to check whether it corrects. You can change a pkg name into main, and change the func Trigger() to main(). After that, you also need to change the relationship test.go pkg name the same to main. After all above, you can execute it in your command line. For example.

% go run filename.go

By the way, almost all unittests I did fake tests. If you wanna to do real, please just change the test data, they will work very well.