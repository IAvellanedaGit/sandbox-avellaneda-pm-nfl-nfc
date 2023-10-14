.PHONY: flatbuffers
flatbuffers: flatbuffers_go flatbuffers_ts

.PHONY: flatbuffers_deps
flatbuffers_deps: flatbuffers_deps_macos

flatbuffers_deps_macos:
	rm -rf /tmp/flatbuffers && \
	rm -rf /tmp/cmake && \
	mkdir -p /tmp/cmake && \
	mkdir -p /tmp/flatbuffers && \
	cd /tmp && \
	cd /tmp/cmake && \
	curl -L -O https://github.com/Kitware/CMake/releases/download/v3.23.1/cmake-3.23.1-macos-universal.tar.gz && \
	tar -xzvf ./cmake-3.23.1-macos-universal.tar.gz && \
	cd /tmp && \
	git clone https://github.com/google/flatbuffers.git -b master && \
    cd flatbuffers && \
	git checkout v2.0.6 && \
    ../cmake/cmake-3.23.1-macos-universal/CMake.app/Contents/bin/cmake -G "Unix Makefiles" && \
    make && \
	sudo make install

.PHONY: flatbuffers_clean
flatbuffers_clean:
	rm -rf /tmp/flatbuffers && \
	rm -rf /tmp/cmake

.PHONY: flatbuffers_go
flatbuffers_go:
	rm -rf output/flatbuffers/go && mkdir -p output/flatbuffers/go
	flatc --go --gen-mutable -o ./output/flatbuffers/go flatbuffers/model.fbs
	rm -rf service/model && mkdir -p service/model
	cp output/flatbuffers/go/model/* service/model/

.PHONY: flatbuffers_ts
flatbuffers_ts:
	rm -rf output/flatbuffers/ts && mkdir -p output/flatbuffers/ts
	flatc --ts --gen-mutable -o ./output/flatbuffers/ts flatbuffers/model.fbs
	rm -rf react-native/ && mkdir -p react-native/
	cp -r output/flatbuffers/ts/model/* react-native/

