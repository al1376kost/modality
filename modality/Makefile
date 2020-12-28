.PHONY: build
build:
	cp -urf /mnt/f/golang/projects/src/modality/cmd/ ./
	cp -urf /mnt/f/golang/projects/src/modality/internal/ ./
	go build -v ./cmd/modal/

.PHONY: build_static
build_static:
	cp /mnt/f/golang/projects/src/modality/mod_frontend/package.json ./mod_frontend/package.json
	cp /mnt/f/golang/projects/src/modality/mod_frontend/vue.config.js ./mod_frontend/vue.config.js
	cp -urf /mnt/f/golang/projects/src/modality/mod_frontend/public ./mod_frontend
	cp -urf /mnt/f/golang/projects/src/modality/mod_frontend/src ./mod_frontend
	cd ./mod_frontend; npm run build;

.PHONY: run
run:
	./modal