LOCAL_PATH := $(call my-dir)
include $(CLEAR_VARS)

LOCAL_SRC := test.c
LOCAL_MODULE := show_macro_example
LOCAL_MODULE_TAGS := test

include $(BUILD_SHARED_LIBRARY)
