# go testでfailしても最後まで実行が終わるように、`make -i test`で実行すること
.PHONY: test
test:
	$(eval when := $(shell date +%Y%m%d_%H%M%S))
	$(eval testLogPath := logs/test/$(when))
	mkdir -p $(testLogPath)
	go test -v -count=1 -cover -json ./... > $(testLogPath)/log-jsons.txt
	go test -v -count=1 -coverprofile=$(testLogPath)/cover.out ./... | tee $(testLogPath)/log.txt