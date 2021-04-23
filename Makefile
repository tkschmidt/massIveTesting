
.PHONY: generate build

MKDIR_P = mkdir -p

generate_build_folder:
	@${MKDIR_P} builds
	@echo "[OK] Build folder exists"

build: generate_build_folder
	@env GOOS=windows GOARCH=amd64 go build -ldflags "-extldflags '-static' -X main.GitCommit=$CI_COMMIT_SHA" -o builds/testMassIVE.exe
	@echo "[OK] Windows build was created!"
	@env GOOS=darwin GOARCH=amd64 go build -ldflags "-extldflags '-static' -X main.GitCommit=$CI_COMMIT_SHA" -o builds/testMassIVE_mac
	@echo "[OK] Mac build was created!"
	@env OOS=linux GOARCH=amd64 go build -ldflags "-extldflags '-static' -X main.GitCommit=$CI_COMMIT_SHA" -o builds/testMassIVE
	@echo "[OK] Linux build was created!"