.PHONY : all
all :

#Test
#preffix := *.txt *.pdf */*.txt */*.pdf
#files := $(wildcard $(addprefix $(SYNC_SRC)/, $(preffix)))
#$(error $(files))
#Test

ifeq ($(strip $(SYNC_SRC)),)
	$(error The SYNC_SRC is NULL.)
else
ifeq ($(wildcard $(strip $(SYNC_SRC))),)
	$(error The SYNC_SRC is not exsit.)
endif
endif

ifeq ($(strip $(SYNC_DES)),)
	$(error The SYNC_DES is NULL.)
else
ifeq ($(wildcard $(strip $(SYNC_DES))),)
	$(error The SYNC_DES is not exsit.)
endif
endif

ifeq ($(strip $(shell which md5sum 2>/dev/null)),)
	$(error Miss md5sum tools.)
endif

empty :=
space := $(empty) $(empty)
sync_files_prefix := "*.chm" "*.pdf" "*.txt"
find_params := -type f $(subst $(space)-name, -o -name,$(addprefix -name , $(sync_files_prefix)))

sync_src_files := $(shell find $(SYNC_SRC) $(find_params))
sync_des_files := $(shell find $(SYNC_DES) $(find_params))


$(foreach des_file, $(sync_des_files), \
	$(eval md5sum_command := md5sum -b $(des_file)) \
	$(eval des_file_sum := $(strip $(word 1,$(shell $(md5sum_command))))) \
	$(if $(sum.$(des_file_sum)), \
		$(eval $(warning The file sum is repetition, the origin file is $(sum.$(des_file_sum)) and this file is $(des_file))), \
		$(eval sum.$(des_file_sum) := $(des_file)) \
	) \
)

$(foreach src_file, $(sync_src_files), \
	$(eval md5sum_command := md5sum -b $(src_file)) \
	$(eval src_file_sum := $(strip $(word 1,$(shell $(md5sum_command))))) \
	$(if $(sum.$(src_file_sum)), \
		$(info Skip file $(src_file) which is be in $(sum.$(src_file_sum))), \
		$(eval rel_path := $(subst $(SYNC_SRC)/,,$(dir $(src_file)))) \
		$(eval des_file := $(subst $(SYNC_SRC)/,$(SYNC_DES)/,$(src_file))) \
		$(if $(wildcard $(des_file)), \
			$(info TODO : Skip file $(src_file), cause there is already the same name file in $(des_file)), \
			$(info mkdir $(SYNC_DES)/$(rel_path)) \
			$(shell mkdir -p $(SYNC_DES)/$(rel_path) 2>/dev/null) \
			$(info cp $(src_file) $(des_file)) \
			$(shell cp $(src_file) $(des_file)) \
		) \
	) \
)
