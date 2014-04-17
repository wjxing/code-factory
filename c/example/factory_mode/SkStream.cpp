#include "SkStream.h"

using namespace skia;

SkStream::SkStream(const char* stream)
    : mpStream(stream) {
}

SkStream::~SkStream() {
}

const char* SkStream::read() {
    return mpStream;
}
