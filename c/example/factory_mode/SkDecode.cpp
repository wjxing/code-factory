#include "SkDecode.h"
#include "SkRegistory.h"

using namespace skia;

typedef SkRegistory<SkDecode*, SkStream*> DecodeReg;

template DecodeReg* SkRegistory<SkDecode*, SkStream*>::gHead;

SkDecode* SkDecode::factory(SkStream* stream) {
    SkDecode* decode = NULL;
    const DecodeReg* cur = DecodeReg::head();
    while (cur) {
        decode = cur->factory()(stream);
        if (decode)
            return decode;
        cur = cur->next();
    }
    return NULL;
}
