go_compile := $(HOME)/bin/go/pkg/tool/linux_amd64/compile
go_link := $(HOME)/bin/go/pkg/tool/linux_amd64/link
my_path := $(HOME)/Shared/Workspace/Golang/src/sample/less1/use_compile
parse_go := $(my_path)/bp-parse/parse.go
parse_obj := $(my_path)/obj/bp/build/parse.a
minibp_obj := $(my_path)/obj/minibp/minibp.a
parse_pkg := bp/build/parse
minibp_pkg := minibp

.PHONY : all clean

all : minibp

minibp : $(minibp_obj)
	$(go_link) -o $@ -L $(my_path)/obj $^

$(minibp_obj) : minibp.go | $(parse_obj)
	$(go_compile) -o $@ -p $(minibp_pkg) -complete -I $(my_path)/obj -pack $^

$(parse_obj) : $(parse_go)
	$(go_compile) -o $@ -p $(pkg_path) -complete -pack $^

clean :
	rm -rf $(parse_obj) $(minibp_obj) minibp
