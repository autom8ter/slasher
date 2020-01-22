version := 0.0.1

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

patch:
	go generate
	bumpversion patch --allow-dirty

tag:
	git tag v$(version)

push:
	git push origin master
	git push origin v$(version)