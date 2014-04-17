#include <iostream>
#include <string.h>
#include <SkDecode.h>
#include "SkGifDecode.h"

using namespace skia;

void SkGifDecode::decode() {
    std::cout << "SkGifDecode" << std::endl;
}

#include <SkRegistory.h>
static SkDecode* factory(SkStream* stream) {
    if (strncmp(stream->read(), "gif", 3) == 0)
        return new SkGifDecode();
    return NULL;
}

static SkRegistory<SkDecode*, SkStream*> gReg(factory);
