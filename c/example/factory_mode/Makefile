LOCAL_MODULE        := factory_test
LOCAL_SRC_FILES     := $(wildcard *.cpp)
LOCAL_SRC_FILES     += $(wildcard plugin/*.cpp)
LOCAL_CFLAGS        := -g -Wall -includeConfig.h
LOCAL_C_INCLUDES    := -I.
CC                  := g++
RM                  := rm

all : $(LOCAL_MODULE)

$(LOCAL_MODULE) : $(LOCAL_SRC_FILES)
	$(CC) $(LOCAL_C_INCLUDES) $(LOCAL_CFLAGS) $^ -o $@

.PHONY : clean

clean :
	$(RM) $(LOCAL_MODULE)

