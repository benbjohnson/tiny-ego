default:

clean:
	find . -type f | grep .ego.go | xargs rm

generate: http/template assets

run: generate
	goo install ./cmd/tiny-ego
	tiny-ego


http/template: http/component http/template/*.ego
	ego ./http/template

http/component: http/component/*.ego
	ego ./http/component


assets: assets/styles/main.css
	go generate ./assets

assets/styles/main.css: assets/styles/main.scss assets/styles/_component.scss assets/styles/_template.scss assets/styles/*.scss
	sass $< > $@

assets/styles/_template.scss: http/template/*.ego
	find http/template -type f | grep ".ego\$$" | xargs fslice -start "/*SCSS" -end "*/" -header "// \$$FILENAME" > $@

assets/styles/_component.scss: http/component/*.ego
	find http/component -type f | grep ".ego\$$" | xargs fslice -start "/*SCSS" -end "*/" -header "// \$$FILENAME" > $@

.PHONY: http/template http/component