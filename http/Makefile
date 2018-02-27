default: template component assets

template: component template/*.ego
	ego ./template

component: component/*.ego
	ego ./component

assets: assets/styles/main.css
	go generate ./assets

assets/styles/main.css: assets/styles/main.scss assets/styles/**/*.scss
	sass $< > $@
