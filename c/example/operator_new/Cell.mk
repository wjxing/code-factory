LOCAL_PATH := $(call my-dir)
include $(CLEAR_VARS)

LOCAL_SRC := test.c
LOCAL_MODULE := $(dir $(LOCAL_PATH))_example
LOCAL_MODULE_TAGS := test

include $(BUILD_EXECUTABLE)
